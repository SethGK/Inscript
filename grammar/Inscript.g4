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

// Function definitions
functionDef         : 'function' '(' paramListOpt ')' block;

// Block for statements
block               : '{' statementListOpt '}';
statementListOpt    : /* empty */ | statementList;

// Optional expression
expressionOpt       : /* empty */ | expression;

// Expression and parameters
expression          : logicalOr;
logicalOr           : logicalAnd ('or' logicalAnd)*;
logicalAnd          : comparison ('and' comparison)*;
comparison          : arith (('==' | '!=' | '<' | '>' | '<=' | '>=') arith)*;
arith               : term (('+' | '-') term)*;
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
                    | fnLiteral
                    ;

// Function literal (expression)
fnLiteral           : 'function' '(' paramListOpt ')' block;

// Literals and collections
listLiteral         : '[' expressionListOpt ']';
tableLiteral        : '{' fieldListOpt '}';
fieldListOpt        : /* empty */ | fieldList;
fieldList           : field (',' field)*;
field               : IDENTIFIER '=' expression;

literal             : INTEGER
                    | FLOAT
                    | STRING
                    | BOOLEAN
                    | 'nil'
                    ;

paramListOpt        : /* empty */ | paramList;
paramList           : IDENTIFIER (',' IDENTIFIER)*;

expressionListOpt   : /* empty */ | expressionList;
expressionList      : expression (',' expression)*;

// Lexer rules
BOOLEAN             : 'true' | 'false';
IDENTIFIER          : LETTER (LETTER | DIGIT | '_')*;
INTEGER             : DIGIT+;
FLOAT               : DIGIT+ '.' DIGIT+;
STRING              : '"' (ESC_SEQ | ~["\\])* '"';

fragment DIGIT      : [0-9];
fragment LETTER     : [a-zA-Z_];
fragment ESC_SEQ    : '\\' ["'\\ntbr];

WS                  : [ \t\r\n]+ -> skip;
LINE_COMMENT        : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT       : '/*' .*? '*/' -> skip;
