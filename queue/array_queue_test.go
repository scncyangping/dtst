/*
@date : 2019/09/06
@author : YaPi
@desc :
*/
package queue

import (
	"dtSt/array"
	"dtSt/common"
	"fmt"
	"math/rand"
	"testing"
)

var q arrayQueue

func init()  {
	a := array.NewSArray()
	for i:=0;i<5;i++{
		newInt := common.NewInt(rand.Intn(100))
		a.Add(newInt)
	}
	q := NewArrayQueue()
	q.array = a
}

func TestArrayQueue(t *testing.T) {
	qq := NewArrayQueue()
	qq.Enqueue(common.NewInt(1))
	qq.Enqueue(common.NewInt(2))
	qq.Enqueue(common.NewInt(3))
	qq.Enqueue(common.NewInt(4))
	fmt.Println(qq)
	qq.Dequeue()
	fmt.Println(qq)
	qq.Dequeue()
	fmt.Println(qq)
	fmt.Println(qq.Size())
	fmt.Println(qq.IsEmpty())
	fmt.Println(qq.GetFront())
}
