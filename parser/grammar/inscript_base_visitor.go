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

func (v *BaseInscriptVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitCompoundStmt(ctx *CompoundStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitElseifListOpt(ctx *ElseifListOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitElseif(ctx *ElseifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitElseBlockOpt(ctx *ElseBlockOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitFunctionDef(ctx *FunctionDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitLogicalOr(ctx *LogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitLogicalAnd(ctx *LogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitArith(ctx *ArithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitFnLiteral(ctx *FnLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitListLiteral(ctx *ListLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitTableLiteral(ctx *TableLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitFieldListOpt(ctx *FieldListOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExpressionOpt(ctx *ExpressionOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExpressionListOpt(ctx *ExpressionListOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitParamListOpt(ctx *ParamListOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInscriptVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
