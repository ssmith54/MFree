package dof

import "Meshfree/geometry"

type Dof_fixture int

const (
	DOF_FREE       Dof_fixture = 0
	DOF_FIXED      Dof_fixture = 1
	DOF_PRESCIRBED Dof_fixture = 2
)

type DOF struct {
	global_dof_nr int
	fix_type      Dof_fixture
	dir           *geometry.Dir
	value         float64
}

// constructor

func NewDOF(nodenum int, dirs *geometry.Dir, dof_type Dof_fixture) DOF {
	return DOF{nodenum, dof_type, dirs, 0}
}

// set direction of degree of freedom
func (dof *DOF) Set_direction(dir *geometry.Dir) {
	dof.dir = dir
}

// set dof_type, e.g fixed free
func (dof *DOF) Set_dof_type(dof_type Dof_fixture) {
	dof.fix_type = dof_type
}

// get direction
func (dof *DOF) Get_direction() *geometry.Dir {
	return dof.dir
}

// get global dof number
func (dof *DOF) Get_global_dof() int {
	return dof.global_dof_nr
}

// set global dof number
func (dof *DOF) Set_global_dof(dof_nr int) {
	dof.global_dof_nr = dof_nr
}

/////////////////////////////////////
// check what type of dof it is
// free
func (dof *DOF) Is_free() bool {
	if dof.fix_type == 0 {
		return true
	} else {
		return false
	}
}

// fixed
func (dof *DOF) Is_fixed() bool {
	if dof.fix_type == 1 {
		return true
	} else {
		return false
	}
}

// prescribed
func (dof *DOF) Is_prescribed() bool {
	if dof.fix_type == 2 {
		return true
	} else {
		return false
	}
}
