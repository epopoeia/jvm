package constants

import (
	"JVM/instructions/base"
	"JVM/rtda"
	"JVM/rtda/heap"
)

type LDC struct {
	base.Index8Instruction
}

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	//case *heap.ClassRef
	default:
		panic("todo: ldc!")
	}
}
