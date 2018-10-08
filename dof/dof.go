package dof

import "Meshfree/geometry"

type dof_fixture int

const (
	dof_free       dof_fixture = 0
	dof_fixed      dof_fixture = 1
	dof_prescribed dof_fixture = 2
)

type DOF struct {
	node_nr       int
	global_dof_nr int
	fix_type      dof_fixture
	dir           *geometry.Dir
}

// set direction of degree of freedom
func (dof *DOF) Set_direction(dir *geometry.Dir) {
	dof.dir = dir
}

// set dof_type, e.g fixed free
func (dof *DOF) Set_dof_type(dof_type dof_fixture) {
	dof.fix_type = dof_type
}

// get direction
func (dof *DOF) get_direction() *geometry.Dir {
	return dof.dir
}

// get global dof number
func (dof *DOF) get_global_dof() int {
	return dof.global_dof_nr
}

// set global dof number
func (dof *DOF) set_global_dof(dof_nr int) {
	dof.global_dof_nr = dof_nr
}

/////////////////////////////////////
// check what type of dof it is
// free
func (dof *DOF) is_free() bool {
	if dof.fix_type == 0 {
		return true
	} else {
		return false
	}
}

// fixed
func (dof *DOF) is_fixed() bool {
	if dof.fix_type == 1 {
		return true
	} else {
		return false
	}
}

// prescribed
func (dof *DOF) is_prescribed() bool {
	if dof.fix_type == 2 {
		return true
	} else {
		return false
	}
}