package arrays_slices

import (
	"reflect"
	"testing"
)

func checkSums(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("should result 15 given collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given numbers %v", got, want, numbers)
		}
	})

}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("should result [2, 9] given args", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("should result [74, 4, 100] given args", func(t *testing.T) {
		got := SumTails([]int{1, 2, 4, 10, 23, 35}, []int{0, 0, 0, 0, 2, 1, 1}, []int{20, 20, 20, 20, 20, 20})
		want := []int{74, 4, 100}

		checkSums(t, got, want)
	})

	t.Run("should result [74, 4, 100] given args", func(t *testing.T) {
		got := SumTails([]int{}, []int{0, 0, 1})
		want := []int{0, 1}

		checkSums(t, got, want)
	})
}
