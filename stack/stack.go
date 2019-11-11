/*
@date : 2019/09/02
@author : YaPi
@desc :
*/
package stack

import "dtSt/common"

type S interface {
	Push(s common.E)
	PoP()common.E
	Peek()common.E
	Size()int
	IsEmpty()bool
}

type Stack struct {
	array []common.E
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) String() string {
	s := ""
	for _,v := range stack.array {
		s += "-->"+ v.String()
	}
	s += " Top"
	return s
}


func (stack *Stack) Push(s common.E) {
	stack.array = append(stack.array,s)
}

func (stack *Stack) PoP() common.E {
	if len(stack.array) == 0 {
		return nil
	}
	obj := stack.array[len(stack.array) - 1]
	stack.array = stack.array[:len(stack.array) - 1]
	return obj
}

func (stack *Stack) Peek() common.E {
	if len(stack.array) == 0 {
		return nil
	}
	return stack.array[len(stack.array) - 1]
}

func (stack *Stack) Size() int {
	return len(stack.array)
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.array) == 0
}
