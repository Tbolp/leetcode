package leetcode

import (
	"math"
	"testing"
)

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := (C - A) * (D - B)
	area2 := (G - E) * (H - F)
	width := 0
	height := 0
	if E >= A && E <= C {
		width = int(math.Min(float64(C), float64(G))) - E
	} else if G >= A && G <= C {
		width = G - int(math.Max(float64(A), float64(E)))
	} else if C <= G && A >= E {
		width = C - A
	}
	if F >= B && F <= D {
		height = int(math.Min(float64(D), float64(H))) - F
	} else if H >= B && H <= D {
		height = H - int(math.Max(float64(F), float64(B)))
	} else if B >= F && D <= H {
		height = D - B
	}
	return area1 + area2 - width*height
}

func Test_computeArea_Usage1(t *testing.T) {
	if computeArea(-3, 0, 3, 4, 0, -1, 9, 2) != 45 {
		t.Fail()
	}
}

func Test_computeArea_Usage2(t *testing.T) {
	if computeArea(-2, -2, 2, 2, -2, -2, 2, 2) != 16 {
		t.Fail()
	}
}

func Test_computeArea_Usage3(t *testing.T) {
	if computeArea(-2, -2, 2, 2, -3, -3, 3, -1) != 24 {
		t.Fail()
	}
}

func Test_computeArea_Usage4(t *testing.T) {
	if computeArea(-2, -2, 2, 2, 1, -3, 3, 3) != 24 {
		t.Fail()
	}
}

func Test_computeArea_Usage5(t *testing.T) {
	if computeArea(-5, 0, 0, 3, -3, -3, 3, 3) != 42 {
		t.Fail()
	}
}
