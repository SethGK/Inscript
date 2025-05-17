// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Inscript

import "github.com/antlr4-go/antlr/v4"

// InscriptListener is a complete listener for a parse tree produced by InscriptParser.
type InscriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterImportStmt is called when entering the importStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterGeExpr is called when entering the geExpr production.
	EnterGeExpr(c *GeExprContext)

	// EnterModExpr is called when entering the modExpr production.
	EnterModExpr(c *ModExprContext)

	// EnterGtExpr is called when entering the gtExpr production.
	EnterGtExpr(c *GtExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterIdivExpr is called when entering the idivExpr production.
	EnterIdivExpr(c *IdivExprContext)

	// EnterSubExpr is called when entering the subExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterNeqExpr is called when entering the neqExpr production.
	EnterNeqExpr(c *NeqExprContext)

	// EnterLtExpr is called when entering the ltExpr production.
	EnterLtExpr(c *LtExprContext)

	// EnterEqExpr is called when entering the eqExpr production.
	EnterEqExpr(c *EqExprContext)

	// EnterExpExpr is called when entering the expExpr production.
	EnterExpExpr(c *ExpExprContext)

	// EnterBitandExpr is called when entering the bitandExpr production.
	EnterBitandExpr(c *BitandExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterLeExpr is called when entering the leExpr production.
	EnterLeExpr(c *LeExprContext)

	// EnterBitorExpr is called when entering the bitorExpr production.
	EnterBitorExpr(c *BitorExprContext)

	// EnterBitxorExpr is called when entering the bitxorExpr production.
	EnterBitxorExpr(c *BitxorExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterDivExpr is called when entering the divExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterShlExpr is called when entering the shlExpr production.
	EnterShlExpr(c *ShlExprContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterShrExpr is called when entering the shrExpr production.
	EnterShrExpr(c *ShrExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterBitnotExpr is called when entering the bitnotExpr production.
	EnterBitnotExpr(c *BitnotExprContext)

	// EnterNegExpr is called when entering the negExpr production.
	EnterNegExpr(c *NegExprContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPrimaryPostfix is called when entering the primaryPostfix production.
	EnterPrimaryPostfix(c *PrimaryPostfixContext)

	// EnterIndexPostfix is called when entering the indexPostfix production.
	EnterIndexPostfix(c *IndexPostfixContext)

	// EnterAttrPostfix is called when entering the attrPostfix production.
	EnterAttrPostfix(c *AttrPostfixContext)

	// EnterCallPostfix is called when entering the callPostfix production.
	EnterCallPostfix(c *CallPostfixContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterTableLiteral is called when entering the tableLiteral production.
	EnterTableLiteral(c *TableLiteralContext)

	// EnterTableKeyValue is called when entering the tableKeyValue production.
	EnterTableKeyValue(c *TableKeyValueContext)

	// EnterTableKey is called when entering the tableKey production.
	EnterTableKey(c *TableKeyContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitGeExpr is called when exiting the geExpr production.
	ExitGeExpr(c *GeExprContext)

	// ExitModExpr is called when exiting the modExpr production.
	ExitModExpr(c *ModExprContext)

	// ExitGtExpr is called when exiting the gtExpr production.
	ExitGtExpr(c *GtExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitIdivExpr is called when exiting the idivExpr production.
	ExitIdivExpr(c *IdivExprContext)

	// ExitSubExpr is called when exiting the subExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitNeqExpr is called when exiting the neqExpr production.
	ExitNeqExpr(c *NeqExprContext)

	// ExitLtExpr is called when exiting the ltExpr production.
	ExitLtExpr(c *LtExprContext)

	// ExitEqExpr is called when exiting the eqExpr production.
	ExitEqExpr(c *EqExprContext)

	// ExitExpExpr is called when exiting the expExpr production.
	ExitExpExpr(c *ExpExprContext)

	// ExitBitandExpr is called when exiting the bitandExpr production.
	ExitBitandExpr(c *BitandExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitLeExpr is called when exiting the leExpr production.
	ExitLeExpr(c *LeExprContext)

	// ExitBitorExpr is called when exiting the bitorExpr production.
	ExitBitorExpr(c *BitorExprContext)

	// ExitBitxorExpr is called when exiting the bitxorExpr production.
	ExitBitxorExpr(c *BitxorExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitDivExpr is called when exiting the divExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitShlExpr is called when exiting the shlExpr production.
	ExitShlExpr(c *ShlExprContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitShrExpr is called when exiting the shrExpr production.
	ExitShrExpr(c *ShrExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitBitnotExpr is called when exiting the bitnotExpr production.
	ExitBitnotExpr(c *BitnotExprContext)

	// ExitNegExpr is called when exiting the negExpr production.
	ExitNegExpr(c *NegExprContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPrimaryPostfix is called when exiting the primaryPostfix production.
	ExitPrimaryPostfix(c *PrimaryPostfixContext)

	// ExitIndexPostfix is called when exiting the indexPostfix production.
	ExitIndexPostfix(c *IndexPostfixContext)

	// ExitAttrPostfix is called when exiting the attrPostfix production.
	ExitAttrPostfix(c *AttrPostfixContext)

	// ExitCallPostfix is called when exiting the callPostfix production.
	ExitCallPostfix(c *CallPostfixContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitTableLiteral is called when exiting the tableLiteral production.
	ExitTableLiteral(c *TableLiteralContext)

	// ExitTableKeyValue is called when exiting the tableKeyValue production.
	ExitTableKeyValue(c *TableKeyValueContext)

	// ExitTableKey is called when exiting the tableKey production.
	ExitTableKey(c *TableKeyContext)
}
