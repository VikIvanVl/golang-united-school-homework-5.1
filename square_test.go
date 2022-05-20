package golang_united_school_homework_5_1

import (
	"testing"
)

func Test_End(t *testing.T) {
	in := &Square{
		start: Point{0, 0},
		a:     2,
	}
	excepted := Point{2, -2}

	got := in.End()
	if excepted != got {
		t.Log("Invalid result for End")
		t.Fail()
	}
}

func Test_Area(t *testing.T) {
	in := &Square{
		start: Point{0, 0},
		a:     2,
	}
	excepted := uint(4)

	got := in.Area()
	if excepted != got {
		t.Log("Invalid result for Area")
		t.Fail()
	}
}

func Test_Perimeter(t *testing.T) {
	in := &Square{
		start: Point{0, 0},
		a:     2,
	}
	excepted := uint(8)

	got := in.Perimeter()
	if excepted != got {
		t.Log("Invalid result for Perimeter")
		t.Fail()
	}
}
