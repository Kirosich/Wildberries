package main

import "fmt"

type Point struct {
	x int
	y int
}

func CreatePoint(x int, y int) *Point {
	point := &Point{
		x: x,
		y: y,
	}
	return point
}

func (p *Point) FindDistance() int {
	if p.x > p.y {
		return p.x - p.y
	}
	return p.y - p.x
}

func main() {
	Point1 := CreatePoint(5, 1)
	Point2 := CreatePoint(3, 4)
	fmt.Println(Point1.FindDistance())
	fmt.Println(Point2.FindDistance())
}
