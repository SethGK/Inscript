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
		"", "'='", "'print'", "'('", "')'", "'return'", "'if'", "'elseif'",
		"'else'", "'while'", "'for'", "'in'", "'function'", "'{'", "'}'", "','",
		"'or'", "'and'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'^'", "'not'", "'['", "']'", "'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"BOOLEAN", "IDENTIFIER", "INTEGER", "FLOAT", "STRING", "WS", "LINE_COMMENT",
		"BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statementList", "statement", "simpleStmt", "assignment",
		"exprStmt", "printStmt", "returnStmt", "compoundStmt", "ifStmt", "elseifListOpt",
		"elseifList", "elseif", "elseBlockOpt", "whileStmt", "forStmt", "functionDef",
		"block", "statementListOpt", "expressionOpt", "expressionListOpt", "expressionList",
		"paramListOpt", "paramList", "expression", "logicalOr", "logicalAnd",
		"comparison", "arith", "term", "factor", "unary", "primary", "atom",
		"listLiteral", "tableLiteral", "fieldListOpt", "fieldList", "field",
		"literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 301, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1,
		85, 8, 1, 11, 1, 12, 1, 86, 1, 2, 1, 2, 3, 2, 91, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 97, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 117,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 127, 8,
		10, 1, 11, 4, 11, 130, 8, 11, 11, 11, 12, 11, 131, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 141, 8, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 3, 18, 165, 8,
		18, 1, 19, 1, 19, 3, 19, 169, 8, 19, 1, 20, 1, 20, 3, 20, 173, 8, 20, 1,
		21, 1, 21, 1, 21, 5, 21, 178, 8, 21, 10, 21, 12, 21, 181, 9, 21, 1, 22,
		1, 22, 3, 22, 185, 8, 22, 1, 23, 1, 23, 1, 23, 5, 23, 190, 8, 23, 10, 23,
		12, 23, 193, 9, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 5, 25, 200, 8, 25,
		10, 25, 12, 25, 203, 9, 25, 1, 26, 1, 26, 1, 26, 5, 26, 208, 8, 26, 10,
		26, 12, 26, 211, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 216, 8, 27, 10, 27,
		12, 27, 219, 9, 27, 1, 28, 1, 28, 1, 28, 5, 28, 224, 8, 28, 10, 28, 12,
		28, 227, 9, 28, 1, 29, 1, 29, 1, 29, 5, 29, 232, 8, 29, 10, 29, 12, 29,
		235, 9, 29, 1, 30, 1, 30, 1, 30, 5, 30, 240, 8, 30, 10, 30, 12, 30, 243,
		9, 30, 1, 31, 1, 31, 1, 31, 3, 31, 248, 8, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 259, 8, 32, 10, 32, 12, 32,
		262, 9, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 3, 33, 273, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 36, 1, 36, 3, 36, 285, 8, 36, 1, 37, 1, 37, 1, 37, 5, 37, 290,
		8, 37, 10, 37, 12, 37, 293, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 39, 0, 0, 40, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 0, 5, 1, 0, 18, 23, 1, 0, 24, 25, 1, 0,
		26, 28, 2, 0, 24, 25, 30, 30, 2, 0, 33, 34, 36, 38, 293, 0, 80, 1, 0, 0,
		0, 2, 84, 1, 0, 0, 0, 4, 90, 1, 0, 0, 0, 6, 96, 1, 0, 0, 0, 8, 98, 1, 0,
		0, 0, 10, 102, 1, 0, 0, 0, 12, 104, 1, 0, 0, 0, 14, 109, 1, 0, 0, 0, 16,
		116, 1, 0, 0, 0, 18, 118, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 129, 1,
		0, 0, 0, 24, 133, 1, 0, 0, 0, 26, 140, 1, 0, 0, 0, 28, 142, 1, 0, 0, 0,
		30, 146, 1, 0, 0, 0, 32, 152, 1, 0, 0, 0, 34, 158, 1, 0, 0, 0, 36, 164,
		1, 0, 0, 0, 38, 168, 1, 0, 0, 0, 40, 172, 1, 0, 0, 0, 42, 174, 1, 0, 0,
		0, 44, 184, 1, 0, 0, 0, 46, 186, 1, 0, 0, 0, 48, 194, 1, 0, 0, 0, 50, 196,
		1, 0, 0, 0, 52, 204, 1, 0, 0, 0, 54, 212, 1, 0, 0, 0, 56, 220, 1, 0, 0,
		0, 58, 228, 1, 0, 0, 0, 60, 236, 1, 0, 0, 0, 62, 247, 1, 0, 0, 0, 64, 249,
		1, 0, 0, 0, 66, 272, 1, 0, 0, 0, 68, 274, 1, 0, 0, 0, 70, 278, 1, 0, 0,
		0, 72, 284, 1, 0, 0, 0, 74, 286, 1, 0, 0, 0, 76, 294, 1, 0, 0, 0, 78, 298,
		1, 0, 0, 0, 80, 81, 3, 2, 1, 0, 81, 82, 5, 0, 0, 1, 82, 1, 1, 0, 0, 0,
		83, 85, 3, 4, 2, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 91, 3, 6, 3, 0, 89,
		91, 3, 16, 8, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 5, 1, 0, 0,
		0, 92, 97, 3, 8, 4, 0, 93, 97, 3, 10, 5, 0, 94, 97, 3, 12, 6, 0, 95, 97,
		3, 14, 7, 0, 96, 92, 1, 0, 0, 0, 96, 93, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		96, 95, 1, 0, 0, 0, 97, 7, 1, 0, 0, 0, 98, 99, 3, 64, 32, 0, 99, 100, 5,
		1, 0, 0, 100, 101, 3, 48, 24, 0, 101, 9, 1, 0, 0, 0, 102, 103, 3, 48, 24,
		0, 103, 11, 1, 0, 0, 0, 104, 105, 5, 2, 0, 0, 105, 106, 5, 3, 0, 0, 106,
		107, 3, 40, 20, 0, 107, 108, 5, 4, 0, 0, 108, 13, 1, 0, 0, 0, 109, 110,
		5, 5, 0, 0, 110, 111, 3, 38, 19, 0, 111, 15, 1, 0, 0, 0, 112, 117, 3, 18,
		9, 0, 113, 117, 3, 28, 14, 0, 114, 117, 3, 30, 15, 0, 115, 117, 3, 32,
		16, 0, 116, 112, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0,
		116, 115, 1, 0, 0, 0, 117, 17, 1, 0, 0, 0, 118, 119, 5, 6, 0, 0, 119, 120,
		3, 48, 24, 0, 120, 121, 3, 34, 17, 0, 121, 122, 3, 20, 10, 0, 122, 123,
		3, 26, 13, 0, 123, 19, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 127, 3, 22,
		11, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 21, 1, 0, 0, 0,
		128, 130, 3, 24, 12, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 23, 1, 0, 0, 0, 133, 134, 5,
		7, 0, 0, 134, 135, 3, 48, 24, 0, 135, 136, 3, 34, 17, 0, 136, 25, 1, 0,
		0, 0, 137, 141, 1, 0, 0, 0, 138, 139, 5, 8, 0, 0, 139, 141, 3, 34, 17,
		0, 140, 137, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 27, 1, 0, 0, 0, 142,
		143, 5, 9, 0, 0, 143, 144, 3, 48, 24, 0, 144, 145, 3, 34, 17, 0, 145, 29,
		1, 0, 0, 0, 146, 147, 5, 10, 0, 0, 147, 148, 5, 35, 0, 0, 148, 149, 5,
		11, 0, 0, 149, 150, 3, 48, 24, 0, 150, 151, 3, 34, 17, 0, 151, 31, 1, 0,
		0, 0, 152, 153, 5, 12, 0, 0, 153, 154, 5, 3, 0, 0, 154, 155, 3, 44, 22,
		0, 155, 156, 5, 4, 0, 0, 156, 157, 3, 34, 17, 0, 157, 33, 1, 0, 0, 0, 158,
		159, 5, 13, 0, 0, 159, 160, 3, 36, 18, 0, 160, 161, 5, 14, 0, 0, 161, 35,
		1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 165, 3, 2, 1, 0, 164, 162, 1, 0,
		0, 0, 164, 163, 1, 0, 0, 0, 165, 37, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0,
		167, 169, 3, 48, 24, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169,
		39, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 173, 3, 42, 21, 0, 172, 170,
		1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 41, 1, 0, 0, 0, 174, 179, 3, 48,
		24, 0, 175, 176, 5, 15, 0, 0, 176, 178, 3, 48, 24, 0, 177, 175, 1, 0, 0,
		0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		43, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 185, 3,
		46, 23, 0, 184, 182, 1, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 45, 1, 0, 0,
		0, 186, 191, 5, 35, 0, 0, 187, 188, 5, 15, 0, 0, 188, 190, 5, 35, 0, 0,
		189, 187, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 47, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 195, 3,
		50, 25, 0, 195, 49, 1, 0, 0, 0, 196, 201, 3, 52, 26, 0, 197, 198, 5, 16,
		0, 0, 198, 200, 3, 52, 26, 0, 199, 197, 1, 0, 0, 0, 200, 203, 1, 0, 0,
		0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 51, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 204, 209, 3, 54, 27, 0, 205, 206, 5, 17, 0, 0, 206, 208,
		3, 54, 27, 0, 207, 205, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1,
		0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 53, 1, 0, 0, 0, 211, 209, 1, 0, 0,
		0, 212, 217, 3, 56, 28, 0, 213, 214, 7, 0, 0, 0, 214, 216, 3, 56, 28, 0,
		215, 213, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217,
		218, 1, 0, 0, 0, 218, 55, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 225, 3,
		58, 29, 0, 221, 222, 7, 1, 0, 0, 222, 224, 3, 58, 29, 0, 223, 221, 1, 0,
		0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0,
		226, 57, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 233, 3, 60, 30, 0, 229,
		230, 7, 2, 0, 0, 230, 232, 3, 60, 30, 0, 231, 229, 1, 0, 0, 0, 232, 235,
		1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 59, 1, 0,
		0, 0, 235, 233, 1, 0, 0, 0, 236, 241, 3, 62, 31, 0, 237, 238, 5, 29, 0,
		0, 238, 240, 3, 62, 31, 0, 239, 237, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0,
		241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 61, 1, 0, 0, 0, 243, 241,
		1, 0, 0, 0, 244, 245, 7, 3, 0, 0, 245, 248, 3, 62, 31, 0, 246, 248, 3,
		64, 32, 0, 247, 244, 1, 0, 0, 0, 247, 246, 1, 0, 0, 0, 248, 63, 1, 0, 0,
		0, 249, 260, 3, 66, 33, 0, 250, 251, 5, 31, 0, 0, 251, 252, 3, 48, 24,
		0, 252, 253, 5, 32, 0, 0, 253, 259, 1, 0, 0, 0, 254, 255, 5, 3, 0, 0, 255,
		256, 3, 40, 20, 0, 256, 257, 5, 4, 0, 0, 257, 259, 1, 0, 0, 0, 258, 250,
		1, 0, 0, 0, 258, 254, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0,
		0, 0, 260, 261, 1, 0, 0, 0, 261, 65, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0,
		263, 273, 3, 78, 39, 0, 264, 273, 5, 35, 0, 0, 265, 273, 3, 68, 34, 0,
		266, 273, 3, 70, 35, 0, 267, 268, 5, 3, 0, 0, 268, 269, 3, 48, 24, 0, 269,
		270, 5, 4, 0, 0, 270, 273, 1, 0, 0, 0, 271, 273, 3, 32, 16, 0, 272, 263,
		1, 0, 0, 0, 272, 264, 1, 0, 0, 0, 272, 265, 1, 0, 0, 0, 272, 266, 1, 0,
		0, 0, 272, 267, 1, 0, 0, 0, 272, 271, 1, 0, 0, 0, 273, 67, 1, 0, 0, 0,
		274, 275, 5, 31, 0, 0, 275, 276, 3, 40, 20, 0, 276, 277, 5, 32, 0, 0, 277,
		69, 1, 0, 0, 0, 278, 279, 5, 13, 0, 0, 279, 280, 3, 72, 36, 0, 280, 281,
		5, 14, 0, 0, 281, 71, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 285, 3, 74,
		37, 0, 284, 282, 1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 73, 1, 0, 0, 0,
		286, 291, 3, 76, 38, 0, 287, 288, 5, 15, 0, 0, 288, 290, 3, 76, 38, 0,
		289, 287, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 75, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 295, 5,
		35, 0, 0, 295, 296, 5, 1, 0, 0, 296, 297, 3, 48, 24, 0, 297, 77, 1, 0,
		0, 0, 298, 299, 7, 4, 0, 0, 299, 79, 1, 0, 0, 0, 25, 86, 90, 96, 116, 126,
		131, 140, 164, 168, 172, 179, 184, 191, 201, 209, 217, 225, 233, 241, 247,
		258, 260, 272, 284, 291,
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
	InscriptParserT__0          = 1
	InscriptParserT__1          = 2
	InscriptParserT__2          = 3
	InscriptParserT__3          = 4
	InscriptParserT__4          = 5
	InscriptParserT__5          = 6
	InscriptParserT__6          = 7
	InscriptParserT__7          = 8
	InscriptParserT__8          = 9
	InscriptParserT__9          = 10
	InscriptParserT__10         = 11
	InscriptParserT__11         = 12
	InscriptParserT__12         = 13
	InscriptParserT__13         = 14
	InscriptParserT__14         = 15
	InscriptParserT__15         = 16
	InscriptParserT__16         = 17
	InscriptParserT__17         = 18
	InscriptParserT__18         = 19
	InscriptParserT__19         = 20
	InscriptParserT__20         = 21
	InscriptParserT__21         = 22
	InscriptParserT__22         = 23
	InscriptParserT__23         = 24
	InscriptParserT__24         = 25
	InscriptParserT__25         = 26
	InscriptParserT__26         = 27
	InscriptParserT__27         = 28
	InscriptParserT__28         = 29
	InscriptParserT__29         = 30
	InscriptParserT__30         = 31
	InscriptParserT__31         = 32
	InscriptParserT__32         = 33
	InscriptParserBOOLEAN       = 34
	InscriptParserIDENTIFIER    = 35
	InscriptParserINTEGER       = 36
	InscriptParserFLOAT         = 37
	InscriptParserSTRING        = 38
	InscriptParserWS            = 39
	InscriptParserLINE_COMMENT  = 40
	InscriptParserBLOCK_COMMENT = 41
)

// InscriptParser rules.
const (
	InscriptParserRULE_program           = 0
	InscriptParserRULE_statementList     = 1
	InscriptParserRULE_statement         = 2
	InscriptParserRULE_simpleStmt        = 3
	InscriptParserRULE_assignment        = 4
	InscriptParserRULE_exprStmt          = 5
	InscriptParserRULE_printStmt         = 6
	InscriptParserRULE_returnStmt        = 7
	InscriptParserRULE_compoundStmt      = 8
	InscriptParserRULE_ifStmt            = 9
	InscriptParserRULE_elseifListOpt     = 10
	InscriptParserRULE_elseifList        = 11
	InscriptParserRULE_elseif            = 12
	InscriptParserRULE_elseBlockOpt      = 13
	InscriptParserRULE_whileStmt         = 14
	InscriptParserRULE_forStmt           = 15
	InscriptParserRULE_functionDef       = 16
	InscriptParserRULE_block             = 17
	InscriptParserRULE_statementListOpt  = 18
	InscriptParserRULE_expressionOpt     = 19
	InscriptParserRULE_expressionListOpt = 20
	InscriptParserRULE_expressionList    = 21
	InscriptParserRULE_paramListOpt      = 22
	InscriptParserRULE_paramList         = 23
	InscriptParserRULE_expression        = 24
	InscriptParserRULE_logicalOr         = 25
	InscriptParserRULE_logicalAnd        = 26
	InscriptParserRULE_comparison        = 27
	InscriptParserRULE_arith             = 28
	InscriptParserRULE_term              = 29
	InscriptParserRULE_factor            = 30
	InscriptParserRULE_unary             = 31
	InscriptParserRULE_primary           = 32
	InscriptParserRULE_atom              = 33
	InscriptParserRULE_listLiteral       = 34
	InscriptParserRULE_tableLiteral      = 35
	InscriptParserRULE_fieldListOpt      = 36
	InscriptParserRULE_fieldList         = 37
	InscriptParserRULE_field             = 38
	InscriptParserRULE_literal           = 39
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementList() IStatementListContext
	EOF() antlr.TerminalNode

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

func (s *ProgramContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(InscriptParserEOF, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.StatementList()
	}
	{
		p.SetState(81)
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

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
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

func (s *StatementListContext) Statement(i int) IStatementContext {
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

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InscriptParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&544437450348) != 0) {
		{
			p.SetState(83)
			p.Statement()
		}

		p.SetState(86)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleStmt() ISimpleStmtContext
	CompoundStmt() ICompoundStmtContext

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

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) CompoundStmt() ICompoundStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundStmtContext)
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
	p.EnterRule(localctx, 4, InscriptParserRULE_statement)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.SimpleStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.CompoundStmt()
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

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	ExprStmt() IExprStmtContext
	PrintStmt() IPrintStmtContext
	ReturnStmt() IReturnStmtContext

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_simpleStmt
	return p
}

func InitEmptySimpleStmtContext(p *SimpleStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_simpleStmt
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
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

func (s *SimpleStmtContext) ExprStmt() IExprStmtContext {
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

func (s *SimpleStmtContext) PrintStmt() IPrintStmtContext {
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

func (s *SimpleStmtContext) ReturnStmt() IReturnStmtContext {
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

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InscriptParserRULE_simpleStmt)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.ExprStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.PrintStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.ReturnStmt()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	Expression() IExpressionContext

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

func (s *AssignmentContext) Primary() IPrimaryContext {
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Primary()
	}
	{
		p.SetState(99)
		p.Match(InscriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Expression()
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
	p.EnterRule(localctx, 10, InscriptParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Expression()
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
	ExpressionListOpt() IExpressionListOptContext

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

func (s *PrintStmtContext) ExpressionListOpt() IExpressionListOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListOptContext)
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
	p.EnterRule(localctx, 12, InscriptParserRULE_printStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(InscriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(InscriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.ExpressionListOpt()
	}
	{
		p.SetState(107)
		p.Match(InscriptParserT__3)
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
	ExpressionOpt() IExpressionOptContext

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

func (s *ReturnStmtContext) ExpressionOpt() IExpressionOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionOptContext)
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
	p.EnterRule(localctx, 14, InscriptParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(InscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.ExpressionOpt()
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

// ICompoundStmtContext is an interface to support dynamic dispatch.
type ICompoundStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	ForStmt() IForStmtContext
	FunctionDef() IFunctionDefContext

	// IsCompoundStmtContext differentiates from other interfaces.
	IsCompoundStmtContext()
}

type CompoundStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundStmtContext() *CompoundStmtContext {
	var p = new(CompoundStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_compoundStmt
	return p
}

func InitEmptyCompoundStmtContext(p *CompoundStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_compoundStmt
}

func (*CompoundStmtContext) IsCompoundStmtContext() {}

func NewCompoundStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundStmtContext {
	var p = new(CompoundStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_compoundStmt

	return p
}

func (s *CompoundStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundStmtContext) IfStmt() IIfStmtContext {
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

func (s *CompoundStmtContext) WhileStmt() IWhileStmtContext {
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

func (s *CompoundStmtContext) ForStmt() IForStmtContext {
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

func (s *CompoundStmtContext) FunctionDef() IFunctionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *CompoundStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterCompoundStmt(s)
	}
}

func (s *CompoundStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitCompoundStmt(s)
	}
}

func (s *CompoundStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitCompoundStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) CompoundStmt() (localctx ICompoundStmtContext) {
	localctx = NewCompoundStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, InscriptParserRULE_compoundStmt)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.IfStmt()
		}

	case InscriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.WhileStmt()
		}

	case InscriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)
			p.ForStmt()
		}

	case InscriptParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(115)
			p.FunctionDef()
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

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Block() IBlockContext
	ElseifListOpt() IElseifListOptContext
	ElseBlockOpt() IElseBlockOptContext

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

func (s *IfStmtContext) Block() IBlockContext {
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

func (s *IfStmtContext) ElseifListOpt() IElseifListOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifListOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifListOptContext)
}

func (s *IfStmtContext) ElseBlockOpt() IElseBlockOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockOptContext)
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
	p.EnterRule(localctx, 18, InscriptParserRULE_ifStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(InscriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Expression()
	}
	{
		p.SetState(120)
		p.Block()
	}
	{
		p.SetState(121)
		p.ElseifListOpt()
	}
	{
		p.SetState(122)
		p.ElseBlockOpt()
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

// IElseifListOptContext is an interface to support dynamic dispatch.
type IElseifListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElseifList() IElseifListContext

	// IsElseifListOptContext differentiates from other interfaces.
	IsElseifListOptContext()
}

type ElseifListOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifListOptContext() *ElseifListOptContext {
	var p = new(ElseifListOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseifListOpt
	return p
}

func InitEmptyElseifListOptContext(p *ElseifListOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseifListOpt
}

func (*ElseifListOptContext) IsElseifListOptContext() {}

func NewElseifListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifListOptContext {
	var p = new(ElseifListOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_elseifListOpt

	return p
}

func (s *ElseifListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifListOptContext) ElseifList() IElseifListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifListContext)
}

func (s *ElseifListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterElseifListOpt(s)
	}
}

func (s *ElseifListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitElseifListOpt(s)
	}
}

func (s *ElseifListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitElseifListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ElseifListOpt() (localctx IElseifListOptContext) {
	localctx = NewElseifListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, InscriptParserRULE_elseifListOpt)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserEOF, InscriptParserT__1, InscriptParserT__2, InscriptParserT__4, InscriptParserT__5, InscriptParserT__7, InscriptParserT__8, InscriptParserT__9, InscriptParserT__11, InscriptParserT__12, InscriptParserT__13, InscriptParserT__23, InscriptParserT__24, InscriptParserT__29, InscriptParserT__30, InscriptParserT__32, InscriptParserBOOLEAN, InscriptParserIDENTIFIER, InscriptParserINTEGER, InscriptParserFLOAT, InscriptParserSTRING:
		p.EnterOuterAlt(localctx, 1)

	case InscriptParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.ElseifList()
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

// IElseifListContext is an interface to support dynamic dispatch.
type IElseifListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElseif() []IElseifContext
	Elseif(i int) IElseifContext

	// IsElseifListContext differentiates from other interfaces.
	IsElseifListContext()
}

type ElseifListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifListContext() *ElseifListContext {
	var p = new(ElseifListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseifList
	return p
}

func InitEmptyElseifListContext(p *ElseifListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseifList
}

func (*ElseifListContext) IsElseifListContext() {}

func NewElseifListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifListContext {
	var p = new(ElseifListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_elseifList

	return p
}

func (s *ElseifListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifListContext) AllElseif() []IElseifContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifContext); ok {
			len++
		}
	}

	tst := make([]IElseifContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifContext); ok {
			tst[i] = t.(IElseifContext)
			i++
		}
	}

	return tst
}

func (s *ElseifListContext) Elseif(i int) IElseifContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifContext); ok {
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

	return t.(IElseifContext)
}

func (s *ElseifListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterElseifList(s)
	}
}

func (s *ElseifListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitElseifList(s)
	}
}

func (s *ElseifListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitElseifList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ElseifList() (localctx IElseifListContext) {
	localctx = NewElseifListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, InscriptParserRULE_elseifList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == InscriptParserT__6 {
		{
			p.SetState(128)
			p.Elseif()
		}

		p.SetState(131)
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

// IElseifContext is an interface to support dynamic dispatch.
type IElseifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Block() IBlockContext

	// IsElseifContext differentiates from other interfaces.
	IsElseifContext()
}

type ElseifContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifContext() *ElseifContext {
	var p = new(ElseifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseif
	return p
}

func InitEmptyElseifContext(p *ElseifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseif
}

func (*ElseifContext) IsElseifContext() {}

func NewElseifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifContext {
	var p = new(ElseifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_elseif

	return p
}

func (s *ElseifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifContext) Expression() IExpressionContext {
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

func (s *ElseifContext) Block() IBlockContext {
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

func (s *ElseifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterElseif(s)
	}
}

func (s *ElseifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitElseif(s)
	}
}

func (s *ElseifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitElseif(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Elseif() (localctx IElseifContext) {
	localctx = NewElseifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, InscriptParserRULE_elseif)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(InscriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Expression()
	}
	{
		p.SetState(135)
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

// IElseBlockOptContext is an interface to support dynamic dispatch.
type IElseBlockOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsElseBlockOptContext differentiates from other interfaces.
	IsElseBlockOptContext()
}

type ElseBlockOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockOptContext() *ElseBlockOptContext {
	var p = new(ElseBlockOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseBlockOpt
	return p
}

func InitEmptyElseBlockOptContext(p *ElseBlockOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_elseBlockOpt
}

func (*ElseBlockOptContext) IsElseBlockOptContext() {}

func NewElseBlockOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockOptContext {
	var p = new(ElseBlockOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_elseBlockOpt

	return p
}

func (s *ElseBlockOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockOptContext) Block() IBlockContext {
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

func (s *ElseBlockOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterElseBlockOpt(s)
	}
}

func (s *ElseBlockOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitElseBlockOpt(s)
	}
}

func (s *ElseBlockOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitElseBlockOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ElseBlockOpt() (localctx IElseBlockOptContext) {
	localctx = NewElseBlockOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, InscriptParserRULE_elseBlockOpt)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserEOF, InscriptParserT__1, InscriptParserT__2, InscriptParserT__4, InscriptParserT__5, InscriptParserT__8, InscriptParserT__9, InscriptParserT__11, InscriptParserT__12, InscriptParserT__13, InscriptParserT__23, InscriptParserT__24, InscriptParserT__29, InscriptParserT__30, InscriptParserT__32, InscriptParserBOOLEAN, InscriptParserIDENTIFIER, InscriptParserINTEGER, InscriptParserFLOAT, InscriptParserSTRING:
		p.EnterOuterAlt(localctx, 1)

	case InscriptParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(InscriptParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Block()
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

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.EnterRule(localctx, 28, InscriptParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(InscriptParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Expression()
	}
	{
		p.SetState(144)
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
	IDENTIFIER() antlr.TerminalNode
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

func (s *ForStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 30, InscriptParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(InscriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(InscriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(InscriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Expression()
	}
	{
		p.SetState(150)
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

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamListOpt() IParamListOptContext
	Block() IBlockContext

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_functionDef
	return p
}

func InitEmptyFunctionDefContext(p *FunctionDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_functionDef
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) ParamListOpt() IParamListOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListOptContext)
}

func (s *FunctionDefContext) Block() IBlockContext {
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

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

func (s *FunctionDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitFunctionDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, InscriptParserRULE_functionDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(InscriptParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(InscriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.ParamListOpt()
	}
	{
		p.SetState(155)
		p.Match(InscriptParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementListOpt() IStatementListOptContext

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

func (s *BlockContext) StatementListOpt() IStatementListOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListOptContext)
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
	p.EnterRule(localctx, 34, InscriptParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(InscriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.StatementListOpt()
	}
	{
		p.SetState(160)
		p.Match(InscriptParserT__13)
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

// IStatementListOptContext is an interface to support dynamic dispatch.
type IStatementListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementList() IStatementListContext

	// IsStatementListOptContext differentiates from other interfaces.
	IsStatementListOptContext()
}

type StatementListOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListOptContext() *StatementListOptContext {
	var p = new(StatementListOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_statementListOpt
	return p
}

func InitEmptyStatementListOptContext(p *StatementListOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_statementListOpt
}

func (*StatementListOptContext) IsStatementListOptContext() {}

func NewStatementListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListOptContext {
	var p = new(StatementListOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_statementListOpt

	return p
}

func (s *StatementListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListOptContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *StatementListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterStatementListOpt(s)
	}
}

func (s *StatementListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitStatementListOpt(s)
	}
}

func (s *StatementListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitStatementListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) StatementListOpt() (localctx IStatementListOptContext) {
	localctx = NewStatementListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, InscriptParserRULE_statementListOpt)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__13:
		p.EnterOuterAlt(localctx, 1)

	case InscriptParserT__1, InscriptParserT__2, InscriptParserT__4, InscriptParserT__5, InscriptParserT__8, InscriptParserT__9, InscriptParserT__11, InscriptParserT__12, InscriptParserT__23, InscriptParserT__24, InscriptParserT__29, InscriptParserT__30, InscriptParserT__32, InscriptParserBOOLEAN, InscriptParserIDENTIFIER, InscriptParserINTEGER, InscriptParserFLOAT, InscriptParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.StatementList()
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

// IExpressionOptContext is an interface to support dynamic dispatch.
type IExpressionOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpressionOptContext differentiates from other interfaces.
	IsExpressionOptContext()
}

type ExpressionOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionOptContext() *ExpressionOptContext {
	var p = new(ExpressionOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expressionOpt
	return p
}

func InitEmptyExpressionOptContext(p *ExpressionOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expressionOpt
}

func (*ExpressionOptContext) IsExpressionOptContext() {}

func NewExpressionOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionOptContext {
	var p = new(ExpressionOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_expressionOpt

	return p
}

func (s *ExpressionOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionOptContext) Expression() IExpressionContext {
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

func (s *ExpressionOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterExpressionOpt(s)
	}
}

func (s *ExpressionOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitExpressionOpt(s)
	}
}

func (s *ExpressionOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitExpressionOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ExpressionOpt() (localctx IExpressionOptContext) {
	localctx = NewExpressionOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, InscriptParserRULE_expressionOpt)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Expression()
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

// IExpressionListOptContext is an interface to support dynamic dispatch.
type IExpressionListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionList() IExpressionListContext

	// IsExpressionListOptContext differentiates from other interfaces.
	IsExpressionListOptContext()
}

type ExpressionListOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListOptContext() *ExpressionListOptContext {
	var p = new(ExpressionListOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expressionListOpt
	return p
}

func InitEmptyExpressionListOptContext(p *ExpressionListOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expressionListOpt
}

func (*ExpressionListOptContext) IsExpressionListOptContext() {}

func NewExpressionListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListOptContext {
	var p = new(ExpressionListOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_expressionListOpt

	return p
}

func (s *ExpressionListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListOptContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterExpressionListOpt(s)
	}
}

func (s *ExpressionListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitExpressionListOpt(s)
	}
}

func (s *ExpressionListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitExpressionListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ExpressionListOpt() (localctx IExpressionListOptContext) {
	localctx = NewExpressionListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, InscriptParserRULE_expressionListOpt)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__3, InscriptParserT__31:
		p.EnterOuterAlt(localctx, 1)

	case InscriptParserT__2, InscriptParserT__11, InscriptParserT__12, InscriptParserT__23, InscriptParserT__24, InscriptParserT__29, InscriptParserT__30, InscriptParserT__32, InscriptParserBOOLEAN, InscriptParserIDENTIFIER, InscriptParserINTEGER, InscriptParserFLOAT, InscriptParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.ExpressionList()
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

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, InscriptParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Expression()
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserT__14 {
		{
			p.SetState(175)
			p.Match(InscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Expression()
		}

		p.SetState(181)
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

// IParamListOptContext is an interface to support dynamic dispatch.
type IParamListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamList() IParamListContext

	// IsParamListOptContext differentiates from other interfaces.
	IsParamListOptContext()
}

type ParamListOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListOptContext() *ParamListOptContext {
	var p = new(ParamListOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_paramListOpt
	return p
}

func InitEmptyParamListOptContext(p *ParamListOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_paramListOpt
}

func (*ParamListOptContext) IsParamListOptContext() {}

func NewParamListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListOptContext {
	var p = new(ParamListOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_paramListOpt

	return p
}

func (s *ParamListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListOptContext) ParamList() IParamListContext {
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

func (s *ParamListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterParamListOpt(s)
	}
}

func (s *ParamListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitParamListOpt(s)
	}
}

func (s *ParamListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitParamListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) ParamListOpt() (localctx IParamListOptContext) {
	localctx = NewParamListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, InscriptParserRULE_paramListOpt)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__3:
		p.EnterOuterAlt(localctx, 1)

	case InscriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.ParamList()
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

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

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

func (s *ParamListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(InscriptParserIDENTIFIER)
}

func (s *ParamListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, i)
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
	p.EnterRule(localctx, 46, InscriptParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(InscriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserT__14 {
		{
			p.SetState(187)
			p.Match(InscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(193)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOr() ILogicalOrContext

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

func (s *ExpressionContext) LogicalOr() ILogicalOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, InscriptParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.LogicalOr()
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

// ILogicalOrContext is an interface to support dynamic dispatch.
type ILogicalOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAnd() []ILogicalAndContext
	LogicalAnd(i int) ILogicalAndContext

	// IsLogicalOrContext differentiates from other interfaces.
	IsLogicalOrContext()
}

type LogicalOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrContext() *LogicalOrContext {
	var p = new(LogicalOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_logicalOr
	return p
}

func InitEmptyLogicalOrContext(p *LogicalOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_logicalOr
}

func (*LogicalOrContext) IsLogicalOrContext() {}

func NewLogicalOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrContext {
	var p = new(LogicalOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_logicalOr

	return p
}

func (s *LogicalOrContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrContext) AllLogicalAnd() []ILogicalAndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndContext); ok {
			tst[i] = t.(ILogicalAndContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrContext) LogicalAnd(i int) ILogicalAndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndContext); ok {
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

	return t.(ILogicalAndContext)
}

func (s *LogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterLogicalOr(s)
	}
}

func (s *LogicalOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitLogicalOr(s)
	}
}

func (s *LogicalOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitLogicalOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) LogicalOr() (localctx ILogicalOrContext) {
	localctx = NewLogicalOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, InscriptParserRULE_logicalOr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.LogicalAnd()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserT__15 {
		{
			p.SetState(197)
			p.Match(InscriptParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.LogicalAnd()
		}

		p.SetState(203)
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

// ILogicalAndContext is an interface to support dynamic dispatch.
type ILogicalAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparison() []IComparisonContext
	Comparison(i int) IComparisonContext

	// IsLogicalAndContext differentiates from other interfaces.
	IsLogicalAndContext()
}

type LogicalAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndContext() *LogicalAndContext {
	var p = new(LogicalAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_logicalAnd
	return p
}

func InitEmptyLogicalAndContext(p *LogicalAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_logicalAnd
}

func (*LogicalAndContext) IsLogicalAndContext() {}

func NewLogicalAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndContext {
	var p = new(LogicalAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_logicalAnd

	return p
}

func (s *LogicalAndContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndContext) AllComparison() []IComparisonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonContext); ok {
			len++
		}
	}

	tst := make([]IComparisonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonContext); ok {
			tst[i] = t.(IComparisonContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndContext) Comparison(i int) IComparisonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
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

	return t.(IComparisonContext)
}

func (s *LogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterLogicalAnd(s)
	}
}

func (s *LogicalAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitLogicalAnd(s)
	}
}

func (s *LogicalAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitLogicalAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) LogicalAnd() (localctx ILogicalAndContext) {
	localctx = NewLogicalAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, InscriptParserRULE_logicalAnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Comparison()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserT__16 {
		{
			p.SetState(205)
			p.Match(InscriptParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Comparison()
		}

		p.SetState(211)
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

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCmpOp returns the cmpOp token.
	GetCmpOp() antlr.Token

	// SetCmpOp sets the cmpOp token.
	SetCmpOp(antlr.Token)

	// Getter signatures
	AllArith() []IArithContext
	Arith(i int) IArithContext

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	cmpOp  antlr.Token
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) GetCmpOp() antlr.Token { return s.cmpOp }

func (s *ComparisonContext) SetCmpOp(v antlr.Token) { s.cmpOp = v }

func (s *ComparisonContext) AllArith() []IArithContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithContext); ok {
			len++
		}
	}

	tst := make([]IArithContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithContext); ok {
			tst[i] = t.(IArithContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) Arith(i int) IArithContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithContext); ok {
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

	return t.(IArithContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, InscriptParserRULE_comparison)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Arith()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16515072) != 0 {
		{
			p.SetState(213)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ComparisonContext).cmpOp = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16515072) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ComparisonContext).cmpOp = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(214)
			p.Arith()
		}

		p.SetState(219)
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

// IArithContext is an interface to support dynamic dispatch.
type IArithContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext

	// IsArithContext differentiates from other interfaces.
	IsArithContext()
}

type ArithContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyArithContext() *ArithContext {
	var p = new(ArithContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_arith
	return p
}

func InitEmptyArithContext(p *ArithContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_arith
}

func (*ArithContext) IsArithContext() {}

func NewArithContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithContext {
	var p = new(ArithContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_arith

	return p
}

func (s *ArithContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithContext) GetOp() antlr.Token { return s.op }

func (s *ArithContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ArithContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *ArithContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterArith(s)
	}
}

func (s *ArithContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitArith(s)
	}
}

func (s *ArithContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitArith(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Arith() (localctx IArithContext) {
	localctx = NewArithContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, InscriptParserRULE_arith)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Term()
	}
	p.SetState(225)
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
			{
				p.SetState(221)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ArithContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == InscriptParserT__23 || _la == InscriptParserT__24) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ArithContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(222)
				p.Term()
			}

		}
		p.SetState(227)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
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

	return t.(IFactorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, InscriptParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Factor()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469762048) != 0 {
		{
			p.SetState(229)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469762048) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(230)
			p.Factor()
		}

		p.SetState(235)
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

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnary() []IUnaryContext
	Unary(i int) IUnaryContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) AllUnary() []IUnaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryContext); ok {
			len++
		}
	}

	tst := make([]IUnaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryContext); ok {
			tst[i] = t.(IUnaryContext)
			i++
		}
	}

	return tst
}

func (s *FactorContext) Unary(i int) IUnaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
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

	return t.(IUnaryContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, InscriptParserRULE_factor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Unary()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserT__28 {
		{
			p.SetState(237)
			p.Match(InscriptParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Unary()
		}

		p.SetState(243)
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

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary() IUnaryContext
	Primary() IPrimaryContext

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryContext) Primary() IPrimaryContext {
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

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, InscriptParserRULE_unary)
	var _la int

	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__23, InscriptParserT__24, InscriptParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1124073472) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(245)
			p.Unary()
		}

	case InscriptParserT__2, InscriptParserT__11, InscriptParserT__12, InscriptParserT__30, InscriptParserT__32, InscriptParserBOOLEAN, InscriptParserIDENTIFIER, InscriptParserINTEGER, InscriptParserFLOAT, InscriptParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Primary()
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

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Atom() IAtomContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllExpressionListOpt() []IExpressionListOptContext
	ExpressionListOpt(i int) IExpressionListOptContext

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

func (s *PrimaryContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
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

func (s *PrimaryContext) AllExpressionListOpt() []IExpressionListOptContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListOptContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListOptContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListOptContext); ok {
			tst[i] = t.(IExpressionListOptContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) ExpressionListOpt(i int) IExpressionListOptContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListOptContext); ok {
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

	return t.(IExpressionListOptContext)
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
	p.EnterRule(localctx, 64, InscriptParserRULE_primary)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Atom()
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(258)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case InscriptParserT__30:
				{
					p.SetState(250)
					p.Match(InscriptParserT__30)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(251)
					p.Expression()
				}
				{
					p.SetState(252)
					p.Match(InscriptParserT__31)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case InscriptParserT__2:
				{
					p.SetState(254)
					p.Match(InscriptParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(255)
					p.ExpressionListOpt()
				}
				{
					p.SetState(256)
					p.Match(InscriptParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode
	ListLiteral() IListLiteralContext
	TableLiteral() ITableLiteralContext
	Expression() IExpressionContext
	FunctionDef() IFunctionDefContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Literal() ILiteralContext {
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

func (s *AtomContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *AtomContext) ListLiteral() IListLiteralContext {
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

func (s *AtomContext) TableLiteral() ITableLiteralContext {
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

func (s *AtomContext) Expression() IExpressionContext {
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

func (s *AtomContext) FunctionDef() IFunctionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, InscriptParserRULE_atom)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__32, InscriptParserBOOLEAN, InscriptParserINTEGER, InscriptParserFLOAT, InscriptParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Literal()
		}

	case InscriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Match(InscriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InscriptParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.ListLiteral()
		}

	case InscriptParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(266)
			p.TableLiteral()
		}

	case InscriptParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(267)
			p.Match(InscriptParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Expression()
		}
		{
			p.SetState(269)
			p.Match(InscriptParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InscriptParserT__11:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(271)
			p.FunctionDef()
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

// IListLiteralContext is an interface to support dynamic dispatch.
type IListLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionListOpt() IExpressionListOptContext

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

func (s *ListLiteralContext) ExpressionListOpt() IExpressionListOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListOptContext)
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
	p.EnterRule(localctx, 68, InscriptParserRULE_listLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(InscriptParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.ExpressionListOpt()
	}
	{
		p.SetState(276)
		p.Match(InscriptParserT__31)
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
	FieldListOpt() IFieldListOptContext

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

func (s *TableLiteralContext) FieldListOpt() IFieldListOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListOptContext)
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
	p.EnterRule(localctx, 70, InscriptParserRULE_tableLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(InscriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.FieldListOpt()
	}
	{
		p.SetState(280)
		p.Match(InscriptParserT__13)
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

// IFieldListOptContext is an interface to support dynamic dispatch.
type IFieldListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldList() IFieldListContext

	// IsFieldListOptContext differentiates from other interfaces.
	IsFieldListOptContext()
}

type FieldListOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListOptContext() *FieldListOptContext {
	var p = new(FieldListOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_fieldListOpt
	return p
}

func InitEmptyFieldListOptContext(p *FieldListOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_fieldListOpt
}

func (*FieldListOptContext) IsFieldListOptContext() {}

func NewFieldListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListOptContext {
	var p = new(FieldListOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_fieldListOpt

	return p
}

func (s *FieldListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListOptContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *FieldListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterFieldListOpt(s)
	}
}

func (s *FieldListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitFieldListOpt(s)
	}
}

func (s *FieldListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitFieldListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) FieldListOpt() (localctx IFieldListOptContext) {
	localctx = NewFieldListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, InscriptParserRULE_fieldListOpt)
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InscriptParserT__13:
		p.EnterOuterAlt(localctx, 1)

	case InscriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.FieldList()
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

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_fieldList
	return p
}

func InitEmptyFieldListContext(p *FieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_fieldList
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldListContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterFieldList(s)
	}
}

func (s *FieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitFieldList(s)
	}
}

func (s *FieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, InscriptParserRULE_fieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Field()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InscriptParserT__14 {
		{
			p.SetState(287)
			p.Match(InscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.Field()
		}

		p.SetState(293)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InscriptParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InscriptParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InscriptParserIDENTIFIER, 0)
}

func (s *FieldContext) Expression() IExpressionContext {
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

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InscriptListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InscriptVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InscriptParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, InscriptParserRULE_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(InscriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(InscriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Expression()
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
	INTEGER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

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

func (s *LiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(InscriptParserINTEGER, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(InscriptParserFLOAT, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(InscriptParserSTRING, 0)
}

func (s *LiteralContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(InscriptParserBOOLEAN, 0)
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
	p.EnterRule(localctx, 78, InscriptParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&506806140928) != 0) {
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
