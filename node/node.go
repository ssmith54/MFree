package node

import (
	"Meshfree/geometry"
)

type Node struct {
	coords  geometry.Point
	node_nr int
}

// Create a new node in the domain
func NewNode(x, y, z float64, nr int) *Node {
	node := Node{geometry.NewPoint(x, y, z), nr}
	return &node
}

// return coordinates of a node
func (node *Node) GetNodalCoordinates() (float64, float64, float64) {
	x, y, z := node.coords.GetPointCoordinates()
	return x, y, z
}

func GetNodalDistance(a, b *Node) float64 {
	return geometry.GetDistance(&a.coords, &b.coords)
}

func (node *Node) GetPoint() *geometry.Point {
	return &node.coords
}

func FindNodesIn(nodes *[]Node, shape geometry.Shape) []Node {
	n := make([]Node, 0)
	for _, node := range *nodes {
		isIn := shape.IsPointInside(node.GetPoint())

		if isIn == true {
			n = append(n, node)
		}
	}
	return n
}
