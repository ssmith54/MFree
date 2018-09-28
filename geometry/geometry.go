package geometry

import "math"

type Point struct {
	x float64
	y float64
	z float64
}

// Return corodinates of point
func (point *Point) GetPointCoordinates() (float64, float64, float64) {
	x := point.x
	y := point.y
	z := point.z
	return x, y, z
}

func NewPoint(x, y, z float64) Point {
	return Point{x, y, z}
}

// get distance between points
func GetDistance(a *Point, b *Point) float64 {
	distance := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
	return distance
}

func GetClosestPoint() {

}
