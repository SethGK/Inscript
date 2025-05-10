package vm

import (
	"fmt"

	"github.com/SethGK/Inscript/internal/compiler" // Import the compiler package to use Bytecode and OpCode
	// You will need to import your runtime value types here
	// "github.com/SethGK/Inscript/internal/vm/value"
)

// VM represents the virtual machine that executes bytecode.
type VM struct {
	constants    []interface{}          // The constant pool from the compiled bytecode
	instructions []compiler.Instruction // The instructions to execute

	stack []interface{} // The main data stack
	sp    int           // Stack pointer: points to the next available slot on the stack

	globals []interface{} // Global variables

	// Frame management (for function calls)
	// frames []*Frame // Call frame stack
	// framesIndex int // Current frame pointer

	// Other potential fields: builtins, error handling, etc.
}

// New creates a new VM instance with the given bytecode and global variable count.
func New(bytecode *compiler.Bytecode, numGlobals int) *VM {
	// Initialize globals with nil values based on the number of global variables defined by the compiler.
	globals := make([]interface{}, numGlobals)

	vm := &VM{
		constants:    bytecode.Constants,
		instructions: bytecode.Instructions,
		stack:        make([]interface{}, 2048), // Arbitrary stack size for now
		sp:           0,                         // Stack starts empty

		globals: globals,

		// Initialize frames (placeholder)
		// frames: make([]*Frame, 1024), // Arbitrary frame stack size
		// framesIndex: 1, // Start with the main frame at index 0 (implicitly handled by initial VM state)
	}
	// The initial state of the VM implicitly represents the "main" frame executing the program bytecode.
	// A dedicated Frame struct would be needed when implementing actual function calls.

	return vm
}

// Run executes the loaded bytecode.
// It returns the value left on the stack after execution (if any) or an error.
func (vm *VM) Run() (interface{}, error) {
	// The main fetch-decode-execute loop
	for ip := 0; ip < len(vm.instructions); ip++ {
		instr := vm.instructions[ip]
		op := instr.Op

		// Decode and execute the instruction
		switch op {
		case compiler.OpConstant:
			// OpConstant operand is the index into the constant pool.
			constantIndex := instr.Operands[0]
			// Push the constant value onto the stack.
			vm.push(vm.constants[constantIndex])

		case compiler.OpAdd:
			// OpAdd pops two operands, adds them, and pushes the result.
			right := vm.pop()
			left := vm.pop()

			// Perform the addition. This requires type checking and handling.
			// For simplicity, assume they are both integers for now.
			// You will need proper runtime value types and arithmetic logic later.
			leftVal, okLeft := left.(int64) // Assuming int64 for IntegerLiteral
			rightVal, okRight := right.(int64)
			if !okLeft || !okRight {
				// TODO: Handle type error (e.g., adding non-integers)
				return nil, fmt.Errorf("type error: cannot add non-integers")
			}
			result := leftVal + rightVal
			vm.push(result) // Push the result back as a Go int64

		case compiler.OpSub:
			// OpSub pops two operands, subtracts right from left, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			// TODO: Implement subtraction with type checking
			leftVal, okLeft := left.(int64)
			rightVal, okRight := right.(int64)
			if !okLeft || !okRight {
				return nil, fmt.Errorf("type error: cannot subtract non-integers")
			}
			result := leftVal - rightVal
			vm.push(result)

		case compiler.OpMul:
			// OpMul pops two operands, multiplies them, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			// TODO: Implement multiplication with type checking
			leftVal, okLeft := left.(int64)
			rightVal, okRight := right.(int64)
			if !okLeft || !okRight {
				return nil, fmt.Errorf("type error: cannot multiply non-integers")
			}
			result := leftVal * rightVal
			vm.push(result)

		case compiler.OpDiv:
			// OpDiv pops two operands, divides left by right, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			// TODO: Implement division with type checking and division by zero handling
			leftVal, okLeft := left.(int64)
			rightVal, okRight := right.(int64)
			if !okLeft || !okRight {
				return nil, fmt.Errorf("type error: cannot divide non-integers")
			}
			if rightVal == 0 {
				return nil, fmt.Errorf("runtime error: division by zero")
			}
			result := leftVal / rightVal // Integer division for now
			vm.push(result)

		case compiler.OpPop:
			// OpPop simply removes the top element from the stack.
			vm.pop()

		case compiler.OpNull:
			// OpNull pushes a null value onto the stack.
			vm.push(nil) // Using Go's nil for now

		case compiler.OpTrue:
			// OpTrue pushes a boolean true onto the stack.
			vm.push(true) // Using Go's bool for now

		case compiler.OpFalse:
			// OpFalse pushes a boolean false onto the stack.
			vm.push(false) // Using Go's bool for now

		case compiler.OpPrint:
			// OpPrint operand is the number of arguments to print.
			numArgs := instr.Operands[0]
			// Pop the arguments from the stack and print them.
			args := make([]interface{}, numArgs)
			for i := numArgs - 1; i >= 0; i-- {
				args[i] = vm.pop()
			}
			// Print the arguments separated by spaces, followed by a newline.
			fmt.Println(args...) // fmt.Println handles multiple arguments

		case compiler.OpSetGlobal:
			// OpSetGlobal operand is the index of the global variable.
			globalIndex := instr.Operands[0]
			// Pop the value from the stack and store it in the global variable slot.
			value := vm.pop()
			vm.globals[globalIndex] = value

		case compiler.OpGetGlobal:
			// OpGetGlobal operand is the index of the global variable.
			globalIndex := instr.Operands[0]
			// Get the value from the global variable slot and push it onto the stack.
			vm.push(vm.globals[globalIndex])

		// TODO: Implement other opcodes:
		// OpMod, OpPow, OpMinus, OpNot
		// OpEqual, OpNotEqual, OpGreaterThan, OpLessThan, OpGreaterEqual, OpLessEqual
		// OpJump, OpJumpNotTruthy, OpJumpTruthy (need to adjust ip)
		// OpGetLocal, OpSetLocal (need frame management)
		// OpArray, OpHash, OpIndex, OpSetIndex
		// OpGetIterator, OpIteratorNext (need iterator implementation)
		// OpCall, OpReturnValue, OpReturn (need frame management)

		default:
			// Handle unknown opcodes (compiler bug or corrupted bytecode)
			return nil, fmt.Errorf("unknown opcode: %d", op)
		}
	}

	// After the loop finishes (or hits an OpReturn from the main program),
	// the result of the program execution is typically the value left on the stack.
	// For a simple script, this might be the result of the last expression statement.
	// If the stack is empty, the result is implicitly null.
	if vm.sp == 0 {
		return nil, nil // Stack is empty, return nil
	}
	return vm.pop(), nil // Return the top value from the stack
}

// push pushes a value onto the stack.
func (vm *VM) push(value interface{}) {
	if vm.sp >= len(vm.stack) {
		// TODO: Handle stack overflow
		panic("stack overflow")
	}
	vm.stack[vm.sp] = value
	vm.sp++
}

// pop pops a value from the stack.
func (vm *VM) pop() interface{} {
	if vm.sp == 0 {
		// TODO: Handle stack underflow
		panic("stack underflow")
	}
	vm.sp--
	value := vm.stack[vm.sp]
	// Optional: zero out the stack slot to help garbage collection
	vm.stack[vm.sp] = nil
	return value
}

// peek returns the value at the top of the stack without popping it.
func (vm *VM) peek(distance int) interface{} {
	if vm.sp-1-distance < 0 {
		// TODO: Handle stack underflow
		panic("stack underflow")
	}
	return vm.stack[vm.sp-1-distance]
}

// TODO: Define runtime value types in internal/vm/value.go
// This will replace interface{} with concrete types like IntValue, StringValue, etc.
