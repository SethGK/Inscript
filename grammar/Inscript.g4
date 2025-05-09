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
forStmt             : 'for' IDENTIFIER 'in' expression block; // Changed: identifier -> IDENTIFIER
functionDef         : 'function' '(' paramListOpt ')' block;

block               : '{' statementListOpt '}';
statementListOpt    : /* empty */ | statementList;

expressionOpt       : /* empty */ | expression;
expressionListOpt   : /* empty */ | expressionList;
expressionList      : expression (',' expression)*;

paramListOpt        : /* empty */ | paramList;
paramList           : IDENTIFIER (',' IDENTIFIER)*; // Changed: identifier -> IDENTIFIER

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
                    | IDENTIFIER // Changed: identifier -> IDENTIFIER
                    | listLiteral
                    | tableLiteral
                    | '(' expression ')'
                    ;

listLiteral         : '[' expressionListOpt ']';
tableLiteral        : '{' fieldListOpt '}';
fieldListOpt        : /* empty */ | fieldList;
fieldList           : field (',' field)*;
field               : IDENTIFIER '=' expression; // Changed: identifier -> IDENTIFIER

literal             : INTEGER
                    | FLOAT
                    | STRING
                    | BOOLEAN
                    | 'nil'
                    ;

// Lexer rules
IDENTIFIER          : LETTER (LETTER | DIGIT | '_')*; // Changed: identifier -> IDENTIFIER (now a lexer rule)

INTEGER             : DIGIT+;
FLOAT               : DIGIT+ '.' DIGIT+;
STRING              : '"' (ESC_SEQ | ~["\\])* '"';
BOOLEAN             : 'true' | 'false';


// Fragments (These are correct as helper rules for lexer rules like IDENTIFIER, INTEGER, FLOAT)
fragment DIGIT      : [0-9];
fragment LETTER     : [a-zA-Z_];
fragment ESC_SEQ    : '\\' ["'\\ntbr];

// Skip whitespace and comments
WS                  : [ \t\r\n]+ -> skip;
LINE_COMMENT        : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT       : '/*' .*? '*/' -> skip;