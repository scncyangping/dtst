/*
@date : 2019/09/05
@author : YaPi
@desc : 二分搜索树
*/
package BstTree

import (
	"dtSt/common"
	"dtSt/queue"
	"fmt"
)


// 树节点
type BstNode struct {
	// 值
	e 			common.E
	left,right	*BstNode
}

func (n BstNode) CompareTo(e common.E) int {
	return n.e.CompareTo(e)
}

func (n BstNode) String() string {
	return n.e.String()
}

func NewNode(e common.E) *BstNode {
	return &BstNode{e: e}
}

// 二分搜索树
type Bst struct {
	node 	*BstNode
	size 	int
}

func NewBst() *Bst {
	return &Bst{}
}

// 二分搜索树前序遍历
func (b *Bst)DLR(){
	// 先遍历当前节点
	// 再遍历左子树
	// 再遍历右子树
	dlr(b.node)
}

func dlr(node *BstNode)  {
	if node == nil{
		return
	}
	fmt.Println(node.e)
	dlr(node.left)
	dlr(node.right)
}

// 二分搜索树中序遍历
func (b *Bst)LDR(){
	// 先遍历当前节点
	// 再遍历左子树
	// 再遍历右子树
	ldr(b.node)
}

func ldr(node *BstNode)  {
	if node == nil{
		return
	}
	ldr(node.left)
	fmt.Println(node.e)
	ldr(node.right)
}


// 二分搜索树后序遍历
func (b *Bst)LRD(){
	// 先遍历当前节点
	// 再遍历左子树
	// 再遍历右子树
	lrd(b.node)
}

func lrd(node *BstNode)  {
	if node == nil{
		return
	}
	lrd(node.left)
	lrd(node.right)
	fmt.Println(node.e)
}

func contains(node *BstNode, d common.E)bool{
	if node == nil {
		return false
	}
	r := node.e.CompareTo(d)
	if r == -1{
		return contains(node.right,d)
	}else if r == 1{
		// 当前节点的值比它大，那么需要在左子树中去查
		return contains(node.left,d)
	}else{
		return true
	}
}

func add(b *BstNode,d common.E)(*BstNode,bool){
	isAdd := false
	// 将nil当作一个为NULL的子节点,那么,判断当程序找到下一个节点是
	// NULL的时候，则代表需要新加一个节点
	if b == nil {
		return NewNode(d),true
	}
	// 若b不为空的时候，则继续向下判断
	// 说明当前节点大于需要新增的节点
	res := b.e.CompareTo(d)
	if res == 1{
		b.left,isAdd = add(b.left,d)
	}else if res == -1{
		b.right,isAdd = add(b.right,d)
	}else{
		// 相等 不做操作
	}
	return b,isAdd
}

func removeMind(node *BstNode)*BstNode{
	if node.left == nil {
		rNode := node.right
		node.right = nil
		return rNode
	}
	node.left = removeMind(node.left)
	return node
}

func removeMax(node *BstNode)*BstNode{
	// 若右孩子为空，则当前节点为最大节点
	if node.right == nil {
		// 找到最大节点，找到其左孩子，左孩子为删除掉当前节点后的最大节点
		lNode := node.left
		node.left = nil
		return lNode
	}
	node.right = removeMax(node.right)
	return node
}


// 找二分搜索树的最小节点
func mind(n *BstNode)*BstNode{
	if n.left == nil{
		return n
	}
	// 获取节点的最右节点
	//lNode := mind(n.left)
	//if lNode != nil{
	//	return lNode
	//}else {
	//	return n
	//}
	return mind(n.left)
}

// 找二分搜索树的最大节点
func max(n *BstNode)*BstNode  {
	// 若当前节点为空 返回nil
	if n.right == nil{
		return n
	}
	// 获取节点的最右节点
	return max(n.right)
}

// 删除任意节点
func removeElement(node *BstNode,e common.E)*BstNode{
	// 非空判断
	if node == nil{
		return nil
	}
	// 和当前节点的值进行比较
	switch node.CompareTo(e) {
	case 1:
		// 当前节点的值比需要删除的值大
		// 需要在当前节点的左孩子进行查找
		node.left = removeElement(node.left,e)
	case -1:
		// 当前节点的值比需要删除的值小
		// 需要在当前节点的右孩子进行查找
		node.right = removeElement(node.right,e)
	case 0:
		// 需要删除的值就是当前值
		// 若当前节点没有左孩子和右孩子
		if node.left == nil && node.right == nil{
			node = nil
		}else if node.right != nil && node.left == nil {
			// 若当前节点只有右孩子 没有左孩子
			// 需要找到右子树当中最小的元素 用最小的元素当作当前节点的值
			n := mind(node.right)
			// 删除最小的元素 并返回删除过后的树
			rNode := removeMind(node.right)
			n.right = rNode
			node = n
		}else if node.left !=nil && node.right == nil {
			// 若当前节点只有左孩子 没有右孩子
			// 找到左子树当中最大的元素 用最大的元素当作当前节点值
			n := max(node.left)
			lNode := removeMax(node.left)
			n.left = lNode
			node = n
		}else if node.left != nil && node.right != nil {
			// 当前节点既有左孩子 也有右孩子
			// 找到左子树最小的值 作为当前节点的值
			n := mind(node.right)
			rNode := removeMind(node.right)
			n.left = node.left
			n.right = rNode
			node = n
		}
	default:
	}
	return node
}

func (b *Bst)RemoveElement(e common.E)  {
	n := removeElement(b.node,e)
	if n!= nil{
		b.size --
	}
}


func (b *Bst)RemoveMind()  {
	b.node = removeMind(b.node)

}

func (b *Bst)RemoveMax()  {
	b.node = removeMax(b.node)
}


func (b *Bst)Max()*BstNode  {
	return max(b.node)
}

func (b *Bst)Mind()*BstNode  {
	return mind(b.node)
}

// 层序遍历(又叫广度优先遍历) 借助队列
func (b *Bst)LevelOrder()(l []BstNode){
	// 借助队列实现
	q := queue.NewArrayQueue()
	node := b.node
	q.Enqueue(b.node)
	for !q.IsEmpty() {
		node = q.Dequeue().(*BstNode)
		// 将节点添加到数组里面返回
		l= append(l, *node)
		if node.left != nil {
			q.Enqueue(node.left)
		}
		if node.right != nil{
			q.Enqueue(node.right)
		}
	}
	return
}


func (b *Bst) Add(d common.E) {
	rNode,ok := add(b.node,d)
	b.node = rNode
	if ok {
		b.size ++
	}
}

func (b *Bst) Size() int {
	return b.size
}

func (b *Bst) IsEmpty() bool {
	return b.size == 0
}

func (b *Bst) Remove(d common.E) {
	removeElement(b.node,d)
}

func (b *Bst) Contains(d common.E) bool {
	return contains(b.node,d)
}
// 版本1
//func (b *Bst) Add(d int) {
//	if b.node == nil {
//		b.node = newNode(d)
//		b.size++
//		return
//	}else {
//		add(b.node, d)
//	}
//}
//func add(b *Node,d int){
//	// 判断当前节点和此根节点是否相等
//	// 若相等则不做操作 退出
//	if b.IsEqual(d){
//		return
//	}
//	// 判断当前值和传入值的比较
//	switch b.Compare(d) {
//	case 0:
//		// 相等就不做处理
//		return
//	case 1:
//		// 这个值比当前值小
//		// 应该插入左子树
//		// 判断左子树是否存在
//		if b.left != nil {
//			add(b.left,d)
//		}else {
//			// 不存在左子树
//			// 新建一个 并将当前左子树节点加入
//			b.left = newNode(d)
//		}
//	case -1:
//		// 这个值比当前值大
//		// 应该插入右子树
//		if b.right != nil {
//			add(b.right,d)
//		}else {
//			// 不存在左子树
//			// 新建一个 并将当前左子树节点加入
//			b.right = newNode(d)
//		}
//	default:
//		return
//	}
//}






