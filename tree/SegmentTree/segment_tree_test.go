/*
@date : 2019/09/11
@author : YaPi
@desc :
*/
package SegmentTree

import (
	"dtSt/common"
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	arr := make([]common.Segment,0)

	arr = append(arr, common.NewInt(-2))
	arr = append(arr, common.NewInt(0))
	arr = append(arr, common.NewInt(3))
	arr = append(arr, common.NewInt(-5))
	arr = append(arr, common.NewInt(2))
	arr = append(arr, common.NewInt(-1))

	sTree := NewSegmentTree(arr)

	fmt.Print(sTree)

	fmt.Println(sTree.query(0,2))
	fmt.Println(sTree.query(1,2))
	fmt.Println(sTree.query(3,5))

	sTree.Set(5,common.NewInt(5))
	fmt.Print(sTree)

	fmt.Println(sTree.query(0,2))
	fmt.Println(sTree.query(1,2))
	fmt.Println(sTree.query(3,5))
}
