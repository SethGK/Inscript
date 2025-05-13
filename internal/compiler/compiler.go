// Package compiler compiles the Inscript AST into bytecode.
package compiler

import (
	"encoding/binary" // Needed for binary encoding/decoding in patchJump
	"fmt"
	"os" // Needed for Fprintf to stderr

	"github.com/SethGK/Inscript/internal/ast"
	"github.com/SethGK/Inscript/internal/types"
)

// Compiler translates an AST into bytecode.
type Compiler struct {
	// The bytecode being compiled for the current compilation unit (program or function)
	instructions Instructions  // Use Instructions type from this package
	constants    []types.Value // Constant pool for the current compilation unit - Referring to types.Value

	// Global symbol table (unique for the entire program)
	// This is shared across all function compilers.
	globals *SymbolTable // Referring to SymbolTable (defined in symboltable.go, same package)
	// Removed numGlobals field - the count is now in c.globals.numDefinitions

	// Scope management
	symbolTableStack []*SymbolTable // Referring to SymbolTable (defined in symboltable.go, same package)
	currentScope     *SymbolTable   // Referring to SymbolTable (defined in symboltable.go, same package)

	// Bytecode buffer for the current compilation unit
	currentBytecode *Bytecode // Use Bytecode type from this package
}

// New creates a new top-level Compiler instance for a program.
func New() *Compiler {
	globalTable := NewSymbolTable() // Referring to NewSymbolTable (defined in symboltable.go, same package) - isFunctionScope is false

	// The initial currentBytecode is for the program itself.
	programBytecode := NewBytecode() // Referring to NewBytecode (defined in code.go, same package)

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
	// c.DefineGlobal("print", Builtin) // Example - Referring to Builtin (defined in symboltable.go, same package)

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
	funcScope := NewEnclosedSymbolTable(c.currentScope, true) // Referring to NewEnclosedSymbolTable (defined in symboltable.go, same package)

	// Each function has its own bytecode and constant pool.
	funcBytecode := NewBytecode() // Referring to NewBytecode (defined in code.go, same package)

	// Create a new compiler instance for this function's body.
	// It shares the global symbol table and uses its numDefinitions for global count.
	funcComp := &Compiler{
		instructions: funcBytecode.Instructions, // Start with the function's instruction slice ([]byte)
		constants:    funcBytecode.Constants,    // Start with the function's constant slice ([]types.Value)
		globals:      c.globals,                 // Share the global table
		// numGlobals field removed
		symbolTableStack: []*SymbolTable{funcScope}, // Stack starts with function scope
		currentScope:     funcScope,
		currentBytecode:  funcBytecode, // Compiling this function's body
	}
	return funcComp
}

// EnterScope pushes a new nested scope onto the stack.
// isFunctionScope is true if entering a function body's scope.
// This method is now primarily used for block scopes within the main program.
func (c *Compiler) EnterScope(isFunctionScope bool) {
	// We only enter non-function scopes here (block scopes in main program).
	// Function scopes are handled by newFunctionCompiler.
	if isFunctionScope {
		panic("EnterScope(true) should not be called on the main compiler instance")
	}

	// Block scope is enclosed in the current active scope.
	// Pass false to indicate this is NOT a function scope.
	newTable := NewEnclosedSymbolTable(c.currentScope, false) // Referring to NewEnclosedSymbolTable (defined in symboltable.go, same package)

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
func (c *Compiler) DefineGlobal(name string, kind SymbolKind) *Symbol { // Referring to SymbolKind, Symbol (defined in symboltable.go, same package)
	// Check if it already exists in the global table
	symbol, ok := c.globals.store[name]
	if ok {
		// If it exists, return the existing symbol.
		// This handles cases where a global is used before being assigned.
		return symbol
	}

	// Define a new global symbol
	// The index is the current number of definitions in the global table.
	symbol = &Symbol{Name: name, Kind: kind, Index: c.globals.numDefinitions} // Referring to Symbol (defined in symboltable.go, same package)
	c.globals.store[name] = symbol
	c.globals.numDefinitions++ // Manually increment here since we are not using the SymbolTable.Define method directly for globals in this helper
	return symbol
}

// DefineLocal defines a symbol in the current scope's symbol table.
// Returns the created Symbol.
func (c *Compiler) DefineLocal(name string) *Symbol { // Referring to Symbol (defined in symboltable.go, same package)
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
	symbol := &Symbol{Name: name, Kind: Local, Index: c.currentScope.numDefinitions} // Referring to Symbol, Local (defined in symboltable.go, same package)
	c.currentScope.store[name] = symbol
	c.currentScope.numDefinitions++ // Manually increment here since we are not using the SymbolTable.Define method directly for locals in this helper
	return symbol
}

// Compile compiles the AST program into bytecode.
// This method is called on the top-level Compiler instance.
func (c *Compiler) Compile(program *ast.Program) (*Bytecode, error) { // Returning *Bytecode
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
	c.emit(OpNull)   // Use OpNull
	c.emit(OpReturn) // Use OpReturn

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
	switch stmt := s.(type) {
	case *ast.ExprStmt:
		// Compile the expression. It leaves its result on the stack.
		if err := c.compileExpression(stmt.Expr); err != nil {
			return err
		}
		// For an expression statement, the result is not used, so pop it.
		c.emit(OpPop) // Use OpPop

	case *ast.AssignStmt:
		// Handle the assignment target based on its type.
		switch target := stmt.Target.(type) {
		case *ast.Identifier:
			// Assignment to a variable (identifier).
			// Compile the value expression first. It leaves the value on the stack.
			if err := c.compileExpression(stmt.Value); err != nil {
				return err
			}

			// Look up the symbol in the current and outer scopes.
			symbol, ok := c.currentScope.Resolve(target.Name)
			if !ok {
				// If the variable is not found in any scope, define it.
				// Variables defined in the main program (including nested blocks) are globals.
				// Variables defined in function scopes are locals.
				if c.currentScope.isFunctionScope {
					// Define as Local in the current function scope.
					// Corrected call to Define: provide the SymbolKind.
					symbol = c.currentScope.Define(target.Name, Local) // Pass name and Local kind
				} else {
					// If not in a function scope (global or block in main program), define as Global.
					// Corrected call to Define: provide the SymbolKind.
					symbol = c.globals.Define(target.Name, Global) // Pass name and Global kind
				}
			}

			// Emit the appropriate instruction based on the symbol's kind.
			switch symbol.Kind {
			case Global: // Referring to Global (defined in symboltable.go, same package)
				// Assign to a global variable. Operand is the global index.
				c.emit(OpSetGlobal, symbol.Index) // Use OpSetGlobal
			case Local, Parameter: // Parameters are also accessed as locals within the function frame (defined in symboltable.go, same package)
				// Assign to a local variable or parameter. Operand is the local/parameter index within the frame.
				// This case should ONLY be reached if compiling within a function scope.
				if !c.currentScope.isFunctionScope {
					return fmt.Errorf("internal compiler error: attempting to set local/parameter outside function scope for '%s'", symbol.Name)
				}
				// The operand for locals is 1 byte (uint8).
				c.emit(OpSetLocal, symbol.Index) // Use OpSetLocal
			case Builtin: // Referring to Builtin (defined in symboltable.go, same package)
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
			if err := c.compileExpression(stmt.Value); err != nil { // Compile stmt.Value LAST
				return err
			}

			// Stack top is now [..., aggregate, index, value]. Correct order for OpSetIndex.
			// Emit OpSetIndex. The VM should pop value, index, aggregate, and perform the assignment.
			c.emit(OpSetIndex) // Use OpSetIndex

		// Add cases for other potential assignment targets if your grammar supports them,
		// e.g., field access like `my_object.field = value`.

		default:
			return fmt.Errorf("unsupported assignment target type: %T", stmt.Target)
		}

	case *ast.PrintStmt:
		// Compile each expression to be printed. They leave their results on the stack.
		for _, e := range stmt.Exprs {
			if err := c.compileExpression(e); err != nil {
				return err
			}
		}
		// Emit the print instruction with the number of arguments to print.
		c.emit(OpPrint, len(stmt.Exprs)) // Use OpPrint

	case *ast.ReturnStmt:
		// Compile the return value expression, if present.
		if stmt.Expr != nil {
			if err := c.compileExpression(stmt.Expr); err != nil {
				return err
			}
			// Emit OpReturnValue if a value is returned.
			c.emit(OpReturnValue) // Use OpReturnValue
		} else {
			// Emit OpReturn if no value is returned (implicitly returns null).
			c.emit(OpReturn) // Use OpReturn
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

	case *ast.FuncDefStmt: // Assuming this is the AST node for function literals (e.g., `function() { ... }`)
		// Create a new compiler specifically for this function's body.
		// This compiler will manage the function's local scope and bytecode.
		funcCompiler := c.newFunctionCompiler()

		// Define parameters as local variables in the function's scope.
		// Assuming ast.FuncDefStmt has a field like `Params []string`
		for _, paramName := range stmt.Params {
			// DefineLocal internally calls Define with Local kind and increments numDefinitions for the function scope.
			funcCompiler.DefineLocal(paramName)
		}

		// Compile the function body.
		// The compileBlockStmt method handles entering/leaving the block's scope.
		if err := funcCompiler.compileBlockStmt(stmt.Body); err != nil {
			return err // Propagate compilation errors from the function body
		}

		// After compiling the body, ensure the last instruction is a return.
		// If the block doesn't end with an explicit return, add an implicit return nil.
		// Use the helper to get the opcode of the last emitted instruction.
		lastOpcode := GetLastOpcode(funcCompiler.instructions)
		if lastOpcode != OpReturn && lastOpcode != OpReturnValue {
			funcCompiler.emit(OpNull)   // Push nil onto the stack
			funcCompiler.emit(OpReturn) // Return from the function
		}

		// Finalize the function's bytecode.
		// The instructions and constants are in funcCompiler.instructions and funcCompiler.constants.
		// The number of locals is the number of definitions in the function's scope (parameters + locals).
		funcBytecode := NewBytecode()
		funcBytecode.Instructions = funcCompiler.instructions
		funcBytecode.Constants = funcCompiler.constants
		funcBytecode.NumLocals = funcCompiler.currentScope.numDefinitions // Number of parameters + locals
		funcBytecode.NumParameters = len(stmt.Params)                     // Store the number of parameters
		// NumGlobals is 0 for function bytecode

		// Create the Function value object.
		// This object represents the compiled function at runtime. Now in types package.
		fn := &types.CompiledFunction{
			Instructions:  funcBytecode.Instructions, // Bytecode.Instructions is []byte
			NumParameters: len(stmt.Params),
			NumLocals:     funcBytecode.NumLocals, // Also store NumLocals in CompiledFunction
			// TODO: Add Free variables/closure information here later if implementing closures.
			// This would involve capturing symbols from outer scopes that are used in the function body.
		}

		// Wrap the CompiledFunction in a Closure value object. Now in types package.
		closure := &types.Closure{Fn: fn}

		// Add the compiled Closure object to the *current* compiler's constant pool (the compiler compiling the code *containing* the function definition).
		// This makes the function object available as a constant that can be pushed onto the stack.
		c.emitConstant(closure) // emitConstant should add the value to c.constants and emit OpConstant

		// The compiled expression (the function literal) leaves the function object on the stack.
		// The calling compileStatement (e.g., AssignStmt) will then handle what to do with this value (e.g., assign it to a variable).

		return nil // Successfully compiled the function literal statement (it's a statement that compiles an expression)

	default:
		// This case should ideally not be reached.
		return fmt.Errorf("unsupported statement type: %T", s)
	}
	return nil
}

// compileExpression emits bytecode for an expression node.
func (c *Compiler) compileExpression(e ast.Expression) error {
	// Defensive check: Ensure we are not trying to compile a statement as an expression.
	// This directly addresses the "ImpossibleAssert" error message.
	// This check is crucial because ast.FuncDefStmt is a Statement, not an Expression.
	if _, ok := e.(ast.Statement); ok {
		return fmt.Errorf("internal compiler error: attempting to compile a statement (%T) as an expression", e)
	}

	switch expr := e.(type) {
	case *ast.IntegerLiteral:
		// Add the integer value to the constant pool and emit OpConstant.
		c.emitConstant(types.NewInteger(expr.Value)) // Referring to types.NewInteger (from types package)

	case *ast.FloatLiteral:
		// Add the float value to the constant pool and emit OpConstant.
		c.emitConstant(types.NewFloat(expr.Value)) // Referring to types.NewFloat (from types package)

	case *ast.StringLiteral:
		// Add the string value to the constant pool and emit OpConstant.
		c.emitConstant(types.NewString(expr.Value)) // Referring to types.NewString (from types package)

	case *ast.BooleanLiteral:
		// Emit OpTrue or OpFalse directly.
		if expr.Value {
			c.emit(OpTrue) // Use OpTrue
		} else {
			c.emit(OpFalse) // Use OpFalse
		}

	case *ast.NilLiteral:
		// Emit OpNull directly.
		c.emit(OpNull) // Use OpNull

	case *ast.Identifier:
		// Look up the symbol (variable) in the current and outer scopes.
		symbol, ok := c.currentScope.Resolve(expr.Name)
		if !ok {
			// If the symbol is not found in any scope (including globals), it's an error.
			return fmt.Errorf("undefined variable '%s'", expr.Name)
		}

		// Emit the appropriate instruction based on the symbol's kind and scope.
		switch symbol.Kind {
		case Global: // Referring to Global (defined in symboltable.go, same package)
			// Get value of a global variable. Operand is the global index.
			c.emit(OpGetGlobal, symbol.Index) // Use OpGetGlobal
		case Local, Parameter: // Parameters are also accessed as locals within the function frame (defined in symboltable.go, same package)
			// Get value of a local variable or parameter. Operand is the local/parameter index.
			// This case should ONLY be reached if compiling within a function scope.
			if !c.currentScope.isFunctionScope {
				return fmt.Errorf("internal compiler error: attempting to get local/parameter outside function scope for '%s'", symbol.Name)
			}
			// The operand for locals is 1 byte (uint8).
			c.emit(OpGetLocal, symbol.Index) // Use OpGetLocal
		case Builtin: // Referring to Builtin (defined in symboltable.go, same package)
			// Getting a builtin usually means pushing the builtin function object onto the stack.
			// If builtins are stored in the global table, OpGetGlobal is appropriate.
			c.emit(OpGetGlobal, symbol.Index) // Use OpGetGlobal
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
			jumpFalsePos := c.emit(OpJumpNotTruthy, 0) // Use OpJumpNotTruthy

			// If the left side was true, pop its value (we only need the right side's value for the result).
			c.emit(OpPop) // Use OpPop

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
			jumpTruePos := c.emit(OpJumpTruthy, 0) // Use OpJumpTruthy

			// If the left side was false, pop its value (we only need the right side's value for the result).
			c.emit(OpPop) // Use OpPop

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil {
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpTruePos, len(c.instructions)) // Pass the target position (current length)
			// The result of the expression (either Left's truthy value or Right's value) is now on the stack.

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
				c.emit(OpAdd) // Use OpAdd
			case "-":
				c.emit(OpSub) // Use OpSub
			case "*":
				c.emit(OpMul) // Use OpMul
			case "/":
				c.emit(OpDiv) // Use OpDiv
			case "%":
				c.emit(OpMod) // Use OpMod
			case "^":
				c.emit(OpPow) // Use OpPow
			case "==":
				c.emit(OpEqual) // Use OpEqual
			case "!=":
				c.emit(OpNotEqual) // Use OpNotEqual
			case "<":
				c.emit(OpLessThan) // Use OpLessThan
			case ">":
				c.emit(OpGreaterThan) // Use OpGreaterThan
			case "<=":
				c.emit(OpLessEqual) // Use OpLessEqual
			case ">=":
				c.emit(OpGreaterEqual) // Use OpGreaterEqual
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
			c.emit(OpMinus) // Use OpMinus
		case "not":
			c.emit(OpNot) // Use OpNot
		case "+":
			// Unary plus is typically a no-op semantically, but we still compile the expression.
			// No opcode is needed unless you want to enforce type checks at runtime (e.g., ensure the operand is a number).
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
		c.emit(OpArray, count) // Use OpArray

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
		c.emit(OpHash, count) // Use OpHash

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
		c.emit(OpIndex) // Use OpIndex

	case *ast.CallExpr:
		// Compile the callee expression (the function to call). Leaves the function object on the stack.
		if err := c.compileExpression(expr.Callee); err != nil {
			return err
		}
		// Compile each argument expression. They leave their results on the stack in order.
		argCount := len(expr.Args)
		for _, a := range expr.Args {
			if err := c.compileExpression(a); err != nil {
				return err
			}
		}
		// Stack top is now [..., function_object, arg1, arg2, ..., argN].
		// Emit OpCall with the number of arguments.
		// The VM will pop the arguments and the function object, set up a new call frame, and execute the function's bytecode.
		c.emit(OpCall, argCount) // Use OpCall

	// Removed the case for *ast.FuncDefStmt from compileExpression
	// case *ast.FuncDefStmt: // This case should NOT be here

	// TODO: Add cases for other expression types as you implement them
	// case *ast.PrefixExpr: // Already handled by UnaryExpr
	// case *ast.InfixExpr: // Already handled by BinaryExpr
	// case *ast.IfExpr: // If expressions (ternary operator or similar)
	// case *ast.WhileExpr: // While expressions (if your language supports them)
	// case *ast.ForExpr: // For expressions (if your language supports them)
	// case *ast.TableAccessExpr: // For accessing fields of a table/object (e.g., obj.field)

	default:
		// If an expression type is not handled, return an error.
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
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 9999) // Use OpJumpNotTruthy

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
	// if c.lastInstructionIs(OpPop) { // Use OpPop
	// 	c.removeLastPop()
	// }

	// Emit OpJump with a placeholder operand to jump over the else-if and else blocks.
	jumpPos := c.emit(OpJump, 9999) // Use OpJump

	// Patch the OpJumpNotTruthy operand to point to the instruction after the 'then' block.
	afterThenPos := len(c.instructions) // Use c.instructions
	c.patchJump(jumpNotTruthyPos, afterThenPos)

	// Handle else-if blocks.
	for _, elseif := range stmt.ElseIfs {
		// Compile the elseif condition.
		if err := c.compileExpression(elseif.Cond); err != nil {
			return err
		}
		// Emit OpJumpNotTruthy for the elseif condition.
		elseifJumpNotTruthyPos := c.emit(OpJumpNotTruthy, 9999) // Use OpJumpNotTruthy

		// Compile the elseif body.
		if err := c.compileBlockStmt(elseif.Body); err != nil {
			return err
		}

		// Emit OpJump after elseif body to jump over remaining else-if/else blocks.
		elseifJumpPos := c.emit(OpJump, 9999) // Use OpJump

		// Patch the OpJumpNotTruthy for this elseif to point after its body.
		afterElseifBodyPos := len(c.instructions) // Use c.instructions
		c.patchJump(elseifJumpNotTruthyPos, afterElseifBodyPos)

		// The previous jump (either from the main 'then' block or a previous elseif)
		// should now jump to the instruction after this elseif block.
		c.patchJump(jumpPos, afterElseifBodyPos) // Update the main jump target

		// The next jump (from this elseif's body) will jump past the remaining blocks.
		jumpPos = elseifJumpPos
	}

	// Handle the 'else' block.
	if stmt.Else != nil {
		if err := c.compileBlockStmt(stmt.Else); err != nil {
			return err
		}
	} else {
		// If there is no 'else' block and the condition was false,
		// the jumpNotTruthy jumps here. We need to push a default value (like nil)
		// onto the stack for the if expression's result.
		c.emit(OpNull) // Use OpNull
	}

	// Patch the final jump (from the 'then' block or the last elseif) to point after the 'else' block (or the pushed nil).
	afterElsePos := len(c.instructions) // Use c.instructions
	c.patchJump(jumpPos, afterElsePos)

	return nil
}

// compileWhile compiles a while loop statement.
func (c *Compiler) compileWhile(stmt *ast.WhileStmt) error {
	// Mark the position before the condition for the loop's jump back.
	loopStartPos := len(c.instructions) // Use c.instructions

	// Compile the loop condition.
	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}

	// Emit OpJumpNotTruthy with a placeholder operand. This jumps out of the loop if the condition is false.
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 9999) // Use OpJumpNotTruthy

	// Compile the loop body.
	if err := c.compileBlockStmt(stmt.Body); err != nil {
		return err
	}

	// After the loop body, emit OpJump to jump back to the start of the loop condition.
	c.emit(OpJump, loopStartPos) // Use OpJump

	// Patch the OpJumpNotTruthy operand to point to the instruction immediately after the loop.
	afterLoopPos := len(c.instructions) // Use c.instructions
	c.patchJump(jumpNotTruthyPos, afterLoopPos)

	// The while loop itself doesn't leave a value on the stack.
	// If the language requires a value (e.g., for use in an expression context),
	// you might need to push a default value (like nil) here.
	// Assuming for now it's a statement and doesn't leave a value.

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
	c.EnterScope(c.currentScope.isFunctionScope) // Pass the function scope status of the outer scope

	var loopVarSymbol *Symbol // Use Symbol from this package

	// Define the loop variable in the current (loop's block) scope.
	// The kind (Global or Local) depends on whether the outer scope is a function scope.
	if c.currentScope.isFunctionScope {
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

// emit adds an instruction to the current compilation unit's instructions.
// It returns the starting position of the emitted instruction.
func (c *Compiler) emit(op OpCode, operands ...int) int { // Use OpCode
	ins := Make(op, operands...) // Use Make
	pos := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	// Note: For tracking last/previous instructions for optimizations like removing OpPop,
	// you would need to add logic here to update those fields in the Compiler struct
	// or the CompilationScope if using scopes for instruction tracking.
	return pos
}

// emitConstant adds a constant to the current compilation unit's constant pool
// and emits an OpConstant instruction.
func (c *Compiler) emitConstant(value types.Value) {
	// Add the value to the constant pool.
	constantIndex := len(c.constants)
	c.constants = append(c.constants, value)

	// Emit OpConstant instruction with the index of the added constant.
	c.emit(OpConstant, constantIndex) // Use OpConstant
}

// patchJump updates the operand of a jump instruction at a given position.
func (c *Compiler) patchJump(jumpPos int, targetPos int) {
	// The jump instruction is at jumpPos. The operand (the jump target)
	// starts after the opcode byte.
	opcode := OpCode(c.instructions[jumpPos]) // Use OpCode
	def, ok := Lookup(opcode)                 // Use Lookup
	if !ok || len(def.OperandWidths) == 0 {
		// This should not happen for jump instructions.
		fmt.Fprintf(os.Stderr, "ERROR: patchJump called on non-jump instruction at position %d\n", jumpPos)
		return
	}

	// Assuming jump operands are 2 bytes (uint16).
	if def.OperandWidths[0] != 2 {
		fmt.Fprintf(os.Stderr, "ERROR: patchJump called on jump instruction with non-2-byte operand at position %d\n", jumpPos)
		return
	}

	// Encode the target position as a 2-byte big-endian integer.
	targetBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(targetBytes, uint16(targetPos))

	// Copy the target bytes into the instruction slice at the correct offset.
	// The offset is 1 byte for the opcode.
	copy(c.instructions[jumpPos+1:], targetBytes)
}

// TODO: Implement helper functions like lastInstructionIs, removeLastPop, changeOperand
// if they are needed for optimizations or specific language features.
// The provided code snippet uses lastInstructionIs and removeLastPop in the previous version.
// We should add them back if they are required for the logic (e.g., in IfStmt compilation).

// lastInstructionIs checks if the last emitted instruction is of the given opcode.
// This requires tracking the position and opcode of the last instruction.
// We can add fields to the Compiler struct or CompilationScope for this.
// For simplicity now, let's re-implement based on the instruction slice directly,
// but a more efficient approach would track it during emit.
func (c *Compiler) lastInstructionIs(op OpCode) bool { // Use OpCode
	if len(c.instructions) == 0 {
		return false
	}
	// This is a simplified check that only works if the last instruction has no operands.
	// A proper implementation needs to parse the last instruction's length.
	// Using GetLastOpcode from the bytecode package is better.
	return GetLastOpcode(c.instructions) == op // Use GetLastOpcode
}

// removeLastPop removes the last OpPop instruction if it exists.
// This is needed for expression statements within contexts where the value is used (e.g., If expressions).
func (c *Compiler) removeLastPop() {
	// Find the last instruction's position and opcode.
	currentLen := len(c.instructions)
	if currentLen == 0 {
		return // Nothing to remove
	}

	// Iterate backward to find the start of the last instruction.
	// This requires parsing instruction lengths from the end.
	// A more efficient approach is to track the last instruction's position during emit.
	// For now, let's assume OpPop is always 1 byte and is the last instruction.
	// TODO: Implement proper backward parsing or track last instruction position.

	// Simplified assumption: last instruction is OpPop (1 byte)
	if c.instructions[currentLen-1] == byte(OpPop) { // Use OpPop
		c.instructions = c.instructions[:currentLen-1]
		// Need to update last/prev instruction tracking if implemented.
	}
}

// changeOperand changes the operand of an instruction at a given position.
// This is used by patchJump, but might be needed for other optimizations.
// It's already implemented within patchJump for jump instructions.
// If needed for other opcodes, a more general implementation would be required.
