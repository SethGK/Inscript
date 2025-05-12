package vm // Package vm, as it's in the internal/vm directory

import (
	// Needed for binary encoding/decoding (e.g., ReadUint16/Uint8)
	"fmt"
	"io"
	"math" // Needed for math.Pow in the power function
	"os"
	"strings"

	// Import the compiler package to access its types (Opcode, Instructions, Bytecode)
	"github.com/SethGK/Inscript/internal/compiler"
	// Import the new types package
	"github.com/SethGK/Inscript/internal/types"
	// No import needed for value.go as its content is moved to types/value.go
)

const StackSize = 2048
const GlobalsSize = 65536 // This constant is not strictly needed if sizing vm.globals dynamically
const MaxFrames = 1024

// VM represents the Inscript Virtual Machine.
type VM struct {
	constants []types.Value // Constant pool - Referring to Value from types package

	stack []types.Value // The main operand stack - Referring to Value from types package
	sp    int           // Stack pointer: points to the next free slot on the stack

	globals []types.Value // Global variables - Referring to Value from types package

	frames      []*Frame // Call frames for function execution - Referring to Frame (defined below in vm package)
	framesIndex int      // Current frame index - points to the next free frame slot

	// Output writer for print statements
	outputWriter io.Writer
}

// Frame represents a single call frame for function execution.
type Frame struct {
	closure     *types.Closure // The closure being executed
	ip          int            // Instruction pointer: points to the next instruction to execute
	basePointer int            // Base pointer: points to the first slot in the stack used by this frame (for locals and arguments)
	locals      []types.Value  // Local variables and parameters for this frame
}

// NewFrame creates a new call frame.
// basePointer is the stack index where this frame's locals and arguments begin.
func NewFrame(closure *types.Closure, basePointer int) *Frame {
	// The number of locals needed for a frame is the number of parameters + declared local variables.
	// This is stored in the CompiledFunction.NumLocals.
	locals := make([]types.Value, closure.Fn.NumLocals)

	// Copy arguments from the stack into the beginning of the locals slice.
	// Arguments are on the stack right before the function object, starting at basePointer + 1.
	// The number of arguments is closure.Fn.NumParameters.
	// This copying is done in the VM's OpCall case, not here.
	// The NewFrame function just creates the frame structure with the correct locals slice size.

	return &Frame{
		closure:     closure,
		ip:          -1, // Start at -1 so the first instruction is at index 0 after increment
		basePointer: basePointer,
		locals:      locals, // Allocate space for locals and parameters
	}
}

// Instructions returns the instructions of the compiled function in this frame.
func (f *Frame) Instructions() compiler.Instructions {
	return f.closure.Fn.Instructions
}

// New creates a new VM instance.
// It takes the bytecode compiled by the compiler package.
func New(bytecode *compiler.Bytecode) *VM { // Takes *compiler.Bytecode
	// The main program is treated as a function (the entry point).
	// Create a CompiledFunction and a Closure for the main program's bytecode.
	// These types are now in the types package.
	mainFn := &types.CompiledFunction{Instructions: bytecode.Instructions, NumLocals: bytecode.NumLocals, NumParameters: bytecode.NumParameters}
	mainClosure := &types.Closure{Fn: mainFn}
	// The base pointer for the main frame is 0, as it starts at the bottom of the stack.
	// The main program's frame should have 0 locals, as variables are global.
	mainFrame := NewFrame(mainClosure, 0) // NewFrame is defined in vm package. NumLocals for mainFn should be 0.

	// Initialize the frames stack with the main frame.
	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame

	return &VM{
		constants:    bytecode.Constants,                       // This should now work if Bytecode.Constants is []types.Value
		stack:        make([]types.Value, StackSize),           // Referring to Value from types package
		sp:           0,                                        // Stack starts empty, sp is the next available slot
		globals:      make([]types.Value, bytecode.NumGlobals), // Allocate space for globals based on compiler count - This now uses the corrected NumGlobals from the compiler
		frames:       frames,
		framesIndex:  1,         // Start with the main frame at index 0, next frame will be at index 1
		outputWriter: os.Stdout, // Default output to stdout
	}
}

// currentFrame returns the currently executing call frame.
func (vm *VM) currentFrame() *Frame { // Referring to Frame (defined below in vm package)
	return vm.frames[vm.framesIndex-1]
}

// pushFrame pushes a new call frame onto the frame stack.
func (vm *VM) pushFrame(frame *Frame) { // Takes *Frame (defined below in vm package)
	if vm.framesIndex >= MaxFrames {
		// TODO: Return a runtime error instead of panicking
		panic("stack overflow: too many function calls")
	}
	vm.frames[vm.framesIndex] = frame
	vm.framesIndex++
}

// popFrame pops the current call frame from the frame stack.
func (vm *VM) popFrame() *Frame { // Returning *Frame (defined below in vm package)
	if vm.framesIndex == 0 {
		// Should not happen if the main frame is always present
		panic("attempting to pop the last frame (main frame)")
	}
	vm.framesIndex--
	frame := vm.frames[vm.framesIndex]
	vm.frames[vm.framesIndex] = nil // Clear the reference
	return frame
}

// StackTop returns the value at the top of the stack without popping it.
func (vm *VM) StackTop() types.Value { // Returning Value from types package
	if vm.sp == 0 {
		return nil // Stack is empty
	}
	return vm.stack[vm.sp-1]
}

// push pushes a value onto the stack.
func (vm *VM) push(obj types.Value) { // Takes Value from types package
	if vm.sp >= StackSize {
		// TODO: Return a runtime error instead of panicking
		panic("stack overflow")
	}
	vm.stack[vm.sp] = obj
	vm.sp++
}

// pop pops a value from the stack.
func (vm *VM) pop() types.Value { // Returning Value from types package
	if vm.sp == 0 {
		// TODO: Return a runtime error instead of panicking
		panic("stack underflow")
	}
	vm.sp--
	obj := vm.stack[vm.sp]
	vm.stack[vm.sp] = nil // Avoid memory leaks by clearing the reference
	return obj
}

// Run executes the compiled bytecode.
func (vm *VM) Run() error {
	// The instruction pointer (ip) is now managed by the current call frame.
	var ip int
	var instructions compiler.Instructions // Using compiler.Instructions (from compiler package)
	var currentFrame *Frame                // Referring to Frame (defined below in vm package)

	// The loop continues as long as there are frames on the stack.
	// The main program's frame is the first one.
	for vm.framesIndex > 0 {
		currentFrame = vm.currentFrame()
		instructions = currentFrame.Instructions() // This returns compiler.Instructions

		// Increment the instruction pointer before fetching the opcode,
		// so ip points to the *current* instruction being executed.
		currentFrame.ip++
		ip = currentFrame.ip

		// If the instruction pointer is out of bounds for the current frame's instructions,
		// it means the function has finished executing without an explicit return.
		// This case should ideally be handled by the compiler emitting an implicit return,
		// but as a safeguard, we can pop the frame here.
		if ip >= len(instructions) {
			vm.popFrame() // Function finished, pop the frame
			// If this was the last frame (main program), the loop condition will become false.
			continue
		}

		opcode := compiler.OpCode(instructions[ip]) // Using compiler.OpCode (from compiler package)

		switch opcode {
		case compiler.OpConstant: // Using compiler.OpConstant (from compiler package)
			// Operand is the index of the constant in the constant pool.
			constantIndex := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			vm.push(vm.constants[constantIndex])
			currentFrame.ip += 2 // Operand (2 bytes for uint16)

		case compiler.OpAdd: // Using compiler.OpAdd (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.add(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpSub: // Using compiler.OpSub (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.subtract(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpMul: // Using compiler.OpMul (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.multiply(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpDiv: // Using compiler.OpDiv (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.divide(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpMod: // Using compiler.OpMod (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.modulo(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpPow: // Using compiler.OpPow (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.power(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpMinus: // Using compiler.OpMinus (from compiler package)
			operand := vm.pop()
			result := vm.negate(operand)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpNot: // Using compiler.OpNot (from compiler package)
			operand := vm.pop()
			result := vm.logicalNot(operand)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpEqual: // Using compiler.OpEqual (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.equal(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpNotEqual: // Using compiler.OpNotEqual (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.notEqual(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpGreaterThan: // Using compiler.OpGreaterThan (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.greaterThan(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpLessThan: // Using compiler.OpLessThan (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.lessThan(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpGreaterEqual: // Using compiler.OpGreaterEqual (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.greaterEqual(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpLessEqual: // Using compiler.OpLessEqual (from compiler package)
			right := vm.pop()
			left := vm.pop()
			result := vm.lessEqual(left, right)
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpTrue: // Using compiler.OpTrue (from compiler package)
			vm.push(&types.Boolean{Value: true}) // Referring to Boolean from types package
			// ip already incremented by the loop

		case compiler.OpFalse: // Using compiler.OpFalse (from compiler package)
			vm.push(&types.Boolean{Value: false}) // Referring to Boolean from types package
			// ip already incremented by the loop

		case compiler.OpNull: // Using compiler.OpNull (from compiler package)
			vm.push(&types.Nil{}) // Referring to Nil from types package
			// ip already incremented by the loop

		case compiler.OpPop: // Using compiler.OpPop (from compiler package)
			vm.pop()
			// ip already incremented by the loop

		case compiler.OpSetGlobal: // Using compiler.OpSetGlobal (from compiler package)
			globalIndex := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			vm.globals[globalIndex] = vm.pop()
			currentFrame.ip += 2 // Operand (2 bytes for uint16)

		case compiler.OpGetGlobal: // Using compiler.OpGetGlobal (from compiler package)
			globalIndex := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			// Check if the global index is within the bounds of the globals slice
			if int(globalIndex) >= len(vm.globals) {
				return fmt.Errorf("runtime error: global variable index out of bounds: %d (max %d)", globalIndex, len(vm.globals)-1)
			}
			vm.push(vm.globals[globalIndex]) // This was the line causing the panic
			currentFrame.ip += 2             // Operand (2 bytes for uint16)

		case compiler.OpSetLocal: // Using compiler.OpSetLocal (from compiler package)
			// Operand for local index is 1 byte (uint8)
			localIndex := compiler.ReadUint8(instructions[ip+1:]) // Using compiler.ReadUint8 (from compiler package)
			currentFrame.ip += 1                                  // Operand (1 byte for uint8)
			// Locals are stored in the current frame's locals array.
			// The operand is the index within that array.
			// Check if the local index is within the bounds of the frame's locals slice
			if int(localIndex) >= len(currentFrame.locals) {
				return fmt.Errorf("runtime error: local variable index out of bounds: %d (max %d)", localIndex, len(currentFrame.locals)-1)
			}
			currentFrame.locals[localIndex] = vm.pop()

		case compiler.OpGetLocal: // Using compiler.OpGetLocal (from compiler package)
			// Operand for local index is 1 byte (uint8)
			localIndex := compiler.ReadUint8(instructions[ip+1:]) // Using compiler.ReadUint8 (from compiler package)
			currentFrame.ip += 1                                  // Operand (1 byte for uint8)
			// Get the local variable from the current frame's locals array.
			// Check if the local index is within the bounds of the frame's locals slice
			if int(localIndex) >= len(currentFrame.locals) {
				return fmt.Errorf("runtime error: local variable index out of bounds: %d (max %d)", localIndex, len(currentFrame.locals)-1)
			}
			vm.push(currentFrame.locals[localIndex])

		case compiler.OpArray: // Using compiler.OpArray (from compiler package)
			numElements := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			currentFrame.ip += 2                                    // Operand (2 bytes for uint16)
			// Elements are on the stack, pop them in reverse order to build the list
			elements := make([]types.Value, numElements) // Referring to Value from types package
			for i := int(numElements) - 1; i >= 0; i-- {
				elements[i] = vm.pop()
			}
			vm.push(&types.List{Elements: elements}) // Referring to List from types package

		case compiler.OpHash: // Using compiler.OpHash (from compiler package)
			numPairs := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			currentFrame.ip += 2                                 // Operand (2 bytes for uint16)
			// Key-value pairs are on the stack: key1, value1, key2, value2, ...
			// Pop them in reverse order: valueN, keyN, ..., value1, key1
			fields := make(map[string]types.Value) // Referring to Value from types package
			for i := 0; i < int(numPairs); i++ {
				value := vm.pop()
				key := vm.pop()
				keyStr, ok := key.(*types.String) // Assuming keys are strings
				if !ok {
					// TODO: Handle non-string keys or return runtime error
					fmt.Fprintf(os.Stderr, "runtime error: hash key must be a string, got %s\n", key.Type())
					return fmt.Errorf("runtime error: hash key must be a string, got %s", key.Type())
				}
				fields[keyStr.Value] = value
			}
			vm.push(&types.Table{Fields: fields}) // Referring to Table from types package

		case compiler.OpIndex: // Using compiler.OpIndex (from compiler package)
			index := vm.pop()
			aggregate := vm.pop()
			// Call the GetIndex method on the aggregate value
			result, err := aggregate.GetIndex(index)
			if err != nil {
				// TODO: Handle runtime error
				fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
				return err
			}
			vm.push(result)
			// ip already incremented by the loop

		case compiler.OpSetIndex: // Using compiler.OpSetIndex (from compiler package)
			value := vm.pop()
			index := vm.pop()
			aggregate := vm.pop()
			// Call the SetIndex method on the aggregate value
			err := aggregate.SetIndex(index, value)
			if err != nil {
				// TODO: Handle runtime error
				fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
				return err
			}
			// The result of assignment is typically the assigned value or the aggregate itself.
			// Let's push the assigned value back onto the stack for now.
			vm.push(value)
			// ip already incremented by the loop

		case compiler.OpJump: // Using compiler.OpJump (from compiler package)
			pos := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			currentFrame.ip = int(pos) - 1                  // Set ip to target - 1, loop increment will make it target

		case compiler.OpJumpNotTruthy: // Using compiler.OpJumpNotTruthy (from compiler package)
			pos := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			currentFrame.ip += 2                            // Move past the operand (2 bytes)
			condition := vm.pop()
			if !isTruthy(condition) { // isTruthy is defined below in vm package
				currentFrame.ip = int(pos) - 1 // Set ip to target - 1
			}
			// ip will be incremented by the loop in the next iteration

		case compiler.OpJumpTruthy: // Using compiler.OpJumpTruthy (from compiler package)
			pos := compiler.ReadUint16(instructions[ip+1:]) // Using compiler.ReadUint16 (from compiler package)
			currentFrame.ip += 2                            // Move past the operand (2 bytes)
			condition := vm.pop()
			if isTruthy(condition) { // isTruthy is defined below in vm package
				currentFrame.ip = int(pos) - 1 // Set ip to target - 1
			}
			// ip will be incremented by the loop in the next iteration

		case compiler.OpPrint: // Using compiler.OpPrint (from compiler package)
			// Operand for print arguments count is 1 byte (uint8)
			numArgs := compiler.ReadUint8(instructions[ip+1:]) // Using compiler.ReadUint8 (from compiler package)
			currentFrame.ip += 1                               // Operand (1 byte for uint8)

			// Arguments are on the stack in reverse order of compilation.
			// Pop them and print.
			args := make([]string, numArgs)
			// Pop from top of stack backwards
			for i := int(numArgs) - 1; i >= 0; i-- {
				args[i] = vm.pop().Inspect() // Assuming Inspect() gives a string representation
			}
			fmt.Fprintln(vm.outputWriter, strings.Join(args, " "))

		case compiler.OpCall: // Using compiler.OpCall (from compiler package)
			// Operand for call arguments count is 1 byte (uint8)
			numArgs := compiler.ReadUint8(instructions[ip+1:]) // Using compiler.ReadUint8 (from compiler package)
			currentFrame.ip += 1                               // Operand (1 byte for uint8)

			// The function object is below the arguments on the stack.
			// Stack: [..., function, arg1, arg2, ..., argN]
			// The base pointer for the new frame is the position *before* the function object.
			// The arguments are at stack[vm.sp - numArgs] to stack[vm.sp - 1].
			// The function object is at stack[vm.sp - numArgs - 1].
			// The base pointer for the new frame is vm.sp - int(numArgs) - 1.
			basePointer := vm.sp - int(numArgs) - 1

			// Get the function object from the stack (it's at basePointer).
			function := vm.stack[basePointer]
			closure, ok := function.(*types.Closure) // Expecting a Closure object from types package
			if !ok {
				// TODO: Handle calling non-closure values (e.g., calling a number)
				return fmt.Errorf("runtime error: cannot call value of type %s", function.Type())
			}

			// Get the compiled function from the closure.
			fn := closure.Fn // Fn is *types.CompiledFunction

			if int(numArgs) != fn.NumParameters {
				return fmt.Errorf("runtime error: wrong number of arguments: expected %d, got %d", fn.NumParameters, numArgs)
			}

			// Create a new call frame for the function.
			// The arguments are already on the stack in the correct position relative to the new frame's base pointer.
			// The NewFrame function will copy the arguments into the frame's locals array.
			newFrame := NewFrame(closure, basePointer) // NewFrame is defined in vm package

			// Copy arguments from the stack into the new frame's locals.
			// Arguments are located on the stack starting at basePointer + 1.
			// The new frame's locals array starts at index 0.
			for i := 0; i < fn.NumParameters; i++ {
				newFrame.locals[i] = vm.stack[basePointer+1+i]
			}

			// Push the new frame onto the frame stack.
			vm.pushFrame(newFrame)

			// The VM will continue execution from the start of the new frame's instructions
			// in the next iteration of the main loop.
			// The stack pointer (vm.sp) remains where it is for now; the new frame
			// will manage its locals relative to its base pointer.

		case compiler.OpReturn: // Using compiler.OpReturn (from compiler package)
			// This opcode is typically emitted by the compiler at the end of a function
			// or for a `return` statement without a value.
			// It implicitly returns nil.

			// Pop the current frame.
			frame := vm.popFrame() // Referring to popFrame (defined below in vm package)

			// The return value is implicitly null.
			returnValue := &types.Nil{} // Referring to Nil from types package

			// Restore the stack pointer to the state before the function call.
			// This is the base pointer of the popped frame.
			vm.sp = frame.basePointer

			// Push the return value onto the stack of the previous frame.
			vm.push(returnValue)

		case compiler.OpReturnValue: // Using compiler.OpReturnValue (from compiler package)
			// This opcode is emitted by the compiler for a `return expression` statement.
			// The return value is already on top of the stack.
			returnValue := vm.pop()

			// Pop the current frame.
			frame := vm.popFrame() // Referring to popFrame (defined below in vm package)

			// Restore the stack pointer to the state before the function call.
			// This is the base pointer of the popped frame.
			vm.sp = frame.basePointer

			// Push the return value onto the stack of the previous frame.
			vm.push(returnValue)

		case compiler.OpGetIterator:
			iterable := vm.pop()
			it, err := iterable.GetIterator()
			if err != nil {
				return fmt.Errorf("runtime error: %v", err)
			}
			vm.push(it)

		case compiler.OpIteratorNext:
			// Pop the iterator
			iteratorVal := vm.pop()
			iterator, ok := iteratorVal.(types.Iterator)
			if !ok {
				return fmt.Errorf("runtime error: cannot call Next on non-iterator value of type %s", iteratorVal.Type())
			}

			// Next() → (value, ok, err)
			value, success, err := iterator.Next()
			if err != nil {
				return fmt.Errorf("runtime error: iteration error: %v", err)
			}

			// Push iterator back, then value and the success flag directly
			vm.push(iteratorVal)
			vm.push(value)
			vm.push(&types.Boolean{Value: success})

		default:
			// TODO: Return a runtime error for unknown opcode
			return fmt.Errorf("runtime error: unknown opcode %v at %d", opcode, ip)
		}
	}

	// Execution finishes when the main frame is popped (vm.framesIndex becomes 0).
	// The result of the program is the value left on the stack (if any, usually the result of the last statement).
	// For a program, the compiler adds OpNull and OpReturn at the end, so the final result should be nil.
	// We can pop the final result here if needed, or just let the VM stop.
	// The compiler's final OpReturn will handle returning from the main frame.

	return nil
}

// Helper functions for arithmetic and comparison operations.
// These should handle type checking and return appropriate results or errors.
// These helpers operate on types.Value and return types.Value.

// add performs addition on two values.
func (vm *VM) add(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values first to avoid panics
	if left == nil || right == nil {
		// TODO: Implement proper runtime error reporting
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform addition on nil values\n")
		return &types.Nil{} // Return nil value or a specific error value - Referring to Nil from types package
	}

	// Handle addition based on the types of the operands
	switch left := left.(type) {
	case *types.Integer: // Referring to Integer from types package
		switch right := right.(type) {
		case *types.Integer: // Referring to Integer from types package
			return &types.Integer{Value: left.Value + right.Value} // Referring to Integer from types package
		case *types.Float: // Referring to Float from types package
			return &types.Float{Value: float64(left.Value) + right.Value} // Referring to Float from types package
		case *types.String: // Referring to String from types package
			// String concatenation
			return &types.String{Value: fmt.Sprintf("%d%s", left.Value, right.Value)} // Referring to String from types package
		default:
			// Incompatible types for addition
			fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: integer + %s\n", right.Type())
			return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
		}
	case *types.Float: // Referring to Float from types package
		switch right := right.(type) {
		case *types.Integer: // Referring to Integer from types package
			return &types.Float{Value: left.Value + float64(right.Value)} // Referring to Float from types package
		case *types.Float: // Referring to Float from types package
			return &types.Float{Value: left.Value + right.Value} // Referring to Float from types package
		case *types.String: // Referring to String from types package
			// String concatenation
			return &types.String{Value: fmt.Sprintf("%f%s", left.Value, right.Value)} // Referring to String from types package
		default:
			// Incompatible types for addition
			fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: float + %s\n", right.Type())
			return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
		}
	case *types.String: // Referring to String from types package
		switch right := right.(type) {
		case *types.Integer: // Referring to Integer from types package
			// String concatenation
			return &types.String{Value: fmt.Sprintf("%s%d", left.Value, right.Value)} // Referring to String from types package
		case *types.Float: // Referring to Float from types package
			// String concatenation
			return &types.String{Value: fmt.Sprintf("%s%f", left.Value, right.Value)} // Referring to String from types package
		case *types.String: // Referring to String from types package
			// String concatenation
			return &types.String{Value: left.Value + right.Value} // Referring to String from types package
		default:
			// Incompatible types for addition
			fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: string + %s\n", right.Type())
			return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
		}
	case *types.List: // Referring to List from types package
		switch right := right.(type) {
		case *types.List: // Referring to List from types package
			// List concatenation
			newList := make([]types.Value, len(left.Elements)+len(right.Elements))
			copy(newList, left.Elements)
			copy(newList[len(left.Elements):], right.Elements)
			return &types.List{Elements: newList} // Referring to List from types package
		default:
			fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: list + %s\n", right.Type())
			return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
		}

	default:
		// Unsupported type for addition
		fmt.Fprintf(os.Stderr, "runtime error: unsupported type for addition: %s\n", left.Type())
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
}

// subtract performs subtraction on two values.
func (vm *VM) subtract(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform subtraction on nil values\n")
		return &types.Nil{} // Referring to Nil from types package
	}

	// Subtraction is typically only defined for numbers.
	leftInt, isLeftInt := left.(*types.Integer)      // Referring to Integer from types package
	leftFloat, isLeftFloat := left.(*types.Float)    // Referring to Float from types package
	rightInt, isRightInt := right.(*types.Integer)   // Referring to Integer from types package
	rightFloat, isRightFloat := right.(*types.Float) // Referring to Float from types package

	if isLeftInt && isRightInt {
		return &types.Integer{Value: leftInt.Value - rightInt.Value} // Referring to Integer from types package
	}
	if isLeftFloat && isRightFloat {
		return &types.Float{Value: leftFloat.Value - rightFloat.Value} // Referring to Float from types package
	}
	if isLeftInt && isRightFloat {
		return &types.Float{Value: float64(leftInt.Value) - rightFloat.Value} // Referring to Float from types package
	}
	if isLeftFloat && isRightInt {
		return &types.Float{Value: leftFloat.Value - float64(rightInt.Value)} // Referring to Float from types package
	}

	// Incompatible types for subtraction
	fmt.Fprintf(os.Stderr, "runtime error: unsupported types for subtraction: %s - %s\n", left.Type(), right.Type())
	return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
}

// multiply performs multiplication on two values.
func (vm *VM) multiply(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform multiplication on nil values\n")
		return &types.Nil{} // Referring to Nil from types package
	}

	// Multiplication is typically only defined for numbers.
	leftInt, isLeftInt := left.(*types.Integer)      // Referring to Integer from types package
	leftFloat, isLeftFloat := left.(*types.Float)    // Referring to Float from types package
	rightInt, isRightInt := right.(*types.Integer)   // Referring to Integer from types package
	rightFloat, isRightFloat := right.(*types.Float) // Referring to Float from types package

	if isLeftInt && isRightInt {
		return &types.Integer{Value: leftInt.Value * rightInt.Value} // Referring to Integer from types package
	}
	if isLeftFloat && isRightFloat {
		return &types.Float{Value: leftFloat.Value * rightFloat.Value} // Referring to Float from types package
	}
	if isLeftInt && isRightFloat {
		return &types.Float{Value: float64(leftInt.Value) * rightFloat.Value} // Referring to Float from types package
	}
	if isLeftFloat && isRightInt {
		return &types.Float{Value: leftFloat.Value * float64(rightInt.Value)} // Referring to Float from types package
	}

	// Incompatible types for multiplication
	fmt.Fprintf(os.Stderr, "runtime error: unsupported types for multiplication: %s * %s\n", left.Type(), right.Type())
	return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
}

// divide performs division on two values.
func (vm *VM) divide(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform division on nil values\n")
		return &types.Nil{} // Referring to Nil from types package
	}

	// Division is typically only defined for numbers.
	leftInt, isLeftInt := left.(*types.Integer)      // Referring to Integer from types package
	leftFloat, isLeftFloat := left.(*types.Float)    // Referring to Float from types package
	rightInt, isRightInt := right.(*types.Integer)   // Referring to Integer from types package
	rightFloat, isRightFloat := right.(*types.Float) // Referring to Float from types package

	// Handle division by zero
	if (isRightInt && rightInt.Value == 0) || (isRightFloat && rightFloat.Value == 0.0) {
		fmt.Fprintf(os.Stderr, "runtime error: division by zero\n")
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}

	if isLeftInt && isRightInt {
		// Integer division
		return &types.Integer{Value: leftInt.Value / rightInt.Value} // Referring to Integer from types package
	}
	if isLeftFloat && isRightFloat {
		return &types.Float{Value: leftFloat.Value / rightFloat.Value} // Referring to Float from types package
	}
	if isLeftInt && isRightFloat {
		return &types.Float{Value: float64(leftInt.Value) / rightFloat.Value} // Referring to Float from types package
	}
	if isLeftFloat && isRightInt {
		return &types.Float{Value: leftFloat.Value / float64(rightInt.Value)} // Referring to Float from types package
	}

	// Incompatible types for division
	fmt.Fprintf(os.Stderr, "runtime error: unsupported types for division: %s / %s\n", left.Type(), right.Type())
	return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
}

// modulo performs the modulo operation on two values.
func (vm *VM) modulo(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform modulo on nil values\n")
		return &types.Nil{} // Referring to Nil from types package
	}

	// Modulo is typically only defined for integers.
	leftInt, isLeftInt := left.(*types.Integer)    // Referring to Integer from types package
	rightInt, isRightInt := right.(*types.Integer) // Referring to Integer from types package

	if isLeftInt && isRightInt {
		// Handle modulo by zero
		if rightInt.Value == 0 {
			fmt.Fprintf(os.Stderr, "runtime error: modulo by zero\n")
			return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
		}
		return &types.Integer{Value: leftInt.Value % rightInt.Value} // Referring to Integer from types package
	}

	// Incompatible types for modulo
	fmt.Fprintf(os.Stderr, "runtime error: unsupported types for modulo: %s %% %s\n", left.Type(), right.Type())
	return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
}

// power performs the power operation (left^right) on two values.
func (vm *VM) power(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform power on nil values\n")
		return &types.Nil{} // Referring to Nil from types package
	}

	// Power is typically defined for numbers.
	leftInt, isLeftInt := left.(*types.Integer)      // Referring to Integer from types package
	leftFloat, isLeftFloat := left.(*types.Float)    // Referring to Float from types package
	rightInt, isRightInt := right.(*types.Integer)   // Referring to Integer from types package
	rightFloat, isRightFloat := right.(*types.Float) // Referring to Float from types package

	if isLeftInt && isRightInt {
		// Integer power
		return &types.Integer{Value: int64(math.Pow(float64(leftInt.Value), float64(rightInt.Value)))} // Referring to Integer from types package
	}
	if isLeftFloat && isRightFloat {
		return &types.Float{Value: math.Pow(leftFloat.Value, rightFloat.Value)} // Referring to Float from types package
	}
	if isLeftInt && isRightFloat {
		return &types.Float{Value: math.Pow(float64(leftInt.Value), rightFloat.Value)} // Referring to Float from types package
	}
	if isLeftFloat && isRightInt {
		return &types.Float{Value: math.Pow(leftFloat.Value, float64(rightInt.Value))} // Referring to Float from types package
	}

	// Incompatible types for power
	fmt.Fprintf(os.Stderr, "runtime error: unsupported types for power: %s ^ %s\n", left.Type(), right.Type())
	return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
}

// negate performs numerical negation on a value.
func (vm *VM) negate(operand types.Value) types.Value { // Referring to Value from types package
	// Check for nil value
	if operand == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot perform negation on nil value\n")
		return &types.Nil{} // Referring to Nil from types package
	}

	// Negation is typically only defined for numbers.
	switch operand := operand.(type) {
	case *types.Integer: // Referring to Integer from types package
		return &types.Integer{Value: -operand.Value} // Referring to Integer from types package
	case *types.Float: // Referring to Float from types package
		return &types.Float{Value: -operand.Value} // Referring to Float from types package
	default:
		// Unsupported type for negation
		fmt.Fprintf(os.Stderr, "runtime error: unsupported type for negation: %s\n", operand.Type())
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
}

// logicalNot performs logical NOT on a value.
func (vm *VM) logicalNot(operand types.Value) types.Value { // Referring to Value from types package
	// Check for nil value
	if operand == nil {
		// 'not nil' is typically true
		return &types.Boolean{Value: true} // Referring to Boolean from types package
	}

	// Logical NOT is typically defined for booleans.
	// For other types, we can define truthiness rules (e.g., numbers other than 0 are true, non-empty strings/lists/tables are true).
	// Based on your compiler's OpJumpNotTruthy logic, it seems you have a concept of truthiness.
	// Let's use the isTruthy helper here.
	return &types.Boolean{Value: !isTruthy(operand)} // Referring to Boolean from types package
}

// equal checks for equality between two values.
func (vm *VM) equal(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Equals method defined on the Value interface.
	// Handle nil comparisons explicitly.
	if left == nil && right == nil {
		return &types.Boolean{Value: true} // Referring to Boolean from types package
	}
	if left == nil || right == nil {
		return &types.Boolean{Value: false} // Referring to Boolean from types package
	}
	return &types.Boolean{Value: left.Equals(right)} // Referring to Boolean from types package
}

// notEqual checks for inequality between two values.
func (vm *VM) notEqual(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Equals method and negate the result.
	if left == nil && right == nil {
		return &types.Boolean{Value: false} // Referring to Boolean from types package
	}
	if left == nil || right == nil {
		return &types.Boolean{Value: true} // Referring to Boolean from types package
	}
	return &types.Boolean{Value: !left.Equals(right)} // Referring to Boolean from types package
}

// greaterThan checks if left is greater than right.
func (vm *VM) greaterThan(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method defined on the Value interface.
	// Handle nil comparisons.
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot compare nil values\n")
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	comparison, err := left.Compare(right)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	return &types.Boolean{Value: comparison > 0} // Referring to Boolean from types package
}

// lessThan checks if left is less than right.
func (vm *VM) lessThan(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method defined on the Value interface.
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot compare nil values\n")
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	comparison, err := left.Compare(right)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	return &types.Boolean{Value: comparison < 0} // Referring to Boolean from types package
}

// greaterEqual checks if left is greater than or equal to right.
func (vm *VM) greaterEqual(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method defined on the Value interface.
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot compare nil values\n")
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	comparison, err := left.Compare(right)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	return &types.Boolean{Value: comparison >= 0} // Referring to Boolean from types package
}

// lessEqual checks if left is less than or equal to right.
func (vm *VM) lessEqual(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method defined on the Value interface.
	if left == nil || right == nil {
		fmt.Fprintf(os.Stderr, "runtime error: cannot compare nil values\n")
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	comparison, err := left.Compare(right)
	if err != nil {
		fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Nil{} // TODO: Return a specific error value - Referring to Nil from types package
	}
	return &types.Boolean{Value: comparison <= 0} // Referring to Boolean from types package
}

// isTruthy determines the truthiness of a value.
// This is used for conditional jumps (OpJumpNotTruthy, OpJumpTruthy).
func isTruthy(obj types.Value) bool { // Referring to Value from types package
	if obj == nil {
		return false // nil is falsy
	}
	switch obj := obj.(type) {
	case *types.Boolean: // Referring to Boolean from types package
		return obj.Value // Boolean value itself
	case *types.Nil: // Referring to Nil from types package
		return false // Nil is falsy
	case *types.Integer: // Referring to Integer from types package
		return obj.Value != 0 // Integers are truthy if non-zero
	case *types.Float: // Referring to Float from types package
		return obj.Value != 0.0 // Floats are truthy if non-zero
	case *types.String: // Referring to String from types package
		return len(obj.Value) > 0 // Strings are truthy if non-empty
	case *types.List: // Referring to List from types package
		return len(obj.Elements) > 0 // Lists are truthy if non-empty
	case *types.Table: // Referring to Table from types package
		return len(obj.Fields) > 0 // Tables are truthy if non-empty
	case *types.CompiledFunction: // Referring to CompiledFunction from types package
		return true // Functions are truthy
	case *types.Closure: // Referring to Closure from types package
		return true // Closures are truthy
	case types.Iterator: // Referring to Iterator from types package
		// Iterators are typically truthy as long as they haven't finished iteration.
		// This might require a method on the Iterator interface to check if it's done.
		// For now, let's assume iterators are always truthy if they exist.
		return true
	case *types.Error: // Referring to Error from types package
		return false // Errors are typically falsy
	default:
		// Default to falsy for unknown types, or error depending on language design.
		return false
	}
}
