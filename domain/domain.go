package domain

import (
	"Meshfree/geometry"
)

type Node struct {
	coords  geometry.Point
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
