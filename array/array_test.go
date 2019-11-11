/*
@date : 2019/08/30
@author : YaPi
@desc :
*/
package array

import (
	"dtSt/common"
	"fmt"
	"math/rand"
	"testing"
)


func TestSArray(t *testing.T) {
	a := NewSArray()
	for i:=0;i<5;i++{
		newInt := common.NewInt(rand.Intn(100))
		a.Add(newInt)
	}
	fmt.Println(a)
	fmt.Println("添加元素1")
	a.Add(common.NewInt(1))
	fmt.Println(a)
	fmt.Println("删除第二个元素")
	a.Remove(2)
	fmt.Println(a)
	fmt.Println("查找元素为81的元素: ",a.Find(common.NewInt(81)))
	fmt.Println("获取第二个元素: ",a.Get(2))
	fmt.Println("数组头添加一个元素111")
	a.AddFirst(common.NewInt(111))
	fmt.Println(a)
	fmt.Println("数组尾添加一个元素123")
	a.AddLast(common.NewInt(123))
	fmt.Println(a)
	fmt.Println("删除第一个元素")
	a.RemoveFirst()
	fmt.Println(a)
	fmt.Println("删除最后一个元素")
	a.RemoveLast()
	fmt.Println(a)
	fmt.Println("删除指定为1的元素")
	a.RemoveElement(common.NewInt(1))
	fmt.Println(a)
}
