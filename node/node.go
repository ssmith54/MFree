package node

import (
	"Meshfree/dof"
	"Meshfree/geometry"
	"fmt"
)

type Node struct {
	coords  geometry.Point
	node_nr int
	dofs    *[]dof.DOF
}

func Get_current_location() {

}

func Get_location() {

}

// Create a new node in the domain
func NewNode(x, y, z float64, nr int) *Node {
	node := Node{geometry.NewPoint(x, y, z), nr, new([]dof.DOF)}
	return &node
}

func (node *Node) GetNodenNr() int {
	return node.node_nr
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

func FindNodesIn(nodes *[]Node, shape geometry.Shape) *[]*Node {
	n := make([]*Node, 0)
	for i, node := range *nodes {
		isIn := shape.IsPointInside(node.GetPoint())

		if isIn == true {
			n = append(n, &((*nodes)[i]))
		}
	}
	return &n
}
func (node *Node) get_dof(dir *geometry.Dir) *dof.DOF {
	for _, dof := range *node.dofs {
		if dof.Get_direction().Is_Equal(dir) == true {
			return &dof
		}
	}
	return nil
}
func SetNodalDofs(nodes *[]*Node, dir *geometry.Dir, fix_type dof.Dof_fixture, dim int) {
	for _, node := range *nodes {
		dof := node.get_dof(dir)
		// check if dof already existed
		// if not create it
		if dof == nil {
			node.CreateDOF(dir, fix_type, dim)
			// else modify existing one
		} else {
			dof.Set_dof_type(fix_type)
		}
	}
}

func (node *Node) CreateDOF(dir *geometry.Dir, dof_type dof.Dof_fixture, global_dof_number int) {
	(*node.dofs) = append(*node.dofs, dof.NewDOF(global_dof_number, dir, dof_type))
}

func PrintNodes(nodes *[]*Node) {
	fmt.Printf("Nodes:\n")
	for _, node := range *nodes {
		fmt.Printf("%v\n", *node)

	}
}

// used in incremental analysis
