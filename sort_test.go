package mergesort

import (
	"testing"
)

func sortAndCompare(iterable, expected []int, t *testing.T) {
	result := Sort(iterable)

	if len(result) != len(iterable) {
		t.Errorf("Length mismatch: result: %d, input: %d\n", len(iterable), len(result))
	}

	for i := 0; i < len(iterable); i++ {
		if expected[i] != result[i] {
			t.Error("Mismatch at index: %d: exp[i]: %d, result[i]: %d", i,
				expected[i], result[i])
		}
	}

}

func TestSort(t *testing.T) {
	input := []int{1, 4, 3, 2}
	exp := []int{1, 2, 3, 4}

	sortAndCompare(input, exp, t)
}

func TestEqualElements(t *testing.T) {
	input := []int{1, 3, 2, 3, 3, 2, 3, 3, 2, 3}
	exp := []int{1, 2, 2, 2, 3, 3, 3, 3, 3, 3}
	sortAndCompare(input, exp, t)
}

func TestEmptyInput(t *testing.T) {
	input := []int{}
	exp := []int{}
	sortAndCompare(input, exp, t)
}
