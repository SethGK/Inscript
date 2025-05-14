// Package vm, as it's in the internal/vm directory
package vm

import (
	// Needed for binary encoding/decoding (e.g., ReadUint16/Uint8)
	"encoding/binary" // Added binary import
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
	// The locals slice holds the parameters and declared local variables.
	locals := make([]types.Value, closure.Fn.NumLocals)

	// --- Debug Print: NewFrame ---
	// Use Inspect() instead of String() on closure.Fn
	fmt.Printf("DEBUG (NewFrame): Created frame for %s. NumLocals: %d, locals slice length: %d\n", closure.Fn.Inspect(), closure.Fn.NumLocals, len(locals))
	// --- End Debug Print ---

	// Arguments are copied into the beginning of the locals slice in the VM's OpCall case.
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
	// The main program's bytecode is the entry point. It has 0 parameters and locals.
	mainFn := &types.CompiledFunction{Instructions: bytecode.Instructions, NumLocals: 0, NumParameters: 0} // NumLocals and NumParameters for main program should be 0
	mainClosure := &types.Closure{Fn: mainFn}
	// The base pointer for the main frame is 0, as it starts at the bottom of the stack.
	mainFrame := NewFrame(mainClosure, 0) // NewFrame is defined in vm package. basePointer is 0.

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
	// --- Debug Print: Push ---
	// Get the current frame's base pointer for context
	currentBasePtr := 0
	currentFrame := vm.currentFrame()              // Get current frame
	if vm.framesIndex > 0 && currentFrame != nil { // Add nil check for currentFrame
		currentBasePtr = currentFrame.basePointer
	}
	fmt.Printf("DEBUG (Push): Pushing %s at index %d (sp before: %d, basePtr: %d)\n", obj.Inspect(), vm.sp, vm.sp, currentBasePtr)
	// --- End Debug Print ---
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
	// --- Debug Print: Pop ---
	// Get the current frame's base pointer for context
	currentBasePtr := 0
	currentFrame := vm.currentFrame()              // Get current frame
	if vm.framesIndex > 0 && currentFrame != nil { // Add nil check for currentFrame
		currentBasePtr = currentFrame.basePointer
	}
	fmt.Printf("DEBUG (Pop): Popped %s from index %d (sp after: %d, basePtr: %d)\n", obj.Inspect(), vm.sp, vm.sp, currentBasePtr)
	// --- End Debug Print ---
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
		// --- Debug Print: Start of Run loop iteration ---
		fmt.Printf("DEBUG (Run Loop): Start of iteration. framesIndex: %d, currentFrame basePtr: %d, currentFrame locals length: %d\n", vm.framesIndex, currentFrame.basePointer, len(currentFrame.locals))
		// --- End Debug Print ---

		instructions = currentFrame.Instructions() // This returns compiler.Instructions
		// --- Debug Print: Instructions slice for current frame ---
		fmt.Printf("DEBUG (Run Loop): Current Frame Instructions: %v\n", instructions)
		// --- End Debug Print ---

		// Increment the instruction pointer before fetching the opcode,
		// so ip points to the *current* instruction being executed.
		currentFrame.ip++
		ip = currentFrame.ip

		// If the instruction pointer is out of bounds for the current frame's instructions,
		// it means the function has finished executing without an explicit return.
		// This case should ideally be handled by the compiler emitting an implicit return,
		// but as a safeguard, we can pop the frame here.
		if ip >= len(instructions) {
			// This should ideally not be reached if the compiler emits a final OpReturn/OpReturnValue
			// However, as a safeguard:
			// If the main frame finishes, the loop condition (vm.framesIndex > 0) will become false.
			// If a function frame finishes without a return, it implicitly returns nil.
			// Let's ensure an implicit nil return if we reach the end of function instructions.
			if vm.framesIndex > 1 { // Not the main frame
				// Implicit return nil
				// The compiler should ideally push nil before the end of instructions if no explicit return.
				// If it doesn't, we need to handle it here.
				// Assuming compiler emits OpNull before the end if needed.
				// Just execute the return logic.
				frame := vm.popFrame()
				vm.sp = frame.basePointer // Restore stack pointer
				// The implicit nil return value should be at vm.stack[vm.sp] if compiler is correct.
				continue // Continue loop to process the caller's next instruction
			}
			// If it's the main frame and we reached the end, the loop will terminate.
			break // Exit the loop if main frame finishes
		}

		// --- Debug Print: Raw Instruction Bytes ---
		// Print the byte at the current ip and the next few bytes
		bytesToPrint := 5 // Adjust as needed
		endIndex := ip + bytesToPrint
		if endIndex > len(instructions) {
			endIndex = len(instructions)
		}
		fmt.Printf("DEBUG (Run Loop): Raw bytes at ip %d: %v\n", ip, instructions[ip:endIndex])
		// --- End Debug Print ---

		opcode := compiler.OpCode(instructions[ip]) // Using compiler.OpCode (from compiler package)
		// --- Debug Print: Fetched Opcode ---
		fmt.Printf("DEBUG (Run Loop): Fetched opcode: %v (int: %d) at ip %d\n", opcode, int(opcode), ip)
		// --- End Debug Print ---

		// Get the instruction definition to determine operand widths
		def, ok := compiler.Lookup(opcode)
		if !ok {
			return fmt.Errorf("runtime error: unknown opcode: %v at %d", opcode, ip)
		}

		// Calculate the size of the current instruction (opcode + operands)
		instructionSize := 1 // Opcode size
		for _, width := range def.OperandWidths {
			instructionSize += width
		}

		// Ensure there are enough bytes remaining for the instruction and its operands
		if ip+instructionSize > len(instructions) {
			return fmt.Errorf("runtime error: incomplete instruction at %d. Expected %d bytes, but only %d remaining", ip, instructionSize, len(instructions)-ip)
		}

		// --- Debug Print: Opcode value before dispatch ---
		fmt.Printf("DEBUG (Run Loop): Opcode value before dispatch: %d\n", int(opcode))
		// --- End Debug Print ---

		// Using if-else if for opcode dispatch
		if opcode == compiler.OpConstant {
			// Operand is the index of the constant in the constant pool.
			constantIndex := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			vm.push(vm.constants[constantIndex])
		} else if opcode == compiler.OpAdd {
			right := vm.pop()
			left := vm.pop()
			result := vm.add(left, right)
			vm.push(result)
		} else if opcode == compiler.OpSub {
			right := vm.pop()
			left := vm.pop()
			result := vm.subtract(left, right)
			vm.push(result)
		} else if opcode == compiler.OpMul {
			right := vm.pop()
			left := vm.pop()
			result := vm.multiply(left, right)
			vm.push(result)
		} else if opcode == compiler.OpDiv {
			right := vm.pop()
			left := vm.pop()
			result := vm.divide(left, right)
			vm.push(result)
		} else if opcode == compiler.OpMod {
			right := vm.pop()
			left := vm.pop()
			result := vm.modulo(left, right)
			vm.push(result)
		} else if opcode == compiler.OpPow {
			right := vm.pop()
			left := vm.pop()
			result := vm.power(left, right)
			vm.push(result)
		} else if opcode == compiler.OpMinus {
			operand := vm.pop()
			result := vm.negate(operand)
			vm.push(result)
		} else if opcode == compiler.OpNot {
			operand := vm.pop()
			result := vm.logicalNot(operand)
			vm.push(result)
		} else if opcode == compiler.OpEqual {
			right := vm.pop()
			left := vm.pop()
			result := vm.equal(left, right)
			vm.push(result)
		} else if opcode == compiler.OpNotEqual {
			right := vm.pop()
			left := vm.pop()
			result := vm.notEqual(left, right)
			vm.push(result)
		} else if opcode == compiler.OpGreaterThan {
			right := vm.pop()
			left := vm.pop()
			result := vm.greaterThan(left, right)
			vm.push(result)
		} else if opcode == compiler.OpLessThan {
			right := vm.pop()
			left := vm.pop()
			result := vm.lessThan(left, right)
			vm.push(result)
		} else if opcode == compiler.OpGreaterEqual {
			right := vm.pop()
			left := vm.pop()
			result := vm.greaterEqual(left, right)
			vm.push(result)
		} else if opcode == compiler.OpLessEqual {
			right := vm.pop()
			left := vm.pop()
			result := vm.lessEqual(left, right)
			vm.push(result)
		} else if opcode == compiler.OpTrue {
			vm.push(&types.Boolean{Value: true}) // Referring to Boolean from types package
		} else if opcode == compiler.OpFalse {
			vm.push(&types.Boolean{Value: false}) // Referring to Boolean from types package
		} else if opcode == compiler.OpNull {
			vm.push(&types.Nil{}) // Referring to Nil from types package
		} else if opcode == compiler.OpPop {
			vm.pop()
		} else if opcode == compiler.OpSetGlobal {
			globalIndex := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			vm.globals[globalIndex] = vm.pop()
		} else if opcode == compiler.OpGetGlobal {
			globalIndex := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			// Check if the global index is within the bounds of the globals slice
			if int(globalIndex) >= len(vm.globals) {
				return fmt.Errorf("runtime error: global variable index out of bounds: %d (max %d)", globalIndex, len(vm.globals)-1)
			}
			vm.push(vm.globals[globalIndex])
		} else if opcode == compiler.OpSetLocal {
			// Operand for local index is 1 byte (uint8)
			localIndex := instructions[ip+1] // Read uint8 directly
			// Locals are stored in the current frame's locals array.
			// The operand is the index within that array.
			// Check if the local index is within the bounds of the frame's locals slice
			if int(localIndex) >= len(currentFrame.locals) {
				return fmt.Errorf("runtime error: local variable index out of bounds: %d (max %d)", localIndex, len(currentFrame.locals)-1)
			}
			valueToSet := vm.pop()
			currentFrame.locals[localIndex] = valueToSet
		} else if opcode == compiler.OpGetLocal {
			// Operand for local index is 1 byte (uint8)
			localIndex := instructions[ip+1] // Read uint8 directly
			// Get the local variable from the current frame's locals array.
			// *** Debug Print: OpGetLocal ***
			fmt.Printf("DEBUG (OpGetLocal): Accessing local index %d. currentFrame.locals length: %d\n", localIndex, len(currentFrame.locals))
			// *** End Debug Print ---
			// Check if the local index is within the bounds of the frame's locals slice
			if int(localIndex) >= len(currentFrame.locals) {
				return fmt.Errorf("runtime error: local variable index out of bounds: %d (max %d)", localIndex, len(currentFrame.locals)-1)
			}
			// Push from the frame's locals slice onto the main stack
			vm.push(currentFrame.locals[localIndex])
		} else if opcode == compiler.OpGetFree {
			// --- Debug Print: OpGetFree Entered ---
			fmt.Printf("DEBUG (OpGetFree): Entered OpGetFree case at ip %d. Opcode value inside case: %d\n", ip, int(opcode))
			// --- End Debug Print ---
			// Operand is a single-byte index into the current closure's Free slice
			freeIndex := instructions[ip+1] // Read uint8 directly
			// Push the captured free variable onto the stack
			if int(freeIndex) >= len(currentFrame.closure.Free) {
				return fmt.Errorf("runtime error: free variable index out of bounds: %d (max %d)", freeIndex, len(currentFrame.closure.Free)-1)
			}
			vm.push(currentFrame.closure.Free[freeIndex])
		} else if opcode == compiler.OpClosure {
			// Operands: 2‑byte constant pool index, then 1‑byte free count
			constIndex := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			freeCount := int(instructions[ip+3])                       // Read uint8 directly

			// Look up the compiled function
			fnVal := vm.constants[constIndex]
			compiledFn, ok := fnVal.(*types.CompiledFunction)
			if !ok {
				return fmt.Errorf("runtime error: constant %d is not a CompiledFunction, got %s", constIndex, fnVal.Type())
			}

			// Pop freeCount values from the stack (in reverse order) to capture
			frees := make([]types.Value, freeCount)
			for i := freeCount - 1; i >= 0; i-- {
				frees[i] = vm.pop()
			}

			// Build and push the closure
			clo := &types.Closure{Fn: compiledFn, Free: frees}
			vm.push(clo)
		} else if opcode == compiler.OpArray {
			numElements := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			// Elements are on the stack, pop them in reverse order to build the list
			elements := make([]types.Value, numElements) // Referring to Value from types package
			for i := int(numElements) - 1; i >= 0; i-- {
				elements[i] = vm.pop()
			}
			vm.push(&types.List{Elements: elements}) // Referring to List from types package
		} else if opcode == compiler.OpHash {
			numPairs := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian

			// Elements are on the stack in reverse order (valueN, keyN, ... value1, key1)
			// We need to pop them and store them to build the ordered list.
			// Pop all pairs first.
			poppedPairs := make([]struct {
				Key   types.Value
				Value types.Value
			}, numPairs)
			// Pop in reverse order of compilation (which is value, then key)
			for i := int(numPairs) - 1; i >= 0; i-- {
				poppedPairs[i].Value = vm.pop() // Pop value
				poppedPairs[i].Key = vm.pop()   // Pop key
			}

			// Now build the ordered structure for the Table from the popped pairs in the correct order (key1, value1, ...)
			orderedPairs := make([]types.TablePair, numPairs) // Use the new TablePair struct
			// lookupMap is not needed here in the VM, only in the Table type itself.
			// lookupMap := make(map[string]int, numPairs)

			// Iterate through the popped pairs in their original compilation order (key1, value1, ...)
			for i := 0; i < int(numPairs); i++ {
				keyStr, ok := poppedPairs[i].Key.(*types.String)
				if !ok {
					return fmt.Errorf("runtime error: hash key must be a string, got %s", poppedPairs[i].Key.Type())
				}
				orderedPairs[i] = types.TablePair{Key: keyStr.Value, Value: poppedPairs[i].Value}
				// lookupMap[keyStr.Value] = i // Store index for Lookup
			}

			// Push the new Table with the ordered structure.
			// Use the NewTable helper which handles initializing the Lookup map correctly.
			vm.push(types.NewTable(orderedPairs)) // Use the NewTable helper
		} else if opcode == compiler.OpIndex {
			// Stack top is now [..., aggregate, index].
			index := vm.pop()
			aggregate := vm.pop()
			// Call the GetIndex method on the aggregate value.
			// This method is now implemented in the types package to handle the new Table structure.
			result, err := aggregate.GetIndex(index)
			if err != nil {
				// TODO: Handle runtime error
				fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
				return err
			}
			vm.push(result)
		} else if opcode == compiler.OpSetIndex {
			// Stack top is now [..., aggregate, index, value].
			value := vm.pop()
			index := vm.pop()
			aggregate := vm.pop()
			// Call the SetIndex method on the aggregate value.
			// This method is now implemented in the types package to handle the new Table structure.
			err := aggregate.SetIndex(index, value)
			if err != nil {
				// TODO: Handle runtime error
				fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
				return err
			}
			// The result of assignment is typically the assigned value or the aggregate itself.
			// Let's push the assigned value back onto the stack for now.
			vm.push(value)
		} else if opcode == compiler.OpJump { // Added Jump cases back
			pos := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			currentFrame.ip = int(pos) - 1                      // Set ip to target - 1, loop increment will make it target
			// ip is NOT advanced after the dispatch block
		} else if opcode == compiler.OpJumpNotTruthy { // Added Jump cases back
			pos := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			condition := vm.pop()
			if !isTruthy(condition) { // isTruthy is defined below in vm package
				currentFrame.ip = int(pos) - 1 // Set ip to target - 1
			} else {
				// If not jumping, advance ip past the operand
				currentFrame.ip += 2
			}
			// ip is NOT advanced after the dispatch block
		} else if opcode == compiler.OpJumpTruthy { // Added Jump cases back
			pos := binary.BigEndian.Uint16(instructions[ip+1:]) // Using binary.BigEndian
			condition := vm.pop()
			if isTruthy(condition) { // isTruthy is defined below in vm package
				currentFrame.ip = int(pos) - 1 // Set ip to target - 1
			} else {
				// If not jumping, advance ip past the operand
				currentFrame.ip += 2
			}
			// ip is NOT advanced after the dispatch block
		} else if opcode == compiler.OpPrint { // Added OpPrint case back
			// --- Debug Print: At start of OpPrint ---
			// Get the current frame's base pointer for context
			currentBasePtr := 0
			currentFrame := vm.currentFrame()              // Get current frame
			if vm.framesIndex > 0 && currentFrame != nil { // Add nil check
				currentBasePtr = currentFrame.basePointer
			}
			fmt.Printf("DEBUG (OpPrint): Starting - Stack: %v, sp: %d, basePtr: %d\n", vm.stack[:vm.sp], vm.sp, currentBasePtr)
			// --- End Debug Print ---

			// Operand for print arguments count is 1 byte (uint8)
			numArgs := instructions[ip+1] // Read uint8 directly

			// Arguments are on the stack in reverse order of compilation.
			// Pop them and print.
			args := make([]string, numArgs)
			// Pop from top of stack backwards
			for i := int(numArgs) - 1; i >= 0; i-- {
				args[i] = vm.pop().Inspect() // Assuming Inspect() gives a string representation
			}
			fmt.Fprintln(vm.outputWriter, strings.Join(args, " "))
		} else if opcode == compiler.OpCall {
			// Operand for call arguments count is 1 byte (uint8)
			numArgs := instructions[ip+1] // Read uint8 directly

			// --- Debug Print: Before OpCall setup ---
			fmt.Printf("DEBUG (OpCall): Before setup - Stack: %v, sp: %d\n", vm.stack[:vm.sp], vm.sp)
			// --- End Debug Print ---

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

			// --- Debug Print: Called Function Instructions ---
			fmt.Printf("DEBUG (OpCall): Called Function Instructions: %v\n", fn.Instructions)
			// --- End Debug Print ---

			if int(numArgs) != fn.NumParameters {
				return fmt.Errorf("runtime error: wrong number of arguments: expected %d, got %d", fn.NumParameters, numArgs)
			}

			// Create a new call frame for the function.
			// The base pointer for the new frame is the index on the main stack
			// where this frame's locals and arguments begin.
			// Since we copy arguments to a separate locals slice, this base pointer
			// should point to the location where the function object *was* on the main stack.
			// This space will be used for the new frame's temporary stack operations.
			newFrame := NewFrame(closure, basePointer) // basePointer is correct here.

			// Copy arguments from the stack into the new frame's locals.
			// Arguments are located on the stack starting at basePointer + 1.
			// The new frame's locals array starts at index 0.
			for i := 0; i < fn.NumParameters; i++ {
				newFrame.locals[i] = vm.stack[basePointer+1+i]
			}

			// --- FIX: Reset and clear the main stack space used by the call setup ---
			// Explicitly nil out the stack slots used by the function object and arguments.
			for i := basePointer; i < vm.sp; i++ {
				vm.stack[i] = nil
			}
			// Reset the main stack pointer to the basePointer. This makes the space
			// from basePointer onwards available for the new frame's temporary stack operations.
			vm.sp = basePointer

			// --- Debug Print: After OpCall setup ---
			// Print the absolute stack state, not relative to basePointer
			fmt.Printf("DEBUG (OpCall): After setup - Stack: %v, sp: %d, newFrameBasePtr: %d, newFrameLocals: %v\n", vm.stack[:vm.sp], vm.sp, newFrame.basePointer, newFrame.locals)
			// --- End Debug Print ---

			// Push the new frame onto the frame stack.
			vm.pushFrame(newFrame)

			// The VM will continue execution from the start of the new frame's instructions
			// in the next iteration of the main loop.
			// The stack pointer (vm.sp) has been adjusted correctly for the new frame.
			// The ip will be advanced by the loop in the next iteration, starting from the new frame's ip (-1 + 1 = 0).

		} else if opcode == compiler.OpReturnValue {
			// This opcode is emitted by the compiler for a `return expression` statement.
			// The return value is already on top of the stack.

			// --- Debug Print: Before OpReturnValue pop ---
			// Get the current frame's base pointer for context
			currentBasePtr := 0
			currentFrame := vm.currentFrame()              // Get current frame
			if vm.framesIndex > 0 && currentFrame != nil { // Add nil check
				currentBasePtr = currentFrame.basePointer
			}
			fmt.Printf("DEBUG (OpReturnValue): Before pop - Stack: %v, sp: %d, basePtr: %d\n", vm.stack[:vm.sp], vm.sp, currentBasePtr)
			// --- End Debug Print ---

			returnValue := vm.pop() // Pop the return value from the current frame's stack

			// --- Debug Print: After OpReturnValue pop, before frame pop ---
			// Get the current frame's base pointer for context (before popFrame)
			currentBasePtr = 0
			if vm.framesIndex > 0 {
				currentBasePtr = vm.currentFrame().basePointer
			}
			fmt.Printf("DEBUG (OpReturnValue): After pop - Stack: %v, sp: %d, basePtr: %d, retValue: %s\n", vm.stack[:vm.sp], vm.sp, currentBasePtr, returnValue.Inspect())
			// --- End Debug Print ---

			// Pop the current frame.
			frame := vm.popFrame() // Referring to popFrame (defined below in vm package)

			// Restore the stack pointer to the state before the function call.
			// This is the base pointer of the popped frame.
			vm.sp = frame.basePointer

			// --- Debug Print: After stack restore, before push ---
			fmt.Printf("DEBUG (OpReturnValue): After stack restore - Stack: %v, sp: %d, frameBasePtr: %d\n", vm.stack[:vm.sp], vm.sp, frame.basePointer)
			// --- End Debug Print ---

			// Push the return value onto the stack of the previous frame at the restored vm.sp.
			vm.push(returnValue)

			// --- Debug Print: After push ---
			fmt.Printf("DEBUG (OpReturnValue): After push - Stack: %v, sp: %d\n", vm.stack[:vm.sp], vm.sp)
			// --- End Debug Print ---
			// The loop will continue and process the caller's next instruction.

		} else if opcode == compiler.OpReturn {
			// This opcode is typically emitted by the compiler at the end of a function
			// or for a `return` statement without a value.
			// It implicitly returns nil. The compiler should push nil before OpReturn.

			// --- Debug Print: Before OpReturn frame pop ---
			// Get the current frame's base pointer for context
			currentBasePtr := 0
			currentFrame := vm.currentFrame()              // Get current frame
			if vm.framesIndex > 0 && currentFrame != nil { // Add nil check
				currentBasePtr = currentFrame.basePointer
			}
			fmt.Printf("DEBUG (OpReturn): Before frame pop - Stack: %v, sp: %d, basePtr: %d\n", vm.stack[:vm.sp], vm.sp, currentBasePtr)
			// --- End Debug Print ---

			// Pop the current frame.
			frame := vm.popFrame() // Referring to popFrame (defined below in vm package)

			// The return value (nil) should be on the stack, pushed by the compiler (OpNull).
			// We don't pop it here, as it will be left on the stack for the caller.

			// Restore the stack pointer to the state before the function call.
			// This is the base pointer of the popped frame.
			// The implicit nil return value should be at vm.stack[vm.sp] if compiler is correct.
			vm.sp = frame.basePointer

			// --- Debug Print: After stack restore ---
			fmt.Printf("DEBUG (OpReturn): After stack restore - Stack: %v, sp: %d, frameBasePtr: %d\n", vm.stack[:vm.sp], vm.sp, frame.basePointer)
			// --- End Debug Print ---

			// Do NOT push nil here. The compiler is responsible for pushing the return value (nil).
			// The value is already at vm.stack[frame.basePointer] if the compiler emitted OpNull before OpReturn.
			// The loop will continue and process the caller's next instruction.

		} else if opcode == compiler.OpGetIterator {
			iterable := vm.pop()
			it, err := iterable.GetIterator()
			if err != nil {
				return fmt.Errorf("runtime error: %v", err)
			}
			vm.push(it)
		} else if opcode == compiler.OpIteratorNext {
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
		} else {
			// This case should ideally not be reached.
			return fmt.Errorf("runtime error: unknown opcode: %v at %d", opcode, ip)
		}

		// Advance the instruction pointer by the size of the executed instruction.
		// This must happen *after* the dispatch block for most opcodes,
		// but *not* for jump instructions which set the ip directly.
		// Check if the opcode was a jump instruction.
		isJumpInstruction := opcode == compiler.OpJump || opcode == compiler.OpJumpNotTruthy || opcode == compiler.OpJumpTruthy

		if !isJumpInstruction {
			currentFrame.ip += (instructionSize - 1) // ip was already incremented by 1 at the start of the loop
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
		// fmt.Fprintf(os.Stderr, "runtime error: cannot perform addition on nil values\n")
		return &types.Error{Message: "cannot perform addition on nil values"} // Return specific error value
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
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: integer + %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for addition: integer + %s", right.Type())} // Return specific error value
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
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: float + %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for addition: float + %s", right.Type())} // Return specific error value
		}
	case *types.String: // Referring to String from types package
		// String concatenation
		switch right := right.(type) {
		case *types.Integer: // Referring to Integer from types package
			return &types.String{Value: fmt.Sprintf("%s%d", left.Value, right.Value)} // Referring to String from types package
		case *types.Float: // Referring to Float from types package
			return &types.String{Value: fmt.Sprintf("%s%f", left.Value, right.Value)} // Referring to String from types package
		case *types.String: // Referring to String from types package
			return &types.String{Value: left.Value + right.Value} // Referring to String from types package
		case *types.Boolean: // Referring to Boolean from types package
			return &types.String{Value: fmt.Sprintf("%s%t", left.Value, right.Value)} // Referring to String from types package
		case *types.Nil: // Referring to Nil from types package
			return &types.String{Value: left.Value + "nil"} // Referring to String from types package
		default:
			// For other types, use their Inspect() representation
			return &types.String{Value: left.Value + right.Inspect()} // Referring to String from types package
		}
	default:
		// Incompatible types for addition
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for addition: %s + %s\n", left.Type(), right.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported types for addition: %s + %s", left.Type(), right.Type())} // Return specific error value
	}
}

// subtract performs subtraction.
func (vm *VM) subtract(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values first
	if left == nil || right == nil {
		// fmt.Fprintf(os.Stderr, "runtime error: cannot perform subtraction on nil values\n")
		return &types.Error{Message: "cannot perform subtraction on nil values"}
	}

	// Handle subtraction based on types
	switch left := left.(type) {
	case *types.Integer:
		switch right := right.(type) {
		case *types.Integer:
			return &types.Integer{Value: left.Value - right.Value}
		case *types.Float:
			return &types.Float{Value: float64(left.Value) - right.Value}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for subtraction: integer - %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for subtraction: integer - %s", right.Type())}
		}
	case *types.Float:
		switch right := right.(type) {
		case *types.Integer:
			return &types.Float{Value: left.Value - float64(right.Value)}
		case *types.Float:
			return &types.Float{Value: left.Value - right.Value}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for subtraction: float - %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for subtraction: float - %s", right.Type())}
		}
	default:
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for subtraction: %s - %s\n", left.Type(), right.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported types for subtraction: %s - %s", left.Type(), right.Type())}
	}
}

// multiply performs multiplication.
func (vm *VM) multiply(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values first
	if left == nil || right == nil {
		// fmt.Fprintf(os.Stderr, "runtime error: cannot perform multiplication on nil values\n")
		return &types.Error{Message: "cannot perform multiplication on nil values"}
	}

	// Handle multiplication based on types
	switch left := left.(type) {
	case *types.Integer:
		switch right := right.(type) {
		case *types.Integer:
			return &types.Integer{Value: left.Value * right.Value}
		case *types.Float:
			return &types.Float{Value: float64(left.Value) * right.Value}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for multiplication: integer * %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for multiplication: integer * %s", right.Type())}
		}
	case *types.Float:
		switch right := right.(type) {
		case *types.Integer:
			return &types.Float{Value: left.Value * float64(right.Value)}
		case *types.Float:
			return &types.Float{Value: left.Value * right.Value}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for multiplication: float * %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for multiplication: float * %s", right.Type())}
		}
	default:
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for multiplication: %s * %s\n", left.Type(), right.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported types for multiplication: %s * %s", left.Type(), right.Type())}
	}
}

// divide performs division.
func (vm *VM) divide(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values first
	if left == nil || right == nil {
		// fmt.Fprintf(os.Stderr, "runtime error: cannot perform division on nil values\n")
		return &types.Error{Message: "cannot perform division on nil values"}
	}

	// Handle division based on types
	switch left := left.(type) {
	case *types.Integer:
		switch right := right.(type) {
		case *types.Integer:
			if right.Value == 0 {
				// fmt.Fprintf(os.Stderr, "runtime error: division by zero\n")
				return &types.Error{Message: "division by zero"} // Or a specific Error type
			}
			return &types.Integer{Value: left.Value / right.Value} // Integer division results in float
		case *types.Float:
			if right.Value == 0.0 {
				// fmt.Fprintf(os.Stderr, "runtime error: division by zero\n")
				return &types.Error{Message: "division by zero"} // Or a specific Error type
			}
			return &types.Float{Value: float64(left.Value) / right.Value}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for division: integer / %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for division: integer / %s", right.Type())}
		}
	case *types.Float:
		switch right := right.(type) {
		case *types.Integer:
			if right.Value == 0 {
				// fmt.Fprintf(os.Stderr, "runtime error: division by zero\n")
				return &types.Error{Message: "division by zero"} // Or a specific Error type
			}
			return &types.Float{Value: left.Value / float64(right.Value)}
		case *types.Float:
			if right.Value == 0.0 {
				// fmt.Fprintf(os.Stderr, "runtime error: division by zero\n")
				return &types.Error{Message: "division by zero"} // Or a specific Error type
			}
			return &types.Float{Value: left.Value / right.Value}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for division: float / %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for division: float / %s", right.Type())}
		}
	default:
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for division: %s / %s\n", left.Type(), right.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported types for division: %s / %s", left.Type(), right.Type())}
	}
}

// modulo performs the modulo operation.
func (vm *VM) modulo(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values first
	if left == nil || right == nil {
		// fmt.Fprintf(os.Stderr, "runtime error: cannot perform modulo on nil values\n")
		return &types.Error{Message: "cannot perform modulo on nil values"}
	}

	// Modulo is typically only defined for integers
	leftInt, okLeft := left.(*types.Integer)
	rightInt, okRight := right.(*types.Integer)

	if !okLeft || !okRight {
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for modulo: %s %% %s\n", left.Type(), right.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported types for modulo: %s %% %s", left.Type(), right.Type())}
	}

	if rightInt.Value == 0 {
		// fmt.Fprintf(os.Stderr, "runtime error: modulo by zero\n")
		return &types.Error{Message: "modulo by zero"} // Or a specific Error type
	}

	return &types.Integer{Value: leftInt.Value % rightInt.Value}
}

// power performs exponentiation.
func (vm *VM) power(left, right types.Value) types.Value { // Referring to Value from types package
	// Check for nil values first
	if left == nil || right == nil {
		// fmt.Fprintf(os.Stderr, "runtime error: cannot perform power on nil values\n")
		return &types.Error{Message: "cannot perform power on nil values"}
	}

	// Handle power based on types
	switch left := left.(type) {
	case *types.Integer:
		switch right := right.(type) {
		case *types.Integer:
			// Integer power
			return &types.Float{Value: math.Pow(float64(left.Value), float64(right.Value))} // Result is float
		case *types.Float:
			// Integer base, float exponent
			return &types.Float{Value: math.Pow(float64(left.Value), right.Value)}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for power: integer ^ %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for power: integer ^ %s", right.Type())}
		}
	case *types.Float:
		switch right := right.(type) {
		case *types.Integer:
			// Float base, integer exponent
			return &types.Float{Value: math.Pow(left.Value, float64(right.Value))}
		case *types.Float:
			// Float power
			return &types.Float{Value: math.Pow(left.Value, right.Value)}
		default:
			// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for power: float ^ %s\n", right.Type())
			return &types.Error{Message: fmt.Sprintf("unsupported types for power: float ^ %s", right.Type())}
		}
	default:
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported types for power: %s ^ %s\n", left.Type(), right.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported types for power: %s ^ %s", left.Type(), right.Type())}
	}
}

// negate performs unary negation.
func (vm *VM) negate(operand types.Value) types.Value { // Referring to Value from types package
	if operand == nil {
		// fmt.Fprintf(os.Stderr, "runtime error: cannot negate nil value\n")
		return &types.Error{Message: "cannot negate nil value"}
	}

	switch operand := operand.(type) {
	case *types.Integer:
		return &types.Integer{Value: -operand.Value}
	case *types.Float:
		return &types.Float{Value: -operand.Value}
	default:
		// fmt.Fprintf(os.Stderr, "runtime error: unsupported type for negation: -%s\n", operand.Type())
		return &types.Error{Message: fmt.Sprintf("unsupported type for negation: -%s", operand.Type())}
	}
}

// logicalNot performs logical negation.
func (vm *VM) logicalNot(operand types.Value) types.Value { // Referring to Value from types package
	// The isTruthy helper determines the boolean value of any type.
	return &types.Boolean{Value: !isTruthy(operand)} // Referring to Boolean from types package
}

// equal checks for equality.
func (vm *VM) equal(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Equals method defined on the Value interface.
	// This allows each type to define its own equality logic.
	return &types.Boolean{Value: left.Equals(right)} // Referring to Boolean from types package
}

// notEqual checks for inequality.
func (vm *VM) notEqual(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Equals method and negate the result.
	return &types.Boolean{Value: !left.Equals(right)} // Referring to Boolean from types package
}

// greaterThan checks if left > right.
func (vm *VM) greaterThan(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method defined on the Value interface.
	cmp, err := left.Compare(right)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Error{Message: fmt.Sprintf("comparison error: %v", err)} // Or a specific Error type
	}
	return &types.Boolean{Value: cmp > 0} // Referring to Boolean from types package
}

// lessThan checks if left < right.
func (vm *VM) lessThan(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method.
	cmp, err := left.Compare(right)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Error{Message: fmt.Sprintf("comparison error: %v", err)} // Or a specific Error type
	}
	return &types.Boolean{Value: cmp < 0} // Referring to Boolean from types package
}

// greaterEqual checks if left >= right.
func (vm *VM) greaterEqual(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method.
	cmp, err := left.Compare(right)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Error{Message: fmt.Sprintf("comparison error: %v", err)} // Or a specific Error type
	}
	return &types.Boolean{Value: cmp >= 0} // Referring to Boolean from types package
}

// lessEqual checks if left <= right.
func (vm *VM) lessEqual(left, right types.Value) types.Value { // Referring to Value from types package
	// Use the Compare method.
	cmp, err := left.Compare(right)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "runtime error: %v\n", err)
		return &types.Error{Message: fmt.Sprintf("comparison error: %v", err)} // Or a specific Error type
	}
	return &types.Boolean{Value: cmp <= 0} // Referring to Boolean from types package
}

// isTruthy determines the boolean value of any value.
func isTruthy(obj types.Value) bool { // Referring to Value from types package
	if obj == nil {
		return false
	}
	switch obj := obj.(type) {
	case *types.Boolean: // Referring to Boolean from types package
		return obj.Value
	case *types.Nil: // Referring to Nil from types package
		return false
	case *types.Integer: // Referring to Integer from types package
		return obj.Value != 0
	case *types.Float: // Referring to Float from types package
		return obj.Value != 0.0 // Consider NaN/Inf handling if needed
	case *types.String: // Referring to String from types package
		return obj.Value != ""
	case *types.List: // Referring to List from types package
		return len(obj.Elements) > 0
	case *types.Table: // Referring to Table from types package
		return len(obj.Pairs) > 0 // Check number of pairs
	// FIX: Change the type switch case from *types.Iterator to the concrete types that implement it.
	case *types.StringIterator, *types.ListIterator, *types.TableIterator: // Iterators are generally truthy if they are not exhausted (though this check doesn't verify exhaustion)
		return true
	case *types.CompiledFunction, *types.Closure: // Functions, Closures are generally truthy
		return true
	case *types.Error: // Errors are generally falsy
		return false
	default:
		// Unknown types are considered falsy by default
		return false
	}
}

// SetOutputWriter sets the writer for print statements.
func (vm *VM) SetOutputWriter(writer io.Writer) {
	vm.outputWriter = writer
}
