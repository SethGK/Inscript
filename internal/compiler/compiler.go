package compiler

import (
	"fmt" // Needed for parsing integer literals
	// Needed for string concatenation

	"github.com/SethGK/Inscript/internal/ast" // Use your actual module path
	// You might need to import your runtime value types here later
	// "github.com/SethGK/Inscript/internal/vm/value"
)

// Compiler translates an AST into bytecode.
type Compiler struct {
	// The bytecode being compiled for the current compilation unit (program or function)
	instructions []Instruction
	constants    []interface{} // Constant pool for the current compilation unit

	// Global symbol table (unique for the entire program)
	// This is shared across all function compilers.
	globals *SymbolTable
	// We need to track the number of global variables defined for the VM.
	numGlobals int

	// Stack of symbol tables for managing nested scopes (functions, blocks)
	symbolTableStack []*SymbolTable
	currentScope     *SymbolTable // Pointer to the current active scope

	// Information about the current compilation unit (for functions)
	// This helps determine if we are compiling a function body
	// and manage local variable/parameter counts.
	currentBytecode *Bytecode // Pointer to the Bytecode being built for the current unit
}

// New creates a new top-level Compiler instance for a program.
func New() *Compiler {
	globalTable := NewSymbolTable() // Create the single global symbol table
	// The global table itself doesn't have an outer scope.
	// We'll use its `store` and manage `numDefinitions` separately for globals.

	// The initial currentBytecode is for the program itself.
	programBytecode := NewBytecode()

	c := &Compiler{
		instructions:     programBytecode.Instructions, // Start with the program's instruction slice
		constants:        programBytecode.Constants,    // Start with the program's constant slice
		globals:          globalTable,                  // The global table
		numGlobals:       0,                            // Start with 0 globals
		symbolTableStack: []*SymbolTable{globalTable},  // Start with global scope on stack
		currentScope:     globalTable,
		currentBytecode:  programBytecode, // Compiling the main program initially
	}

	// Pre-define builtins if any (e.g., "print" could be a builtin function)
	// For now, let's assume print is handled by OpPrint directly and not a callable builtin.
	// If you add builtins, define them in the global table here:
	// c.DefineGlobal("print", Builtin) // Example

	return c
}

// newFunctionCompiler creates a compiler instance for a function body.
// It inherits the global symbol table but has its own bytecode buffer
// and a symbol table nested within the parent compiler's current scope.
func (c *Compiler) newFunctionCompiler() *Compiler {
	// Create a new symbol table enclosed in the parent's *current* scope.
	// This captures variables from outer scopes (for closures).
	funcScope := NewEnclosedSymbolTable(c.currentScope)

	// Each function has its own bytecode and constant pool.
	funcBytecode := NewBytecode()

	// Create a new compiler instance for this function's body.
	// It shares the global symbol table but has its own instruction/constant buffers.
	funcComp := &Compiler{
		instructions:     funcBytecode.Instructions, // Start with the function's instruction slice
		constants:        funcBytecode.Constants,    // Start with the function's constant slice
		globals:          c.globals,                 // Share the global table
		numGlobals:       c.numGlobals,              // Inherit global count (will be updated in parent)
		symbolTableStack: []*SymbolTable{funcScope}, // Stack starts with function scope
		currentScope:     funcScope,
		currentBytecode:  funcBytecode, // Compiling this function's body
	}
	return funcComp
}

// EnterScope pushes a new nested scope onto the stack.
// isFunctionScope is true if entering a function body's scope.
func (c *Compiler) EnterScope(isFunctionScope bool) {
	var newTable *SymbolTable
	if isFunctionScope {
		// Function scope is enclosed in the parent compiler's *current* scope.
		// This is handled when newFunctionCompiler is called.
		// This method should be called *after* creating the new function compiler
		// and setting its initial scope.
		// For the main compiler, this method is used for block scopes.
		// panic("EnterScope(true) should not be called on the main compiler instance") // Re-enable this check if needed
		// For block scopes:
		newTable = NewEnclosedSymbolTable(c.currentScope)
	} else {
		// Block scope is enclosed in the current active scope.
		newTable = NewEnclosedSymbolTable(c.currentScope)
	}
	c.symbolTableStack = append(c.symbolTableStack, newTable)
	c.currentScope = newTable
}

// LeaveScope pops the current scope from the stack.
func (c *Compiler) LeaveScope() {
	if len(c.symbolTableStack) <= 1 {
		// Should not pop the global scope from the stack using this method.
		// The global scope is the base of the stack.
		panic("attempting to leave global scope using LeaveScope")
	}
	c.symbolTableStack = c.symbolTableStack[:len(c.symbolTableStack)-1]
	c.currentScope = c.symbolTableStack[len(c.symbolTableStack)-1]
}

// DefineGlobal defines a symbol in the top-level global symbol table.
// Returns the created Symbol.
func (c *Compiler) DefineGlobal(name string, kind SymbolKind) *Symbol {
	// Check if it already exists in the global table
	symbol, ok := c.globals.store[name]
	if ok {
		// If it exists, return the existing symbol.
		// This handles cases where a global is used before being assigned.
		return symbol
	}

	// Define a new global symbol
	symbol = &Symbol{Name: name, Kind: kind, Index: c.numGlobals}
	c.globals.store[name] = symbol
	c.numGlobals++ // Increment the global count
	return symbol
}

// Compile compiles the AST program into bytecode.
// This method is called on the top-level Compiler instance.
func (c *Compiler) Compile(program *ast.Program) (*Bytecode, error) {
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
	c.emit(OpNull)
	c.emit(OpReturn) // Return from the program's "main" execution context

	// Finalize the program's bytecode object.
	// The instructions and constants are already in c.instructions and c.constants.
	c.currentBytecode.Instructions = c.instructions
	c.currentBytecode.Constants = c.constants
	c.currentBytecode.NumLocals = c.currentScope.NumDefinitions() // Num locals at global scope is 0 usually
	c.currentBytecode.NumParameters = 0                           // Program has no parameters

	// The number of global variables is stored in c.numGlobals.
	// You might need to add this to the Bytecode struct or pass it to the VM separately.
	// Let's add a field to the Bytecode struct for the total number of globals.
	// This requires modifying the Bytecode struct definition. (Done in code.go)
	// c.currentBytecode.NumGlobals = c.numGlobals // Add this field to Bytecode

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
		c.emit(OpPop)

	case *ast.AssignStmt:
		// Compile the value expression first. It leaves the value on the stack.
		if err := c.compileExpression(stmt.Value); err != nil {
			return err
		}

		// Handle the assignment target based on its type.
		switch target := stmt.Target.(type) {
		case *ast.Identifier:
			// Assignment to a variable (identifier).
			// Look up the symbol in the current and outer scopes.
			symbol, ok := c.currentScope.Resolve(target.Name)
			if !ok {
				// If the variable is not found in any scope, define it as a Global.
				// This implements implicit global variable declaration on first assignment.
				// A language requiring explicit 'var' would handle this differently.
				symbol = c.DefineGlobal(target.Name, Global) // Define in the global table
			}

			// Emit the appropriate instruction based on the symbol's kind.
			switch symbol.Kind {
			case Global:
				// Assign to a global variable. Operand is the global index.
				c.emit(OpSetGlobal, symbol.Index)
			case Local, Parameter:
				// Assign to a local variable or parameter. Operand is the local/parameter index within the frame.
				c.emit(OpSetLocal, symbol.Index)
			case Builtin:
				// Cannot assign to a builtin.
				return fmt.Errorf("cannot assign to builtin '%s'", symbol.Name)
			default:
				return fmt.Errorf("unsupported symbol kind for assignment target: %v", symbol.Kind)
			}

		case *ast.IndexExpr:
			// Assignment to an indexed element (e.g., `my_list[index] = value`).
			// The value is already on the stack from compiling stmt.Value.
			// Compile the aggregate (e.g., `my_list`). This leaves the aggregate on the stack below the value.
			if err := c.compileExpression(target.Primary); err != nil {
				return err
			}
			// Compile the index expression (e.g., integer for list, string/value for table). Leaves index on stack.
			if err := c.compileExpression(target.Index); err != nil {
				return err
			}
			// Stack top is now [..., aggregate, index, value].
			// Emit OpSetIndex. The VM should pop value, index, aggregate, perform the set operation,
			// and typically push the updated aggregate back onto the stack (or the assigned value, depending on language semantics).
			c.emit(OpSetIndex) // Requires OpSetIndex opcode and VM implementation.

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
		c.emit(OpPrint, len(stmt.Exprs))

	case *ast.ReturnStmt:
		// Compile the return value expression, if present.
		if stmt.Expr != nil {
			if err := c.compileExpression(stmt.Expr); err != nil {
				return err
			}
			// Emit OpReturnValue if a value is returned.
			c.emit(OpReturnValue)
		} else {
			// Emit OpReturn if no value is returned (implicitly returns null).
			c.emit(OpReturn)
		}
		// Note: A return statement typically causes an immediate jump out of the function.
		// The VM handles this when it encounters OpReturnValue or OpReturn.
		// No explicit jump instruction is needed here in the compiler after the return opcode.

	case *ast.BlockStmt:
		// Delegate compilation of the block to the compileBlockStmt method.
		return c.compileBlockStmt(stmt)

	case *ast.IfStmt:
		return c.compileIf(stmt) // Delegate to helper for if-elseif-else logic

	case *ast.WhileStmt:
		return c.compileWhile(stmt) // Delegate to helper for while loop logic

	case *ast.ForStmt:
		return c.compileFor(stmt) // Delegate to helper for for loop logic (needs implementation)

	case *ast.FuncDefStmt:
		return c.compileFuncDef(stmt) // Delegate to helper for function definition (needs implementation)

	default:
		return fmt.Errorf("unsupported statement type: %T", s)
	}
	return nil
}

// compileBlockStmt compiles a block of statements.
// It handles entering and leaving the block's scope.
func (c *Compiler) compileBlockStmt(block *ast.BlockStmt) error {
	// Enter a new scope for the block.
	// Block scopes are not function scopes.
	c.EnterScope(false)

	// Compile all statements within the block.
	for _, st := range block.Stmts {
		if err := c.compileStatement(st); err != nil {
			// Ensure we leave the scope on error before returning.
			c.LeaveScope()
			return err
		}
	}

	// Leave the block scope.
	c.LeaveScope()

	return nil
}

// compileExpression emits bytecode for an expression node.
func (c *Compiler) compileExpression(e ast.Expression) error {
	switch expr := e.(type) {
	case *ast.IntegerLiteral:
		// Add the integer value to the constant pool and emit OpConstant.
		c.emitConstant(expr.Value)

	case *ast.FloatLiteral:
		// Add the float value to the constant pool and emit OpConstant.
		c.emitConstant(expr.Value)

	case *ast.StringLiteral:
		// Add the string value to the constant pool and emit OpConstant.
		c.emitConstant(expr.Value)

	case *ast.BooleanLiteral:
		// Emit OpTrue or OpFalse directly.
		if expr.Value {
			c.emit(OpTrue) // Assuming OpTrue exists
		} else {
			c.emit(OpFalse) // Assuming OpFalse exists
		}

	case *ast.NilLiteral:
		// Emit OpNull directly.
		c.emit(OpNull)

	case *ast.Identifier:
		// Look up the symbol (variable) in the current and outer scopes.
		symbol, ok := c.currentScope.Resolve(expr.Name)
		if !ok {
			// If the symbol is not found in any scope (including globals), it's an error.
			return fmt.Errorf("undefined variable '%s'", expr.Name)
		}

		// Emit the appropriate instruction based on the symbol's kind.
		switch symbol.Kind {
		case Global:
			// Get value of a global variable. Operand is the global index.
			c.emit(OpGetGlobal, symbol.Index)
		case Local, Parameter:
			// Get value of a local variable or parameter. Operand is the local/parameter index.
			c.emit(OpGetLocal, symbol.Index)
		case Builtin:
			// Getting a builtin usually means pushing the builtin function object onto the stack.
			// If builtins are stored in the global table, OpGetGlobal is appropriate.
			c.emit(OpGetGlobal, symbol.Index) // Assuming builtins are in the global table
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
			jumpFalsePos := c.emit(OpJumpNotTruthy, 0) // Placeholder jump target

			// If the left side was true, pop its value (we only need the right side's value for the result).
			c.emit(OpPop)

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil {
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpFalsePos)
			// The result of the expression (either Left's falsy value or Right's value) is now on the stack.

		} else if expr.Operator == "or" {
			// Compile the left side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Left); err != nil {
				return err
			}
			// If the left side is true (or not nil), jump over the right side compilation.
			// The truthy value is already on the stack and is the result of the 'or' if short-circuited.
			jumpTruePos := c.emit(OpJumpTruthy, 0) // Placeholder jump target

			// If the left side was false, pop its value (we only need the right side's value for the result).
			c.emit(OpPop)

			// Compile the right side. It leaves a boolean value on the stack.
			if err := c.compileExpression(expr.Right); err != nil {
				return err
			}

			// Patch the jump instruction to point to the instruction immediately after the right side compilation.
			c.patchJump(jumpTruePos)
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
				c.emit(OpAdd)
			case "-":
				c.emit(OpSub)
			case "*":
				c.emit(OpMul)
			case "/":
				c.emit(OpDiv)
			case "%":
				c.emit(OpMod)
			case "^":
				c.emit(OpPow)
			case "==":
				c.emit(OpEqual)
			case "!=":
				c.emit(OpNotEqual)
			case "<":
				c.emit(OpLessThan)
			case ">":
				c.emit(OpGreaterThan)
			case "<=":
				c.emit(OpLessEqual)
			case ">=":
				c.emit(OpGreaterEqual)
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
			c.emit(OpMinus)
		case "not":
			c.emit(OpNot)
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
		c.emit(OpArray, count)

	case *ast.TableLiteral:
		// Compile each field (key and value) of the table.
		// Keys and values are pushed onto the stack alternatingly: key1, value1, key2, value2, ...
		count := len(expr.Fields)
		for _, f := range expr.Fields {
			// Table keys are identifiers in the grammar, but typically evaluated to strings or other hashable values.
			// Assuming keys are treated as strings in the bytecode/VM.
			// Add the key string to the constant pool and emit OpConstant.
			c.emitConstant(f.Key)
			// Compile the value expression.
			if err := c.compileExpression(f.Value); err != nil {
				return err
			}
		}
		// Emit OpHash with the number of fields (key-value pairs).
		// The VM will pop 2*count values from the stack (key, value pairs) and create a hash/table.
		c.emit(OpHash, count)

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
		c.emit(OpIndex) // Requires OpIndex opcode and VM implementation.

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
		c.emit(OpCall, argCount) // Requires OpCall opcode and VM implementation.

	default:
		// This case should ideally not be reached.
		return fmt.Errorf("unsupported expression type: %T", e)
	}
	return nil
}

// compileIf handles if-elseif-else statements.
func (c *Compiler) compileIf(i *ast.IfStmt) error {
	// Compile the main condition. Leaves a boolean on the stack.
	if err := c.compileExpression(i.Cond); err != nil {
		return err
	}

	// If the condition is false, jump to the start of the first elseif or the else block (or the end if none).
	// We emit a placeholder jump instruction and store its position to patch later.
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 0) // Placeholder target 0

	// Compile the 'then' block (the body of the if statement).
	// Use the compileBlockStmt method.
	if err := c.compileBlockStmt(i.Then); err != nil {
		return err
	}

	// After the 'then' block, unconditionally jump to the end of the entire if-elseif-else chain.
	// We emit another placeholder jump.
	jumpEndPos := c.emit(OpJump, 0) // Placeholder target 0

	// Patch the OpJumpNotTruthy from the condition: its target is the instruction *after* the 'then' block.
	c.patchJump(jumpNotTruthyPos) // Patch to the instruction right after the OpJumpEnd

	// Compile the elseif chain.
	// Collect all jumps from elseif bodies that should go to the final end.
	jumpToEndPositions := []int{jumpEndPos} // Start with the jump from the 'then' block

	for _, elif := range i.ElseIfs {
		// Compile the elseif condition. Leaves a boolean on the stack.
		if err := c.compileExpression(elif.Cond); err != nil {
			return err
		}
		// If the elseif condition is false, jump to the start of the next elseif or else block.
		// Emit a placeholder jump.
		pos := c.emit(OpJumpNotTruthy, 0) // Placeholder target 0

		// Compile the elseif body.
		// Use the compileBlockStmt method.
		if err := c.compileBlockStmt(elif.Body); err != nil {
			return err
		}

		// After the elseif body, unconditionally jump to the end of the entire chain.
		// Emit another placeholder jump and add its position to the list of jumps to patch later.
		currentElseIfJumpEnd := c.emit(OpJump, 0) // Emit jump after elseif body
		jumpToEndPositions = append(jumpToEndPositions, currentElseIfJumpEnd)

		// Patch the OpJumpNotTruthy from the elseif condition: its target is the instruction *after* this elseif body's jump.
		c.patchJump(pos) // Patch to the instruction right after the OpJumpEnd for this elseif
	}

	// Compile the optional else block.
	if i.Else != nil {
		// If there's an else block, its code starts immediately after the last elseif (or the then block if no elseifs).
		// Use the compileBlockStmt method.
		if err := c.compileBlockStmt(i.Else); err != nil {
			return err
		}
	}
	// If there's no else block, the code continues here, which is the target for the jumps that skipped the else.

	// Patch all collected OpJump instructions (from the 'then' block and each elseif body)
	// to the instruction *after* the entire if-elseif-else chain.
	for _, jumpPos := range jumpToEndPositions {
		c.patchJump(jumpPos) // Patch all collected jumps to the final end
	}

	return nil
}

// compileWhile handles while loops.
func (c *Compiler) compileWhile(w *ast.WhileStmt) error {
	// Mark the starting position of the loop condition check.
	// Jumps from the body will return here.
	loopStartPos := len(c.instructions)

	// Compile the loop condition. Leaves a boolean on the stack.
	if err := c.compileExpression(w.Cond); err != nil {
		return err
	}

	// If the condition is false, jump out of the loop.
	// Emit a placeholder jump and store its position.
	jumpNotTruthy := c.emit(OpJumpNotTruthy, 0) // Placeholder target 0

	// Compile the loop body.
	// Use the compileBlockStmt method.
	if err := c.compileBlockStmt(w.Body); err != nil {
		return err
	}

	// After executing the body, unconditionally jump back to the start of the loop condition check.
	c.emit(OpJump, loopStartPos)

	// Patch the OpJumpNotTruthy instruction: its target is the instruction immediately after this jump back.
	c.patchJump(jumpNotTruthy) // Patch to the instruction right after the OpJump to loopStartPos

	return nil
}

// compileFor handles for loops (for IDENTIFIER 'in' expression block).
// This requires support for iterables and iterators in the VM.
func (c *Compiler) compileFor(f *ast.ForStmt) error {
	// TODO: Implement for loop compilation.
	// This is complex and depends on how you handle iterables and iterators in your VM.
	// Typical approach:
	// 1. Compile the Iterable expression (f.Iterable). It should result in an iterable object on the stack.
	if err := c.compileExpression(f.Iterable); err != nil {
		return err
	}

	// 2. Emit an instruction to get an iterator from the iterable (e.g., OpGetIterator). Leaves iterator on stack.
	c.emit(OpGetIterator) // Requires OpGetIterator and VM support

	// Mark the start of the loop body/check:
	loopStartPos := len(c.instructions)

	// 3. Emit an instruction to get the next value from the iterator (e.g., OpIteratorNext).
	//    This instruction typically pushes the next value and a boolean (true if successful, false if done) onto the stack.
	c.emit(OpIteratorNext) // Requires OpIteratorNext and VM support

	// 4. Pop the boolean result of OpIteratorNext.
	// 5. Emit OpJumpIfFalse to jump out of the loop if iteration is done. Patch target later.
	//    The boolean is the top element after OpIteratorNext.
	jumpEndPos := c.emit(OpJumpNotTruthy, 0) // Placeholder target 0
	// The value is now below the boolean on the stack.

	// 6. Pop the next value (the iterated element).
	//    The VM's OpIteratorNext pushes value THEN boolean. So after OpJumpNotTruthy (which consumes boolean),
	//    the value is at the top.
	//    c.emit(OpPop) // Pop the value - actually, we need to use it!

	// 7. Define the loop variable (`f.Variable`) in the current scope (it's a local variable).
	//    // Loop variables are typically local to the loop body's scope.
	//    // We need to enter a scope before compiling the body and define the variable there.
	//    // This requires restructuring how the loop variable is handled relative to the block scope.
	//    // A common pattern: the loop variable is defined *outside* the block scope, but its value is set inside.
	//    // Let's assume for now the loop variable is defined in the scope *containing* the for loop.
	//    // A better approach: the loop variable is local to the *implicit* scope of the loop body.
	//    // Let's define it as a local within the scope entered for the block.

	//    // Enter a new scope for the loop body (this scope contains the loop variable).
	c.EnterScope(false) // Not a function scope

	//    // Define the loop variable symbol in the *newly entered* scope.
	loopVarSymbol := c.currentScope.Define(f.Variable, Local)
	//    // The index is assigned by Define.

	//    // 8. Emit OpSetLocal to store the current iterated value into the loop variable.
	//    //    The iterated value is on top of the stack after OpIteratorNext and OpJumpNotTruthy.
	c.emit(OpSetLocal, loopVarSymbol.Index) // Use the index assigned by Define

	// 9. Compile the loop body (`f.Body`).
	// Use the compileBlockStmt method.
	if err := c.compileBlockStmt(f.Body); err != nil {
		// Leave the scope on error
		c.LeaveScope()
		return err
	}

	//    // The result of the body is likely on the stack; pop it if it's an expression statement.
	//    // The compileStatement for ExprStmt already emits OpPop, so this might not be needed here.

	// 10. Emit OpJump back to the loop start check.
	c.emit(OpJump, loopStartPos)

	// 11. Leave the loop body scope.
	c.LeaveScope()

	// 12. Patch the exit jump.
	c.patchJump(jumpEndPos) // Patch to the instruction right after the OpPop iterator

	// 13. Clean up the iterator on the stack (e.g., emit OpPop after the loop).
	c.emit(OpPop) // Pop the iterator object

	return nil // For loop compilation implemented
}

// compileFuncDef handles function definitions (function(params) { body }).
// This requires creating a separate compilation unit (Bytecode) for the function body.
func (c *Compiler) compileFuncDef(f *ast.FuncDefStmt) error {
	// 1. Create a new compiler instance for the function body.
	// This new compiler is enclosed in the *current* scope of the parent compiler.
	funcCompiler := c.newFunctionCompiler()

	// 2. Enter the function's scope in the new compiler.
	// This creates the symbol table for parameters and local variables.
	funcCompiler.EnterScope(true) // Mark as function scope

	// 3. Define parameters in the function's symbol table. Parameters are locals at the start of the frame.
	// The index of a parameter corresponds to its position in the argument list (0-based).
	for i, paramName := range f.Params {
		// Define the parameter in the function's current scope.
		paramSymbol := funcCompiler.currentScope.Define(paramName, Parameter)
		// Set the index to its position (0-based).
		paramSymbol.Index = i
	}
	// Store the number of parameters in the function's bytecode.
	funcCompiler.currentBytecode.NumParameters = len(f.Params)

	// 4. Compile the function body using the new compiler instance.
	// The function body is a BlockStmt.
	// Use the compileBlockStmt method on the function compiler instance.
	if err := funcCompiler.compileBlockStmt(f.Body); err != nil {
		// Leave the scope on error
		funcCompiler.LeaveScope()
		// Handle error, potentially propagate it from the sub-compiler.
		return err
	}

	// 5. Add an implicit return null at the end of the function body if it doesn't end with return.
	// Check the last instruction emitted by the function compiler.
	lastInstrIndex := len(funcCompiler.instructions) - 1
	if lastInstrIndex < 0 || (funcCompiler.instructions[lastInstrIndex].Op != OpReturnValue && funcCompiler.instructions[lastInstrIndex].Op != OpReturn) {
		// If the last instruction is not a return, add an implicit return null.
		funcCompiler.emit(OpNull)
		funcCompiler.emit(OpReturn)
	}

	// 6. Leave the function's scope.
	funcCompiler.LeaveScope()

	// 7. Finalize the function's bytecode object.
	// The instructions and constants are in funcCompiler.instructions and funcCompiler.constants.
	funcCompiler.currentBytecode.Instructions = funcCompiler.instructions
	funcCompiler.currentBytecode.Constants = funcCompiler.constants
	// Store the number of local variables defined within the function's scope (including parameters).
	// The symbol table's numDefinitions counts both parameters and locals defined within the body.
	// The VM needs the total number of slots for locals (parameters + locals).
	funcCompiler.currentBytecode.NumLocals = funcCompiler.currentScope.NumDefinitions() // Total slots needed for frame

	// 8. Create a Function Object runtime value. This object needs:
	//    - The compiled funcBytecode
	//    - Number of parameters (already stored in funcBytecode)
	//    - Number of local variables (already stored in funcBytecode)
	//    - (If supporting closures) References to variables from parent scopes that the function uses.
	//      This is complex and requires analyzing which variables from outer scopes are referenced
	//      within the function body and capturing them. This is often done by the compiler
	//      identifying "free variables" and the VM/runtime creating a "closure" object.
	//      For a basic implementation without closures, you might just need the bytecode, params, and locals count.

	//    Let's define a simple placeholder Function object struct for now.
	//    You'll need to define this in your VM/runtime value types package.
	type Function struct {
		Bytecode      *Bytecode
		NumParameters int
		NumLocals     int // Total slots needed for the frame (params + locals)
		// FreeVariables []*value.Value // For closures (complex)
	}
	// Create the function object.
	funcObject := &Function{
		Bytecode:      funcCompiler.currentBytecode,
		NumParameters: funcCompiler.currentBytecode.NumParameters,
		NumLocals:     funcCompiler.currentBytecode.NumLocals,
	}

	// 9. Add the Function Object to the constant pool of the *outer* compiler (the one compiling the FuncDefStmt).
	// This makes the function object available as a constant that can be pushed onto the stack.
	funcObjConstantIndex := c.addConstant(funcObject) // Use the parent compiler's constant pool

	// 10. Emit OpConstant in the outer compiler to push the Function Object onto the stack.
	// When the program execution reaches this point, it pushes the function object,
	// which can then be assigned to a variable or immediately called.
	c.emit(OpConstant, funcObjConstantIndex)

	// If the function definition is part of an assignment (e.g., `myFunc = function() { ... }`),
	// the assignment logic will handle storing the function object (which is now on top of the stack)
	// into the variable. If it's a standalone function definition (anonymous function not assigned),
	// the OpPop after an expression statement will discard the function object.

	return nil // Function definition compilation implemented
}

// emit adds an instruction to the bytecode being compiled for the current unit.
// It returns the starting position (index) of the instruction.
func (c *Compiler) emit(op OpCode, operands ...int) int {
	// TODO: Add operand encoding/decoding logic based on Definition if using []byte
	// For now, simple integer operands are stored directly in the Instruction struct
	instr := Instruction{Op: op, Operands: operands}
	pos := len(c.instructions)
	c.instructions = append(c.instructions, instr)
	return pos
}

// emitConstant adds a constant value to the constant pool of the current unit
// and emits OpConstant instruction.
func (c *Compiler) emitConstant(val interface{}) {
	// TODO: Check if constant already exists to avoid duplicates (Optimization)
	c.constants = append(c.constants, val)
	c.emit(OpConstant, len(c.constants)-1)
}

// patchJump sets the operand of a previously emitted jump instruction
// to the current position in the instruction stream.
// This is used to fix forward jumps (e.g., in if/while statements).
func (c *Compiler) patchJump(pos int) {
	// The target address is the instruction *after* the current end of the instructions.
	target := len(c.instructions)
	if pos < 0 || pos >= len(c.instructions) || len(c.instructions[pos].Operands) == 0 {
		// Basic validation for the position and if it has an operand.
		fmt.Printf("WARNING: Attempted to patch invalid jump at position %d\n", pos)
		return
	}
	// Assumes the jump target operand is the first (and usually only) operand.
	c.instructions[pos].Operands[0] = target
}

// addConstant adds a constant to the current compilation unit's constant pool.
// This is a helper used internally by emitConstant and compileFuncDef.
func (c *Compiler) addConstant(obj interface{}) int {
	// TODO: Check for duplicates
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

// GetBytecode returns the compiled bytecode for the current compilation unit.
// This is called after the top-level Compile finishes (for the program)
// or after a function body is compiled (by the parent compiler).
func (c *Compiler) GetBytecode() *Bytecode {
	// Ensure the Bytecode struct holds the final instructions and constants.
	c.currentBytecode.Instructions = c.instructions
	c.currentBytecode.Constants = c.constants
	// NumLocals and NumParameters should be set during function compilation.
	// For the main program, NumLocals is typically 0.

	return c.currentBytecode
}

// NumGlobalVariables returns the total number of global variables defined in the program.
// This is only meaningful for the top-level compiler instance after program compilation.
func (c *Compiler) NumGlobalVariables() int {
	// The number of globals is tracked separately as numGlobals.
	return c.numGlobals
}
