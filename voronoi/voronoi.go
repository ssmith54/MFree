package voronoi

// #cgo CFLAGS: -Wall -I./C -I./C/jc_voronoi
// #cgo LDFLAGS: -lm -L./C/gpc -lgpc
//#include "C/jc_voronoi/jc_voronoi.h"
//#include "callVoronoi.h"
//#include "clipVoronoi.h"
//#include "./C/gpc/gpc.h"
//
import "C"
import (
	"unsafe"
)

type Voronoi struct {
	polygon      int
	num_polygons int
}

type _C_voronoi C.struct_gpc_polygon

func NewVoronoi(polygonIn int, num_polygonsIn int) *Voronoi {
	return &Voronoi{polygon: polygonIn, num_polygons: num_polygonsIn}
}

func GenerateClippedVoronoi(points []float64, boundary []int) *Voronoi {
	// make a call to jc_voronoi, just needs a set of points.

	var v Voronoi
	numpoints := C.int(len(points)) / 2

	// allocate C array for the points
	cPoints := ((*C.double)(unsafe.Pointer(&points[0])))
	cBoundary := ((*C.size_t)(unsafe.Pointer(&boundary[0])))
	numBoundary := C.int(len(boundary))
	var voronoi_out **C.gpc_polygon
	// allocate return structure
	//voronoi_out := C.malloc(C.size_t(1*gpc_polygon**)    )
	// clip said polygon, needs a GPC_polygon so will have to convert
	C.callVoronoi(&voronoi_out, cPoints, numpoints)
	// return a voronoi struct
	//
	C.clipVoronoi(&voronoi_out, cPoints, cBoundary, numpoints, numBoundary)
	// put voronoi out into a go slice so we can iterate over each voronoi cell

	return &v
}

func (voronoi *Voronoi) PrintVoronoi() {

}
