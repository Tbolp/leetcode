package leetcode

type Trie struct {
	Val   byte
	Ended bool
	Nodes []Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	obj := Trie{}
	obj.Nodes = []Trie{}
	return obj
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.Ended = true
		return
	}
	for ind := range this.Nodes {
		if this.Nodes[ind].Val == word[0] {
			this.Nodes[ind].Insert(word[1:])
			return
		}
	}
	this.Nodes = append(this.Nodes, Constructor())
	this.Nodes[len(this.Nodes)-1].Val = word[0]
	this.Nodes[len(this.Nodes)-1].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		if this.Ended {
			return true
		}
		return false
	}
	for ind := range this.Nodes {
		if this.Nodes[ind].Val == word[0] {
			return this.Nodes[ind].Search(word[1:])
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	for ind := range this.Nodes {
		if this.Nodes[ind].Val == prefix[0] {
			return this.Nodes[ind].StartsWith(prefix[1:])
		}
	}
	return false
}
