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

func New_node(x, y, z float64, nr int) *Node {
	node := Node{Point{x, y, z}, nr}
	return &node
}

func Get_distance(a *Point, b *Point) float64 {
	distance := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))

	return distance
}

func (domain *Domain) Add_nodes(nodes ...*Node) int {
	for _, nodes := range nodes {
		domain.Nodes = append(domain.Nodes, *nodes)
	}
	return 1
}
