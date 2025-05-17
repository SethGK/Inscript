package ast

import (
	"go/token"
)

// Node is the interface for all AST nodes.
type Node interface {
	Pos() token.Pos // Position of the first token associated with the node
}

// Statement is the interface for all statement nodes.
type Statement interface {
	Node
	statementNode() // Marker method to indicate it's a statement node
}

// Expression is the interface for all expression nodes.
type Expression interface {
	Node
	expressionNode() // Marker method to indicate it's an expression node
}

// --- Custom Token Type ---
// Define a simple Token struct to hold operator information
// This bypasses potential issues with go/token.Token if the compiler
// is having trouble resolving it in this package context.
type Token struct {
	Type    int       // The token type (using the integer ID from ANTLR)
	Pos     token.Pos // The position of the token in the source
	Literal string    // The literal text of the token
}

// --- Statement Nodes ---

// Program represents the root of the AST.
type Program struct {
	Stmts    []Statement
	PosToken token.Pos // Position of the first token (usually the start of the file)
}

func (p *Program) Pos() token.Pos { return p.PosToken }

// BlockStmt represents a block of statements enclosed in braces `{ ... }`.
type BlockStmt struct {
	Stmts    []Statement
	PosToken token.Pos // Position of the opening brace '{'
}

func (b *BlockStmt) statementNode() {}
func (b *BlockStmt) Pos() token.Pos { return b.PosToken }

// ExprStmt represents a statement that is just an expression.
type ExprStmt struct {
	Expr     Expression
	PosToken token.Pos // Position of the start of the expression
}

func (e *ExprStmt) statementNode() {}
func (e *ExprStmt) Pos() token.Pos { return e.PosToken }

// AssignStmt represents an assignment statement: `target op value`.
type AssignStmt struct {
	Target   Expression // Target can be Identifier, IndexExpr, AttrExpr
	Op       Token      // Assignment operator (using custom Token struct)
	Value    Expression
	PosToken token.Pos // Position of the target
}

func (a *AssignStmt) statementNode() {}
func (a *AssignStmt) Pos() token.Pos { return a.PosToken }

// IfStmt represents an if-else if-else statement.
type IfStmt struct {
	Cond     Expression
	Then     *BlockStmt
	ElseIfs  []ElseIf   // Slice of ElseIf blocks
	Else     *BlockStmt // Optional else block (nil if not present)
	PosToken token.Pos  // Position of the 'if' keyword
}

// ElseIf represents an else if block within an IfStmt.
type ElseIf struct {
	Cond     Expression
	Body     *BlockStmt
	PosToken token.Pos // Position of the 'else' keyword for this block
}

func (i *IfStmt) statementNode() {}
func (i *IfStmt) Pos() token.Pos { return i.PosToken }

// WhileStmt represents a while loop: `while cond { body }`.
type WhileStmt struct {
	Cond     Expression
	Body     *BlockStmt
	PosToken token.Pos // Position of the 'while' keyword
}

func (w *WhileStmt) statementNode() {}
func (w *WhileStmt) Pos() token.Pos { return w.PosToken }

// ForStmt represents a for-in loop: `for variable in iterable { body }`.
type ForStmt struct {
	Variable string // The loop variable name
	Iterable Expression
	Body     *BlockStmt
	PosToken token.Pos // Position of the 'for' keyword
}

func (f *ForStmt) statementNode() {}
func (f *ForStmt) Pos() token.Pos { return f.PosToken }

// FunctionDef represents a function definition statement: `function name(params) -> type? { body }`.
type FunctionDef struct {
	Name       string // Function name
	Params     []Param
	ReturnType *TypeAnnotation // Optional return type annotation (nil if not present)
	Body       *BlockStmt
	PosToken   token.Pos // Position of the 'function' keyword
}

func (f *FunctionDef) statementNode() {}
func (f *FunctionDef) Pos() token.Pos { return f.PosToken }

// Param represents a function parameter: `name = defaultValue? : type?` or `... name`.
type Param struct {
	Name         string
	DefaultValue Expression      // Optional default value (nil if not present)
	Type         *TypeAnnotation // Optional type annotation (nil if not present)
	IsVariadic   bool            // True if this is a variadic parameter (...)
	PosToken     token.Pos       // Position of the parameter name or '...'
}

// TypeAnnotation represents a type annotation: `identifier`.
type TypeAnnotation struct {
	Name     string    // The type name (currently just an identifier)
	PosToken token.Pos // Position of the identifier
}

// BreakStmt represents a break statement.
type BreakStmt struct {
	PosToken token.Pos // Position of the 'break' keyword
}

func (b *BreakStmt) statementNode() {}
func (b *BreakStmt) Pos() token.Pos { return b.PosToken }

// ContinueStmt represents a continue statement.
type ContinueStmt struct {
	PosToken token.Pos // Position of the 'continue' keyword
}

func (c *ContinueStmt) statementNode() {}
func (c *ContinueStmt) Pos() token.Pos { return c.PosToken }

// ReturnStmt represents a return statement: `return expression?`.
type ReturnStmt struct {
	Expr     Expression // Optional return value (nil if not present)
	PosToken token.Pos  // Position of the 'return' keyword
}

func (r *ReturnStmt) statementNode() {}
func (r *ReturnStmt) Pos() token.Pos { return r.PosToken }

// ImportStmt represents an import statement: `import string`.
type ImportStmt struct {
	Path     string    // The path string (unquoted)
	PosToken token.Pos // Position of the 'import' keyword
}

func (i *ImportStmt) statementNode() {}
func (i *ImportStmt) Pos() token.Pos { return i.PosToken }

// PrintStmt represents a print statement: `print(expressions...)`.
type PrintStmt struct {
	Exprs    []Expression // List of expressions to print
	PosToken token.Pos    // Position of the 'print' keyword
}

func (p *PrintStmt) statementNode() {}
func (p *PrintStmt) Pos() token.Pos { return p.PosToken }

// --- Expression Nodes ---

// BinaryExpr represents a binary operation: `left operator right`.
type BinaryExpr struct {
	Left     Expression
	Operator Token // The operator token (using custom Token struct)
	Right    Expression
	PosToken token.Pos // Position of the operator
}

func (b *BinaryExpr) expressionNode() {}
func (b *BinaryExpr) Pos() token.Pos  { return b.PosToken }

// UnaryExpr represents a unary operation: `operator expression`.
type UnaryExpr struct {
	Operator Token // The operator token (using custom Token struct)
	Expr     Expression
	PosToken token.Pos // Position of the operator
}

func (u *UnaryExpr) expressionNode() {}
func (u *UnaryExpr) Pos() token.Pos  { return u.PosToken }

// CallExpr represents a function call: `callee(args...)`.
type CallExpr struct {
	Callee   Expression   // The expression being called (e.g., an Identifier or another CallExpr)
	Args     []Expression // List of arguments
	PosToken token.Pos    // Position of the opening parenthesis '('
}

func (c *CallExpr) expressionNode() {}
func (c *CallExpr) Pos() token.Pos  { return c.PosToken }

// IndexExpr represents an index access: `primary[index]`.
type IndexExpr struct {
	Primary  Expression // The expression being indexed (e.g., an Identifier or CallExpr)
	Index    Expression // The index expression
	PosToken token.Pos  // Position of the opening bracket '['
}

func (i *IndexExpr) expressionNode() {}
func (i *IndexExpr) Pos() token.Pos  { return i.PosToken }

// AttrExpr represents an attribute access: `primary.attribute`.
type AttrExpr struct {
	Primary   Expression // The expression whose attribute is being accessed
	Attribute string     // The attribute name (Identifier text)
	PosToken  token.Pos  // Position of the dot '.'
}

func (a *AttrExpr) expressionNode() {}
func (a *AttrExpr) Pos() token.Pos  { return a.PosToken }

// Identifier represents a variable or function name.
type Identifier struct {
	Name     string
	PosToken token.Pos // Position of the identifier token
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Pos() token.Pos  { return i.PosToken }

// --- Literal Nodes ---

// IntegerLiteral represents an integer literal (e.g., 123).
type IntegerLiteral struct {
	Value    int64
	PosToken token.Pos // Position of the number token
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) Pos() token.Pos  { return i.PosToken }

// FloatLiteral represents a floating-point literal (e.g., 1.23).
type FloatLiteral struct {
	Value    float64
	PosToken token.Pos // Position of the number token
}

func (f *FloatLiteral) expressionNode() {}
func (f *FloatLiteral) Pos() token.Pos  { return f.PosToken }

// StringLiteral represents a string literal (e.g., "hello" or 'world').
type StringLiteral struct {
	Value    string    // The unquoted string value
	PosToken token.Pos // Position of the string token
}

func (s *StringLiteral) expressionNode() {}
func (s *StringLiteral) Pos() token.Pos  { return s.PosToken }

// BooleanLiteral represents a boolean literal (`true` or `false`).
type BooleanLiteral struct {
	Value    bool
	PosToken token.Pos // Position of the boolean token
}

func (b *BooleanLiteral) expressionNode() {}
func (b *BooleanLiteral) Pos() token.Pos  { return b.PosToken }

// NilLiteral represents the `nil` literal.
type NilLiteral struct {
	PosToken token.Pos // Position of the 'nil' token
}

func (n *NilLiteral) expressionNode() {}
func (n *NilLiteral) Pos() token.Pos  { return n.PosToken }

// ListLiteral represents a list literal: `[elements...]`.
type ListLiteral struct {
	Elements []Expression
	PosToken token.Pos // Position of the opening bracket '['
}

func (l *ListLiteral) expressionNode() {}
func (l *ListLiteral) Pos() token.Pos  { return l.PosToken }

// TableLiteral represents a table literal: `{ fields... }`.
type TableLiteral struct {
	Fields   []TableField
	PosToken token.Pos // Position of the opening brace '{'
}

func (t *TableLiteral) expressionNode() {}
func (t *TableLiteral) Pos() token.Pos  { return t.PosToken }

// TableField represents a key-value pair in a table literal: `key = value`.
type TableField struct {
	Key      Expression // Key can be Identifier, StringLiteral, or other Expression (from tableKey rule)
	Value    Expression
	PosToken token.Pos // Position of the key or the '=' sign
}

// FunctionLiteral represents a function definition used as an expression: `function(params) -> type? { body }`.
type FunctionLiteral struct {
	Params     []Param
	ReturnType *TypeAnnotation // Optional return type annotation (nil if not present)
	Body       *BlockStmt
	PosToken   token.Pos // Position of the 'function' keyword
}

func (f *FunctionLiteral) expressionNode() {}
func (f *FunctionLiteral) Pos() token.Pos  { return f.PosToken }
