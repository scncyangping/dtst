/*
@date : 2019/09/16
@author : YaPi
@desc :
*/
package Trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Add("hello")
	trie.Add("word")

	fmt.Println(trie.Contains("hello"))
	fmt.Println(trie.Contains("hello"))
	fmt.Println(trie.Contains("Hello"))

}
