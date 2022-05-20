package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (squareValue Square) End() Point {
	var pointValue Point
	pointValue.x = squareValue.start.x + int(squareValue.a)
	pointValue.y = squareValue.start.y - int(squareValue.a)
	return pointValue
}

func (squareValue Square) Area() uint {
	return squareValue.a * squareValue.a
}

func (squareValue Square) Perimeter() uint {
	return 4 * squareValue.a
}
