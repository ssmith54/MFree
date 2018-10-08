package coordinatesystem

import "Meshfree/geometry"

// cartesian

type Cartesian struct {
	X *geometry.Dir
	Y *geometry.Dir
	Z *geometry.Dir
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
