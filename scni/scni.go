package scni

import (
	"Meshfree/geometry"
	"Meshfree/shapefunctions"
	"Meshfree/voronoi"
	"fmt"
	"runtime"
	"sync"
	"time"

	"gonum.org/v1/gonum/mat"
)

type SCNI struct {
	attribute      float64
	Bmat           *mat.Dense
	contribute     []int
	node_nr        int
	volume         float64
	num_neighbours int
}

// create a cell function which just takes a polygon as its input
func createCell(meshfree *shapefunctions.Meshfree, poly *geometry.Polygon) *SCNI {

	// variables
	var nx_1, ny_1, seg_length_1 float64 = 0, 0, 0
	var nx_0, ny_0, seg_length_0 float64 = 0, 0, 0
	var n_0, n_1 int = 0, 0
	var bI1, bI2 float64 = 0, 0
	scni := new(SCNI)
	bI1v := make([]float64, 0)
	bI2v := make([]float64, 0)
	neigh_cell := make([]int, 0)

	// get points
	points := poly.GetPolygonPoints()
	numPoints := len(*points)
	isCw := poly.IsClockwise()

	// first segment
	seg_0 := geometry.NewSegment(&(*points)[0], &(*points)[1])
	nx_0, ny_0 = seg_0.Normal(isCw)
	seg_length_0 = seg_0.Length()
	// loop over each vertex of the polygon
	for i := 0; i < numPoints; i++ {
		// create segments
		// check for recursive, i.e numPoints + 1 = 0
		if i < numPoints-1 {
			n_1 = i + 1
			n_0 = i
		} else {
			n_1 = 0
			n_0 = i
		}
		seg_1 := geometry.NewSegment(&(*points)[n_0], &(*points)[n_1])
		nx_1, ny_1 = seg_1.Normal(isCw)
		//fmt.Printf("nx = %v ny = %v \n", nx_1, ny_1)
		seg_length_1 = seg_1.Length()
		phi, _, neighbours, err := meshfree.ComputeMeshfree(&(*points)[n_1], 1, false)
		if err != nil {
			panic(err)
		}

		for k := 0; k < len(*neighbours); k++ {
			bI1 = 0.5 * (nx_0*seg_length_0 + nx_1*seg_length_1) * phi.AtVec(k)
			bI2 = 0.5 * (ny_0*seg_length_0 + ny_1*seg_length_1) * phi.AtVec(k)
			indx := findInt(neigh_cell, (*neighbours)[k])
			if indx >= 0 {
				bI1v[indx] += bI1
				bI2v[indx] += bI2
			} else {
				neigh_cell = append(neigh_cell, (*neighbours)[k])
				bI1v = append(bI1v, bI1)
				bI2v = append(bI2v, bI2)
			}

		}
		// update previous variables
		nx_0, ny_0, seg_length_0 = nx_1, ny_1, seg_length_1
	} // end of loop over segments
	scni.num_neighbours = len(neigh_cell)
	scni.contribute = neigh_cell

	// assume its just a 2D plane strain or plane stress problem at the moment
	Bmat := mat.NewDense(3, 2*scni.num_neighbours, nil)
	// now have bi1 and bi2 vectors, need to make B matrix
	for i := 0; i < scni.num_neighbours; i++ {
		Bmat.Set(0, 2*i, bI1v[i])
		Bmat.Set(1, 2*i+1, bI2v[i])
		Bmat.Set(2, 2*i, bI2v[i])
		Bmat.Set(2, 2*i+1, bI1v[i])
	}
	scni.Bmat = Bmat
	// set volume of cell
	scni.volume = poly.ReturnArea()
	return scni
}

func findInt(a []int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
func CreateSCNI(meshfree *shapefunctions.Meshfree, voronoi *voronoi.Voronoi) {

	numnodes := meshfree.GetDomain().GetNumNodes()
	cells := *voronoi.ReturnVoronoiCells()

	// how many processors to run
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	// create the channel
	SCNIs := make([]*SCNI, numnodes)
	//phi_vec := make([]float64, meshfree.GetDomain().GetNumNodes())
	start := time.Now()
	fmt.Printf("Creating SCNI cells\n ")
	// create cells
	wg.Add(numnodes)
	for i, cell := range cells {
		go func(cell geometry.Polygon, i int) {
			defer wg.Done()
			SCNIs[i] = createCell(meshfree, &cell)
		}(cell, i)
	}

	wg.Wait()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Finished creating SCNI cells in %v\n", elapsed)

	var area float64 = 0
	for i := 0; i < len(SCNIs); i++ {
		area = area + SCNIs[i].volume

	}
	fmt.Printf("Example B=\n%v\n", mat.Formatted(SCNIs[0].Bmat, mat.Prefix(""), mat.Squeeze()))

	fmt.Printf("area = %v\n", area)
}
