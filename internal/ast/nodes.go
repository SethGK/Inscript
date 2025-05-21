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
	stmtNode() // Marker method to indicate it's a statement node
}

// Expression is the interface for all expression nodes.
type Expression interface {
	Node
	exprNode() // Marker method to indicate it's an expression node
}

// --- Custom Token Type ---
// Define a simple Token struct to hold operator information
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
func (p *Program) stmtNode()      {} // Program is also a statement for top-level compilation

// BlockStmt represents a block of statements enclosed in braces `{ ... }`.
type BlockStmt struct {
	Stmts    []Statement
	PosToken token.Pos // Position of the opening brace '{'
}

func (b *BlockStmt) stmtNode()      {}
func (b *BlockStmt) Pos() token.Pos { return b.PosToken }

// ExprStmt represents a statement that is just an expression.
type ExprStmt struct {
	Expr     Expression
	PosToken token.Pos // Position of the start of the expression
}

func (e *ExprStmt) stmtNode()      {}
func (e *ExprStmt) Pos() token.Pos { return e.PosToken }

// AssignStmt represents an assignment statement: `target op value`.
type AssignStmt struct {
	Target   Expression // Target can be Identifier, IndexExpr, AttrExpr
	Op       Token      // Assignment operator (using custom Token struct)
	Value    Expression
	PosToken token.Pos // Position of the target
}

func (a *AssignStmt) stmtNode()      {}
func (a *AssignStmt) Pos() token.Pos { return a.PosToken }

// IfStmt represents an if-else if-else statement.
type IfStmt struct {
	Cond     Expression
	Then     *BlockStmt
	ElseIfs  []ElseIf   // Slice of ElseIf blocks
	Else     *BlockStmt // Optional else block (nil if not present)
	PosToken token.Pos  // Position of the 'if' keyword
}

// ElseIf represents an 'else if' branch.
type ElseIf struct {
	Cond     Expression
	Body     *BlockStmt
	PosToken token.Pos // Position of the 'else' keyword for this block
}

func (i *IfStmt) stmtNode()      {}
func (i *IfStmt) Pos() token.Pos { return i.PosToken }

// WhileStmt represents a while loop: `while cond { body }`.
type WhileStmt struct {
	Cond     Expression
	Body     *BlockStmt
	PosToken token.Pos // Position of the 'while' keyword
}

func (w *WhileStmt) stmtNode()      {}
func (w *WhileStmt) Pos() token.Pos { return w.PosToken }

// ForStmt represents a for-in loop: `for variable in iterable { body }`.
type ForStmt struct {
	Variable string // The loop variable name
	Iterable Expression
	Body     *BlockStmt
	PosToken token.Pos // Position of the 'for' keyword
}

func (f *ForStmt) stmtNode()      {}
func (f *ForStmt) Pos() token.Pos { return f.PosToken }

// FunctionDef represents a function definition statement: `function name(params) -> type? { body }`.
type FunctionDef struct {
	Name       string // Function name
	Params     []Param
	ReturnType *TypeAnnotation // Optional return type annotation (nil if not present)
	Body       *BlockStmt
	PosToken   token.Pos // Position of the 'function' keyword
}

func (f *FunctionDef) stmtNode()      {}
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

func (b *BreakStmt) stmtNode()      {}
func (b *BreakStmt) Pos() token.Pos { return b.PosToken }

// ContinueStmt represents a continue statement.
type ContinueStmt struct {
	PosToken token.Pos // Position of the 'continue' keyword
}

func (c *ContinueStmt) stmtNode()      {}
func (c *ContinueStmt) Pos() token.Pos { return c.PosToken }

// ReturnStmt represents a return statement: `return expression?`.
type ReturnStmt struct {
	Expr     Expression // Optional return value (nil if not present)
	PosToken token.Pos  // Position of the 'return' keyword
}

func (r *ReturnStmt) stmtNode()      {}
func (r *ReturnStmt) Pos() token.Pos { return r.PosToken }

// ImportStmt represents an import statement: `import string`.
type ImportStmt struct {
	Path     string    // The path string (unquoted)
	PosToken token.Pos // Position of the 'import' keyword
}

func (i *ImportStmt) stmtNode()      {}
func (i *ImportStmt) Pos() token.Pos { return i.PosToken }

// PrintStmt represents a print statement: `print(expressions...)`.
type PrintStmt struct {
	Exprs    []Expression // List of expressions to print
	PosToken token.Pos    // Position of the 'print' keyword
}

func (p *PrintStmt) stmtNode()      {}
func (p *PrintStmt) Pos() token.Pos { return p.PosToken }

// --- Expression Nodes ---

// BinaryExpr represents a binary operation: `left operator right`.
type BinaryExpr struct {
	Left     Expression
	Operator Token // The operator token (using custom Token struct)
	Right    Expression
	PosToken token.Pos // Position of the operator
}

func (b *BinaryExpr) exprNode()      {}
func (b *BinaryExpr) Pos() token.Pos { return b.PosToken }

// UnaryExpr represents a unary operation: `operator expression`.
type UnaryExpr struct {
	Operator Token // The operator token (using custom Token struct)
	Expr     Expression
	PosToken token.Pos // Position of the operator
}

func (u *UnaryExpr) exprNode()      {}
func (u *UnaryExpr) Pos() token.Pos { return u.PosToken }

// CallExpr represents a function call: `callee(args...)`.
type CallExpr struct {
	Callee   Expression   // The expression being called (e.g., an Identifier or another CallExpr)
	Args     []Expression // List of arguments
	PosToken token.Pos    // Position of the opening parenthesis '('
}

func (c *CallExpr) exprNode()      {}
func (c *CallExpr) Pos() token.Pos { return c.PosToken }

// IndexExpr represents an index access (e.g., list[index], table[key]).
type IndexExpr struct {
	Primary  Expression // The expression being indexed (list, table, string)
	Index    Expression // The index or key expression
	PosToken token.Pos  // Position of the opening bracket '['
}

func (i *IndexExpr) exprNode()      {}
func (i *IndexExpr) Pos() token.Pos { return i.PosToken }

// AttrExpr represents an attribute access (e.g., obj.attribute).
type AttrExpr struct {
	Primary   Expression // The expression whose attribute is being accessed
	Attribute string     // The attribute name (Identifier text)
	PosToken  token.Pos  // Position of the dot '.'
}

func (a *AttrExpr) exprNode()      {}
func (a *AttrExpr) Pos() token.Pos { return a.PosToken }

// Identifier represents a variable or function name.
type Identifier struct {
	Name     string
	PosToken token.Pos // Position of the identifier token
}

func (i *Identifier) exprNode()      {}
func (i *Identifier) Pos() token.Pos { return i.PosToken }

// --- Literal Nodes ---

// IntegerLiteral represents an integer literal (e.g., 123).
type IntegerLiteral struct {
	Value    int64
	PosToken token.Pos // Position of the number token
}

func (i *IntegerLiteral) exprNode()      {}
func (i *IntegerLiteral) Pos() token.Pos { return i.PosToken }

// FloatLiteral represents a floating-point literal (e.g., 1.23).
type FloatLiteral struct {
	Value    float64
	PosToken token.Pos // Position of the number token
}

func (f *FloatLiteral) exprNode()      {}
func (f *FloatLiteral) Pos() token.Pos { return f.PosToken }

// StringLiteral represents a string literal (e.g., "hello" or 'world').
type StringLiteral struct {
	Value    string    // The unquoted string value
	PosToken token.Pos // Position of the string token
}

func (s *StringLiteral) exprNode()      {}
func (s *StringLiteral) Pos() token.Pos { return s.PosToken }

// BooleanLiteral represents a boolean literal (`true` or `false`).
type BooleanLiteral struct {
	Value    bool
	PosToken token.Pos // Position of the boolean token
}

func (b *BooleanLiteral) exprNode()      {}
func (b *BooleanLiteral) Pos() token.Pos { return b.PosToken }

// NilLiteral represents the `nil` literal.
type NilLiteral struct {
	PosToken token.Pos // Position of the 'nil' token
}

func (n *NilLiteral) exprNode()      {}
func (n *NilLiteral) Pos() token.Pos { return n.PosToken }

// ListLiteral represents a list literal: `[elements...]`.
type ListLiteral struct {
	Elements []Expression
	PosToken token.Pos // Position of the opening bracket '['
}

func (l *ListLiteral) exprNode()      {}
func (l *ListLiteral) Pos() token.Pos { return l.PosToken }

// TableLiteral represents a table literal: `{ fields... }`.
type TableLiteral struct {
	Fields   []*TableField // Corrected to []*TableField
	PosToken token.Pos     // Position of the opening brace '{'
}

func (t *TableLiteral) exprNode()      {}
func (t *TableLiteral) Pos() token.Pos { return t.PosToken }

// TableField represents a key-value pair in a table literal: `key = value`.
type TableField struct {
	Key      Expression // Key can be Identifier, StringLiteral, or other Expression (from tableKey rule)
	Value    Expression
	PosToken token.Pos // Position of the key or the '=' sign
}

func (t *TableField) Pos() token.Pos { return t.PosToken } // TableField is not a Statement or Expression

// TupleLiteral represents a tuple literal (e.g., (1, 2, "a")).
type TupleLiteral struct {
	Elements []Expression
	PosToken token.Pos // Position of the opening parenthesis '('
}

func (t *TupleLiteral) exprNode()      {}
func (t *TupleLiteral) Pos() token.Pos { return t.PosToken }

// FunctionLiteral represents a function definition used as an expression: `function(params) -> type? { body }`.
// Note: This is an expression, distinct from FunctionDef (which is a statement).
type FunctionLiteral struct {
	Params     []Param
	ReturnType *TypeAnnotation
	Body       *BlockStmt
	PosToken   token.Pos
}

func (f *FunctionLiteral) exprNode()      {}
func (f *FunctionLiteral) Pos() token.Pos { return f.PosToken }
