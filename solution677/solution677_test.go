package solution677

type MapSum struct {
	Ended bool
	Val   int
	Nodes map[byte]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	obj := MapSum{}
	obj.Nodes = map[byte]*MapSum{}
	return obj
}

func (this *MapSum) Insert(key string, val int) {
	if key == "" {
		return
	}
	if len(key) == 1 {
		if _, ok := this.Nodes[key[0]]; !ok {
			this.Nodes[key[0]] = &MapSum{
				Nodes: map[byte]*MapSum{},
			}
		}
		this.Nodes[key[0]].Val = val
		this.Nodes[key[0]].Ended = true
		return
	}
	if _, ok := this.Nodes[key[0]]; !ok {
		this.Nodes[key[0]] = &MapSum{
			Nodes: map[byte]*MapSum{},
		}
	}
	this.Nodes[key[0]].Insert(key[1:], val)
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for i := range prefix {
		if v, ok := cur.Nodes[prefix[i]]; ok {
			cur = v
		} else {
			return 0
		}
	}
	return cur.sum()
}

func (this *MapSum) sum() int {
	sum := 0
	if this.Ended {
		sum += this.Val
	}
	for _, v := range this.Nodes {
		sum += v.sum()
	}
	return sum
}
