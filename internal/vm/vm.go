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

type VM struct {
	constants []types.Value

	stack []types.Value
	sp    int // Stack pointer

	globals []types.Value

	frames      []*Frame // Call frames for function execution
	framesIndex int

	outputWriter io.Writer
}

// Frame represents a single call frame for function execution.
type Frame struct {
	closure     *types.Closure // The closure being executed
	ip          int            // Instruction pointer
	basePointer int
}

// NewFrame creates a new call frame.
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

// push pushes a value onto the stack.
func (vm *VM) push(obj types.Value) error {
	if vm.sp >= StackSize {
		return types.NewError("stack overflow")
	}
	vm.stack[vm.sp] = obj
	vm.sp++
	return nil
}

// pop pops a value from the stack.
func (vm *VM) pop() types.Value {
	if vm.sp == 0 {
		return types.NewError("stack underflow: attempted to pop from empty stack")
	}
	vm.sp--
	obj := vm.stack[vm.sp]
	vm.stack[vm.sp] = nil
	return obj
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
				// If not the main program, pop the frame and push nil as return value
				frame := vm.popFrame()
				vm.sp = frame.basePointer - 1 // Clean up stack for arguments and locals
				if err := vm.push(&types.Nil{}); err != nil {
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
			vm.pop()

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
			operand := vm.pop()
			err = vm.push(types.NewBoolean(!isTruthy(operand)))
			if err != nil {
				return err
			}

		case compiler.OpBitNot: // Bitwise NOT
			operand := vm.pop()
			result, bitNotErr := vm.executeUnaryBitwiseNot(operand)
			if bitNotErr != nil {
				return bitNotErr
			}
			err = vm.push(result)
			if err != nil {
				return err
			}

		case compiler.OpMinus: // Negation
			operand := vm.pop()
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
				// If condition is falsy, we jump. The falsy value remains on stack as the result.
				currentFrame.ip += offset
			} else {
				// If condition is truthy, we don't jump. The next instruction (OpPop) will remove it.
				// No action needed here.
			}

		case compiler.OpJump:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			currentFrame.ip += offset

		case compiler.OpJumpTruthy:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			condition := vm.StackTop() // Peek the condition
			if isTruthy(condition) {
				// If condition is truthy, we jump. The truthy value remains on stack as the result.
				currentFrame.ip += offset
			} else {
				// If condition is falsy, we don't jump. The next instruction (OpPop) will remove it.
				// No action needed here.
			}

		case compiler.OpSetGlobal:
			globalIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			value := vm.pop()
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
			value := vm.pop()
			if currentFrame.basePointer+localIndex >= len(vm.stack) { // Check against stack size
				return types.NewError("local variable index out of bounds: %d (basePointer %d, stack size %d)", localIndex, currentFrame.basePointer, len(vm.stack))
			}
			vm.stack[currentFrame.basePointer+localIndex] = value

		case compiler.OpGetLocal:
			localIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			if currentFrame.basePointer+localIndex >= len(vm.stack) { // Check against stack size
				return types.NewError("local variable index out of bounds: %d (basePointer %d, stack size %d)", localIndex, currentFrame.basePointer, len(vm.stack))
			}
			err = vm.push(vm.stack[currentFrame.basePointer+localIndex])
			if err != nil {
				return err
			}

		case compiler.OpSetFree: // New opcode for setting free variables
			freeIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			value := vm.pop()
			if int(freeIndex) >= len(currentFrame.closure.Free) {
				return types.NewError("free variable index out of bounds for set: %d (max %d)", freeIndex, len(currentFrame.closure.Free)-1)
			}
			currentFrame.closure.Free[freeIndex] = value

		case compiler.OpGetFree:
			freeIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			if int(freeIndex) >= len(currentFrame.closure.Free) {
				return types.NewError("free variable index out of bounds: %d (max %d)", freeIndex, len(currentFrame.closure.Free)-1)
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
			for i := freeCount - 1; i >= 0; i-- {
				frees[i] = vm.pop()
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
			for i := numElements - 1; i >= 0; i-- {
				elements[i] = vm.pop()
			}
			err = vm.push(types.NewList(elements...))
			if err != nil {
				return err
			}

		case compiler.OpTable:
			numPairs, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead

			pairs := make([]types.TablePair, numPairs)
			for i := numPairs - 1; i >= 0; i-- {
				value := vm.pop()
				keyVal := vm.pop()
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
			index := vm.pop()
			aggregate := vm.pop()
			result, getErr := aggregate.GetIndex(index)
			if getErr != nil {
				return types.NewError("runtime error: %s", getErr.Error())
			}
			err = vm.push(result)
			if err != nil {
				return err
			}

		case compiler.OpSetIndex:
			value := vm.pop()
			index := vm.pop()
			aggregate := vm.pop()
			setErr := aggregate.SetIndex(index, value)
			if setErr != nil {
				return types.NewError("runtime error: %s", setErr.Error())
			}
			err = vm.push(value)
			if err != nil {
				return err
			}

		case compiler.OpPrint:
			numExprs, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			args := make([]string, numExprs)
			// Pop elements from the stack (they come off right-to-left)
			for i := 0; i < int(numExprs); i++ {
				// The first item popped (rightmost in print statement) goes to the last slot in args.
				// The last item popped (leftmost in print statement) goes to the first slot in args.
				args[int(numExprs)-1-i] = vm.pop().Inspect()
			}
			// Join the arguments with a space and print the single line.
			fmt.Fprintln(vm.outputWriter, strings.Join(args, " "))

		case compiler.OpGetIter:
			iterable := vm.pop()
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

			iteratorVal := vm.pop() // Pop the iterator to work with it
			iterator, ok := iteratorVal.(types.Iterator)
			if !ok {
				return types.NewError("object is not an iterator: %s", iteratorVal.Type())
			}

			nextVal, hasNext, iterErr := iterator.Next()
			if iterErr != nil {
				return types.NewError("runtime error during iteration: %s", iterErr.Error())
			}

			if !hasNext {
				// If no more elements, push nil and jump to the 'after loop' instruction
				if err := vm.push(&types.Nil{}); err != nil {
					return err
				}
				currentFrame.ip += offset // Jump to exit if no more elements
			} else {
				// If there are more elements, push the next value
				if err := vm.push(nextVal); err != nil {
					return err
				}
				// And push the iterator back for the next iteration
				if err := vm.push(iterator); err != nil {
					return err
				}
			}

		case compiler.OpCall:
			numArgs, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead

			calleePos := vm.sp - int(numArgs) - 1
			callee := vm.stack[calleePos]

			closure, ok := callee.(*types.Closure)
			if !ok {
				return types.NewError("call target is not a function or closure: %s", callee.Type())
			}

			if int(numArgs) != closure.Fn.NumParameters {
				return types.NewError("wrong number of arguments: expected %d, got %d",
					closure.Fn.NumParameters, numArgs)
			}

			newFrame := NewFrame(closure, calleePos)

			err = vm.pushFrame(newFrame)
			if err != nil {
				return err
			}

			// Adjust stack pointer for new frame's locals (arguments are already on stack)
			vm.sp = newFrame.basePointer + closure.Fn.NumLocals

		case compiler.OpReturnValue:
			returnValue := vm.pop() // Pop the return value

			poppedFrame := vm.popFrame() // Pop the current frame

			// Clean up stack for arguments and locals of the popped frame
			vm.sp = poppedFrame.basePointer

			err = vm.push(returnValue) // Push the return value onto the previous frame's stack
			if err != nil {
				return err
			}

		case compiler.OpReturn:
			poppedFrame := vm.popFrame() // Pop the current frame

			// Clean up stack for arguments and locals of the popped frame
			vm.sp = poppedFrame.basePointer

			err = vm.push(&types.Nil{}) // Push nil as implicit return value
			if err != nil {
				return err
			}

		case compiler.OpImport: // Handle import statement
			pathVal := vm.pop()
			pathStr, ok := pathVal.(*types.String)
			if !ok {
				return types.NewError("import path must be a string, got %s", pathVal.Type())
			}
			// TODO: Implement actual module loading/compilation/execution here.
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
	right := vm.pop()
	left := vm.pop()

	// String concatenation for OpAdd
	if op == compiler.OpAdd && left.Type() == types.STRING_OBJ && right.Type() == types.STRING_OBJ {
		leftStr := left.(*types.String).Value
		rightStr := right.(*types.String).Value
		return vm.push(types.NewString(leftStr + rightStr))
	}

	// Type checking for numeric operations
	if (left.Type() != types.INTEGER_OBJ && left.Type() != types.FLOAT_OBJ) ||
		(right.Type() != types.INTEGER_OBJ && right.Type() != types.FLOAT_OBJ) {
		return types.NewError("type mismatch for %s: %s %s %s", op.String(), left.Type(), op.String(), right.Type())
	}

	var result types.Value

	switch op {
	case compiler.OpAdd:
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
	case compiler.OpDiv: // Standard division (float result if any operand is float)
		if (right.Type() == types.INTEGER_OBJ && right.(*types.Integer).Value == 0) ||
			(right.Type() == types.FLOAT_OBJ && right.(*types.Float).Value == 0.0) {
			return types.NewError("division by zero")
		}
		lVal := toFloat64(left)
		rVal := toFloat64(right)
		result = types.NewFloat(lVal / rVal)
	case compiler.OpMod:
		if left.Type() != types.INTEGER_OBJ || right.Type() != types.INTEGER_OBJ {
			return types.NewError("type mismatch for modulo: expected Integer %% Integer, got %s %% %s", left.Type(), right.Type())
		}
		if right.(*types.Integer).Value == 0 {
			return types.NewError("modulo by zero")
		}
		result = types.NewInteger(left.(*types.Integer).Value % right.(*types.Integer).Value)
	case compiler.OpPow: // Power operation
		lVal := toFloat64(left)
		rVal := toFloat64(right)
		result = types.NewFloat(math.Pow(lVal, rVal))
	case compiler.OpIDiv: // Integer division
		if left.Type() != types.INTEGER_OBJ || right.Type() != types.INTEGER_OBJ {
			return types.NewError("type mismatch for integer division: expected Integer // Integer, got %s // %s", left.Type(), right.Type())
		}
		if right.(*types.Integer).Value == 0 {
			return types.NewError("integer division by zero")
		}
		result = types.NewInteger(left.(*types.Integer).Value / right.(*types.Integer).Value)
	default:
		return types.NewError("unsupported binary operation: %s", op)
	}

	return vm.push(result)
}

// executeBitwiseOperation handles binary bitwise operations.
func (vm *VM) executeBitwiseOperation(op compiler.Opcode) error {
	right := vm.pop()
	left := vm.pop()

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
func toFloat64(val types.Value) float64 {
	if i, ok := val.(*types.Integer); ok {
		return float64(i.Value)
	}
	if f, ok := val.(*types.Float); ok {
		return f.Value
	}
	return 0.0 // Should ideally be an error or panic in a robust system
}

// executeComparisonOperation handles comparison operations.
func (vm *VM) executeComparisonOperation(op compiler.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	var result bool
	var err error

	switch op {
	case compiler.OpEqual:
		result = left.Equals(right)
	case compiler.OpNotEqual:
		result = !left.Equals(right)
	case compiler.OpGreaterThan:
		var cmp int
		cmp, err = left.Compare(right)
		if err != nil {
			return types.NewError("runtime error: %s", err.Error())
		}
		result = cmp > 0
	case compiler.OpLessThan:
		var cmp int
		cmp, err = left.Compare(right)
		if err != nil {
			return types.NewError("runtime error: %s", err.Error())
		}
		result = cmp < 0
	case compiler.OpGreaterEqual:
		var cmp int
		cmp, err = left.Compare(right)
		if err != nil {
			return types.NewError("runtime error: %s", err.Error())
		}
		result = cmp >= 0
	case compiler.OpLessEqual:
		var cmp int
		cmp, err = left.Compare(right)
		if err != nil {
			return types.NewError("runtime error: %s", err.Error())
		}
		result = cmp <= 0
	default:
		return types.NewError("unsupported comparison operation: %s", op)
	}
	return vm.push(types.NewBoolean(result))
}
