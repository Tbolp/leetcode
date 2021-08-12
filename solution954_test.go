package leetcode

import "sort"

func canReorderDoubled(arr []int) bool {
	hm := map[int]int{}
	for i := range arr {
		hm[arr[i]]++
	}
	keys := make([]int, 0, len(hm))
	for k := range hm {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if hm[k] == 0 {
			continue
		}
		next := 0
		if k < 0 {
			if k%2 != 0 {
				return false
			}
			next = k / 2
		} else {
			next = k * 2
		}
		if nv, ok := hm[next]; ok {
			if hm[k] > nv {
				return false
			}
			hm[next] -= hm[k]
			hm[k] = 0
		} else {
			return false
		}
	}
	return true
}
