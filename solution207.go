package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	skipset := map[int]bool{}
	for _, edge := range prerequisites {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	for i := 0; i < numCourses; i++ {
		if skipset[i] == false {
			if insertSet(i, graph, map[int]bool{}, skipset) == false {
				return false
			}
		}
	}
	return true
}

func insertSet(node int, graph map[int][]int, set map[int]bool, skipset map[int]bool) bool {
	skipset[node] = true
	set[node] = true
	for _, val := range graph[node] {
		tmp := map[int]bool{}
		for k, v := range set {
			tmp[k] = v
		}
		if tmp[val] == true {
			return false
		}
		if insertSet(val, graph, tmp, skipset) == false {
			return false
		}
	}
	return true
}
