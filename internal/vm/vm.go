package vm

import (
	"fmt"
	"math"

	"github.com/SethGK/Inscript/internal/compiler" // Import the compiler package
	// Removed: "github.com/SethGK/Inscript/internal/vm/value" // Import the new value package
)

// Frame represents a function call frame.
type Frame struct {
	// Function *value.Function // The function being executed (need Function value type)
	bytecode *compiler.Bytecode // The bytecode for this frame (program or function)
	ip       int                // Instruction pointer: current instruction index in this frame's bytecode
	basePtr  int                // Base pointer: stack index where locals for this frame start
	// ReturnAddr int // Instruction pointer in the caller's frame to return to (needed for OpCall/OpReturn)
}

// NewFrame creates a new Frame for executing bytecode.
// For the main program, bytecode is the program's bytecode, basePtr is 0.
// For a function call, bytecode is the function's bytecode, basePtr is where args/locals start on the stack.
func NewFrame(bytecode *compiler.Bytecode, basePtr int) *Frame {
	return &Frame{
		bytecode: bytecode,
		ip:       0, // Start at the first instruction of the frame's bytecode
		basePtr:  basePtr,
		// ReturnAddr: 0, // Placeholder
	}
}

// VM represents the virtual machine that executes bytecode.
type VM struct {
	// Constants and instructions are now part of the current Frame's bytecode.
	// constants    []Value          // The constant pool from the compiled bytecode (now in Frame)
	// instructions []compiler.Instruction // The instructions to execute (now in Frame)

	stack []Value // The main data stack (now using Value)
	sp    int     // Stack pointer: points to the next available slot on the stack

	globals []Value // Global variables (now using Value)

	// Frame management
	frames []*Frame // Call frame stack
	// framesIndex int // Current frame pointer (index into frames slice)
	currentFrame *Frame // Pointer to the current active frame
}

// New creates a new VM instance with the given program bytecode and number of globals.
func New(programBytecode *compiler.Bytecode, numGlobals int) *VM {
	// Convert constants from interface{} to Value for the program bytecode.
	// The compiler stores constants as interface{}, the VM needs Value types.
	programConstants := make([]Value, len(programBytecode.Constants))
	for i, c := range programBytecode.Constants {
		// This conversion logic needs to handle all possible types in the constant pool
		// based on what the compiler puts there.
		switch v := c.(type) {
		case int64:
			programConstants[i] = &Integer{Value: v}
		case float64:
			programConstants[i] = &Float{Value: v}
		case string:
			programConstants[i] = &String{Value: v}
		case bool:
			programConstants[i] = GoBoolToBoolean(v)
		case nil: // Compiler might put Go nil for ast.NilLiteral
			programConstants[i] = NULL
		// TODO: Add cases for other constant types from compiler (e.g., function objects)
		default:
			// This indicates a compiler bug or unsupported constant type
			// Use fmt.Errorf and handle it, or panic for now.
			panic(fmt.Sprintf("unsupported constant type in VM: %T", v)) // Or return error
		}
	}
	// Replace the bytecode's constants with the converted Value slice.
	// This is a bit of a hack; ideally, the compiler would store Value types directly.
	// But given the current structure, this works for now.
	// NOTE: The Bytecode struct's Constants field is []interface{}, so we can't
	// directly assign programConstants ([]Value) to it.
	// Instead, we'll store the converted constants separately in the Frame or VM,
	// or modify the Bytecode struct to hold []Value.
	// Let's modify the Frame struct to hold []Value constants for simplicity now.

	// Create the initial frame for the main program.
	// The main program's locals start at stack index 0 (though it won't have locals in this design).
	// The main program's bytecode is the initial bytecode.
	// Pass the *original* bytecode to the frame, but the VM will use the converted constants.
	mainFrame := NewFrame(programBytecode, 0) // Base pointer for main frame is 0

	// Initialize globals with Null values based on the number of global variables.
	globals := make([]Value, numGlobals)
	for i := range globals {
		globals[i] = NULL // Initialize global slots to nil
	}

	vm := &VM{
		// constants:    constants, // Constants are in the frame's bytecode
		// instructions: bytecode.Instructions, // Instructions are in the frame's bytecode
		stack: make([]Value, 2048), // Arbitrary stack size for now (using Value)
		sp:    0,                   // Stack starts empty

		globals: globals,

		// Initialize frames
		frames: []*Frame{mainFrame}, // Start with the main frame on the stack
		// framesIndex: 0, // Index of the current frame (implicitly the last one in the slice)
		currentFrame: mainFrame, // Set the current frame to the main frame
	}

	// Store the converted constants in the VM or Frame for easy access during execution.
	// Let's store them in the Frame for now, as constants are per bytecode unit.
	// This requires adding a Constants field to the Frame struct.
	// For now, let's access them via the bytecode but perform the conversion on the fly
	// in the OpConstant instruction, or better, convert them once per bytecode.
	// Let's revert the assumption that Bytecode.Constants is []Value and fix the VM.Run loop.

	return vm
}

// Run executes the loaded bytecode.
// It returns the value left on the stack after execution (if any) or an error.
func (vm *VM) Run() (Value, error) { // Return Value
	// The main fetch-decode-execute loop
	// ip (instruction pointer) is now managed within the current frame.
	// Access current frame's bytecode and instruction pointer.
	// These variables are updated if the currentFrame changes (e.g., on OpCall/OpReturn)
	var instructions []compiler.Instruction
	var constants []interface{} // Constants are []interface{} in the Bytecode struct

	for { // Use an infinite loop and break on OpReturn from main frame or error
		// Update instructions and constants to point to the current frame's bytecode
		instructions = vm.currentFrame.bytecode.Instructions
		constants = vm.currentFrame.bytecode.Constants // Access constants as []interface{}

		// Check if we've gone past the end of the current frame's instructions
		if vm.currentFrame.ip >= len(instructions) {
			// Reached the end of the current frame's bytecode.
			// If this is the main frame, the program finishes.
			// If this is a function frame, it's an implicit return null.
			if len(vm.frames) == 1 && vm.currentFrame == vm.frames[0] {
				// End of main program bytecode.
				break // Exit the execution loop
			}
			// TODO: Implement actual function return logic (pop frame, restore caller state, push NULL).
			return nil, fmt.Errorf("implicit return in non-main frame (function calls not fully implemented)")
		}

		// Get the current instruction using the frame's instruction pointer.
		instr := instructions[vm.currentFrame.ip]
		op := instr.Op

		// Increment instruction pointer *before* executing, in case of jumps
		vm.currentFrame.ip++

		// Decode and execute the instruction
		switch op {
		case compiler.OpConstant:
			// OpConstant operand is the index into the constant pool.
			constantIndex := instr.Operands[0]
			// Get the constant value from the *current frame's* constant pool (which is []interface{}).
			constantInterface := constants[constantIndex]
			// Convert the interface{} value to a Value type before pushing it onto the stack.
			constantValue, err := vm.convertInterfaceToValue(constantInterface) // Use a helper function
			if err != nil {
				return nil, fmt.Errorf("failed to convert constant to Value: %w", err)
			}
			vm.push(constantValue)

		case compiler.OpAdd:
			// OpAdd pops two operands, adds them, and pushes the result.
			right := vm.pop()
			left := vm.pop()

			// Perform the addition using value types and type-specific logic.
			result, err := vm.add(left, right)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpSub:
			// OpSub pops two operands, subtracts right from left, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			result, err := vm.subtract(left, right)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpMul:
			// OpMul pops two operands, multiplies them, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			result, err := vm.multiply(left, right)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpDiv:
			// OpDiv pops two operands, divides left by right, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			result, err := vm.divide(left, right)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpMod:
			// OpMod pops two operands, performs modulo, and pushes the result.
			right := vm.pop()
			left := vm.pop()
			result, err := vm.modulo(left, right)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpPow:
			// OpPow pops two operands (base, exp), performs power, and pushes the result.
			right := vm.pop() // Exponent
			left := vm.pop()  // Base
			result, err := vm.power(left, right)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpMinus:
			// OpMinus pops one operand, negates it, and pushes the result.
			operand := vm.pop()
			result, err := vm.negate(operand)
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpNot:
			// OpNot pops one operand, performs logical NOT, and pushes the boolean result.
			operand := vm.pop()
			// The result of 'not' is the boolean opposite of the operand's truthiness.
			vm.push(GoBoolToBoolean(!IsTruthy(operand)))

		case compiler.OpEqual:
			right := vm.pop()
			left := vm.pop()
			// Equality comparison
			vm.push(GoBoolToBoolean(vm.isEqual(left, right)))

		case compiler.OpNotEqual:
			right := vm.pop()
			left := vm.pop()
			// Inequality comparison
			vm.push(GoBoolToBoolean(!vm.isEqual(left, right)))

		case compiler.OpGreaterThan:
			right := vm.pop()
			left := vm.pop()
			result, err := vm.compare(left, right, ">")
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpLessThan:
			right := vm.pop()
			left := vm.pop()
			result, err := vm.compare(left, right, "<")
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpGreaterEqual:
			right := vm.pop()
			left := vm.pop()
			result, err := vm.compare(left, right, ">=")
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpLessEqual:
			right := vm.pop()
			left := vm.pop()
			result, err := vm.compare(left, right, "<=")
			if err != nil {
				return nil, err
			}
			vm.push(result)

		case compiler.OpJump:
			// OpJump operand is the absolute target instruction index within the *current frame's* bytecode.
			// Set the current frame's instruction pointer to the target address.
			jumpTarget := instr.Operands[0]
			vm.currentFrame.ip = jumpTarget // ip will be incremented at the start of the next loop iteration

		case compiler.OpJumpNotTruthy:
			// OpJumpNotTruthy operand is the absolute target instruction index within the *current frame's* bytecode.
			// Pop the condition value.
			condition := vm.pop()
			// If the condition is falsy, set the current frame's instruction pointer to the target address.
			jumpTarget := instr.Operands[0]
			if !IsTruthy(condition) {
				vm.currentFrame.ip = jumpTarget // ip will be incremented at the start of the next loop iteration
			}
			// If truthy, continue to the next instruction (ip is already incremented)

		case compiler.OpJumpTruthy:
			// OpJumpTruthy operand is the absolute target instruction index within the *current frame's* bytecode.
			// Pop the condition value.
			condition := vm.pop()
			// If the condition is truthy, set the current frame's instruction pointer to the target address.
			jumpTarget := instr.Operands[0]
			if IsTruthy(condition) {
				vm.currentFrame.ip = jumpTarget // ip will be incremented at the start of the next loop iteration
			}
			// If falsy, continue to the next instruction (ip is already incremented)

		case compiler.OpPop:
			// OpPop simply removes the top element from the stack.
			vm.pop()

		case compiler.OpNull:
			// OpNull pushes the Null singleton onto the stack.
			vm.push(NULL)

		case compiler.OpTrue:
			// OpTrue pushes the True singleton onto the stack.
			vm.push(TRUE)

		case compiler.OpFalse:
			// OpFalse pushes the False singleton onto the stack.
			vm.push(FALSE)

		case compiler.OpPrint:
			// OpPrint operand is the number of arguments to print.
			numArgs := instr.Operands[0]
			// Pop the arguments from the stack and print them using their Inspect() method.
			args := make([]interface{}, numArgs) // Use interface{} for fmt.Println variadic args
			for i := numArgs - 1; i >= 0; i-- {
				// Pop the value and get its string representation
				args[i] = vm.pop().Inspect()
			}
			// Print the arguments separated by spaces, followed by a newline.
			fmt.Println(args...) // fmt.Println handles multiple arguments

		case compiler.OpSetGlobal:
			// OpSetGlobal operand is the index of the global variable.
			globalIndex := instr.Operands[0]
			// Pop the value from the stack and store it in the global variable slot.
			vm.globals[globalIndex] = vm.pop()

		case compiler.OpGetGlobal:
			// OpGetGlobal operand is the index of the global variable.
			globalIndex := instr.Operands[0]
			// Get the value from the global variable slot and push it onto the stack.
			vm.push(vm.globals[globalIndex])

		case compiler.OpSetLocal:
			// OpSetLocal operand is the index of the local variable within the current frame.
			localIndex := instr.Operands[0]
			// The stack index for this local is the frame's base pointer + local index.
			stackIndex := vm.currentFrame.basePtr + localIndex
			// Ensure the stack is large enough for this index.
			if stackIndex >= len(vm.stack) {
				// This indicates a compiler bug or stack overflow.
				// Consider growing the stack or returning an error.
				panic(fmt.Sprintf("stack index out of bounds for local variable assignment: %d", stackIndex)) // Or return error
			}
			// Pop the value from the stack and store it in the local variable slot on the stack.
			vm.stack[stackIndex] = vm.pop()

		case compiler.OpGetLocal:
			// OpGetLocal operand is the index of the local variable within the current frame.
			localIndex := instr.Operands[0]
			// The stack index for this local is the frame's base pointer + local index.
			stackIndex := vm.currentFrame.basePtr + localIndex
			// Ensure the stack is large enough for this index.
			if stackIndex >= len(vm.stack) {
				// This indicates a compiler bug or accessing uninitialized local.
				// Consider returning an error for accessing undefined local.
				panic(fmt.Sprintf("stack index out of bounds for local variable access: %d", stackIndex)) // Or return error
			}
			// Get the value from the local variable slot on the stack and push it onto the stack.
			vm.push(vm.stack[stackIndex])

		case compiler.OpReturn:
			// OpReturn from the main program indicates the end of execution.
			// For function calls, this will involve popping frames.
			// For the main frame, just break the loop.
			if len(vm.frames) == 1 && vm.currentFrame == vm.frames[0] {
				// This is the main program frame.
				break // Exit the execution loop
			}
			// TODO: Implement actual function return logic (pop frame, restore caller state, push return value).
			return nil, fmt.Errorf("OpReturn encountered in non-main frame (function calls not fully implemented)")

		case compiler.OpReturnValue:
			// OpReturnValue pops the return value and returns from the function.
			// TODO: Implement actual function return logic (pop frame, restore caller state, push return value).
			return nil, fmt.Errorf("OpReturnValue encountered (function calls not fully implemented)")

		// TODO: Implement other opcodes:
		// OpArray, OpHash, OpIndex, OpSetIndex
		// OpGetIterator, OpIteratorNext (need iterator implementation)
		// OpCall (need frame management)

		default:
			// Handle unknown opcodes (compiler bug or corrupted bytecode)
			return nil, fmt.Errorf("unknown opcode: %d at instruction %d", op, vm.currentFrame.ip-1) // ip-1 because we incremented already
		}
	}

	// After the loop finishes (either by reaching the end of instructions or hitting OpReturn from main frame),
	// the result of the program execution is typically the value left on the stack.
	// For a simple script, this might be the result of the last expression statement.
	// If the stack is empty, the result is implicitly null.
	if vm.sp == 0 {
		return NULL, nil // Stack is empty, return Null
	}
	return vm.pop(), nil // Return the top value from the stack
}

// convertInterfaceToValue converts an interface{} value from the constant pool
// into a VM Value type.
func (vm *VM) convertInterfaceToValue(val interface{}) (Value, error) {
	switch v := val.(type) {
	case int64:
		return &Integer{Value: v}, nil
	case float64:
		return &Float{Value: v}, nil
	case string:
		return &String{Value: v}, nil
	case bool:
		return GoBoolToBoolean(v), nil
	case nil:
		return NULL, nil
	// TODO: Add cases for other constant types from compiler (e.g., function objects)
	default:
		// This indicates an unsupported constant type from the compiler.
		return nil, fmt.Errorf("unsupported constant type in VM: %T", v)
	}
}

// push pushes a value onto the stack.
func (vm *VM) push(val Value) { // Accepts Value
	if vm.sp >= len(vm.stack) {
		// TODO: Handle stack overflow more gracefully
		// panic("stack overflow") // Or return an error
		// Returning error is better for a robust VM
		// For now, let's just grow the stack (less efficient but avoids panic)
		newStack := make([]Value, len(vm.stack)*2) // Double stack size
		copy(newStack, vm.stack)
		vm.stack = newStack
		// fmt.Println("Stack grown to", len(vm.stack)) // For debugging
	}
	vm.stack[vm.sp] = val
	vm.sp++
}

// pop pops a value from the stack.
func (vm *VM) pop() Value { // Returns Value
	if vm.sp == 0 {
		// TODO: Handle stack underflow
		panic("stack underflow") // Or return an error
	}
	vm.sp--
	val := vm.stack[vm.sp]
	// Optional: zero out the stack slot to help garbage collection
	vm.stack[vm.sp] = nil // Set to nil to release reference
	return val
}

// peek returns the value at the top of the stack without popping it.
// distance 0 is the top, distance 1 is the element below the top, etc.
func (vm *VM) peek(distance int) Value { // Returns Value
	if vm.sp-1-distance < 0 {
		// TODO: Handle stack underflow
		panic("stack underflow") // Or return an error
	}
	return vm.stack[vm.sp-1-distance]
}

// Helper methods for arithmetic operations using value types (already implemented)
func (vm *VM) add(left, right Value) (Value, error) {
	// ... (existing add logic) ...
	switch left := left.(type) {
	case *Integer:
		switch right := right.(type) {
		case *Integer:
			return &Integer{Value: left.Value + right.Value}, nil
		case *Float:
			return &Float{Value: float64(left.Value) + right.Value}, nil
		default:
			return nil, fmt.Errorf("type error: cannot add Integer and %s", right.Type())
		}
	case *Float:
		switch right := right.(type) {
		case *Float:
			return &Float{Value: left.Value + right.Value}, nil
		case *Integer:
			return &Float{Value: left.Value + float64(right.Value)}, nil
		default:
			return nil, fmt.Errorf("type error: cannot add Float and %s", right.Type())
		}
	case *String:
		// String concatenation with other types (convert right to string)
		return &String{Value: left.Value + right.Inspect()}, nil
	default:
		return nil, fmt.Errorf("type error: cannot add %s and %s", left.Type(), right.Type())
	}
}

func (vm *VM) subtract(left, right Value) (Value, error) {
	// ... (existing subtract logic) ...
	switch left := left.(type) {
	case *Integer:
		switch right := right.(type) {
		case *Integer:
			return &Integer{Value: left.Value - right.Value}, nil
		case *Float:
			return &Float{Value: float64(left.Value) - right.Value}, nil
		default:
			return nil, fmt.Errorf("type error: cannot subtract Integer and %s", right.Type())
		}
	case *Float:
		switch right := right.(type) {
		case *Float:
			return &Float{Value: left.Value - right.Value}, nil
		case *Integer:
			return &Float{Value: left.Value - float64(right.Value)}, nil
		default:
			return nil, fmt.Errorf("type error: cannot subtract Float and %s", right.Type())
		}
	}
	return nil, fmt.Errorf("type error: cannot subtract %s and %s", left.Type(), right.Type())
}

func (vm *VM) multiply(left, right Value) (Value, error) {
	// ... (existing multiply logic) ...
	switch left := left.(type) {
	case *Integer:
		switch right := right.(type) {
		case *Integer:
			return &Integer{Value: left.Value * right.Value}, nil
		case *Float:
			return &Float{Value: float64(left.Value) * right.Value}, nil
		default:
			return nil, fmt.Errorf("type error: cannot multiply Integer and %s", right.Type())
		}
	case *Float:
		switch right := right.(type) {
		case *Float:
			return &Float{Value: left.Value * right.Value}, nil
		case *Integer:
			return &Float{Value: left.Value * float64(right.Value)}, nil
		default:
			return nil, fmt.Errorf("type error: cannot multiply Float and %s", right.Type())
		}
	}
	return nil, fmt.Errorf("type error: cannot multiply %s and %s", left.Type(), right.Type())
}

func (vm *VM) divide(left, right Value) (Value, error) {
	// ... (existing divide logic) ...
	switch left := left.(type) {
	case *Integer:
		switch right := right.(type) {
		case *Integer:
			if right.Value == 0 {
				return nil, fmt.Errorf("runtime error: division by zero")
			}
			// Decide between integer and float division based on language rules.
			// Current: Integer division if both are integers.
			return &Integer{Value: left.Value / right.Value}, nil
		case *Float:
			if right.Value == 0 {
				return nil, fmt.Errorf("runtime error: division by zero")
			}
			return &Float{Value: float64(left.Value) / right.Value}, nil
		default:
			return nil, fmt.Errorf("type error: cannot divide Integer and %s", right.Type())
		}
	case *Float:
		switch right := right.(type) {
		case *Float:
			if right.Value == 0 {
				return nil, fmt.Errorf("runtime error: division by zero")
			}
			return &Float{Value: left.Value / right.Value}, nil
		case *Integer:
			if right.Value == 0 {
				return nil, fmt.Errorf("runtime error: division by zero")
			}
			return &Float{Value: left.Value / float64(right.Value)}, nil
		default:
			return nil, fmt.Errorf("type error: cannot divide Float and %s", right.Type())
		}
	}
	return nil, fmt.Errorf("type error: cannot divide %s and %s", left.Type(), right.Type())
}

func (vm *VM) modulo(left, right Value) (Value, error) {
	// ... (existing modulo logic) ...
	// Modulo is typically only defined for integers
	leftInt, okLeft := left.(*Integer)
	rightInt, okRight := right.(*Integer)
	if !okLeft || !okRight {
		return nil, fmt.Errorf("type error: modulo operator %% is only defined for integers, got %s and %s", left.Type(), right.Type())
	}
	if rightInt.Value == 0 {
		return nil, fmt.Errorf("runtime error: modulo by zero")
	}
	return &Integer{Value: leftInt.Value % rightInt.Value}, nil
}

func (vm *VM) power(left, right Value) (Value, error) {
	// ... (existing power logic) ...
	// Power (^) can be defined for integers and floats
	leftInt, okLeftInt := left.(*Integer)
	rightInt, okRightInt := right.(*Integer)
	leftFloat, okLeftFloat := left.(*Float)
	rightFloat, okRightFloat := right.(*Float)

	// Handle mixed types by promoting to float
	if okLeftInt && okRightFloat {
		return &Float{Value: math.Pow(float64(leftInt.Value), rightFloat.Value)}, nil
	}
	if okLeftFloat && okRightInt {
		return &Float{Value: math.Pow(leftFloat.Value, float64(rightInt.Value))}, nil
	}
	// Handle same types
	if okLeftInt && okRightInt {
		// Integer power - still using float for simplicity, might need dedicated int power func
		result := math.Pow(float64(leftInt.Value), float64(rightInt.Value))
		// Consider if result should be integer if it's a whole number
		// For now, return float result of math.Pow
		return &Float{Value: result}, nil
	}
	if okLeftFloat && okRightFloat {
		return &Float{Value: math.Pow(leftFloat.Value, rightFloat.Value)}, nil
	}

	return nil, fmt.Errorf("type error: cannot perform power operation on %s and %s", left.Type(), right.Type())
}

func (vm *VM) negate(operand Value) (Value, error) {
	// ... (existing negate logic) ...
	switch operand := operand.(type) {
	case *Integer:
		return &Integer{Value: -operand.Value}, nil
	case *Float:
		return &Float{Value: -operand.Value}, nil
	default:
		return nil, fmt.Errorf("type error: cannot negate type %s", operand.Type())
	}
}

// Helper method for comparison operations (already implemented)
func (vm *VM) compare(left, right Value, op string) (Value, error) {
	// ... (existing compare logic) ...
	// Implement comparison logic based on types and the operator.
	// This will involve many type checks and comparisons.
	// For simplicity, let's only compare numbers and strings for now.
	leftInt, okLeftInt := left.(*Integer)
	rightInt, okRightInt := right.(*Integer)
	leftFloat, okLeftFloat := left.(*Float)
	rightFloat, okRightFloat := right.(*Float)
	leftStr, okLeftStr := left.(*String)
	rightStr, okRightStr := right.(*String)

	if okLeftInt && okRightInt {
		switch op {
		case ">":
			return GoBoolToBoolean(leftInt.Value > rightInt.Value), nil
		case "<":
			return GoBoolToBoolean(leftInt.Value < rightInt.Value), nil
		case ">=":
			return GoBoolToBoolean(leftInt.Value >= rightInt.Value), nil
		case "<=":
			return GoBoolToBoolean(leftInt.Value <= rightInt.Value), nil
		}
	} else if okLeftFloat && okRightFloat {
		switch op {
		case ">":
			return GoBoolToBoolean(leftFloat.Value > rightFloat.Value), nil
		case "<":
			return GoBoolToBoolean(leftFloat.Value < rightFloat.Value), nil
		case ">=":
			return GoBoolToBoolean(leftFloat.Value >= rightFloat.Value), nil
		case "<=":
			return GoBoolToBoolean(leftFloat.Value <= rightFloat.Value), nil
		}
	} else if okLeftInt && okRightFloat {
		switch op {
		case ">":
			return GoBoolToBoolean(float64(leftInt.Value) > rightFloat.Value), nil
		case "<":
			return GoBoolToBoolean(float64(leftInt.Value) < rightFloat.Value), nil
		case ">=":
			return GoBoolToBoolean(float64(leftInt.Value) >= rightFloat.Value), nil
		case "<=":
			return GoBoolToBoolean(float64(leftInt.Value) <= rightFloat.Value), nil
		}
	} else if okLeftFloat && okRightInt {
		switch op {
		case ">":
			return GoBoolToBoolean(leftFloat.Value > float64(rightInt.Value)), nil
		case "<":
			return GoBoolToBoolean(leftFloat.Value < float64(rightInt.Value)), nil
		case ">=":
			return GoBoolToBoolean(leftFloat.Value >= float64(rightInt.Value)), nil
		case "<=":
			return GoBoolToBoolean(leftFloat.Value <= float64(rightInt.Value)), nil
		}
	} else if okLeftStr && okRightStr {
		// String comparisons
		switch op {
		case ">":
			return GoBoolToBoolean(leftStr.Value > rightStr.Value), nil
		case "<":
			return GoBoolToBoolean(leftStr.Value < rightStr.Value), nil
		case ">=":
			return GoBoolToBoolean(leftStr.Value >= rightStr.Value), nil
		case "<=":
			return GoBoolToBoolean(leftStr.Value <= rightStr.Value), nil
		}
	}

	return nil, fmt.Errorf("type error: cannot compare %s and %s with operator %s", left.Type(), right.Type(), op)
}

// Helper method for equality comparison (== and !=) (already implemented)
func (vm *VM) isEqual(left, right Value) bool {
	// ... (existing isEqual logic) ...
	// Implement equality logic. This can be complex depending on types (e.g., comparing arrays/hashes).
	// For basic types, compare values directly.
	// For numbers, allow comparison between int and float.
	leftInt, okLeftInt := left.(*Integer)
	rightInt, okRightInt := right.(*Integer)
	leftFloat, okLeftFloat := left.(*Float)
	rightFloat, okRightFloat := right.(*Float)

	if okLeftInt && okRightInt {
		return leftInt.Value == rightInt.Value
	}
	if okLeftFloat && okRightFloat {
		return leftFloat.Value == rightFloat.Value
	}
	if okLeftInt && okRightFloat {
		return float64(leftInt.Value) == rightFloat.Value
	}
	if okLeftFloat && okRightInt {
		return leftFloat.Value == float64(rightInt.Value)
	}

	// For other types, require exact type match for equality.
	if left.Type() != right.Type() {
		return false
	}

	switch left := left.(type) {
	case *Null:
		return true // nil == nil
	case *Boolean:
		return left.Value == right.(*Boolean).Value
	case *String:
		return left.Value == right.(*String).Value
	// TODO: Add equality for Array, Hash, Function, etc.
	default:
		// For unimplemented types, assume not equal for now.
		return false
	}
}
