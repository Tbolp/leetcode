package solution380

import "math/rand"

type RandomizedSet struct {
	data  map[int]int
	array []int
}

func Constructor() RandomizedSet {
	obj := RandomizedSet{}
	obj.data = map[int]int{}
	obj.array = []int{}
	return obj
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; ok {
		return false
	}
	this.data[val] = len(this.array)
	this.array = append(this.array, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if ind, ok := this.data[val]; ok {
		last := len(this.array) - 1
		this.array[ind] = this.array[last]
		this.data[this.array[ind]] = ind
		this.array = this.array[:last]
		delete(this.data, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.array[rand.Intn(len(this.array))]
}
