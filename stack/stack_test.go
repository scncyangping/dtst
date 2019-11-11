/*
@date : 2019/09/02
@author : YaPi
@desc :
*/
package stack

import (
	"dtSt/common"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	s.Push(common.NewInt(2))
	//for i :=0 ;i<5;i++{
	//	s.Push(common.NewInt(i))
	//}
	t.Log(s)
	t.Log("IsEmpty ",s.IsEmpty())
	t.Log("Size ",s.Size())
	t.Log("PoP ",s.PoP())
	t.Log("Peek ",s.Peek())
	t.Log("Size ",s.Size())
	t.Log("PoP ",s.PoP())
	t.Log("Size ",s.Size())
	t.Log(s)


}
//// 括号匹配
//func IsValid(s string) bool {
//	la :="({["
//	ra :=")}]"
//	stack := NewStack()
//	for i := 0 ;i<len(s); i++{
//		if strings.Contains(ra, string(s[i])){
//			if !stack.IsEmpty(){
//				if stack.Peek() == string(s[i]){
//					stack.PoP()
//				}
//			}
//		}else if strings.Contains(la, string(s[i])){
//			stack.Push(string(s[i]))
//		}
//
//	}
//	return stack.IsEmpty()
//}

