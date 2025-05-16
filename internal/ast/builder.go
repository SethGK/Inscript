package ast

import (
	"fmt" // Import fmt for debugging prints
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
		stmt := stmtCtx.Accept(v).(Statement)
		prog.Stmts = append(prog.Stmts, stmt)
	}
	return prog
}

// VisitStatement dispatches a StatementContext to either your SimpleStmt or CompoundStmt visitor
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.SimpleStmt() != nil {
		return ctx.SimpleStmt().Accept(v)
	}
	if ctx.CompoundStmt() != nil {
		return ctx.CompoundStmt().Accept(v)
	}
	panic(fmt.Sprintf("unexpected statement node: %s", ctx.GetText()))
}

// VisitSimpleStmt handles simple statements.
func (v *ASTBuilder) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting SimpleStmt")
	if ctx.Assignment() != nil {
		result := ctx.Assignment().Accept(v)
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (Assignment), returning %T\n", result)
		return result
	}
	if ctx.ExprStmt() != nil {
		result := ctx.ExprStmt().Accept(v)
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (ExprStmt), returning %T\n", result)
		return result
	}
	if ctx.PrintStmt() != nil {
		result := ctx.PrintStmt().Accept(v)
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (PrintStmt), returning %T\n", result)
		return result
	}
	if ctx.ReturnStmt() != nil {
		result := ctx.ReturnStmt().Accept(v)
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (ReturnStmt), returning %T\n", result)
		return result
	}
	fmt.Printf("WARNING: Visiting SimpleStmt, unexpected context: %T\n", ctx)
	return nil
}

// VisitCompoundStmt handles compound statements.
func (v *ASTBuilder) VisitCompoundStmt(ctx *parser.CompoundStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting CompoundStmt")
	if fd := ctx.FunctionDef(); fd != nil {
		fnExpr := fd.Accept(v).(Expression)
		return &ExprStmt{Expr: fnExpr, PosToken: token.Pos(fd.GetStart().GetStart())}
	}
	if ctx.IfStmt() != nil {
		return ctx.IfStmt().Accept(v)
	}
	if ctx.WhileStmt() != nil {
		return ctx.WhileStmt().Accept(v)
	}
	if ctx.ForStmt() != nil {
		return ctx.ForStmt().Accept(v)
	}
	fmt.Printf("WARNING: Visiting CompoundStmt, unexpected context: %T\n", ctx)
	return nil
}

// VisitAssignment builds an AssignStmt node.
func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	raw := ctx.Primary().Accept(v)
	target, ok := raw.(Expression)
	if !ok || target == nil {
		panic(fmt.Sprintf("VisitAssignment: invalid target %T", raw))
	}
	value := ctx.Expression().Accept(v).(Expression)
	return &AssignStmt{Target: target, Value: value, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitExprStmt builds an ExprStmt node.
func (v *ASTBuilder) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	expr := ctx.Expression().Accept(v).(Expression)
	return &ExprStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitPrintStmt builds a PrintStmt node.
func (v *ASTBuilder) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting PrintStmt")
	var exprs []Expression
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		if visited, ok := el.Accept(v).([]Expression); ok {
			exprs = visited
		}
	}
	return &PrintStmt{Exprs: exprs, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitReturnStmt builds a ReturnStmt node.
func (v *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	var expr Expression
	if ctx.ExpressionOpt().Expression() != nil {
		expr = ctx.ExpressionOpt().Expression().Accept(v).(Expression)
	}
	return &ReturnStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitIfStmt builds an IfStmt node.
func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := ctx.Expression().Accept(v).(Expression)
	then := ctx.Block().Accept(v).(*BlockStmt)
	var elseIfs []ElseIf
	if elOpt := ctx.ElseifListOpt(); elOpt != nil {
		for _, eis := range elOpt.AllElseif() {
			econd := eis.Expression().Accept(v).(Expression)
			body := eis.Block().Accept(v).(*BlockStmt)
			elseIfs = append(elseIfs, ElseIf{Cond: econd, Body: body, PosToken: token.Pos(eis.GetStart().GetStart())})
		}
	}
	var elseBlk *BlockStmt
	if ebOpt := ctx.ElseBlockOpt(); ebOpt != nil && ebOpt.Block() != nil {
		elseBlk = ebOpt.Block().Accept(v).(*BlockStmt)
	}
	return &IfStmt{Cond: cond, Then: then, ElseIfs: elseIfs, Else: elseBlk, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitWhileStmt builds a WhileStmt node.
func (v *ASTBuilder) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting WhileStmt")
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &WhileStmt{Cond: cond, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitForStmt builds a ForStmt node.
func (v *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting ForStmt")
	varName := ctx.IDENTIFIER().GetText()
	iter := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &ForStmt{Variable: varName, Iterable: iter, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitFunctionDef builds a FunctionLiteral node.
func (v *ASTBuilder) VisitFunctionDef(ctx *parser.FunctionDefContext) interface{} {
	fmt.Println("DEBUG: Visiting FunctionDef (as Expression)")
	var params []string
	if pl := ctx.ParamListOpt().ParamList(); pl != nil {
		for _, id := range pl.AllIDENTIFIER() {
			params = append(params, id.GetText())
		}
	}
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &FunctionLiteral{Params: params, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitBlock builds a BlockStmt node.
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	block := &BlockStmt{PosToken: token.Pos(ctx.GetStart().GetStart())}
	for i, stmtCtx := range ctx.AllStatement() {
		stmt := stmtCtx.Accept(v).(Statement)
		block.Stmts = append(block.Stmts, stmt)
		if _, isRet := stmt.(*ReturnStmt); isRet {
			remaining := ctx.AllStatement()[i+1:]
			if len(remaining) > 0 {
				fmt.Printf("WARNING: Unreachable code after return at line %d\n", stmtCtx.GetStart().GetLine())
			}
			break
		}
	}
	return block
}

// VisitExpressionList handles a comma-separated list of expressions.
func (v *ASTBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	var exprs []Expression
	for _, e := range ctx.AllExpression() {
		exprs = append(exprs, e.Accept(v).(Expression))
	}
	return exprs
}

// VisitParamList handles a comma-separated list for function parameters.
func (v *ASTBuilder) VisitParamList(ctx *parser.ParamListContext) interface{} {
	var params []string
	for _, id := range ctx.AllIDENTIFIER() {
		params = append(params, id.GetText())
	}
	return params
}

// VisitExpression simply visits the top-level logicalOr.
func (v *ASTBuilder) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return ctx.LogicalOr().Accept(v)
}

// VisitLogicalOr handles 'or' operations.
func (v *ASTBuilder) VisitLogicalOr(ctx *parser.LogicalOrContext) interface{} {
	expr := ctx.LogicalAnd(0).Accept(v).(Expression)
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__14 { // 'or' ⇒ T__14
			right := ctx.LogicalAnd((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitLogicalAnd handles 'and' operations.
func (v *ASTBuilder) VisitLogicalAnd(ctx *parser.LogicalAndContext) interface{} {
	expr := ctx.Comparison(0).Accept(v).(Expression)
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__15 { // 'and' ⇒ T__15
			right := ctx.Comparison((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitComparison handles comparison operations.
func (v *ASTBuilder) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	expr := ctx.Arith(0).Accept(v).(Expression)
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__17, parser.InscriptParserT__18,
			parser.InscriptParserT__19, parser.InscriptParserT__20,
			parser.InscriptParserT__21, parser.InscriptParserT__22:
			right := ctx.Arith((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitArith handles '+' and '-' operations.
func (v *ASTBuilder) VisitArith(ctx *parser.ArithContext) interface{} {
	expr := ctx.Term(0).Accept(v).(Expression)
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__22: // '+' ⇒ T__22
			right := ctx.Term((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: "+", Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		case parser.InscriptParserT__23: // '-' ⇒ T__23
			right := ctx.Term((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: "-", Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitTerm handles '*', '/', '%' operations.
func (v *ASTBuilder) VisitTerm(ctx *parser.TermContext) interface{} {
	expr := ctx.Factor(0).Accept(v).(Expression)
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__24: // '*' ⇒ T__24
			right := ctx.Factor((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: "*", Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		case parser.InscriptParserT__25: // '/' ⇒ T__25
			right := ctx.Factor((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: "/", Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		case parser.InscriptParserT__26: // '%' ⇒ T__26
			right := ctx.Factor((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: "%", Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitFactor handles '^' operations.
func (v *ASTBuilder) VisitFactor(ctx *parser.FactorContext) interface{} {
	expr := ctx.Unary(0).Accept(v).(Expression)
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__27 { // '^' ⇒ T__27
			right := ctx.Unary((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: "^", Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitUnary handles unary '+', '-', 'not'.
func (v *ASTBuilder) VisitUnary(ctx *parser.UnaryContext) interface{} {
	if ctx.GetChildCount() == 2 {
		opNode := ctx.GetChild(0).(antlr.TerminalNode)
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__22: // '+' ⇒ T__22
			expr := ctx.Unary().Accept(v).(Expression)
			return &UnaryExpr{Operator: "+", Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		case parser.InscriptParserT__23: // '-' ⇒ T__23
			expr := ctx.Unary().Accept(v).(Expression)
			return &UnaryExpr{Operator: "-", Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		case parser.InscriptParserT__29: // 'not' ⇒ T__29
			expr := ctx.Unary().Accept(v).(Expression)
			return &UnaryExpr{Operator: "not", Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return ctx.Primary().Accept(v)
}

// VisitPrimary handles primary expressions.
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	atomVal := ctx.Atom().Accept(v)
	if atomVal == nil {
		return nil
	}
	expr := atomVal.(Expression)
	for i := 1; i < ctx.GetChildCount(); i++ {
		if term, ok := ctx.GetChild(i).(antlr.TerminalNode); ok && term.GetSymbol().GetText() == "[" {
			index := ctx.Expression((i + 1) / 2).Accept(v).(Expression)
			expr = &IndexExpr{Primary: expr, Index: index, PosToken: token.Pos(term.GetSymbol().GetStart())}
			i += 2
		}
		if term, ok := ctx.GetChild(i).(antlr.TerminalNode); ok && term.GetSymbol().GetText() == "(" {
			argsCtx := ctx.GetChild(i + 1).(*parser.ExpressionListOptContext)
			var args []Expression
			if el := argsCtx.ExpressionList(); el != nil {
				args = el.Accept(v).([]Expression)
			}
			expr = &CallExpr{Callee: expr, Args: args, PosToken: token.Pos(term.GetSymbol().GetStart())}
			i += 2
		}
	}
	return expr
}

// VisitAtom handles literals, identifiers, lists, tables, parens, function literals.
func (v *ASTBuilder) VisitAtom(ctx *parser.AtomContext) interface{} {
	fmt.Printf("DEBUG: VisitAtom - text: %s\n", ctx.GetText())
	if lit := ctx.Literal(); lit != nil {
		return lit.Accept(v)
	}
	if id := ctx.IDENTIFIER(); id != nil {
		return &Identifier{Name: id.GetText(), PosToken: token.Pos(id.GetSymbol().GetStart())}
	}
	if ll := ctx.ListLiteral(); ll != nil {
		return ll.Accept(v)
	}
	if tl := ctx.TableLiteral(); tl != nil {
		return tl.Accept(v)
	}
	if fn := ctx.FnLiteral(); fn != nil {
		return fn.Accept(v)
	}
	if expr := ctx.Expression(); expr != nil {
		return expr.Accept(v)
	}
	return nil
}

// VisitFnLiteral builds a FunctionLiteral from fnLiteral.
func (v *ASTBuilder) VisitFnLiteral(ctx *parser.FnLiteralContext) interface{} {
	var params []string
	if pl := ctx.ParamListOpt().ParamList(); pl != nil {
		for _, id := range pl.AllIDENTIFIER() {
			params = append(params, id.GetText())
		}
	}
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &FunctionLiteral{Params: params, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitListLiteral builds a ListLiteral.
func (v *ASTBuilder) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	var elems []Expression
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		elems = el.Accept(v).([]Expression)
	}
	return &ListLiteral{Elements: elems, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitTableLiteral builds a TableLiteral.
func (v *ASTBuilder) VisitTableLiteral(ctx *parser.TableLiteralContext) interface{} {
	var fields []Field
	if fl := ctx.FieldListOpt(); fl != nil {
		for _, fctx := range fl.AllField() {
			if f, ok := fctx.Accept(v).(Field); ok {
				fields = append(fields, f)
			}
		}
	}
	return &TableLiteral{Fields: fields, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitField builds a Field for table literals.
func (v *ASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	value := ctx.Expression().Accept(v).(Expression)
	return Field{Key: ctx.IDENTIFIER().GetText(), Value: value, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitLiteral handles integer, float, string, boolean, nil.
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if i := ctx.INTEGER(); i != nil {
		val, _ := strconv.ParseInt(i.GetText(), 10, 64)
		return &IntegerLiteral{Value: val, PosToken: token.Pos(i.GetSymbol().GetStart())}
	}
	if f := ctx.FLOAT(); f != nil {
		val, _ := strconv.ParseFloat(f.GetText(), 64)
		return &FloatLiteral{Value: val, PosToken: token.Pos(f.GetSymbol().GetStart())}
	}
	if s := ctx.STRING(); s != nil {
		t := s.GetText()
		return &StringLiteral{Value: t[1 : len(t)-1], PosToken: token.Pos(s.GetSymbol().GetStart())}
	}
	if b := ctx.BOOLEAN(); b != nil {
		val, _ := strconv.ParseBool(b.GetText())
		return &BooleanLiteral{Value: val, PosToken: token.Pos(b.GetSymbol().GetStart())}
	}
	if nilTok := ctx.GetToken(parser.InscriptParserT__32, 0); nilTok != nil {
		return &NilLiteral{PosToken: token.Pos(nilTok.GetSymbol().GetStart())}
	}
	return nil
}

var _ parser.InscriptVisitor = (*ASTBuilder)(nil)
