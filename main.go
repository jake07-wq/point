package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	digits := []rune{'0', '0' + 1, '0' + 2, '0' + 3, '0' + 4, '0' + 5, '0' + 6, '0' + 7, '0' + 8, '0' + 9}

	x1 := digits[points.x/10]
	x2 := digits[points.x%10]
	y1 := digits[points.y/10]
	y2 := digits[points.y%10]

	output := []rune{
		'x', ' ', '=', ' ', x1, x2, ',', ' ',
		'y', ' ', '=', ' ', y1, y2, '\n',
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
