grammar Inscript;

// -------------------------
// Parser Rules
// -------------------------

program
    : statement (SEPARATOR statement)* SEPARATOR? EOF
    ;

statement
    : simpleStmt
    | compoundStmt
    ;

simpleStmt
    : assignment
    | exprStmt
    | printStmt
    | returnStmt
    ;

assignment
    : primary '=' expression
    ;

exprStmt
    : expression
    ;

printStmt
    : 'print' '(' expressionListOpt ')'
    ;

returnStmt
    : 'return' expressionOpt
    ;

compoundStmt
    : ifStmt
    | whileStmt
    | forStmt
    | functionDef
    ;

// If / elseif / else
ifStmt
    : 'if' expression block elseifListOpt elseBlockOpt
    ;

elseifListOpt
    : /* empty */
    | elseif+
    ;

elseif
    : 'elseif' expression block
    ;

elseBlockOpt
    : /* empty */
    | 'else' block
    ;

// Loops
whileStmt
    : 'while' expression block
    ;

forStmt
    : 'for' IDENTIFIER 'in' expression block
    ;

// Function definition (statement)
functionDef
    : 'function' '(' paramListOpt ')' block
    ;

// Block of statements, with optional separators
block
    : '{' statementListOpt '}'
    ;

statementListOpt
    : (statement (SEPARATOR statement)*)? SEPARATOR?
    ;

// Expressions
expression
    : logicalOr
    ;

logicalOr
    : logicalAnd ('or' logicalAnd)*
    ;

logicalAnd
    : comparison ('and' comparison)*
    ;

comparison
    : arith (('==' | '!=' | '<' | '>' | '<=' | '>=') arith)*
    ;

arith
    : term (('+' | '-') term)*
    ;

term
    : factor (('*' | '/' | '%') factor)*
    ;

factor
    : unary ('^' unary)*
    ;

unary
    : ('+' | '-' | 'not') unary
    | primary
    ;

primary
    : atom (   '[' expression ']'        // index lookup
             | '(' expressionListOpt ')' // function call
            )*
    ;

atom
    : literal
    | IDENTIFIER
    | listLiteral
    | tableLiteral
    | '(' expression ')'
    | fnLiteral
    ;

fnLiteral
    : 'function' '(' paramListOpt ')' block
    ;

listLiteral
    : '[' expressionListOpt ']'
    ;

tableLiteral
    : '{' fieldListOpt '}'
    ;

fieldListOpt
    : /* empty */
    | field (',' field)*
    ;

field
    : IDENTIFIER '=' expression
    ;

// Optional expression or parameter lists
expressionOpt
    : /* empty */
    | expression
    ;

expressionListOpt
    : /* empty */
    | expressionList
    ;

expressionList
    : expression (',' expression)*
    ;

paramListOpt
    : /* empty */
    | paramList
    ;

paramList
    : IDENTIFIER (',' IDENTIFIER)*
    ;

// Literals
literal
    : INTEGER
    | FLOAT
    | STRING
    | BOOLEAN
    | 'nil'
    ;

// -------------------------
// Lexer Rules
// -------------------------

// Treat either newline or semicolon as statement terminator
SEPARATOR
    : ';'
    | NEWLINE
    ;

// Newlines become separators, not just skipped
NEWLINE
    : '\r'? '\n'
    ;

// Whitespace (spaces/tabs) still skipped
WS
    : [ \t]+ -> skip
    ;

// Comments
LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> skip
    ;

// Basic tokens
BOOLEAN
    : 'true'
    | 'false'
    ;

IDENTIFIER
    : LETTER (LETTER | DIGIT | '_')*
    ;

INTEGER
    : DIGIT+
    ;

FLOAT
    : DIGIT+ '.' DIGIT+
    ;

STRING
    : '"' (ESC_SEQ | ~["\\])* '"'
    ;

fragment DIGIT
    : [0-9]
    ;

fragment LETTER
    : [a-zA-Z_]
    ;

fragment ESC_SEQ
    : '\\' ["'\\ntbr]
    ;
