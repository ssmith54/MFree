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

type Rectangle struct {
	vertex_list []Point // 4 points that define a rectangle
	area        float64 // area
}

type Cylinder struct {
	radius float64
	p1     Point // p2------p2+r
	p2     Point // p1------p1+r
}

type Circle struct {
	radius float64
	center Point
}

func subtract_points(a, b *Point) *Point {

	cx := a.x - b.x
	cy := a.y - b.y
	cz := a.z - b.z

	return &(Point{cx, cy, cz})

}

type Segment struct {
	p1 *Point
	p2 *Point
}

func NewSegment(p1, p2 *Point) *Segment {
	return &(Segment{p1, p2})

}

type Dir struct {
	e1 float64
	e2 float64
	e3 float64
}

func (segment *Segment) normal() *Dir {
	nx := -(segment.p2.y - segment.p1.y)
	ny := segment.p2.x - segment.p2.x
	len := math.Sqrt(math.Pow(nx, 2) + math.Pow(ny, 2))
	nx = nx / len
	ny = ny / len
	// return direction
	return &(Dir{nx, ny, 0})
}

func (segment *Segment) length() float64 {
	return math.Sqrt(math.Pow((segment.p2.x-segment.p1.x), 2) +
		math.Pow((segment.p2.y-segment.p1.y), 2))
}

// IN ROUTINES, FINDS IF A POINT IS INSIDE A PRIMIATIVE SHAPE
func (point *Point) InCircle(circle *Circle) bool {
	// shift coordinates
	shifted_point := subtract_points(point, &circle.center)
	shifted_radius := math.Sqrt(math.Pow(shifted_point.x, 2) + math.Pow(shifted_point.y, 2))
	if shifted_radius <= circle.radius {
		return true
	} else {
		return false
	}

}
func (point *Point) InCylinder() bool {

	return true
}

func (point *Point) InRectangle() bool {

	return true
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

}

func (polygon *Polygon) GetArea() {

	var X1, Y1, X2, Y2, area float64
	area = 0
	for i := 0; i < polygon.num_vertices; i++ {

		if i == polygon.num_vertices-1 {
			X1 = polygon.vertex_list[i].x
			Y1 = polygon.vertex_list[i].y
			X2 = polygon.vertex_list[0].x
			Y2 = polygon.vertex_list[0].y
		} else {
			X1 = polygon.vertex_list[i].x
			Y1 = polygon.vertex_list[i].y
			X2 = polygon.vertex_list[i+1].x
			Y2 = polygon.vertex_list[i+1].y
		}

		area = area + (X1*Y2-Y1*X2)/2

	}
	polygon.area = math.Abs(area)

}

func (polygon *Polygon) GetNumVertices() int {
	return polygon.num_vertices
}

func (polygon *Polygon) GetPolyCoordinates() ([]float64, []float64) {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for _, point := range polygon.vertex_list {
		xtemp, ytemp, _ := point.GetPointCoordinates()

		x = append(x, xtemp)
		y = append(y, ytemp)
	}

	return x, y
}

func (polygon *Polygon) IsClockwise() bool {

	x, y := polygon.GetPolyCoordinates()

	// sum over edges
	var sum float64 = 0
	for i := 0; i < polygon.num_vertices; i++ {
		if i == polygon.num_vertices-1 {
			sum = sum + (x[0]-x[i])*(y[0]-y[i])
		} else {
			sum = sum + (x[i+1]-x[i])*(y[i+1]-y[i])
		}

	}
	if sum > 0 {
		return true
	} else {
		return false
	}
}
