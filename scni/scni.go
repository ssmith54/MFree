package scni

import (
	"Meshfree/shapefunctions"
	"Meshfree/voronoi"

	"gonum.org/v1/gonum/mat"
)

type SCNI struct {
	attribute  float64
	Bmat       *mat.Dense
	contribute []int
	node_nr    int
	volume     float64
}

func CreateSCNI(meshfree *shapefunctions.Meshfree, voronoi *voronoi.Voronoi) {
	cells := *voronoi.ReturnVoronoiCells()

	//phi_vec := make([]float64, meshfree.GetDomain().GetNumNodes())
	for _, cell := range cells {
		points := cell.GetPolygonPoints()

		//isCw := cell.IsClockwise()
		// find shape functions at each point
		_, _, _, err := meshfree.ComputeMeshfree(&(*points)[0], 1, false)
		if err != nil {
			panic(err)
		}

	}

}
