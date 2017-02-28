package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
	Z int
}

func createPoint(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func calculateDistance(a *Point, b *Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X - b.X), 2) + math.Pow(float64(a.Y - b.Y), 2))
}

func main() {
	pointA := createPoint(0, 6)
	var pointB *Point
	fmt.Println(calculateDistance(pointA, pointB))
	fmt.Println("done!")
}
