package rtda

import "myjvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
