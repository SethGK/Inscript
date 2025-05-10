package compiler

import (
	"fmt"
	"strings"
	// You might need to import your runtime value types here later
	// "your_module_path/internal/vm/value"
)

// Instruction represents a single bytecode instruction.
// In a real high-performance VM, this would likely be just a sequence of bytes
// for efficiency. For development and debugging, a struct is easier to work with.
type Instruction struct {
	Op       OpCode // The operation code
	Operands []int  // Operands for the instruction (e.g., constant index, jump target, local index)
	// The interpretation of Operands depends on the OpCode.
	// In a production compiler, operands would be encoded into bytes based on their size (OperandWidths).
}

// Bytecode represents the compiled result of a compilation unit (program or function body).
// It contains the sequence of instructions and the pool of constants used by those instructions.
type Bytecode struct {
	Instructions  []Instruction // The sequence of bytecode instructions
	Constants     []interface{} // The constant pool (literals, function objects, etc.)
	NumLocals     int           // Number of local variables needed for this compilation unit (for function frames)
	NumParameters int           // Number of parameters (for function frames) - only relevant for function bytecode
	// NumGlobals int // Might be added here or managed by the VM based on global symbol table size
}

// NewBytecode creates a new, empty Bytecode struct.
func NewBytecode() *Bytecode {
	return &Bytecode{
		Instructions:  []Instruction{},
		Constants:     []interface{}{},
		NumLocals:     0, // Default to 0
		NumParameters: 0, // Default to 0
	}
}

// AddInstruction appends an instruction to the bytecode's instruction slice.
// It returns the starting position (index) of the newly added instruction.
// This position is useful for patching jumps later.
func (b *Bytecode) AddInstruction(op OpCode, operands ...int) int {
	instr := Instruction{Op: op, Operands: operands}
	pos := len(b.Instructions)
	b.Instructions = append(b.Instructions, instr)
	return pos
}

// ReplaceInstructionAt patches an existing instruction at a given position.
// This is primarily used by the compiler to fill in jump targets
// after the target instruction's position is known.
func (b *Bytecode) ReplaceInstructionAt(pos int, op OpCode, operands ...int) {
	// Basic validation for the provided position.
	if pos < 0 || pos >= len(b.Instructions) {
		fmt.Printf("ERROR: Attempted to patch instruction at invalid position %d\n", pos)
		// In a real compiler, this might be a fatal error or added to an error list.
		return
	}

	// Optional: Basic check to ensure the new instruction's operand count
	// matches the definition for the opcode. This helps catch compiler bugs.
	def, ok := Lookup(op) // Lookup needs access to definitions from opcode.go
	if !ok || len(operands) != len(def.OperandWidths) {
		fmt.Printf("WARNING: Patching instruction %d with mismatched operands for opcode %v (Expected %d, Got %d)\n",
			pos, op, len(def.OperandWidths), len(operands))
		// Continue with the patch, but warn.
	}

	b.Instructions[pos] = Instruction{Op: op, Operands: operands}
}

// AddConstant adds a constant value to the constant pool.
// Returns the index of the added constant.
// TODO: Implement checking for duplicate constants to avoid adding the same value multiple times (optimization).
func (b *Bytecode) AddConstant(obj interface{}) int {
	b.Constants = append(b.Constants, obj)
	return len(b.Constants) - 1 // Return the index of the newly added constant
}

// FormatInstructions converts the bytecode instructions into a human-readable string format.
// This is invaluable for debugging the compiler and understanding the generated code.
func (b *Bytecode) FormatInstructions() string {
	var s strings.Builder // Use strings.Builder for efficient string concatenation
	for i, instr := range b.Instructions {
		// Format each instruction and append it to the string builder.
		s.WriteString(fmt.Sprintf("%04d %s\n", i, formatInstruction(&instr))) // %04d formats the index with leading zeros
	}
	return s.String()
}

// formatInstruction is a helper function to format a single instruction.
// It uses the opcode definitions to print the instruction name and its operands.
func formatInstruction(instr *Instruction) string {
	// Look up the definition for the opcode to get its name and operand information.
	def, ok := Lookup(instr.Op) // Lookup needs access to definitions from opcode.go
	if !ok {
		return fmt.Sprintf("ERROR: unknown opcode %d", instr.Op)
	}

	// Check if the number of provided operands matches the definition.
	operandCount := len(def.OperandWidths)
	if len(instr.Operands) != operandCount {
		return fmt.Sprintf("ERROR: opcode %s has %d operands but received %d",
			def.Name, operandCount, len(instr.Operands))
	}

	// Format the operands.
	operandsStr := ""
	if operandCount > 0 {
		// For simplicity, we just print the integer values of the operands.
		// A more advanced formatter might look up constant values from the pool
		// or symbol names from the symbol table based on the opcode type.
		operandsStr = fmt.Sprintf("%v", instr.Operands)
		operandsStr = strings.Trim(operandsStr, "[]") // Remove the brackets from the slice printout
	}

	// Combine the opcode name and formatted operands.
	return fmt.Sprintf("%s %s", def.Name, operandsStr)
}

// Note: If you implement operand encoding into []byte, you would add functions here
// to read operands from a byte slice based on the opcode definition.
