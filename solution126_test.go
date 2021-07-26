package leetcode

import (
	"fmt"
	"testing"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	flag := false
	startind := 0
	for i, v := range wordList {
		if v == beginWord {
			flag = true
			startind = i
		}
	}
	if !flag {
		wordList = append(wordList, beginWord)
		startind = len(wordList) - 1
	}
	endind := 0
	flag = false
	for i, v := range wordList {
		if v == endWord {
			flag = true
			endind = i
		}
	}
	if !flag {
		return [][]string{}
	}
	type void struct{}
	null := void{}
	aset := map[int][]int{}
	bset := map[int]void{}
	aset[startind] = []int{}
	for i := range wordList {
		if i != startind {
			bset[i] = null
		}
	}
	checklist := map[int]void{}
	checklist[startind] = null
	for {
		tmplist := map[int]void{}
		for v := range checklist {
			for k := range bset {
				if findLadders_isNear(wordList[v], wordList[k]) {
					tmplist[k] = null
					aset[v] = append(aset[v], k)
				}
			}
		}
		checklist = tmplist
		for v := range checklist {
			delete(bset, v)
			aset[v] = []int{}
		}
		if len(checklist) == 0 {
			break
		}
	}
	if _, ok := aset[endind]; !ok {
		return [][]string{}
	}
	fmt.Println(aset)
	ctx := findLadders_{}
	ctx.Res = [][]string{}
	ctx.Path = aset
	ctx.Endind = endind
	ctx.List = wordList
	ctx.traversal(startind)
	return ctx.Res
}

func findLadders_isNear(a string, b string) bool {
	dis := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dis++
		}
	}
	return dis <= 1
}

type findLadders_ struct {
	Res    [][]string
	Path   map[int][]int
	List   []string
	Endind int
	tmp    []string
}

func (this *findLadders_) traversal(root int) {
	if root == this.Endind {
		res := make([]string, len(this.tmp))
		copy(res, this.tmp)
		res = append(res, this.List[root])
		this.Res = append(this.Res, res)
		return
	}
	for _, c := range this.Path[root] {
		this.tmp = append(this.tmp, this.List[root])
		this.traversal(c)
		this.tmp = this.tmp[:len(this.tmp)-1]
	}
}

func Test_findLadders_Usage1(t *testing.T) {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func Test_findLadders_Usage2(t *testing.T) {
	fmt.Println(findLadders("kiss", "tusk", []string{"miss", "dusk", "kiss", "musk", "tusk", "diss", "disk", "sang", "ties", "muss"}))
}

func Test_findLadders_Usage3(t *testing.T) {
	fmt.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
}
