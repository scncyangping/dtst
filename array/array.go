/*
@date : 2019/08/30
@author : YaPi
@desc :
*/
package array

import (
	"dtSt/common"
	"strconv"
)
// 自定义数组结构
type SArray struct {
	size 		int
	data 		[]common.E
}

// 实现stringer接口
func (s *SArray) String() string {
	str := "数组长度 : "+ strconv.Itoa(s.size) + " 【 ";
	for i:=0;i<len(s.data);i++{
		str += s.data[i].String() + "  "
	}
	str += " 】"
	return str
}

func NewSArray() *SArray {
	return &SArray{}
}


// 自定义数组实现的方法
type SArrayInterface interface {
	// 添加
	Add(e common.E)
	// 删除指定位置的元素
	Remove(i int)common.E
	// 数组容量(不是数组中元素的总数)
	Capacity()int
	// 数组中元素的总数
	Size()int
	// 查找元素common.E在数组中的位置
	Find(e common.E)int
	// 数组最后一个位置添加元素
	AddLast(e common.E)
	// 数组头添加元素
	AddFirst(e common.E)
	// 获取指定位置的元素
	Get(i int)common.E
	// 设置指定位置的元素
	Set(i int,e common.E)
	// 删除第一个元素
	RemoveFirst()common.E
	// 删除最后一个元素
	RemoveLast()common.E
	// 删除找到的第一个元素
	RemoveElement(e common.E)
}


func (s *SArray) Add(e common.E) {
	s.data = append(s.data, e)
	s.size ++
}

func (s *SArray) Remove(i int) common.E {
	if i < 0 || i >= s.size {
		return nil
	}
	e := s.data[i]
	s.data = append(s.data[:i], s.data[i+1:]...)
	s.size--
	return e
}

func (s *SArray) Capacity()int {
	return cap(s.data)
}

func (s *SArray) Size()int {
	return s.size
}

func (s *SArray) Find(e common.E) int {
	index := -1
	for i,d := range  s.data{
		if d.CompareTo(e) == 0{
			index = i
			break
		}
	}
	return index
}

func (s *SArray) AddLast(e common.E) {
	s.data = append(s.data, e)
	s.size ++
}

func (s *SArray) AddFirst(e common.E) {
	lNode := s.data[s.size-1]
	for i:=s.size-2;i>=0;i--{
		s.data[i+1] = s.data[i]
	}
	s.data[0] = e
	s.size ++
	s.data = append(s.data, lNode)
}

func (s *SArray) Get(i int) common.E {
	if i < 0 || i >= s.size {
		return nil
	}
	return s.data[i]
}

func (s *SArray) Set(i int, e common.E) {
	if i < 0 || i >= s.size {
		return
	}
	s.data[i] = e
}

func (s *SArray) RemoveFirst()common.E {
	return s.Remove(0)
}

func (s *SArray) RemoveLast()common.E {
	return s.Remove(s.size-1)
}

func (s *SArray) RemoveElement(e common.E) {
	index := s.Find(e)
	if index != -1{
		s.Remove(index)
	}
}
