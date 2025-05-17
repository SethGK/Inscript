// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Inscript

import "github.com/antlr4-go/antlr/v4"

type BaseInscriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseInscriptVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTarget(ctx *TargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitFuncDef(ctx *FuncDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitGeExpr(ctx *GeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitModExpr(ctx *ModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitGtExpr(ctx *GtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitIdivExpr(ctx *IdivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitSubExpr(ctx *SubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitNeqExpr(ctx *NeqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitLtExpr(ctx *LtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitEqExpr(ctx *EqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExpExpr(ctx *ExpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBitandExpr(ctx *BitandExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitLeExpr(ctx *LeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBitorExpr(ctx *BitorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBitxorExpr(ctx *BitxorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitShlExpr(ctx *ShlExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitShrExpr(ctx *ShrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBitnotExpr(ctx *BitnotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitNegExpr(ctx *NegExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitPrimaryPostfix(ctx *PrimaryPostfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitIndexPostfix(ctx *IndexPostfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitAttrPostfix(ctx *AttrPostfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitCallPostfix(ctx *CallPostfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitListLiteral(ctx *ListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTableLiteral(ctx *TableLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTableKeyValue(ctx *TableKeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTableKey(ctx *TableKeyContext) interface{} {
	return v.VisitChildren(ctx)
}
