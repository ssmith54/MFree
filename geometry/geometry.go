package geometry

// contains geometric primatives

import (
	"errors"
	"math"
)

// represent a polygon using this struct
type Polygon struct {
	vertex_list  []Point
	num_vertices int
	area         float64
}

// a 3D polygon, internal representation is the same as
// polygon, however different operations will be defined on polytopes
type Polytope struct {
	vertex_list   []Point
	num_verticies int
}

// representation of a geometric point
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
func CreatePolygon(x_points, y_points []float64) (*Polygon, error) {

	if len(x_points) != len(y_points) {
		return nil, errors.New("Cannot create polygon, x and y were different sizes")
	}
	// initialise polygon
	poly := new(Polygon)

	for i := 0; i < len(x_points); i++ {
		poly.vertex_list = append(poly.vertex_list, NewPoint(x_points[i], y_points[i], 0))
		poly.num_vertices++
	}

	poly.GetArea()

	return poly, nil
	//	err := errors.New("x and y were not the same size")
	// raise an error if the size of x and y is difrent

}

func (polygon *Polygon) GetArea() {

}
