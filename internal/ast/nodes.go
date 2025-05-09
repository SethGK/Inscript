// internal/ast/nodes.go
package ast

import (
	"go/token"
)

// Node is the interface for all AST nodes.
type Node interface {
	Pos() token.Pos
}

// Statement is the interface for statement nodes.
type Statement interface {
	Node
	isStatement()
}

// Expression is the interface for expression nodes.
type Expression interface {
	Node
	isExpression()
}

// Program is the root AST node.
type Program struct {
	Stmts    []Statement // top-level statements
	PosToken token.Pos   // position of first token
}

func (p *Program) Pos() token.Pos { return p.PosToken }

// BlockStmt represents a `{ ... }` block of statements.
type BlockStmt struct {
	Stmts    []Statement
	PosToken token.Pos
}

func (b *BlockStmt) isStatement()   {}
func (b *BlockStmt) Pos() token.Pos { return b.PosToken }

// ExprStmt wraps an expression as a statement.
type ExprStmt struct {
	Expr     Expression
	PosToken token.Pos
}

func (e *ExprStmt) isStatement()   {}
func (e *ExprStmt) Pos() token.Pos { return e.PosToken }

// AssignStmt represents `target = value`.
type AssignStmt struct {
	Target   Expression
	Value    Expression
	PosToken token.Pos
}

func (a *AssignStmt) isStatement()   {}
func (a *AssignStmt) Pos() token.Pos { return a.PosToken }

// PrintStmt represents `print(exprList)`.
type PrintStmt struct {
	Exprs    []Expression
	PosToken token.Pos
}

func (p *PrintStmt) isStatement()   {}
func (p *PrintStmt) Pos() token.Pos { return p.PosToken }

// ReturnStmt represents `return expr?`.
type ReturnStmt struct {
	Expr     Expression // may be nil
	PosToken token.Pos
}

func (r *ReturnStmt) isStatement()   {}
func (r *ReturnStmt) Pos() token.Pos { return r.PosToken }

// IfStmt represents an if-elseif-else chain.
type IfStmt struct {
	Cond     Expression
	Then     *BlockStmt
	ElseIfs  []ElseIf
	Else     *BlockStmt // may be nil
	PosToken token.Pos
}
type ElseIf struct {
	Cond     Expression
	Body     *BlockStmt
	PosToken token.Pos
}

func (i *IfStmt) isStatement()   {}
func (i *IfStmt) Pos() token.Pos { return i.PosToken }

// WhileStmt represents `while cond { ... }`.
type WhileStmt struct {
	Cond     Expression
	Body     *BlockStmt
	PosToken token.Pos
}

func (w *WhileStmt) isStatement()   {}
func (w *WhileStmt) Pos() token.Pos { return w.PosToken }

// ForStmt represents `for ident in expr { ... }`.
type ForStmt struct {
	Variable string
	Iterable Expression
	Body     *BlockStmt
	PosToken token.Pos
}

func (f *ForStmt) isStatement()   {}
func (f *ForStmt) Pos() token.Pos { return f.PosToken }

// FuncDefStmt represents `function(params) { body }`.
type FuncDefStmt struct {
	Name     string // optional: empty for anonymous
	Params   []string
	Body     *BlockStmt
	PosToken token.Pos
}

func (f *FuncDefStmt) isStatement()   {}
func (f *FuncDefStmt) Pos() token.Pos { return f.PosToken }

// -- Expression nodes --

// BinaryExpr represents `left op right`.
type BinaryExpr struct {
	Left     Expression
	Operator string
	Right    Expression
	PosToken token.Pos
}

func (b *BinaryExpr) isExpression()  {}
func (b *BinaryExpr) Pos() token.Pos { return b.PosToken }

// UnaryExpr represents `op expr`.
type UnaryExpr struct {
	Operator string
	Expr     Expression
	PosToken token.Pos
}

func (u *UnaryExpr) isExpression()  {}
func (u *UnaryExpr) Pos() token.Pos { return u.PosToken }

// Literal nodes

type IntegerLiteral struct {
	Value    int64
	PosToken token.Pos
}

func (i *IntegerLiteral) isExpression()  {}
func (i *IntegerLiteral) Pos() token.Pos { return i.PosToken }

type FloatLiteral struct {
	Value    float64
	PosToken token.Pos
}

func (f *FloatLiteral) isExpression()  {}
func (f *FloatLiteral) Pos() token.Pos { return f.PosToken }

type StringLiteral struct {
	Value    string
	PosToken token.Pos
}

func (s *StringLiteral) isExpression()  {}
func (s *StringLiteral) Pos() token.Pos { return s.PosToken }

type BooleanLiteral struct {
	Value    bool
	PosToken token.Pos
}

func (b *BooleanLiteral) isExpression()  {}
func (b *BooleanLiteral) Pos() token.Pos { return b.PosToken }

// NilLiteral represents `nil`.
type NilLiteral struct{ PosToken token.Pos }

func (n *NilLiteral) isExpression()  {}
func (n *NilLiteral) Pos() token.Pos { return n.PosToken }

// Identifier represents a variable name.
type Identifier struct {
	Name     string
	PosToken token.Pos
}

func (i *Identifier) isExpression()  {}
func (i *Identifier) Pos() token.Pos { return i.PosToken }

// CallExpr represents a function call: `fn(args...)`.
type CallExpr struct {
	Callee   Expression
	Args     []Expression
	PosToken token.Pos
}

func (c *CallExpr) isExpression()  {}
func (c *CallExpr) Pos() token.Pos { return c.PosToken }

// IndexExpr represents `primary[index]`.
type IndexExpr struct {
	Primary  Expression
	Index    Expression
	PosToken token.Pos
}

func (i *IndexExpr) isExpression()  {}
func (i *IndexExpr) Pos() token.Pos { return i.PosToken }

// ListLiteral represents `[expr,...]`.
type ListLiteral struct {
	Elements []Expression
	PosToken token.Pos
}

func (l *ListLiteral) isExpression()  {}
func (l *ListLiteral) Pos() token.Pos { return l.PosToken }

// TableLiteral represents `{ key=expr,... }`.
type TableLiteral struct {
	Fields   []Field
	PosToken token.Pos
}

func (t *TableLiteral) isExpression()  {}
func (t *TableLiteral) Pos() token.Pos { return t.PosToken }

// Field is a key/value pair in a table literal.
type Field struct {
	Key      string
	Value    Expression
	PosToken token.Pos
}

// Note: Field does not implement Node directly; accessed via its container.
