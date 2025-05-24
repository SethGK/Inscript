// Package vm implements the Inscript virtual machine.
package vm

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings" // Added for strings.Join

	"github.com/SethGK/Inscript/internal/compiler"
	"github.com/SethGK/Inscript/internal/types"
)

const StackSize = 2048
const MaxFrames = 1024

// VM represents the Inscript Virtual Machine.
type VM struct {
	constants []types.Value

	stack []types.Value
	sp    int // Stack pointer: points to the next free slot on the stack

	globals []types.Value

	frames      []*Frame // Call frames for function execution
	framesIndex int      // Current frame index - points to the next free frame slot

	outputWriter io.Writer
}

// Frame represents a single call frame for function execution.
type Frame struct {
	closure     *types.Closure // The closure being executed
	ip          int            // Instruction pointer: points to the next instruction to execute
	basePointer int            // Base pointer: points to the first slot in the stack used by this frame (for locals and arguments)
}

// NewFrame creates a new call frame.
// basePointer is the stack index where this frame's locals and arguments begin.
func NewFrame(closure *types.Closure, basePointer int) *Frame {
	return &Frame{
		closure:     closure,
		ip:          -1, // Start at -1 so the first instruction is at index 0 after increment
		basePointer: basePointer,
	}
}

// Instructions returns the instructions of the compiled function in this frame.
func (f *Frame) Instructions() compiler.Instructions {
	return f.closure.Fn.Instructions
}

// New creates a new VM instance.
// It takes the bytecode compiled by the compiler package.
func New(bytecode *compiler.Bytecode) *VM {
	// The main program is treated as a function (the entry point).
	mainFn := &types.CompiledFunction{
		Instructions:  bytecode.Instructions,
		NumLocals:     bytecode.NumLocals,
		NumParameters: 0, // Main program has no parameters
		FreeCount:     0,
	}
	mainClosure := &types.Closure{Fn: mainFn, Free: []types.Value{}}
	mainFrame := NewFrame(mainClosure, 0) // Base pointer for main program is 0

	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame

	return &VM{
		constants:    bytecode.Constants,
		stack:        make([]types.Value, StackSize),
		sp:           0,
		globals:      make([]types.Value, bytecode.NumGlobals),
		frames:       frames,
		framesIndex:  1,         // Start with the main frame at index 0, next frame will be at index 1
		outputWriter: os.Stdout, // Default output to stdout
	}
}

// currentFrame returns the currently executing call frame.
func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.framesIndex-1]
}

// pushFrame pushes a new call frame onto the frame stack.
func (vm *VM) pushFrame(frame *Frame) error {
	if vm.framesIndex >= MaxFrames {
		return types.NewError("frame stack overflow")
	}
	vm.frames[vm.framesIndex] = frame
	vm.framesIndex++
	return nil
}

// popFrame pops the current call frame from the frame stack.
func (vm *VM) popFrame() *Frame {
	if vm.framesIndex == 0 {
		return nil // Should not happen if the main frame is always present
	}
	vm.framesIndex--
	frame := vm.frames[vm.framesIndex]
	vm.frames[vm.framesIndex] = nil // Clear the reference
	return frame
}

// StackTop returns the value at the top of the stack without popping it.
func (vm *VM) StackTop() types.Value {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) push(obj types.Value) error {
	if vm.sp >= StackSize { // StackSize is your max stack capacity
		return types.NewError("stack overflow")
	}
	vm.stack[vm.sp] = obj
	vm.sp++
	fmt.Printf("DEBUG: PUSHED %s to SP=%d\n", obj.Inspect(), vm.sp-1) // vm.sp-1 is the index it was pushed to
	return nil
}

func (vm *VM) pop() (types.Value, error) {
	if vm.sp == 0 {
		return nil, types.NewError("stack empty")
	}
	vm.sp--
	popped := vm.stack[vm.sp]
	vm.stack[vm.sp] = nil                                                // Clear reference to allow GC
	fmt.Printf("DEBUG: POPPED %s from SP=%d\n", popped.Inspect(), vm.sp) // vm.sp is the index it was popped from
	return popped, nil
}

// Run executes the compiled bytecode.
func (vm *VM) Run() error {
	var err error

	for vm.framesIndex > 0 {
		currentFrame := vm.currentFrame()
		instructions := currentFrame.Instructions()

		currentFrame.ip++
		ip := currentFrame.ip

		if ip >= len(instructions) {
			if vm.framesIndex > 1 {
				// If not the main program, pop the frame and push nil as return value.
				// The stack pointer should be reset to the base pointer of the previous frame.
				frame := vm.popFrame()
				vm.sp = frame.basePointer                     // Reset sp to where the callee was
				if err := vm.push(&types.Nil{}); err != nil { // Push return value
					return err
				}
				continue
			}
			break // Exit loop if main program finishes
		}

		opcode := compiler.ReadOpcode(instructions, ip)

		switch opcode {
		case compiler.OpConstant:
			constantIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			err = vm.push(vm.constants[constantIndex])
			if err != nil {
				return err
			}

		case compiler.OpPop:
			if _, err := vm.pop(); err != nil {
				return err
			}

		case compiler.OpTrue:
			err = vm.push(types.NewBoolean(true))
			if err != nil {
				return err
			}
		case compiler.OpFalse:
			err = vm.push(types.NewBoolean(false))
			if err != nil {
				return err
			}
		case compiler.OpNull:
			err = vm.push(&types.Nil{})
			if err != nil {
				return err
			}

		case compiler.OpAdd, compiler.OpSub, compiler.OpMul, compiler.OpDiv, compiler.OpMod, compiler.OpPow, compiler.OpIDiv:
			err = vm.executeBinaryOperation(opcode)
			if err != nil {
				return err
			}

		case compiler.OpBitAnd, compiler.OpBitOr, compiler.OpBitXor, compiler.OpShl, compiler.OpShr:
			err = vm.executeBitwiseOperation(opcode)
			if err != nil {
				return err
			}

		case compiler.OpBang: // Logical NOT
			operand, err := vm.pop()
			if err != nil {
				return err
			}
			err = vm.push(types.NewBoolean(!isTruthy(operand)))
			if err != nil {
				return err
			}

		case compiler.OpBitNot: // Bitwise NOT
			operand, err := vm.pop()
			if err != nil {
				return err
			}
			result, bitNotErr := vm.executeUnaryBitwiseNot(operand)
			if bitNotErr != nil {
				return bitNotErr
			}
			err = vm.push(result)
			if err != nil {
				return err
			}

		case compiler.OpMinus: // Negation
			operand, err := vm.pop()
			if err != nil {
				return err
			}
			switch val := operand.(type) {
			case *types.Integer:
				if err := vm.push(types.NewInteger(-val.Value)); err != nil {
					return err
				}
			case *types.Float:
				if err := vm.push(types.NewFloat(-val.Value)); err != nil {
					return err
				}
			default:
				return types.NewError("unsupported type for negation: %s", operand.Type())
			}

		case compiler.OpEqual, compiler.OpNotEqual, compiler.OpGreaterThan, compiler.OpLessThan, compiler.OpGreaterEqual, compiler.OpLessEqual:
			err = vm.executeComparisonOperation(opcode)
			if err != nil {
				return err
			}

		case compiler.OpJumpNotTruthy:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			condition := vm.StackTop() // Peek the condition
			if !isTruthy(condition) {
				currentFrame.ip += offset
			}
			// The value is popped by a subsequent OpPop in the compiler's output for logical AND/OR

		case compiler.OpJump:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			currentFrame.ip += offset

		case compiler.OpJumpTruthy:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			condition := vm.StackTop() // Peek the condition
			if isTruthy(condition) {
				currentFrame.ip += offset
			}
			// The value is popped by a subsequent OpPop in the compiler's output for logical AND/OR

		case compiler.OpSetGlobal:
			globalIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			value, err := vm.pop()
			if err != nil {
				return err
			}
			if int(globalIndex) >= len(vm.globals) {
				return types.NewError("global variable index out of bounds: %d (max %d)", globalIndex, len(vm.globals)-1)
			}
			vm.globals[globalIndex] = value

		case compiler.OpGetGlobal:
			globalIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			if int(globalIndex) >= len(vm.globals) {
				return types.NewError("global variable index out of bounds: %d (max %d)", globalIndex, len(vm.globals)-1)
			}
			err = vm.push(vm.globals[globalIndex])
			if err != nil {
				return err
			}

		case compiler.OpSetLocal:
			localIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			value, err := vm.pop()
			if err != nil {
				return err
			}
			vm.stack[currentFrame.basePointer+localIndex] = value
			fmt.Printf("DEBUG: OpSetLocal %d setting local %d to %s (at stack index %d)\n", localIndex, localIndex, value.Inspect(), currentFrame.basePointer+localIndex)

		case compiler.OpGetLocal:
			localIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			valToPush := vm.stack[currentFrame.basePointer+localIndex]
			fmt.Printf("DEBUG: OpGetLocal %d getting local %d, value: %s (from stack index %d)\n", localIndex, localIndex, valToPush.Inspect(), currentFrame.basePointer+localIndex)
			err = vm.push(valToPush)
			if err != nil {
				return err
			}

		case compiler.OpSetFree: // New opcode for setting free variables
			freeIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			value, err := vm.pop()
			if err != nil {
				return err
			}
			if int(freeIndex) >= len(currentFrame.closure.Free) {
				return types.NewError("free variable index %d out of bounds for closure with %d free variables. This indicates a compiler bug.", freeIndex, len(currentFrame.closure.Free))
			}
			currentFrame.closure.Free[freeIndex] = value

		case compiler.OpGetFree:
			freeIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			if int(freeIndex) >= len(currentFrame.closure.Free) {
				return types.NewError("free variable index %d out of bounds for closure with %d free variables. This indicates a compiler bug.", freeIndex, len(currentFrame.closure.Free))
			}
			err = vm.push(currentFrame.closure.Free[freeIndex])
			if err != nil {
				return err
			}

		case compiler.OpClosure:
			constIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			freeCount, bytesRead2 := compiler.ReadOperand(instructions, ip+1+bytesRead, 1)
			currentFrame.ip += bytesRead + bytesRead2

			fnVal := vm.constants[constIndex]
			compiledFn, ok := fnVal.(*types.CompiledFunction)
			if !ok {
				return types.NewError("constant %d is not a CompiledFunction, got %s", constIndex, fnVal.Type())
			}

			frees := make([]types.Value, freeCount)
			// Free variables are pushed onto the stack before OpClosure.
			// Pop them in reverse order to maintain their original order.
			for i := freeCount - 1; i >= 0; i-- {
				freedVal, err := vm.pop()
				if err != nil {
					return err
				}
				frees[i] = freedVal
			}
			clo := &types.Closure{Fn: compiledFn, Free: frees}
			err = vm.push(clo)
			if err != nil {
				return err
			}

		case compiler.OpArray:
			numElements, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			elements := make([]types.Value, numElements)
			// Elements are pushed in order, so pop them in reverse to build the list
			for i := numElements - 1; i >= 0; i-- {
				el, err := vm.pop()
				if err != nil {
					return err
				}
				elements[i] = el
			}
			err = vm.push(types.NewList(elements...))
			if err != nil {
				return err
			}

		case compiler.OpTable:
			numPairs, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead

			pairs := make([]types.TablePair, numPairs)
			// Pairs are pushed as (key, value) pairs. Pop value then key.
			// Pop in reverse order of how they were pushed.
			for i := numPairs - 1; i >= 0; i-- {
				value, err := vm.pop()
				if err != nil {
					return err
				}
				keyVal, err := vm.pop()
				if err != nil {
					return err
				}
				keyStr, ok := keyVal.(*types.String)
				if !ok {
					return types.NewError("table key must be a string, got %s", keyVal.Type())
				}
				pairs[i] = types.TablePair{Key: keyStr.Value, Value: value}
			}
			err = vm.push(types.NewTable(pairs))
			if err != nil {
				return err
			}

		case compiler.OpIndex:
			index, err := vm.pop()
			if err != nil {
				return err
			}
			aggregate, err := vm.pop()
			if err != nil {
				return err
			}
			result, getErr := aggregate.GetIndex(index)
			if getErr != nil {
				return types.NewError("runtime error: %s", getErr.Error())
			}
			err = vm.push(result)
			if err != nil {
				return err
			}

		case compiler.OpSetIndex:
			value, err := vm.pop()
			if err != nil {
				return err
			}
			index, err := vm.pop()
			if err != nil {
				return err
			}
			aggregate, err := vm.pop()
			if err != nil {
				return err
			}
			setErr := aggregate.SetIndex(index, value)
			if setErr != nil {
				return types.NewError("runtime error: %s", setErr.Error())
			}
			// SetIndex typically leaves the assigned value on stack
			err = vm.push(value)
			if err != nil {
				return err
			}

		case compiler.OpPrint:
			numExprs, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			fmt.Printf("DEBUG: Stack before OpPrint (SP=%d, numExprs=%d):\n", vm.sp, numExprs)
			for i := 0; i < vm.sp; i++ {
				fmt.Printf("  [%d]: %+v\n", i, vm.stack[i].Inspect())
			}
			args := make([]string, numExprs)
			// Pop elements from the stack (they come off right-to-left)
			// and place them into the 'args' slice in the correct left-to-right order.
			for i := 0; i < int(numExprs); i++ {
				poppedVal, err := vm.pop()
				if err != nil {
					return err
				}
				args[int(numExprs)-1-i] = poppedVal.Inspect()
			}
			// Join the arguments with a space and print the single line.
			fmt.Fprintln(vm.outputWriter, strings.Join(args, " "))

		case compiler.OpGetIter:
			iterable, err := vm.pop()
			if err != nil {
				return err
			}
			iter, getIterErr := iterable.GetIterator()
			if getIterErr != nil {
				return types.NewError("runtime error: %s", getIterErr.Error())
			}
			if err := vm.push(iter); err != nil {
				return err
			}

		case compiler.OpIterNext:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead

			// Peek at the iterator (don't pop it yet)
			if vm.sp == 0 {
				return types.NewError("stack underflow: no iterator on stack")
			}

			iteratorVal := vm.stack[vm.sp-1] // Peek, don't pop
			iterator, ok := iteratorVal.(types.Iterator)
			if !ok {
				return types.NewError("object is not an iterator: %s", iteratorVal.Type())
			}

			nextVal, hasNext, iterErr := iterator.Next()
			if iterErr != nil {
				return types.NewError("runtime error during iteration: %s", iterErr.Error())
			}

			if !hasNext {
				_, err := vm.pop()
				if err != nil {
					return err
				}
				currentFrame.ip += int(offset)
				fmt.Printf("DEBUG: OpIterNext finished, jumping to %d. Stack top: %v\n", currentFrame.ip, vm.stack[vm.sp-1]) // Add this
				continue
			} else {

				if err := vm.push(nextVal); err != nil {
					return err
				}
			}

		case compiler.OpCall:
			numArgs, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead

			calleePos := vm.sp - int(numArgs) - 1
			if calleePos < 0 || calleePos >= vm.sp {
				return types.NewError("runtime error: invalid callee position on stack. SP=%d, NumArgs=%d", vm.sp, numArgs)
			}
			callee := vm.stack[calleePos]
			closure, ok := callee.(*types.Closure)
			if !ok {
				return types.NewError("call target is not a function or closure: %s", callee.Type())
			}
			if int(numArgs) != closure.Fn.NumParameters {
				return types.NewError("wrong number of arguments: expected %d, got %d",
					closure.Fn.NumParameters, numArgs)
			}

			argStart := calleePos + 1
			newFrame := NewFrame(closure, argStart)
			if err := vm.pushFrame(newFrame); err != nil {
				return err
			}

			vm.sp = newFrame.basePointer + closure.Fn.NumLocals

		case compiler.OpReturnValue:
			returnValue, err := vm.pop()
			if err != nil {
				return err
			}
			poppedFrame := vm.popFrame()

			if vm.framesIndex > 0 {

				vm.sp = poppedFrame.basePointer - 1
			} else {

				vm.sp = poppedFrame.basePointer
			}

			fmt.Printf("DEBUG: OpReturnValue is pushing: %+v\n", returnValue)
			if err := vm.push(returnValue); err != nil {
				return err
			}

		case compiler.OpReturn:
			poppedFrame := vm.popFrame()

			if vm.framesIndex > 0 {

				vm.sp = poppedFrame.basePointer - 1
			} else {
				vm.sp = poppedFrame.basePointer
			}

			if err := vm.push(&types.Nil{}); err != nil {
				return err
			}

		case compiler.OpImport: // Handle import statement
			pathVal, err := vm.pop()
			if err != nil {
				return err
			}
			pathStr, ok := pathVal.(*types.String)
			if !ok {
				return types.NewError("import path must be a string, got %s", pathVal.Type())
			}
			// TODO: Implement actual module loading/compilation/execution here.
			// This would typically involve a module cache to avoid re-importing.
			// For now, we'll just print a debug message and push nil.
			fmt.Fprintf(vm.outputWriter, "DEBUG: Attempting to import module: %s (actual import logic not yet implemented)\n", pathStr.Value)
			err = vm.push(&types.Nil{}) // Push nil for now, representing the imported module
			if err != nil {
				return err
			}

		default:
			return types.NewError("unknown opcode: %s", opcode.String()) // Use opcode.String() for better error messages
		}
	}
	return nil
}

// LastPoppedStackElem returns the last element popped from the stack.
// Note: This method's behavior is tricky. If you need the value *after* a pop,
// it's already returned by `pop()`. This method returns the element at `vm.stack[vm.sp]`
// which is the new top of the stack after `sp` has been decremented by `pop()`.
// For practical purposes, it's often better to capture the return value of `pop()`.
func (vm *VM) LastPoppedStackElem() types.Value {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp]
}

// isTruthy determines the truthiness of a value.
func isTruthy(obj types.Value) bool {
	if b, ok := obj.(*types.Boolean); ok {
		return b.Value
	}
	if _, isNil := obj.(*types.Nil); isNil {
		return false
	}
	return true // All other values are truthy
}

// executeBinaryOperation handles arithmetic and power operations.
func (vm *VM) executeBinaryOperation(op compiler.Opcode) error {
	right, err := vm.pop()
	if err != nil {
		return err
	}
	left, err := vm.pop()
	if err != nil {
		return err
	}

	// --- NEW LIST CONCATENATION LOGIC ---
	if op == compiler.OpAdd && left.Type() == types.LIST_OBJ && right.Type() == types.LIST_OBJ {
		leftList := left.(*types.List)
		rightList := right.(*types.List)

		// Create a new slice for the combined elements
		newElements := make([]types.Value, len(leftList.Elements)+len(rightList.Elements))
		copy(newElements, leftList.Elements)
		copy(newElements[len(leftList.Elements):], rightList.Elements)

		return vm.push(types.NewList(newElements...)) // FIX: Added '...' spread operator
	}
	// --- END NEW LIST CONCATENATION LOGIC ---

	// String concatenation for OpAdd (keep this, it's correct)
	if op == compiler.OpAdd && left.Type() == types.STRING_OBJ && right.Type() == types.STRING_OBJ {
		leftStr := left.(*types.String).Value
		rightStr := right.(*types.String).Value
		return vm.push(types.NewString(leftStr + rightStr))
	}

	// Type checking for numeric operations
	if (left.Type() != types.INTEGER_OBJ && left.Type() != types.FLOAT_OBJ) ||
		(right.Type() != types.INTEGER_OBJ && right.Type() != types.FLOAT_OBJ) {
		// If it's not numbers and not handled above (like lists or strings), it's an error
		return types.NewError("type mismatch for %s: %s %s %s", op.String(), left.Type(), op.String(), right.Type())
	}

	var result types.Value

	switch op {
	case compiler.OpAdd: // Now this will only handle numbers (lists and strings are handled above)
		if left.Type() == types.INTEGER_OBJ && right.Type() == types.INTEGER_OBJ {
			result = types.NewInteger(left.(*types.Integer).Value + right.(*types.Integer).Value)
		} else {
			lVal := toFloat64(left)
			rVal := toFloat64(right)
			result = types.NewFloat(lVal + rVal)
		}
	case compiler.OpSub:
		if left.Type() == types.INTEGER_OBJ && right.Type() == types.INTEGER_OBJ {
			result = types.NewInteger(left.(*types.Integer).Value - right.(*types.Integer).Value)
		} else {
			lVal := toFloat64(left)
			rVal := toFloat64(right)
			result = types.NewFloat(lVal - rVal)
		}
	case compiler.OpMul:
		if left.Type() == types.INTEGER_OBJ && right.Type() == types.INTEGER_OBJ {
			result = types.NewInteger(left.(*types.Integer).Value * right.(*types.Integer).Value)
		} else {
			lVal := toFloat64(left)
			rVal := toFloat64(right)
			result = types.NewFloat(lVal * rVal)
		}
	case compiler.OpDiv:
		if left.Type() == types.INTEGER_OBJ && right.Type() == types.INTEGER_OBJ {
			// Integer division with float result if not evenly divisible (common in many languages)
			if right.(*types.Integer).Value == 0 {
				return types.NewError("division by zero")
			}
			result = types.NewFloat(float64(left.(*types.Integer).Value) / float64(right.(*types.Integer).Value))
		} else {
			lVal := toFloat64(left)
			rVal := toFloat64(right)
			if rVal == 0.0 {
				return types.NewError("division by zero")
			}
			result = types.NewFloat(lVal / rVal)
		}
	case compiler.OpMod:
		if left.Type() == types.INTEGER_OBJ && right.Type() == types.INTEGER_OBJ {
			if right.(*types.Integer).Value == 0 {
				return types.NewError("modulo by zero")
			}
			result = types.NewInteger(left.(*types.Integer).Value % right.(*types.Integer).Value)
		} else {
			// Modulo for floats is generally fmod in math libraries
			lVal := toFloat64(left)
			rVal := toFloat64(right)
			if rVal == 0.0 {
				return types.NewError("modulo by zero")
			}
			result = types.NewFloat(math.Mod(lVal, rVal))
		}
	case compiler.OpPow:
		lVal := toFloat64(left)
		rVal := toFloat64(right)
		result = types.NewFloat(math.Pow(lVal, rVal))
	case compiler.OpIDiv: // Integer division (floor division)
		lVal := toFloat64(left)
		rVal := toFloat64(right)
		if rVal == 0.0 {
			return types.NewError("integer division by zero")
		}
		result = types.NewInteger(int64(math.Floor(lVal / rVal)))
	default:
		return types.NewError("unknown operator for binary operation: %s", op.String())
	}

	return vm.push(result)
}

// executeBitwiseOperation handles binary bitwise operations.
func (vm *VM) executeBitwiseOperation(op compiler.Opcode) error {
	right, err := vm.pop()
	if err != nil {
		return err
	}
	left, err := vm.pop()
	if err != nil {
		return err
	}

	if left.Type() != types.INTEGER_OBJ || right.Type() != types.INTEGER_OBJ {
		return types.NewError("type mismatch for bitwise operation %s: expected Integer %s Integer, got %s %s %s", op.String(), op.String(), left.Type(), op.String(), right.Type())
	}

	lVal := left.(*types.Integer).Value
	rVal := right.(*types.Integer).Value
	var result int64

	switch op {
	case compiler.OpBitAnd:
		result = lVal & rVal
	case compiler.OpBitOr:
		result = lVal | rVal
	case compiler.OpBitXor:
		result = lVal ^ rVal
	case compiler.OpShl:
		if rVal < 0 {
			return types.NewError("negative shift amount for left shift")
		}
		result = lVal << uint(rVal)
	case compiler.OpShr:
		if rVal < 0 {
			return types.NewError("negative shift amount for right shift")
		}
		result = lVal >> uint(rVal)
	default:
		return types.NewError("unsupported bitwise operation: %s", op)
	}
	return vm.push(types.NewInteger(result))
}

// executeUnaryBitwiseNot handles the unary bitwise NOT operation.
func (vm *VM) executeUnaryBitwiseNot(operand types.Value) (types.Value, error) {
	if operand.Type() != types.INTEGER_OBJ {
		return nil, types.NewError("type mismatch for bitwise NOT: expected Integer, got %s", operand.Type())
	}
	val := operand.(*types.Integer).Value
	return types.NewInteger(^val), nil
}

// toFloat64 safely converts an Integer or Float Value to float64.
func toFloat64(obj types.Value) float64 {
	if i, ok := obj.(*types.Integer); ok {
		return float64(i.Value)
	}
	if f, ok := obj.(*types.Float); ok {
		return f.Value
	}
	return 0.0 // Should be caught by type checks earlier
}

// executeComparisonOperation handles comparison operations.
func (vm *VM) executeComparisonOperation(op compiler.Opcode) error {
	right, err := vm.pop()
	if err != nil {
		return err
	}
	left, err := vm.pop()
	if err != nil {
		return err
	}

	var result bool
	var errCmp error

	switch op {
	case compiler.OpEqual:
		result = left.Equals(right)
	case compiler.OpNotEqual:
		result = !left.Equals(right)
	case compiler.OpGreaterThan:
		var cmp int
		cmp, errCmp = left.Compare(right)
		if errCmp != nil {
			return types.NewError("runtime error: %s", errCmp.Error())
		}
		result = cmp > 0
	case compiler.OpLessThan:
		var cmp int
		cmp, errCmp = left.Compare(right)
		if errCmp != nil {
			return types.NewError("runtime error: %s", errCmp.Error())
		}
		result = cmp < 0
	case compiler.OpGreaterEqual: // New comparison
		var cmp int
		cmp, errCmp = left.Compare(right)
		if errCmp != nil {
			return types.NewError("runtime error: %s", errCmp.Error())
		}
		result = cmp >= 0
	case compiler.OpLessEqual: // New comparison
		var cmp int
		cmp, errCmp = left.Compare(right)
		if errCmp != nil {
			return types.NewError("runtime error: %s", errCmp.Error())
		}
		result = cmp <= 0
	default:
		return types.NewError("unsupported comparison operation: %s", op)
	}
	return vm.push(types.NewBoolean(result))
}
