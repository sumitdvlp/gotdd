package interf

import "testing"

func TestPerimeter(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("Got %g want %g", got, want)
		}
	}

	rectangle := Rectangle{4.0, 3.0}
	got := rectangle.Perimeter()
	want := 14.0
	assertCorrectMessage(t, got, want)

	got = rectangle.Area()
	want = 12
	assertCorrectMessage(t, got, want)

	cir := Circle{3}
	got = cir.Area()
	want = 28.274333882308138
	assertCorrectMessage(t, got, want)
}
