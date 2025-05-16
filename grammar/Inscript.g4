// Inscript.g4

grammar Inscript;

// Parser Rules

program: (statement? NEWLINE*)* EOF;

statement
    : exprStmt
    | assignment
    | ifStmt
    | whileStmt
    | forStmt
    | funcDef
    | breakStmt
    | continueStmt
    | returnStmt
    | importStmt
    | printStmt
    | block
    ;

block: LBRACE (statement? NEWLINE*)* RBRACE;

exprStmt: expression;

assignment
    : target ASSIGN expression
    ;

target
    : IDENTIFIER
    | indexExpr
    | attrExpr
    ;

ifStmt
    : IF expression block (ELSE block)?
    ;

whileStmt: WHILE expression block;

forStmt
    : FOR IDENTIFIER IN expression block
    ;

funcDef
    : FUNCTION IDENTIFIER LPAREN paramList? RPAREN (ARROW typeAnnotation)? block
    ;

paramList
    : param (COMMA param)* (COMMA)?
    ;

param
    : IDENTIFIER (ASSIGN expression)? (COLON typeAnnotation)?
    | ELLIPSIS IDENTIFIER
    ;

typeAnnotation: IDENTIFIER;

breakStmt: BREAK;
continueStmt: CONTINUE;
returnStmt: RETURN expression?;
importStmt: IMPORT STRING;
printStmt: PRINT LPAREN (expression (COMMA expression)*)? RPAREN;

expression
    : expression POW expression         #expExpr
    | expression MUL expression         #mulExpr
    | expression DIV expression         #divExpr
    | expression IDIV expression        #idivExpr
    | expression MOD expression         #modExpr
    | expression ADD expression         #addExpr
    | expression SUB expression         #subExpr
    | expression BITAND expression      #bitandExpr
    | expression BITOR expression       #bitorExpr
    | expression BITXOR expression      #bitxorExpr
    | expression SHL expression         #shlExpr
    | expression SHR expression         #shrExpr
    | expression LT expression          #ltExpr
    | expression LE expression          #leExpr
    | expression GT expression          #gtExpr
    | expression GE expression          #geExpr
    | expression EQ expression          #eqExpr
    | expression NEQ expression         #neqExpr
    | expression AND expression         #andExpr
    | expression OR expression          #orExpr
    | NOT expression                    #notExpr
    | BITNOT expression                 #bitnotExpr
    | SUB expression                    #negExpr
    | primary                           #primaryExpr
    ;

primary
    : literal
    | IDENTIFIER
    | LPAREN expression RPAREN
    | LPAREN expression (COMMA expression)+ RPAREN // tuple
    | listLiteral
    | tableLiteral
    | callExpr
    | indexExpr
    | attrExpr
    ;

callExpr: expression LPAREN (expression (COMMA expression)*)? RPAREN;
indexExpr: expression LBRACK expression RBRACK;
attrExpr: expression DOT IDENTIFIER;

literal
    : NUMBER
    | STRING
    | TRUE
    | FALSE
    | NIL
    ;

listLiteral: LBRACK (expression (COMMA expression)*)? RBRACK;
tableLiteral: LBRACE (tableEntry (COMMA tableEntry)*)? RBRACE;
tableEntry: (expression | STRING | NUMBER) ASSIGN expression;

// Lexer Rules

FUNCTION: 'function';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
FOR: 'for';
IN: 'in';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
IMPORT: 'import';
PRINT: 'print';
TRUE: 'true';
FALSE: 'false';
NIL: 'nil';
AND: 'and';
OR: 'or';
NOT: 'not';
POW: '\^\^';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
IDIV: '//' ;
MOD: '%';
BITAND: '&';
BITOR: '\|';
BITXOR: '\^';
BITNOT: '~';
SHL: '<<';
SHR: '>>';

EQ: '==';
NEQ: '!=';
LT: '<';
LE: '<=';
GT: '>';
GE: '>=';
ASSIGN: '=';
ARROW: '->';

LPAREN: '(';
RPAREN: ')';
LBRACK: '[';
RBRACK: ']';
LBRACE: '{';
RBRACE: '}';
COMMA: ',';
DOT: '.';
COLON: ':';
ELLIPSIS: '...';

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
NUMBER: [0-9]+ ('.' [0-9]+)?;
STRING
    : '"' (ESC_SEQ | ~["\\\n\r])* '"'
    | '\'' (ESC_SEQ | ~['\\\n\r])* '\''
    | '"""' .*? '"""'
    | "'''" .*? "'''"
    ;

fragment ESC_SEQ: '\\' [btnr"'\\];

COMMENT: '#' ~[\n\r]* -> skip;
MULTILINE_COMMENT: '#' .*? '#' -> skip;
BLOCK_COMMENT: '/\*' .*? '\*/' -> skip;

NEWLINE: [\r\n]+;
WS: [ \t]+ -> skip;
