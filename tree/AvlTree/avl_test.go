/*
@date : 2019/09/17
@author : YaPi
@desc :
*/
package AvlTree

import (
	"dtSt/common"
	"fmt"
	"testing"
)

func TestAVL_AvlAdd(t *testing.T) {
	bst := NewAvl()
	//bst.Add(common.NewInt(5))
	//bst.Add(common.NewInt(3))
	//bst.Add(common.NewInt(6))
	//bst.Add(common.NewInt(8))
	//bst.Add(common.NewInt(4))
	//bst.Add(common.NewInt(2))

	bst.Add(common.NewInt(12))
	bst.Add(common.NewInt(7))
	bst.Add(common.NewInt(16))
	bst.Add(common.NewInt(5))
	bst.Add(common.NewInt(9))
	bst.Add(common.NewInt(15))
	bst.Add(common.NewInt(19))
	bst.Add(common.NewInt(2))
	bst.Add(common.NewInt(6))
	bst.Add(common.NewInt(8))
	bst.Add(common.NewInt(10))
	bst.Add(common.NewInt(18))
	bst.Add(common.NewInt(20))
	bst.Add(common.NewInt(3))
	bst.Add(common.NewInt(21))
	fmt.Println("is balance ",bst.IsBalanced())
	bst.Add(common.NewInt(22))
	fmt.Println("is balance ",bst.IsBalanced())
	bst.RemoveElement(common.NewInt(21))
	fmt.Println("is balance ",bst.IsBalanced())
	bst.RemoveElement(common.NewInt(22))
	fmt.Println("is balance ",bst.IsBalanced())
	//
	//list := make([]common.E,0)
	//
	//bst.LDR(list)
	//
	//for _,v := range list{
	//	fmt.Println("--》",v)
	//}
	//
	//fmt.Println("is balance ",bst.IsBalanced())
}
