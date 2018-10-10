package geometry

// contains geometric primatives

import (
	"errors"
	"math"
)

type Shape interface {
	IsPointInside(point *Point) bool
}

// represent a polygon using this struct
type Polygon struct {
	vertex_list  []Point
	num_vertices int
	area         float64
}

type Rectangle struct {
	A Point
	B Point
	C Point
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

func CreateDir(e1, e2, e3 float64) *Dir {
	return &(Dir{e1, e2, e3})
}

func (dir *Dir) Is_Equal(dir_1 *Dir) bool {
	if dir.e1 == dir_1.e1 && dir.e2 == dir_1.e2 && dir.e3 == dir_1.e3 {
		return true
	} else {
		return false
	}
}

type Vector struct {
	e1 float64
	e2 float64
	e3 float64
}

func createVector(p1, p2 Point) *Vector {
	return &(Vector{p2.x - p1.x, p2.y - p1.y, p2.z - p1.z})
}

func (segment *Segment) Normal(isCW bool) (float64, float64) {
	nx := segment.p2.y - segment.p1.y
	ny := segment.p2.x - segment.p1.x
	len := math.Sqrt(math.Pow(nx, 2) + math.Pow(ny, 2))
	nx = nx / len
	ny = ny / len
	// return direction
	if isCW == true {
		return nx, -ny
	} else {
		return -nx, ny
	}
}

func (segment *Segment) Length() float64 {
	return math.Sqrt(math.Pow((segment.p2.x-segment.p1.x), 2) +
		math.Pow((segment.p2.y-segment.p1.y), 2))
}

// IN ROUTINES, FINDS IF A POINT IS INSIDE A PRIMIATIVE SHAPE
func (circle *Circle) InCircle(points *[]Point) *[]bool {
	// shift coordinates
	isIn := make([]bool, len(*points))
	for i := 0; i < len(*points); i++ {
		shifted_point := subtract_points(&(*points)[i], &circle.center)
		shifted_radius := math.Sqrt(math.Pow(shifted_point.x, 2) + math.Pow(shifted_point.y, 2))
		if shifted_radius <= circle.radius {
			isIn[i] = true
		} else {
			isIn[i] = false
		}
	}
	return &isIn

}
func (cylinder *Cylinder) InCylinder(points *[]Point) bool {

	return true
}

// create a Rectangle
// b ------------
// |            |
// |            |
// a------------c
//
func CreateRectangle(ax, ay, bx, by, cx, cy float64) *Rectangle {
	return &(Rectangle{NewPoint(ax, ay, 0), NewPoint(bx, by, 0), NewPoint(cx, cy, 0)})
}

func (point *Point) dot(a *Point) float64 {
	return point.x*a.x + point.y*a.y + point.z*a.z
}

func (vector *Vector) dot(a *Vector) float64 {

	return vector.e1*a.e1 + vector.e2*a.e2 + vector.e3*a.e3

}

func (rectangle *Rectangle) IsPointInside(point *Point) bool {
	AB := createVector(rectangle.B, rectangle.A)
	BC := createVector(rectangle.C, rectangle.B)

	var isIn bool
	AM := createVector((*point), rectangle.A)
	BM := createVector((*point), rectangle.B)
	if (0 <= AB.dot(AM)) && (AB.dot(AM) <= AB.dot(AB)) &&
		(BC.dot(BM) <= BC.dot(BC)) {
		isIn = true
	} else {
		isIn = false
	}
	return isIn
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
func (polygon *Polygon) ReturnArea() float64 {
	return polygon.area

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

func (polygon *Polygon) GetPolygonPoints() *[]Point {
	return &polygon.vertex_list

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
