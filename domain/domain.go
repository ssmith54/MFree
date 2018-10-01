package domain

// #cgo CFLAGS: -Wall -I./C -I./C/triangle
// #cgo LDFLAGS: -L./C/triangle -ltriangle -lm
//#include "C/trigen.h"
//
import "C"
import (
	"Meshfree/geometry"
	"Meshfree/voronoi"
	"unsafe"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Node struct {
	coords  geometry.Point
	node_nr int
}

type Domain struct {
	Nodes         []Node
	num_nodes     int
	voronoi       *voronoi.Voronoi
	dim           int
	boundaryNodes []int
}

// Create a new domain
func NewDomain(nodes []Node, numnodes int) Domain {
	return Domain{nodes, numnodes, nil, 0, nil}
}

// Create a new node in the domain
func NewNode(x, y, z float64, nr int) *Node {
	node := Node{geometry.NewPoint(x, y, z), nr}
	return &node
}

// add nodes to domain
func (domain *Domain) AddNodes(nodes ...*Node) int {
	for _, nodes := range nodes {
		domain.Nodes = append(domain.Nodes, *nodes)
		domain.num_nodes = domain.num_nodes + 1
	}
	return 1
}

// return coordinates of a node
func (domain *Domain) GetNodalCoordinates(indx int) (float64, float64, float64) {
	x, y, z := domain.Nodes[indx].coords.GetPointCoordinates()
	return x, y, z
}

// Get number of nodes
func (domain *Domain) GetNumNodes() int {
	return domain.num_nodes
}

func GetNodalDistance(a, b *Node) float64 {
	return geometry.GetDistance(&a.coords, &b.coords)
}

// Copy a domain
func (domain *Domain) copyDomain() *Domain {
	return &Domain{domain.Nodes, domain.num_nodes, nil, 0, nil}
}

// update the domain based on the displacement
func (domain *Domain) UpdateDomain() {

}

// read in nodes from a files

// generate nodes using triangle ( A two dimensional meshfree generator )
func (domain *Domain) TriGen(fileName string, options string) {

	// convert options and filename into a C string
	fileName_C := C.CString(fileName)
	options_C := C.CString(options)
	// define the C outputs and inputs
	var points *C.double
	var boundary *C.int
	var num_points C.int
	var num_boundary C.int

	C.trigen(&points, &boundary, options_C, fileName_C, &num_points, &num_boundary)

	// points
	unsafePtr_points := unsafe.Pointer(points)
	arrayPtr_points := (*[1 << 30]C.double)(unsafePtr_points)
	length_points := int(num_points)
	slice_points := arrayPtr_points[0 : 2*length_points]
	nodes := make([]Node, num_points)
	//

	for i := 0; i < int(num_points); i++ {
		nodes[i] = *(NewNode(float64(slice_points[2*i]), float64(slice_points[2*i+1]), 0, i))
	}
	// store boundary

	// void pointer to c array
	unsafePtr_boundary := unsafe.Pointer(boundary)
	// Convert C array into go array so we can access elements
	arrayPtr_boundary := (*[1 << 30]C.int)(unsafePtr_boundary)
	// find number of nodes to store
	length_boundary := int(num_boundary)
	boundary_nodes := arrayPtr_boundary[0:(2 * (length_boundary - 1))]
	// Put boundary nodes into domain
	domain.boundaryNodes = make([]int, num_boundary)
	for i := 0; i < length_boundary-1; i++ {
		domain.boundaryNodes[i] = int(boundary_nodes[2*i] - 1)

	}
	// set up domain
	domain.Nodes = nodes
	domain.num_nodes = int(num_points)

	// Need to free the C memory
}

// for 3D geometry
func (domain *Domain) TetGen(fileName []string) {

}

// generate the voronoi diagram
func (domain *Domain) GenerateVoronoi() {

	// have to convert nodes it raw points
	var x, y float64
	points := make([]float64, domain.num_nodes*2)
	for i := 0; i < domain.num_nodes; i++ {
		x, y, _ = domain.GetNodalCoordinates(i)
		points[2*i] = x
		points[2*i+1] = y
	}
	domain.voronoi = voronoi.GenerateVoronoi(points, domain.boundaryNodes)

}

func (domain *Domain) PrintNodesToFile(filename string) {

}

// print the domain to an img file
func (domain *Domain) PrintNodesToImg(imagename string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Nodes"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	var X, Y float64

	// scatter nodes
	pts := make(plotter.XYs, domain.num_nodes)
	for i := range pts {
		X, Y, _ = domain.GetNodalCoordinates(i)
		pts[i].X = X
		pts[i].Y = Y
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	pts_boundary := make(plotter.XYs, len(domain.boundaryNodes))
	for i := range pts_boundary {
		indx := domain.boundaryNodes[i]
		X, Y, _ = domain.GetNodalCoordinates(indx)
		pts_boundary[i].X = X
		pts_boundary[i].Y = Y
	}
	lpLine, _, err := plotter.NewLinePoints(pts_boundary)
	if err != nil {
		panic(err)
	}

	// add line and scatter to plot
	p.Add(lpLine, s)

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

}
