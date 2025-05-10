package compiler

// OpCode represents the type of a bytecode instruction.
type OpCode byte

// Define the instruction opcodes.
const (
	// Constants and stack ops
	OpConstant OpCode = iota // Push a constant onto the stack (operand is constant pool index)
	OpPop                    // Pop the top value from the stack (used for expression statements)
	OpNull                   // Push a null value onto the stack
	OpTrue                   // Push boolean true onto stack
	OpFalse                  // Push boolean false onto stack

	// Arithmetic
	OpAdd // Pop two operands, add them, push result
	OpSub // Pop two operands, subtract, push result
	OpMul // Pop two operands, multiply, push result
	OpDiv // Pop two operands, divide, push result
	OpMod // Pop two operands, modulo, push result
	OpPow // Pop two operands, power (base, exp), push result

	// Prefix/Unary
	OpMinus // Pop operand, negate (-), push result
	OpNot   // Pop operand, logical NOT, push boolean

	// Comparisons
	OpEqual        // Pop two operands, compare for equality (==), push boolean
	OpNotEqual     // Pop two operands, compare for inequality (!=), push boolean
	OpGreaterThan  // Pop two operands, compare for greater than (>), push boolean
	OpLessThan     // Pop two operands, compare for less than (<), push boolean
	OpGreaterEqual // Pop two operands, compare for greater than or equal (>=), push boolean
	OpLessEqual    // Pop two operands, compare for less than or equal (<=), push boolean

	// Jumps (operand is jump target address - absolute instruction index)
	OpJump          // Unconditional jump
	OpJumpNotTruthy // Jump if the top of stack is false/nil
	OpJumpTruthy    // Jump if the top of stack is true/not nil

	// Variables (operands are symbol indices)
	OpGetGlobal // Load a global variable onto the stack
	OpSetGlobal // Pop value, store in global variable
	OpGetLocal  // Load a local variable onto the stack (relative to frame pointer)
	OpSetLocal  // Pop value, store in local variable (relative to frame pointer)

	// Data structures (operands are element/field counts)
	OpArray    // Create an array/list (operand is number of elements)
	OpHash     // Create a hash/table (operand is number of key-value pairs)
	OpIndex    // Index into an array/hash (pop index, pop aggregate, push value)
	OpSetIndex // Set element at index (pop value, pop index, pop aggregate, push aggregate)

	// Iteration (for loops)
	OpGetIterator  // Get an iterator from an iterable (pop iterable, push iterator)
	OpIteratorNext // Get the next value from an iterator (pop iterator, push value, push boolean - true if successful)

	// Function call/return (operand for OpCall is argument count)
	OpCall        // Call a function
	OpReturnValue // Return from function with a value
	OpReturn      // Return from function without a value (implicitly null)

	// Print (operand is number of arguments to print)
	OpPrint
)

// Definition provides information about an opcode, including the number of operands and their widths.
type Definition struct {
	Name          string
	OperandWidths []int // Widths in bytes for each operand
}

// definitions maps each OpCode to its Definition.
var definitions = map[OpCode]*Definition{
	OpConstant: {"OpConstant", []int{2}}, // Index into constant pool (up to 65535)
	OpPop:      {"OpPop", []int{}},
	OpNull:     {"OpNull", []int{}},
	OpTrue:     {"OpTrue", []int{}},
	OpFalse:    {"OpFalse", []int{}},

	OpAdd: {"OpAdd", []int{}},
	OpSub: {"OpSub", []int{}},
	OpMul: {"OpMul", []int{}},
	OpDiv: {"OpDiv", []int{}},
	OpMod: {"OpMod", []int{}},
	OpPow: {"OpPow", []int{}},

	OpMinus: {"OpMinus", []int{}},
	OpNot:   {"OpNot", []int{}},

	OpEqual:        {"OpEqual", []int{}},
	OpNotEqual:     {"OpNotEqual", []int{}},
	OpGreaterThan:  {"OpGreaterThan", []int{}},
	OpLessThan:     {"OpLessThan", []int{}},
	OpGreaterEqual: {"OpGreaterEqual", []int{}},
	OpLessEqual:    {"OpLessEqual", []int{}},

	OpJump:          {"OpJump", []int{2}},          // Target address (absolute instruction index)
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}}, // Target address (absolute instruction index)
	OpJumpTruthy:    {"OpJumpTruthy", []int{2}},    // Target address (absolute instruction index)

	OpGetGlobal: {"OpGetGlobal", []int{2}}, // Index into global symbol table
	OpSetGlobal: {"OpSetGlobal", []int{2}}, // Index into global symbol table
	OpGetLocal:  {"OpGetLocal", []int{1}},  // Index into current frame's locals (up to 255)
	OpSetLocal:  {"OpSetLocal", []int{1}},  // Index into current frame's locals (up to 255)

	OpArray:    {"OpArray", []int{2}}, // Number of elements (up to 65535)
	OpHash:     {"OpHash", []int{2}},  // Number of key-value pairs (up to 65535)
	OpIndex:    {"OpIndex", []int{}},
	OpSetIndex: {"OpSetIndex", []int{}},

	OpGetIterator:  {"OpGetIterator", []int{}},
	OpIteratorNext: {"OpIteratorNext", []int{}},

	OpCall:        {"OpCall", []int{1}}, // Number of arguments
	OpReturnValue: {"OpReturnValue", []int{}},
	OpReturn:      {"OpReturn", []int{}},

	OpPrint: {"OpPrint", []int{1}}, // Number of args to print
}

// Lookup returns the Definition for an OpCode.
// Returns (nil, false) if the opcode is not defined.
func Lookup(op OpCode) (*Definition, bool) {
	def, ok := definitions[op]
	return def, ok
}

// Note: Functions for encoding/decoding instructions from byte slices
// would typically go here if you move to a []byte representation.
/*
func MakeInstruction(op OpCode, operands ...int) ([]byte, error) { ... }
func ReadOperands(def *Definition, instructions []byte) ([]int, int) { ... }
func ReadUint16(ins []byte) uint16 { ... }
func ReadUint8(ins []byte) uint8 { ... }
*/
