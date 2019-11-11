/*
@date : 2019/09/09
@author : YaPi
@desc :
*/
package set

import (
	"dtSt/common"
	"dtSt/tree"
)

type TreeSet struct {
	b		*tree.Bst
}

func NewTreeSet() *TreeSet {
	return &TreeSet{b : tree.NewBst()}
}

func (t *TreeSet) Add(d common.E) {
	t.b.Add(d)
}

func (t *TreeSet) Remove(d common.E) {
	t.b.RemoveElement(d)
}

func (t *TreeSet) Contains(d common.E) bool {
	return t.b.Contains(d)
}

func (t *TreeSet) Size() int {
	return t.b.Size()
}

func (t *TreeSet) IsEmpty() bool {
	return t.b.IsEmpty()
}

