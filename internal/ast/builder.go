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
	fmt.Println("Visiting Program") // Debug print
	prog := &Program{PosToken: token.Pos(ctx.GetStart().GetStart())}
	// The grammar ensures StatementList is present if there are statements
	if ctx.StatementList() != nil {
		// VisitStatementList returns []Statement, which we append to the program's Stmts
		visitedStmts := ctx.StatementList().Accept(v)                            // Get the result first
		fmt.Printf("  Visited StatementList, returned type: %T\n", visitedStmts) // Debug print
		if stmts, ok := visitedStmts.([]Statement); ok {
			prog.Stmts = append(prog.Stmts, stmts...)
		} else {
			fmt.Printf("WARNING: VisitStatementList did not return []Statement, got %T\n", visitedStmts) // Debug print
		}
	}
	fmt.Printf("Finished Visiting Program, returning %+v\n", prog) // Debug print
	return prog
}

// VisitStatementList handles one or more statements
func (v *ASTBuilder) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	fmt.Println("Visiting StatementList") // Debug print
	var stmts []Statement
	// Iterate through all statement contexts and visit them
	for _, sctx := range ctx.AllStatement() {
		// Each statement visit should return a Statement node
		stmt := sctx.Accept(v)                                                              // Call Accept on the StatementContext
		fmt.Printf("  Visited statement context type: %T, returned type: %T\n", sctx, stmt) // Debug print
		if stmtNode, ok := stmt.(Statement); ok {
			stmts = append(stmts, stmtNode)
		} else {
			fmt.Printf("  WARNING: Visited statement context type %T did not return a Statement interface, got %T\n", sctx, stmt) // Debug print
		}
	}
	fmt.Printf("Finished Visiting StatementList, returning %d statements\n", len(stmts)) // Debug print
	return stmts                                                                         // Returns []Statement
}

// VisitStatement handles a single statement (either simple or compound).
// This method is crucial for ensuring the result from sub-visits is propagated.
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("Visiting Statement") // Debug print
	// A statement context has only one child which is either a simpleStmt or compoundStmt
	if ctx.SimpleStmt() != nil {
		result := ctx.SimpleStmt().Accept(v)
		fmt.Printf("Finished Visiting Statement (SimpleStmt), returning type: %T\n", result) // Debug print
		return result                                                                        // Return the result from visiting the simple statement
	}
	if ctx.CompoundStmt() != nil {
		result := ctx.CompoundStmt().Accept(v)
		fmt.Printf("Finished Visiting Statement (CompoundStmt), returning type: %T\n", result) // Debug print
		return result                                                                          // Return the result from visiting the compound statement
	}
	fmt.Printf("WARNING: Visiting Statement, unexpected context type: %T\n", ctx) // Debug print
	return nil                                                                    // Should not happen if grammar is correct
}

// VisitSimpleStmt handles a simple statement (assignment, exprStmt, printStmt, returnStmt).
// This method is crucial for ensuring the result from sub-visits is propagated.
func (v *ASTBuilder) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	fmt.Println("Visiting SimpleStmt") // Debug print
	// A simpleStmt context has only one child which is one of the simple statement types
	if ctx.Assignment() != nil {
		result := ctx.Assignment().Accept(v)
		fmt.Printf("Finished Visiting SimpleStmt (Assignment), returning type: %T\n", result) // Debug print
		return result                                                                         // Return the result from visiting the assignment
	}
	if ctx.ExprStmt() != nil {
		result := ctx.ExprStmt().Accept(v)
		fmt.Printf("Finished Visiting SimpleStmt (ExprStmt), returning type: %T\n", result) // Debug print
		return result                                                                       // Return the result from visiting the expression statement
	}
	if ctx.PrintStmt() != nil {
		result := ctx.PrintStmt().Accept(v)
		fmt.Printf("Finished Visiting SimpleStmt (PrintStmt), returning type: %T\n", result) // Debug print
		return result                                                                        // Return the result from visiting the print statement
	}
	if ctx.ReturnStmt() != nil {
		result := ctx.ReturnStmt().Accept(v)
		fmt.Printf("Finished Visiting SimpleStmt (ReturnStmt), returning type: %T\n", result) // Debug print
		return result                                                                         // Return the result from visiting the return statement
	}
	fmt.Printf("WARNING: Visiting SimpleStmt, unexpected context type: %T\n", ctx) // Debug print
	return nil                                                                     // Should not happen if grammar is correct
}

// --- Simple statements ---

// VisitAssignment builds an AssignStmt node.
func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	fmt.Println("Visiting Assignment") // Debug print
	// Target must be a Primary expression (e.g., identifier, index, call)
	target := ctx.Primary().Accept(v).(Expression)
	// Value is any Expression
	value := ctx.Expression().Accept(v).(Expression)
	assignStmt := &AssignStmt{Target: target, Value: value, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting Assignment, returning %+v\n", assignStmt) // Debug print
	return assignStmt
}

// VisitExprStmt builds an ExprStmt node.
func (v *ASTBuilder) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	fmt.Println("Visiting ExprStmt") // Debug print
	// The statement is just an expression
	expr := ctx.Expression().Accept(v).(Expression)
	exprStmt := &ExprStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting ExprStmt, returning %+v\n", exprStmt) // Debug print
	return exprStmt
}

// VisitPrintStmt builds a PrintStmt node.
func (v *ASTBuilder) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	fmt.Println("Visiting PrintStmt") // Debug print
	var exprs []Expression
	// Check if ExpressionListOpt has an actual ExpressionList
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		// Visit the expression list to get the slice of expressions
		// CORRECTED: Cast to []Expression instead of []Statement
		visitedExprs := el.Accept(v)                                                                // Get the result first
		fmt.Printf("  Visited ExpressionListOpt.ExpressionList, returned type: %T\n", visitedExprs) // Debug print
		if visitedExprsSlice, ok := visitedExprs.([]Expression); ok {                               // <--- This should be correct now
			exprs = visitedExprsSlice
		} else {
			fmt.Printf("  WARNING: Visiting ExpressionListOpt.ExpressionList did not return []Expression, got %T\n", visitedExprs) // Debug print
		}
	}
	printStmtNode := &PrintStmt{Exprs: exprs, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting PrintStmt, returning %+v\n", printStmtNode) // Debug print
	return printStmtNode                                                      // Ensure *ast.PrintStmt is returned
}

// VisitReturnStmt builds a ReturnStmt node.
func (v *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	fmt.Println("Visiting ReturnStmt") // Debug print
	var expr Expression
	// Check if ExpressionOpt has an actual Expression
	if ctx.ExpressionOpt().Expression() != nil {
		// Visit the optional expression
		expr = ctx.ExpressionOpt().Expression().Accept(v).(Expression)
	}
	returnStmtNode := &ReturnStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting ReturnStmt, returning %+v\n", returnStmtNode) // Debug print
	return returnStmtNode
}

// --- Compound statements ---

// VisitIfStmt builds an IfStmt node.
func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	fmt.Println("Visiting IfStmt") // Debug print
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
	fmt.Printf("Finished Visiting IfStmt, returning %+v\n", ifStmtNode) // Debug print
	return ifStmtNode
}

// VisitWhileStmt builds a WhileStmt node.
func (v *ASTBuilder) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	fmt.Println("Visiting WhileStmt") // Debug print
	// Visit the loop condition and body block
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	whileStmtNode := &WhileStmt{Cond: cond, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting WhileStmt, returning %+v\n", whileStmtNode) // Debug print
	return whileStmtNode
}

// VisitForStmt builds a ForStmt node.
func (v *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	fmt.Println("Visiting ForStmt") // Debug print
	// Get the variable name from the IDENTIFIER token
	varName := ctx.IDENTIFIER().GetText()
	// Visit the iterable expression
	iter := ctx.Expression().Accept(v).(Expression)
	// Visit the loop body block
	body := ctx.Block().Accept(v).(*BlockStmt)
	forStmtNode := &ForStmt{Variable: varName, Iterable: iter, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting ForStmt, returning %+v\n", forStmtNode) // Debug print
	return forStmtNode
}

// VisitFunctionDef builds a FuncDefStmt node.
func (v *ASTBuilder) VisitFunctionDef(ctx *parser.FunctionDefContext) interface{} {
	fmt.Println("Visiting FunctionDef") // Debug print
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
	// Note: Function name is not part of the grammar rule itself,
	// it would typically be assigned via an assignment statement,
	// e.g., `myFunc = function() { ... }`. The AST node reflects this.
	funcDefStmtNode := &FuncDefStmt{Name: "", Params: params, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting FunctionDef, returning %+v\n", funcDefStmtNode) // Debug print
	return funcDefStmtNode
}

// VisitBlock builds a BlockStmt node.
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	fmt.Println("Visiting Block") // Debug print
	var stmts []Statement
	// Check for optional statement list within the block
	if sl := ctx.StatementListOpt().StatementList(); sl != nil {
		// Visit the statement list to get the slice of statements
		visitedStmts := sl.Accept(v)                                                              // Get result first
		fmt.Printf("  Visited StatementListOpt.StatementList, returned type: %T\n", visitedStmts) // Debug print
		if visitedStmtsSlice, ok := visitedStmts.([]Statement); ok {
			stmts = visitedStmtsSlice
		} else {
			fmt.Printf("  WARNING: Visiting StatementListOpt.StatementList did not return []Statement, got %T\n", visitedStmts) // Debug print
		}
	}
	blockStmtNode := &BlockStmt{Stmts: stmts, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting Block, returning %+v\n", blockStmtNode) // Debug print
	return blockStmtNode
}

// VisitExpressionList handles a comma-separated list of expressions.
func (v *ASTBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	fmt.Println("Visiting ExpressionList") // Debug print
	var exprs []Expression
	// Iterate through all expression contexts in the list
	for _, e := range ctx.AllExpression() {
		// Visit each expression
		expr := e.Accept(v).(Expression)
		exprs = append(exprs, expr)
	}
	fmt.Printf("Finished Visiting ExpressionList, returning %d expressions\n", len(exprs)) // Debug print
	return exprs                                                                           // Returns []Expression
}

// VisitParamList handles a comma-separated list of identifiers for function parameters.
func (v *ASTBuilder) VisitParamList(ctx *parser.ParamListContext) interface{} {
	fmt.Println("Visiting ParamList") // Debug print
	var params []string
	// Iterate through all IDENTIFIER tokens in the parameter list
	for _, id := range ctx.AllIDENTIFIER() {
		params = append(params, id.GetText())
	}
	fmt.Printf("Finished Visiting ParamList, returning %d parameters\n", len(params)) // Debug print
	return params
}

// --- Expressions ---

// VisitExpression simply visits the top-level logicalOr expression.
func (v *ASTBuilder) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	fmt.Println("Visiting Expression") // Debug print
	expr := ctx.LogicalOr().Accept(v)
	fmt.Printf("Finished Visiting Expression, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitLogicalOr handles 'or' operations (left-associative).
func (v *ASTBuilder) VisitLogicalOr(ctx *parser.LogicalOrContext) interface{} {
	fmt.Println("Visiting LogicalOr") // Debug print
	// Start with the first logicalAnd expression
	expr := ctx.LogicalAnd(0).Accept(v).(Expression)
	// Iterate through children to find 'or' tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is an 'or' token using the generated token constant
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__15 { // Corrected token name
			right := ctx.LogicalAnd((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	fmt.Printf("Finished Visiting LogicalOr, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitLogicalAnd handles 'and' operations (left-associative).
func (v *ASTBuilder) VisitLogicalAnd(ctx *parser.LogicalAndContext) interface{} {
	fmt.Println("Visiting LogicalAnd") // Debug print
	// Start with the first comparison expression
	expr := ctx.Comparison(0).Accept(v).(Expression)
	// Iterate through children to find 'and' tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Corrected: Removed the extra .GetSymbol() call
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__16 { // Corrected token name
			right := ctx.Comparison((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	fmt.Printf("Finished Visiting LogicalAnd, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitComparison handles comparison operations (chainable, but typically left-associative evaluation).
func (v *ASTBuilder) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	fmt.Println("Visiting Comparison") // Debug print
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
			parser.InscriptParserT__22: // '>=' // Corrected token names
			right := ctx.Arith((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	fmt.Printf("Finished Visiting Comparison, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitArith handles addition and subtraction operations (left-associative).
func (v *ASTBuilder) VisitArith(ctx *parser.ArithContext) interface{} {
	fmt.Println("Visiting Arith") // Debug print
	// Start with the first term expression
	expr := ctx.Term(0).Accept(v).(Expression)
	// Iterate through children to find '+' or '-' operator tokens
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is an '+' or '-' token using generated token constants
		switch opNode.GetSymbol().GetTokenType() {
		case parser.InscriptParserT__23, // '+'
			parser.InscriptParserT__24: // '-' // Corrected token names
			right := ctx.Term((i + 1) / 2).Accept(v).(Expression)
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	fmt.Printf("Finished Visiting Arith, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitTerm handles multiplication, division, and modulo operations (left-associative).
func (v *ASTBuilder) VisitTerm(ctx *parser.TermContext) interface{} {
	fmt.Println("Visiting Term") // Debug print
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
			parser.InscriptParserT__27: // '%' // Corrected token names
			right := ctx.Factor((i + 1) / 2).Accept(v).(Expression) // Get the corresponding right factor
			// Build a BinaryExpr for each term operation
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	fmt.Printf("Finished Visiting Term, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitFactor handles exponentiation operations (right-associative).
func (v *ASTBuilder) VisitFactor(ctx *parser.FactorContext) interface{} {
	fmt.Println("Visiting Factor") // Debug print
	// Start with the first unary expression
	expr := ctx.Unary(0).Accept(v).(Expression)
	// Iterate through subsequent '^' operators and right-hand operands
	// Exponentiation is right-associative, so we build the AST from right to left conceptually,
	// or by chaining from left to right here, ensuring the right side is the result of the next factor.
	// The AST structure will naturally represent the right-associativity if built this way.
	for i := 1; i < ctx.GetChildCount(); i += 2 { // Iterate through children after the first Unary
		opNode := ctx.GetChild(i).(antlr.TerminalNode)
		// Check if the child is a '^' token using the generated token constant
		if opNode.GetSymbol().GetTokenType() == parser.InscriptParserT__28 { // Corrected token name for '^'
			right := ctx.Unary((i + 1) / 2).Accept(v).(Expression) // Get the corresponding right unary
			// Build a BinaryExpr for each exponentiation operation
			// The position token should ideally be the operator itself, but the grammar doesn't label it.
			// Using the operator's position is better.
			expr = &BinaryExpr{Left: expr, Operator: opNode.GetText(), Right: right, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
		}
	}
	fmt.Printf("Finished Visiting Factor, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitUnary handles unary operations ('+', '-', 'not').
func (v *ASTBuilder) VisitUnary(ctx *parser.UnaryContext) interface{} {
	fmt.Println("Visiting Unary") // Debug print
	// Check if it's a unary operation (has 2 children: operator and expression)
	if ctx.GetChildCount() == 2 {
		opNode := ctx.GetChild(0).(antlr.TerminalNode) // Get the operator token node
		op := opNode.GetText()                         // Get the operator text
		// Ensure we visit the correct child for the expression (it's the second child)
		if exprCtx, ok := ctx.GetChild(1).(*parser.UnaryContext); ok { // Unary can be recursive
			expr := exprCtx.Accept(v).(Expression)
			unaryExpr := &UnaryExpr{Operator: op, Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
			fmt.Printf("Finished Visiting Unary (op), returning type: %T\n", unaryExpr) // Debug print
			return unaryExpr
		}
		// Fallback if the child is not a UnaryContext (e.g., a Primary)
		if primaryCtx, ok := ctx.GetChild(1).(*parser.PrimaryContext); ok {
			expr := primaryCtx.Accept(v).(Expression)
			unaryExpr := &UnaryExpr{Operator: op, Expr: expr, PosToken: token.Pos(opNode.GetSymbol().GetStart())}
			fmt.Printf("Finished Visiting Unary (primary), returning type: %T\n", unaryExpr) // Debug print
			return unaryExpr
		}
		// Handle unexpected child type if necessary
		fmt.Printf("WARNING: Unary child 1 has unexpected type: %T\n", ctx.GetChild(1)) // Debug print
		return nil                                                                      // Or return an error
	}
	// If not a unary operation, it must be a primary expression
	primaryExpr := ctx.Primary().Accept(v)
	fmt.Printf("Finished Visiting Unary (primary only), returning type: %T\n", primaryExpr) // Debug print
	return primaryExpr
}

// VisitPrimary handles primary expressions, including indexing and function calls.
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	fmt.Println("Visiting Primary") // Debug print
	// Start with the base atom
	expr := ctx.Atom().Accept(v).(Expression)

	// Iterate through all children after the atom to find suffixes
	// The children are ordered: atom suffix1 suffix2 ...
	for i := 1; i < ctx.GetChildCount(); i++ {
		suffixChild := ctx.GetChild(i)

		// Check if the suffix is an indexing operation (starts with '[')
		if termNode, ok := suffixChild.(antlr.TerminalNode); ok && termNode.GetSymbol().GetTokenType() == parser.InscriptParserT__30 { // Token name for '['
			fmt.Println("  Visiting Primary (Index suffix)") // Debug print
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
					fmt.Printf("  WARNING: Primary index child %d has unexpected type: %T\n", i+1, ctx.GetChild(i+1)) // Debug print
					break
				}
			} else {
				// Handle unexpected end of input after '['
				fmt.Println("  WARNING: Unexpected end of input after '[' in Primary") // Debug print
				break
			}
		} else if termNode, ok := suffixChild.(antlr.TerminalNode); ok && termNode.GetSymbol().GetTokenType() == parser.InscriptParserT__2 { // Token name for '('
			fmt.Println("  Visiting Primary (Call suffix)") // Debug print
			// The suffix is a function call (starts with '(')
			// The structure is primary -> atom '(' expressionListOpt ')'
			// The ExpressionListOpt is the next child at index i+1
			if i+1 < ctx.GetChildCount() {
				// Assert the child is an ExpressionListOptContext before calling Accept
				if expressionListOptCtx, ok := ctx.GetChild(i + 1).(*parser.ExpressionListOptContext); ok {
					var args []Expression
					// Check if there's an actual expression list
					if el := expressionListOptCtx.ExpressionList(); el != nil {
						// Visit the expression list to get the arguments
						visitedArgs := el.Accept(v)                                                                                  // Get result first
						fmt.Printf("  Visited ExpressionListOpt.ExpressionList for CallExpr args, returned type: %T\n", visitedArgs) // Debug print
						if visitedArgsSlice, ok := visitedArgs.([]Expression); ok {
							args = visitedArgsSlice
						} else {
							fmt.Printf("  WARNING: Visiting ExpressionListOpt.ExpressionList for CallExpr args did not return []Expression, got %T\n", visitedArgs) // Debug print
						}
					}
					// Wrap the current expression in a CallExpr node using the correct field name 'Callee'
					expr = &CallExpr{Callee: expr, Args: args, PosToken: token.Pos(termNode.GetSymbol().GetStart())} // Use '(' token position
					i += 2                                                                                           // Skip the expressionListOpt and the closing ')'
				} else {
					// Handle unexpected child type
					fmt.Printf("  WARNING: Primary call child %d has unexpected type: %T\n", i+1, ctx.GetChild(i+1)) // Debug print
					break
				}
			} else {
				// Handle unexpected end of input after '('
				fmt.Println("  WARNING: Unexpected end of input after '(' in Primary") // Debug print
				break
			}
		} else {
			// If the child is neither '[' nor '(', it's an unexpected child in primary
			// The parser should ideally catch this.
			fmt.Printf("  WARNING: Unexpected child type in Primary suffix at index %d: %T\n", i, suffixChild) // Debug print
		}
	}
	fmt.Printf("Finished Visiting Primary, returning type: %T\n", expr) // Debug print
	return expr
}

// VisitAtom handles the different types of atoms (literals, identifiers, lists, tables, parenthesized expressions).
func (v *ASTBuilder) VisitAtom(ctx *parser.AtomContext) interface{} {
	fmt.Println("Visiting Atom") // Debug print
	// Check which alternative the atom matches and visit accordingly
	if lit := ctx.Literal(); lit != nil {
		result := lit.Accept(v)
		fmt.Printf("Finished Visiting Atom (Literal), returning type: %T\n", result) // Debug print
		return result
	}
	if id := ctx.IDENTIFIER(); id != nil {
		// Build an Identifier node
		identifierNode := &Identifier{Name: id.GetText(), PosToken: token.Pos(id.GetSymbol().GetStart())}
		fmt.Printf("Finished Visiting Atom (Identifier), returning type: %T\n", identifierNode) // Debug print
		return identifierNode
	}
	if ll := ctx.ListLiteral(); ll != nil {
		result := ll.Accept(v)
		fmt.Printf("Finished Visiting Atom (ListLiteral), returning type: %T\n", result) // Debug print
		return result
	}
	if tl := ctx.TableLiteral(); tl != nil {
		result := tl.Accept(v)
		fmt.Printf("Finished Visiting Atom (TableLiteral), returning type: %T\n", result) // Debug print
		return result
	}
	// If none of the above, it must be a parenthesized expression
	// The structure is atom -> '(' expression ')'
	// The expression is the child at index 1. We need to assert its type.
	if exprCtx, ok := ctx.GetChild(1).(*parser.ExpressionContext); ok { // Assert type to ExpressionContext
		result := exprCtx.Accept(v)
		fmt.Printf("Finished Visiting Atom (Parenthesized Expression), returning type: %T\n", result) // Debug print
		return result
	}
	// This case should ideally not be reached if the parser succeeded.
	fmt.Printf("WARNING: Visiting Atom, unexpected context type: %T\n", ctx) // Debug print
	return nil
}

// VisitListLiteral builds a ListLiteral node.
func (v *ASTBuilder) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	fmt.Println("Visiting ListLiteral") // Debug print
	var elems []Expression
	// Check for optional expression list within the brackets
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		// Visit the expression list to get the elements
		visitedElems := el.Accept(v)                                                                                // Get result first
		fmt.Printf("  Visited ExpressionListOpt.ExpressionList for ListLiteral, returned type: %T\n", visitedElems) // Debug print
		if visitedElemsSlice, ok := visitedElems.([]Expression); ok {
			elems = visitedElemsSlice
		} else {
			fmt.Printf("  WARNING: Visiting ExpressionListOpt.ExpressionList for ListLiteral did not return []Expression, got %T\n", visitedElems) // Debug print
		}
	}
	listLiteralNode := &ListLiteral{Elements: elems, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting ListLiteral, returning %+v\n", listLiteralNode) // Debug print
	return listLiteralNode
}

// VisitTableLiteral builds a TableLiteral node.
func (v *ASTBuilder) VisitTableLiteral(ctx *parser.TableLiteralContext) interface{} {
	fmt.Println("Visiting TableLiteral") // Debug print
	var fields []Field
	// Check for optional field list within the braces
	if fl := ctx.FieldListOpt().FieldList(); fl != nil {
		// Iterate through all field contexts
		for _, fctx := range fl.AllField() {
			// Get the field key (identifier) and visit the field value
			key := fctx.IDENTIFIER().GetText()
			val := fctx.Expression().Accept(v).(Expression)
			// Append a new Field struct to the slice
			fields = append(fields, Field{Key: key, Value: val, PosToken: token.Pos(fctx.GetStart().GetStart())})
		}
	}
	tableLiteralNode := &TableLiteral{Fields: fields, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting TableLiteral, returning %+v\n", tableLiteralNode) // Debug print
	return tableLiteralNode
}

// VisitField builds a Field struct for table literals.
// Note: This is called by VisitTableLiteral, it doesn't return an AST node itself.
func (v *ASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	// This method is visited, but the logic to build the Field struct
	// is handled directly in VisitTableLiteral to collect all fields.
	// This method can be left empty or return nil.
	fmt.Println("Visiting Field (should be handled by VisitTableLiteral)") // Debug print
	return nil
}

// VisitExpressionListOpt handles the optional expression list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitExpressionListOpt(ctx *parser.ExpressionListOptContext) interface{} {
	fmt.Println("Visiting ExpressionListOpt") // Debug print
	if ctx.ExpressionList() != nil {
		result := ctx.ExpressionList().Accept(v)
		fmt.Printf("Finished Visiting ExpressionListOpt (has list), returning type: %T\n", result) // Debug print
		return result
	}
	fmt.Println("Finished Visiting ExpressionListOpt (empty), returning nil") // Debug print
	return nil                                                                // Represents an empty list
}

// VisitParamListOpt handles the optional parameter list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitParamListOpt(ctx *parser.ParamListOptContext) interface{} {
	fmt.Println("Visiting ParamListOpt") // Debug print
	if ctx.ParamList() != nil {
		result := ctx.ParamList().Accept(v)
		fmt.Printf("Finished Visiting ParamListOpt (has list), returning type: %T\n", result) // Debug print
		return result
	}
	fmt.Println("Finished Visiting ParamListOpt (empty), returning nil") // Debug print
	return nil                                                           // Represents no parameters
}

// VisitStatementListOpt handles the optional statement list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitStatementListOpt(ctx *parser.StatementListOptContext) interface{} {
	fmt.Println("Visiting StatementListOpt") // Debug print
	if ctx.StatementList() != nil {
		result := ctx.StatementList().Accept(v)
		fmt.Printf("Finished Visiting StatementListOpt (has list), returning type: %T\n", result) // Debug print
		return result
	}
	fmt.Println("Finished Visiting StatementListOpt (empty), returning nil") // Debug print
	return nil                                                               // Represents an empty block
}

// VisitExpressionOpt handles the optional expression.
// If the expression is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitExpressionOpt(ctx *parser.ExpressionOptContext) interface{} {
	fmt.Println("Visiting ExpressionOpt") // Debug print
	if ctx.Expression() != nil {
		result := ctx.Expression().Accept(v)
		fmt.Printf("Finished Visiting ExpressionOpt (has expr), returning type: %T\n", result) // Debug print
		return result
	}
	fmt.Println("Finished Visiting ExpressionOpt (empty), returning nil") // Debug print
	return nil                                                            // Represents no expression
}

// VisitElseifListOpt handles the optional elseif list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitElseifListOpt(ctx *parser.ElseifListOptContext) interface{} {
	fmt.Println("Visiting ElseifListOpt") // Debug print
	if ctx.ElseifList() != nil {
		result := ctx.ElseifList().Accept(v)
		fmt.Printf("Finished Visiting ElseifListOpt (has list), returning type: %T\n", result) // Debug print
		return result
	}
	fmt.Println("Finished Visiting ElseifListOpt (empty), returning nil") // Debug print
	return nil                                                            // Represents no elseif clauses
}

// VisitElseifList handles one or more elseif clauses.
func (v *ASTBuilder) VisitElseifList(ctx *parser.ElseifListContext) interface{} {
	fmt.Println("Visiting ElseifList") // Debug print
	var elseIfs []ElseIf
	for _, eis := range ctx.AllElseif() {
		// Visit each elseif clause, which should return an ElseIf struct
		if elseIf, ok := eis.Accept(v).(ElseIf); ok {
			elseIfs = append(elseIfs, elseIf)
		} else {
			fmt.Printf("WARNING: Visiting ElseifList, elseif context type %T did not return ElseIf, got %T\n", eis, eis.Accept(v)) // Debug print
		}
	}
	fmt.Printf("Finished Visiting ElseifList, returning %d ElseIfs\n", len(elseIfs)) // Debug print
	return elseIfs
}

// VisitElseif handles a single elseif clause.
// Note: This is called by VisitElseifList, it doesn't return an AST node itself,
// but a struct used within the IfStmt node.
func (v *ASTBuilder) VisitElseif(ctx *parser.ElseifContext) interface{} {
	fmt.Println("Visiting Elseif") // Debug print
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	elseIfStruct := ElseIf{Cond: cond, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting Elseif, returning %+v\n", elseIfStruct) // Debug print
	return elseIfStruct                                                   // Return the struct value
}

// VisitElseBlockOpt handles the optional else block.
// If the block is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitElseBlockOpt(ctx *parser.ElseBlockOptContext) interface{} {
	fmt.Println("Visiting ElseBlockOpt") // Debug print
	if ctx.Block() != nil {
		result := ctx.Block().Accept(v)
		fmt.Printf("Finished Visiting ElseBlockOpt (has block), returning type: %T\n", result) // Debug print
		return result
	}
	fmt.Println("Finished Visiting ElseBlockOpt (empty), returning nil") // Debug print
	return nil                                                           // Represents no else block
}

// VisitLiteral handles the different types of literals.
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	fmt.Println("Visiting Literal") // Debug print
	// Check which type of literal it is and parse/build the corresponding AST node
	if i := ctx.INTEGER(); i != nil {
		val, _ := strconv.ParseInt(i.GetText(), 10, 64) // Error handling omitted for brevity
		integerLiteralNode := &IntegerLiteral{Value: val, PosToken: token.Pos(i.GetSymbol().GetStart())}
		fmt.Printf("Finished Visiting Literal (Integer), returning %+v\n", integerLiteralNode) // Debug print
		return integerLiteralNode
	}
	if f := ctx.FLOAT(); f != nil {
		val, _ := strconv.ParseFloat(f.GetText(), 64) // Error handling omitted for brevity
		floatLiteralNode := &FloatLiteral{Value: val, PosToken: token.Pos(f.GetSymbol().GetStart())}
		fmt.Printf("Finished Visiting Literal (Float), returning %+v\n", floatLiteralNode) // Debug print
		return floatLiteralNode
	}
	if s := ctx.STRING(); s != nil {
		// Remove quotes from the string literal
		text := s.GetText()[1 : len(s.GetText())-1]
		// TODO: Handle escape sequences within the string
		stringLiteralNode := &StringLiteral{Value: text, PosToken: token.Pos(s.GetSymbol().GetStart())}
		fmt.Printf("Finished Visiting Literal (String), returning %+v\n", stringLiteralNode) // Debug print
		return stringLiteralNode
	}
	if b := ctx.BOOLEAN(); b != nil {
		booleanLiteralNode := &BooleanLiteral{Value: b.GetText() == "true", PosToken: token.Pos(b.GetSymbol().GetStart())}
		fmt.Printf("Finished Visiting Literal (Boolean), returning %+v\n", booleanLiteralNode) // Debug print
		return booleanLiteralNode
	}
	// Must be 'nil'
	nilLiteralNode := &NilLiteral{PosToken: token.Pos(ctx.GetStart().GetStart())}
	fmt.Printf("Finished Visiting Literal (Nil), returning %+v\n", nilLiteralNode) // Debug print
	return nilLiteralNode
}

// Ensure ASTBuilder satisfies parser.InscriptVisitor
var _ parser.InscriptVisitor = (*ASTBuilder)(nil)
