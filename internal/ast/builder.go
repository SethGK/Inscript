package ast

import (
	"fmt"
	"go/token"
	"strconv"

	parser "github.com/SethGK/Inscript/parser/grammar"
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
		if stmtNode, ok := stmtCtx.Accept(v).(Statement); ok && stmtNode != nil {
			prog.Stmts = append(prog.Stmts, stmtNode)
		}
	}
	return prog
}

// VisitStatement visits a statement rule and dispatches to the appropriate handler.
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
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
		fmt.Printf("WARNING: VisitStatement encountered unexpected child context: %T at line %d\n", ctx.GetChild(0), ctx.GetStart().GetLine())
		return nil
	}
}

// VisitBlock builds a BlockStmt node.
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	block := &BlockStmt{PosToken: token.Pos(ctx.GetStart().GetStart())}
	for i, stmtCtx := range ctx.AllStatement() {
		if stmtNode, ok := stmtCtx.Accept(v).(Statement); ok && stmtNode != nil {
			block.Stmts = append(block.Stmts, stmtNode)
			if _, isRet := stmtNode.(*ReturnStmt); isRet {
				remaining := ctx.AllStatement()[i+1:]
				if len(remaining) > 0 {
					fmt.Printf("WARNING: Unreachable code after return in block starting at line %d\n", stmtCtx.GetStart().GetLine())
				}
				break
			}
		}
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
	target := ctx.Target().Accept(v).(Expression)
	antlrOpToken := ctx.GetChild(1).(antlr.TerminalNode).GetSymbol()
	opToken := Token{
		Type:    antlrOpToken.GetTokenType(),
		Pos:     token.Pos(antlrOpToken.GetStart()),
		Literal: antlrOpToken.GetText(),
	}
	value := ctx.Expression().Accept(v).(Expression)
	switch target.(type) {
	case *Identifier, *IndexExpr, *AttrExpr:
		// Valid target types
	default:
		fmt.Printf("ERROR: Invalid assignment target type %T at line %d\n", target, ctx.GetStart().GetLine())
		return nil
	}

	return &AssignStmt{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Target:   target,
		Op:       opToken,
		Value:    value,
	}
}

// VisitTarget builds the Expression node for the assignment target.
func (v *ASTBuilder) VisitTarget(ctx *parser.TargetContext) interface{} {
	if ctx.IDENTIFIER() != nil && ctx.PostfixExpr() == nil {
		idToken := ctx.IDENTIFIER().GetSymbol()
		return &Identifier{PosToken: token.Pos(idToken.GetStart()), Name: idToken.GetText()}
	} else if ctx.PostfixExpr() != nil {
		primary := ctx.PostfixExpr().Accept(v).(Expression)
		if ctx.LBRACK() != nil {
			index := ctx.Expression().Accept(v).(Expression)
			lbrackToken := ctx.LBRACK().GetSymbol()
			return &IndexExpr{PosToken: token.Pos(lbrackToken.GetStart()), Primary: primary, Index: index}
		} else if ctx.DOT() != nil {
			attrToken := ctx.IDENTIFIER().GetSymbol()
			dotToken := ctx.DOT().GetSymbol()
			return &AttrExpr{PosToken: token.Pos(dotToken.GetStart()), Primary: primary, Attribute: attrToken.GetText()}
		}
	}
	fmt.Printf("WARNING: VisitTarget encountered unexpected context: %T at line %d\n", ctx, ctx.GetStart().GetLine())
	return nil
}

// VisitIfStmt builds an IfStmt node.
func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := ctx.Expression().Accept(v).(Expression)
	thenBlock := ctx.Block(0).Accept(v).(*BlockStmt)

	var elseIfs []ElseIf
	var elseBlock *BlockStmt
	if ctx.ELSE() != nil {
		if len(ctx.AllBlock()) > 1 {
			elseBlock = ctx.Block(1).Accept(v).(*BlockStmt)
		} else {
			fmt.Printf("ERROR: ELSE keyword without a following block at line %d\n", ctx.ELSE().GetSymbol().GetLine())
			elseBlock = nil
		}
	}

	return &IfStmt{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Cond:     cond,
		Then:     thenBlock,
		ElseIfs:  elseIfs,
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
		params = ctx.ParamList().Accept(v).([]Param)
	}

	var returnType *TypeAnnotation
	if ctx.TypeAnnotation() != nil {
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
		param := paramCtx.Accept(v).(Param)
		params = append(params, param)
	}
	return params
}

// VisitParam handles a single function parameter.
func (v *ASTBuilder) VisitParam(ctx *parser.ParamContext) interface{} {
	posToken := token.Pos(ctx.GetStart().GetStart())
	isVariadic := ctx.ELLIPSIS() != nil

	var name string
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
		posToken = token.Pos(ctx.IDENTIFIER().GetSymbol().GetStart())
	} else if isVariadic && ctx.ELLIPSIS() != nil && ctx.GetChildCount() > 1 {
		if idToken, ok := ctx.GetChild(1).(antlr.TerminalNode); ok && idToken.GetSymbol().GetTokenType() == parser.InscriptParserIDENTIFIER {
			name = idToken.GetText()
			posToken = token.Pos(ctx.ELLIPSIS().GetSymbol().GetStart())
		} else {
			fmt.Printf("ERROR: Expected identifier after '...' in parameter list at line %d\n", ctx.GetStart().GetLine())
			name = ""
		}
	} else {
		fmt.Printf("ERROR: Invalid parameter definition at line %d\n", ctx.GetStart().GetLine())
		name = ""
	}

	var defaultValue Expression
	if ctx.ASSIGN() != nil && ctx.Expression() != nil { // Corrected: ctx.Expression() is fine here as it's singular
		defaultValue = ctx.Expression().Accept(v).(Expression) // No index needed for singular optional expression
	}

	var paramType *TypeAnnotation
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
	if ctx.Expression() != nil {
		expr = ctx.Expression().Accept(v).(Expression)
	}
	return &ReturnStmt{PosToken: token.Pos(ctx.GetStart().GetStart()), Expr: expr}
}

// VisitImportStmt builds an ImportStmt node.
func (v *ASTBuilder) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	stringToken := ctx.STRING().GetSymbol()
	path := stringToken.GetText()
	if len(path) >= 2 && (path[0] == '"' && path[len(path)-1] == '"' || path[0] == '\'' && path[len(path)-1] == '\'') {
		path = path[1 : len(path)-1]
	} else if len(path) >= 6 && (path[0:3] == `"""` && path[len(path)-3:] == `"""` || path[0:3] == `'''` && path[len(path)-3:] == `'''`) {
		path = path[3 : len(path)-3]
	}
	return &ImportStmt{
		PosToken: token.Pos(ctx.GetStart().GetStart()),
		Path:     path,
	}
}

// VisitPrintStmt builds a PrintStmt node.
func (v *ASTBuilder) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	var exprs []Expression
	for _, exprCtx := range ctx.AllExpression() { // Corrected: Use AllExpression()
		if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
			exprs = append(exprs, exprNode)
		}
	}
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
	antlrOpToken := ctx.POW().GetSymbol()
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
	antlrOpToken := ctx.MUL().GetSymbol()
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
	antlrOpToken := ctx.DIV().GetSymbol()
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
	antlrOpToken := ctx.IDIV().GetSymbol()
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
	antlrOpToken := ctx.MOD().GetSymbol()
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
	antlrOpToken := ctx.ADD().GetSymbol()
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
	antlrOpToken := ctx.SUB().GetSymbol()
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
	antlrOpToken := ctx.BITAND().GetSymbol()
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
	antlrOpToken := ctx.BITOR().GetSymbol()
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
	antlrOpToken := ctx.BITXOR().GetSymbol()
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
	antlrOpToken := ctx.SHL().GetSymbol()
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
	antlrOpToken := ctx.SHR().GetSymbol()
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
	antlrOpToken := ctx.LT().GetSymbol()
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
	antlrOpToken := ctx.LE().GetSymbol()
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
	antlrOpToken := ctx.GT().GetSymbol()
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
	antlrOpToken := ctx.GE().GetSymbol()
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
	antlrOpToken := ctx.EQ().GetSymbol()
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
	antlrOpToken := ctx.NEQ().GetSymbol()
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
	antlrOpToken := ctx.AND().GetSymbol()
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
	antlrOpToken := ctx.OR().GetSymbol()
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
	antlrOpToken := ctx.NOT().GetSymbol()
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
	antlrOpToken := ctx.BITNOT().GetSymbol()
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
	antlrOpToken := ctx.SUB().GetSymbol()
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
		args = ctx.ArgList().Accept(v).([]Expression)
	}
	return &CallExpr{
		PosToken: token.Pos(ctx.LPAREN().GetSymbol().GetStart()),
		Callee:   callee,
		Args:     args,
	}
}

// VisitIndexPostfix handles an index access postfix operation.
func (v *ASTBuilder) VisitIndexPostfix(ctx *parser.IndexPostfixContext) interface{} {
	primary := ctx.PostfixExpr().Accept(v).(Expression)
	index := ctx.Expression().Accept(v).(Expression)
	return &IndexExpr{
		PosToken: token.Pos(ctx.LBRACK().GetSymbol().GetStart()),
		Primary:  primary,
		Index:    index,
	}
}

// VisitAttrPostfix handles an attribute access postfix operation.
func (v *ASTBuilder) VisitAttrPostfix(ctx *parser.AttrPostfixContext) interface{} {
	primary := ctx.PostfixExpr().Accept(v).(Expression)
	attrToken := ctx.IDENTIFIER().GetSymbol()
	return &AttrExpr{
		PosToken:  token.Pos(ctx.DOT().GetSymbol().GetStart()),
		Primary:   primary,
		Attribute: attrToken.GetText(),
	}
}

// VisitArgList handles a comma-separated list of arguments.
func (v *ASTBuilder) VisitArgList(ctx *parser.ArgListContext) interface{} {
	var args []Expression
	for _, exprCtx := range ctx.AllExpression() { // Corrected: Use AllExpression()
		if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
			args = append(args, exprNode)
		}
	}
	return args
}

// --- Primary Expression Visitor Methods ---

// VisitLiteral handles literal expressions.
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.NUMBER() != nil {
		numStr := ctx.NUMBER().GetText()
		if i, err := strconv.ParseInt(numStr, 10, 64); err == nil {
			return &IntegerLiteral{PosToken: token.Pos(ctx.NUMBER().GetSymbol().GetStart()), Value: i}
		} else if f, err := strconv.ParseFloat(numStr, 64); err == nil {
			return &FloatLiteral{PosToken: token.Pos(ctx.NUMBER().GetSymbol().GetStart()), Value: f}
		}
		fmt.Printf("ERROR: Could not parse number literal '%s' at line %d\n", numStr, ctx.NUMBER().GetSymbol().GetLine())
		return nil
	} else if ctx.STRING() != nil {
		strToken := ctx.STRING().GetSymbol()
		val := strToken.GetText()
		if len(val) >= 2 && (val[0] == '"' && val[len(val)-1] == '"' || val[0] == '\'' && val[len(val)-1] == '\'') {
			val = val[1 : len(val)-1]
		} else if len(val) >= 6 && (val[0:3] == `"""` && val[len(val)-3:] == `"""` || val[0:3] == `'''` && val[len(val)-3:] == `'''`) {
			val = val[3 : len(val)-3]
		}
		return &StringLiteral{PosToken: token.Pos(strToken.GetStart()), Value: val}
	} else if ctx.TRUE() != nil {
		return &BooleanLiteral{PosToken: token.Pos(ctx.TRUE().GetSymbol().GetStart()), Value: true}
	} else if ctx.FALSE() != nil {
		return &BooleanLiteral{PosToken: token.Pos(ctx.FALSE().GetSymbol().GetStart()), Value: false}
	} else if ctx.NIL() != nil {
		return &NilLiteral{PosToken: token.Pos(ctx.NIL().GetSymbol().GetStart())}
	}
	return nil
}

// VisitTerminal handles terminal nodes (like identifiers within primary).
func (v *ASTBuilder) VisitTerminal(node antlr.TerminalNode) interface{} {
	if node.GetSymbol().GetTokenType() == parser.InscriptParserIDENTIFIER {
		return &Identifier{PosToken: token.Pos(node.GetSymbol().GetStart()), Name: node.GetText()}
	}
	return nil
}

// VisitPrimary handles primary expressions.
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(v)
	} else if ctx.IDENTIFIER() != nil {
		idToken := ctx.IDENTIFIER().GetSymbol()
		return &Identifier{PosToken: token.Pos(idToken.GetStart()), Name: idToken.GetText()}
	} else if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		// Check if it's a single expression in parentheses or a tuple
		allExprs := ctx.AllExpression() // Get all expressions in parentheses
		if len(allExprs) == 1 && ctx.COMMA(0) == nil {
			return allExprs[0].Accept(v) // Single expression
		} else if len(allExprs) > 0 { // If there's at least one expression and it's not a single expr (implying commas)
			elements := make([]Expression, len(allExprs))
			for i, exprCtx := range allExprs { // Corrected: Iterate over allExprs
				elements[i] = exprCtx.Accept(v).(Expression)
			}
			return &TupleLiteral{PosToken: token.Pos(ctx.LPAREN().GetSymbol().GetStart()), Elements: elements}
		}
	} else if ctx.ListLiteral() != nil {
		return ctx.ListLiteral().Accept(v)
	} else if ctx.TableLiteral() != nil {
		return ctx.TableLiteral().Accept(v)
	}
	return nil
}

// VisitListLiteral builds a ListLiteral node.
func (v *ASTBuilder) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	var elements []Expression
	for _, exprCtx := range ctx.AllExpression() { // Corrected: Use AllExpression()
		if exprNode, ok := exprCtx.Accept(v).(Expression); ok {
			elements = append(elements, exprNode)
		}
	}
	return &ListLiteral{PosToken: token.Pos(ctx.LBRACK().GetSymbol().GetStart()), Elements: elements}
}

// VisitTableLiteral builds a TableLiteral node.
func (v *ASTBuilder) VisitTableLiteral(ctx *parser.TableLiteralContext) interface{} {
	var fields []*TableField
	for _, kvCtx := range ctx.AllTableKeyValue() {
		if field, ok := kvCtx.Accept(v).(*TableField); ok {
			fields = append(fields, field)
		}
	}
	return &TableLiteral{PosToken: token.Pos(ctx.LBRACE().GetSymbol().GetStart()), Fields: fields}
}

// VisitTableKeyValue builds a TableField node.
func (v *ASTBuilder) VisitTableKeyValue(ctx *parser.TableKeyValueContext) interface{} {
	key := ctx.TableKey().Accept(v).(Expression)
	value := ctx.Expression().Accept(v).(Expression) // Corrected: No index needed for singular expression
	return &TableField{
		PosToken: token.Pos(ctx.TableKey().GetStart().GetStart()),
		Key:      key,
		Value:    value,
	}
}

// VisitTableKey builds the Expression for a table key.
func (v *ASTBuilder) VisitTableKey(ctx *parser.TableKeyContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(v)
	} else if ctx.STRING() != nil {
		strToken := ctx.STRING().GetSymbol()
		val := strToken.GetText()
		if len(val) >= 2 && (val[0] == '"' && val[len(val)-1] == '"' || val[0] == '\'' && val[len(val)-1] == '\'') {
			val = val[1 : len(val)-1]
		} else if len(val) >= 6 && (val[0:3] == `"""` && val[len(val)-3:] == `"""` || val[0:3] == `'''` && val[len(val)-3:] == `'''`) {
			val = val[3 : len(val)-3]
		}
		return &StringLiteral{PosToken: token.Pos(strToken.GetStart()), Value: val}
	} else if ctx.IDENTIFIER() != nil {
		idToken := ctx.IDENTIFIER().GetSymbol()
		return &Identifier{PosToken: token.Pos(idToken.GetStart()), Name: idToken.GetText()}
	}
	return nil
}
