package rtda

import "JVM/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
