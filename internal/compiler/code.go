package compiler

import (
	"bytes"
	"fmt"
)

// OpCode represents a single operation code.
type OpCode byte

// Define your opcodes here. Use iota to make them sequential.
const (
	OpConstant OpCode = iota // Push a constant onto the stack (operand: constant index)
	OpAdd                    // Pop two, add, push result
	OpSub                    // Pop two, subtract, push result
	OpMul                    // Pop two, multiply, push result
	OpDiv                    // Pop two, divide, push result
	OpMod                    // Pop two, modulo, push result
	OpPow                    // Pop two (base, exp), power, push result
	OpMinus                  // Pop one, negate, push result
	OpNot                    // Pop one, logical NOT, push boolean

	OpEqual        // Pop two, compare ==, push boolean
	OpNotEqual     // Pop two, compare !=, push boolean
	OpGreaterThan  // Pop two, compare >, push boolean
	OpLessThan     // Pop two, compare <, push boolean
	OpGreaterEqual // Pop two, compare >=, push boolean
	OpLessEqual    // Pop two, compare <=, push boolean

	OpJump          // Unconditional jump (operand: target instruction index)
	OpJumpNotTruthy // Pop condition, jump if falsy (operand: target instruction index)
	OpJumpTruthy    // Pop condition, jump if truthy (operand: target instruction index)

	OpPop // Pop the top element from the stack

	OpNull  // Push null onto the stack
	OpTrue  // Push boolean true onto the stack
	OpFalse // Push boolean false onto the stack

	OpPrint // Pop N arguments (operand: N), print them

	OpSetGlobal // Pop value, set global variable (operand: global index)
	OpGetGlobal // Get global variable, push value (operand: global index)

	OpSetLocal // Pop value, set local variable (operand: local index within frame)
	OpGetLocal // Get local variable, push value (operand: local index within frame)

	OpArray    // Create an array from stack elements (operand: number of elements)
	OpHash     // Create a hash from stack key-value pairs (operand: number of pairs)
	OpIndex    // Pop index, pop aggregate, get indexed value, push result
	OpSetIndex // Pop value, pop index, pop aggregate, set indexed value

	OpCall        // Pop function object and args, call function (operand: number of arguments)
	OpReturnValue // Pop return value, return from function
	OpReturn      // Return from function (implicitly returns null)

	OpGetIterator  // Pop iterable, push iterator
	OpIteratorNext // Pop iterator, get next value/done, push value, push boolean - true if successful
)

// Definition describes an opcode and its operand widths.
type Definition struct {
	Name          string
	OperandWidths []int // Widths in bytes for each operand
}

// definitions maps OpCode to its Definition.
var definitions = map[OpCode]*Definition{
	OpConstant: {"OpConstant", []int{2}}, // Constant index (up to 65535)
	OpAdd:      {"OpAdd", []int{}},
	OpSub:      {"OpSub", []int{}},
	OpMul:      {"OpMul", []int{}},
	OpDiv:      {"OpDiv", []int{}},
	OpMod:      {"OpMod", []int{}},
	OpPow:      {"OpPow", []int{}},
	OpMinus:    {"OpMinus", []int{}},
	OpNot:      {"OpNot", []int{}},

	OpEqual:        {"OpEqual", []int{}},
	OpNotEqual:     {"OpNotEqual", []int{}},
	OpGreaterThan:  {"OpGreaterThan", []int{}},
	OpLessThan:     {"OpLessThan", []int{}},
	OpGreaterEqual: {"OpGreaterEqual", []int{}},
	OpLessEqual:    {"OpLessEqual", []int{}},

	OpJump:          {"OpJump", []int{2}},          // Target instruction index (up to 65535)
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}}, // Target instruction index
	OpJumpTruthy:    {"OpJumpTruthy", []int{2}},    // Target instruction index

	OpPop: {"OpPop", []int{}},

	OpNull:  {"OpNull", []int{}},
	OpTrue:  {"OpTrue", []int{}},
	OpFalse: {"OpFalse", []int{}},

	OpPrint: {"OpPrint", []int{1}}, // Number of arguments (up to 255)

	OpSetGlobal: {"OpSetGlobal", []int{2}}, // Global variable index (up to 65535)
	OpGetGlobal: {"OpGetGlobal", []int{2}}, // Global variable index

	OpSetLocal: {"OpSetLocal", []int{1}}, // Local variable index (up to 255 locals per frame)
	OpGetLocal: {"OpGetLocal", []int{1}}, // Local variable index

	OpArray:    {"OpArray", []int{2}}, // Number of elements (up to 65535)
	OpHash:     {"OpHash", []int{2}},  // Number of key-value pairs
	OpIndex:    {"OpIndex", []int{}},
	OpSetIndex: {"OpSetIndex", []int{}},

	OpCall:        {"OpCall", []int{1}}, // Number of arguments (up to 255)
	OpReturnValue: {"OpReturnValue", []int{}},
	OpReturn:      {"OpReturn", []int{}},

	OpGetIterator:  {"OpGetIterator", []int{}},
	OpIteratorNext: {"OpIteratorNext", []int{}},
}

// Lookup returns the Definition for a given OpCode.
func Lookup(op OpCode) (*Definition, bool) {
	def, ok := definitions[op]
	return def, ok
}

// Instruction represents a single compiled instruction.
// For simplicity, we store operands as int slice for now.
// In a real VM, operands would be encoded directly in a byte slice.
type Instruction struct {
	Op       OpCode
	Operands []int
}

// Bytecode represents a compiled program or function body.
type Bytecode struct {
	Instructions []Instruction // The sequence of instructions
	Constants    []interface{} // The constant pool (using interface{} to store various Go types)

	// Information for function calls and frames
	NumLocals     int // Total number of local variables (including parameters) needed for a frame
	NumParameters int // Number of parameters the function expects

	// Total number of global variables defined in the program (only for the main program bytecode)
	NumGlobals int
}

// NewBytecode creates an empty Bytecode object.
func NewBytecode() *Bytecode {
	return &Bytecode{
		Instructions:  make([]Instruction, 0),
		Constants:     make([]interface{}, 0),
		NumLocals:     0,
		NumParameters: 0,
		NumGlobals:    0,
	}
}

// FormatInstructions converts the Bytecode's instructions into a human-readable string.
// This is now a method on the Bytecode struct.
func (b *Bytecode) FormatInstructions() string {
	var out bytes.Buffer
	i := 0
	for i < len(b.Instructions) { // Use b.Instructions
		instr := b.Instructions[i] // Use b.Instructions
		def, ok := Lookup(instr.Op)
		if !ok {
			fmt.Fprintf(&out, "ERROR: unknown opcode %d at %d\n", instr.Op, i)
			i++
			continue
		}

		// Format operands
		operands := ""
		for j, operand := range instr.Operands {
			if j > 0 {
				operands += ", "
			}
			operands += fmt.Sprintf("%d", operand)
		}

		fmt.Fprintf(&out, "%04d %s %s\n", i, def.Name, operands)

		// Move to the next instruction.
		// The instruction length depends on the opcode and operand widths.
		// For now, since Instruction struct stores operands directly,
		// we just increment i by 1 per Instruction struct.
		// If using a byte slice, you'd calculate the length here.
		i++
	}
	return out.String()
}
