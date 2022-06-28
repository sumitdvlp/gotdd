package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function
		// call rather than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	assertCorrectMessageArr := func(t testing.TB, gotarr, wantarr []int) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function
		// call rather than inside our test helper.
		t.Helper()
		if !reflect.DeepEqual(gotarr, wantarr) {
			t.Errorf("Got %q want %q", gotarr, wantarr)
		}
	}

	t.Run("Sum multiple arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorrectMessageArr(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want)
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
