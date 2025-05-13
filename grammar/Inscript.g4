grammar Inscript;

// Parser rules
program             : statementList EOF;

statementList       : statement+;
statement           : simpleStmt | compoundStmt;

simpleStmt          : assignment
                    | exprStmt
                    | printStmt
                    | returnStmt
                    ;

assignment          : primary '=' expression;
exprStmt            : expression;
printStmt           : 'print' '(' expressionListOpt ')';
returnStmt          : 'return' expressionOpt;

compoundStmt        : ifStmt
                    | whileStmt
                    | forStmt
                    | functionDef
                    ;

ifStmt              : 'if' expression block elseifListOpt elseBlockOpt;
elseifListOpt       : /* empty */ | elseifList;
elseifList          : elseif+;
elseif              : 'elseif' expression block;
elseBlockOpt        : /* empty */ | 'else' block;

whileStmt           : 'while' expression block;
forStmt             : 'for' IDENTIFIER 'in' expression block;
functionDef         : 'function' '(' paramListOpt ')' block;

block               : '{' statementListOpt '}';
statementListOpt    : /* empty */ | statementList;

expressionOpt       : /* empty */ | expression;
expressionListOpt   : /* empty */ | expressionList;
expressionList      : expression (',' expression)*;

paramListOpt        : /* empty */ | paramList;
paramList           : IDENTIFIER (',' IDENTIFIER)*;

expression          : logicalOr;
logicalOr           : logicalAnd ('or' logicalAnd)*;
logicalAnd          : comparison ('and' comparison)*;
comparison          : arith ( cmpOp=('==' | '!=' | '<' | '>' | '<=' | '>=') arith )*;
arith               : term ( op=('+' | '-') term )*;
term                : factor (('*' | '/' | '%') factor)*;
factor              : unary ('^' unary)*;
unary               : ('+' | '-' | 'not') unary
                    | primary;

primary             : atom ( '[' expression ']' | '(' expressionListOpt ')' )*;

atom                : literal
                    | IDENTIFIER
                    | listLiteral
                    | tableLiteral
                    | '(' expression ')'
                    | functionDef
                    ;

listLiteral         : '[' expressionListOpt ']';
tableLiteral        : '{' fieldListOpt '}';
fieldListOpt        : /* empty */ | fieldList;
fieldList           : field (',' field)*;
field               : IDENTIFIER '=' expression;

literal             : INTEGER
                    | FLOAT
                    | STRING
                    | BOOLEAN // Moved BOOLEAN here
                    | 'nil'
                    ;

// Lexer rules
// IMPORTANT: Order matters! BOOLEAN must come before IDENTIFIER
BOOLEAN             : 'true' | 'false'; // Defined before IDENTIFIER
IDENTIFIER          : LETTER (LETTER | DIGIT | '_')*;

INTEGER             : DIGIT+;
FLOAT               : DIGIT+ '.' DIGIT+;
STRING              : '"' (ESC_SEQ | ~["\\])* '"';


// Fragments (These are correct as helper rules for lexer rules)
fragment DIGIT      : [0-9];
fragment LETTER     : [a-zA-Z_];
fragment ESC_SEQ    : '\\' ["'\\ntbr];

// Skip whitespace and comments
WS                  : [ \t\r\n]+ -> skip;
LINE_COMMENT        : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT       : '/*' .*? '*/' -> skip;
