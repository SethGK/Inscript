// Package compiler compiles the Inscript AST into bytecode.
package compiler

import (
	"encoding/binary" // Needed for binary encoding/decoding in patchJump
	"fmt"

	// Needed for Fprintf to stderr
	"github.com/SethGK/Inscript/internal/ast"
	// code.go and symboltable.go are in the same package (compiler),
	// so they do not need to be imported explicitly.
	"github.com/SethGK/Inscript/internal/types"
)

// Compiler translates an AST into bytecode.
type Compiler struct {
	// The bytecode being compiled for the current compilation unit (program or function)
	instructions Instructions  // Use Instructions type from code.go
	constants    []types.Value // Constant pool for the current compilation unit - Referring to types.Value

	// Global symbol table (unique for the entire program)
	// This is shared across all function compilers.
	globals *SymbolTable // Referring to SymbolTable from symboltable.go
	// Removed numGlobals field - the count is now in c.globals.NumDefinitions()

	// Scope management
	symbolTableStack []*SymbolTable // Referring to SymbolTable from symboltable.go
	currentScope     *SymbolTable   // Referring to SymbolTable from symboltable.go

	// Bytecode buffer for the current compilation unit
	currentBytecode *Bytecode // Use Bytecode type from code.go

	// Flag to indicate if a return statement has been compiled in the current function body
	returned bool
}

// New creates a new top-level Compiler instance for a program.
func New() *Compiler {
	globalTable := NewSymbolTable() // Referring to NewSymbolTable from symboltable.go - isFunctionScope is false

	// The initial currentBytecode is for the program itself.
	programBytecode := NewBytecode() // Referring to NewBytecode from code.go

	c := &Compiler{
		instructions: programBytecode.Instructions, // Start with the program's instruction slice ([]byte)
		constants:    programBytecode.Constants,    // Start with the program's constant slice ([]types.Value)
		globals:      globalTable,                  // The global table
		// numGlobals field removed
		symbolTableStack: []*SymbolTable{globalTable}, // Start with global scope on stack
		currentScope:     globalTable,
		currentBytecode:  programBytecode, // Compiling the main program initially
		returned:         false,           // Main program doesn't have a return flag initially
	}

	// Pre-define builtins if any (e.g., "print" could be a builtin function)
	// For now, let's assume print is handled by OpPrint directly and not a callable builtin.
	// If you add builtins, define them in the global table here:
	// c.DefineGlobal("print", Builtin) // Example - Referring to Builtin from symboltable.go

	return c
}

// NumGlobalVariables returns the total number of global variables defined.
// This now returns the number of definitions in the global symbol table.
func (c *Compiler) NumGlobalVariables() int {
	return c.globals.NumDefinitions() // Use the method from the global symbol table
}

// newFunctionCompiler creates a compiler instance for a function body.
// It inherits the global symbol table but has its own bytecode buffer
// and a symbol table nested within the parent compiler's current scope.
func (c *Compiler) newFunctionCompiler() *Compiler {
	// Create a new symbol table enclosed in the parent's *current* scope.
	// This captures variables from outer scopes (for closures).
	// Pass true to indicate this is a function scope.
	funcScope := NewEnclosedSymbolTable(c.currentScope, true) // Referring to NewEnclosedSymbolTable from symboltable.go

	// Each function has its own bytecode and constant pool.
	// Functions share the *main* compiler's constant pool.
	// The function compiler's constants slice is a reference to the main compiler's constants.
	funcComp := &Compiler{
		instructions:     make(Instructions, 0),     // Start with an empty instruction slice for the function body
		constants:        c.constants,               // Reference the main compiler's constant pool
		globals:          c.globals,                 // Share the global table
		symbolTableStack: []*SymbolTable{funcScope}, // Stack starts with function scope
		currentScope:     funcScope,
		returned:         false, // Initialize returned flag for the function compiler
		// currentBytecode is not used in the function compiler instance itself,
		// its bytecode is built into funcComp.instructions and funcComp.constants
	}
	return funcComp
}

// EnterScope pushes a new nested scope onto the stack.
// isFunctionScope is true if entering a function body's scope.
// This method is now primarily used for block scopes within the main program.
func (c *Compiler) EnterScope(isFunctionScope bool) {
	// We only enter non-function scopes here (block scopes in main program).
	// Function scopes are handled by newFunctionCompiler.
	// *** FIX START: Access isFunctionScope as a lowercase field ***
	if isFunctionScope && !c.currentScope.isFunctionScope {
		// *** FIX END ***
		// Only allow entering a function scope if the current scope is not already one.
		// This prevents accidentally nesting function scopes using EnterScope.
		panic("EnterScope(true) should only be called when entering a function's top-level scope")
	}
	// *** FIX START: Access isFunctionScope as a lowercase field ***
	if isFunctionScope && c.currentScope.isFunctionScope {
		// *** FIX END ***
		// If already in a function scope, entering another function scope is nested,
		// which is handled by newFunctionCompiler, not EnterScope.
		panic("EnterScope(true) should not be called on a compiler instance already in a function scope")
	}

	// Block scope is enclosed in the current active scope.
	// Pass false to indicate this is NOT a function scope.
	newTable := NewEnclosedSymbolTable(c.currentScope, false) // Referring to NewEnclosedSymbolTable from symboltable.go

	c.symbolTableStack = append(c.symbolTableStack, newTable)
	c.currentScope = newTable
	// c.currentScope.numDefinitions is already initialized to 0 by NewEnclosedSymbolTable
}

// LeaveScope pops the current scope from the stack.
func (c *Compiler) LeaveScope() {
	if len(c.symbolTableStack) <= 1 {
		// Should not pop the global scope from the stack using this method.
		// The global scope is the base of the stack.
		panic("attempting to leave global scope using LeaveScope")
	}
	// Pop the current scope.
	poppedScope := c.currentScope
	c.symbolTableStack = c.symbolTableStack[:len(c.symbolTableStack)-1]
	c.currentScope = c.symbolTableStack[len(c.symbolTableStack)-1]

	// When leaving a scope (especially a function scope), we might need to
	// adjust the number of locals for the parent compilation unit.
	// For block scopes within the main program, the locals are conceptually
	// part of the main program's stack space.
	// For function scopes, the poppedScope.numDefinitions is the number of locals
	// for that function. This is needed when compiling the function definition.
	// This logic will be refined when implementing function compilation.

	// For now, in the main compiler, leaving a block scope doesn't change the
	// main program's total local count (which is effectively 0 as globals are separate).
	// The number of locals for a function is determined when compiling the function body.
	_ = poppedScope // Use the variable to avoid unused warning, will be used later
}

// DefineGlobal defines a symbol in the top-level global symbol table.
// Returns the created Symbol.
func (c *Compiler) DefineGlobal(name string) *Symbol {
	// If it already exists, just return it.
	if sym, ok := c.globals.store[name]; ok {
		return sym
	}
	// Otherwise define a new GLOBAL
	return c.globals.Define(name, Global)
}

// DefineLocal defines a symbol in the current scope's symbol table.
// Returns the created Symbol.
func (c *Compiler) DefineLocal(name string) *Symbol { // Referring to Symbol from symboltable.go
	// Check if the symbol is already defined in the current scope.
	// This prevents re-declaring variables in the same scope (if language rules require this).
	// For now, we allow shadowing outer scopes but not re-declaration in the same scope.
	// *** FIX START: Access store as a lowercase field ***
	_, ok := c.currentScope.store[name]
	// *** FIX END ***
	if ok {
		// Variable already defined in the current scope.
		// Depending on language semantics, you might want to return an error here.
		// For now, let's allow it, and the new definition will shadow the old one in this scope.
		// A stricter language would return fmt.Errorf("variable '%s' already defined in this scope", name)
	}

	// Define a new local symbol in the current scope.
	// The index is the current number of definitions in this scope.
	// *** FIX START: Call NumDefinitions() method and remove increment ***
	symbol := &Symbol{Name: name, Kind: Local, Index: c.currentScope.NumDefinitions()} // Use the method
	c.currentScope.store[name] = symbol                                                // Access store as a lowercase field
	c.currentScope.numDefinitions++                                                    // Manually increment the lowercase field
	// *** FIX END ***
	return symbol
}

// Compile compiles the AST program into bytecode.
// This method is called on the top-level Compiler instance.
func (c *Compiler) Compile(program *ast.Program) (*Bytecode, error) { // Returning *Bytecode from code.go
	// The compiler is already initialized with the global scope and program bytecode.

	// compile all top-level statements
	for _, stmt := range program.Stmts {
		if err := c.compileStatement(stmt); err != nil {
			return nil, err
		}
	}

	// Ensure the current scope is the global scope at the end of program compilation.
	if c.currentScope != c.globals {
		// This indicates a scope mismatch, likely a bug in scope management (Enter/LeaveScope calls).
		return nil, fmt.Errorf("scope mismatch at end of program compilation")
	}

	// Implicit return null at the end of the program execution.
	// This ensures the VM doesn't run off the end of instructions.
	c.emit(OpNull)   // Use OpNull from code.go
	c.emit(OpReturn) // Use OpReturn from code.go

	// Finalize the program's bytecode object.
	// The instructions and constants are already in c.instructions and c.constants.
	c.currentBytecode.Instructions = c.instructions
	c.currentBytecode.Constants = c.constants
	// The number of locals for the main program's bytecode is 0, as variables
	// defined at the top level are globals.
	c.currentBytecode.NumLocals = 0
	c.currentBytecode.NumParameters = 0 // Program has no parameters

	// The number of global variables is the number of definitions in the global symbol table.
	// *** FIX START: Call NumDefinitions() method ***
	c.currentBytecode.NumGlobals = c.globals.NumDefinitions() // Use the method from the global symbol table
	// *** FIX END ***

	return c.currentBytecode, nil
}

// compileStatement emits bytecode for a statement node.
func (c *Compiler) compileStatement(s ast.Statement) error {
	// If we've already compiled a return statement in this function body,
	// subsequent statements are unreachable and should not be compiled.
	if c.returned {
		// --- Debug Print: Skipping unreachable statement ---
		fmt.Printf("DEBUG (Compiler): Skipping unreachable statement type %T after return\n", s)
		// --- End Debug Print ---
		return nil
	}

	// --- Debug Print: Compiling Statement ---
	fmt.Printf("DEBUG (Compiler): Compiling Statement type %T\n", s)
	// --- End Debug Print ---

	switch stmt := s.(type) {
	case *ast.ExprStmt:
		// Compile the expression. It leaves its result on the stack.
		// --- Debug Print: Compiling ExprStmt expression ---
		fmt.Printf("DEBUG (Compiler): Compiling ExprStmt expression type %T\n", stmt.Expr)
		// --- End Debug Print ---
		if err := c.compileExpression(stmt.Expr); err != nil {
			return err
		}
		// Pop the unused result.
		c.emit(OpPop)

	case *ast.AssignStmt:
		// Only identifier targets supported
		ident, ok := stmt.Target.(*ast.Identifier)
		if !ok {
			return fmt.Errorf("unsupported assignment target: %T", stmt.Target)
		}

		// Compile the RHS
		fmt.Printf(
			"DEBUG (Compiler): Compiling AssignStmt value for target '%s' (type %T\n",
			ident.Name, stmt.Value,
		)
		fmt.Printf(
			"DEBUG (Compiler): Before compileExpression for AssignStmt value. Instructions length: %d\n",
			len(c.instructions),
		)
		if err := c.compileExpression(stmt.Value); err != nil {
			return err
		}
		fmt.Printf(
			"DEBUG (Compiler): After compileExpression for AssignStmt value. Instructions length: %d\n",
			len(c.instructions),
		)

		// Resolve or define the symbol
		var sym *Symbol
		if c.currentScope.isFunctionScope {
			// Inside a function: parameters and locals
			sym, _ = c.currentScope.Resolve(ident.Name)
			if sym == nil {
				// Wasn't a parameter or existing local → define a new local
				sym = c.currentScope.DefineLocal(ident.Name)
			}
			c.emit(OpSetLocal, sym.Index)
		} else {
			// Top‑level or main‑program block: globals
			sym, _ = c.globals.Resolve(ident.Name)
			if sym == nil {
				sym = c.DefineGlobal(ident.Name)
			}
			c.emit(OpSetGlobal, sym.Index)
		}

	case *ast.PrintStmt:
		for _, e := range stmt.Exprs {
			fmt.Printf("DEBUG (Compiler): Compiling PrintStmt expression (type %T\n", e)
			fmt.Printf("DEBUG (Compiler): Before compileExpression for PrintStmt expression. Instructions length: %d\n", len(c.instructions))
			if err := c.compileExpression(e); err != nil {
				return err
			}
			fmt.Printf("DEBUG (Compiler): After compileExpression for PrintStmt expression. Instructions length: %d\n", len(c.instructions))
		}
		c.emit(OpPrint, len(stmt.Exprs))

	case *ast.ReturnStmt:
		if stmt.Expr != nil {
			fmt.Printf("DEBUG (Compiler): Compiling ReturnStmt expression (type %T\n", stmt.Expr)
			fmt.Printf("DEBUG (Compiler): Before compileExpression for ReturnStmt expression. Instructions length: %d\n", len(c.instructions))
			if err := c.compileExpression(stmt.Expr); err != nil {
				return err
			}
			fmt.Printf("DEBUG (Compiler): After compileExpression for ReturnStmt expression. Instructions length: %d\n", len(c.instructions))
			c.emit(OpReturnValue)
		} else {
			c.emit(OpNull)
			c.emit(OpReturn)
		}
		c.returned = true

	case *ast.BlockStmt:
		return c.compileBlockStmt(stmt)

	case *ast.IfStmt:
		return c.compileIf(stmt)

	case *ast.WhileStmt:
		return c.compileWhile(stmt)

	case *ast.ForStmt:
		return c.compileFor(stmt)

	default:
		fmt.Printf("DEBUG (Compiler): Unsupported Statement type %T\n", s)
		return fmt.Errorf("unsupported statement type: %T", s)
	}

	return nil
}

// compileExpression emits bytecode for an expression node.
func (c *Compiler) compileExpression(e ast.Expression) error {
	// --- Debug Print: Compiling Expression type %T ---
	fmt.Printf("DEBUG (Compiler): Compiling Expression type %T\n", e)
	// --- End Debug Print ---

	switch expr := e.(type) {
	case *ast.IntegerLiteral:
		c.emitConstant(types.NewInteger(expr.Value)) // Assuming types.NewInteger exists

	case *ast.FloatLiteral:
		c.emitConstant(types.NewFloat(expr.Value)) // Assuming types.NewFloat exists

	case *ast.StringLiteral:
		c.emitConstant(types.NewString(expr.Value)) // Assuming types.NewString exists

	case *ast.BooleanLiteral:
		if expr.Value {
			c.emit(OpTrue)
		} else {
			c.emit(OpFalse)
		}

	case *ast.NilLiteral:
		c.emit(OpNull)

	case *ast.Identifier:
		symbol, ok := c.currentScope.Resolve(expr.Name) // Assuming Resolve method exists
		if !ok {
			return fmt.Errorf("undefined variable '%s'", expr.Name)
		}
		switch symbol.Kind { // Assuming Kind field exists on Symbol
		case Global:
			c.emit(OpGetGlobal, symbol.Index) // Assuming Index field exists on Symbol
		case Local, Parameter: // Assuming Local and Parameter kinds exist
			// *** FIX START: Access isFunctionScope as a lowercase field ***
			if !c.currentScope.isFunctionScope {
				// *** FIX END ***
				return fmt.Errorf("getting local '%s' outside function scope", symbol.Name)
			}
			c.emit(OpGetLocal, symbol.Index) // Assuming Index field exists on Symbol
		case Builtin: // Assuming Builtin kind exists
			c.emit(OpGetGlobal, symbol.Index) // Builtins are stored in globals
		// *** FIX START: Add case for Free symbol kind ***
		case Free: // Referring to Free from symboltable.go
			// Get a free variable. Operand is the index in the function's freeSymbols slice.
			c.emit(OpGetFree, symbol.Index) // Use OpGetFree from code.go
		// *** FIX END ***
		default:
			return fmt.Errorf("unsupported symbol kind for get: %v", symbol.Kind)
		}

	case *ast.BinaryExpr: // Assuming ast.BinaryExpr exists
		// Handle short-circuiting for 'and' and 'or'.
		if expr.Operator == "and" { // Assuming Operator field exists
			// Compile the left side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Left); err != nil { // Assuming Left field exists
				return err
			}
			// If the left side is false (or nil), jump over the right side compilation.
			// The falsy value is already on the stack and is the result of the 'and' if short-circuited.
			jumpFalsePos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy from code.go

			// If the left side was true, pop its value (we only need the right side's value for the result).
			c.emit(OpPop) // Use OpPop from code.go

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil { // Assuming Right field exists
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpFalsePos, len(c.instructions)) // Pass the target position (current length)
			// The result of the expression (either Left's falsy value or Right's value) is now on the stack.

		} else if expr.Operator == "or" { // Assuming Operator field exists
			// Compile the left side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Left); err != nil { // Assuming Left field exists
				return err
			}
			// If the left side is true (or not nil), jump over the right side compilation.
			// The truthy value is already on the stack and is the result of the 'or' if short-circuited.
			jumpTruePos := c.emit(OpJumpTruthy, 0) // Use OpJumpTruthy from code.go

			// If the left side was false, pop its value (we only need the right side's value for the result).
			c.emit(OpPop) // Use OpPop from code.go

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil { // Assuming Right field exists
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpTruePos, len(c.instructions)) // Pass the target position (current length)
			// The result of the expression (either Left's falsy value or Right's value) is now on the stack.

		} else {
			// For other binary operators (+, -, *, /, %, ^, ==, !=, <, >, <=, >=).
			// Compile the left side first. It leaves its result on the stack.
			if err := c.compileExpression(expr.Left); err != nil { // Assuming Left field exists
				return err
			}
			// Compile the right side second. It leaves its result on the stack.
			if err := c.compileExpression(expr.Right); err != nil { // Assuming Right field exists
				return err
			}

			// Emit the opcode corresponding to the operator.
			// The VM will pop the two operands, perform the operation, and push the result.
			switch expr.Operator { // Assuming Operator field exists
			case "+":
				c.emit(OpAdd) // Use OpAdd from code.go
			case "-":
				c.emit(OpSub) // Use OpSub from code.go
			case "*":
				c.emit(OpMul) // Use OpMul from code.go
			case "/":
				c.emit(OpDiv) // Use OpDiv from code.go
			case "%":
				c.emit(OpMod) // Use OpMod from code.go
			case "^":
				c.emit(OpPow) // Use OpPow from code.go
			case "==":
				c.emit(OpEqual) // Use OpEqual from code.go
			case "!=":
				c.emit(OpNotEqual) // Use OpNotEqual from code.go
			case "<":
				c.emit(OpLessThan) // Use OpLessThan from code.go
			case ">":
				c.emit(OpGreaterThan) // Use OpGreaterThan from code.go
			case "<=":
				c.emit(OpLessEqual) // Use OpLessEqual from code.go
			case ">=":
				c.emit(OpGreaterEqual) // Use OpGreaterEqual from code.go
			default:
				// This case should ideally not be reached if the parser/AST builder is correct.
				return fmt.Errorf("unknown binary operator %s", expr.Operator)
			}
		}

	case *ast.UnaryExpr: // Assuming ast.UnaryExpr exists
		// Compile the expression the unary operator applies to. It leaves its result on the stack.
		if err := c.compileExpression(expr.Expr); err != nil { // Assuming Expr field exists
			return err
		}

		// Emit the opcode corresponding to the unary operator.
		// The VM will pop the operand, perform the operation, and push the result.
		switch expr.Operator { // Assuming Operator field exists
		case "-":
			c.emit(OpMinus) // Use OpMinus from code.go
		case "not":
			c.emit(OpNot) // Use OpNot from code.go
		case "+":
			// Unary plus is typically a no-op semantically, but we still compile the expression.
			// No opcode is needed unless you want to enforce type checks at runtime (e.g., enforce the operand is a number).
			// For now, we do nothing after compiling the expression.
		default:
			// This case should ideally not be reached.
			return fmt.Errorf("unknown unary operator %s", expr.Operator)
		}

	case *ast.ListLiteral: // Assuming ast.ListLiteral exists
		// Compile each element of the list. They leave their results on the stack in order.
		count := len(expr.Elements) // Assuming Elements field exists
		for _, el := range expr.Elements {
			if err := c.compileExpression(el); err != nil {
				return err
			}
		}
		// Emit OpArray with the number of elements.
		// The VM will pop the elements from the stack and create a list.
		c.emit(OpArray, count) // Use OpArray from code.go

	case *ast.TableLiteral: // Assuming ast.TableLiteral exists
		// Compile each field (key and value) of the table.
		// Keys and values are pushed onto the stack alternatingly: key1, value1, key2, value2, ...
		count := len(expr.Fields)       // Assuming Fields field exists
		for _, f := range expr.Fields { // Assuming f is a struct with Key and Value fields
			// Table keys are identifiers in the grammar, but typically evaluated to strings or other hashable values.
			// Assuming keys are treated as strings in the bytecode/VM.
			// Add the key string to the constant pool and emit OpConstant.
			c.emitConstant(types.NewString(f.Key)) // Referring to types.NewString (from types package), Assuming Key field exists on f
			// Compile the value expression.
			if err := c.compileExpression(f.Value); err != nil { // Assuming Value field exists on f
				return err
			}
		}
		// Emit OpHash with the number of fields (key-value pairs).
		// The VM will pop 2*count values from the stack (key, value pairs) and create a hash/table.
		c.emit(OpHash, count) // Use OpHash from code.go

	case *ast.IndexExpr: // Assuming ast.IndexExpr exists
		// Compile the primary expression (the aggregate, e.g., list or table). Leaves aggregate on stack.
		if err := c.compileExpression(expr.Primary); err != nil { // Assuming Primary field exists
			return err
		}
		// Compile the index expression (e.g., integer for list, string/value for table). Leaves index on stack.
		if err := c.compileExpression(expr.Index); err != nil { // Assuming Index field exists
			return err
		}
		// Stack top is now [..., aggregate, index].
		// Emit OpIndex. The VM should pop index, pop aggregate, perform the index lookup, and push the result.
		c.emit(OpIndex) // Use OpIndex from code.go

	case *ast.CallExpr: // Assuming ast.CallExpr exists
		// Compile the callee (the function or callable object). Leaves the callable on the stack.
		if err := c.compileExpression(expr.Callee); err != nil { // Assuming Callee field exists
			return err
		}
		// Compile each argument. Arguments are pushed onto the stack in order.
		for _, arg := range expr.Args { // Assuming Args field exists
			if err := c.compileExpression(arg); err != nil {
				return err
			}
		}
		// Emit OpCall with the number of arguments.
		// The VM will pop the arguments and the callable, set up a new frame, and execute the callable.
		c.emit(OpCall, len(expr.Args)) // Use OpCall from code.go

	case *ast.FunctionLiteral: // Assuming ast.FunctionLiteral exists
		// Compile the function literal itself. This doesn't compile the body yet.
		// The body is compiled into a separate set of instructions.

		// Create a new compiler instance for the function body.
		// This compiler will have its own instruction stream and a symbol table
		// nested within the current scope of the outer compiler.
		funcCompiler := c.newFunctionCompiler()

		// Enter the function's scope. This is handled by newFunctionCompiler,
		// which creates a scope enclosed in the parent's current scope.
		// Parameters are defined as locals in the function's symbol table.
		for _, param := range expr.Params { // Assuming Params field exists in ast.FunctionLiteral
			// Define parameter as a Local in the function's scope.
			// *** FIX START: Pass 'param' directly as it's a string ***
			funcCompiler.currentScope.DefineParameter(param) // Assuming param is string, DefineLocal correctly increments numDefinitions
			// *** FIX END ***
		}

		// Compile the function body (the block statement).
		// This compiles statements within the function's instruction stream.
		if err := funcCompiler.compileStatement(expr.Body); err != nil { // Assuming Body field exists
			return err
		}

		// After compiling the body, check if the last instruction is a return.
		// If not, implicitly return null.
		// Check the last instruction in the function compiler's instruction stream.
		lastInstruction := GetLastOpcode(funcCompiler.instructions) // Use GetLastOpcode

		// --- Debug Print: Function Instructions after body compilation ---
		fmt.Printf("DEBUG (Compiler): Function Instructions after body compilation (before implicit return check): %v\n", funcCompiler.instructions)
		// --- End Debug Print ---

		// Check if the last instruction is OpReturnValue or OpReturn.
		// You must ensure these match your actual definitions in code.go.
		isReturn := lastInstruction == OpReturnValue || lastInstruction == OpReturn // Use OpReturnValue and OpReturn from code.go

		// --- Debug Print: Checking if last instruction is OpReturnValue ---
		fmt.Printf("DEBUG (Compiler): Checking if last instruction is %v. Current instructions length: %d\n", OpReturnValue, len(funcCompiler.instructions))
		fmt.Printf("DEBUG (Compiler): Last instruction is %v. Check result: %t\n", lastInstruction, lastInstruction == OpReturnValue)
		// --- End Debug Print ---
		// --- Debug Print: Checking if last instruction is OpReturn ---
		fmt.Printf("DEBUG (Compiler): Checking if last instruction is %v. Current instructions length: %d\n", OpReturn, len(funcCompiler.instructions))
		fmt.Printf("DEBUG (Compiler): Last instruction is %v. Check result: %t\n", lastInstruction, lastInstruction == OpReturn)
		// --- End Debug Print ---

		// *** NOTE: The current implementation of GetLastOpcode in your code.go is a simplification
		// that only returns the last byte. This will be incorrect for instructions with operands.
		// For a robust check, you would need to parse the last instruction's opcode and operands
		// to determine its true type. However, for the specific case of checking for OpReturn/OpReturnValue
		// which have no operands, this simplified check might work if they are the last byte.
		// Be aware of this limitation if you add other no-operand opcodes or if OpReturn/OpReturnValue
		// ever gain operands. ***

		if !isReturn {
			// --- Debug Print: Emitting implicit nil return for FunctionLiteral ---
			fmt.Println("DEBUG (Compiler): Emitting implicit nil return for FunctionLiteral")
			// --- End Debug Print ---
			// Implicit return null.
			funcCompiler.emit(OpNull)   // Use OpNull from code.go
			funcCompiler.emit(OpReturn) // Use OpReturn from code.go
		}

		// The compiled function body instructions are in funcCompiler.instructions.
		// The constants used within the function body are added to the main compiler's constants.
		// The number of locals needed for this function is funcCompiler.currentScope.NumDefinitions().
		// The number of parameters is len(expr.Params).

		// Create a Function object that holds the compiled instructions and metadata.
		// Add this Function object to the main compiler's constant pool.
		// The operand for OpClosure is the index of this Function object in the constant pool.
		// The operand for OpClosure also includes the number of free variables.
		// Free variables are symbols resolved in outer scopes but used within this function.
		// The free variables are stored in the function's symbol table's freeSymbols slice.
		numLocals := funcCompiler.currentScope.NumDefinitions() // Assuming NumDefinitions exists
		numParameters := len(expr.Params)                       // Assuming Params field exists in ast.FunctionLiteral
		freeSymbols := funcCompiler.currentScope.freeSymbols    // Access freeSymbols as a lowercase field

		// Create the CompiledFunction object.
		compiledFn := &types.CompiledFunction{ // Referring to types.CompiledFunction, assuming it exists
			Instructions:  funcCompiler.instructions,
			NumLocals:     numLocals,
			NumParameters: numParameters,
			// Free variables are not stored directly in the CompiledFunction object,
			// but their count and values are needed for OpClosure.
			FreeCount: len(freeSymbols), // Added FreeCount field, assuming it exists on types.CompiledFunction
		}

		// Add the CompiledFunction object to the main compiler's constants.
		fnConstantIndex := c.addConstant(compiledFn) // Use addConstant method

		// Emit OpClosure. Operand is the index of the CompiledFunction constant and the number of free variables.
		// The free variables themselves will be pushed onto the stack by the VM when OpClosure is executed.
		// The compiler needs to ensure the free variables' values are on the stack *before* OpClosure.
		// This requires compiling the free variables' Get operations *before* emitting OpClosure.

		// Compile the Get operations for free variables.
		// These will push the values of the free variables onto the stack.
		// The order here must match the order the VM expects free variables to be on the stack
		// when executing OpClosure. This order is determined by how freeSymbols are collected
		// in the symbol table.
		// before emitting OpClosure, push the actual free values:
		for _, s := range freeSymbols {
			switch s.Kind {
			case Global:
				c.emit(OpGetGlobal, s.Index)
			case Local, Parameter:
				// locals in an _outer_ function scope are still accessed with GetLocal
				c.emit(OpGetLocal, s.Index)
			case Free:
				// the real free variable capture
				c.emit(OpGetFree, s.Index)
			default:
				return fmt.Errorf("unsupported kind for free variable '%s': %v", s.Name, s.Kind)
			}
		}

		// now we know we actually pushed len(freeSymbols) values,
		// so OpClosure’s second operand is correct:
		c.emit(OpClosure, fnConstantIndex, len(freeSymbols))

	default:
		// This case should ideally not be reached.
		// --- Debug Print: Unsupported Expression ---
		fmt.Printf("DEBUG (Compiler): Unsupported Expression type %T\n", e)
		// --- End Debug Print ---
		return fmt.Errorf("unsupported expression type: %T", e)
	}
	return nil
}

// compileBlockStmt compiles a block statement.
// It enters a new scope, compiles the statements within the block, and leaves the scope.
func (c *Compiler) compileBlockStmt(block *ast.BlockStmt) error { // Assuming ast.BlockStmt exists
	// Block statements introduce a new scope.
	// Pass false to indicate this is NOT a function scope.
	c.EnterScope(false) // Pass false for IsFunctionScope

	// Compile each statement in the block.
	for _, stmt := range block.Stmts { // Assuming Stmts field exists
		if err := c.compileStatement(stmt); err != nil {
			// If there's an error, leave the scope before returning the error.
			c.LeaveScope()
			return err
		}
	}

	// Leave the scope after compiling all statements in the block.
	c.LeaveScope()

	return nil
}

// compileIf compiles an if-elseif-else statement.
func (c *Compiler) compileIf(stmt *ast.IfStmt) error { // Assuming ast.IfStmt exists
	// Compile the main condition.
	if err := c.compileExpression(stmt.Cond); err != nil { // Assuming field is Cond
		return err
	}

	// Emit OpJumpNotTruthy with a placeholder operand. This jumps if the condition is false.
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy from code.go

	// Compile the 'then' block.
	// *** FIX START: Use stmt.Then instead of stmt.Consequence ***
	if err := c.compileStatement(stmt.Then); err != nil { // Assuming field is Then
		return err
	}
	// *** FIX END ***

	// Emit OpJump with a placeholder operand to jump over the else-if and else blocks.
	jumpToEndPos := c.emit(OpJump, 0) // Use OpJump from code.go

	// Patch the OpJumpNotTruthy from the condition to point to the instruction
	// immediately after the 'then' block (which is the OpJump or the start of else-if/else).
	afterThenPos := len(c.instructions)
	c.patchJump(jumpNotTruthyPos, afterThenPos)

	// Compile the alternative (else-if or else) block, if present.
	if stmt.Else != nil { // Assuming field is Else
		if err := c.compileStatement(stmt.Else); err != nil {
			return err
		}
	}

	// Patch the OpJump instruction after the consequence: the target is the position *after* the alternative block (or after the consequence if no alternative).
	// The jump target is the current length of instructions.
	c.patchJump(jumpToEndPos, len(c.instructions)) // Pass the target position

	return nil
}

// compileWhile compiles a while loop.
func (c *Compiler) compileWhile(stmt *ast.WhileStmt) error { // Assuming ast.WhileStmt exists
	// Mark the start of the loop condition for the jump back.
	loopStartPos := len(c.instructions)

	// Compile the loop condition.
	if err := c.compileExpression(stmt.Cond); err != nil { // Assuming field is Cond
		return err
	}

	// Emit OpJumpNotTruthy with a placeholder operand. This jumps out of the loop if the condition is false.
	jumpOutOfLoopPos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy from code.go

	// Compile the loop body.
	if err := c.compileStatement(stmt.Body); err != nil { // Assuming field is Body
		return err
	}

	// Emit OpJump with an operand to jump back to the start of the loop condition.
	c.emit(OpJump, loopStartPos) // Use OpJump

	// Patch the OpJumpNotTruthy from the condition to point to the instruction immediately after the loop body (the OpJump).
	// The target is the instruction immediately after the OpJump.
	afterLoopBodyPos := len(c.instructions)
	c.patchJump(jumpOutOfLoopPos, afterLoopBodyPos)

	return nil
}

// compileFor compiles a for loop statement.
func (c *Compiler) compileFor(stmt *ast.ForStmt) error {
	// Compile the iterable expression.
	if err := c.compileExpression(stmt.Iterable); err != nil {
		return err
	}
	// Get an iterator for the iterable.
	c.emit(OpGetIterator) // Use OpGetIterator

	// Enter a new scope for the loop variable and loop body.
	// This scope's function scope status is inherited from the outer scope.
	// *** FIX START: Access isFunctionScope as a lowercase field ***
	c.EnterScope(c.currentScope.isFunctionScope) // Pass the function scope status of the outer scope
	// *** FIX END ***

	var loopVarSymbol *Symbol // Use Symbol from this package

	// Define the loop variable in the current (loop's block) scope.
	// The kind (Global or Local) depends on whether the outer scope is a function scope.
	// *** FIX START: Access isFunctionScope as a lowercase field ***
	if c.currentScope.isFunctionScope {
		// *** FIX END ***
		// If the outer scope is a function scope, the loop variable is local to this function frame.
		loopVarSymbol = c.DefineLocal(stmt.Variable) // Define as Local in the current (loop) scope
	} else {
		// If the outer scope is NOT a function scope (main program or block in main),
		// the loop variable is a global variable.
		// We define it in the global table, even though we are in a nested block scope.
		// This flattens the variable into the global scope for the main program.
		loopVarSymbol = c.globals.Define(stmt.Variable, Global) // Define as Global in the global table
	}

	// Mark the position for the loop's jump back (to get the next item).
	loopStartPos := len(c.instructions) // Use c.instructions

	// Get the next value from the iterator. OpIteratorNext pushes value, boolean (true if successful).
	c.emit(OpIteratorNext) // Use OpIteratorNext

	// Emit OpJumpNotTruthy with a placeholder operand. This jumps out of the loop if the boolean is false (iteration done).
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 9999) // Use OpJumpNotTruthy

	// The next value is now on top of the stack. Assign it to the loop variable.
	// Use the correct opcode based on the loop variable's kind.
	switch loopVarSymbol.Kind {
	case Global:
		c.emit(OpSetGlobal, loopVarSymbol.Index) // Use OpSetGlobal
	case Local:
		c.emit(OpSetLocal, loopVarSymbol.Index) // Use OpSetLocal
	default:
		// This should not happen based on the definition logic above.
		return fmt.Errorf("internal compiler error: unexpected symbol kind for loop variable '%s' after definition: %v", stmt.Variable, loopVarSymbol.Kind)
	}

	// Compile the loop body.
	if err := c.compileBlockStmt(stmt.Body); err != nil {
		// If there's an error, leave the scope before returning.
		c.LeaveScope()
		return err
	}

	// After the loop body, emit OpJump to jump back to the start of the loop (to get the next item).
	c.emit(OpJump, loopStartPos) // Use OpJump

	// Patch the OpJumpNotTruthy operand to point to the instruction immediately after the loop.
	afterLoopPos := len(c.instructions) // Use c.instructions
	c.patchJump(jumpNotTruthyPos, afterLoopPos)

	// Leave the scope for the loop variable and loop body.
	c.LeaveScope()

	// The iterator is still on the stack after the loop finishes. Pop it.
	c.emit(OpPop) // Use OpPop

	// The for loop itself doesn't leave a value on the stack.
	// If the language requires a value, you might need to push a default value (like nil) here.

	return nil
}

// emit adds an opcode and its operands to the instructions.
// It returns the starting position of the emitted instruction.
func (c *Compiler) emit(op OpCode, operands ...int) int { // Use OpCode from code.go
	// Get the definition of the opcode to know the operand widths.
	def, ok := Lookup(op) // Use Lookup from code.go
	if !ok {
		// This should not happen if all opcodes are defined.
		panic(fmt.Sprintf("unknown opcode %v", op))
	}

	// Calculate the instruction length (opcode + operands).
	instructionLen := 1 // Opcode byte
	for _, width := range def.OperandWidths {
		instructionLen += width
	}

	// Ensure enough capacity in the instructions slice.
	// Grow the slice if needed.
	newInstructions := make(Instructions, instructionLen) // Use Instructions from code.go
	newInstructions[0] = byte(op)

	// Encode and add operands.
	offset := 1
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 1: // 1-byte operand (uint8)
			newInstructions[offset] = byte(operand)
		case 2: // 2-byte operand (uint16)
			binary.BigEndian.PutUint16(newInstructions[offset:], uint16(operand)) // Use binary.BigEndian
		// Add cases for other operand widths if needed (e.g., 4 bytes for uint32)
		default:
			// This should not happen if operand widths are defined correctly.
			panic(fmt.Sprintf("unsupported operand width %d", width))
		}
		offset += width
	}

	// Get the current position before appending.
	pos := len(c.instructions)

	// Append the new instruction to the instructions slice.
	c.instructions = append(c.instructions, newInstructions...)

	// --- Debug Print: Emitted Opcode ---
	fmt.Printf("DEBUG (Compiler): Emitted Opcode %v at position %d\n", op, pos)
	// --- End Debug Print ---

	// Return the starting position of the emitted instruction.
	return pos
}

// emitConstant adds a value to the constant pool, emits OpConstant, and returns the instruction position.
func (c *Compiler) emitConstant(obj types.Value) int {
	// Add the constant to the pool and get its index.
	constIndex := c.addConstant(obj)
	// Emit the OpConstant instruction with the constant's index as the operand.
	// OpConstant has a 2-byte operand for the constant pool index.
	return c.emit(OpConstant, constIndex) // Use OpConstant from code.go
}

// addConstant adds a value to the constant pool and returns its index.
func (c *Compiler) addConstant(obj types.Value) int { // Takes types.Value
	// Check if the constant already exists to avoid duplicates.
	// This requires implementing an Equals method on all Value types.
	// Assuming types.Value has an Equals method.
	for i, constant := range c.constants {
		if constant.Equals(obj) { // Assuming Equals method exists on types.Value
			return i // Return index of existing constant
		}
	}

	// If not found, add the new constant.
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1 // Return index of the newly added constant
}

// lastInstructionIs checks if the last emitted instruction is the given opcode.
// *** NOTE: This relies on the simplified GetLastOpcode which only reads the last byte.
// This is only correct for opcodes without operands if they are the last instruction. ***
// This method is currently unused, but kept for potential future use.
func (c *Compiler) lastInstructionIs(op OpCode) bool { // Use OpCode from code.go
	// --- Debug Print: Checking last instruction ---
	fmt.Printf("DEBUG (Compiler): Checking if last instruction is %v. Current instructions length: %d\n", op, len(c.instructions))
	// --- End Debug Print ---
	if len(c.instructions) == 0 {
		// --- Debug Print: lastInstructionIs returning false (empty instructions) ---
		fmt.Println("DEBUG (Compiler): lastInstructionIs returning false (empty instructions)")
		// --- End Debug Print ---
		return false
	}
	// Get the last emitted instruction's opcode using the standalone function.
	lastOp := GetLastOpcode(c.instructions) // Use GetLastOpcode
	// --- Debug Print: lastInstructionIs check result ---
	fmt.Printf("DEBUG (Compiler): Last instruction is %v. Check result: %t\n", lastOp, lastOp == op)
	// --- End Debug Print ---
	return lastOp == op
}

// removeLastPop removes the last emitted OpPop instruction.
// This is useful for expression statements that are the last statement in a block
// and whose value needs to be the result of the block/function.
// This method is currently unused, but kept for potential future use.
func (c *Compiler) removeLastPop() {
	// Get the definition of OpPop to know its length (1 byte).
	// OpPop has no operands, so its size is always 1 byte.
	opPopSize := 1

	// Ensure the last instruction is indeed OpPop.
	if !c.lastInstructionIs(OpPop) { // Use OpPop from code.go
		// Should not happen if called correctly.
		panic("last instruction is not OpPop")
	}
	// Remove the last instruction (OpPop).
	c.instructions = c.instructions[:len(c.instructions)-opPopSize]
}

// patchJump modifies the operand of a jump instruction.
// It takes the position of the jump instruction and the target position.
func (c *Compiler) patchJump(jumpPos int, targetPos int) {
	// Get the definition of the jump opcode (OpJump or OpJumpNotTruthy).
	// Assuming they both have a 2-byte operand for the jump target.
	def, ok := Lookup(OpCode(c.instructions[jumpPos])) // Use Lookup from code.go
	if !ok {
		// Should not happen.
		panic(fmt.Sprintf("opcode at position %d not found", jumpPos))
	}
	// Ensure the operand width is 2 bytes.
	if len(def.OperandWidths) != 1 || def.OperandWidths[0] != 2 {
		// Should not happen if jump opcodes are defined correctly.
		panic(fmt.Sprintf("opcode at position %d is not a jump instruction with 2-byte operand", jumpPos))
	}

	// The target position is relative to the start of the instructions.
	// The operand is the relative jump offset.
	// For a forward jump, the operand is the distance from the *end* of the jump instruction
	// to the target instruction.
	// Jump instruction starts at jumpPos, its operand is at jumpPos + 1.
	// The operand is 2 bytes, so the instruction ends at jumpPos + 1 + 2 = jumpPos + 3.
	// The distance is targetPos - (jumpPos + 3).
	// However, in many VMs, the jump target is simply the absolute instruction index.
	// Let's assume the operand is the absolute target position for simplicity.
	// If your VM expects a relative jump, you'll need to calculate the offset.
	// Assuming absolute target for now.

	// Encode the target position as a 2-byte unsigned integer.
	operand := uint16(targetPos)
	// Write the operand into the instructions slice at the correct position.
	binary.BigEndian.PutUint16(c.instructions[jumpPos+1:], operand) // Use binary.BigEndian
}

// Bytecode returns the compiled bytecode for the current compilation unit.
// This method is typically called on the main compiler instance after compiling the program.
func (c *Compiler) Bytecode() *Bytecode { // Returning *Bytecode from code.go
	// The bytecode is already finalized in the Compile method for the main program.
	// For function compilers, this method would return a new Bytecode object
	// containing the function's instructions and a reference to the shared constants.
	// Since this method is likely intended for the main compiler, we return currentBytecode.
	// Ensure Bytecode object is populated correctly in the Compile method.
	return c.currentBytecode
}

// Instructions returns the current instruction stream being built.
func (c *Compiler) Instructions() Instructions { // Returning Instructions from code.go
	return c.instructions
}

// Constants returns the current constant pool.
func (c *Compiler) Constants() []types.Value { // Returning []types.Value
	return c.constants
}

// LastInstruction returns the last emitted instruction and its operands.
// *** NOTE: This relies on the simplified GetLastOpcode which only reads the last byte.
// This is only correct for opcodes without operands if they are the last instruction. ***
func (c *Compiler) LastInstruction() (OpCode, []int) { // Returning OpCode and []int
	// Get the last instruction's opcode using the standalone function.
	if len(c.instructions) == 0 {
		return 0, nil // Or return an error/sentinel value
	}
	lastOpCode := GetLastOpcode(c.instructions) // Use GetLastOpcode

	// Get the definition to determine operand widths.
	def, ok := Lookup(lastOpCode) // Use Lookup from code.go
	if !ok {
		// Should not happen for valid opcodes.
		return lastOpCode, nil // Or return an error/sentinel value
	}

	// Extract operands. This is simplified and assumes fixed-width operands.
	// A more robust approach would parse operands based on definition widths.
	operands := []int{}
	offset := 1 // Start after the opcode
	for _, width := range def.OperandWidths {
		switch width {
		case 1:
			// *** FIX START: Correct bounds checking and operand reading ***
			if len(c.instructions) >= len(c.instructions)-(len(c.instructions)-offset)+1 { // Check bounds from the end
				operands = append(operands, int(ReadUint8(c.instructions[len(c.instructions)-offset:]))) // Use ReadUint8
			} else {
				// Not enough bytes for operand
				break // Exit operand reading loop
			}
		case 2:
			if len(c.instructions) >= len(c.instructions)-(len(c.instructions)-offset)+2 { // Check bounds from the end
				operands = append(operands, int(ReadUint16(c.instructions[len(c.instructions)-offset:]))) // Use ReadUint16
			} else {
				// Not enough bytes for operand
				break // Exit operand reading loop
			}
			// *** FIX END ***
		}
		offset += width
	}

	return lastOpCode, operands
}

// RemoveLastInstruction removes the last emitted instruction from the stream.
// This is useful for optimizations or backtracking in compilation.
// *** NOTE: This relies on the simplified GetLastOpcode and assumes the last instruction
// is correctly identified and its length is accurately determined based on the last byte. ***
func (c *Compiler) RemoveLastInstruction() {
	if len(c.instructions) == 0 {
		return // Nothing to remove
	}

	// Get the last instruction's opcode using the standalone function.
	lastOpCode := GetLastOpcode(c.instructions) // Use GetLastOpcode
	def, ok := Lookup(lastOpCode)               // Use Lookup from code.go
	if !ok {
		// Should not happen for valid opcodes.
		return
	}

	// Calculate the instruction length based on the definition.
	instructionLen := 1 // Opcode byte
	for _, width := range def.OperandWidths {
		instructionLen += width
	}

	// Truncate the instruction slice.
	if len(c.instructions) >= instructionLen {
		c.instructions = c.instructions[:len(c.instructions)-instructionLen]
	} else {
		// Should not happen if emit is correct, but handle defensively
		c.instructions = c.instructions[:0] // Clear instructions if length is inconsistent
	}
}

// CurrentScope returns the current symbol table scope.
func (c *Compiler) CurrentScope() *SymbolTable { // Returning *SymbolTable
	return c.currentScope
}

// GlobalScope returns the global symbol table.
func (c *Compiler) GlobalScope() *SymbolTable { // Returning *SymbolTable
	return c.globals
}

// numDefinitions in the current scope (for local/parameter counting)
// This method name conflicts with the field in SymbolTable.
// Accessing the field directly is clearer.
// *** FIX START: Remove unused method ***
// func (c *Compiler) numDefinitions() int {
// 	// Accessing the field directly from symboltable.SymbolTable
// 	return c.currentScope.numDefinitions
// }
// *** FIX END ***

// FreeSymbols in the current scope (for closure compilation)
// This method name conflicts with the field in SymbolTable.
// Accessing the field directly is clearer.
func (c *Compiler) FreeSymbols() []*Symbol { // Returning []*Symbol
	return c.currentScope.freeSymbols
}
