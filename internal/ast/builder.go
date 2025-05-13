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
	fmt.Println("DEBUG: Visiting Program") // Debug print
	prog := &Program{PosToken: token.Pos(ctx.GetStart().GetStart())}
	// The grammar ensures StatementList is present if there are statements
	if ctx.StatementList() != nil {
		fmt.Println("DEBUG: Visiting StatementList from VisitProgram") // Debug print
		// VisitStatementList returns []Statement, which we append to the program's Stmts
		visitedStmts := ctx.StatementList().Accept(v)                                                       // Get the result first
		fmt.Printf("DEBUG: VisitStatementList returned type: %T, value: %+v\n", visitedStmts, visitedStmts) // Debug print
		if stmts, ok := visitedStmts.([]Statement); ok {
			prog.Stmts = append(prog.Stmts, stmts...)
			fmt.Printf("DEBUG: Appended %d statements to Program\n", len(stmts)) // Debug print
		} else {
			fmt.Printf("WARNING: VisitStatementList did not return []Statement, got %T\n", visitedStmts) // Debug print
		}
	}
	fmt.Printf("DEBUG: Finished Visiting Program, returning %+v (Stmts count: %d)\n", prog, len(prog.Stmts)) // Debug print
	return prog
}

// VisitStatementList handles one or more statements
func (v *ASTBuilder) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	fmt.Println("DEBUG: Visiting StatementList") // Debug print
	var stmts []Statement
	// Iterate through all statement contexts and visit them
	for i, sctx := range ctx.AllStatement() {
		fmt.Printf("DEBUG: Visiting statement context #%d (type: %T) from VisitStatementList\n", i, sctx) // Debug print
		// Each statement visit should return a Statement node
		stmt := sctx.Accept(v)                                                                    // Call Accept on the StatementContext
		fmt.Printf("DEBUG: Statement context #%d returned type: %T, value: %+v\n", i, stmt, stmt) // Debug print
		if stmtNode, ok := stmt.(Statement); ok {
			stmts = append(stmts, stmtNode)
			fmt.Printf("DEBUG: Appended statement #%d (type: %T) to StatementList\n", i, stmtNode) // Debug print
		} else {
			fmt.Printf("  WARNING: Visited statement context type %T did not return a Statement interface, got %T\n", sctx, stmt) // Debug print
		}
	}
	fmt.Printf("DEBUG: Finished Visiting StatementList, returning %d statements\n", len(stmts)) // Debug print
	return stmts                                                                                // Returns []Statement
}

// VisitStatement handles a single statement (either simple or compound).
// This method is crucial for ensuring the result from sub-visits is propagated.
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("DEBUG: Visiting Statement") // Debug print
	// A statement context has only one child which is either a simpleStmt or compoundStmt
	if ctx.SimpleStmt() != nil {
		fmt.Println("DEBUG: Visiting SimpleStmt from VisitStatement") // Debug print
		result := ctx.SimpleStmt().Accept(v)
		fmt.Printf("DEBUG: SimpleStmt.Accept(v) returned type: %T, value: %+v\n", result, result)   // Debug print
		fmt.Printf("DEBUG: Finished Visiting Statement (SimpleStmt), returning type: %T\n", result) // Debug print
		return result                                                                               // Return the result from visiting the simple statement
	}
	if ctx.CompoundStmt() != nil {
		fmt.Println("DEBUG: Visiting CompoundStmt from VisitStatement") // Debug print
		result := ctx.CompoundStmt().Accept(v)
		fmt.Printf("DEBUG: CompoundStmt.Accept(v) returned type: %T, value: %+v\n", result, result)   // Debug print
		fmt.Printf("DEBUG: Finished Visiting Statement (CompoundStmt), returning type: %T\n", result) // Debug print
		return result                                                                                 // Return the result from visiting the compound statement
	}
	fmt.Printf("WARNING: Visiting Statement, unexpected context type: %T\n", ctx) // Debug print
	return nil                                                                    // Should not happen if grammar is correct
}

// VisitSimpleStmt handles a simple statement (assignment, exprStmt, printStmt, returnStmt).
// This method is crucial for ensuring the result from sub-visits is propagated.
func (v *ASTBuilder) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting SimpleStmt") // Debug print
	// A simpleStmt context has only one child which is one of the simple statement types
	if ctx.Assignment() != nil {
		fmt.Println("DEBUG: Visiting Assignment from VisitSimpleStmt") // Debug print
		result := ctx.Assignment().Accept(v)
		fmt.Printf("DEBUG: Assignment.Accept(v) returned type: %T, value: %+v\n", result, result)    // Debug print
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (Assignment), returning type: %T\n", result) // Debug print
		return result                                                                                // Return the result from visiting the assignment
	}
	if ctx.ExprStmt() != nil {
		fmt.Println("DEBUG: Visiting ExprStmt from VisitSimpleStmt") // Debug print
		result := ctx.ExprStmt().Accept(v)
		fmt.Printf("DEBUG: ExprStmt.Accept(v) returned type: %T, value: %+v\n", result, result)    // Debug print
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (ExprStmt), returning type: %T\n", result) // Debug print
		return result                                                                              // Return the result from visiting the expression statement
	}
	if ctx.PrintStmt() != nil {
		fmt.Println("DEBUG: Visiting PrintStmt from VisitSimpleStmt") // Debug print
		result := ctx.PrintStmt().Accept(v)
		fmt.Printf("DEBUG: PrintStmt.Accept(v) returned type: %T, value: %+v\n", result, result)    // Debug print
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (PrintStmt), returning type: %T\n", result) // Debug print
		return result                                                                               // Return the result from visiting the print statement
	}
	if ctx.ReturnStmt() != nil {
		fmt.Println("DEBUG: Visiting ReturnStmt from VisitSimpleStmt") // Debug print
		result := ctx.ReturnStmt().Accept(v)
		fmt.Printf("DEBUG: ReturnStmt.Accept(v) returned type: %T, value: %+v\n", result, result)    // Debug print
		fmt.Printf("DEBUG: Finished Visiting SimpleStmt (ReturnStmt), returning type: %T\n", result) // Debug print
		return result                                                                                // Return the result from visiting the return statement
	}
	fmt.Printf("WARNING: Visiting SimpleStmt, unexpected context type: %T\n", ctx) // Debug print
	return nil                                                                     // Should not happen if grammar is correct
}

// VisitCompoundStmt handles compound statements (if, while, for, functionDef).
// Note: If functionDef is only used as an expression (in atom), this case might be removed.
func (v *ASTBuilder) VisitCompoundStmt(ctx *parser.CompoundStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting CompoundStmt") // Debug print
	// A compoundStmt context has only one child which is one of the compound statement types
	if ctx.IfStmt() != nil {
		fmt.Println("DEBUG: Visiting IfStmt from VisitCompoundStmt") // Debug print
		result := ctx.IfStmt().Accept(v)
		fmt.Printf("DEBUG: IfStmt.Accept(v) returned type: %T, value: %+v\n", result, result)      // Debug print
		fmt.Printf("DEBUG: Finished Visiting CompoundStmt (IfStmt), returning type: %T\n", result) // Debug print
		return result                                                                              // Return the result from visiting the if statement
	}
	if ctx.WhileStmt() != nil {
		fmt.Println("DEBUG: Visiting WhileStmt from VisitCompoundStmt") // Debug print
		result := ctx.WhileStmt().Accept(v)
		fmt.Printf("DEBUG: WhileStmt.Accept(v) returned type: %T, value: %+v\n", result, result)      // Debug print
		fmt.Printf("DEBUG: Finished Visiting CompoundStmt (WhileStmt), returning type: %T\n", result) // Debug print
		return result                                                                                 // Return the result from visiting the while statement
	}
	if ctx.ForStmt() != nil {
		fmt.Println("DEBUG: Visiting ForStmt from VisitCompoundStmt") // Debug print
		result := ctx.ForStmt().Accept(v)
		fmt.Printf("DEBUG: ForStmt.Accept(v) returned type: %T, value: %+v\n", result, result)      // Debug print
		fmt.Printf("DEBUG: Finished Visiting CompoundStmt (ForStmt), returning type: %T\n", result) // Debug print
		return result                                                                               // Return the result from visiting the for statement
	}
	// Removed the case for ctx.FunctionDef() here, assuming functionDef is only an expression now.
	// If you still need standalone named function declarations, this logic needs adjustment.

	fmt.Printf("WARNING: Visiting CompoundStmt, unexpected context type: %T\n", ctx) // Debug print
	return nil                                                                       // Should not happen if grammar is correct and functionDef is removed from compoundStmt
}

// --- Simple statements ---

// VisitAssignment builds an AssignStmt node.
func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	// fmt.Println("DEBUG: Visiting Assignment") // Less verbose debug print
	// Target must be a Primary expression (e.g., identifier, index, call)
	target := ctx.Primary().Accept(v).(Expression)
	// Value is any Expression
	value := ctx.Expression().Accept(v).(Expression)
	return &AssignStmt{Target: target, Value: value, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitExprStmt builds an ExprStmt node.
func (v *ASTBuilder) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	// fmt.Println("DEBUG: Visiting ExprStmt") // Less verbose debug print
	// The statement is just an expression
	expr := ctx.Expression().Accept(v).(Expression)
	return &ExprStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitPrintStmt builds a PrintStmt node.
func (v *ASTBuilder) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting PrintStmt") // Debug print
	var exprs []Expression
	// Check if ExpressionListOpt has an actual ExpressionList
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		// Visit the expression list to get the slice of expressions
		visitedExprs := el.Accept(v)                                                                     // Get the result first
		fmt.Printf("DEBUG: Visited ExpressionListOpt.ExpressionList, returned type: %T\n", visitedExprs) // Debug print
		if visitedExprsSlice, ok := visitedExprs.([]Expression); ok {
			exprs = visitedExprsSlice
			fmt.Printf("DEBUG: Collected %d expressions for PrintStmt\n", len(exprs)) // Debug print
		} else {
			fmt.Printf("  WARNING: Visiting ExpressionListOpt.ExpressionList did not return []Expression, got %T\n", visitedExprs) // Debug print
		}
	}
	printStmtNode := &PrintStmt{Exprs: exprs, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("DEBUG: Finished Visiting PrintStmt, returning %+v\n", printStmtNode) // Debug print
	return printStmtNode                                                             // Ensure *ast.PrintStmt is returned
}

// VisitReturnStmt builds a ReturnStmt node.
func (v *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// fmt.Println("DEBUG: Visiting ReturnStmt") // Less verbose debug print
	var expr Expression
	// Check if ExpressionOpt has an actual Expression
	if ctx.ExpressionOpt().Expression() != nil {
		// Visit the optional expression
		expr = ctx.ExpressionOpt().Expression().Accept(v).(Expression)
	}
	return &ReturnStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// --- Compound statements --- (continued)

// VisitIfStmt builds an IfStmt node.
func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting IfStmt") // Debug print
	// Visit the main condition and the 'then' block
	cond := ctx.Expression().Accept(v).(Expression)
	then := ctx.Block().Accept(v).(*BlockStmt)

	var elseIfs []ElseIf
	// Check for optional elseif clauses
	if elOpt := ctx.ElseifListOpt().ElseifList(); elOpt != nil {
		// Iterate through all elseif contexts
		for _, eis := range elOpt.AllElseif() {
			// Visit the elseif condition and block
			ec := eis.Expression().Accept(v).(Expression)
			body := eis.Block().Accept(v).(*BlockStmt)
			// Append a new ElseIf struct to the slice
			elseIfs = append(elseIfs, ElseIf{Cond: ec, Body: body, PosToken: token.Pos(eis.GetStart().GetStart())})
		}
	}

	var elseBlk *BlockStmt
	// Check for an optional else block
	if ctx.ElseBlockOpt().Block() != nil {
		// Visit the else block
		elseBlk = ctx.ElseBlockOpt().Block().Accept(v).(*BlockStmt)
	}

	ifStmtNode := &IfStmt{Cond: cond, Then: then, ElseIfs: elseIfs, Else: elseBlk, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("DEBUG: Finished Visiting IfStmt, returning %+v\n", ifStmtNode) // Debug print
	return ifStmtNode
}

// VisitWhileStmt builds a WhileStmt node.
func (v *ASTBuilder) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting WhileStmt") // Debug print
	// Visit the loop condition and body block
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	whileStmtNode := &WhileStmt{Cond: cond, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("DEBUG: Finished Visiting WhileStmt, returning %+v\n", whileStmtNode) // Debug print
	return whileStmtNode
}

// VisitForStmt builds a ForStmt node.
func (v *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	fmt.Println("DEBUG: Visiting ForStmt") // Debug print
	// Get the variable name from the IDENTIFIER token
	varName := ctx.IDENTIFIER().GetText()
	// Visit the iterable expression
	iter := ctx.Expression().Accept(v).(Expression)
	// Visit the loop body block
	body := ctx.Block().Accept(v).(*BlockStmt)
	forStmtNode := &ForStmt{Variable: varName, Iterable: iter, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("DEBUG: Finished Visiting ForStmt, returning %+v\n", forStmtNode) // Debug print
	return forStmtNode
}

// VisitFunctionDef builds a FunctionLiteral node.
// This method is now expected to be called when functionDef appears as an atom/expression.
func (v *ASTBuilder) VisitFunctionDef(ctx *parser.FunctionDefContext) interface{} {
	fmt.Println("DEBUG: Visiting FunctionDef (as Expression)") // Debug print
	var params []string
	// Check for optional parameter list
	if pl := ctx.ParamListOpt().ParamList(); pl != nil {
		// Iterate through all IDENTIFIER tokens in the parameter list
		for _, id := range pl.AllIDENTIFIER() {
			params = append(params, id.GetText())
		}
	}
	// Visit the function body block
	body := ctx.Block().Accept(v).(*BlockStmt)
	// Create and return a FunctionLiteral node (no Name field for literals)
	funcLiteralNode := &FunctionLiteral{Params: params, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("DEBUG: Finished Visiting FunctionDef (as Expression), returning %+v\n", funcLiteralNode) // Debug print
	return funcLiteralNode                                                                               // Return *ast.FunctionLiteral
}

// VisitBlock builds a BlockStmt node.
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	// fmt.Println("DEBUG: Visiting Block") // Less verbose debug print
	var stmts []Statement
	// Check for optional statement list within the block
	if sl := ctx.StatementListOpt().StatementList(); sl != nil {
		// Visit the statement list to get the slice of statements
		visitedStmts := sl.Accept(v) // Get result first
		// fmt.Printf("DEBUG: Visited StatementListOpt.StatementList, returned type: %T\n", visitedStmts) // Less verbose debug print
		if visitedStmtsSlice, ok := visitedStmts.([]Statement); ok {
			stmts = visitedStmtsSlice
		}
	}
	return &BlockStmt{Stmts: stmts, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitExpressionList handles a comma-separated list of expressions.
func (v *ASTBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	// fmt.Println("DEBUG: Visiting ExpressionList") // Less verbose debug print
	var exprs []Expression
	// Iterate through all expression contexts in the list
	for _, e := range ctx.AllExpression() {
		// Visit each expression
		expr := e.Accept(v).(Expression)
		exprs = append(exprs, expr)
	}
	return exprs // Returns []Expression
}

// VisitParamList handles a comma-separated list of identifiers for function parameters.
func (v *ASTBuilder) VisitParamList(ctx *parser.ParamListContext) interface{} {
	// fmt.Println("DEBUG: Visiting ParamList") // Less verbose debug print
	var params []string
	// Iterate through all IDENTIFIER tokens in the parameter list
	for _, id := range ctx.AllIDENTIFIER() {
		params = append(params, id.GetText())
	}
	return params
}

// --- Expressions ---

// VisitExpression simply visits the top-level logicalOr expression.
func (v *ASTBuilder) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	// fmt.Println("DEBUG: Visiting Expression") // Less verbose debug print
	return ctx.LogicalOr().Accept(v)
}

// VisitLogicalOr handles 'or' operations (left-associative).
func (v *ASTBuilder) VisitLogicalOr(ctx *parser.LogicalOrContext) interface{} {
	// fmt.Println("DEBUG: Visiting LogicalOr") // Less verbose debug print
	// Start with the first logicalAnd expression
	expr := ctx.LogicalAnd(0).Accept(v).(Expression)
	// Iterate through children to find 'or' tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is an 'or' token using the generated token constant
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__15 { // Corrected token name based on list
			right := ctx.LogicalAnd((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitLogicalAnd handles 'and' operations (left-associative).
func (v *ASTBuilder) VisitLogicalAnd(ctx *parser.LogicalAndContext) interface{} {
	// fmt.Println("DEBUG: Visiting LogicalAnd") // Less verbose debug print
	// Start with the first comparison expression
	expr := ctx.Comparison(0).Accept(v).(Expression)
	// Iterate through children to find 'and' tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__16 { // Corrected token name based on list
			right := ctx.Comparison((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitComparison handles comparison operations (chainable, but typically left-associative evaluation).
func (v *ASTBuilder) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	// fmt.Println("DEBUG: Visiting Comparison") // Less verbose debug print
	// Start with the first arithmetic expression
	expr := ctx.Arith(0).Accept(v).(Expression)
	// Iterate through children to find comparison operator tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is a comparison operator token using generated token constants
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__17, // '=='
			parser.InscriptParserT__18, // '!='
			parser.InscriptParserT__19, // '<'
			parser.InscriptParserT__20, // '>'
			parser.InscriptParserT__21, // '<='
			parser.InscriptParserT__22: // '>=' // Corrected token names based on list
			right := ctx.Arith((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitArith handles addition and subtraction operations (left-associative).
func (v *ASTBuilder) VisitArith(ctx *parser.ArithContext) interface{} {
	// fmt.Println("DEBUG: Visiting Arith") // Less verbose debug print
	// Start with the first term expression
	expr := ctx.Term(0).Accept(v).(Expression)
	// Iterate through children to find '+' or '-' operator tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is an '+' or '-' token using generated token constants
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__23, // '+'
			parser.InscriptParserT__24: // '-' // Corrected token names based on list
			right := ctx.Term((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitTerm handles multiplication, division, and modulo operations (left-associative).
func (v *ASTBuilder) VisitTerm(ctx *parser.TermContext) interface{} {
	// fmt.Println("DEBUG: Visiting Term") // Less verbose debug print
	// Start with the first factor expression
	expr := ctx.Factor(0).Accept(v).(Expression)
	// Iterate through subsequent '*', '/', or '%' operators and right-hand operands
	// The operators are not labeled, so we access them by child index.
	// Children are ordered: factor op factor op factor ...
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode) // Get the operator token node
		// Check if the child is a '*', '/', or '%' token using generated token constants
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__25, // '*'
			parser.InscriptParserT__26, // '/'
			parser.InscriptParserT__27: // '%' // Corrected token names based on list
			right := ctx.Factor((i + 1) / 2).Accept(v).(Expression) // Get the corresponding right factor
			// Build a BinaryExpr for each term operation
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitFactor handles exponentiation operations (right-associative).
func (v *ASTBuilder) VisitFactor(ctx *parser.FactorContext) interface{} {
	// fmt.Println("DEBUG: Visiting Factor") // Less verbose debug print
	// Start with the first unary expression
	expr := ctx.Unary(0).Accept(v).(Expression)
	// Iterate through subsequent '^' operators and right-hand operands
	// Exponentiation is right-associative, so we build the AST from right to left conceptually,
	// or by chaining from left to right here, ensuring the right side is the result of the next factor.
	// The AST structure will naturally represent the right-associativity if built this way.
	for i := 1; i < ctx.GetChildCount(); i += 2 { // Iterate through children after the first Unary
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is a '^' token using the generated token constant
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__28 { // Corrected token name for '^' based on list
			right := ctx.Unary((i + 1) / 2).Accept(v).(Expression) // Get the corresponding right unary
			// Build a BinaryExpr for each exponentiation operation
			// The position token should ideally be the operator itself, but the grammar doesn't label it.
			// Using the operator's position is better.
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	return expr
}

// VisitUnary handles unary operations ('+', '-', 'not').
func (v *ASTBuilder) VisitUnary(ctx *parser.UnaryContext) interface{} {
	// fmt.Println("DEBUG: Visiting Unary") // Less verbose debug print
	// Check if it's a unary operation (has 2 children: operator and expression)
	if ctx.GetChildCount() == 2 {
		opNode := ctx.GetChild(0).(antlr.TerminalNode) // Get the operator token node
		op := opNode.GetText()                         // Get the operator text
		// Ensure we visit the correct child for the expression (it's the second child)
		// Check the token type of the operator
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__23, // '+'
			parser.InscriptParserT__24, // '-'
			parser.InscriptParserT__29: // 'not' // Corrected token names based on list
			if exprCtx, ok := ctx.GetChild(1).(*parser.UnaryContext); ok { // Unary can be recursive
				expr := exprCtx.Accept(v).(Expression)
				return &UnaryExpr{Operator: op, Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
			}
			// Fallback if the child is not a UnaryContext (e.g., a Primary)
			if primaryCtx, ok := ctx.GetChild(1).(*parser.PrimaryContext); ok {
				expr := primaryCtx.Accept(v).(Expression)
				return &UnaryExpr{Operator: op, Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
			}
			// Handle unexpected child type if necessary
			return nil // Or return an error
		default:
			// If the first child is a terminal node but not a unary operator, this case should not be reached
			// if the grammar is correct.
			return nil // Or return an error
		}
	}
	// If not a unary operation, it must be a primary expression
	return ctx.Primary().Accept(v)
}

// VisitPrimary handles primary expressions, including indexing and function calls.
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	// fmt.Println("DEBUG: Visiting Primary") // Less verbose debug print
	// Start with the base atom
	expr := ctx.Atom().Accept(v).(Expression)

	// Iterate through all children after the atom to find suffixes
	// The children are ordered: atom suffix1 suffix2 ...
	for i := 1; i < ctx.GetChildCount(); i++ {
		suffixChild := ctx.GetChild(i)

		// Check if the suffix is an indexing operation (starts with '[')
		if termNode, ok := suffixChild.(antlr.TerminalNode); ok && termNode.GetSymbol().GetTokenType() == parser.InscriptParserT__30 { // Corrected token name for '[' based on list
			// The next child should be the expression inside the brackets, followed by ']'
			// The structure is primary -> atom '[' expression ']'
			// We need to get the expression child which is at index i+1
			if i+1 < ctx.GetChildCount() {
				// Assert the child is an ExpressionContext before calling Accept
				if indexExprCtx, ok := ctx.GetChild(i + 1).(*parser.ExpressionContext); ok {
					indexExpr := indexExprCtx.Accept(v).(Expression)
					// Wrap the current expression in an IndexExpr node using the correct field name 'Primary'
					expr = &IndexExpr{Primary: expr, Index: indexExpr, PosToken: token.Pos(termNode.GetSymbol().GetStart())} // Use '[' token position
					i += 2                                                                                                   // Skip the expression and the closing ']'
				} else {
					// Handle unexpected child type (should ideally be caught by parser)
					break
				}
			} else {
				// Handle unexpected end of input after '['
				break
			}
		} else if termNode, ok := suffixChild.(antlr.TerminalNode); ok && termNode.GetSymbol().GetTokenType() == parser.InscriptParserT__2 { // Corrected token name for '(' based on list
			// The suffix is a function call (starts with '(')
			// The structure is primary -> atom '(' expressionListOpt ')'
			// The ExpressionListOpt is the next child at index i+1
			if i+1 < ctx.GetChildCount() {
				// Assert the child is an ExpressionListOptContext before calling calling Accept
				if expressionListOptCtx, ok := ctx.GetChild(i + 1).(*parser.ExpressionListOptContext); ok {
					var args []Expression
					// Check if there's an actual expression list
					if el := expressionListOptCtx.ExpressionList(); el != nil {
						// Visit the expression list to get the arguments
						if visitedArgs, ok := el.Accept(v).([]Expression); ok {
							args = visitedArgs
						}
					}
					// Wrap the current expression in a CallExpr node using the correct field name 'Callee'
					expr = &CallExpr{Callee: expr, Args: args, PosToken: token.Pos(termNode.GetSymbol().GetStart())} // Use '(' token position
					i += 2                                                                                           // Skip the expressionListOpt and the closing ')'
				} else {
					// Handle unexpected child type
					break
				}
			} else {
				// Handle unexpected end of input after '('
				break
			}
		}
		// If the child is neither '[' nor '(', it's an unexpected child in primary
		// The parser should ideally catch this.
	}

	return expr
}

// VisitAtom handles the different types of atoms (literals, identifiers, lists, tables, parenthesized expressions, function definitions).
func (v *ASTBuilder) VisitAtom(ctx *parser.AtomContext) interface{} {
	// fmt.Println("DEBUG: Visiting Atom") // Less verbose debug print
	// Check which alternative the atom matches and visit accordingly
	if lit := ctx.Literal(); lit != nil {
		return lit.Accept(v)
	}
	if id := ctx.IDENTIFIER(); id != nil {
		// Build an Identifier node
		return &Identifier{Name: id.GetText(), PosToken: token.Pos(id.GetSymbol().GetStart())}
	}
	if ll := ctx.ListLiteral(); ll != nil {
		return ll.Accept(v)
	}
	if tl := ctx.TableLiteral(); tl != nil {
		return tl.Accept(v)
	}
	// Handle the new case for function definitions as atoms/expressions
	if fd := ctx.FunctionDef(); fd != nil {
		// Visit the function definition context. VisitFunctionDef should now return *ast.FunctionLiteral
		return fd.Accept(v) // VisitFunctionDef builds and returns a FunctionLiteral
	}
	// If none of the above match, it must be a parenthesized expression '(' expression ')'
	// The expression is the child at index 1. We need to assert its type.
	// This check needs to be after the other alternatives.
	if exprCtx := ctx.Expression(); exprCtx != nil { // Handles parenthesized expressions '(' expression ')'
		return exprCtx.Accept(v)
	}

	// If none of the above match, it's an unexpected atom type (should ideally be caught by parser)
	return nil // Or return an error
}

// VisitListLiteral builds a ListLiteral node.
func (v *ASTBuilder) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	// fmt.Println("DEBUG: Visiting ListLiteral") // Less verbose debug print
	var elements []Expression
	// Check for optional expression list
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		// Visit the expression list to get the slice of elements
		if visitedElements, ok := el.Accept(v).([]Expression); ok {
			elements = visitedElements
		}
	}
	return &ListLiteral{Elements: elements, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitTableLiteral builds a TableLiteral node.
func (v *ASTBuilder) VisitTableLiteral(ctx *parser.TableLiteralContext) interface{} {
	// fmt.Println("DEBUG: Visiting TableLiteral") // Less verbose debug print
	var fields []Field
	// Check for optional field list
	if fl := ctx.FieldListOpt().FieldList(); fl != nil {
		// Iterate through all field contexts
		for _, f := range fl.AllField() {
			// Visit each field to get the key/value pair
			// VisitField returns a Field struct, not a pointer
			if fieldStruct, ok := f.Accept(v).(Field); ok { // VisitField should return ast.Field
				fields = append(fields, fieldStruct) // Append the Field struct
			}
		}
	}
	return &TableLiteral{Fields: fields, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitField builds a Field struct for table literals.
// Note: This is called by VisitTableLiteral, it doesn't return an AST node itself,
// but a struct used within the TableLiteral node.
func (v *ASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	// fmt.Println("DEBUG: Visiting Field") // Less verbose debug print
	// Get the key name from the IDENTIFIER token
	keyName := ctx.IDENTIFIER().GetText()
	// Visit the value expression
	value := ctx.Expression().Accept(v).(Expression)
	// Return a Field struct
	return Field{Key: keyName, Value: value, PosToken: token.Pos(ctx.GetStart().GetStart())} // Return Field struct, not pointer
}

// --- Literal nodes ---

// VisitLiteral handles the different types of literals.
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	// fmt.Println("DEBUG: Visiting Literal") // Less verbose debug print
	// Check which literal type matches and build the corresponding AST node
	if i := ctx.INTEGER(); i != nil {
		val, _ := strconv.ParseInt(i.GetText(), 10, 64) // Handle potential error
		return &IntegerLiteral{Value: val, PosToken: token.Pos(i.GetSymbol().GetStart())}
	}
	if f := ctx.FLOAT(); f != nil {
		val, _ := strconv.ParseFloat(f.GetText(), 64) // Handle potential error
		return &FloatLiteral{Value: val, PosToken: token.Pos(f.GetSymbol().GetStart())}
	}
	if s := ctx.STRING(); s != nil {
		// Remove quotes from the string literal
		text := s.GetText()
		// TODO: Handle escape sequences correctly
		value := text[1 : len(text)-1]
		return &StringLiteral{Value: value, PosToken: token.Pos(s.GetSymbol().GetStart())}
	}
	if b := ctx.BOOLEAN(); b != nil {
		val, _ := strconv.ParseBool(b.GetText()) // Handle potential error
		return &BooleanLiteral{Value: val, PosToken: token.Pos(b.GetSymbol().GetStart())}
	}
	if nilToken := ctx.GetToken(parser.InscriptParserT__32, 0); nilToken != nil { // Check for 'nil' token using T__32
		return &NilLiteral{PosToken: token.Pos(nilToken.GetSymbol().GetStart())}
	}
	// If none of the above match, it's an unexpected literal type (should ideally be caught by parser)
	return nil // Or return an error
}

// --- Other visitor methods (optional, based on your grammar and needs) ---
// You might have other Visit methods for optional rules or specific grammar parts.
// If you don't explicitly define a Visit method for a rule, the BaseInscriptVisitor's
// default implementation is used, which typically visits children.

// Example of a default Visit method if needed (usually not necessary unless you need custom traversal)
// func (v *ASTBuilder) VisitTerminal(node antlr.TerminalNode) interface{} {
// 	// fmt.Printf("DEBUG: Visiting Terminal: %s\n", node.GetText())
// 	return node // Return the terminal node itself, or nil, or a specific value
// }

// Example of a default VisitErrorNode method for handling parser errors during traversal
// func (v *ASTBuilder) VisitErrorNode(node antlr.ErrorNode) interface{} {
// 	fmt.Printf("ERROR: Visiting ErrorNode: %s\n", node.GetText())
// 	// Depending on your error handling strategy, you might return an error, nil,
// 	// or a special error AST node here.
// 	return nil
// }

// Ensure ASTBuilder satisfies parser.InscriptVisitor
var _ parser.InscriptVisitor = (*ASTBuilder)(nil)
