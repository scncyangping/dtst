/*
@date : 2019/09/11
@author : YaPi
@desc :
*/
package SegmentTree

import "dtSt/common"

type SegmentTree struct {
	data		[]common.Segment
	tree 		[]common.Segment
}

func (s *SegmentTree) String() string {
	ss := "[ "
	for _,v := range s.tree{
		if v != nil {
			ss += " "+ v.String()
		}
	}
	ss += " ]"
	return ss
}

func NewSegmentTree(d []common.Segment) *SegmentTree {
	// 需要 4n的元素才能存储完长度为len的数组所生成的线段树
	tre := &SegmentTree{
		data: d,
		tree: make([]common.Segment,len(d) * 4)}

	tre.buildSegmentTree(0,0,len(d) - 1)
	return tre
}

func (s *SegmentTree)IsEmpty()bool{
	return len(s.data) == 0
}

func (s *SegmentTree)GetSize()int{
	return len(s.data)
}

func (s *SegmentTree)Get(index int)common.Segment{
	if index < 0 || index >= len(s.data){
		return nil
	}
	return s.data[index]
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

// 构建线段树
// 在位置为i的地方构建区间为[l,r]的线段树
func (s *SegmentTree)buildSegmentTree(i,l,r int)  {
	// 说明已经不能再分了，就是一个单元素了
	if l == r{
		s.tree[i] = s.data[l]
		return
	}
	leftTressIndex := leftChild(i)
	rightTreeIndex := rightChild(i)
	// 不使用 (l + r) /2 避免越界
	mid := l + (r-l) / 2
	s.buildSegmentTree(leftTressIndex,l,mid)
	s.buildSegmentTree(rightTreeIndex,mid+1,r)

	s.tree[i] = s.tree[leftTressIndex].Merge(s.tree[rightTreeIndex])
}

func (s *SegmentTree)query(queryL,queryR int)common.Segment  {
	if queryL < 0 || queryR < 0 || queryL >=len(s.data) || queryR >= len(s.data){
		return nil
	}

	return s.queryInTree(0,0,len(s.data)-1,queryL,queryR)
}

// 在以treeIndex为根的，范围为l到r的线段树中查找区间为queryL,到queryR的数据
func (s *SegmentTree)queryInTree(treeIndex,l,r,queryL,queryR int)common.Segment  {
	if l == queryL && r == queryR{
		return s.tree[treeIndex]
	}
	mid := l + (r -l)/2
	leftTressIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	// 需要查找的和左子树没关系
	if queryL >= mid +1 {
		return s.queryInTree(rightTreeIndex,mid + 1,r,queryL,queryR)
	}else if queryR <= mid {
		// 右子树没关系 在左孩子中查找
		return s.queryInTree(leftTressIndex,l,mid,queryL,queryR)
	}else {
		// 意味着所关注的区间一部分在左孩子，一部分在右孩子

		leftResult := s.queryInTree(leftTressIndex,l,mid,queryL,mid)

		rightResult := s.queryInTree(rightTreeIndex,mid+1,r,mid+1,queryR)

		return leftResult.Merge(rightResult)
	}
}

func (s *SegmentTree)Set(index int,e common.Segment)  {

	if index <0 || index >= len(s.data){
		return
	}
	s.set(0,0, len(s.data)-1,index,e)
}


func (s *SegmentTree)set(treeIndex,l,r,index int,e common.Segment) {
	if l == r {
		s.tree[treeIndex] = e
		return
	}
	mid := l + (r -l)/2

	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)
	if index >= mid +1 {
		s.set(rightTreeIndex,mid + 1,r,index,e)
	}else {
		s.set(leftTreeIndex,l,mid,index,e)
	}
	s.tree[treeIndex] = s.tree[leftTreeIndex].Merge(s.tree[rightTreeIndex])
}
