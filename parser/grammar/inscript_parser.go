// Code generated from grammar/Inscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Inscript

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type InscriptParser struct {
	*antlr.BaseParser
}

var InscriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func inscriptParserInit() {
	staticData := &InscriptParserStaticData
	staticData.LiteralNames = []string{
		"", "'function'", "'if'", "'else'", "'while'", "'for'", "'in'", "'break'",
		"'continue'", "'return'", "'import'", "'print'", "'true'", "'false'",
		"'nil'", "'and'", "'or'", "'not'", "'^^'", "'+'", "'-'", "'*'", "'/'",
		"'//'", "'%'", "'&'", "'|'", "'^'", "'~'", "'<<'", "'>>'", "'=='", "'!='",
		"'<'", "'<='", "'>'", "'>='", "'='", "'+='", "'-='", "'*='", "'/='",
		"'^^='", "'->'", "'('", "')'", "'['", "']'", "'{'", "'}'", "','", "'.'",
		"':'", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "FUNCTION", "IF", "ELSE", "WHILE", "FOR", "IN", "BREAK", "CONTINUE",
		"RETURN", "IMPORT", "PRINT", "TRUE", "FALSE", "NIL", "AND", "OR", "NOT",
		"POW", "ADD", "SUB", "MUL", "DIV", "IDIV", "MOD", "BITAND", "BITOR",
		"BITXOR", "BITNOT", "SHL", "SHR", "EQ", "NEQ", "LT", "LE", "GT", "GE",
		"ASSIGN", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "POW_ASSIGN",
		"ARROW", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
		"COMMA", "DOT", "COLON", "ELLIPSIS", "IDENTIFIER", "NUMBER", "STRING",
		"COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "block", "exprStmt", "assignment", "target",
		"ifStmt", "whileStmt", "forStmt", "funcDef", "paramList", "param", "typeAnnotation",
		"breakStmt", "continueStmt", "returnStmt", "importStmt", "printStmt",
		"expression", "unaryExpr", "postfixExpr", "argList", "primary", "literal",
		"listLiteral", "tableLiteral", "tableKeyValue", "tableKey",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 351, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0, 61, 9, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 77, 8, 1, 1, 2, 1, 2, 5, 2, 81, 8, 2, 10, 2, 12, 2, 84, 9,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 104, 8, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 111, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 127, 8, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 132, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 139,
		8, 10, 10, 10, 12, 10, 142, 9, 10, 1, 10, 3, 10, 145, 8, 10, 1, 11, 1,
		11, 1, 11, 3, 11, 150, 8, 11, 1, 11, 1, 11, 3, 11, 154, 8, 11, 1, 11, 1,
		11, 3, 11, 158, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 3, 15, 168, 8, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 5, 17, 178, 8, 17, 10, 17, 12, 17, 181, 9, 17, 3, 17, 183, 8,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 5, 18, 250, 8, 18, 10, 18, 12, 18, 253, 9, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 262, 8, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 270, 8, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 281, 8, 20, 10, 20, 12,
		20, 284, 9, 20, 1, 21, 1, 21, 1, 21, 5, 21, 289, 8, 21, 10, 21, 12, 21,
		292, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 4, 22, 304, 8, 22, 11, 22, 12, 22, 305, 1, 22, 1, 22, 1, 22,
		1, 22, 3, 22, 312, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5,
		24, 320, 8, 24, 10, 24, 12, 24, 323, 9, 24, 3, 24, 325, 8, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 333, 8, 25, 10, 25, 12, 25, 336,
		9, 25, 3, 25, 338, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 3, 27, 349, 8, 27, 1, 27, 0, 2, 36, 40, 28, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 0, 2, 1, 0, 37, 42, 2, 0, 12, 14, 55, 56, 388,
		0, 59, 1, 0, 0, 0, 2, 76, 1, 0, 0, 0, 4, 78, 1, 0, 0, 0, 6, 87, 1, 0, 0,
		0, 8, 89, 1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 105, 1, 0, 0, 0, 14, 112,
		1, 0, 0, 0, 16, 116, 1, 0, 0, 0, 18, 122, 1, 0, 0, 0, 20, 135, 1, 0, 0,
		0, 22, 157, 1, 0, 0, 0, 24, 159, 1, 0, 0, 0, 26, 161, 1, 0, 0, 0, 28, 163,
		1, 0, 0, 0, 30, 165, 1, 0, 0, 0, 32, 169, 1, 0, 0, 0, 34, 172, 1, 0, 0,
		0, 36, 186, 1, 0, 0, 0, 38, 261, 1, 0, 0, 0, 40, 263, 1, 0, 0, 0, 42, 285,
		1, 0, 0, 0, 44, 311, 1, 0, 0, 0, 46, 313, 1, 0, 0, 0, 48, 315, 1, 0, 0,
		0, 50, 328, 1, 0, 0, 0, 52, 341, 1, 0, 0, 0, 54, 348, 1, 0, 0, 0, 56, 58,
		3, 2, 1, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5,
		0, 0, 1, 63, 1, 1, 0, 0, 0, 64, 77, 3, 6, 3, 0, 65, 77, 3, 8, 4, 0, 66,
		77, 3, 12, 6, 0, 67, 77, 3, 14, 7, 0, 68, 77, 3, 16, 8, 0, 69, 77, 3, 18,
		9, 0, 70, 77, 3, 26, 13, 0, 71, 77, 3, 28, 14, 0, 72, 77, 3, 30, 15, 0,
		73, 77, 3, 32, 16, 0, 74, 77, 3, 34, 17, 0, 75, 77, 3, 4, 2, 0, 76, 64,
		1, 0, 0, 0, 76, 65, 1, 0, 0, 0, 76, 66, 1, 0, 0, 0, 76, 67, 1, 0, 0, 0,
		76, 68, 1, 0, 0, 0, 76, 69, 1, 0, 0, 0, 76, 70, 1, 0, 0, 0, 76, 71, 1,
		0, 0, 0, 76, 72, 1, 0, 0, 0, 76, 73, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76,
		75, 1, 0, 0, 0, 77, 3, 1, 0, 0, 0, 78, 82, 5, 48, 0, 0, 79, 81, 3, 2, 1,
		0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83,
		1, 0, 0, 0, 83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 49, 0, 0,
		86, 5, 1, 0, 0, 0, 87, 88, 3, 36, 18, 0, 88, 7, 1, 0, 0, 0, 89, 90, 3,
		10, 5, 0, 90, 91, 7, 0, 0, 0, 91, 92, 3, 36, 18, 0, 92, 9, 1, 0, 0, 0,
		93, 104, 5, 54, 0, 0, 94, 95, 3, 40, 20, 0, 95, 96, 5, 46, 0, 0, 96, 97,
		3, 36, 18, 0, 97, 98, 5, 47, 0, 0, 98, 104, 1, 0, 0, 0, 99, 100, 3, 40,
		20, 0, 100, 101, 5, 51, 0, 0, 101, 102, 5, 54, 0, 0, 102, 104, 1, 0, 0,
		0, 103, 93, 1, 0, 0, 0, 103, 94, 1, 0, 0, 0, 103, 99, 1, 0, 0, 0, 104,
		11, 1, 0, 0, 0, 105, 106, 5, 2, 0, 0, 106, 107, 3, 36, 18, 0, 107, 110,
		3, 4, 2, 0, 108, 109, 5, 3, 0, 0, 109, 111, 3, 4, 2, 0, 110, 108, 1, 0,
		0, 0, 110, 111, 1, 0, 0, 0, 111, 13, 1, 0, 0, 0, 112, 113, 5, 4, 0, 0,
		113, 114, 3, 36, 18, 0, 114, 115, 3, 4, 2, 0, 115, 15, 1, 0, 0, 0, 116,
		117, 5, 5, 0, 0, 117, 118, 5, 54, 0, 0, 118, 119, 5, 6, 0, 0, 119, 120,
		3, 36, 18, 0, 120, 121, 3, 4, 2, 0, 121, 17, 1, 0, 0, 0, 122, 123, 5, 1,
		0, 0, 123, 124, 5, 54, 0, 0, 124, 126, 5, 44, 0, 0, 125, 127, 3, 20, 10,
		0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		131, 5, 45, 0, 0, 129, 130, 5, 43, 0, 0, 130, 132, 3, 24, 12, 0, 131, 129,
		1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 3, 4,
		2, 0, 134, 19, 1, 0, 0, 0, 135, 140, 3, 22, 11, 0, 136, 137, 5, 50, 0,
		0, 137, 139, 3, 22, 11, 0, 138, 136, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0,
		140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142,
		140, 1, 0, 0, 0, 143, 145, 5, 50, 0, 0, 144, 143, 1, 0, 0, 0, 144, 145,
		1, 0, 0, 0, 145, 21, 1, 0, 0, 0, 146, 149, 5, 54, 0, 0, 147, 148, 5, 37,
		0, 0, 148, 150, 3, 36, 18, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0,
		0, 150, 153, 1, 0, 0, 0, 151, 152, 5, 52, 0, 0, 152, 154, 3, 24, 12, 0,
		153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 158, 1, 0, 0, 0, 155,
		156, 5, 53, 0, 0, 156, 158, 5, 54, 0, 0, 157, 146, 1, 0, 0, 0, 157, 155,
		1, 0, 0, 0, 158, 23, 1, 0, 0, 0, 159, 160, 5, 54, 0, 0, 160, 25, 1, 0,
		0, 0, 161, 162, 5, 7, 0, 0, 162, 27, 1, 0, 0, 0, 163, 164, 5, 8, 0, 0,
		164, 29, 1, 0, 0, 0, 165, 167, 5, 9, 0, 0, 166, 168, 3, 36, 18, 0, 167,
		166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 31, 1, 0, 0, 0, 169, 170, 5,
		10, 0, 0, 170, 171, 5, 56, 0, 0, 171, 33, 1, 0, 0, 0, 172, 173, 5, 11,
		0, 0, 173, 182, 5, 44, 0, 0, 174, 179, 3, 36, 18, 0, 175, 176, 5, 50, 0,
		0, 176, 178, 3, 36, 18, 0, 177, 175, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 182, 174, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 185, 5, 45, 0, 0, 185, 35, 1, 0, 0, 0, 186, 187, 6, 18,
		-1, 0, 187, 188, 3, 38, 19, 0, 188, 251, 1, 0, 0, 0, 189, 190, 10, 20,
		0, 0, 190, 191, 5, 18, 0, 0, 191, 250, 3, 36, 18, 21, 192, 193, 10, 19,
		0, 0, 193, 194, 5, 21, 0, 0, 194, 250, 3, 36, 18, 20, 195, 196, 10, 18,
		0, 0, 196, 197, 5, 22, 0, 0, 197, 250, 3, 36, 18, 19, 198, 199, 10, 17,
		0, 0, 199, 200, 5, 23, 0, 0, 200, 250, 3, 36, 18, 18, 201, 202, 10, 16,
		0, 0, 202, 203, 5, 24, 0, 0, 203, 250, 3, 36, 18, 17, 204, 205, 10, 15,
		0, 0, 205, 206, 5, 19, 0, 0, 206, 250, 3, 36, 18, 16, 207, 208, 10, 14,
		0, 0, 208, 209, 5, 20, 0, 0, 209, 250, 3, 36, 18, 15, 210, 211, 10, 13,
		0, 0, 211, 212, 5, 25, 0, 0, 212, 250, 3, 36, 18, 14, 213, 214, 10, 12,
		0, 0, 214, 215, 5, 26, 0, 0, 215, 250, 3, 36, 18, 13, 216, 217, 10, 11,
		0, 0, 217, 218, 5, 27, 0, 0, 218, 250, 3, 36, 18, 12, 219, 220, 10, 10,
		0, 0, 220, 221, 5, 29, 0, 0, 221, 250, 3, 36, 18, 11, 222, 223, 10, 9,
		0, 0, 223, 224, 5, 30, 0, 0, 224, 250, 3, 36, 18, 10, 225, 226, 10, 8,
		0, 0, 226, 227, 5, 33, 0, 0, 227, 250, 3, 36, 18, 9, 228, 229, 10, 7, 0,
		0, 229, 230, 5, 34, 0, 0, 230, 250, 3, 36, 18, 8, 231, 232, 10, 6, 0, 0,
		232, 233, 5, 35, 0, 0, 233, 250, 3, 36, 18, 7, 234, 235, 10, 5, 0, 0, 235,
		236, 5, 36, 0, 0, 236, 250, 3, 36, 18, 6, 237, 238, 10, 4, 0, 0, 238, 239,
		5, 31, 0, 0, 239, 250, 3, 36, 18, 5, 240, 241, 10, 3, 0, 0, 241, 242, 5,
		32, 0, 0, 242, 250, 3, 36, 18, 4, 243, 244, 10, 2, 0, 0, 244, 245, 5, 15,
		0, 0, 245, 250, 3, 36, 18, 3, 246, 247, 10, 1, 0, 0, 247, 248, 5, 16, 0,
		0, 248, 250, 3, 36, 18, 2, 249, 189, 1, 0, 0, 0, 249, 192, 1, 0, 0, 0,
		249, 195, 1, 0, 0, 0, 249, 198, 1, 0, 0, 0, 249, 201, 1, 0, 0, 0, 249,
		204, 1, 0, 0, 0, 249, 207, 1, 0, 0, 0, 249, 210, 1, 0, 0, 0, 249, 213,
		1, 0, 0, 0, 249, 216, 1, 0, 0, 0, 249, 219, 1, 0, 0, 0, 249, 222, 1, 0,
		0, 0, 249, 225, 1, 0, 0, 0, 249, 228, 1, 0, 0, 0, 249, 231, 1, 0, 0, 0,
		249, 234, 1, 0, 0, 0, 249, 237, 1, 0, 0, 0, 249, 240, 1, 0, 0, 0, 249,
		243, 1, 0, 0, 0, 249, 246, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249,
		1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 37, 1, 0, 0, 0, 253, 251, 1, 0,
		0, 0, 254, 255, 5, 17, 0, 0, 255, 262, 3, 38, 19, 0, 256, 257, 5, 28, 0,
		0, 257, 262, 3, 38, 19, 0, 258, 259, 5, 20, 0, 0, 259, 262, 3, 38, 19,
		0, 260, 262, 3, 40, 20, 0, 261, 254, 1, 0, 0, 0, 261, 256, 1, 0, 0, 0,
		261, 258, 1, 0, 0, 0, 261, 260, 1, 0, 0, 0, 262, 39, 1, 0, 0, 0, 263, 264,
		6, 20, -1, 0, 264, 265, 3, 44, 22, 0, 265, 282, 1, 0, 0, 0, 266, 267, 10,
		3, 0, 0, 267, 269, 5, 44, 0, 0, 268, 270, 3, 42, 21, 0, 269, 268, 1, 0,
		0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 281, 5, 45, 0, 0,
		272, 273, 10, 2, 0, 0, 273, 274, 5, 46, 0, 0, 274, 275, 3, 36, 18, 0, 275,
		276, 5, 47, 0, 0, 276, 281, 1, 0, 0, 0, 277, 278, 10, 1, 0, 0, 278, 279,
		5, 51, 0, 0, 279, 281, 5, 54, 0, 0, 280, 266, 1, 0, 0, 0, 280, 272, 1,
		0, 0, 0, 280, 277, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0, 282, 280, 1, 0, 0,
		0, 282, 283, 1, 0, 0, 0, 283, 41, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285,
		290, 3, 36, 18, 0, 286, 287, 5, 50, 0, 0, 287, 289, 3, 36, 18, 0, 288,
		286, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291,
		1, 0, 0, 0, 291, 43, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 312, 3, 46,
		23, 0, 294, 312, 5, 54, 0, 0, 295, 296, 5, 44, 0, 0, 296, 297, 3, 36, 18,
		0, 297, 298, 5, 45, 0, 0, 298, 312, 1, 0, 0, 0, 299, 300, 5, 44, 0, 0,
		300, 303, 3, 36, 18, 0, 301, 302, 5, 50, 0, 0, 302, 304, 3, 36, 18, 0,
		303, 301, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305,
		306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 5, 45, 0, 0, 308, 312,
		1, 0, 0, 0, 309, 312, 3, 48, 24, 0, 310, 312, 3, 50, 25, 0, 311, 293, 1,
		0, 0, 0, 311, 294, 1, 0, 0, 0, 311, 295, 1, 0, 0, 0, 311, 299, 1, 0, 0,
		0, 311, 309, 1, 0, 0, 0, 311, 310, 1, 0, 0, 0, 312, 45, 1, 0, 0, 0, 313,
		314, 7, 1, 0, 0, 314, 47, 1, 0, 0, 0, 315, 324, 5, 46, 0, 0, 316, 321,
		3, 36, 18, 0, 317, 318, 5, 50, 0, 0, 318, 320, 3, 36, 18, 0, 319, 317,
		1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 322, 1, 0,
		0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 316, 1, 0, 0, 0,
		324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 5, 47, 0, 0, 327,
		49, 1, 0, 0, 0, 328, 337, 5, 48, 0, 0, 329, 334, 3, 52, 26, 0, 330, 331,
		5, 50, 0, 0, 331, 333, 3, 52, 26, 0, 332, 330, 1, 0, 0, 0, 333, 336, 1,
		0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 338, 1, 0, 0,
		0, 336, 334, 1, 0, 0, 0, 337, 329, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338,
		339, 1, 0, 0, 0, 339, 340, 5, 49, 0, 0, 340, 51, 1, 0, 0, 0, 341, 342,
		3, 54, 27, 0, 342, 343, 5, 37, 0, 0, 343, 344, 3, 36, 18, 0, 344, 53, 1,
		0, 0, 0, 345, 349, 3, 36, 18, 0, 346, 349, 5, 56, 0, 0, 347, 349, 5, 54,
		0, 0, 348, 345, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 347, 1, 0, 0, 0,
		349, 55, 1, 0, 0, 0, 29, 59, 76, 82, 103, 110, 126, 131, 140, 144, 149,
		153, 157, 167, 179, 182, 249, 251, 261, 269, 280, 282, 290, 305, 311, 321,
		324, 334, 337, 348,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// InscriptParserInit initializes any static state used to implement InscriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInscriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InscriptParserInit() {
	staticData := &InscriptParserStaticData
	staticData.once.Do(inscriptParserInit)
}

// NewInscriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInscriptParser(input antlr.TokenStream) *InscriptParser {
	InscriptParserInit()
	this := new(InscriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &InscriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Inscript.g4"

	return this
}

// InscriptParser tokens.
const (
	InscriptParserEOF           = antlr.TokenEOF
	InscriptParserFUNCTION      = 1
	InscriptParserIF            = 2
	InscriptParserELSE          = 3
	InscriptParserWHILE         = 4
	InscriptParserFOR           = 5
	InscriptParserIN            = 6
	InscriptParserBREAK         = 7
	InscriptParserCONTINUE      = 8
	InscriptParserRETURN        = 9
	InscriptParserIMPORT        = 10
	InscriptParserPRINT         = 11
	InscriptParserTRUE          = 12
	InscriptParserFALSE         = 13
	InscriptParserNIL           = 14
	InscriptParserAND           = 15
	InscriptParserOR            = 16
	InscriptParserNOT           = 17
	InscriptParserPOW           = 18
	InscriptParserADD           = 19
	InscriptParserSUB           = 20
	InscriptParserMUL           = 21
	InscriptParserDIV           = 22
	InscriptParserIDIV          = 23
	InscriptParserMOD           = 24
	InscriptParserBITAND        = 25
	InscriptParserBITOR         = 26
	InscriptParserBITXOR        = 27
	InscriptParserBITNOT        = 28
	InscriptParserSHL           = 29
	InscriptParserSHR           = 30
	InscriptParserEQ            = 31
	InscriptParserNEQ           = 32
	InscriptParserLT            = 33
	InscriptParserLE            = 34
	InscriptParserGT            = 35
	InscriptParserGE            = 36
	InscriptParserASSIGN        = 37
	InscriptParserADD_ASSIGN    = 38
	InscriptParserSUB_ASSIGN    = 39
	InscriptParserMUL_ASSIGN    = 40
	InscriptParserDIV_ASSIGN    = 41
	InscriptParserPOW_ASSIGN    = 42
	InscriptParserARROW         = 43
	InscriptParserLPAREN        = 44
	InscriptParserRPAREN        = 45
	InscriptParserLBRACK        = 46
	InscriptParserRBRACK        = 47
	InscriptParserLBRACE        = 48
	InscriptParserRBRACE        = 49
	InscriptParserCOMMA         = 50
	InscriptParserDOT           = 51
	InscriptParserCOLON         = 52
	InscriptParserELLIPSIS      = 53
	InscriptParserIDENTIFIER    = 54
	InscriptParserNUMBER        = 55
	InscriptParserSTRING        = 56
	InscriptParserCOMMENT       = 57
	InscriptParserBLOCK_COMMENT = 58
	InscriptParserWS            = 59
)

// InscriptParser rules.
const (
	InscriptParserRULE_program        = 0
	InscriptParserRULE_statement      = 1
	InscriptParserRULE_block          = 2
	InscriptParserRULE_exprStmt       = 3
	InscriptParserRULE_assignment     = 4
	InscriptParserRULE_target         = 5
	InscriptParserRULE_ifStmt         = 6
	InscriptParserRULE_whileStmt      = 7
	InscriptParserRULE_forStmt        = 8
	InscriptParserRULE_funcDef        = 9
	InscriptParserRULE_paramList      = 10
	InscriptParserRULE_param          = 11
	InscriptParserRULE_typeAnnotation = 12
	InscriptParserRULE_breakStmt      = 13
	InscriptParserRULE_continueStmt   = 14
	InscriptParserRULE_returnStmt     = 15
	InscriptParserRULE_importStmt     = 16
	InscriptParserRULE_printStmt      = 17
	InscriptParserRULE_expression     = 18
	InscriptParserRULE_unaryExpr      = 19
	InscriptParserRULE_postfixExpr    = 20
	InscriptParserRULE_argList        = 21
	InscriptParserRULE_primary        = 22
	InscriptParserRULE_literal        = 23
	InscriptParserRULE_listLiteral    = 24
	InscriptParserRULE_tableLiteral   = 25
	InscriptParserRULE_tableKeyValue  = 26
	InscriptParserRULE_tableKey       = 27
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(InscriptParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InscriptParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126470225742954422) != 0 {
		{
			p.SetState(56)
			p.Statement()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(InscriptParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprStmt() IExprStmtContext
	Assignment() IAssignmentContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	ForStmt() IForStmtContext
	FuncDef() IFuncDefContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	ReturnStmt() IReturnStmtContext
	ImportStmt() IImportStmtContext
	PrintStmt() IPrintStmtContext
	Block() IBlockContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) FuncDef() IFuncDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) ImportStmt() IImportStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *StatementContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InscriptParserRULE_statement)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.ExprStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.IfStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.WhileStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.ForStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.FuncDef()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.BreakStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(71)
			p.ContinueStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(72)
			p.ReturnStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(73)
			p.ImportStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(74)
			p.PrintStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(75)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(InscriptParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(InscriptParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InscriptParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(InscriptParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126470225742954422) != 0 {
		{
			p.SetState(79)
			p.Statement()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(InscriptParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InscriptParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Target() ITargetContext
	Expression() IExpressionContext
	ASSIGN() antlr.TerminalNode
	ADD_ASSIGN() antlr.TerminalNode
	SUB_ASSIGN() antlr.TerminalNode
	MUL_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	POW_ASSIGN() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Target() ITargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserASSIGN, 0)
}

func (s *AssignmentContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserADD_ASSIGN, 0)
}

func (s *AssignmentContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserSUB_ASSIGN, 0)
}

func (s *AssignmentContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserMUL_ASSIGN, 0)
}

func (s *AssignmentContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserDIV_ASSIGN, 0)
}

func (s *AssignmentContext) POW_ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserPOW_ASSIGN, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InscriptParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Target()
	}
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8658654068736) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(91)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	PostfixExpr() IPostfixExprContext
	LBRACK() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACK() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_target
	return p
}

func InitEmptyTargetContext(p *TargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_target
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *TargetContext) PostfixExpr() IPostfixExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExprContext)
}

func (s *TargetContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(InscriptParserLBRACK, 0)
}

func (s *TargetContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TargetContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(InscriptParserRBRACK, 0)
}

func (s *TargetContext) DOT() antlr.TerminalNode {
	return s.GetToken(InscriptParserDOT, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (s *TargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitTarget(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, InscriptParserRULE_target)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.postfixExpr(0)
		}
		{
			p.SetState(95)
			p.Match(InscriptParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.expression(0)
		}
		{
			p.SetState(97)
			p.Match(InscriptParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.postfixExpr(0)
		}
		{
			p.SetState(100)
			p.Match(InscriptParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(InscriptParserIF, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(InscriptParserELSE, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, InscriptParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(InscriptParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.expression(0)
	}
	{
		p.SetState(107)
		p.Block()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InscriptParserELSE {
		{
			p.SetState(108)
			p.Match(InscriptParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(InscriptParserWHILE, 0)
}

func (s *WhileStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, InscriptParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(InscriptParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.expression(0)
	}
	{
		p.SetState(114)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(InscriptParserFOR, 0)
}

func (s *ForStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *ForStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(InscriptParserIN, 0)
}

func (s *ForStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, InscriptParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(InscriptParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(InscriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(InscriptParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.expression(0)
	}
	{
		p.SetState(120)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ParamList() IParamListContext
	ARROW() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_funcDef
	return p
}

func InitEmptyFuncDefContext(p *FuncDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_funcDef
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(InscriptParserFUNCTION, 0)
}

func (s *FuncDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *FuncDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserLPAREN, 0)
}

func (s *FuncDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserRPAREN, 0)
}

func (s *FuncDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDefContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FuncDefContext) ARROW() antlr.TerminalNode {
	return s.GetToken(InscriptParserARROW, 0)
}

func (s *FuncDefContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (s *FuncDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitFuncDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, InscriptParserRULE_funcDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(InscriptParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(InscriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(InscriptParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InscriptParserELLIPSIS || _la == InscriptParserIDENTIFIER {
		{
			p.SetState(125)
			p.ParamList()
		}

	}
	{
		p.SetState(128)
		p.Match(InscriptParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InscriptParserARROW {
		{
			p.SetState(129)
			p.Match(InscriptParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.TypeAnnotation()
		}

	}
	{
		p.SetState(133)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserCOMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, InscriptParserRULE_paramList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Param()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(136)
				p.Match(InscriptParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(137)
				p.Param()
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InscriptParserCOMMA {
		{
			p.SetState(143)
			p.Match(InscriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	ELLIPSIS() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *ParamContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserASSIGN, 0)
}

func (s *ParamContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(InscriptParserCOLON, 0)
}

func (s *ParamContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *ParamContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(InscriptParserELLIPSIS, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, InscriptParserRULE_param)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == InscriptParserASSIGN {
			{
				p.SetState(147)
				p.Match(InscriptParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(148)
				p.expression(0)
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == InscriptParserCOLON {
			{
				p.SetState(151)
				p.Match(InscriptParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(152)
				p.TypeAnnotation()
			}

		}

	case InscriptParserELLIPSIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(InscriptParserELLIPSIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_typeAnnotation
	return p
}

func InitEmptyTypeAnnotationContext(p *TypeAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_typeAnnotation
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitTypeAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, InscriptParserRULE_typeAnnotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(InscriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(InscriptParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, InscriptParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(InscriptParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(InscriptParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, InscriptParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(InscriptParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(InscriptParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, InscriptParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(InscriptParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(166)
			p.expression(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_importStmt
	return p
}

func InitEmptyImportStmtContext(p *ImportStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_importStmt
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(InscriptParserIMPORT, 0)
}

func (s *ImportStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(InscriptParserSTRING, 0)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ImportStmt() (localctx IImportStmtContext) {
	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, InscriptParserRULE_importStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(InscriptParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(InscriptParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(InscriptParserPRINT, 0)
}

func (s *PrintStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserLPAREN, 0)
}

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserRPAREN, 0)
}

func (s *PrintStmtContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrintStmtContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserCOMMA)
}

func (s *PrintStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserCOMMA, i)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, InscriptParserRULE_printStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(InscriptParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(InscriptParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126470225742950400) != 0 {
		{
			p.SetState(174)
			p.expression(0)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == InscriptParserCOMMA {
			{
				p.SetState(175)
				p.Match(InscriptParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(176)
				p.expression(0)
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(184)
		p.Match(InscriptParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GeExprContext struct {
	ExpressionContext
}

func NewGeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GeExprContext {
	var p = new(GeExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *GeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GeExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GeExprContext) GE() antlr.TerminalNode {
	return s.GetToken(InscriptParserGE, 0)
}

func (s *GeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterGeExpr(s)
	}
}

func (s *GeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitGeExpr(s)
	}
}

func (s *GeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitGeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModExprContext struct {
	ExpressionContext
}

func NewModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExprContext {
	var p = new(ModExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ModExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(InscriptParserMOD, 0)
}

func (s *ModExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterModExpr(s)
	}
}

func (s *ModExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitModExpr(s)
	}
}

func (s *ModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitModExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GtExprContext struct {
	ExpressionContext
}

func NewGtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GtExprContext {
	var p = new(GtExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *GtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GtExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GtExprContext) GT() antlr.TerminalNode {
	return s.GetToken(InscriptParserGT, 0)
}

func (s *GtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterGtExpr(s)
	}
}

func (s *GtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitGtExpr(s)
	}
}

func (s *GtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitGtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(InscriptParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdivExprContext struct {
	ExpressionContext
}

func NewIdivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdivExprContext {
	var p = new(IdivExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdivExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IdivExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdivExprContext) IDIV() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDIV, 0)
}

func (s *IdivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterIdivExpr(s)
	}
}

func (s *IdivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitIdivExpr(s)
	}
}

func (s *IdivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitIdivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubExprContext struct {
	ExpressionContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SubExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(InscriptParserSUB, 0)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

func (s *SubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NeqExprContext struct {
	ExpressionContext
}

func NewNeqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NeqExprContext {
	var p = new(NeqExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NeqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeqExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *NeqExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NeqExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(InscriptParserNEQ, 0)
}

func (s *NeqExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterNeqExpr(s)
	}
}

func (s *NeqExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitNeqExpr(s)
	}
}

func (s *NeqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitNeqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LtExprContext struct {
	ExpressionContext
}

func NewLtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LtExprContext {
	var p = new(LtExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LtExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LtExprContext) LT() antlr.TerminalNode {
	return s.GetToken(InscriptParserLT, 0)
}

func (s *LtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterLtExpr(s)
	}
}

func (s *LtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitLtExpr(s)
	}
}

func (s *LtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitLtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqExprContext struct {
	ExpressionContext
}

func NewEqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExprContext {
	var p = new(EqExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(InscriptParserEQ, 0)
}

func (s *EqExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterEqExpr(s)
	}
}

func (s *EqExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitEqExpr(s)
	}
}

func (s *EqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitEqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpExprContext struct {
	ExpressionContext
}

func NewExpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpExprContext {
	var p = new(ExpExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpExprContext) POW() antlr.TerminalNode {
	return s.GetToken(InscriptParserPOW, 0)
}

func (s *ExpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterExpExpr(s)
	}
}

func (s *ExpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitExpExpr(s)
	}
}

func (s *ExpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitExpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitandExprContext struct {
	ExpressionContext
}

func NewBitandExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitandExprContext {
	var p = new(BitandExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitandExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitandExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitandExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitandExprContext) BITAND() antlr.TerminalNode {
	return s.GetToken(InscriptParserBITAND, 0)
}

func (s *BitandExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterBitandExpr(s)
	}
}

func (s *BitandExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitBitandExpr(s)
	}
}

func (s *BitandExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitBitandExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExpressionContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(InscriptParserADD, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LeExprContext struct {
	ExpressionContext
}

func NewLeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LeExprContext {
	var p = new(LeExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LeExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LeExprContext) LE() antlr.TerminalNode {
	return s.GetToken(InscriptParserLE, 0)
}

func (s *LeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterLeExpr(s)
	}
}

func (s *LeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitLeExpr(s)
	}
}

func (s *LeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitLeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitorExprContext struct {
	ExpressionContext
}

func NewBitorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitorExprContext {
	var p = new(BitorExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitorExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitorExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitorExprContext) BITOR() antlr.TerminalNode {
	return s.GetToken(InscriptParserBITOR, 0)
}

func (s *BitorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterBitorExpr(s)
	}
}

func (s *BitorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitBitorExpr(s)
	}
}

func (s *BitorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitBitorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitxorExprContext struct {
	ExpressionContext
}

func NewBitxorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitxorExprContext {
	var p = new(BitxorExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitxorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitxorExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitxorExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitxorExprContext) BITXOR() antlr.TerminalNode {
	return s.GetToken(InscriptParserBITXOR, 0)
}

func (s *BitxorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterBitxorExpr(s)
	}
}

func (s *BitxorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitBitxorExpr(s)
	}
}

func (s *BitxorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitBitxorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExpressionContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(InscriptParserMUL, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivExprContext struct {
	ExpressionContext
}

func NewDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExprContext {
	var p = new(DivExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DivExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DivExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(InscriptParserDIV, 0)
}

func (s *DivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterDivExpr(s)
	}
}

func (s *DivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitDivExpr(s)
	}
}

func (s *DivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShlExprContext struct {
	ExpressionContext
}

func NewShlExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShlExprContext {
	var p = new(ShlExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ShlExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShlExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ShlExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShlExprContext) SHL() antlr.TerminalNode {
	return s.GetToken(InscriptParserSHL, 0)
}

func (s *ShlExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterShlExpr(s)
	}
}

func (s *ShlExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitShlExpr(s)
	}
}

func (s *ShlExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitShlExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShrExprContext struct {
	ExpressionContext
}

func NewShrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShrExprContext {
	var p = new(ShrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ShrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ShrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShrExprContext) SHR() antlr.TerminalNode {
	return s.GetToken(InscriptParserSHR, 0)
}

func (s *ShrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterShrExpr(s)
	}
}

func (s *ShrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitShrExpr(s)
	}
}

func (s *ShrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitShrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(InscriptParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *InscriptParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, InscriptParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewUnaryExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(187)
		p.UnaryExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(190)
					p.Match(InscriptParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(191)
					p.expression(21)
				}

			case 2:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(193)
					p.Match(InscriptParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(194)
					p.expression(20)
				}

			case 3:
				localctx = NewDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(196)
					p.Match(InscriptParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(197)
					p.expression(19)
				}

			case 4:
				localctx = NewIdivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(198)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(199)
					p.Match(InscriptParserIDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(200)
					p.expression(18)
				}

			case 5:
				localctx = NewModExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(202)
					p.Match(InscriptParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(203)
					p.expression(17)
				}

			case 6:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(205)
					p.Match(InscriptParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(206)
					p.expression(16)
				}

			case 7:
				localctx = NewSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(208)
					p.Match(InscriptParserSUB)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(209)
					p.expression(15)
				}

			case 8:
				localctx = NewBitandExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(211)
					p.Match(InscriptParserBITAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(212)
					p.expression(14)
				}

			case 9:
				localctx = NewBitorExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(214)
					p.Match(InscriptParserBITOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(215)
					p.expression(13)
				}

			case 10:
				localctx = NewBitxorExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(217)
					p.Match(InscriptParserBITXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(218)
					p.expression(12)
				}

			case 11:
				localctx = NewShlExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(220)
					p.Match(InscriptParserSHL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(221)
					p.expression(11)
				}

			case 12:
				localctx = NewShrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(223)
					p.Match(InscriptParserSHR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(224)
					p.expression(10)
				}

			case 13:
				localctx = NewLtExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(226)
					p.Match(InscriptParserLT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(227)
					p.expression(9)
				}

			case 14:
				localctx = NewLeExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(229)
					p.Match(InscriptParserLE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(230)
					p.expression(8)
				}

			case 15:
				localctx = NewGtExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(232)
					p.Match(InscriptParserGT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(233)
					p.expression(7)
				}

			case 16:
				localctx = NewGeExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(235)
					p.Match(InscriptParserGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(236)
					p.expression(6)
				}

			case 17:
				localctx = NewEqExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(238)
					p.Match(InscriptParserEQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(239)
					p.expression(5)
				}

			case 18:
				localctx = NewNeqExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(241)
					p.Match(InscriptParserNEQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(242)
					p.expression(4)
				}

			case 19:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(244)
					p.Match(InscriptParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(245)
					p.expression(3)
				}

			case 20:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_expression)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(247)
					p.Match(InscriptParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(248)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_unaryExpr
	return p
}

func InitEmptyUnaryExprContext(p *UnaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_unaryExpr
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) CopyAll(ctx *UnaryExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotExprContext struct {
	UnaryExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(InscriptParserNOT, 0)
}

func (s *NotExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegExprContext struct {
	UnaryExprContext
}

func NewNegExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExprContext {
	var p = new(NegExprContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *NegExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(InscriptParserSUB, 0)
}

func (s *NegExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *NegExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterNegExpr(s)
	}
}

func (s *NegExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitNegExpr(s)
	}
}

func (s *NegExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitNegExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitnotExprContext struct {
	UnaryExprContext
}

func NewBitnotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitnotExprContext {
	var p = new(BitnotExprContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *BitnotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitnotExprContext) BITNOT() antlr.TerminalNode {
	return s.GetToken(InscriptParserBITNOT, 0)
}

func (s *BitnotExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *BitnotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterBitnotExpr(s)
	}
}

func (s *BitnotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitBitnotExpr(s)
	}
}

func (s *BitnotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitBitnotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixExpressionContext struct {
	UnaryExprContext
}

func NewPostfixExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) PostfixExpr() IPostfixExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExprContext)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitPostfixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, InscriptParserRULE_unaryExpr)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Match(InscriptParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.UnaryExpr()
		}

	case InscriptParserBITNOT:
		localctx = NewBitnotExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.Match(InscriptParserBITNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.UnaryExpr()
		}

	case InscriptParserSUB:
		localctx = NewNegExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)
			p.Match(InscriptParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.UnaryExpr()
		}

	case InscriptParserTRUE, InscriptParserFALSE, InscriptParserNIL, InscriptParserLPAREN, InscriptParserLBRACK, InscriptParserLBRACE, InscriptParserIDENTIFIER, InscriptParserNUMBER, InscriptParserSTRING:
		localctx = NewPostfixExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(260)
			p.postfixExpr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixExprContext is an interface to support dynamic dispatch.
type IPostfixExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPostfixExprContext differentiates from other interfaces.
	IsPostfixExprContext()
}

type PostfixExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExprContext() *PostfixExprContext {
	var p = new(PostfixExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_postfixExpr
	return p
}

func InitEmptyPostfixExprContext(p *PostfixExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_postfixExpr
}

func (*PostfixExprContext) IsPostfixExprContext() {}

func NewPostfixExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExprContext {
	var p = new(PostfixExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_postfixExpr

	return p
}

func (s *PostfixExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExprContext) CopyAll(ctx *PostfixExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PostfixExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryPostfixContext struct {
	PostfixExprContext
}

func NewPrimaryPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryPostfixContext {
	var p = new(PrimaryPostfixContext)

	InitEmptyPostfixExprContext(&p.PostfixExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixExprContext))

	return p
}

func (s *PrimaryPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryPostfixContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryPostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterPrimaryPostfix(s)
	}
}

func (s *PrimaryPostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitPrimaryPostfix(s)
	}
}

func (s *PrimaryPostfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitPrimaryPostfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexPostfixContext struct {
	PostfixExprContext
}

func NewIndexPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexPostfixContext {
	var p = new(IndexPostfixContext)

	InitEmptyPostfixExprContext(&p.PostfixExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixExprContext))

	return p
}

func (s *IndexPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexPostfixContext) PostfixExpr() IPostfixExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExprContext)
}

func (s *IndexPostfixContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(InscriptParserLBRACK, 0)
}

func (s *IndexPostfixContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexPostfixContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(InscriptParserRBRACK, 0)
}

func (s *IndexPostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterIndexPostfix(s)
	}
}

func (s *IndexPostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitIndexPostfix(s)
	}
}

func (s *IndexPostfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitIndexPostfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type AttrPostfixContext struct {
	PostfixExprContext
}

func NewAttrPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttrPostfixContext {
	var p = new(AttrPostfixContext)

	InitEmptyPostfixExprContext(&p.PostfixExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixExprContext))

	return p
}

func (s *AttrPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPostfixContext) PostfixExpr() IPostfixExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExprContext)
}

func (s *AttrPostfixContext) DOT() antlr.TerminalNode {
	return s.GetToken(InscriptParserDOT, 0)
}

func (s *AttrPostfixContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *AttrPostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterAttrPostfix(s)
	}
}

func (s *AttrPostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitAttrPostfix(s)
	}
}

func (s *AttrPostfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitAttrPostfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallPostfixContext struct {
	PostfixExprContext
}

func NewCallPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallPostfixContext {
	var p = new(CallPostfixContext)

	InitEmptyPostfixExprContext(&p.PostfixExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixExprContext))

	return p
}

func (s *CallPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallPostfixContext) PostfixExpr() IPostfixExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExprContext)
}

func (s *CallPostfixContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserLPAREN, 0)
}

func (s *CallPostfixContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserRPAREN, 0)
}

func (s *CallPostfixContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *CallPostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterCallPostfix(s)
	}
}

func (s *CallPostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitCallPostfix(s)
	}
}

func (s *CallPostfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitCallPostfix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) PostfixExpr() (localctx IPostfixExprContext) {
	return p.postfixExpr(0)
}

func (p *InscriptParser) postfixExpr(_p int) (localctx IPostfixExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPostfixExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostfixExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, InscriptParserRULE_postfixExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPrimaryPostfixContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(264)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCallPostfixContext(p, NewPostfixExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_postfixExpr)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(267)
					p.Match(InscriptParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(269)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126470225742950400) != 0 {
					{
						p.SetState(268)
						p.ArgList()
					}

				}
				{
					p.SetState(271)
					p.Match(InscriptParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewIndexPostfixContext(p, NewPostfixExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_postfixExpr)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(273)
					p.Match(InscriptParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(274)
					p.expression(0)
				}
				{
					p.SetState(275)
					p.Match(InscriptParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewAttrPostfixContext(p, NewPostfixExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, InscriptParserRULE_postfixExpr)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(278)
					p.Match(InscriptParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(279)
					p.Match(InscriptParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, InscriptParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.expression(0)
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserCOMMA {
		{
			p.SetState(286)
			p.Match(InscriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.expression(0)
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ListLiteral() IListLiteralContext
	TableLiteral() ITableLiteralContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserLPAREN, 0)
}

func (s *PrimaryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(InscriptParserRPAREN, 0)
}

func (s *PrimaryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserCOMMA)
}

func (s *PrimaryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserCOMMA, i)
}

func (s *PrimaryContext) ListLiteral() IListLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListLiteralContext)
}

func (s *PrimaryContext) TableLiteral() ITableLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableLiteralContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, InscriptParserRULE_primary)
	var _la int

	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Match(InscriptParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.expression(0)
		}
		{
			p.SetState(297)
			p.Match(InscriptParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(299)
			p.Match(InscriptParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.expression(0)
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == InscriptParserCOMMA {
			{
				p.SetState(301)
				p.Match(InscriptParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(302)
				p.expression(0)
			}

			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(307)
			p.Match(InscriptParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(309)
			p.ListLiteral()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(310)
			p.TableLiteral()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NIL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InscriptParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(InscriptParserSTRING, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(InscriptParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(InscriptParserFALSE, 0)
}

func (s *LiteralContext) NIL() antlr.TerminalNode {
	return s.GetToken(InscriptParserNIL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, InscriptParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&108086391056920576) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListLiteralContext is an interface to support dynamic dispatch.
type IListLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListLiteralContext differentiates from other interfaces.
	IsListLiteralContext()
}

type ListLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListLiteralContext() *ListLiteralContext {
	var p = new(ListLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_listLiteral
	return p
}

func InitEmptyListLiteralContext(p *ListLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_listLiteral
}

func (*ListLiteralContext) IsListLiteralContext() {}

func NewListLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListLiteralContext {
	var p = new(ListLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_listLiteral

	return p
}

func (s *ListLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ListLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(InscriptParserLBRACK, 0)
}

func (s *ListLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(InscriptParserRBRACK, 0)
}

func (s *ListLiteralContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListLiteralContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserCOMMA)
}

func (s *ListLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserCOMMA, i)
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterListLiteral(s)
	}
}

func (s *ListLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitListLiteral(s)
	}
}

func (s *ListLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitListLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ListLiteral() (localctx IListLiteralContext) {
	localctx = NewListLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, InscriptParserRULE_listLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(InscriptParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126470225742950400) != 0 {
		{
			p.SetState(316)
			p.expression(0)
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == InscriptParserCOMMA {
			{
				p.SetState(317)
				p.Match(InscriptParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(318)
				p.expression(0)
			}

			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(326)
		p.Match(InscriptParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableLiteralContext is an interface to support dynamic dispatch.
type ITableLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllTableKeyValue() []ITableKeyValueContext
	TableKeyValue(i int) ITableKeyValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTableLiteralContext differentiates from other interfaces.
	IsTableLiteralContext()
}

type TableLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableLiteralContext() *TableLiteralContext {
	var p = new(TableLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_tableLiteral
	return p
}

func InitEmptyTableLiteralContext(p *TableLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_tableLiteral
}

func (*TableLiteralContext) IsTableLiteralContext() {}

func NewTableLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableLiteralContext {
	var p = new(TableLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_tableLiteral

	return p
}

func (s *TableLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TableLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(InscriptParserLBRACE, 0)
}

func (s *TableLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(InscriptParserRBRACE, 0)
}

func (s *TableLiteralContext) AllTableKeyValue() []ITableKeyValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableKeyValueContext); ok {
			len++
		}
	}

	tst := make([]ITableKeyValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableKeyValueContext); ok {
			tst[i] = t.(ITableKeyValueContext)
			i++
		}
	}

	return tst
}

func (s *TableLiteralContext) TableKeyValue(i int) ITableKeyValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableKeyValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableKeyValueContext)
}

func (s *TableLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserCOMMA)
}

func (s *TableLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserCOMMA, i)
}

func (s *TableLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterTableLiteral(s)
	}
}

func (s *TableLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitTableLiteral(s)
	}
}

func (s *TableLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitTableLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) TableLiteral() (localctx ITableLiteralContext) {
	localctx = NewTableLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, InscriptParserRULE_tableLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(InscriptParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126470225742950400) != 0 {
		{
			p.SetState(329)
			p.TableKeyValue()
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == InscriptParserCOMMA {
			{
				p.SetState(330)
				p.Match(InscriptParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(331)
				p.TableKeyValue()
			}

			p.SetState(336)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(339)
		p.Match(InscriptParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableKeyValueContext is an interface to support dynamic dispatch.
type ITableKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TableKey() ITableKeyContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsTableKeyValueContext differentiates from other interfaces.
	IsTableKeyValueContext()
}

type TableKeyValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableKeyValueContext() *TableKeyValueContext {
	var p = new(TableKeyValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_tableKeyValue
	return p
}

func InitEmptyTableKeyValueContext(p *TableKeyValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_tableKeyValue
}

func (*TableKeyValueContext) IsTableKeyValueContext() {}

func NewTableKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableKeyValueContext {
	var p = new(TableKeyValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_tableKeyValue

	return p
}

func (s *TableKeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TableKeyValueContext) TableKey() ITableKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableKeyContext)
}

func (s *TableKeyValueContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(InscriptParserASSIGN, 0)
}

func (s *TableKeyValueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TableKeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableKeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableKeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterTableKeyValue(s)
	}
}

func (s *TableKeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitTableKeyValue(s)
	}
}

func (s *TableKeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitTableKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) TableKeyValue() (localctx ITableKeyValueContext) {
	localctx = NewTableKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, InscriptParserRULE_tableKeyValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.TableKey()
	}
	{
		p.SetState(342)
		p.Match(InscriptParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableKeyContext is an interface to support dynamic dispatch.
type ITableKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsTableKeyContext differentiates from other interfaces.
	IsTableKeyContext()
}

type TableKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableKeyContext() *TableKeyContext {
	var p = new(TableKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_tableKey
	return p
}

func InitEmptyTableKeyContext(p *TableKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_tableKey
}

func (*TableKeyContext) IsTableKeyContext() {}

func NewTableKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableKeyContext {
	var p = new(TableKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_tableKey

	return p
}

func (s *TableKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *TableKeyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TableKeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(InscriptParserSTRING, 0)
}

func (s *TableKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *TableKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterTableKey(s)
	}
}

func (s *TableKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitTableKey(s)
	}
}

func (s *TableKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitTableKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) TableKey() (localctx ITableKeyContext) {
	localctx = NewTableKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, InscriptParserRULE_tableKey)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Match(InscriptParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(347)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *InscriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PostfixExprContext = nil
		if localctx != nil {
			t = localctx.(*PostfixExprContext)
		}
		return p.PostfixExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *InscriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *InscriptParser) PostfixExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 20:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
