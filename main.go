/*
@date : 2019/09/02
@author : YaPi
@desc :
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "werdsfasjfaj"
	ss := strings.Split(s,"")
	for k,v := range ss{
		fmt.Println(k," ",v)
	}
}
