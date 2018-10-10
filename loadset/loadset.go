package loadset

import (
	"Meshfree/geometry"
	"Meshfree/node"
)

type Nodal_loads struct {
	nodes     *[]node.Node
	load_dir  *geometry.Dir
	magnitude float64
}

type Pressure_loads struct {
	surface  float64 // need to be able to generate a nodal surface
	load_dir *geometry.Dir
	pressure float64 // can be a function of time, temperature ect
}

func CreatePressureLoad() *Pressure_loads {
	return &(Pressure_loads{})
}

func CreateNodalLoad() *Nodal_loads {
	return &(Nodal_loads{})
}
