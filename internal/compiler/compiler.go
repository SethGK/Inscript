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
}

// New creates a new top-level Compiler.
func New() *Compiler {
	global := NewSymbolTable()
	c := &Compiler{
		instructions: make(Instructions, 0),
		constants:    make([]types.Value, 0),
		globals:      global,
		symbolStack:  []*SymbolTable{global},
		currentScope: global,
		returned:     false,
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
	// implicit return
	c.emit(OpNull)
	c.emit(OpReturn)

	bc := &Bytecode{
		Instructions:  c.instructions,
		Constants:     c.constants,
		NumLocals:     0,
		NumParameters: 0,
		NumGlobals:    c.globals.NumDefinitions(),
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
		ident := stmt.Target.(*ast.Identifier)
		if err := c.compileExpression(stmt.Value); err != nil {
			return err
		}
		sym, ok := c.currentScope.Resolve(ident.Name)
		if ok {
			switch sym.Kind {
			case Global:
				c.emit(OpSetGlobal, sym.Index)
			case Local, Parameter:
				c.emit(OpSetLocal, sym.Index)
			default:
				return fmt.Errorf("cannot assign to %s %s", sym.Kind, ident.Name)
			}
		} else {
			if c.currentScope == c.globals {
				sym = c.globals.DefineGlobal(ident.Name)
				c.emit(OpSetGlobal, sym.Index)
			} else {
				sym = c.currentScope.DefineLocal(ident.Name)
				c.emit(OpSetLocal, sym.Index)
			}
		}

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
		for _, s2 := range stmt.Stmts {
			if err := c.compileStatement(s2); err != nil {
				return err
			}
		}

	case *ast.IfStmt:
		return c.compileIf(stmt)
	case *ast.WhileStmt:
		return c.compileWhile(stmt)
	case *ast.ForStmt:
		return c.compileFor(stmt)
	default:
		return fmt.Errorf("unsupported statement: %T", s)
	}
	return nil
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
	default:
		return fmt.Errorf("unsupported expression: %T", e)
	}
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

// compileIf compiles an if statement with optional else.
func (c *Compiler) compileIf(stmt *ast.IfStmt) error {
	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}
	jumpNotTruthyPos := c.emit(OpJumpNotTruthy, 0)
	if err := c.compileStatement(stmt.Then); err != nil {
		return err
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
	if err := c.compileExpression(stmt.Cond); err != nil {
		return err
	}
	exitPos := c.emit(OpJumpNotTruthy, 0)
	if err := c.compileStatement(stmt.Body); err != nil {
		return err
	}
	c.emit(OpJump, loopStart)
	afterPos := len(c.instructions)
	c.patchJump(exitPos, afterPos)
	return nil
}

// compileFor compiles a for-in loop.
func (c *Compiler) compileFor(stmt *ast.ForStmt) error {
	if err := c.compileExpression(stmt.Iterable); err != nil {
		return err
	}
	c.emit(OpGetIter)
	loopStart := len(c.instructions)
	exitPos := c.emit(OpIterNext, 0)
	sym := c.currentScope.DefineLocal(stmt.Variable)
	c.emit(OpSetLocal, sym.Index)
	if err := c.compileStatement(stmt.Body); err != nil {
		return err
	}
	c.emit(OpJump, loopStart)
	afterLoop := len(c.instructions)
	c.patchJump(exitPos, afterLoop)
	return nil
}

// patchJump fixes a jump operand.
func (c *Compiler) patchJump(jumpPos, target int) {
	offset := target - (jumpPos + 3)
	c.instructions[jumpPos+1] = byte(offset >> 8)
	c.instructions[jumpPos+2] = byte(offset)
}
