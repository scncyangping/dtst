/*
@date : 2019/09/06
@author : YaPi
@desc :
*/
package common

import (
	"strconv"
)

// 自定义可比较类型
type E interface {
	// 0 相等 1 大于 -1 小于
	CompareTo(e E)int
	// 返回打印内容
	String()string
}


// 集合接口需实现的方法 二叉树使用
type Set interface {
	Add(d E)
	Remove(d E)
	Contains(d E)bool
	Size()int
	IsEmpty()bool
}


type Segment interface {
	// 输入两个节点a,b，根据业务需求返回合适的节点，如：
	// 若线段树是用于存储一段区间的最大值，那么就应该返回阿a 和 b中的最大值
	// 若线段树是用于存储一段区间的最小值，那么就应该返回阿a 和 b中的最小值
	// 根据具体的结果返回
	Merge(b Segment)Segment
	String()string
}

type NewInt int
// stringer接口
func (ee NewInt) String() string {
	return strconv.Itoa(int(ee))
}
// 数组、链表、树等自定义接口的比较接口
func (ee NewInt) CompareTo(e E) int {

	if e == nil{
		return 0
	}

	eeInt := int(ee)

	eInt := int(e.(NewInt))

	if eeInt > eInt {
		return 1
	}else if eeInt == eInt{
		return 0
	}else {
		return -1
	}
	return 0
}
// 线段树接口
// 构建一个求最大值的线段树方法
func (ee NewInt) Merge(b Segment) Segment {
	if b == nil{
		return ee
	}
	bb := b.(NewInt)
	// 最大值
	if ee.CompareTo(bb) == 1 {
		// ret
		return ee
	}else if ee.CompareTo(bb) == 0{
		return ee
	}else {
		return b
	}

	//// 求和
	//eeInt := int(ee)
	//
	//eInt := int(bb)
	//
	//return NewInt(eeInt+eInt)

}


