package compiler // Package compiler, as it's in the internal/compiler directory

import (
	"bytes"
	"encoding/binary" // Needed for binary encoding/decoding
	"fmt"

	// Import the new types package to refer to types.Value
	"github.com/SethGK/Inscript/internal/types"
	// No import needed for vm package anymore
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
	OpGetFree
	OpClosure

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

	OpGetFree: {"OpGetFree", []int{1}},
	OpClosure: {"OpClosure", []int{2, 1}},

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

// Make creates an instruction byte slice from an opcode and operands.
func Make(op OpCode, operands ...int) []byte {
	def, ok := Lookup(op)
	if !ok {
		// TODO: Return an error or panic for unknown opcode
		return []byte{} // Return empty slice for now
	}

	instructionLen := 1 // Opcode byte
	for _, width := range def.OperandWidths {
		instructionLen += width
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 1:
			instruction[offset] = byte(operand)
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width
	}

	return instruction
}

// Instructions is a slice of instruction bytes.
type Instructions []byte

// ReadUint16 reads a 16-bit unsigned integer from an instruction slice.
func ReadUint16(is Instructions) uint16 {
	// Ensure the slice has at least 2 bytes
	if len(is) < 2 {
		// TODO: Handle error properly
		return 0
	}
	return binary.BigEndian.Uint16(is)
}

// ReadUint8 reads an 8-bit unsigned integer from an instruction slice.
func ReadUint8(is Instructions) uint8 {
	// Ensure the slice has at least 1 byte
	if len(is) < 1 {
		// TODO: Handle error properly
		return 0
	}
	return uint8(is[0])
}

// IntToBytes converts an integer to a 2-byte slice (big endian).
func IntToBytes(i int) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

// Bytecode represents a compiled program or function body.
type Bytecode struct {
	Instructions Instructions  // The sequence of instruction bytes
	Constants    []types.Value // The constant pool (stores types.Value types) - Corrected type

	// Information for function calls and frames
	NumLocals     int // Total number of local variables (including parameters) needed for a frame
	NumParameters int // Number of parameters the function expects

	// Total number of global variables defined in the program (only for the main program bytecode)
	NumGlobals int
}

// NewBytecode creates an empty Bytecode object.
func NewBytecode() *Bytecode {
	return &Bytecode{
		Instructions:  make([]byte, 0),        // Instructions is now []byte
		Constants:     make([]types.Value, 0), // Constants is now []types.Value
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
		instr := b.Instructions[i:] // Slice from current position
		opcode := OpCode(instr[0])  // Get opcode from the byte slice

		def, ok := Lookup(opcode)
		if !ok {
			fmt.Fprintf(&out, "ERROR: unknown opcode %d at %d\n", opcode, i)
			i++
			continue
		}

		// Read operands based on definition
		operands := ""
		operandOffset := 1 // Start after the opcode byte
		for j, width := range def.OperandWidths {
			if j > 0 {
				operands += ", "
			}
			// Ensure there are enough bytes for the operand
			if operandOffset+width > len(instr) {
				fmt.Fprintf(&out, "ERROR: not enough bytes for operand of %s at %d\n", def.Name, i)
				// Attempt to print remaining bytes as hex for debugging
				fmt.Fprintf(&out, " (remaining bytes: %x)\n", instr[operandOffset:])
				i += len(instr) - i // Skip to end to avoid infinite loop
				continue
			}

			switch width {
			case 1:
				operand := ReadUint8(instr[operandOffset:])
				operands += fmt.Sprintf("%d", operand)
			case 2:
				operand := ReadUint16(instr[operandOffset:])
				operands += fmt.Sprintf("%d", operand)
			}
			operandOffset += width
		}

		fmt.Fprintf(&out, "%04d %s %s\n", i, def.Name, operands)

		// Move to the next instruction.
		// The instruction length is 1 (opcode) + total operand bytes
		i += 1 + operandOffset - 1 // 1 (opcode) + total operand bytes
	}
	return out.String()
}

// GetLastOpcode returns the opcode of the last instruction in the byte slice.
// This is a helper for the compiler's internal checks.
// Note: This is a simplification and assumes the last instruction is valid.
func GetLastOpcode(instrs Instructions) OpCode {
	if len(instrs) == 0 {
		return 0 // Return a zero value for Opcode
	}
	// This is a simplification. A proper implementation would parse the last instruction.
	// For now, assuming the last byte is the opcode for simple cases like OpReturn/OpReturnValue.
	// This is incorrect in general if opcodes have operands.
	// TODO: Implement proper last instruction parsing from []byte if needed for more complex checks.
	return OpCode(instrs[len(instrs)-1])
}
