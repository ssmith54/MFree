package domain

// #cgo CFLAGS: -Wall -I./C -I./C/triangle
// #cgo LDFLAGS: -L./C/triangle -ltriangle -lm
//#include "C/trigen.h"
//
import "C"
import (
	"Meshfree/geometry"
	"Meshfree/voronoi"
	"fmt"
	"unsafe"
)

type Node struct {
	coords  geometry.Point
	node_nr int
}

type Domain struct {
	Nodes            []Node
	num_nodes        int
	voronoi          *voronoi.Voronoi
	dim              int
	boundarySegments []int
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
	//fileName_C := C.CString(fileName)
	//options_C := C.CString(options)
	// define the C outputs and inputs
	var points *C.double
	var boundary *C.int
	var num_points C.int
	var num_boundary C.int

	fileNameIn := C.CString("preform")
	optionsIn := C.CString("pDa1q0")
	C.trigen(&points, &boundary, optionsIn, fileNameIn, &num_points, &num_boundary)

	// points
	unsafePtr_points := unsafe.Pointer(points)
	arrayPtr_points := (*[1 << 30]C.double)(unsafePtr_points)
	length_points := int(num_points)
	slice_points := arrayPtr_points[0 : 2*length_points]
	nodes := make([]Node, num_points)
	//

	fmt.Printf("There are %v points :%v \n", num_points, slice_points)
	for i := 0; i < int(num_points); i++ {
		nodes[i] = *(NewNode(float64(slice_points[2*i]), float64(slice_points[2*i+1]), 0, i))
	}
	fmt.Printf("Finding boundary points \n")
	// store boundary
	unsafePtr_boundary := unsafe.Pointer(boundary)
	arrayPtr_boundary := (*[1 << 30]C.int)(unsafePtr_boundary)
	length_boundary := int(num_boundary)
	fmt.Printf("numBoundary = %v", num_boundary)

	boundary_nodes := arrayPtr_boundary[0:(2*length_boundary - 1)]
	domain.boundarySegments = make([]int, 2*num_boundary-1)
	for i := 0; i < 2*length_boundary-1; i++ {
		domain.boundarySegments[i] = int(boundary_nodes[i])

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

}
