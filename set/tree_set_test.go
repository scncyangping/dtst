/*
@date : 2019/09/09
@author : YaPi
@desc :
*/
package set

import (
	"dtSt/common"
	"fmt"
	"testing"
)

func TestTreeSet_Add(t *testing.T) {
	treeSet := NewTreeSet()

	newInt := common.NewInt(1)
	treeSet.Add(newInt)

	fmt.Println(treeSet.IsEmpty())

	fmt.Println(treeSet.Size())

	fmt.Println(treeSet.Contains(common.NewInt(2)))
	treeSet.Add(common.NewInt(2))
	fmt.Println(treeSet.Contains(common.NewInt(2)))
	treeSet.Remove(common.NewInt(2))
	fmt.Println(treeSet.Contains(common.NewInt(2)))
	treeSet.Remove(common.NewInt(1))
	fmt.Println(treeSet.Size())
	treeSet.Remove(common.NewInt(2))
	fmt.Println(treeSet.Size())

}
