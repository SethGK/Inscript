// Package compiler defines opcodes and helpers for bytecode instructions.
package compiler

import (
	"fmt"
	"strings" // Added for String() method in Instructions

	"github.com/SethGK/Inscript/internal/types"
)

// Opcode is a single byte code that identifies the operation.
type Opcode byte

// Opcode definitions (iota starts at 0 and increments)
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
	OpJumpTruthy

	OpNull // explicit nil/None

	OpGetGlobal
	OpSetGlobal
	OpGetLocal
	OpSetLocal
	OpGetFree
	OpSetFree

	OpArray // create array literal

	OpCall
	OpReturnValue
	OpReturn

	OpPrint

	// Iterator opcodes
	OpGetIter
	OpIterNext

	// New Opcodes (continue sequence from OpIterNext)
	OpPow
	OpIDiv
	OpBitAnd
	OpBitOr
	OpBitXor
	OpShl
	OpShr
	OpGreaterEqual
	OpLessEqual
	OpBitNot
	OpClosure
	OpIndex
	OpSetIndex
	OpTable
	OpImport
)

// Instruction widths by opcode: number and byte-width of each operand.
// This map MUST be declared AFTER the Opcode constants are defined using iota.
var operandWidths = map[Opcode][]int{
	OpConstant:      {2}, // constant pool index (uint16)
	OpPop:           {},  // Added: no operands
	OpTrue:          {},  // Added: no operands
	OpFalse:         {},  // Added: no operands
	OpEqual:         {},  // Added: no operands
	OpNotEqual:      {},  // Added: no operands
	OpGreaterThan:   {},  // Added: no operands
	OpLessThan:      {},  // Added: no operands
	OpBang:          {},  // Added: no operands (logical NOT)
	OpMinus:         {},  // Added: no operands (unary negation)
	OpAdd:           {},  // Added: no operands
	OpSub:           {},  // Added: no operands
	OpMul:           {},  // Added: no operands
	OpDiv:           {},  // Added: no operands
	OpMod:           {},  // Added: no operands
	OpJumpNotTruthy: {2}, // jump offset
	OpJump:          {2}, // jump offset
	OpJumpTruthy:    {2}, // jump offset for 'or'
	OpNull:          {},  // Added: no operands
	OpGetGlobal:     {2}, // global name index
	OpSetGlobal:     {2}, // global name index
	OpGetLocal:      {1}, // local slot index
	OpSetLocal:      {1}, // local slot index
	OpGetFree:       {1}, // free slot index
	OpSetFree:       {1}, // free slot index
	OpArray:         {2}, // Corrected: Number of elements (uint16)
	OpCall:          {1}, // argument count
	OpReturnValue:   {},  // Added: no operands
	OpReturn:        {},  // Added: no operands
	OpPrint:         {1}, // number of expressions
	OpGetIter:       {},  // no operands
	OpIterNext:      {2}, // exit jump offset

	// New Opcodes and their operand widths
	OpPow:          {},     // no operands
	OpIDiv:         {},     // no operands
	OpBitAnd:       {},     // no operands
	OpBitOr:        {},     // no operands
	OpBitXor:       {},     // no operands
	OpShl:          {},     // no operands
	OpShr:          {},     // no operands
	OpGreaterEqual: {},     // no operands
	OpLessEqual:    {},     // no operands
	OpBitNot:       {},     // no operands
	OpClosure:      {2, 1}, // constant pool index (uint16), free variable count (uint8)
	OpIndex:        {},     // no operands (pops aggregate, index)
	OpSetIndex:     {},     // no operands (pops aggregate, index, value)
	OpTable:        {2},    // number of key-value pairs (uint16)
	OpImport:       {2},    // string constant index for path (uint16)
}

// Instructions is a slice of bytecode instructions.
type Instructions []byte

// Make constructs an instruction byte-slice for op and operands.
func Make(op Opcode, operands ...int) []byte {
	widths, ok := operandWidths[op]
	if !ok {
		panic(fmt.Sprintf("unknown opcode %d (%s) in Make function", op, Opcode(op).String()))
	}

	size := 1
	for _, w := range widths {
		size += w
	}

	ins := make([]byte, size)
	ins[0] = byte(op)
	offset := 1

	if len(operands) != len(widths) {
		panic(fmt.Sprintf("mismatched operand count for opcode %s: expected %d, got %d",
			Opcode(op).String(), len(widths), len(operands)))
	}

	for i, operand := range operands {
		width := widths[i] // This will now be safe as len(operands) matches len(widths)
		switch width {
		case 2:
			ins[offset] = byte(operand >> 8)
			ins[offset+1] = byte(operand)
		case 1:
			ins[offset] = byte(operand)
		case 0:
			// no operand, do nothing
		default:
			panic(fmt.Sprintf("unsupported operand width %d for opcode %s", width, Opcode(op).String()))
		}
		offset += width
	}

	return ins
}

// ReadOperand reads a single operand of given byte-width and returns the value and bytes consumed.
func ReadOperand(ins Instructions, offset, width int) (int, int) {
	switch width {
	case 2:
		val := int16(ins[offset])<<8 | int16(ins[offset+1])
		return int(val), 2
	case 1:
		return int(ins[offset]), 1
	case 0:
		return 0, 0
	default:
		panic(fmt.Sprintf("unsupported operand width %d", width))
	}
}

// ReadOperands decodes all operands for op starting at offset.
func ReadOperands(op Opcode, ins Instructions, offset int) ([]int, int) {
	widths, ok := operandWidths[op]
	if !ok {
		// This case should ideally not be hit if the bytecode is valid
		// and all opcodes are defined in operandWidths.
		return []int{}, 0
	}

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
	var out strings.Builder
	i := 0
	for i < len(ins) {
		op := ReadOpcode(ins, i)

		// Check if opcode is defined before trying to read operands
		// This ensures we don't panic on an unknown opcode in operandWidths map
		_, ok := operandWidths[op]
		if !ok {
			fmt.Fprintf(&out, "%04d ERROR: unknown opcode %d\n", i, op)
			i++ // Advance to avoid infinite loop on unknown opcode
			continue
		}

		operands, bytesRead := ReadOperands(op, ins, i+1)

		// Use op.String() directly for the opcode name
		fmt.Fprintf(&out, "%04d %s", i, op.String())
		for _, operand := range operands {
			fmt.Fprintf(&out, " %d", operand)
		}
		fmt.Fprintln(&out)

		i += 1 + bytesRead
	}
	return out.String()
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
	case OpJumpTruthy:
		return "OpJumpTruthy"
	case OpNull:
		return "OpNull"
	case OpGetGlobal:
		return "OpGetGlobal"
	case OpSetGlobal:
		return "OpSetGlobal"
	case OpGetLocal:
		return "OpGetLocal"
	case OpSetLocal:
		return "OpSetLocal"
	case OpGetFree:
		return "OpGetFree"
	case OpSetFree:
		return "OpSetFree"
	case OpArray:
		return "OpArray"
	case OpCall:
		return "OpCall"
	case OpReturnValue:
		return "OpReturnValue"
	case OpReturn:
		return "OpReturn"
	case OpPrint:
		return "OpPrint"
	case OpGetIter:
		return "OpGetIter"
	case OpIterNext:
		return "OpIterNext"
	case OpPow:
		return "OpPow"
	case OpIDiv:
		return "OpIDiv"
	case OpBitAnd:
		return "OpBitAnd"
	case OpBitOr:
		return "OpBitOr"
	case OpBitXor:
		return "OpBitXor"
	case OpShl:
		return "OpShl"
	case OpShr:
		return "OpShr"
	case OpGreaterEqual:
		return "OpGreaterEqual"
	case OpLessEqual:
		return "OpLessEqual"
	case OpBitNot:
		return "OpBitNot"
	case OpClosure:
		return "OpClosure"
	case OpIndex:
		return "OpIndex"
	case OpSetIndex:
		return "OpSetIndex"
	case OpTable:
		return "OpTable"
	case OpImport:
		return "OpImport"
	default:
		return fmt.Sprintf("Opcode(%d)", op)
	}
}

// Bytecode holds compiled instructions and metadata.
type Bytecode struct {
	Instructions  Instructions
	Constants     []types.Value // use types.Value for constant pool
	NumLocals     int
	NumParameters int
	NumGlobals    int
}
