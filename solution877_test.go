package leetcode

func stoneGame(piles []int) bool {
	return true
	// return stoneGame_alex(piles, 0, 0)
}

func stoneGame_alex(piles []int, alex int, lee int) bool {
	if len(piles) == 0 {
		return alex > lee
	}
	if stoneGame_lee(piles[1:], alex+piles[0], lee) {
		return true
	}
	if stoneGame_lee(piles[:len(piles)-1], alex+piles[len(piles)-1], lee) {
		return true
	}
	return false
}

func stoneGame_lee(piles []int, alex int, lee int) bool {
	if len(piles) == 0 {
		return alex > lee
	}
	if stoneGame_alex(piles[1:], alex, lee+piles[0]) {
		return true
	}
	if stoneGame_alex(piles[:len(piles)-1], alex, lee+piles[len(piles)-1]) {
		return true
	}
	return false
}
