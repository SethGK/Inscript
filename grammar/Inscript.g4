grammar Inscript;

// Parser Rules

program
    : (statement)*
        EOF
    ;

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

block
    : LBRACE // Statements in a block, with whitespace (including newlines) implicitly skipped
        (statement)*
        RBRACE
    ;

exprStmt: expression;

assignment
    : target (ASSIGN | ADD_ASSIGN | SUB_ASSIGN | MUL_ASSIGN | DIV_ASSIGN | POW_ASSIGN) expression
    ;

target
    : IDENTIFIER
    | postfixExpr LBRACK expression RBRACK
    | postfixExpr DOT IDENTIFIER
    ;

ifStmt
    : IF expression block (ELSE block)?
    ;

whileStmt
    : WHILE expression block
    ;

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

// Left-recursive expression rules, unchanged
expression
    : unaryExpr                         #unaryExpression
    | expression POW expression         #expExpr
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
    ;

unaryExpr
    : NOT unaryExpr                     #notExpr
    | BITNOT unaryExpr                  #bitnotExpr
    | SUB unaryExpr                     #negExpr
    | postfixExpr                       #postfixExpression
    ;

postfixExpr
    : primary                                          #primaryPostfix
    | postfixExpr LPAREN argList? RPAREN              #callPostfix
    | postfixExpr LBRACK expression RBRACK            #indexPostfix
    | postfixExpr DOT IDENTIFIER                      #attrPostfix
    ;

argList: expression (COMMA expression)*;

primary
    : literal
    | IDENTIFIER
    | LPAREN expression RPAREN
    | LPAREN expression (COMMA expression)+ RPAREN     // tuple
    | listLiteral
    | tableLiteral
    ;

literal
    : NUMBER
    | STRING
    | TRUE
    | FALSE
    | NIL
    ;

listLiteral: LBRACK (expression (COMMA expression)*)? RBRACK;
tableLiteral: LBRACE (tableKeyValue (COMMA tableKeyValue)*)? RBRACE;
tableKeyValue: tableKey ASSIGN expression;
tableKey: expression | STRING | IDENTIFIER;

// Lexer Rules (unchanged)...

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
POW: '^^';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
IDIV: '//' ;
MOD: '%';
BITAND: '&';
BITOR: '|';
BITXOR: '^';
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
ADD_ASSIGN: '+=';
SUB_ASSIGN: '-=';
MUL_ASSIGN: '*=';
DIV_ASSIGN: '/=';
POW_ASSIGN: '^^=';
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
    : '"' ( ESC_SEQ | ~["\\\r\n] )* '"'
    | '\'' ( ESC_SEQ | ~['\\\r\n] )* '\''
    | '"""' .*? '"""'
    | '\'\'\'' .*? '\'\'\''
    ;

fragment ESC_SEQ: '\\' [btnr"'\\];

COMMENT: '#' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

WS: [ \t\r\n]+ -> skip; 
