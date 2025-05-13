// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Inscript

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by InscriptParser.
type InscriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by InscriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by InscriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by InscriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by InscriptParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by InscriptParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#compoundStmt.
	VisitCompoundStmt(ctx *CompoundStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#elseifListOpt.
	VisitElseifListOpt(ctx *ElseifListOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#elseifList.
	VisitElseifList(ctx *ElseifListContext) interface{}

	// Visit a parse tree produced by InscriptParser#elseif.
	VisitElseif(ctx *ElseifContext) interface{}

	// Visit a parse tree produced by InscriptParser#elseBlockOpt.
	VisitElseBlockOpt(ctx *ElseBlockOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by InscriptParser#functionDef.
	VisitFunctionDef(ctx *FunctionDefContext) interface{}

	// Visit a parse tree produced by InscriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by InscriptParser#statementListOpt.
	VisitStatementListOpt(ctx *StatementListOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#expressionOpt.
	VisitExpressionOpt(ctx *ExpressionOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#expressionListOpt.
	VisitExpressionListOpt(ctx *ExpressionListOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by InscriptParser#paramListOpt.
	VisitParamListOpt(ctx *ParamListOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by InscriptParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by InscriptParser#logicalOr.
	VisitLogicalOr(ctx *LogicalOrContext) interface{}

	// Visit a parse tree produced by InscriptParser#logicalAnd.
	VisitLogicalAnd(ctx *LogicalAndContext) interface{}

	// Visit a parse tree produced by InscriptParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by InscriptParser#arith.
	VisitArith(ctx *ArithContext) interface{}

	// Visit a parse tree produced by InscriptParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by InscriptParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by InscriptParser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by InscriptParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by InscriptParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by InscriptParser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by InscriptParser#tableLiteral.
	VisitTableLiteral(ctx *TableLiteralContext) interface{}

	// Visit a parse tree produced by InscriptParser#fieldListOpt.
	VisitFieldListOpt(ctx *FieldListOptContext) interface{}

	// Visit a parse tree produced by InscriptParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by InscriptParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by InscriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
