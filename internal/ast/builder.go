package ast

import (
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
	prog := &Program{PosToken: token.Pos(ctx.GetStart().GetStart())}
	// The grammar ensures StatementList is present if there are statements
	if ctx.StatementList() != nil {
		// VisitStatementList returns []Statement, which we append to the program's Stmts
		if stmts, ok := ctx.StatementList().Accept(v).([]Statement); ok {
			prog.Stmts = append(prog.Stmts, stmts...)
		}
	}
	return prog
}

// VisitStatementList handles one or more statements
func (v *ASTBuilder) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	var stmts []Statement
	// Iterate through all statement contexts and visit them
	for _, sctx := range ctx.AllStatement() {
		// Each statement visit should return a Statement node
		if stmt, ok := sctx.Accept(v).(Statement); ok {
			stmts = append(stmts, stmt)
		}
	}
	return stmts // Returns []Statement
}

// VisitStatement handles a single statement (either simple or compound).
// This method is crucial for ensuring the result from sub-visits is propagated.
func (v *ASTBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	// A statement context has only one child which is either a simpleStmt or compoundStmt
	if ctx.SimpleStmt() != nil {
		return ctx.SimpleStmt().Accept(v) // Return the result from visiting the simple statement
	}
	if ctx.CompoundStmt() != nil {
		return ctx.CompoundStmt().Accept(v) // Return the result from visiting the compound statement
	}
	return nil // Should not happen if grammar is correct
}

// VisitSimpleStmt handles a simple statement (assignment, exprStmt, printStmt, returnStmt).
// This method is crucial for ensuring the result from sub-visits is propagated.
func (v *ASTBuilder) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	// A simpleStmt context has only one child which is one of the simple statement types
	if ctx.Assignment() != nil {
		return ctx.Assignment().Accept(v) // Return the result from visiting the assignment
	}
	if ctx.ExprStmt() != nil {
		return ctx.ExprStmt().Accept(v) // Return the result from visiting the expression statement
	}
	if ctx.PrintStmt() != nil {
		return ctx.PrintStmt().Accept(v) // Return the result from visiting the print statement
	}
	if ctx.ReturnStmt() != nil {
		return ctx.ReturnStmt().Accept(v) // Return the result from visiting the return statement
	}
	return nil // Should not happen if grammar is correct
}

// --- Simple statements ---

// VisitAssignment builds an AssignStmt node.
func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	// Target must be a Primary expression (e.g., identifier, index, call)
	target := ctx.Primary().Accept(v).(Expression)
	// Value is any Expression
	value := ctx.Expression().Accept(v).(Expression)
	return &AssignStmt{Target: target, Value: value, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitExprStmt builds an ExprStmt node.
func (v *ASTBuilder) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	// The statement is just an expression
	expr := ctx.Expression().Accept(v).(Expression)
	return &ExprStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitPrintStmt builds a PrintStmt node.
func (v *ASTBuilder) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	var exprs []Expression
	// Check if ExpressionListOpt has an actual ExpressionList
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		// Visit the expression list to get the slice of expressions
		if visitedExprs, ok := el.Accept(v).([]Expression); ok {
			exprs = visitedExprs
		}
	}
	return &PrintStmt{Exprs: exprs, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitReturnStmt builds a ReturnStmt node.
func (v *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	var expr Expression
	// Check if ExpressionOpt has an actual Expression
	if ctx.ExpressionOpt().Expression() != nil {
		// Visit the optional expression
		expr = ctx.ExpressionOpt().Expression().Accept(v).(Expression)
	}
	return &ReturnStmt{Expr: expr, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// --- Compound statements ---

// VisitIfStmt builds an IfStmt node.
func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
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

	return &IfStmt{Cond: cond, Then: then, ElseIfs: elseIfs, Else: elseBlk, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitWhileStmt builds a WhileStmt node.
func (v *ASTBuilder) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	// Visit the loop condition and body block
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &WhileStmt{Cond: cond, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitForStmt builds a ForStmt node.
func (v *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	// Get the variable name from the IDENTIFIER token
	varName := ctx.IDENTIFIER().GetText()
	// Visit the iterable expression
	iter := ctx.Expression().Accept(v).(Expression)
	// Visit the loop body block
	body := ctx.Block().Accept(v).(*BlockStmt)
	return &ForStmt{Variable: varName, Iterable: iter, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitFunctionDef builds a FuncDefStmt node.
func (v *ASTBuilder) VisitFunctionDef(ctx *parser.FunctionDefContext) interface{} {
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
	return &FuncDefStmt{Name: "", Params: params, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitBlock builds a BlockStmt node.
func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	var stmts []Statement
	// Check for optional statement list within the block
	if sl := ctx.StatementListOpt().StatementList(); sl != nil {
		// Visit the statement list to get the slice of statements
		if visitedStmts, ok := sl.Accept(v).([]Statement); ok {
			stmts = visitedStmts
		}
	}
	return &BlockStmt{Stmts: stmts, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitExpressionList handles a comma-separated list of expressions.
func (v *ASTBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	var exprs []Expression
	// Iterate through all expression contexts in the list
	for _, e := range ctx.AllExpression() {
		// Visit each expression
		exprs = append(exprs, e.Accept(v).(Expression))
	}
	return exprs // Returns []Expression
}

// VisitParamList handles a comma-separated list of identifiers for function parameters.
func (v *ASTBuilder) VisitParamList(ctx *parser.ParamListContext) interface{} {
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
	return ctx.LogicalOr().Accept(v)
}

// VisitLogicalOr handles 'or' operations (left-associative).
func (v *ASTBuilder) VisitLogicalOr(ctx *parser.LogicalOrContext) interface{} {
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
	return expr
}

// VisitLogicalAnd handles 'and' operations (left-associative).
func (v *ASTBuilder) VisitLogicalAnd(ctx *parser.LogicalAndContext) interface{} {
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
	return expr
}

// VisitComparison handles comparison operations (chainable, but typically left-associative evaluation).
func (v *ASTBuilder) VisitComparison(ctx *parser.ComparisonContext) interface{} {
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
	return expr
}

// VisitArith handles addition and subtraction operations (left-associative).
func (v *ASTBuilder) VisitArith(ctx *parser.ArithContext) interface{} {
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
	return expr
}

// VisitTerm handles multiplication, division, and modulo operations (left-associative).
func (v *ASTBuilder) VisitTerm(ctx *parser.TermContext) interface{} {
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
	return expr
}

// VisitFactor handles exponentiation operations (right-associative).
func (v *ASTBuilder) VisitFactor(ctx *parser.FactorContext) interface{} {
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
	return expr
}

// VisitUnary handles unary operations ('+', '-', 'not').
func (v *ASTBuilder) VisitUnary(ctx *parser.UnaryContext) interface{} {
	// Check if it's a unary operation (has 2 children: operator and expression)
	if ctx.GetChildCount() == 2 {
		opNode := ctx.GetChild(0).(antlr.TerminalNode) // Get the operator token node
		op := opNode.GetText()                         // Get the operator text
		// Ensure we visit the correct child for the expression (it's the second child)
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
	}
	// If not a unary operation, it must be a primary expression
	return ctx.Primary().Accept(v)
}

// VisitPrimary handles primary expressions, including indexing and function calls.
func (v *ASTBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	// Start with the base atom
	expr := ctx.Atom().Accept(v).(Expression)

	// Iterate through all children after the atom to find suffixes
	// The children are ordered: atom suffix1 suffix2 ...
	for i := 1; i < ctx.GetChildCount(); i++ {
		suffixChild := ctx.GetChild(i)

		// Check if the suffix is an indexing operation (starts with '[')
		if termNode, ok := suffixChild.(antlr.TerminalNode); ok && termNode.GetSymbol().GetTokenType() == parser.InscriptParserT__30 { // Token name for '['
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
		} else if termNode, ok := suffixChild.(antlr.TerminalNode); ok && termNode.GetSymbol().GetTokenType() == parser.InscriptParserT__2 { // Token name for '('
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

// VisitAtom handles the different types of atoms (literals, identifiers, lists, tables, parenthesized expressions).
func (v *ASTBuilder) VisitAtom(ctx *parser.AtomContext) interface{} {
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
	// If none of the above, it must be a parenthesized expression
	// The structure is atom -> '(' expression ')'
	// The expression is the child at index 1. We need to assert its type.
	if exprCtx, ok := ctx.GetChild(1).(*parser.ExpressionContext); ok { // Assert type to ExpressionContext
		return exprCtx.Accept(v)
	}
	// This case should ideally not be reached if the parser succeeded.
	return nil
}

// VisitListLiteral builds a ListLiteral node.
func (v *ASTBuilder) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	var elems []Expression
	// Check for optional expression list within the brackets
	if el := ctx.ExpressionListOpt().ExpressionList(); el != nil {
		// Visit the expression list to get the elements
		if visitedElems, ok := el.Accept(v).([]Expression); ok {
			elems = visitedElems
		}
	}
	return &ListLiteral{Elements: elems, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitTableLiteral builds a TableLiteral node.
func (v *ASTBuilder) VisitTableLiteral(ctx *parser.TableLiteralContext) interface{} {
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
	return &TableLiteral{Fields: fields, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitField builds a Field struct for table literals.
// Note: This is called by VisitTableLiteral, it doesn't return an AST node itself.
func (v *ASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	// This method is visited, but the logic to build the Field struct
	// is handled directly in VisitTableLiteral to collect all fields.
	// This method can be left empty or return nil.
	return nil
}

// VisitExpressionListOpt handles the optional expression list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitExpressionListOpt(ctx *parser.ExpressionListOptContext) interface{} {
	if ctx.ExpressionList() != nil {
		return ctx.ExpressionList().Accept(v)
	}
	return nil // Represents an empty list
}

// VisitParamListOpt handles the optional parameter list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitParamListOpt(ctx *parser.ParamListOptContext) interface{} {
	if ctx.ParamList() != nil {
		return ctx.ParamList().Accept(v)
	}
	return nil // Represents no parameters
}

// VisitStatementListOpt handles the optional statement list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitStatementListOpt(ctx *parser.StatementListOptContext) interface{} {
	if ctx.StatementList() != nil {
		return ctx.StatementList().Accept(v)
	}
	return nil // Represents an empty block
}

// VisitExpressionOpt handles the optional expression.
// If the expression is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitExpressionOpt(ctx *parser.ExpressionOptContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(v)
	}
	return nil // Represents no expression
}

// VisitElseifListOpt handles the optional elseif list.
// If the list is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitElseifListOpt(ctx *parser.ElseifListOptContext) interface{} {
	if ctx.ElseifList() != nil {
		return ctx.ElseifList().Accept(v)
	}
	return nil // Represents no elseif clauses
}

// VisitElseifList handles one or more elseif clauses.
func (v *ASTBuilder) VisitElseifList(ctx *parser.ElseifListContext) interface{} {
	var elseIfs []ElseIf
	for _, eis := range ctx.AllElseif() {
		// Visit each elseif clause, which should return an ElseIf struct
		if elseIf, ok := eis.Accept(v).(ElseIf); ok {
			elseIfs = append(elseIfs, elseIf)
		}
	}
	return elseIfs
}

// VisitElseif handles a single elseif clause.
// Note: This is called by VisitElseifList, it doesn't return an AST node itself,
// but a struct used within the IfStmt node.
func (v *ASTBuilder) VisitElseif(ctx *parser.ElseifContext) interface{} {
	cond := ctx.Expression().Accept(v).(Expression)
	body := ctx.Block().Accept(v).(*BlockStmt)
	return ElseIf{Cond: cond, Body: body, PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// VisitElseBlockOpt handles the optional else block.
// If the block is present, it visits it. Otherwise, it returns nil.
func (v *ASTBuilder) VisitElseBlockOpt(ctx *parser.ElseBlockOptContext) interface{} {
	if ctx.Block() != nil {
		return ctx.Block().Accept(v)
	}
	return nil // Represents no else block
}

// VisitLiteral handles the different types of literals.
func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	// Check which type of literal it is and parse/build the corresponding AST node
	if i := ctx.INTEGER(); i != nil {
		val, _ := strconv.ParseInt(i.GetText(), 10, 64) // Error handling omitted for brevity
		return &IntegerLiteral{Value: val, PosToken: token.Pos(i.GetSymbol().GetStart())}
	}
	if f := ctx.FLOAT(); f != nil {
		val, _ := strconv.ParseFloat(f.GetText(), 64) // Error handling omitted for brevity
		return &FloatLiteral{Value: val, PosToken: token.Pos(f.GetSymbol().GetStart())}
	}
	if s := ctx.STRING(); s != nil {
		// Remove quotes from the string literal
		text := s.GetText()[1 : len(s.GetText())-1]
		// TODO: Handle escape sequences within the string
		return &StringLiteral{Value: text, PosToken: token.Pos(s.GetSymbol().GetStart())}
	}
	if b := ctx.BOOLEAN(); b != nil {
		return &BooleanLiteral{Value: b.GetText() == "true", PosToken: token.Pos(b.GetSymbol().GetStart())}
	}
	// Must be 'nil'
	return &NilLiteral{PosToken: token.Pos(ctx.GetStart().GetStart())}
}

// Ensure ASTBuilder satisfies parser.InscriptVisitor
var _ parser.InscriptVisitor = (*ASTBuilder)(nil)
