package leetcode

import "testing"

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	m := len(word1)
	n := len(word2)
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	cache[0][0] = 1
	status := true
	for i := 0; i < m; i++ {
		if i != 0 {
			cache[i][0] = cache[i-1][0] + 1
		}
		if word1[i] == word2[0] && status {
			cache[i][0]--
			status = false
		}
	}
	cache[0][0] = 1
	status = true
	for j := 0; j < n; j++ {
		if j != 0 {
			cache[0][j] = cache[0][j-1] + 1
		}
		if word1[0] == word2[j] && status {
			cache[0][j]--
			status = false
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			val := cache[i-1][j-1]
			if word1[i] != word2[j] {
				val++
			}
			if val > cache[i-1][j]+1 {
				val = cache[i-1][j] + 1
			}
			if val > cache[i][j-1]+1 {
				val = cache[i][j-1] + 1
			}
			cache[i][j] = val
		}
	}
	return cache[m-1][n-1]
}

func Test_minDistance_Usage1(t *testing.T) {
	if minDistance("a", "ab") != 1 {
		t.FailNow()
	}
}

func Test_minDistance_Usage2(t *testing.T) {
	if minDistance("ab", "a") != 1 {
		t.FailNow()
	}
}

func Test_minDistance_Usage3(t *testing.T) {
	if minDistance("a", "a") != 0 {
		t.FailNow()
	}
}

func Test_minDistance_Usage4(t *testing.T) {
	if minDistance("pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopically") != 27 {
		t.FailNow()
	}
}
