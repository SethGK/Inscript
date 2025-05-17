// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Inscript

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by InscriptParser.
type InscriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by InscriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by InscriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by InscriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by InscriptParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by InscriptParser#target.
	VisitTarget(ctx *TargetContext) interface{}

	// Visit a parse tree produced by InscriptParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#funcDef.
	VisitFuncDef(ctx *FuncDefContext) interface{}

	// Visit a parse tree produced by InscriptParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by InscriptParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by InscriptParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by InscriptParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#geExpr.
	VisitGeExpr(ctx *GeExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#modExpr.
	VisitModExpr(ctx *ModExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#gtExpr.
	VisitGtExpr(ctx *GtExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#idivExpr.
	VisitIdivExpr(ctx *IdivExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#subExpr.
	VisitSubExpr(ctx *SubExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#neqExpr.
	VisitNeqExpr(ctx *NeqExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#ltExpr.
	VisitLtExpr(ctx *LtExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#eqExpr.
	VisitEqExpr(ctx *EqExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#expExpr.
	VisitExpExpr(ctx *ExpExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#bitandExpr.
	VisitBitandExpr(ctx *BitandExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#leExpr.
	VisitLeExpr(ctx *LeExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#bitorExpr.
	VisitBitorExpr(ctx *BitorExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#bitxorExpr.
	VisitBitxorExpr(ctx *BitxorExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#divExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#shlExpr.
	VisitShlExpr(ctx *ShlExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by InscriptParser#shrExpr.
	VisitShrExpr(ctx *ShrExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#bitnotExpr.
	VisitBitnotExpr(ctx *BitnotExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#negExpr.
	VisitNegExpr(ctx *NegExprContext) interface{}

	// Visit a parse tree produced by InscriptParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by InscriptParser#primaryPostfix.
	VisitPrimaryPostfix(ctx *PrimaryPostfixContext) interface{}

	// Visit a parse tree produced by InscriptParser#indexPostfix.
	VisitIndexPostfix(ctx *IndexPostfixContext) interface{}

	// Visit a parse tree produced by InscriptParser#attrPostfix.
	VisitAttrPostfix(ctx *AttrPostfixContext) interface{}

	// Visit a parse tree produced by InscriptParser#callPostfix.
	VisitCallPostfix(ctx *CallPostfixContext) interface{}

	// Visit a parse tree produced by InscriptParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by InscriptParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by InscriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by InscriptParser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by InscriptParser#tableLiteral.
	VisitTableLiteral(ctx *TableLiteralContext) interface{}

	// Visit a parse tree produced by InscriptParser#tableKeyValue.
	VisitTableKeyValue(ctx *TableKeyValueContext) interface{}

	// Visit a parse tree produced by InscriptParser#tableKey.
	VisitTableKey(ctx *TableKeyContext) interface{}
}
