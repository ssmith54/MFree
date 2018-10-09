package coordinatesystem

import "Meshfree/geometry"

// cartesian
type CoordinateSystem interface {
	CreateBasis()
	ReturnBasis(dim int) *[]*geometry.Dir
}
type Cartesian struct {
	X *geometry.Dir
	Y *geometry.Dir
	Z *geometry.Dir
}

// returns a slice of the basis vectors for the domain
func (cart *Cartesian) ReturnBasis(dim int) *[]*geometry.Dir {
	dirs := make([]*geometry.Dir, dim)
	if dim == 1 {
		dirs[0] = cart.Y
	} else if dim == 2 {
		dirs[0] = cart.X
		dirs[1] = cart.Y
	} else if dim == 3 {
		dirs[0] = cart.X
		dirs[1] = cart.Y
		dirs[2] = cart.Z
	} else {
		// raise an error maybe
	}
	return &dirs
}

func (cs *Cartesian) CreateBasis() {
	cs.X = geometry.CreateDir(1, 0, 0)
	cs.Y = geometry.CreateDir(0, 1, 0)
	cs.Z = geometry.CreateDir(0, 0, 1)
}

func CreateCartesian() *Cartesian {
	cs := new(Cartesian)
	cs.X = geometry.CreateDir(1, 0, 0)
	cs.Y = geometry.CreateDir(0, 1, 0)
	cs.Z = geometry.CreateDir(0, 0, 1)
	return cs
}

// cylindircal
type Cylindrical int

const (
	r Cylindrical = 0
	z Cylindrical = 1
)

// spherical ( probably will not need this)
