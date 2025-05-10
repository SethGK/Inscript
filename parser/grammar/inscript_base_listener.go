// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Inscript
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

// EnterStatementList is called when production statementList is entered.
func (s *BaseInscriptListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseInscriptListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseInscriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseInscriptListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseInscriptListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseInscriptListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseInscriptListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseInscriptListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseInscriptListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseInscriptListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseInscriptListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseInscriptListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseInscriptListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseInscriptListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterCompoundStmt is called when production compoundStmt is entered.
func (s *BaseInscriptListener) EnterCompoundStmt(ctx *CompoundStmtContext) {}

// ExitCompoundStmt is called when production compoundStmt is exited.
func (s *BaseInscriptListener) ExitCompoundStmt(ctx *CompoundStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseInscriptListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseInscriptListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterElseifListOpt is called when production elseifListOpt is entered.
func (s *BaseInscriptListener) EnterElseifListOpt(ctx *ElseifListOptContext) {}

// ExitElseifListOpt is called when production elseifListOpt is exited.
func (s *BaseInscriptListener) ExitElseifListOpt(ctx *ElseifListOptContext) {}

// EnterElseifList is called when production elseifList is entered.
func (s *BaseInscriptListener) EnterElseifList(ctx *ElseifListContext) {}

// ExitElseifList is called when production elseifList is exited.
func (s *BaseInscriptListener) ExitElseifList(ctx *ElseifListContext) {}

// EnterElseif is called when production elseif is entered.
func (s *BaseInscriptListener) EnterElseif(ctx *ElseifContext) {}

// ExitElseif is called when production elseif is exited.
func (s *BaseInscriptListener) ExitElseif(ctx *ElseifContext) {}

// EnterElseBlockOpt is called when production elseBlockOpt is entered.
func (s *BaseInscriptListener) EnterElseBlockOpt(ctx *ElseBlockOptContext) {}

// ExitElseBlockOpt is called when production elseBlockOpt is exited.
func (s *BaseInscriptListener) ExitElseBlockOpt(ctx *ElseBlockOptContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseInscriptListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseInscriptListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseInscriptListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseInscriptListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterFunctionDef is called when production functionDef is entered.
func (s *BaseInscriptListener) EnterFunctionDef(ctx *FunctionDefContext) {}

// ExitFunctionDef is called when production functionDef is exited.
func (s *BaseInscriptListener) ExitFunctionDef(ctx *FunctionDefContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseInscriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseInscriptListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementListOpt is called when production statementListOpt is entered.
func (s *BaseInscriptListener) EnterStatementListOpt(ctx *StatementListOptContext) {}

// ExitStatementListOpt is called when production statementListOpt is exited.
func (s *BaseInscriptListener) ExitStatementListOpt(ctx *StatementListOptContext) {}

// EnterExpressionOpt is called when production expressionOpt is entered.
func (s *BaseInscriptListener) EnterExpressionOpt(ctx *ExpressionOptContext) {}

// ExitExpressionOpt is called when production expressionOpt is exited.
func (s *BaseInscriptListener) ExitExpressionOpt(ctx *ExpressionOptContext) {}

// EnterExpressionListOpt is called when production expressionListOpt is entered.
func (s *BaseInscriptListener) EnterExpressionListOpt(ctx *ExpressionListOptContext) {}

// ExitExpressionListOpt is called when production expressionListOpt is exited.
func (s *BaseInscriptListener) ExitExpressionListOpt(ctx *ExpressionListOptContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseInscriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseInscriptListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterParamListOpt is called when production paramListOpt is entered.
func (s *BaseInscriptListener) EnterParamListOpt(ctx *ParamListOptContext) {}

// ExitParamListOpt is called when production paramListOpt is exited.
func (s *BaseInscriptListener) ExitParamListOpt(ctx *ParamListOptContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseInscriptListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseInscriptListener) ExitParamList(ctx *ParamListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInscriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInscriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalOr is called when production logicalOr is entered.
func (s *BaseInscriptListener) EnterLogicalOr(ctx *LogicalOrContext) {}

// ExitLogicalOr is called when production logicalOr is exited.
func (s *BaseInscriptListener) ExitLogicalOr(ctx *LogicalOrContext) {}

// EnterLogicalAnd is called when production logicalAnd is entered.
func (s *BaseInscriptListener) EnterLogicalAnd(ctx *LogicalAndContext) {}

// ExitLogicalAnd is called when production logicalAnd is exited.
func (s *BaseInscriptListener) ExitLogicalAnd(ctx *LogicalAndContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseInscriptListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseInscriptListener) ExitComparison(ctx *ComparisonContext) {}

// EnterArith is called when production arith is entered.
func (s *BaseInscriptListener) EnterArith(ctx *ArithContext) {}

// ExitArith is called when production arith is exited.
func (s *BaseInscriptListener) ExitArith(ctx *ArithContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseInscriptListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseInscriptListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseInscriptListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseInscriptListener) ExitFactor(ctx *FactorContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseInscriptListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseInscriptListener) ExitUnary(ctx *UnaryContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseInscriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseInscriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseInscriptListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseInscriptListener) ExitAtom(ctx *AtomContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseInscriptListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseInscriptListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterTableLiteral is called when production tableLiteral is entered.
func (s *BaseInscriptListener) EnterTableLiteral(ctx *TableLiteralContext) {}

// ExitTableLiteral is called when production tableLiteral is exited.
func (s *BaseInscriptListener) ExitTableLiteral(ctx *TableLiteralContext) {}

// EnterFieldListOpt is called when production fieldListOpt is entered.
func (s *BaseInscriptListener) EnterFieldListOpt(ctx *FieldListOptContext) {}

// ExitFieldListOpt is called when production fieldListOpt is exited.
func (s *BaseInscriptListener) ExitFieldListOpt(ctx *FieldListOptContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseInscriptListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseInscriptListener) ExitFieldList(ctx *FieldListContext) {}

// EnterField is called when production field is entered.
func (s *BaseInscriptListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseInscriptListener) ExitField(ctx *FieldContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseInscriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseInscriptListener) ExitLiteral(ctx *LiteralContext) {}
