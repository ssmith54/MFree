package loadset

import (
	"Meshfree/geometry"
	"Meshfree/node"
)

type LoadSet struct {
	Nodal_loads    *[]*Nodal_load
	Pressure_loads *[]*Pressure_load
}

// returns the external load vector, built up from the loads of
// this load set
func (loadset *LoadSet) get_load_vector() {

}

type Nodal_load struct {
	nodes     *[]*node.Node
	load_dir  *geometry.Dir
	magnitude float64
}

type Pressure_load struct {
	surface  float64 // need to be able to generate a nodal surface
	load_dir *geometry.Dir
	pressure float64 // can be a function of time, temperature ect
}

func CreatePressureLoad() *Pressure_load {
	return &(Pressure_load{})
}

func (nodal_load *Nodal_load) get_load_vector() {

}
func (pressure_load *Pressure_load) get_load_vector() {

}

func CreateNodalLoad(nodes_in *[]*node.Node, load_dir_in *geometry.Dir, magnitude_in float64) *Nodal_load {
	return &(Nodal_load{nodes: nodes_in, load_dir: load_dir_in, magnitude: magnitude_in})
}

func (nodal_load *Nodal_load) SetMagnitude(magnitude_in float64) {
	nodal_load.magnitude = magnitude_in
}
