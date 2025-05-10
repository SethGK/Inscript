// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Inscript
import "github.com/antlr4-go/antlr/v4"

// InscriptListener is a complete listener for a parse tree produced by InscriptParser.
type InscriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterCompoundStmt is called when entering the compoundStmt production.
	EnterCompoundStmt(c *CompoundStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterElseifListOpt is called when entering the elseifListOpt production.
	EnterElseifListOpt(c *ElseifListOptContext)

	// EnterElseifList is called when entering the elseifList production.
	EnterElseifList(c *ElseifListContext)

	// EnterElseif is called when entering the elseif production.
	EnterElseif(c *ElseifContext)

	// EnterElseBlockOpt is called when entering the elseBlockOpt production.
	EnterElseBlockOpt(c *ElseBlockOptContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterFunctionDef is called when entering the functionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementListOpt is called when entering the statementListOpt production.
	EnterStatementListOpt(c *StatementListOptContext)

	// EnterExpressionOpt is called when entering the expressionOpt production.
	EnterExpressionOpt(c *ExpressionOptContext)

	// EnterExpressionListOpt is called when entering the expressionListOpt production.
	EnterExpressionListOpt(c *ExpressionListOptContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterParamListOpt is called when entering the paramListOpt production.
	EnterParamListOpt(c *ParamListOptContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogicalOr is called when entering the logicalOr production.
	EnterLogicalOr(c *LogicalOrContext)

	// EnterLogicalAnd is called when entering the logicalAnd production.
	EnterLogicalAnd(c *LogicalAndContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterArith is called when entering the arith production.
	EnterArith(c *ArithContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterTableLiteral is called when entering the tableLiteral production.
	EnterTableLiteral(c *TableLiteralContext)

	// EnterFieldListOpt is called when entering the fieldListOpt production.
	EnterFieldListOpt(c *FieldListOptContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitCompoundStmt is called when exiting the compoundStmt production.
	ExitCompoundStmt(c *CompoundStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitElseifListOpt is called when exiting the elseifListOpt production.
	ExitElseifListOpt(c *ElseifListOptContext)

	// ExitElseifList is called when exiting the elseifList production.
	ExitElseifList(c *ElseifListContext)

	// ExitElseif is called when exiting the elseif production.
	ExitElseif(c *ElseifContext)

	// ExitElseBlockOpt is called when exiting the elseBlockOpt production.
	ExitElseBlockOpt(c *ElseBlockOptContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitFunctionDef is called when exiting the functionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementListOpt is called when exiting the statementListOpt production.
	ExitStatementListOpt(c *StatementListOptContext)

	// ExitExpressionOpt is called when exiting the expressionOpt production.
	ExitExpressionOpt(c *ExpressionOptContext)

	// ExitExpressionListOpt is called when exiting the expressionListOpt production.
	ExitExpressionListOpt(c *ExpressionListOptContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitParamListOpt is called when exiting the paramListOpt production.
	ExitParamListOpt(c *ParamListOptContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogicalOr is called when exiting the logicalOr production.
	ExitLogicalOr(c *LogicalOrContext)

	// ExitLogicalAnd is called when exiting the logicalAnd production.
	ExitLogicalAnd(c *LogicalAndContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitArith is called when exiting the arith production.
	ExitArith(c *ArithContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitTableLiteral is called when exiting the tableLiteral production.
	ExitTableLiteral(c *TableLiteralContext)

	// ExitFieldListOpt is called when exiting the fieldListOpt production.
	ExitFieldListOpt(c *FieldListOptContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
