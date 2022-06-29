package interf

import "testing"

func TestInterfase(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("Got %g want %g", got, want)
		}
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, got, want)
	}

	rectangle := Rectangle{4.0, 3.0}
	got := rectangle.Perimeter()
	want := 14.0
	assertCorrectMessage(t, got, want)

	want = 12
	checkArea(t, rectangle, want)

	cir := Circle{3}
	want = 28.274333882308138
	checkArea(t, cir, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

func TestInterfaceTable(t *testing.T) {
	areaTest := []struct{
		name string
		shape Shape
		hasArea float64
	}
	{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _,tt := range areaTest{
		t.Run(tt.name, func(t *testing.T) {
			got:= tt.Shape.Area()
			want:= tt.hasArea
			if got!=want{
				t.Errorf("%#v got %g wanted %g",tt.name, got, want)
			}
		})
	}

}
