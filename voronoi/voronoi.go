package voronoi

// #cgo CFLAGS: -Wall -I./C
// #cgo LDFLAGS: -lm -L/C/gpc -lgpc
//#include "C/jc_voronoi/jc_voronoi.h"
//
type Voronoi struct {
	polygon      int
	num_polygons int
}

func NewVoronoi(polygonIn int, num_polygonsIn int) *Voronoi {
	return &Voronoi{polygon: polygonIn, num_polygons: num_polygonsIn}
}

func GenerateVoronoi() {
	// make a call to jc_voronoi, just needs a set of points.

	// clip said polygon, needs a GPC_polygon so will have to convert

	// return a voronoi struct
}
