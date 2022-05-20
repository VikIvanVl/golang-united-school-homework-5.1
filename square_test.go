package golang_united_school_homework_5_1

import (
	"testing"
)

func TestEnd(t *testing.T) {
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

func TestArea(t *testing.T) {
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

func TestPerimeter(t *testing.T) {
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
