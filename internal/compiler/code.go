// Package bytecode defines the opcodes and helpers for making and decoding instructions.
package compiler

import (
	"fmt"
)

// Opcode is a single byte code that identifies the operation.
type Opcode byte

// Instruction widths by opcode: number and byte-width of each operand.
var operandWidths = map[Opcode][]int{
	OpConstant:      {2}, // constant pool index (uint16)
	OpSetGlobal:     {2}, // global name index
	OpJump:          {2}, // jump offset
	OpJumpNotTruthy: {2}, // jump offset
	OpCall:          {1}, // argument count
	OpPrint:         {1}, // number of expressions
}

// Opcode definitions
const (
	OpConstant Opcode = iota
	OpPop

	OpTrue
	OpFalse
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpLessThan

	OpBang  // logical NOT
	OpMinus // negation

	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod

	OpJumpNotTruthy
	OpJump

	OpNull // explicit nil/None

	OpGetGlobal
	OpSetGlobal

	OpArray // create array literal

	OpCall
	OpReturnValue
	OpReturn

	OpPrint
)

// Instructions is a slice of bytecode instructions.
type Instructions []byte

// Make constructs an instruction byte-slice for op and operands.
func Make(op Opcode, operands ...int) []byte {
	widths := operandWidths[op]
	size := 1
	for _, w := range widths {
		size += w
	}

	ins := make([]byte, size)
	ins[0] = byte(op)
	offset := 1

	for i, operand := range operands {
		width := widths[i]
		switch width {
		case 2:
			// big-endian uint16
			ins[offset] = byte(operand >> 8)
			ins[offset+1] = byte(operand)
		case 1:
			ins[offset] = byte(operand)
		default:
			panic(fmt.Sprintf("unsupported operand width %d", width))
		}
		offset += width
	}

	return ins
}

// ReadOperand reads a single operand of given byte-width and returns the value and bytes consumed.
func ReadOperand(ins Instructions, offset, width int) (int, int) {
	switch width {
	case 2:
		hi := int(ins[offset])
		lo := int(ins[offset+1])
		return (hi << 8) | lo, 2
	case 1:
		return int(ins[offset]), 1
	default:
		panic(fmt.Sprintf("unsupported operand width %d", width))
	}
}

// ReadOperands decodes all operands for op starting at offset.
// It returns the slice of operand values and total bytes read.
func ReadOperands(op Opcode, ins Instructions, offset int) ([]int, int) {
	widths := operandWidths[op]
	operands := make([]int, len(widths))
	bytesRead := 0

	for i, width := range widths {
		operand, n := ReadOperand(ins, offset+bytesRead, width)
		operands[i] = operand
		bytesRead += n
	}
	return operands, bytesRead
}

// ReadOpcode returns the Opcode at the given offset.
func ReadOpcode(ins Instructions, offset int) Opcode {
	return Opcode(ins[offset])
}

// String returns a human-readable representation of instructions for debugging.
func (ins Instructions) String() string {
	var out string
	i := 0
	for i < len(ins) {
		op := ReadOpcode(ins, i)
		operands, bytesRead := ReadOperands(op, ins, i+1)

		out += fmt.Sprintf("%04d %s", i, op.String())
		for _, operand := range operands {
			out += fmt.Sprintf(" %d", operand)
		}
		out += "\n"

		i += 1 + bytesRead
	}
	return out
}

// String maps opcodes to names.
func (op Opcode) String() string {
	switch op {
	case OpConstant:
		return "OpConstant"
	case OpPop:
		return "OpPop"
	case OpTrue:
		return "OpTrue"
	case OpFalse:
		return "OpFalse"
	case OpEqual:
		return "OpEqual"
	case OpNotEqual:
		return "OpNotEqual"
	case OpGreaterThan:
		return "OpGreaterThan"
	case OpLessThan:
		return "OpLessThan"
	case OpBang:
		return "OpBang"
	case OpMinus:
		return "OpMinus"
	case OpAdd:
		return "OpAdd"
	case OpSub:
		return "OpSub"
	case OpMul:
		return "OpMul"
	case OpDiv:
		return "OpDiv"
	case OpMod:
		return "OpMod"
	case OpJumpNotTruthy:
		return "OpJumpNotTruthy"
	case OpJump:
		return "OpJump"
	case OpNull:
		return "OpNull"
	case OpGetGlobal:
		return "OpGetGlobal"
	case OpSetGlobal:
		return "OpSetGlobal"
	case OpCall:
		return "OpCall"
	case OpReturnValue:
		return "OpReturnValue"
	case OpReturn:
		return "OpReturn"
	case OpPrint:
		return "OpPrint"
	default:
		return fmt.Sprintf("Opcode(%d)", op)
	}
}
