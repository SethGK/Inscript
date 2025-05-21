// Package vm implements the Inscript virtual machine.
package vm

import (
	"fmt"
	"io"
	"math"
	"os"

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
				frame := vm.popFrame()
				vm.sp = frame.basePointer - 1
				if err := vm.push(&types.Nil{}); err != nil {
					return err
				}
				continue
			}
			break
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

		case compiler.OpAdd, compiler.OpSub, compiler.OpMul, compiler.OpDiv, compiler.OpMod:
			err = vm.executeBinaryOperation(opcode)
			if err != nil {
				return err
			}

		// Placeholder for OpPow (assuming opcode 100)
		case compiler.Opcode(100):
			right := vm.pop()
			left := vm.pop()
			result, powErr := vm.executePower(left, right)
			if powErr != nil {
				return powErr
			}
			err = vm.push(result)
			if err != nil {
				return err
			}

		case compiler.OpBang:
			operand := vm.pop()
			err = vm.push(types.NewBoolean(!isTruthy(operand)))
			if err != nil {
				return err
			}

		case compiler.OpMinus:
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

		case compiler.OpEqual, compiler.OpNotEqual, compiler.OpGreaterThan, compiler.OpLessThan:
			err = vm.executeComparisonOperation(opcode)
			if err != nil {
				return err
			}

		// Placeholder for OpGreaterEqual (assuming opcode 101)
		case compiler.Opcode(101):
			right := vm.pop()
			left := vm.pop()
			cmp, cmpErr := left.Compare(right)
			if cmpErr != nil {
				return types.NewError("runtime error: %s", cmpErr.Error())
			}
			err = vm.push(types.NewBoolean(cmp >= 0))
			if err != nil {
				return err
			}
		// Placeholder for OpLessEqual (assuming opcode 102)
		case compiler.Opcode(102):
			right := vm.pop()
			left := vm.pop()
			cmp, cmpErr := left.Compare(right)
			if cmpErr != nil {
				return types.NewError("runtime error: %s", cmpErr.Error())
			}
			err = vm.push(types.NewBoolean(cmp <= 0))
			if err != nil {
				return err
			}

		case compiler.OpJumpNotTruthy:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			condition := vm.pop()
			if !isTruthy(condition) {
				currentFrame.ip += offset
			}

		case compiler.OpJump:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			currentFrame.ip += offset

		case compiler.OpJumpTruthy:
			offset, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			currentFrame.ip += bytesRead
			condition := vm.pop()
			if isTruthy(condition) {
				currentFrame.ip += offset
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
			if currentFrame.basePointer+localIndex >= vm.sp {
				return types.NewError("local variable index out of bounds: %d (basePointer %d, sp %d)", localIndex, currentFrame.basePointer, vm.sp)
			}
			vm.stack[currentFrame.basePointer+localIndex] = value

		case compiler.OpGetLocal:
			localIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 1)
			currentFrame.ip += bytesRead
			if currentFrame.basePointer+localIndex >= vm.sp {
				return types.NewError("local variable index out of bounds: %d (basePointer %d, sp %d)", localIndex, currentFrame.basePointer, vm.sp)
			}
			err = vm.push(vm.stack[currentFrame.basePointer+localIndex])
			if err != nil {
				return err
			}

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

		// Placeholder for OpClosure (assuming opcode 103)
		case compiler.Opcode(103):
			constIndex, bytesRead := compiler.ReadOperand(instructions, ip+1, 2)
			freeCount, bytesRead2 := compiler.ReadOperand(instructions, ip+1+bytesRead, 1)
			currentFrame.ip += bytesRead + bytesRead2

			fnVal := vm.constants[constIndex]
			compiledFn, ok := fnVal.(*types.CompiledFunction)
			if !ok {
				return types.NewError("constant %d is not a CompiledFunction, got %s", constIndex, fnVal.Type())
			}

			frees := make([]types.Value, freeCount)
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

		// Placeholder for OpTable (assuming opcode 104)
		case compiler.Opcode(104):
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

		// Placeholder for OpIndex (assuming opcode 105)
		case compiler.Opcode(105):
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

		// Placeholder for OpSetIndex (assuming opcode 106)
		case compiler.Opcode(106):
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
			for i := int(numExprs) - 1; i >= 0; i-- {
				args[i] = vm.pop().Inspect()
			}
			fmt.Fprintln(vm.outputWriter, args)

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
			iteratorVal := vm.pop()
			iterator, ok := iteratorVal.(types.Iterator)
			if !ok {
				return types.NewError("object is not an iterator: %s", iteratorVal.Type())
			}

			nextVal, hasNext, iterErr := iterator.Next()
			if iterErr != nil {
				return types.NewError("runtime error during iteration: %s", iterErr.Error())
			}

			if !hasNext {
				currentFrame.ip += offset
				vm.push(&types.Nil{})
			} else {
				if err := vm.push(nextVal); err != nil {
					return err
				}
			}
			if err := vm.push(iterator); err != nil {
				return err
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

			vm.sp = newFrame.basePointer + closure.Fn.NumLocals

		case compiler.OpReturnValue:
			returnValue := vm.pop()

			poppedFrame := vm.popFrame()

			vm.sp = poppedFrame.basePointer

			err = vm.push(returnValue)
			if err != nil {
				return err
			}

		case compiler.OpReturn:
			poppedFrame := vm.popFrame()

			vm.sp = poppedFrame.basePointer

			err = vm.push(&types.Nil{})
			if err != nil {
				return err
			}

		default:
			return types.NewError("unknown opcode: %s", opcode)
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
	return true
}

// executeBinaryOperation handles arithmetic operations.
func (vm *VM) executeBinaryOperation(op compiler.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	if op == compiler.OpAdd && left.Type() == types.STRING_OBJ && right.Type() == types.STRING_OBJ {
		leftStr := left.(*types.String).Value
		rightStr := right.(*types.String).Value
		return vm.push(types.NewString(leftStr + rightStr))
	}

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
	case compiler.OpDiv:
		if (right.Type() == types.INTEGER_OBJ && right.(*types.Integer).Value == 0) ||
			(right.Type() == types.FLOAT_OBJ && right.(*types.Float).Value == 0.0) {
			return types.NewError("division by zero")
		}
		if left.Type() == types.INTEGER_OBJ && right.Type() == types.INTEGER_OBJ {
			result = types.NewInteger(left.(*types.Integer).Value / right.(*types.Integer).Value)
		} else {
			lVal := toFloat64(left)
			rVal := toFloat64(right)
			result = types.NewFloat(lVal / rVal)
		}
	case compiler.OpMod:
		if left.Type() != types.INTEGER_OBJ || right.Type() != types.INTEGER_OBJ {
			return types.NewError("type mismatch for modulo: expected Integer %% Integer, got %s %% %s", left.Type(), right.Type())
		}
		if right.(*types.Integer).Value == 0 {
			return types.NewError("modulo by zero")
		}
		result = types.NewInteger(left.(*types.Integer).Value % right.(*types.Integer).Value)
	default:
		return types.NewError("unsupported binary operation: %s", op)
	}

	return vm.push(result)
}

// executePower handles the power operation (base^exponent).
func (vm *VM) executePower(left, right types.Value) (types.Value, error) {
	if (left.Type() != types.INTEGER_OBJ && left.Type() != types.FLOAT_OBJ) ||
		(right.Type() != types.INTEGER_OBJ && right.Type() != types.FLOAT_OBJ) {
		return nil, types.NewError("type mismatch for power: %s ^ %s", left.Type(), right.Type())
	}

	lVal := toFloat64(left)
	rVal := toFloat64(right)

	return types.NewFloat(math.Pow(lVal, rVal)), nil
}

// toFloat64 safely converts an Integer or Float Value to float64.
func toFloat64(val types.Value) float64 {
	if i, ok := val.(*types.Integer); ok {
		return float64(i.Value)
	}
	if f, ok := val.(*types.Float); ok {
		return f.Value
	}
	return 0.0
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
	default:
		return types.NewError("unsupported comparison operation: %s", op)
	}
	return vm.push(types.NewBoolean(result))
}
