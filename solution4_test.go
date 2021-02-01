package leetcode

import "testing"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1+l2 == 1 {
		if l1 == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
	stop := (l1 + l2 - 1) / 2
	count := 0
	for count != stop {
		if len(nums2) == 0 {
			nums1 = nums1[stop-count:]
			break
		}
		if len(nums1) == 0 {
			nums2 = nums2[stop-count:]
			break
		}
		l1pos := (stop - count - 1) / 2
		l2pos := (stop - count - 1) / 2
		if l1pos >= len(nums1) {
			l1pos = len(nums1) - 1
		}
		if l2pos >= len(nums2) {
			l2pos = len(nums2) - 1
		}
		if nums1[l1pos] > nums2[l2pos] {
			nums2 = nums2[l2pos+1:]
			count += l2pos + 1

		} else {
			nums1 = nums1[l1pos+1:]
			count += l1pos + 1

		}
	}
	first := 0.0
	second := 0.0
	if len(nums1) == 0 {
		first = float64(nums2[0])
		second = float64(nums2[1])
	} else if len(nums2) == 0 {
		first = float64(nums1[0])
		second = float64(nums1[1])
	} else {
		if nums1[0] < nums2[0] {
			first = float64(nums1[0])
			second = float64(nums2[0])
			if len(nums1) > 1 && float64(nums1[1]) < second {
				second = float64(nums1[1])
			}

		} else {
			first = float64(nums2[0])
			second = float64(nums1[0])
			if len(nums2) > 1 && float64(nums2[1]) < second {
				second = float64(nums2[1])
			}
		}
	}
	if (l1+l2)%2 != 0 {
		return first
	} else {
		return (first + second) / 2
	}
}

func Test_findMedianSortedArrays_Usage1(t *testing.T) {
	findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6, 7})
}

func Test_findMedianSortedArrays_Usage2(t *testing.T) {
	findMedianSortedArrays([]int{1, 3}, []int{2})
}
