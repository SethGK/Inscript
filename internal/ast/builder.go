package ast

import (
	"fmt"
	"go/token" // We still need go/token for token.Pos
	"strconv"

	parser "github.com/SethGK/Inscript/parser/grammar" // Assuming this is the correct import path
	"github.com/antlr4-go/antlr/v4"
)

// ASTBuilder implements the ANTLR InscriptVisitor to build our AST.
type ASTBuilder struct {
	*parser.BaseInscriptVisitor
}

// NewASTBuilder creates a new ASTBuilder.
func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{BaseInscriptVisitor: &parser.BaseInscriptVisitor{}}
}

// VisitProgram builds the root Program node.
func (v *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	prog := &Program{PosToken: token.Pos(ctx.GetStart().GetStart())}
	for _, stmtCtx := range ctx.AllStatement() {
		// The grammar allows NEWLINE* between statements, so we need to check if the
		// visited context is actually a Statement node.
		if stmtNode, ok := stmtCtx.Accept(v).(Statement); ok && stmtNode != nil {
			prog.Stmts = append(prog.Stmts, stmtNode)
		}
	}
	return prog
}

// VisitStatement visits a statement rule and dispatches to the appropriate handler.
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	// The statement rule in the grammar is a simple choice between all statement types.
	// We can directly visit the child context that is not nil.
	switch {
	case ctx.ExprStmt() != nil:
		return ctx.ExprStmt().Accept(v)
	case ctx.Assignment() != nil:
		return ctx.Assignment().Accept(v)
	case ctx.IfStmt() != nil:
		return ctx.IfStmt().Accept(v)
	case ctx.WhileStmt() != nil:
		return ctx.WhileStmt().Accept(v)
	case ctx.ForStmt() != nil:
		return ctx.ForStmt().Accept(v)
	case ctx.FuncDef() != nil:
		// Note: FuncDef is a statement in the grammar, not an expression literal here.
		return ctx.FuncDef().Accept(v)
	case ctx.BreakStmt() != nil:
		return ctx.BreakStmt().Accept(v)
	case ctx.ContinueStmt() != nil:
		return ctx.ContinueStmt().Accept(v)
	case ctx.ReturnStmt() != nil:
		return ctx.ReturnStmt().Accept(v)
	case ctx.ImportStmt() != nil:
		return ctx.ImportStmt().Accept(v)
	case ctx.PrintStmt() != nil:
		return ctx.PrintStmt().Accept(v)
	case ctx.Block() != nil:
		return ctx.Block().Accept(v)
	default:
		// This case should ideally not be reached if the grammar is correct and parsing succeeded.
		fmt.Printf("WARNING: VisitStatement encountered unexpected child context: %T at line %d\n", ctx.GetChild(0), ctx.GetStart().GetLine())
		return nil // Or panic, depending on desired error handling
	}
}

// VisitBlock builds a BlockStmt node.
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	block := &BlockStmt{PosToken: token.Pos(ctx.GetStart().GetStart())}
	for i, stmtCtx := range ctx.AllStatement() {
		if stmtNode, ok := stmtCtx.Accept(v).(Statement); ok && stmtNode != nil {
			block.Stmts = append(block.Stmts, stmtNode)
			// Optional: Add unreachable code check after return
			if _, isRet := stmtNode.(*ReturnStmt); isRet {
				remaining := ctx.AllStatement()[i+1:]
				if len(remaining) > 0 {
					// A more robust approach might involve walking the remaining contexts
					// to find the start position of the first unreachable statement.
					fmt.Printf("WARNING: Unreachable code after return in block starting at line %d\n", stmtCtx.GetStart().GetLine())
				}
				break // Stop processing statements after a return
			}
		}
		// Handle potential NEWLINE* between statements within the block
		// The grammar allows NEWLINE* after each statement, so we only process Statement contexts.
	}
	return block
}

// VisitExprStmt builds an ExprStmt node.
func (v *ASTBuilder) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	expr := ctx.Expression().Accept(v).(Expression)
	return &ExprStmt{PosToken: token.Pos(ctx.GetStart().GetStart()), Expr: expr}
}

// VisitAssignment builds an AssignStmt node.
func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	// The target rule can be an identifier, index access, or attribute access.
	// We need to visit the target context.
	target := ctx.Target().Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.GetChild(1).(antlr.TerminalNode).GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(), // Use integer token type directly
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	value := ctx.Expression().Accept(v).(Expression)

	// Ensure the target is a valid assignable expression type
	switch target.(type) {
	case *Identifier, *IndexExpr, *AttrExpr:
		// Valid target types
	default:
		// Handle invalid assignment target error
		// You might want to return an error node or panic
		fmt.Printf("ERROR: Invalid assignment target type %T at line %d\n", target, ctx.GetStart().GetLine())
		// Return a placeholder or nil, depending on error handling strategy
		return nil
	}

	return &AssignStmt{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Target:   target,
		Op:       opToken, // Assign the created custom ast.Token struct
		Value:    value,
	}
}

// VisitTarget builds the Expression node for the assignment target.
func (v *ASTBuilder) VisitTarget(ctx *parser.TargetContext) interface{} {
	// The target rule has three alternatives: Identifier, IndexExpr, AttrExpr
	if ctx.IDENTIFIER() != nil && ctx.PostfixExpr() == nil {
		// Case: IDENTIFIER
		idToken := ctx.IDENTIFIER().GetSymbol()
		return &Identifier{PosToken: token.Pos(idToken.GetStart()), Name: idToken.GetText()}
	} else if ctx.PostfixExpr() != nil {
		// Case: postfixExpr LBRACK expression RBRACK or postfixExpr DOT IDENTIFIER
		primary := ctx.PostfixExpr().Accept(v).(Expression)
		if ctx.LBRACK() != nil {
			// Case: postfixExpr LBRACK expression RBRACK (IndexExpr)
			index := ctx.Expression().Accept(v).(Expression)
			lbrackToken := ctx.LBRACK().GetSymbol()
			return &IndexExpr{PosToken: token.Pos(lbrackToken.GetStart()), Primary: primary, Index: index}
		} else if ctx.DOT() != nil {
			// Case: postfixExpr DOT IDENTIFIER (AttrExpr)
			attrToken := ctx.IDENTIFIER().GetSymbol()
			dotToken := ctx.DOT().GetSymbol()
			return &AttrExpr{PosToken: token.Pos(dotToken.GetStart()), Primary: primary, Attribute: attrToken.GetText()}
		}
	}
	// This case should not be reached if the grammar is correct.
	fmt.Printf("WARNING: VisitTarget encountered unexpected context: %T at line %d\n", ctx, ctx.GetStart().GetLine())
	return nil
}

// VisitIfStmt builds an IfStmt node.
func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := ctx.Expression().Accept(v).(Expression)
	thenBlock := ctx.Block(0).Accept(v).(*BlockStmt)

	var elseIfs []ElseIf
	// The grammar is `ifStmt: IF expression block (ELSE block)?`
	// This means there's only one optional ELSE block. If you intended elif,
	// the grammar needs to be updated. Assuming the provided grammar is correct:
	var elseBlock *BlockStmt
	if ctx.ELSE() != nil {
		// The ELSE is followed by a block
		// Need to check if there's a second block context available
		if len(ctx.AllBlock()) > 1 {
			elseBlock = ctx.Block(1).Accept(v).(*BlockStmt) // Get the second block
		} else {
			fmt.Printf("ERROR: ELSE keyword without a following block at line %d\n", ctx.ELSE().GetSymbol().GetLine())
			// Handle error: maybe return an error node or nil
			elseBlock = nil
		}
	}

	// If your intention was `elif`, the grammar should look something like:
	// ifStmt: IF expression block (ELSE IF expression block)* (ELSE block)?;
	// Since it's not, we'll build the AST based on the provided grammar.
	// If you update the grammar, this visitor needs to be updated too.

	return &IfStmt{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Cond:     cond,
		Then:     thenBlock,
		ElseIfs:  elseIfs, // This will be empty based on the provided grammar
		Else:     elseBlock,
	}
}

// VisitWhileStmt builds a WhileStmt node.
func (v *ASTBuilder) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &WhileStmt{PosToken: token.Pos(ctx.GetStart().GetStart()), Cond: cond, Body: body}
}

// VisitForStmt builds a ForStmt node.
func (v *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	varNameToken := ctx.IDENTIFIER().GetSymbol()
	iter := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &ForStmt{PosToken: token.Pos(ctx.GetStart().GetStart()), Variable: varNameToken.GetText(), Iterable: iter, Body: body}
}

// VisitFuncDef builds a FunctionDef statement node.
func (v *ASTBuilder) VisitFuncDef(ctx *parser.FuncDefContext) interface{} {
	nameToken := ctx.IDENTIFIER().GetSymbol()
	var params []Param
	if ctx.ParamList() != nil {
		// VisitParamList returns []Param
		params = ctx.ParamList().Accept(v).([]Param)
	}

	var returnType *TypeAnnotation
	if ctx.TypeAnnotation() != nil {
		// VisitTypeAnnotation returns *TypeAnnotation
		returnType = ctx.TypeAnnotation().Accept(v).(*TypeAnnotation)
	}

	body := ctx.Block().Accept(v).(*BlockStmt)

	return &FunctionDef{
		PosToken:   token.Pos(ctx.GetStart().GetStart()),
		Name:       nameToken.GetText(),
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

// VisitParamList handles a comma-separated list of parameters.
func (v *ASTBuilder) VisitParamList(ctx *parser.ParamListContext) interface{} {
	var params []Param
	for _, paramCtx := range ctx.AllParam() {
		// VisitParam returns a single Param struct
		param := paramCtx.Accept(v).(Param)
		params = append(params, param)
	}
	return params // Return a slice of Param structs
}

// VisitParam handles a single function parameter.
func (v *ASTBuilder) VisitParam(ctx *parser.ParamContext) interface{} {
	posToken := token.Pos(ctx.GetStart().GetStart())
	isVariadic := ctx.ELLIPSIS() != nil

	var name string
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
		posToken = token.Pos(ctx.IDENTIFIER().GetSymbol().GetStart()) // Use identifier position
	} else if isVariadic && ctx.ELLIPSIS() != nil && ctx.GetChildCount() > 1 {
		// Handle ... identifier case
		if idToken, ok := ctx.GetChild(1).(antlr.TerminalNode); ok && idToken.GetSymbol().GetTokenType() == parser.InscriptParserIDENTIFIER {
			name = idToken.GetText()
			posToken = token.Pos(ctx.ELLIPSIS().GetSymbol().GetStart()) // Use ... token position
		} else {
			// Error: Expected identifier after ...
			fmt.Printf("ERROR: Expected identifier after '...' in parameter list at line %d\n", ctx.GetStart().GetLine())
			name = "" // Or handle error appropriately
		}
	} else {
		// Error: Parameter must be an identifier or variadic identifier
		fmt.Printf("ERROR: Invalid parameter definition at line %d\n", ctx.GetStart().GetLine())
		name = "" // Or handle error appropriately
	}

	var defaultValue Expression
	// Check for ASSIGN and Expression for default value
	if ctx.ASSIGN() != nil && ctx.Expression() != nil {
		defaultValue = ctx.Expression().Accept(v).(Expression)
	}

	var paramType *TypeAnnotation
	// Check for COLON and TypeAnnotation
	if ctx.COLON() != nil && ctx.TypeAnnotation() != nil {
		paramType = ctx.TypeAnnotation().Accept(v).(*TypeAnnotation)
	}

	return Param{
		PosToken:     posToken,
		Name:         name,
		DefaultValue: defaultValue,
		Type:         paramType,
		IsVariadic:   isVariadic,
	}
}

// VisitTypeAnnotation builds a TypeAnnotation node.
func (v *ASTBuilder) VisitTypeAnnotation(ctx *parser.TypeAnnotationContext) interface{} {
	nameToken := ctx.IDENTIFIER().GetSymbol()
	return &TypeAnnotation{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Name:     nameToken.GetText(),
	}
}

// VisitBreakStmt builds a BreakStmt node.
func (v *ASTBuilder) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	return &BreakStmt{PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitContinueStmt builds a ContinueStmt node.
func (v *ASTBuilder) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	return &ContinueStmt{PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitReturnStmt builds a ReturnStmt node.
func (v *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	var expr Expression
	if ctx.Expression() != nil { // Check if the optional expression exists
		expr = ctx.Expression().Accept(v).(Expression)
	}
	return &ReturnStmt{PosToken: token.Pos(ctx.GetStart().GetStart()), Expr: expr}
}

// VisitImportStmt builds an ImportStmt node.
func (v *ASTBuilder) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	// The grammar defines IMPORT followed by STRING
	stringToken := ctx.STRING().GetSymbol()
	// Remove quotes from the string literal value
	path := stringToken.GetText()
	// Handle different quote types if necessary (single, double, triple)
	// For simplicity, assuming standard double quotes for now.
	if len(path) > 1 {
		path = path[1 : len(path)-1] // Remove surrounding quotes
	} else {
		path = "" // Handle empty string literal case
	}

	// TODO: Add proper escape sequence unescaping

	return &ImportStmt{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Path:     path,
	}
}

// VisitPrintStmt builds a PrintStmt node.
func (v *ASTBuilder) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	var exprs []Expression
	// The grammar is PRINT LPAREN (expression (COMMA expression)*)? RPAREN
	// The expressions are directly children of the PrintStmtContext if they exist.
	// We need to iterate through the children and find Expression contexts.
	// A more robust way is to define an `argList` rule for print like for calls.
	// Assuming the current grammar structure:
	// Check if there's a non-nil child after the LPAREN and before the RPAREN
	if ctx.GetChildCount() > 2 {
		// Assuming the structure is PRINT LPAREN expr (COMMA expr)* RPAREN
		// Iterate from the child after LPAREN up to the child before RPAREN
		for i := 2; i < ctx.GetChildCount()-1; i++ {
			child := ctx.GetChild(i)
			// Only process expression contexts, skip commas
			if exprCtx, ok := child.(parser.IExpressionContext); ok {
				if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
					exprs = append(exprs, exprNode)
				}
			}
		}
	}
	// A better grammar would be: printStmt: PRINT LPAREN argList? RPAREN;
	// If you update the grammar, update this visitor accordingly.

	return &PrintStmt{PosToken: token.Pos(ctx.GetStart().GetStart()), Exprs: exprs}
}

// --- Expression Visitor Methods (Using Labeled Alternatives) ---

// VisitUnaryExpression handles the base case for expressions, visiting a unaryExpr.
func (v *ASTBuilder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	return ctx.UnaryExpr().Accept(v)
}

// VisitExpExpr handles the power (^^) binary operation.
func (v *ASTBuilder) VisitExpExpr(ctx *parser.ExpExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.POW().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitMulExpr handles the multiplication (*) binary operation.
func (v *ASTBuilder) VisitMulExpr(ctx *parser.MulExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.MUL().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitDivExpr handles the division (/) binary operation.
func (v *ASTBuilder) VisitDivExpr(ctx *parser.DivExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.DIV().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitIdivExpr handles the integer division (//) binary operation.
func (v *ASTBuilder) VisitIdivExpr(ctx *parser.IdivExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.IDIV().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitModExpr handles the modulo (%) binary operation.
func (v *ASTBuilder) VisitModExpr(ctx *parser.ModExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.MOD().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitAddExpr handles the addition (+) binary operation.
func (v *ASTBuilder) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.ADD().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitSubExpr handles the subtraction (-) binary operation.
func (v *ASTBuilder) VisitSubExpr(ctx *parser.SubExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.SUB().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitBitandExpr handles the bitwise AND (&) binary operation.
func (v *ASTBuilder) VisitBitandExpr(ctx *parser.BitandExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.BITAND().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitBitorExpr handles the bitwise OR (|) binary operation.
func (v *ASTBuilder) VisitBitorExpr(ctx *parser.BitorExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.BITOR().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitBitxorExpr handles the bitwise XOR (^) binary operation.
func (v *ASTBuilder) VisitBitxorExpr(ctx *parser.BitxorExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.BITXOR().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitShlExpr handles the left shift (<<) binary operation.
func (v *ASTBuilder) VisitShlExpr(ctx *parser.ShlExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.SHL().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitShrExpr handles the right shift (>>) binary operation.
func (v *ASTBuilder) VisitShrExpr(ctx *parser.ShrExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.SHR().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitLtExpr handles the less than (<) binary operation.
func (v *ASTBuilder) VisitLtExpr(ctx *parser.LtExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.LT().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitLeExpr handles the less than or equal to (<=) binary operation.
func (v *ASTBuilder) VisitLeExpr(ctx *parser.LeExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.LE().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitGtExpr handles the greater than (>) binary operation.
func (v *ASTBuilder) VisitGtExpr(ctx *parser.GtExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.GT().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitGeExpr handles the greater than or equal to (>=) binary operation.
func (v *ASTBuilder) VisitGeExpr(ctx *parser.GeExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.GE().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitEqExpr handles the equality (==) binary operation.
func (v *ASTBuilder) VisitEqExpr(ctx *parser.EqExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.EQ().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitNeqExpr handles the inequality (!=) binary operation.
func (v *ASTBuilder) VisitNeqExpr(ctx *parser.NeqExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.NEQ().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitAndExpr handles the logical AND (and) binary operation.
func (v *ASTBuilder) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.AND().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// VisitOrExpr handles the logical OR (or) binary operation.
func (v *ASTBuilder) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	left := ctx.Expression(0).Accept(v).(Expression)
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.OR().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	right := ctx.Expression(1).Accept(v).(Expression)
	return &BinaryExpr{PosToken: opToken.Pos, Left: left, Operator: opToken, Right: right}
}

// --- Unary Expression Visitor Methods (Using Labeled Alternatives) ---

// VisitNotExpr handles the logical NOT (not) unary operation.
func (v *ASTBuilder) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.NOT().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	expr := ctx.UnaryExpr().Accept(v).(Expression)
	return &UnaryExpr{PosToken: opToken.Pos, Operator: opToken, Expr: expr}
}

// VisitBitnotExpr handles the bitwise NOT (~) unary operation.
func (v *ASTBuilder) VisitBitnotExpr(ctx *parser.BitnotExprContext) interface{} {
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.BITNOT().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	expr := ctx.UnaryExpr().Accept(v).(Expression)
	return &UnaryExpr{PosToken: opToken.Pos, Operator: opToken, Expr: expr}
}

// VisitNegExpr handles the negation (-) unary operation.
func (v *ASTBuilder) VisitNegExpr(ctx *parser.NegExprContext) interface{} {
	// Get the ANTLR token for the operator
	antlrOpToken := ctx.SUB().GetSymbol()
	// Create a custom ast.Token from the ANTLR token information
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	expr := ctx.UnaryExpr().Accept(v).(Expression)
	return &UnaryExpr{PosToken: opToken.Pos, Operator: opToken, Expr: expr}
}

// VisitPostfixExpression handles the base case for unary expressions, visiting a postfixExpr.
func (v *ASTBuilder) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	return ctx.PostfixExpr().Accept(v)
}

// --- Postfix Expression Visitor Methods (Using Labeled Alternatives) ---

// VisitPrimaryPostfix handles the base case for postfix expressions, visiting a primary.
func (v *ASTBuilder) VisitPrimaryPostfix(ctx *parser.PrimaryPostfixContext) interface{} {
	return ctx.Primary().Accept(v)
}

// VisitCallPostfix handles a function call postfix operation.
func (v *ASTBuilder) VisitCallPostfix(ctx *parser.CallPostfixContext) interface{} {
	callee := ctx.PostfixExpr().Accept(v).(Expression)
	var args []Expression
	if ctx.ArgList() != nil {
		// VisitArgList returns []Expression
		args = ctx.ArgList().Accept(v).([]Expression)
	}
	lparenToken := ctx.LPAREN().GetSymbol()
	return &CallExpr{PosToken: token.Pos(lparenToken.GetStart()), Callee: callee, Args: args}
}

// VisitIndexPostfix handles an index access postfix operation.
func (v *ASTBuilder) VisitIndexPostfix(ctx *parser.IndexPostfixContext) interface{} {
	primary := ctx.PostfixExpr().Accept(v).(Expression)
	index := ctx.Expression().Accept(v).(Expression)
	lbrackToken := ctx.LBRACK().GetSymbol()
	return &IndexExpr{PosToken: token.Pos(lbrackToken.GetStart()), Primary: primary, Index: index}
}

// VisitAttrPostfix handles an attribute access postfix operation.
func (v *ASTBuilder) VisitAttrPostfix(ctx *parser.AttrPostfixContext) interface{} {
	primary := ctx.PostfixExpr().Accept(v).(Expression)
	attrToken := ctx.IDENTIFIER().GetSymbol()
	dotToken := ctx.DOT().GetSymbol()
	return &AttrExpr{PosToken: token.Pos(dotToken.GetStart()), Primary: primary, Attribute: attrToken.GetText()}
}

// VisitArgList handles a comma-separated list of arguments for a function call.
func (v *ASTBuilder) VisitArgList(ctx *parser.ArgListContext) interface{} {
	var args []Expression
	for _, exprCtx := range ctx.AllExpression() {
		if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
			args = append(args, exprNode)
		}
	}
	return args // Return a slice of Expression nodes
}

// --- Primary Expression Visitor Methods ---

// VisitPrimary handles primary expressions (literals, identifiers, parens, lists, tables).
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	// The primary rule has several alternatives. We check which child context is present.
	switch {
	case ctx.Literal() != nil:
		return ctx.Literal().Accept(v)
	case ctx.IDENTIFIER() != nil:
		idToken := ctx.IDENTIFIER().GetSymbol()
		return &Identifier{PosToken: token.Pos(idToken.GetStart()), Name: idToken.GetText()}
	case ctx.LPAREN() != nil && ctx.RPAREN() != nil:
		// This could be a single expression in parens or a tuple.
		// The grammar is: LPAREN expression RPAREN | LPAREN expression (COMMA expression)+ RPAREN
		// We need to distinguish between a single expression and a tuple.
		exprs := ctx.AllExpression()
		// Check if there is at least one COMMA token to identify a tuple
		if len(exprs) == 1 && ctx.COMMA(0) == nil {
			// Single expression in parens
			return exprs[0].Accept(v)
		} else if len(exprs) > 0 {
			// Tuple - build a ListLiteral for now
			var elements []Expression
			for _, exprCtx := range exprs {
				if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
					elements = append(elements, exprNode)
				}
			}
			// Represent a tuple as a ListLiteral for simplicity, or define a specific TupleLiteral node
			lparenToken := ctx.LPAREN().GetSymbol()
			return &ListLiteral{PosToken: token.Pos(lparenToken.GetStart()), Elements: elements}
		}
		// Should not reach here if grammar is followed, but handle empty parens if needed
		fmt.Printf("WARNING: Empty or invalid parens in primary expression at line %d\n", ctx.GetStart().GetLine())
		return nil // Or an EmptyListLiteral/EmptyTupleLiteral node
	case ctx.ListLiteral() != nil:
		return ctx.ListLiteral().Accept(v)
	case ctx.TableLiteral() != nil:
		return ctx.TableLiteral().Accept(v)
	// Removed the case for ctx.FuncDef() here, as FuncDef is not a direct alternative in the primary rule.
	default:
		// Should not be reached
		fmt.Printf("WARNING: VisitPrimary encountered unexpected child context: %T at line %d\n", ctx.GetChild(0), ctx.GetStart().GetLine())
		return nil
	}
}

// VisitLiteral handles different literal types.
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	// The literal rule has choices: NUMBER, STRING, TRUE, FALSE, NIL
	switch {
	case ctx.NUMBER() != nil:
		numToken := ctx.NUMBER().GetSymbol()
		text := numToken.GetText()
		// Check if it's an integer or float
		if _, err := strconv.ParseInt(text, 10, 64); err == nil {
			val, _ := strconv.ParseInt(text, 10, 64)
			return &IntegerLiteral{PosToken: token.Pos(numToken.GetStart()), Value: val}
		} else if _, err := strconv.ParseFloat(text, 64); err == nil {
			val, _ := strconv.ParseFloat(text, 64)
			return &FloatLiteral{PosToken: token.Pos(numToken.GetStart()), Value: val}
		} else {
			// Handle parsing error if necessary
			fmt.Printf("ERROR: Failed to parse number literal '%s' at line %d\n", text, numToken.GetLine())
			return nil // Or an error node
		}
	case ctx.STRING() != nil:
		strToken := ctx.STRING().GetSymbol()
		text := strToken.GetText()
		// Remove quotes and handle escape sequences (basic handling for now)
		// This requires more sophisticated logic to handle triple quotes and escape sequences correctly.
		// For simplicity, just removing outer quotes.
		value := text
		if len(value) >= 2 {
			// Basic check for single or double quotes
			if (value[0] == '"' && value[len(value)-1] == '"') || (value[0] == '\'' && value[len(value)-1] == '\'') {
				value = value[1 : len(value)-1]
			} else if len(value) >= 6 && ((value[0:3] == `"""` && value[len(value)-3:] == `"""`) || (value[0:3] == `'''` && value[len(value)-3:] == `'''`)) {
				value = value[3 : len(value)-3]
			}
			// TODO: Add proper escape sequence unescaping
		}

		return &StringLiteral{PosToken: token.Pos(strToken.GetStart()), Value: value}
	case ctx.TRUE() != nil:
		trueToken := ctx.TRUE().GetSymbol()
		return &BooleanLiteral{PosToken: token.Pos(trueToken.GetStart()), Value: true}
	case ctx.FALSE() != nil:
		falseToken := ctx.FALSE().GetSymbol()
		return &BooleanLiteral{PosToken: token.Pos(falseToken.GetStart()), Value: false}
	case ctx.NIL() != nil:
		nilToken := ctx.NIL().GetSymbol()
		return &NilLiteral{PosToken: token.Pos(nilToken.GetStart())}
	default:
		// Should not be reached
		fmt.Printf("WARNING: VisitLiteral encountered unexpected child context: %T at line %d\n", ctx.GetChild(0), ctx.GetStart().GetLine())
		return nil
	}
}

// VisitListLiteral builds a ListLiteral node.
func (v *ASTBuilder) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	var elements []Expression
	// The grammar is LBRACK (expression (COMMA expression)*)? RBRACK
	// We need to find all Expression contexts within the list.
	for _, exprCtx := range ctx.AllExpression() {
		if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
			elements = append(elements, exprNode)
		}
	}
	lbrackToken := ctx.LBRACK().GetSymbol()
	return &ListLiteral{PosToken: token.Pos(lbrackToken.GetStart()), Elements: elements}
}

// VisitTableLiteral builds a TableLiteral node.
func (v *ASTBuilder) VisitTableLiteral(ctx *parser.TableLiteralContext) interface{} {
	var fields []TableField
	// The grammar is LBRACE (tableKeyValue (COMMA tableKeyValue)*)? RBRACE
	for _, fieldCtx := range ctx.AllTableKeyValue() {
		if fieldNode, ok := fieldCtx.Accept(v).(TableField); ok {
			fields = append(fields, fieldNode)
		}
	}
	lbraceToken := ctx.LBRACE().GetSymbol()
	return &TableLiteral{PosToken: token.Pos(lbraceToken.GetStart()), Fields: fields}
}

// VisitTableKeyValue builds a TableField node for a table literal.
func (v *ASTBuilder) VisitTableKeyValue(ctx *parser.TableKeyValueContext) interface{} {
	// The grammar is tableKey ASSIGN expression
	key := ctx.TableKey().Accept(v).(Expression) // Visit the tableKey rule
	value := ctx.Expression().Accept(v).(Expression)
	assignToken := ctx.ASSIGN().GetSymbol()

	// Ensure the key is a valid type for a table key (Expression, StringLiteral, or Identifier)
	// The VisitTableKey should handle this, but a check here adds safety.
	switch key.(type) {
	case *Identifier, *StringLiteral, Expression: // Expression covers other cases allowed by tableKey
		// Valid key types
	default:
		fmt.Printf("ERROR: Invalid table key type %T at line %d\n", key, ctx.GetStart().GetLine())
		// Return a placeholder or handle error
		return TableField{}
	}

	return TableField{
		PosToken: token.Pos(assignToken.GetStart()),
		Key:      key,
		Value:    value,
	}
}

// VisitTableKey handles the key part of a table key-value pair.
func (v *ASTBuilder) VisitTableKey(ctx *parser.TableKeyContext) interface{} {
	// The grammar is expression | STRING | IDENTIFIER
	switch {
	case ctx.Expression() != nil:
		return ctx.Expression().Accept(v)
	case ctx.STRING() != nil:
		// VisitLiteral will handle the string token and return a StringLiteral
		return ctx.STRING().Accept(v)
	case ctx.IDENTIFIER() != nil:
		// Build an Identifier node directly
		idToken := ctx.IDENTIFIER().GetSymbol()
		return &Identifier{PosToken: token.Pos(idToken.GetStart()), Name: idToken.GetText()}
	default:
		// Should not be reached
		fmt.Printf("WARNING: VisitTableKey encountered unexpected child context: %T at line %d\n", ctx.GetChild(0), ctx.GetStart().GetLine())
		return nil
	}
}

// Removed the duplicate VisitFuncDef method that was intended for expression context.
// The single VisitFuncDef method above handles the statement case correctly.

// Ensure ASTBuilder implements the generated visitor interface.
var _ parser.InscriptVisitor = (*ASTBuilder)(nil)
