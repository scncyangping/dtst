/*
@date : 2019/09/16
@author : YaPi
@desc : 字典树 本质就是一个多叉树
*/
package Trie

import "strings"

type TrieNode struct {
	// 表示该节点是否是单词节点
	isWord		bool
	// 该节点的值
	value 		string
	// 该节点的子节点指针数组
	// nodes		[]*TrieNode
	// 使用map方便获取指定节点
	nodes 		map[string]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{isWord: false,nodes:make(map[string]*TrieNode)}
}

func NewTrieNodeByValue(value string,isWord bool) *TrieNode {
	return &TrieNode{isWord: isWord,nodes:make(map[string]*TrieNode),value:value}
}

type Trie struct {
	root 		*TrieNode
	size 		int
}

func NewTrie() *Trie {
	return &Trie{root:NewTrieNode()}
}

func (t *Trie)GetSize()int{
	return t.size
}

func (t *Trie)Add(word string)  {
	node := t.root
	for _,v := range word{
		n := node.nodes[string(v)]
		if n == nil {
			node.nodes[string(v)] = NewTrieNodeByValue(string(v),false)

		}
		node = node.nodes[string(v)]
	}
	// 该节点是插入单词对应的最后一个节点
	if ! node.isWord {
		node.isWord = true
		t.size ++
	}
}

// 查看是否包含某个单词
func (t *Trie)Contains(word string)bool  {
	node := t.root
	for _,v := range word{
		n := node.nodes[string(v)]
		if n == nil {
			return false
		}
		node = node.nodes[string(v)]
	}
	// 判断该节点是否是一个单词节点
	if node.isWord {
		return true
	}else {
		return false
	}
}

// 用.号匹配任意字符
func (t *Trie)Match(word string)bool  {
	return t.match(t.root,strings.Split(word,""),0)
}
// 查看是否包含某个单词
func (t *Trie)match(node *TrieNode,word []string,index int)bool  {
	if index == len(word){
		return node.isWord
	}
	s := word[index]

	if s != "."{
		if node.nodes[s] == nil {
			return false
		}
		return t.match(node.nodes[s],word,index + 1)
	}else {
		for _,v := range node.nodes{
			if t.match(v,word,index + 1){
				return true
			}
		}
		return false
	}
}

// 查询是否有以prefix前缀的单词
func (t *Trie)IsPrefix(prefix string)bool  {
	node := t.root
	for _,v := range prefix{
		n := node.nodes[string(v)]
		if n == nil {
			return false
		}
		node = node.nodes[string(v)]
	}
	return true
}
