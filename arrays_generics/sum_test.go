package arraysgenerics

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
	t.Run("SumGenerics - collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumGenerics(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{0, 9}, []int{1, 2})
	want := []int{9, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("TestSumAllTails", func(t *testing.T) {
		checkSums := func(t testing.TB, got, want []int) {
			t.Helper()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}

		t.Run("Sum some slice", func(t *testing.T) {
			got := SumAllTails([]int{0, 1, 2}, []int{1, 2})
			want := []int{3, 2}
			checkSums(t, got, want)

		})

		t.Run("Run empty slice safely", func(t *testing.T) {
			got := SumAllTails([]int{}, []int{1, 2})
			want := []int{0, 2}
			checkSums(t, got, want)
		})
	})
	t.Run("TestSumAllTails generics", func(t *testing.T) {
		checkSums := func(t testing.TB, got, want []int) {
			t.Helper()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}

		t.Run("Sum some slice", func(t *testing.T) {
			got := SumAllTailGenerics([]int{0, 1, 2}, []int{1, 2})
			want := []int{3, 2}
			checkSums(t, got, want)

		})

		t.Run("Run empty slice safely", func(t *testing.T) {
			got := SumAllTailGenerics([]int{}, []int{1, 2})
			want := []int{0, 2}
			checkSums(t, got, want)
		})
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(a, b string) string {
			return a + b
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func AssertEqual[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
