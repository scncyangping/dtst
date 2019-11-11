/*
@date : 2019/09/05
@author : YaPi
@desc :
*/
package BstTree

import (
	"dtSt/common"
	"fmt"
	"testing"
)

func TestBst_Add(t *testing.T) {
	bst := NewBst()
	//bst.Add(common.NewInt(5))
	//bst.Add(common.NewInt(3))
	//bst.Add(common.NewInt(6))
	//bst.Add(common.NewInt(8))
	//bst.Add(common.NewInt(4))
	//bst.Add(common.NewInt(2))

	// test2
	bst.Add(common.NewInt(42))
	bst.Add(common.NewInt(33))
	bst.Add(common.NewInt(50))
	bst.Add(common.NewInt(17))
	bst.Add(common.NewInt(37))
	bst.Add(common.NewInt(48))
	bst.Add(common.NewInt(88))
	bst.Add(common.NewInt(12))
	bst.Add(common.NewInt(18))
	bst.Add(common.NewInt(66))
	bst.Add(common.NewInt(6))
	//fmt.Println("max: ",bst.Max())
	//fmt.Println("min: ",bst.Mind())
	//st.RemoveMax()
	//bst.RemoveMind()
	//fmt.Println("max: ",bst.Max())
	//fmt.Println("min: ",bst.Mind())
	// bst.DLR()
	 bst.LDR()
	// bst.LRD()
	bst.RemoveElement(common.NewInt(16))
	for _, v := range bst.LevelOrder() {
		fmt.Print(v, " ")
	}

}
