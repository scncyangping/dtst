/*
@date : 2019/09/09
@author : YaPi
@desc :
*/
package heap

import "dtSt/common"

type MaxHeap struct {
	data		[]common.E
	size 		int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data : make([]common.E,0)}
}

// 数组从0开始
// 父亲节点对应数组下标
func parent(n int)int  {
	if n == 0{
		return n
	}
	return (n-1)/2
}
// 左孩子节点对应数组下标
func leftChild(n int)int{
	return 2 * n +1
}
// 右孩子节点对应数组下标
func rightChild(n int)int{
	return 2 * n +2
}

func (m *MaxHeap)get(i int) common.E {
	if i < 0 || i > len(m.data){
		return nil
	}
	return m.data[i]
}

func (m *MaxHeap)swap(i ,j int){
	if i < 0 || j < 0 || i >= len(m.data) || j >= len(m.data){
		return
	}
	m.data[i],m.data[j] = m.data[j],m.data[i]
}

func (m *MaxHeap)Add(e common.E)  {
	// 添加元素到末尾
	m.data = append(m.data, e)
	// 对数组进行shiftUp操作
	m.shiftUp(len(m.data) - 1)
	m.size ++
}

// 向上重新组装二叉堆
func (m *MaxHeap)shiftUp(i int)  {
	for i >0 && m.get(i).CompareTo(m.get(parent(i))) == 1 {
		m.swap(i,parent(i))
		i = parent(i)
	}
}

func (m *MaxHeap)findMax()common.E  {
	if len(m.data) == 0{
		return nil
	}
	return m.data[0]
}

func (m * MaxHeap)ExtractMax()common.E{
	e := m.findMax()
	if e == nil {
		return e
	}
	// 将最后一个元素的值设置到最大值的位置
	m.swap(0,len(m.data)-1)
	// 将最大值删除
	m.data = m.data[:len(m.data)-1]
	// 从根节点执行down操作
	m.shiftDown(0)
	m.size --
	return e
}

func (m *MaxHeap)shiftDown(k int)  {
	for leftChild(k) < len(m.data){
		// 设置j的值为左孩子对应的数组下标值
		j := leftChild(k)
		// j + 1 是右孩子的数组中的下标
		// 如果右孩子对应的值大于左孩子对应的值，设置j的值为右孩子的下标
		if j + 1 < len(m.data) && m.get(j + 1).CompareTo(m.get(j)) == 1{
				j ++
		}
		// 判断当前节点的值是不是比j对应的孩子节点大，如果大于或等于，就跳过，否则就替换
		if m.get(k).CompareTo(m.get(j)) >= 0{
			break
		}

		m.swap(k,j)
		k = j
	}
}