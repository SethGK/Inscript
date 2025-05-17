// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Inscript

import "github.com/antlr4-go/antlr/v4"

// BaseInscriptListener is a complete listener for a parse tree produced by InscriptParser.
type BaseInscriptListener struct{}

var _ InscriptListener = &BaseInscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseInscriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseInscriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseInscriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseInscriptListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseInscriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseInscriptListener) ExitBlock(ctx *BlockContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseInscriptListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseInscriptListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseInscriptListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseInscriptListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseInscriptListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseInscriptListener) ExitTarget(ctx *TargetContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseInscriptListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseInscriptListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseInscriptListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseInscriptListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseInscriptListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseInscriptListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseInscriptListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseInscriptListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseInscriptListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseInscriptListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseInscriptListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseInscriptListener) ExitParam(ctx *ParamContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseInscriptListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseInscriptListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseInscriptListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseInscriptListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseInscriptListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseInscriptListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseInscriptListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseInscriptListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterImportStmt is called when production importStmt is entered.
func (s *BaseInscriptListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BaseInscriptListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseInscriptListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseInscriptListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterGeExpr is called when production geExpr is entered.
func (s *BaseInscriptListener) EnterGeExpr(ctx *GeExprContext) {}

// ExitGeExpr is called when production geExpr is exited.
func (s *BaseInscriptListener) ExitGeExpr(ctx *GeExprContext) {}

// EnterModExpr is called when production modExpr is entered.
func (s *BaseInscriptListener) EnterModExpr(ctx *ModExprContext) {}

// ExitModExpr is called when production modExpr is exited.
func (s *BaseInscriptListener) ExitModExpr(ctx *ModExprContext) {}

// EnterGtExpr is called when production gtExpr is entered.
func (s *BaseInscriptListener) EnterGtExpr(ctx *GtExprContext) {}

// ExitGtExpr is called when production gtExpr is exited.
func (s *BaseInscriptListener) ExitGtExpr(ctx *GtExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseInscriptListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseInscriptListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterIdivExpr is called when production idivExpr is entered.
func (s *BaseInscriptListener) EnterIdivExpr(ctx *IdivExprContext) {}

// ExitIdivExpr is called when production idivExpr is exited.
func (s *BaseInscriptListener) ExitIdivExpr(ctx *IdivExprContext) {}

// EnterSubExpr is called when production subExpr is entered.
func (s *BaseInscriptListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production subExpr is exited.
func (s *BaseInscriptListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterNeqExpr is called when production neqExpr is entered.
func (s *BaseInscriptListener) EnterNeqExpr(ctx *NeqExprContext) {}

// ExitNeqExpr is called when production neqExpr is exited.
func (s *BaseInscriptListener) ExitNeqExpr(ctx *NeqExprContext) {}

// EnterLtExpr is called when production ltExpr is entered.
func (s *BaseInscriptListener) EnterLtExpr(ctx *LtExprContext) {}

// ExitLtExpr is called when production ltExpr is exited.
func (s *BaseInscriptListener) ExitLtExpr(ctx *LtExprContext) {}

// EnterEqExpr is called when production eqExpr is entered.
func (s *BaseInscriptListener) EnterEqExpr(ctx *EqExprContext) {}

// ExitEqExpr is called when production eqExpr is exited.
func (s *BaseInscriptListener) ExitEqExpr(ctx *EqExprContext) {}

// EnterExpExpr is called when production expExpr is entered.
func (s *BaseInscriptListener) EnterExpExpr(ctx *ExpExprContext) {}

// ExitExpExpr is called when production expExpr is exited.
func (s *BaseInscriptListener) ExitExpExpr(ctx *ExpExprContext) {}

// EnterBitandExpr is called when production bitandExpr is entered.
func (s *BaseInscriptListener) EnterBitandExpr(ctx *BitandExprContext) {}

// ExitBitandExpr is called when production bitandExpr is exited.
func (s *BaseInscriptListener) ExitBitandExpr(ctx *BitandExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseInscriptListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseInscriptListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterLeExpr is called when production leExpr is entered.
func (s *BaseInscriptListener) EnterLeExpr(ctx *LeExprContext) {}

// ExitLeExpr is called when production leExpr is exited.
func (s *BaseInscriptListener) ExitLeExpr(ctx *LeExprContext) {}

// EnterBitorExpr is called when production bitorExpr is entered.
func (s *BaseInscriptListener) EnterBitorExpr(ctx *BitorExprContext) {}

// ExitBitorExpr is called when production bitorExpr is exited.
func (s *BaseInscriptListener) ExitBitorExpr(ctx *BitorExprContext) {}

// EnterBitxorExpr is called when production bitxorExpr is entered.
func (s *BaseInscriptListener) EnterBitxorExpr(ctx *BitxorExprContext) {}

// ExitBitxorExpr is called when production bitxorExpr is exited.
func (s *BaseInscriptListener) ExitBitxorExpr(ctx *BitxorExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseInscriptListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseInscriptListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterDivExpr is called when production divExpr is entered.
func (s *BaseInscriptListener) EnterDivExpr(ctx *DivExprContext) {}

// ExitDivExpr is called when production divExpr is exited.
func (s *BaseInscriptListener) ExitDivExpr(ctx *DivExprContext) {}

// EnterShlExpr is called when production shlExpr is entered.
func (s *BaseInscriptListener) EnterShlExpr(ctx *ShlExprContext) {}

// ExitShlExpr is called when production shlExpr is exited.
func (s *BaseInscriptListener) ExitShlExpr(ctx *ShlExprContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseInscriptListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseInscriptListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterShrExpr is called when production shrExpr is entered.
func (s *BaseInscriptListener) EnterShrExpr(ctx *ShrExprContext) {}

// ExitShrExpr is called when production shrExpr is exited.
func (s *BaseInscriptListener) ExitShrExpr(ctx *ShrExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseInscriptListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseInscriptListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseInscriptListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseInscriptListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterBitnotExpr is called when production bitnotExpr is entered.
func (s *BaseInscriptListener) EnterBitnotExpr(ctx *BitnotExprContext) {}

// ExitBitnotExpr is called when production bitnotExpr is exited.
func (s *BaseInscriptListener) ExitBitnotExpr(ctx *BitnotExprContext) {}

// EnterNegExpr is called when production negExpr is entered.
func (s *BaseInscriptListener) EnterNegExpr(ctx *NegExprContext) {}

// ExitNegExpr is called when production negExpr is exited.
func (s *BaseInscriptListener) ExitNegExpr(ctx *NegExprContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseInscriptListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseInscriptListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPrimaryPostfix is called when production primaryPostfix is entered.
func (s *BaseInscriptListener) EnterPrimaryPostfix(ctx *PrimaryPostfixContext) {}

// ExitPrimaryPostfix is called when production primaryPostfix is exited.
func (s *BaseInscriptListener) ExitPrimaryPostfix(ctx *PrimaryPostfixContext) {}

// EnterIndexPostfix is called when production indexPostfix is entered.
func (s *BaseInscriptListener) EnterIndexPostfix(ctx *IndexPostfixContext) {}

// ExitIndexPostfix is called when production indexPostfix is exited.
func (s *BaseInscriptListener) ExitIndexPostfix(ctx *IndexPostfixContext) {}

// EnterAttrPostfix is called when production attrPostfix is entered.
func (s *BaseInscriptListener) EnterAttrPostfix(ctx *AttrPostfixContext) {}

// ExitAttrPostfix is called when production attrPostfix is exited.
func (s *BaseInscriptListener) ExitAttrPostfix(ctx *AttrPostfixContext) {}

// EnterCallPostfix is called when production callPostfix is entered.
func (s *BaseInscriptListener) EnterCallPostfix(ctx *CallPostfixContext) {}

// ExitCallPostfix is called when production callPostfix is exited.
func (s *BaseInscriptListener) ExitCallPostfix(ctx *CallPostfixContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseInscriptListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseInscriptListener) ExitArgList(ctx *ArgListContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseInscriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseInscriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseInscriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseInscriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseInscriptListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseInscriptListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterTableLiteral is called when production tableLiteral is entered.
func (s *BaseInscriptListener) EnterTableLiteral(ctx *TableLiteralContext) {}

// ExitTableLiteral is called when production tableLiteral is exited.
func (s *BaseInscriptListener) ExitTableLiteral(ctx *TableLiteralContext) {}

// EnterTableKeyValue is called when production tableKeyValue is entered.
func (s *BaseInscriptListener) EnterTableKeyValue(ctx *TableKeyValueContext) {}

// ExitTableKeyValue is called when production tableKeyValue is exited.
func (s *BaseInscriptListener) ExitTableKeyValue(ctx *TableKeyValueContext) {}

// EnterTableKey is called when production tableKey is entered.
func (s *BaseInscriptListener) EnterTableKey(ctx *TableKeyContext) {}

// ExitTableKey is called when production tableKey is exited.
func (s *BaseInscriptListener) ExitTableKey(ctx *TableKeyContext) {}
