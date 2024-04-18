package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 6}

		got := Sum(nums)
		want := 16

		assertMsg(t, got, want)
	})

	t.Run("Collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		assertMsg(t, got, want)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got -> %d, wanted -> %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got -> %d, wanted -> %d", got, want)
		}
	}

	t.Run("Sum of slices tails", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3, 5}, []int{0, 9, 1, 2, 3})
		want := []int{5, 3}

		checkSums(t, got, want)
	})

	t.Run("Sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func assertMsg(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Got -> %d, wanted -> %d", got, want)
	}

}
