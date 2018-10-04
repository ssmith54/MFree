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
	"Meshfree/geometry"
	"unsafe"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

// polygon abstraction

type Voronoi struct {
	polygon_list []geometry.Polygon
	num_polygons int
}

func NewVoronoi(polygonIn []geometry.Polygon, num_polygonsIn int) *Voronoi {
	return &Voronoi{polygon_list: polygonIn, num_polygons: num_polygonsIn}
}

func (voronoi *Voronoi) NumPolygons() int {
	return voronoi.num_polygons
}

func (voronoi *Voronoi) AddPolygon(polygon *geometry.Polygon) {
	voronoi.polygon_list = append(voronoi.polygon_list, *polygon)
	voronoi.num_polygons++
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

	// create the voronoi diagram
	C.callVoronoi(&voronoi_out, cPoints, numpoints)

	// clip voronoi diagram
	var vorPoints *C.double
	var vorNumVertex *C.int
	var totalVertex C.int
	C.clipVoronoi(&vorPoints, &vorNumVertex, &totalVertex, &voronoi_out, cPoints, cBoundary, numpoints, numBoundary)
	// put voronoi out into a go slice so we can iterate over each voronoi cell

	// unsafe poiner to the c array
	unsafePtr_points := unsafe.Pointer(vorPoints)
	// convert unsafe pointer to a pointer of the type *[1 << 30]C.double
	arrayPtr_points := (*[1 << 30]C.double)(unsafePtr_points)
	// slice the array into a go slice, with same backing as the C array
	// make sure to specify the capacity as well as the length
	length_points := 2 * C.int(totalVertex)
	slice_points := arrayPtr_points[0:length_points:length_points]

	unsafePtr_numVertex := unsafe.Pointer(vorNumVertex)
	arrayPtr_numVertex := (*[1 << 30]C.int)(unsafePtr_numVertex)
	length_vertex := numpoints
	slice_numVertex := arrayPtr_numVertex[0:length_vertex:length_vertex]

	// create voronoi diagram
	var count int = 0
	for i := 0; i < int(numpoints); i++ {
		num_vertex := int(slice_numVertex[i])
		x := make([]float64, num_vertex)
		y := make([]float64, num_vertex)
		for k := 0; k < num_vertex; k++ {
			x[k] = float64(slice_points[2*count])
			y[k] = float64(slice_points[2*count+1])
			count++
		}
		poly, err := geometry.CreatePolygon(x, y)
		if err != nil {
			panic(err)
		} else {
			v.AddPolygon(poly)
		}
	}
	return &v
}

func (voronoi *Voronoi) PrintVoronoi() {

}
func (voronoi *Voronoi) ReturnVoronoiCells() *[]geometry.Polygon {
	return &voronoi.polygon_list
}

func (voronoi *Voronoi) PrintVoronoiToImg(filename string) {
	var xmin, xmax, ymin, ymax float64 = 0, 0, 0, 0
	p, _ := plot.New()
	for _, poly := range voronoi.polygon_list {

		x, y := poly.GetPolyCoordinates()
		type XY struct {
			X float64
			Y float64
		}
		var polygon_coords plotter.XYs
		for i := 0; i < poly.GetNumVertices(); i++ {
			temp := XY{x[i], y[i]}
			polygon_coords = append(polygon_coords, temp)

			if x[i] > xmax {
				xmax = x[i]
			}
			if x[i] < xmin {
				xmin = x[i]
			}
			if y[i] > ymax {
				ymax = y[i]
			}
			if y[i] < ymin {
				ymin = x[i]
			}

		}

		polygon, _ := plotter.NewPolygon(polygon_coords)
		p.Add(polygon)
	}

	// this is used to replicate the axis equal option of matlab
	scale_y := (xmax - xmin) / (ymax - ymin)
	scale_x := 1 / scale_y
	if scale_y > 1 {
		outer1 := plotter.XYs{{X: xmin, Y: ymin * scale_y}, {X: 1.1 * xmax, Y: ymin * scale_y}, {X: 1.1 * xmax, Y: ymax * scale_y}, {X: xmin, Y: ymax * scale_y}}
		poly, _ := plotter.NewPolygon(outer1)
		p.Add(poly)
	} else {
		outer1 := plotter.XYs{{X: xmin * scale_x, Y: ymin}, {X: xmax * scale_x, Y: ymin}, {X: xmax * scale_x, Y: ymax}, {X: xmin * scale_x, Y: ymax}}
		poly, _ := plotter.NewPolygon(outer1)
		p.Add(poly)
	}

	p.HideY()
	p.HideX()

	p.Save(1000, 1000, filename)

}
