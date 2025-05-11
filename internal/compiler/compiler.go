// internal/compiler/compiler.go
package compiler

import (
	"fmt"
	// Needed for parsing integer literals
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

// NumGlobalVariables returns the total number of global variables defined.
func (c *Compiler) NumGlobalVariables() int {
	return c.numGlobals
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
		// Delegate to helper for if-elseif-else logic
		return c.compileIf(stmt)

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
	// TODO: Implement while loop compilation.
	// Similar to if, but the jump at the end of the body goes back to the condition.
	return fmt.Errorf("While loop compilation not yet implemented")
}

// compileFor handles for loops (for IDENTIFIER 'in' expression block).
// This requires support for iterables and iterators in the VM.
func (c *Compiler) compileFor(f *ast.ForStmt) error {
	// TODO: Implement for loop compilation.
	return fmt.Errorf("For loop compilation not yet implemented")
}

// compileFuncDef handles function definitions (function(params) { body }).
// This requires creating a separate compilation unit (Bytecode) for the function body.
func (c *Compiler) compileFuncDef(f *ast.FuncDefStmt) error {
	// TODO: Implement function definition compilation.
	return fmt.Errorf("Function definition compilation not yet implemented")
}

// emit adds an instruction to the bytecode.
// It returns the starting position of the emitted instruction.
func (c *Compiler) emit(op OpCode, operands ...int) int {
	// Get the definition of the opcode to know operand widths.
	def, ok := Lookup(op)
	if !ok {
		// This indicates a compiler bug: trying to emit an undefined opcode.
		panic(fmt.Sprintf("undefined opcode %d", op)) // Or return an error
	}

	// Calculate the instruction length (opcode + operands).
	instructionLen := 1 // Opcode byte
	for _, width := range def.OperandWidths {
		instructionLen += width
	}

	// Create the instruction byte slice.
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op) // First byte is the opcode

	// Encode the operands into the instruction slice.
	offset := 1 // Start writing operands after the opcode
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 1:
			// 1-byte operand (uint8)
			instruction[offset] = byte(operand)
		case 2:
			// 2-byte operand (uint16)
			// Use big-endian order (most significant byte first).
			instruction[offset] = byte(operand >> 8) // High byte
			instruction[offset+1] = byte(operand)    // Low byte
		// Add cases for other operand widths if needed (e.g., 4 bytes).
		default:
			// This indicates a compiler bug: unknown operand width in definition.
			panic(fmt.Sprintf("unknown operand width %d for opcode %s", width, def.Name)) // Or return error
		}
		offset += width
	}

	// Store the starting position of the new instruction.
	pos := len(c.instructions)

	// Append the new instruction to the compiler's instruction slice.
	c.instructions = append(c.instructions, Instruction{Op: op, Operands: operands}) // Store parsed operands for now

	// If you were using a []byte slice for instructions, you would append the byte slice here.
	// c.instructions = append(c.instructions, instruction...)

	return pos
}

// emitConstant adds a constant to the constant pool and emits an OpConstant instruction.
// It returns the starting position of the emitted instruction.
func (c *Compiler) emitConstant(value interface{}) int {
	// Add the value to the constant pool.
	constantIndex := c.addConstant(value)

	// Emit an OpConstant instruction with the index of the constant.
	// OpConstant expects a 2-byte operand for the constant index.
	return c.emit(OpConstant, constantIndex)
}

// addConstant adds a value to the constant pool and returns its index.
func (c *Compiler) addConstant(value interface{}) int {
	// Check if the constant already exists in the pool to avoid duplicates.
	// This requires value equality checking, which can be complex for custom types.
	// For now, we'll just append, allowing duplicates. Optimization can be added later.

	// Store the index before appending.
	index := len(c.constants)

	// Append the value to the constant pool.
	c.constants = append(c.constants, value)

	return index
}

// patchJump updates the operand of a jump instruction at a given position.
// The operand is the target instruction address.
func (c *Compiler) patchJump(instrPos int) {
	// The target address is the position of the instruction *after* the jump.
	targetAddress := len(c.instructions)

	// Get the instruction to patch.
	instruction := &c.instructions[instrPos]

	// Get the definition to know the operand width (should be 2 for jump targets).
	def, ok := Lookup(instruction.Op)
	if !ok || len(def.OperandWidths) != 1 || def.OperandWidths[0] != 2 {
		// This indicates a compiler bug: trying to patch a non-jump instruction
		// or a jump instruction with an unexpected operand structure.
		panic(fmt.Sprintf("cannot patch instruction at %d: not a 2-byte jump operand", instrPos)) // Or return error
	}

	// Update the operand with the correct target address.
	instruction.Operands[0] = targetAddress

	// If you were using a []byte slice for instructions, you would need to
	// update the bytes in the slice directly at instrPos + 1 and instrPos + 2.
	/*
		// Ensure the byte slice is large enough
		if instrPos+3 > len(c.currentBytecode.Instructions) {
			panic("bytecode slice too short for patching") // Or return error
		}
		// Encode the target address (uint16) into the byte slice (big-endian).
		c.currentBytecode.Instructions[instrPos+1] = byte(targetAddress >> 8) // High byte
		c.currentBytecode.Instructions[instrPos+2] = byte(targetAddress)    // Low byte
	*/
}
