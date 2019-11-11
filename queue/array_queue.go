/*
@date : 2019/09/06
@author : YaPi
@desc : 使用自定义数组实现队列
*/
package queue

import (
	"dtSt/array"
	"dtSt/common"
)

type arrayQueue struct {
	array 			*array.SArray
}

func (s *arrayQueue) String() string {
	return s.array.String()
}

func NewArrayQueue() *arrayQueue {
	return &arrayQueue{array: array.NewSArray()}
}

func (s *arrayQueue) Size() int {
	return s.array.Size()
}

func (s *arrayQueue) IsEmpty() bool {
	return s.array.Size() == 0
}

func (s *arrayQueue) Enqueue(q common.E) {
	s.array.AddLast(q)
}

func (s *arrayQueue) Dequeue() common.E {
	return s.array.RemoveFirst()
}

func (s *arrayQueue) GetFront()common.E {
	return s.array.Get(0)
}
