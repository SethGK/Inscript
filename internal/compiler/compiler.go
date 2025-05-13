// Package compiler compiles the Inscript AST into bytecode.
package compiler

import (
	"encoding/binary" // Needed for binary encoding/decoding in patchJump
	"fmt"

	// Needed for Fprintf to stderr
	"github.com/SethGK/Inscript/internal/ast"
	"github.com/SethGK/Inscript/internal/types"
	// Assuming SymbolTable, Symbol, SymbolKind, NewSymbolTable, NewEnclosedSymbolTable
	// are defined in symboltable.go in the same package.
	// Assuming Instructions, Bytecode, NewBytecode, OpCode, Definition, Make, Lookup, GetLastOpcode
	// are defined in code.go in the same package.
	// Note: The import path for the generated parser package should be correct.
	// parser "github.com/SethGK/Inscript/parser/grammar"
)

// Compiler translates an AST into bytecode.
type Compiler struct {
	// The bytecode being compiled for the current compilation unit (program or function)
	instructions Instructions  // Use Instructions type from code.go
	constants    []types.Value // Constant pool for the current compilation unit - Referring to types.Value

	// Global symbol table (unique for the entire program)
	// This is shared across all function compilers.
	globals *SymbolTable // Referring to SymbolTable from symboltable.go
	// Removed numGlobals field - the count is now in c.globals.numDefinitions

	// Scope management
	symbolTableStack []*SymbolTable // Referring to SymbolTable from symboltable.go
	currentScope     *SymbolTable   // Referring to SymbolTable from symboltable.go

	// Bytecode buffer for the current compilation unit
	currentBytecode *Bytecode // Use Bytecode type from code.go
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
	return c.globals.numDefinitions // Use the count from the global symbol table
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
	if isFunctionScope && !c.currentScope.isFunctionScope {
		// Only allow entering a function scope if the current scope is not already one.
		// This prevents accidentally nesting function scopes using EnterScope.
		panic("EnterScope(true) should only be called when entering a function's top-level scope")
	}
	if isFunctionScope && c.currentScope.isFunctionScope {
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
// This now increments the global symbol table's definition count.
func (c *Compiler) DefineGlobal(name string, kind SymbolKind) *Symbol { // Referring to SymbolKind, Symbol from symboltable.go
	// Check if it already exists in the global table
	symbol, ok := c.globals.store[name]
	if ok {
		// If it exists, return the existing symbol.
		// This handles cases where a global is used before being assigned.
		return symbol
	}

	// Define a new global symbol
	// The index is the current number of definitions in the global table.
	symbol = &Symbol{Name: name, Kind: kind, Index: c.globals.numDefinitions} // Referring to Symbol from symboltable.go
	c.globals.store[name] = symbol
	c.globals.numDefinitions++ // Manually increment here since we are not using the SymbolTable.Define method directly for globals in this helper
	return symbol
}

// DefineLocal defines a symbol in the current scope's symbol table.
// Returns the created Symbol.
func (c *Compiler) DefineLocal(name string) *Symbol { // Referring to Symbol from symboltable.go
	// Check if the symbol is already defined in the current scope.
	// This prevents re-declaring variables in the same scope (if language rules require this).
	// For now, we allow shadowing outer scopes but not re-declaration in the same scope.
	_, ok := c.currentScope.store[name]
	if ok {
		// Variable already defined in the current scope.
		// Depending on language semantics, you might want to return an error here.
		// For now, let's allow it, and the new definition will shadow the old one in this scope.
		// A stricter language would return fmt.Errorf("variable '%s' already defined in this scope", name)
	}

	// Define a new local symbol in the current scope.
	// The index is the current number of definitions in this scope.
	symbol := &Symbol{Name: name, Kind: Local, Index: c.currentScope.numDefinitions} // Referring to Symbol, Local from symboltable.go
	c.currentScope.store[name] = symbol
	c.currentScope.numDefinitions++ // Manually increment here since we are not using the SymbolTable.Define method directly for locals in this helper
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
	c.currentBytecode.NumGlobals = c.globals.numDefinitions // Use the count from the global symbol table

	return c.currentBytecode, nil
}

// compileStatement emits bytecode for a statement node.
func (c *Compiler) compileStatement(s ast.Statement) error {
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
		// For a standalone expression statement, the result is not used, so pop it.
		// This OpPop is necessary for expressions used as statements (e.g., `1 + 2;`).
		c.emit(OpPop) // Use OpPop from code.go

	case *ast.AssignStmt:
		// Handle the assignment target based on its type.
		switch target := stmt.Target.(type) {
		case *ast.Identifier:
			// Assignment to a variable (identifier).
			// Compile the value expression first. It leaves the value on the stack.
			// --- Debug Print: Compiling AssignStmt value for target '%s' ---
			fmt.Printf("DEBUG (Compiler): Compiling AssignStmt value for target '%s' (type %T)\n", target.Name, stmt.Value)
			// --- End Debug Print ---
			// --- Debug Print: Before compileExpression for AssignStmt value ---
			fmt.Printf("DEBUG (Compiler): Before compileExpression for AssignStmt value. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---
			if err := c.compileExpression(stmt.Value); err != nil {
				return err
			}
			// --- Debug Print: After compileExpression for AssignStmt value ---
			fmt.Printf("DEBUG (Compiler): After compileExpression for AssignStmt value. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---

			// Look up the symbol in the current and outer scopes.
			symbol, ok := c.currentScope.Resolve(target.Name)
			if !ok {
				// If the variable is not found in any scope, define it.
				// Variables defined in the main program (including nested blocks) are globals.
				// Variables defined in function scopes are locals.
				if c.currentScope.isFunctionScope {
					// Define as Local in the current function scope.
					// Corrected call to Define: provide the SymbolKind.
					symbol = c.currentScope.Define(target.Name, Local) // Pass name and Local kind from symboltable.go
				} else {
					// If not in a function scope (global or block in main program), define as Global.
					// Corrected call to Define: provide the SymbolKind.
					symbol = c.globals.Define(target.Name, Global) // Pass name and Global kind from symboltable.go
				}
			}

			// Emit the appropriate instruction based on the symbol's kind.
			// OpSetGlobal and OpSetLocal consume the value from the stack.
			switch symbol.Kind {
			case Global: // Referring to Global from symboltable.go
				// Assign to a global variable. Operand is the global index.
				c.emit(OpSetGlobal, symbol.Index) // Use OpSetGlobal from code.go
			case Local, Parameter: // Parameters are also accessed as locals within the function frame from symboltable.go
				// Assign to a local variable or parameter. Operand is the local/parameter index within the frame.
				// This case should ONLY be reached if compiling within a function scope.
				if !c.currentScope.isFunctionScope {
					return fmt.Errorf("internal compiler error: attempting to set local/parameter outside function scope for '%s'", symbol.Name)
				}
				// The operand for locals is 1 byte (uint8).
				c.emit(OpSetLocal, symbol.Index) // Use OpSetLocal from code.go
			case Builtin: // Referring to Builtin from symboltable.go
				// Cannot assign to a builtin.
				return fmt.Errorf("cannot assign to builtin '%s'", symbol.Name)
			default:
				return fmt.Errorf("unsupported symbol kind for assignment target: %v", symbol.Kind)
			}

		case *ast.IndexExpr:
			// Assignment to an indexed element (e.g., `my_list[index] = value`).
			// The required stack order for OpSetIndex is [..., aggregate, index, value].

			// 1. Compile the aggregate expression (e.g., `my_list`).
			// This pushes the aggregate object onto the stack.
			if err := c.compileExpression(target.Primary); err != nil {
				return err
			}

			// 2. Compile the index expression (e.g., integer for list, string/value for table).
			// This pushes the index value onto the stack.
			if err := c.compileExpression(target.Index); err != nil {
				return err
			}

			// 3. Compile the value expression (the value to be assigned).
			// This pushes the value onto the stack.
			// --- Debug Print: Compiling AssignStmt IndexExpr value ---
			fmt.Printf("DEBUG (Compiler): Compiling AssignStmt IndexExpr value (type %T)\n", stmt.Value)
			// --- End Debug Print ---
			// --- Debug Print: Before compileExpression for AssignStmt IndexExpr value ---
			fmt.Printf("DEBUG (Compiler): Before compileExpression for AssignStmt IndexExpr value. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---
			if err := c.compileExpression(stmt.Value); err != nil { // Compile stmt.Value LAST
				return err
			}
			// --- Debug Print: After compileExpression for AssignStmt IndexExpr value ---
			fmt.Printf("DEBUG (Compiler): After compileExpression for AssignStmt IndexExpr value. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---

			// Stack top is now [..., aggregate, index, value]. Correct order for OpSetIndex.
			// Emit OpSetIndex. The VM should pop value, index, aggregate, and perform the assignment.
			c.emit(OpSetIndex) // Use OpSetIndex from code.go

		// Add cases for other potential assignment targets if your grammar supports them,
		// e.g., field access like `my_object.field = value`.

		default:
			return fmt.Errorf("unsupported assignment target type: %T", stmt.Target)
		}

	case *ast.PrintStmt:
		// Compile each expression to be printed. They leave their results on the stack.
		for _, e := range stmt.Exprs {
			// --- Debug Print: Compiling PrintStmt expression ---
			fmt.Printf("DEBUG (Compiler): Compiling PrintStmt expression (type %T)\n", e)
			// --- End Debug Print ---
			// --- Debug Print: Before compileExpression for PrintStmt expression ---
			fmt.Printf("DEBUG (Compiler): Before compileExpression for PrintStmt expression. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---
			if err := c.compileExpression(e); err != nil {
				return err
			}
			// --- Debug Print: After compileExpression for PrintStmt expression ---
			fmt.Printf("DEBUG (Compiler): After compileExpression for PrintStmt expression. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---
		}
		// Emit the print instruction with the number of arguments to print.
		c.emit(OpPrint, len(stmt.Exprs)) // Use OpPrint from code.go

	case *ast.ReturnStmt:
		// Compile the return value expression, if present.
		if stmt.Expr != nil {
			// --- Debug Print: Compiling ReturnStmt expression ---
			fmt.Printf("DEBUG (Compiler): Compiling ReturnStmt expression (type %T)\n", stmt.Expr)
			// --- End Debug Print ---
			// --- Debug Print: Before compileExpression for ReturnStmt expression ---
			fmt.Printf("DEBUG (Compiler): Before compileExpression for ReturnStmt expression. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---
			if err := c.compileExpression(stmt.Expr); err != nil {
				return err
			}
			// --- Debug Print: After compileExpression for ReturnStmt expression ---
			fmt.Printf("DEBUG (Compiler): After compileExpression for ReturnStmt expression. Instructions length: %d\n", len(c.instructions))
			// --- End Debug Print ---

			// Emit OpReturnValue if a value is returned. OpReturnValue consumes the value from the stack.
			c.emit(OpReturnValue) // Use OpReturnValue (34)
		} else {
			// Emit OpReturn if no value is returned (implicitly returns null).
			// The compiler should ensure OpNull is on the stack before OpReturn,
			// or the VM handles pushing nil for OpReturn. Current compiler
			// emits OpNull before OpReturn in the implicit return case, which is good.
			c.emit(OpReturn) // Use OpReturn (35)
		}
		// Note: A return statement typically causes an immediate jump out of the function.
		// The VM handles this when it encounters OpReturnValue or OpReturn.
		// No explicit jump instruction is needed here in the compiler after the return opcode.

	case *ast.BlockStmt:
		// Delegate compilation of the block to the compileBlockStmt method.
		return c.compileBlockStmt(stmt)

	case *ast.IfStmt:
		// Delegate to helper for if-elseif-else logic
		return c.compileIf(stmt)

	case *ast.WhileStmt:
		return c.compileWhile(stmt) // Delegate to helper for while loop logic

	case *ast.ForStmt:
		return c.compileFor(stmt) // Delegate to helper for for loop logic (needs implementation)

	// Removed case *ast.FuncDefStmt: from compileStatement
	// Assuming function definitions are now handled as expressions via *ast.FunctionLiteral

	default:
		// This case should ideally not be reached.
		// --- Debug Print: Unsupported Statement ---
		fmt.Printf("DEBUG (Compiler): Unsupported Statement type %T\n", s)
		// --- End Debug Print ---
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
		c.emitConstant(types.NewInteger(expr.Value))

	case *ast.FloatLiteral:
		c.emitConstant(types.NewFloat(expr.Value))

	case *ast.StringLiteral:
		c.emitConstant(types.NewString(expr.Value))

	case *ast.BooleanLiteral:
		if expr.Value {
			c.emit(OpTrue)
		} else {
			c.emit(OpFalse)
		}

	case *ast.NilLiteral:
		c.emit(OpNull)

	case *ast.Identifier:
		symbol, ok := c.currentScope.Resolve(expr.Name)
		if !ok {
			return fmt.Errorf("undefined variable '%s'", expr.Name)
		}
		switch symbol.Kind {
		case Global:
			c.emit(OpGetGlobal, symbol.Index)
		case Local, Parameter:
			if !c.currentScope.isFunctionScope {
				return fmt.Errorf("getting local '%s' outside function scope", symbol.Name)
			}
			c.emit(OpGetLocal, symbol.Index)
		case Builtin:
			c.emit(OpGetGlobal, symbol.Index) // Builtins are stored in globals
		default:
			return fmt.Errorf("unsupported symbol kind for get: %v", symbol.Kind)
		}

	case *ast.BinaryExpr:
		// Handle short-circuiting for 'and' and 'or'.
		if expr.Operator == "and" {
			// Compile the left side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Left); err != nil {
				return err
			}
			// If the left side is false (or nil), jump over the right side compilation.
			// The falsy value is already on the stack and is the result of the 'and' if short-circuited.
			jumpFalsePos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy from code.go

			// If the left side was true, pop its value (we only need the right side's value for the result).
			c.emit(OpPop) // Use OpPop from code.go

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil {
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpFalsePos, len(c.instructions)) // Pass the target position (current length)
			// The result of the expression (either Left's falsy value or Right's value) is now on the stack.

		} else if expr.Operator == "or" {
			// Compile the left side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Left); err != nil {
				return err
			}
			// If the left side is true (or not nil), jump over the right side compilation.
			// The truthy value is already on the stack and is the result of the 'or' if short-circuited.
			jumpTruePos := c.emit(OpJumpTruthy, 0) // Use OpJumpTruthy from code.go

			// If the left side was false, pop its value (we only need the right side's value for the result).
			c.emit(OpPop) // Use OpPop from code.go

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil {
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpTruePos, len(c.instructions)) // Pass the target position (current length)
			// The result of the expression (either Left's falsy value or Right's value) is now on the stack.

		} else {
			// For other binary operators (+, -, *, /, %, ^, ==, !=, <, >, <=, >=).
			// Compile the left side first. It leaves its result on the stack.
			if err := c.compileExpression(expr.Left); err != nil {
				return err
			}
			// Compile the right side second. It leaves its result on the stack.
			if err := c.compileExpression(expr.Right); err != nil {
				return err
			}

			// Emit the opcode corresponding to the operator.
			// The VM will pop the two operands, perform the operation, and push the result.
			switch expr.Operator {
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

	case *ast.UnaryExpr:
		// Compile the expression the unary operator applies to. It leaves its result on the stack.
		if err := c.compileExpression(expr.Expr); err != nil {
			return err
		}

		// Emit the opcode corresponding to the unary operator.
		// The VM will pop the operand, perform the operation, and push the result.
		switch expr.Operator {
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

	case *ast.ListLiteral:
		// Compile each element of the list. They leave their results on the stack in order.
		count := len(expr.Elements)
		for _, el := range expr.Elements {
			if err := c.compileExpression(el); err != nil {
				return err
			}
		}
		// Emit OpArray with the number of elements.
		// The VM will pop the elements from the stack and create a list.
		c.emit(OpArray, count) // Use OpArray from code.go

	case *ast.TableLiteral:
		// Compile each field (key and value) of the table.
		// Keys and values are pushed onto the stack alternatingly: key1, value1, key2, value2, ...
		count := len(expr.Fields)
		for _, f := range expr.Fields {
			// Table keys are identifiers in the grammar, but typically evaluated to strings or other hashable values.
			// Assuming keys are treated as strings in the bytecode/VM.
			// Add the key string to the constant pool and emit OpConstant.
			c.emitConstant(types.NewString(f.Key)) // Referring to types.NewString (from types package)
			// Compile the value expression.
			if err := c.compileExpression(f.Value); err != nil {
				return err
			}
		}
		// Emit OpHash with the number of fields (key-value pairs).
		// The VM will pop 2*count values from the stack (key, value pairs) and create a hash/table.
		c.emit(OpHash, count) // Use OpHash from code.go

	case *ast.IndexExpr:
		// Compile the primary expression (the aggregate, e.g., list or table). Leaves aggregate on stack.
		if err := c.compileExpression(expr.Primary); err != nil {
			return err
		}
		// Compile the index expression (e.g., integer for list, string/value for table). Leaves index on stack.
		if err := c.compileExpression(expr.Index); err != nil {
			return err
		}
		// Stack top is now [..., aggregate, index].
		// Emit OpIndex. The VM should pop index, pop aggregate, perform the index lookup, and push the result.
		c.emit(OpIndex) // Use OpIndex from code.go

	case *ast.CallExpr:
		if err := c.compileExpression(expr.Callee); err != nil {
			return err
		}
		for _, arg := range expr.Args {
			if err := c.compileExpression(arg); err != nil {
				return err
			}
		}
		c.emit(OpCall, len(expr.Args))

	case *ast.FunctionLiteral:
		// New function scope
		fnCompiler := c.newFunctionCompiler()
		// Define parameters
		for _, p := range expr.Params {
			fnCompiler.DefineLocal(p)
		}
		// Compile body statements
		for _, stmt := range expr.Body.Stmts {
			if err := fnCompiler.compileStatement(stmt); err != nil {
				return err
			}
		}

		// --- Debug Print: Function Instructions After Body Compilation ---
		fmt.Printf("DEBUG (Compiler): Function Instructions after body compilation: %v\n", fnCompiler.instructions)
		// --- End Debug Print ---

		// Implicit return nil if needed
		// Check if the function's instruction slice is empty or if the last instruction is NOT a return opcode.
		if len(fnCompiler.instructions) == 0 || (!fnCompiler.lastInstructionIs(OpReturn) && !fnCompiler.lastInstructionIs(OpReturnValue)) {
			// --- Debug Print: Emitting implicit nil return for FunctionLiteral ---
			fmt.Printf("DEBUG (Compiler): Emitting implicit nil return for FunctionLiteral\n")
			// --- End Debug Print ---
			fnCompiler.emit(OpNull)
			fnCompiler.emit(OpReturn)
		}
		// The else if fnCompiler.lastInstructionIs(OpPop) block is removed as it's not needed
		// with the corrected implicit return logic and the OpPop is handled in ExprStmt.

		// Gather free symbols
		freeSyms := fnCompiler.currentScope.FreeSymbols()
		// Create CompiledFunction
		compiledFn := &types.CompiledFunction{
			Instructions:  fnCompiler.instructions, // <-- This is the slice being stored
			NumParameters: len(expr.Params),
			NumLocals:     fnCompiler.currentScope.NumDefinitions(),
			FreeCount:     len(freeSyms),
		}
		// Add to constant pool
		constIndex := c.addConstant(compiledFn)
		// Emit closure opcode, capturing frees
		c.emit(OpClosure, constIndex, len(freeSyms))

	default:
		// --- Debug Print: Unsupported Expression ---
		fmt.Printf("DEBUG (Compiler): Unsupported Expression type %T\n", e)
		// --- End Debug Print ---
		return fmt.Errorf("unsupported expression type: %T", e)
	}
	return nil
}

// compileBlockStmt compiles a block statement.
func (c *Compiler) compileBlockStmt(block *ast.BlockStmt) error {
	// Block statements introduce a new scope.
	// Pass false to indicate this is NOT a function scope.
	c.EnterScope(false)

	// Compile each statement in the block.
	for _, stmt := range block.Stmts {
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
func (c *Compiler) compileIf(stmt *ast.IfStmt) error {
	// Compile the main condition.
	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}

	// Emit OpJumpNotTruthy with a placeholder operand. This jumps if the condition is false.
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy from code.go

	// Compile the 'then' block.
	if err := c.compileBlockStmt(stmt.Then); err != nil {
		return err
	}

	// After the 'then' block, if the last instruction was OpPop (from an expression statement), remove it.
	// The result of the if expression should be the value of the last statement in the executed block,
	// or nil if the block is empty or ends with a non-expression statement.
	// If the 'then' block is empty or ends with a statement that doesn't leave a value,
	// we might need to push a default value (like nil) here depending on language semantics.
	// Assuming for now that the last instruction of the block handles the value.
	// If the block is empty, compileBlockStmt might emit nothing.
	// TODO: Handle empty blocks and ensure a value is left on the stack for the if expression.
	// if c.lastInstructionIs(OpPop) { // Use OpPop from code.go
	//  c.removeLastPop()
	// }

	// Emit OpJump with a placeholder operand to jump over the else-if and else blocks.
	jumpToEndPos := c.emit(OpJump, 0) // Use OpJump from code.go

	// Patch the OpJumpNotTruthy from the condition to point to the instruction
	// immediately after the 'then' block (which is the OpJump or the start of else-if/else).
	afterThenPos := len(c.instructions)
	c.patchJump(jumpNotTruthyPos, afterThenPos)

	// Compile else-if blocks
	// TODO: Implement compilation of else-if blocks

	// Compile the else block, if present.
	if stmt.Else != nil {
		if err := c.compileBlockStmt(stmt.Else); err != nil {
			return err
		}
		// After the 'else' block, if the last instruction was OpPop, remove it.
		// TODO: Handle empty else blocks and ensure a value is left on the stack.
		// if c.lastInstructionIs(OpPop) { // Use OpPop from code.go
		//  c.removeLastPop()
		// }
	} else {
		// If there's no else block, and the 'then' block was not executed,
		// the result of the if expression is implicitly nil.
		// We need to push nil here if the condition was false and there was no else.
		// This happens when the OpJumpNotTruthy is taken and there was no else block.
		// So, the target of jumpNotTruthyPos should implicitly push nil if no else.
		// This requires adjusting the jump target or emitting nil after the jump target.
		// A simpler approach is to always push nil after the OpJumpNotTruthy if there's no else block.
		// This is done by placing the nil push *after* the jump target in the bytecode stream.
		// So, if there's no else, the jumpNotTruthy jumps *past* the 'then' block and lands here.
		// We need to ensure nil is pushed here.
		// The current logic is flawed for the "no else" case regarding the expression value.
		// Let's refine the jump targets.

		// Correct logic for if (expr) { block }
		// compile cond -> OpJumpNotTruthy placeholder -> compile block -> OpJump placeholder -> Patch JumpNotTruthy to after block -> Patch Jump to after everything.
		// If no else, the OpJumpNotTruthy should jump to after the OpJump.
		// The value of the if expression when no else and condition is false is nil.
		// We need to ensure nil is on the stack in that case.

		// Let's rethink the if compilation for expression value.
		// The if expression's value is the value of the last expression in the executed block.
		// If no block is executed (condition false, no else), the value is nil.

		// New approach for if (expr) { block } else { block } as an expression:
		// compile cond -> OpJumpNotTruthy placeholder_else_or_end
		// compile then_block -> OpJump placeholder_end
		// patch jumpNotTruthyPos to current instruction (start of else or end)
		// if else_block exists:
		//   compile else_block
		// patch jumpToEndPos to current instruction (after else_block)
		//
		// For expression value:
		// The last instruction of the executed block must leave its value on the stack.
		// If a block is empty or ends with a statement that doesn't leave a value,
		// we need to push nil at the end of that block's compilation.

		// Let's revert to a simpler model for now: if statements do not produce a value.
		// Any value left by the last expression statement in a block is popped.
		// The compiler's current logic for OpPop after ExprStmt supports this.
		// The issue with nil results likely stems from the compiler not leaving
		// the return value on the stack before OpReturnValue, or the VM issue.
		// The if/else compilation itself seems okay for control flow, but might need
		// adjustment if if-expressions are a language feature. Assuming they are not for now.

		// Let's fix the jumps based on the current control flow model.
		// jumpNotTruthyPos jumps to the start of the else block or the end if no else.
		// jumpToEndPos jumps from the end of the then block to the end of the if statement.

		// The current structure:
		// compile cond
		// OpJumpNotTruthy placeholder1
		// compile then_block
		// OpJump placeholder2
		// patch placeholder1 to here (start of else or end)
		// compile else_block (if any)
		// patch placeholder2 to here (end of if)

		// This control flow seems correct for skipping blocks.
		// The value issue is separate.

	}

	// Patch the OpJump from the 'then' block to point to the instruction immediately after the if statement.
	endOfIfPos := len(c.instructions)
	c.patchJump(jumpToEndPos, endOfIfPos)

	return nil
}

// compileWhile compiles a while loop.
func (c *Compiler) compileWhile(stmt *ast.WhileStmt) error {
	// Mark the start of the loop condition for the jump back.
	loopStartPos := len(c.instructions)

	// Compile the loop condition.
	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}

	// Emit OpJumpNotTruthy with a placeholder operand. This jumps out of the loop if the condition is false.
	jumpOutOfLoopPos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy from code.go

	// Compile the loop body.
	if err := c.compileBlockStmt(stmt.Body); err != nil {
		return err
	}

	// Emit OpJump with an operand to jump back to the start of the loop condition.
	c.emit(OpJump, loopStartPos) // Use OpJump from code.go

	// Patch the OpJumpNotTruthy from the condition to point to the instruction immediately after the loop body (the OpJump).
	// The target is the instruction immediately after the OpJump.
	afterLoopBodyPos := len(c.instructions)
	c.patchJump(jumpOutOfLoopPos, afterLoopBodyPos)

	return nil
}

// compileFor compiles a for loop.
func (c *Compiler) compileFor(stmt *ast.ForStmt) error {
	// TODO: Implement for loop compilation.
	// A typical for loop (init; cond; post) body can be compiled as:
	// compile init (if any)
	// loop_start:
	// compile cond
	// OpJumpNotTruthy jump_out_of_loop
	// compile body
	// compile post (if any)
	// OpJump loop_start
	// jump_out_of_loop:
	return fmt.Errorf("for loops are not yet implemented")
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
	for i, constant := range c.constants {
		if constant.Equals(obj) {
			return i // Return index of existing constant
		}
	}

	// If not found, add the new constant.
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1 // Return index of the newly added constant
}

// lastInstructionIs checks if the last emitted instruction is the given opcode.
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
	// Get the last emitted instruction's opcode.
	lastOp := OpCode(c.instructions[len(c.instructions)-1])
	// --- Debug Print: lastInstructionIs check result ---
	fmt.Printf("DEBUG (Compiler): Last instruction is %v. Check result: %t\n", lastOp, lastOp == op)
	// --- End Debug Print ---
	return lastOp == op
}

// removeLastPop removes the last emitted OpPop instruction.
// This is useful for expression statements that are the last statement in a block
// and whose value needs to be the result of the block/function.
func (c *Compiler) removeLastPop() {
	// Get the definition of OpPop to know its length (1 byte).
	// OpPop has no operands, so its size is always 1 byte.
	opPopSize := 1

	// Ensure the last instruction is indeed OpPop.
	if !c.lastInstructionIs(OpPop) {
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

// addConstant helper is defined above.
// emit helper is defined above.
// lastInstructionIs helper is defined above.
// removeLastPop helper is defined above.
// patchJump helper is defined above.
