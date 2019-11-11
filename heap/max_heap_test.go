/*
@date : 2019/09/10
@author : YaPi
@desc :
*/
package heap

import (
	"dtSt/common"
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap()

	for i:=0;i<10;i++{
		t := rand.Intn(1000)
		fmt.Println(t)
		maxHeap.Add(common.NewInt(t))
	}
	fmt.Println()
	test := make([]common.E,0)
	for i:=0;i<10;i++{
		test = append(test, maxHeap.ExtractMax())
	}

	for _,v := range test{
		fmt.Println(v)
	}
	// output
	//		81
	//		887
	//		847
	//		59
	//		81
	//		318
	//		425
	//		540
	//		456
	//		300
	//
	//		887
	//		847
	//		540
	//		456
	//		425
	//		318
	//		300
	//		81
	//		81
	//		59

}
