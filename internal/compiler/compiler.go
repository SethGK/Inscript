// Package compiler compiles the Inscript AST into bytecode.
package compiler

import (
	"fmt"

	"github.com/SethGK/Inscript/internal/ast"
	"github.com/SethGK/Inscript/internal/types"
)

// Compiler translates an AST into bytecode.
type Compiler struct {
	instructions Instructions
	constants    []types.Value

	globals      *SymbolTable
	symbolStack  []*SymbolTable
	currentScope *SymbolTable

	returned bool

	loopJumpStack [][]int // [breakJumpPos, continueJumpPos]
}

// New creates a new top-level Compiler.
func New() *Compiler {
	global := NewSymbolTable()
	c := &Compiler{
		instructions:  make(Instructions, 0),
		constants:     make([]types.Value, 0),
		globals:       global,
		symbolStack:   []*SymbolTable{global},
		currentScope:  global,
		returned:      false,
		loopJumpStack: make([][]int, 0),
	}
	return c
}

// Compile compiles the AST root program and returns a Bytecode struct.
func (c *Compiler) Compile(program *ast.Program) (*Bytecode, error) {
	for _, stmt := range program.Stmts {
		if err := c.compileStatement(stmt); err != nil {
			return nil, err
		}
	}
	c.emit(OpNull)
	c.emit(OpReturn)

	bc := &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
		NumLocals:    0,
		NumGlobals:   c.globals.NumGlobalsInTable(),
	}
	return bc, nil
}

// compileStatement handles statements.
func (c *Compiler) compileStatement(s ast.Statement) error {
	if c.returned {
		return nil
	}
	switch stmt := s.(type) {
	case *ast.ExprStmt:
		if err := c.compileExpression(stmt.Expr); err != nil {
			return err
		}
		c.emit(OpPop)

	case *ast.AssignStmt:
		return c.compileAssignment(stmt)

	case *ast.PrintStmt:
		for _, e := range stmt.Exprs {
			if err := c.compileExpression(e); err != nil {
				return err
			}
		}
		c.emit(OpPrint, len(stmt.Exprs))

	case *ast.ReturnStmt:
		if stmt.Expr != nil {
			if err := c.compileExpression(stmt.Expr); err != nil {
				return err
			}
			c.emit(OpReturnValue)
		} else {
			c.emit(OpNull)
			c.emit(OpReturn)
		}
		c.returned = true

	case *ast.BlockStmt:
		c.enterScope(false)
		for _, s2 := range stmt.Stmts {
			if err := c.compileStatement(s2); err != nil {
				return err
			}
		}
		c.leaveScope()
		c.returned = false

	case *ast.IfStmt:
		return c.compileIf(stmt)
	case *ast.WhileStmt:
		return c.compileWhile(stmt)
	case *ast.ForStmt:
		return c.compileFor(stmt)
	case *ast.FunctionDef:
		return c.compileFuncDef(stmt)
	case *ast.BreakStmt:
		return c.compileBreak()
	case *ast.ContinueStmt:
		return c.compileContinue()
	case *ast.ImportStmt:
		return c.compileImport(stmt)
	default:
		return fmt.Errorf("unsupported statement: %T", s)
	}
	return nil
}

// compileAssignment handles assignment statements including compound assignments.
func (c *Compiler) compileAssignment(stmt *ast.AssignStmt) error {
	if ident, isIdent := stmt.Target.(*ast.Identifier); isIdent {
		sym, ok := c.currentScope.Resolve(ident.Name)
		if !ok {
			// If not found, define it in the current scope.
			// If currentScope is global, it's a global. Otherwise, it's local.
			if c.currentScope == c.globals {
				sym = c.globals.DefineGlobal(ident.Name)
			} else {
				sym = c.currentScope.DefineLocal(ident.Name)
			}
		}

		if stmt.Op.Literal != "=" { // Check for compound assignment
			switch sym.Kind {
			case Global:
				c.emit(OpGetGlobal, sym.Index)
			case Local, Parameter:
				c.emit(OpGetLocal, sym.Index)
			case Free:
				c.emit(OpGetFree, sym.Index)
			case Builtin:
				return fmt.Errorf("cannot assign to builtin '%s'", ident.Name)
			}

			if err := c.compileExpression(stmt.Value); err != nil {
				return err
			}

			switch stmt.Op.Literal {
			case "+=":
				c.emit(OpAdd)
			case "-=":
				c.emit(OpSub)
			case "*=":
				c.emit(OpMul)
			case "/=":
				c.emit(OpDiv)
			case "^^=":
				c.emit(OpPow)
			default:
				return fmt.Errorf("unsupported compound assignment operator: %s", stmt.Op.Literal)
			}
		} else { // Simple assignment
			if err := c.compileExpression(stmt.Value); err != nil {
				return err
			}
		}

		switch sym.Kind {
		case Global:
			c.emit(OpSetGlobal, sym.Index)
		case Local, Parameter:
			c.emit(OpSetLocal, sym.Index)
		case Free:
			c.emit(OpSetFree, sym.Index)
		default:
			return fmt.Errorf("cannot assign to %s %s", sym.Kind, ident.Name)
		}
		return nil
	}

	if indexTarget, isIndex := stmt.Target.(*ast.IndexExpr); isIndex {
		if err := c.compileExpression(indexTarget.Primary); err != nil {
			return err
		}
		if err := c.compileExpression(indexTarget.Index); err != nil {
			return err
		}
		if err := c.compileExpression(stmt.Value); err != nil {
			return err
		}
		c.emit(OpSetIndex)
		return nil
	}

	if attrTarget, isAttr := stmt.Target.(*ast.AttrExpr); isAttr {
		if err := c.compileExpression(attrTarget.Primary); err != nil {
			return err
		}
		c.emitConstant(types.NewString(attrTarget.Attribute))
		if err := c.compileExpression(stmt.Value); err != nil {
			return err
		}
		c.emit(OpSetIndex)
		return nil
	}

	return fmt.Errorf("unsupported assignment target: %T", stmt.Target)
}

// compileExpression handles expressions.
func (c *Compiler) compileExpression(e ast.Expression) error {
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
		sym, ok := c.currentScope.Resolve(expr.Name)
		if !ok {
			return fmt.Errorf("undefined variable '%s'", expr.Name)
		}
		switch sym.Kind {
		case Global:
			c.emit(OpGetGlobal, sym.Index)
		case Local, Parameter:
			c.emit(OpGetLocal, sym.Index)
		case Free:
			c.emit(OpGetFree, sym.Index)
		case Builtin:
			c.emit(OpGetGlobal, sym.Index)
		}

	case *ast.BinaryExpr:
		return c.compileBinaryExpression(expr)
	case *ast.UnaryExpr:
		return c.compileUnaryExpression(expr)
	case *ast.CallExpr:
		return c.compileCallExpression(expr)
	case *ast.IndexExpr:
		return c.compileIndexExpression(expr)
	case *ast.AttrExpr:
		return c.compileAttrExpression(expr)
	case *ast.ListLiteral:
		return c.compileListLiteral(expr)
	case *ast.TableLiteral:
		return c.compileTableLiteral(expr)
	case *ast.TupleLiteral:
		return c.compileTupleLiteral(expr) // Call a separate function for tuple
	default:
		return fmt.Errorf("unsupported expression: %T", e)
	}
	return nil
}

// compileBinaryExpression handles binary operators.
func (c *Compiler) compileBinaryExpression(expr *ast.BinaryExpr) error {
	if expr.Operator.Literal == "and" {
		if err := c.compileExpression(expr.Left); err != nil {
			return err
		}
		jumpFalsePos := c.emit(OpJumpNotTruthy, 0)
		c.emit(OpPop)
		if err := c.compileExpression(expr.Right); err != nil {
			return err
		}
		c.patchJump(jumpFalsePos, len(c.instructions))
		return nil
	}
	if expr.Operator.Literal == "or" {
		if err := c.compileExpression(expr.Left); err != nil {
			return err
		}
		jumpTruePos := c.emit(OpJumpTruthy, 0)
		c.emit(OpPop)
		if err := c.compileExpression(expr.Right); err != nil {
			return err
		}
		c.patchJump(jumpTruePos, len(c.instructions))
		return nil
	}

	if err := c.compileExpression(expr.Left); err != nil {
		return err
	}
	if err := c.compileExpression(expr.Right); err != nil {
		return err
	}

	switch expr.Operator.Literal {
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
	case "^^":
		c.emit(OpPow)
	case "//":
		c.emit(OpIDiv)
	case "&":
		c.emit(OpBitAnd)
	case "|":
		c.emit(OpBitOr)
	case "^":
		c.emit(OpBitXor)
	case "<<":
		c.emit(OpShl)
	case ">>":
		c.emit(OpShr)
	case "==":
		c.emit(OpEqual)
	case "!=":
		c.emit(OpNotEqual)
	case "<":
		c.emit(OpLessThan)
	case "<=":
		c.emit(OpLessEqual)
	case ">":
		c.emit(OpGreaterThan)
	case ">=":
		c.emit(OpGreaterEqual)
	default:
		return fmt.Errorf("unsupported binary operator: %s", expr.Operator.Literal)
	}
	return nil
}

// compileUnaryExpression handles unary operators.
func (c *Compiler) compileUnaryExpression(expr *ast.UnaryExpr) error {
	if err := c.compileExpression(expr.Expr); err != nil {
		return err
	}
	switch expr.Operator.Literal {
	case "-":
		c.emit(OpMinus)
	case "not":
		c.emit(OpBang)
	case "~":
		c.emit(OpBitNot)
	default:
		return fmt.Errorf("unsupported unary operator: %s", expr.Operator.Literal)
	}
	return nil
}

// compileCallExpression handles function calls.
func (c *Compiler) compileCallExpression(expr *ast.CallExpr) error {
	if err := c.compileExpression(expr.Callee); err != nil {
		return err
	}
	for _, arg := range expr.Args {
		if err := c.compileExpression(arg); err != nil {
			return err
		}
	}
	c.emit(OpCall, len(expr.Args))
	return nil
}

// compileIndexExpression handles list/table indexing.
func (c *Compiler) compileIndexExpression(expr *ast.IndexExpr) error {
	if err := c.compileExpression(expr.Primary); err != nil {
		return err
	}
	if err := c.compileExpression(expr.Index); err != nil {
		return err
	}
	c.emit(OpIndex)
	return nil
}

// compileAttrExpression handles table attribute access.
func (c *Compiler) compileAttrExpression(expr *ast.AttrExpr) error {
	if err := c.compileExpression(expr.Primary); err != nil {
		return err
	}
	c.emitConstant(types.NewString(expr.Attribute))
	c.emit(OpIndex)
	return nil
}

// compileListLiteral handles list literal creation.
func (c *Compiler) compileListLiteral(expr *ast.ListLiteral) error {
	for _, el := range expr.Elements {
		if err := c.compileExpression(el); err != nil {
			return err
		}
	}
	numElements := len(expr.Elements)
	c.emit(OpArray, numElements)
	return nil
}

// compileTableLiteral handles table literal creation.
func (c *Compiler) compileTableLiteral(expr *ast.TableLiteral) error {
	for _, field := range expr.Fields {
		// Handle table keys: if it's an Identifier, treat its name as a string literal.
		// Otherwise, compile it as a regular expression.
		switch key := field.Key.(type) {
		case *ast.Identifier:
			c.emitConstant(types.NewString(key.Name))
		case *ast.StringLiteral: // Explicitly handle string literals as keys
			c.emitConstant(types.NewString(key.Value))
		default:
			// For any other expression type (e.g., `{"key" + "name" = 10}`),
			// compile the expression and expect it to resolve to a string at runtime.
			if err := c.compileExpression(field.Key); err != nil {
				return err
			}
		}

		if err := c.compileExpression(field.Value); err != nil {
			return err
		}
	}
	numFields := len(expr.Fields)
	c.emit(OpTable, numFields)
	return nil
}

// compileTupleLiteral handles tuple literal creation.
func (c *Compiler) compileTupleLiteral(expr *ast.TupleLiteral) error {
	for _, el := range expr.Elements {
		if err := c.compileExpression(el); err != nil {
			return err
		}
	}
	c.emit(OpArray, len(expr.Elements)) // Tuples can be represented as immutable arrays
	return nil
}

// compileFuncDef compiles a function definition.
func (c *Compiler) compileFuncDef(stmt *ast.FunctionDef) error {
	// 1. Save the current instructions slice for the outer scope
	outerInstructions := c.instructions
	c.instructions = make(Instructions, 0) // Initialize a NEW slice for this function's instructions

	// 2. Create the function's scope and set it as current.
	c.enterScope(true)          // This creates the function's scope and sets c.currentScope to it.
	funcScope := c.currentScope // Now funcScope truly points to the function's symbol table.

	for _, param := range stmt.Params {
		funcScope.DefineParameter(param.Name) // Define parameters directly in funcScope
	}

	// 3. Compile the function body using the new (function-specific) instructions slice.
	if err := c.compileStatement(stmt.Body); err != nil {
		// IMPORTANT: If there's an error, you must restore instructions before returning
		c.instructions = outerInstructions
		c.leaveScope()
		return err
	}

	if !c.returned {
		c.emit(OpNull)
		c.emit(OpReturn)
	}

	// 4. Capture the compiled instructions for this function.
	functionInstructions := c.instructions

	// Get numDefinitions from funcScope, which correctly accumulated parameters and direct locals.
	functionNumLocals := funcScope.NumDefinitions()
	functionNumParameters := len(stmt.Params)

	freeSymbols := funcScope.FreeSymbols() // Get free symbols from the function's scope

	// 5. Restore the outer scope and its instructions.
	c.leaveScope()
	c.instructions = outerInstructions // Restore the instructions slice for the outer scope

	compiledFn := &types.CompiledFunction{
		Instructions:  functionInstructions, // This is the function's bytecode
		NumLocals:     functionNumLocals,
		NumParameters: functionNumParameters,
		FreeCount:     len(freeSymbols),
	}

	fnConstIndex := len(c.constants)
	c.constants = append(c.constants, compiledFn)

	// Emit instructions to push captured free variables onto the stack
	for _, sym := range freeSymbols {
		// Resolve in the *outer scope* (which is now c.currentScope)
		outerSym, ok := c.currentScope.Resolve(sym.Name)
		if !ok {
			return fmt.Errorf("internal compiler error: free variable '%s' not found in outer scope during closure compilation", sym.Name)
		}

		switch outerSym.Kind {
		case Global:
			c.emit(OpGetGlobal, outerSym.Index)
		case Local, Parameter:
			c.emit(OpGetLocal, outerSym.Index)
		case Free:
			c.emit(OpGetFree, outerSym.Index) // Use outerSym.Index for nested free variables
		case Builtin:
			c.emit(OpGetGlobal, outerSym.Index)
		default:
			return fmt.Errorf("unsupported free variable kind for closure capture: %s for '%s'", outerSym.Kind, sym.Name)
		}
	}

	c.emit(OpClosure, fnConstIndex, len(freeSymbols))

	// Define the function name in the current (outer) scope
	funcSym, ok := c.currentScope.Resolve(stmt.Name)
	if !ok {
		if c.currentScope == c.globals {
			funcSym = c.globals.DefineGlobal(stmt.Name)
		} else {
			funcSym = c.currentScope.DefineLocal(stmt.Name)
		}
	}

	switch funcSym.Kind {
	case Global:
		c.emit(OpSetGlobal, funcSym.Index)
	case Local, Parameter:
		c.emit(OpSetLocal, funcSym.Index)
	case Free:
		c.emit(OpSetFree, funcSym.Index)
	default:
		return fmt.Errorf("cannot assign function to %s %s", funcSym.Kind, stmt.Name)
	}

	c.returned = false

	return nil
}

// compileBreak compiles a break statement.
func (c *Compiler) compileBreak() error {
	if len(c.loopJumpStack) == 0 {
		return fmt.Errorf("break statement outside of a loop")
	}
	loopJumps := c.loopJumpStack[len(c.loopJumpStack)-1]
	breakJumpPos := loopJumps[0]

	c.emit(OpJump, breakJumpPos)
	return nil
}

// compileContinue compiles a continue statement.
func (c *Compiler) compileContinue() error {
	if len(c.loopJumpStack) == 0 {
		return fmt.Errorf("continue statement outside of a loop")
	}
	loopJumps := c.loopJumpStack[len(c.loopJumpStack)-1]
	continueJumpPos := loopJumps[1]

	c.emit(OpJump, continueJumpPos)
	return nil
}

// compileImport handles import statements.
func (c *Compiler) compileImport(stmt *ast.ImportStmt) error {
	c.emitConstant(types.NewString(stmt.Path))
	c.emit(OpImport, 0)
	c.emit(OpPop)
	return nil
}

// emit appends an instruction.
func (c *Compiler) emit(op Opcode, operands ...int) int {
	ins := Make(op, operands...)
	pos := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return pos
}

// emitConstant adds a constant and emits OpConstant.
func (c *Compiler) emitConstant(val types.Value) {
	idx := len(c.constants)
	c.constants = append(c.constants, val)
	c.emit(OpConstant, idx)
}

// patchJump fixes a jump operand.
func (c *Compiler) patchJump(jumpPos, target int) {
	offset := target - (jumpPos + 3)
	c.instructions[jumpPos+1] = byte(offset >> 8)
	c.instructions[jumpPos+2] = byte(offset)
}

// enterScope creates and enters a new symbol table scope.
func (c *Compiler) enterScope(isFunction bool) {
	outer := c.currentScope
	newScope := NewEnclosedSymbolTable(outer, isFunction)
	c.symbolStack = append(c.symbolStack, newScope)
	c.currentScope = newScope
	c.returned = false
}

// leaveScope exits the current symbol table scope.
func (c *Compiler) leaveScope() {
	if len(c.symbolStack) <= 1 {
		return
	}
	c.symbolStack = c.symbolStack[:len(c.symbolStack)-1]
	c.currentScope = c.symbolStack[len(c.symbolStack)-1]
}

// compileIf compiles an if statement with optional else.
func (c *Compiler) compileIf(stmt *ast.IfStmt) error {
	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 0)
	if err := c.compileStatement(stmt.Then); err != nil {
		return err
	}
	if c.returned {
		c.patchJump(jumpNotTruthyPos, len(c.instructions))
		c.returned = false
		return nil
	}

	jumpPos := c.emit(OpJump, 0)
	afterThenPos := len(c.instructions)
	c.patchJump(jumpNotTruthyPos, afterThenPos)

	if stmt.Else != nil {
		if err := c.compileStatement(stmt.Else); err != nil {
			return err
		}
	}
	afterElsePos := len(c.instructions)
	c.patchJump(jumpPos, afterElsePos)
	return nil
}

// compileWhile compiles a while loop.
func (c *Compiler) compileWhile(stmt *ast.WhileStmt) error {
	loopStart := len(c.instructions)

	breakJumpPos := c.emit(OpJump, 0)
	continueJumpPos := len(c.instructions)

	c.loopJumpStack = append(c.loopJumpStack, []int{breakJumpPos, continueJumpPos})

	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}
	exitPos := c.emit(OpJumpNotTruthy, 0)

	if err := c.compileStatement(stmt.Body); err != nil {
		return err
	}

	c.emit(OpJump, loopStart)

	afterLoop := len(c.instructions)
	c.patchJump(exitPos, afterLoop)
	c.patchJump(breakJumpPos, afterLoop)

	c.loopJumpStack = c.loopJumpStack[:len(c.loopJumpStack)-1]

	return nil
}

// compileFor compiles a for-in loop.
func (c *Compiler) compileFor(stmt *ast.ForStmt) error {
	if err := c.compileExpression(stmt.Iterable); err != nil {
		return err
	}
	c.emit(OpGetIter)

	loopStart := len(c.instructions)

	breakJumpPos := c.emit(OpJump, 0)
	continueJumpPos := len(c.instructions)

	c.loopJumpStack = append(c.loopJumpStack, []int{breakJumpPos, continueJumpPos})

	exitPos := c.emit(OpIterNext, 0)

	sym := c.currentScope.DefineLocal(stmt.Variable)
	c.emit(OpSetLocal, sym.Index)

	if err := c.compileStatement(stmt.Body); err != nil {
		return err
	}

	c.emit(OpJump, loopStart)

	afterLoop := len(c.instructions)
	c.patchJump(exitPos, afterLoop)
	c.patchJump(breakJumpPos, afterLoop)

	c.loopJumpStack = c.loopJumpStack[:len(c.loopJumpStack)-1]

	return nil
}
