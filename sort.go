package mergesort

func merge(left, right []int) []int {
	var l, r, elem int
	l_size := len(left)
	r_size := len(right)

	merged := make([]int, 0, l_size+r_size)

	for l < l_size && r < r_size {
		if left[l] < right[r] {
			elem = left[l]
			l++
		} else {
			elem = right[r]
			r++
		}

		merged = append(merged, elem)
	}

	if l < l_size {
		for l < l_size {
			merged = append(merged, left[l])
			l++
		}
	}

	if r < r_size {
		for r < r_size {
			merged = append(merged, right[r])
			r++
		}
	}
	return merged
}

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var mid int = len(arr) / 2

	left := Sort(arr[:mid])
	right := Sort(arr[mid:])

	return merge(left, right)

}
