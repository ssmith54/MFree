package domain

import "math"

type Point struct {
	x float64
	y float64
	z float64
}

type Node struct {
	coords  Point
	node_nr int
}

type Domain struct {
	Nodes     []Node
	num_nodes int
}

// Create a new domain
func NewDomain(nodes []Node, numnodes int) Domain {
	return Domain{nodes, numnodes}
}

// Create a new node in the domain
func NewNode(x, y, z float64, nr int) *Node {
	node := Node{Point{x, y, z}, nr}
	return &node
}

// get distance between points
func GetDistance(a *Point, b *Point) float64 {
	distance := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))

	return distance
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

// Return corodinates of point
func (point *Point) GetPointCoordinates() (float64, float64, float64) {
	x := point.x
	y := point.y
	z := point.z

	return x, y, z
}

// Get number of nodes
func (domain *Domain) GetNumNodes() int {
	return domain.num_nodes
}
