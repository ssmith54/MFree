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
	//
	a := *voronoi.GetVoronoiCells()

	// scni := make([]SCNI, voronoi.NumPolygons()
	// numMeshfreeNodes := meshfree.GetDomain().GetNumNodes()
	// phi_vec := make([]float64, numMeshfreeNodes)
	// for _, cell := range *voronoi.GetVoronoi() {
	// }
	//
	// return &scni

}
