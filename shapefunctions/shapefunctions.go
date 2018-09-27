package shapefunctions

import (
	"Meshfree/domain"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Point struct {
	x float64
	y float64
	z float64
}

type Node struct {
	coords  Point
	node_nr int
}

type Meshfree struct {
	domainSize         float64
	nodalSpacing       []float64
	gamma              []float64
	isConstantSpacing  bool
	isVariousPoints    bool
	basisFunctionRadii []float64
	dim                int
}

func get_distance(a *Point, b *Point) float64 {
	distance := math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
	return distance
}

func (meshfree *domain.Meshfree) get_shifted_coordinates(p *Point, m *mat.Dense) int {
	num_r, num_c := m.Dims()
	if num_c == 1 {
		for i := 0; i < num_r; i++ {
			m.Set(i, 0, -p.x+meshfree.Nodes[i].coords.x)
		}
	}
	if num_c == 2 {
		for i := 0; i < num_r; i++ {
			m.Set(i, 0, -p.x+meshfree.Nodes[i].coords.x)
			m.Set(i, 1, -p.y+meshfree.Nodes[i].coords.y)
		}
	}

	if num_c == 3 {
		for i := 0; i < num_r; i++ {
			m.Set(i, 0, -p.x+meshfree.Nodes[i].coords.x)
			m.Set(i, 1, -p.y+meshfree.Nodes[i].coords.y)
			m.Set(i, 2, -p.z+meshfree.Nodes[i].coords.z)
		}
	}

	return 2
}

// compute meshfree shape functions
func (meshfree *domain.Meshfree) compute_meshfree(p *Point, dim int, compute int, tol float64) int {

	fmt.Printf("----------------------------------------------------------------------------\n")
	fmt.Printf("Constructing meshfree shape functions at p = %v\n\n", p)
	// compute shifted coordindates
	shifted_coordinates := mat.NewDense(meshfree.num_nodes, dim, nil)
	meshfree.get_shifted_coordinates(p, shifted_coordinates)

	// get neighbours of point p
	neighbours := meshfree.get_neighbours(shifted_coordinates)
	fmt.Printf("-%v neighbours of point p=%v\n", len(neighbours), neighbours)

	// find contributing shifted_coordinates xS_c and
	xS_c := mat.NewDense(len(neighbours), dim, nil)
	for i := 0; i < len(neighbours); i++ {
		xS_c.SetRow(i, shifted_coordinates.RawRowView(neighbours[i]))
	}

	// set up storage w(x) and dw() [ if compute = 2]
	weight, weightDer := prior("cubic_spline", neighbours, xS_c, meshfree)

	// branch for either MLS or MAXENT

	// MAXENT
	// Find phi from solving the newton raphson problem
	fmt.Printf("-Starting Newton-Raphson iterations to find phi\n")
	phi, err := f_of_lamdba(dim, weight, xS_c, tol, 100)
	//phiDer := mat.NewDense(len(neighbours), dim, nil)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Printf("Finding shape function derivatives\n ")

	// compute phiDer matricies
	if compute == 2 {
		compute_phiDer(phi, xS_c, weight, weightDer)
	}

	// MLS
	fmt.Printf("----------------------------------------------------------------------------\n")

	return 1
}

func outer_product_vectors(a *mat.VecDense, b *mat.VecDense) *mat.Dense {
	dim := a.Len()

	c := mat.NewDense(dim, dim, nil)

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			c.Set(i, j, a.AtVec(i)*b.AtVec(j))
		} // end of loop over j
	} // end of loop over i
	return c
}

func compute_phiDer(phi *mat.VecDense, xS_c *mat.Dense, weight *mat.VecDense, weightDer *mat.Dense) int {
	// compute phi der = phi * (x_s . (inv(H) - inv(H).A  ) + MA - MC )
	// MA = weightDer/weight,
	// MC = phi.MA
	// A = phi.xs (dyad) MA

	num_points, dim := xS_c.Dims()
	phiDer := mat.NewDense(num_points, dim, nil)
	if dim == 1 {
		// Hessian is just a scalar which is grad(grad(F)) = o
		xS_sq := mat.NewVecDense(num_points, nil)
		xS_sq.MulElemVec(xS_c.ColView(0), xS_c.ColView(0))

		//H := mat.Dot(phi, xS_sq)

	} else {
		// compute matricies needed for derivative
		H := mat.NewDense(dim, dim, nil)
		MA := mat.NewDense(num_points, dim, nil)
		A := mat.NewDense(dim, dim, nil)
		MC_slice := make([]float64, dim)

		// initiailse
		dyad_xS := mat.NewDense(dim, dim, nil)
		dyad_xS_MA := mat.NewDense(dim, dim, nil)

		for i := 0; i < num_points; i++ {
			xS_vec := mat.NewVecDense(dim, xS_c.RawRowView(i))
			// compute Hessian
			dyad_xS = outer_product_vectors(xS_vec, xS_vec)
			dyad_xS.Scale(phi.AtVec(i), dyad_xS)
			H.Add(H, dyad_xS)

			// compute MA
			for j := 0; j < dim; j++ {
				MA.Set(i, j, weightDer.At(i, j)/weight.AtVec(i))
			}

			// compute A
			maRow := make([]float64, dim)
			copy(maRow, MA.RawRowView(i))
			MA_vec := mat.NewVecDense(dim, maRow)
			dyad_xS_MA = outer_product_vectors(xS_vec, MA_vec)
			dyad_xS_MA.Scale(phi.AtVec(i), dyad_xS_MA)
			A.Add(A, dyad_xS_MA)
			MA_vec.ScaleVec(phi.AtVec(i), MA_vec)

			for j := 0; j < dim; j++ {
				MC_slice[j] += MA_vec.AtVec(j)
			}

		} // end of loop over points

		MC_mat := mat.NewDense(num_points, dim, nil)
		phi_mat := mat.NewDense(num_points, num_points, nil)
		for i := 0; i < num_points; i++ {
			MC_mat.SetRow(i, MC_slice)
			phi_mat.Set(i, i, phi.AtVec(i))
		}

		invH := mat.NewDense(dim, dim, nil)
		invH.Inverse(H)

		temp1 := mat.NewDense(dim, dim, nil)
		mat1 := mat.NewDense(dim, dim, nil)
		temp1.Mul(invH, A)
		mat1.Sub(invH, temp1)

		// term_1 = xS_c*(inv(H) - inv(H)*A)
		term_1 := mat.NewDense(num_points, dim, nil)
		term_1.Mul(xS_c, mat1)

		// term 2 = term1 + MA
		term_2 := mat.NewDense(num_points, dim, nil)
		term_2.Add(term_1, MA)

		// term 1 = term2 - MC_mat
		term_1.Reset()
		term_1.Sub(term_2, MC_mat)

		phiDer.Mul(phi_mat, term_1)

		fmt.Printf("H=\n%v\n", mat.Formatted(H, mat.Prefix(""), mat.Squeeze()))
		fmt.Printf("MA=\n%v\n", mat.Formatted(MA, mat.Prefix(""), mat.Squeeze()))
		fmt.Printf("A=\n%v\n", mat.Formatted(A, mat.Prefix(""), mat.Squeeze()))
		fmt.Printf("weight=\n%v\n", mat.Formatted(weight, mat.Prefix(""), mat.Squeeze()))
		fmt.Printf("MC_Mat=\n%v\n", mat.Formatted(MC_mat, mat.Prefix(""), mat.Squeeze()))
		fmt.Printf("weightDer=\n%v\n", mat.Formatted(weightDer, mat.Prefix(""), mat.Squeeze()))

		fmt.Printf("phiDer=\n%v\n", mat.Formatted(phiDer, mat.Prefix(""), mat.Squeeze()))

	} // end of dim 2 or 3

	return 1
}

func f_of_lamdba(dim int, weight *mat.VecDense, xs *mat.Dense, tol float64, max_iter int) (*mat.VecDense, error) {

	// initialisation
	// initial guess of lambda vector
	lambda := mat.NewVecDense(dim, nil)
	// inital residual
	residual := mat.NewVecDense(dim, nil)
	dlam := mat.NewVecDense(dim, nil)
	for i := 0; i < dim; i++ {
		residual.SetVec(i, 100)
		dlam.SetVec(i, 100)
	}
	var num_iterations int = 0

	// set up iterative scheme matricies
	num_points, _ := xs.Dims()
	Zi := mat.NewVecDense(num_points, nil)
	phi := mat.NewVecDense(num_points, nil)
	xs_sq := mat.NewDense(num_points, dim, nil)
	var Z float64 = 0

	// if dim == 1 newton raphson scheme is operating on scalars eg
	// lambda, residual and the hessian are all scalars
	// for consistent return values lambda is stored as a vector
	// raphson equation is grad(F) = 0, where F = log(Z)

	if dim == 1 {
		for i := 0; i < num_points; i++ {
			Zi.SetVec(i, weight.AtVec(i)*math.Exp(-(mat.Dot(xs.RowView(i), lambda))))
			Z += Zi.AtVec(i)
			xs_sq.Set(i, 0, xs.At(i, 0)*xs.At(i, 0))

		}
		phi.ScaleVec(1.00/Z, Zi)

		// begin iterative scheme, F(lambda) = grad(Z) = 0
		// F(lambda) = grad(Z)
		residual := -mat.Dot(phi, xs.ColView(0))
		// grad(F(lambda)) = hessian ( Z )
		hF := mat.Dot(phi, xs_sq.ColView(0)) - residual*residual

		// *************************************
		// 			NEWTON RAPHSON ITERATIONS
		//     dLambda = -inv(hF)*residual
		// *************************************
		for math.Abs(residual) > tol {

			// find Jacobian for Newton iterations
			dlam.SetVec(0, -residual/hF)
			lambda.SetVec(0, lambda.AtVec(0)+dlam.AtVec(0))

			// find hf and residiual at k+1th iteration
			Z = 0

			for i := 0; i < num_points; i++ {
				Zi.SetVec(i, weight.AtVec(i)*math.Exp(-(mat.Dot(xs.RowView(i), lambda))))
				Z += Zi.AtVec(i)
				xs_sq.Set(i, 0, xs.At(i, 0)*xs.At(i, 0))

			}
			phi.ScaleVec(1.00/Z, Zi)

			// begin iterative scheme, F(lambda) = grad(Z) = 0
			// F(lambda) = grad(Z)
			residual = -mat.Dot(phi, xs.ColView(0))
			// grad(F(lambda)) = hessian ( Z )
			hF = mat.Dot(phi, xs_sq.ColView(0)) - residual*residual

			num_iterations++

			if num_iterations > 100 {
				err1 := errors.New("Newton Raphson iterations failed, max number of iterations exceeded")

				return phi, err1
			}

		}
		// end of dim == 1

		// DIM = 2,3
	} else {
		for i := 0; i < num_points; i++ {
			Zi.SetVec(i, weight.AtVec(i)*math.Exp(-(mat.Dot(xs.RowView(i), lambda))))
			Z += Zi.AtVec(i)
		}
		phi.ScaleVec(1.00/Z, Zi)
		// begin iterative scheme, F(lambda) = grad(Z) = 0
		// F(lambda) = grad(Z) = 0
		residual := mat.NewVecDense(dim, nil)
		residual.MulVec(xs.T(), phi)
		residual.ScaleVec(-1, residual)
		// grad(F(lambda)) = hessian ( Z )
		hF := mat.NewDense(dim, dim, nil)
		temp1 := mat.NewVecDense(num_points, nil) // xs(i) .* xs(j)

		//temp2 := mat.NewVecDense(dim, data)// residual ⊗ residual ( outer product)
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				temp1.MulElemVec(xs.ColView(i), xs.ColView(j))
				hF.Set(i, j, mat.Dot(phi, temp1)-residual.AtVec(i)*residual.AtVec(j))
			}

		}

		// *************************************
		// 			NEWTON RAPHSON ITERATIONS
		//     dLambda = -inv(hF)*residual
		// *************************************
		for math.Sqrt(mat.Dot(residual, residual)) > tol {

			// find Jacobian for Newton iterations
			residual.ScaleVec(-1.00, residual)
			hF.Inverse(hF)
			dlam.MulVec(hF, residual)
			lambda.AddVec(lambda, dlam)

			// find hf and residiual at k+1th iteration
			Z = 0
			// partition functionss
			for i := 0; i < num_points; i++ {
				Zi.SetVec(i, weight.AtVec(i)*math.Exp(-(mat.Dot(xs.RowView(i), lambda))))
				Z += Zi.AtVec(i)
			}
			phi.ScaleVec(1.00/Z, Zi)

			// begin iterative scheme, F(lambda) = grad(Z) = 0
			// F(lambda) = grad(Z)
			residual.MulVec(xs.T(), phi)
			residual.ScaleVec(-1, residual)
			// grad(F(lambda)) = hessian ( Z )
			for i := 0; i < dim; i++ {
				for j := 0; j < dim; j++ {
					temp1.MulElemVec(xs.ColView(i), xs.ColView(j))
					hF.Set(i, j, mat.Dot(phi, temp1)-residual.AtVec(i)*residual.AtVec(j))
				}

			}

			num_iterations++
			if num_iterations > 100 {
				err1 := errors.New("Newton Raphson iterations failed, max number of iterations exceeded")
				return phi, err1
			}

		}
		// end of dim == 2
	}
	fmt.Printf("******SUCCESS*******\n")
	fmt.Printf("Newton raphson scheme converged in %v iterations \nphi=\n%v\n", num_iterations, mat.Formatted(phi, mat.Prefix(""), mat.Squeeze()))

	return phi, nil

}

func (meshfree *domain.Meshfree) set_basis_function_radii() {

	for i := 0; i < meshfree.num_nodes; i++ {
		if meshfree.isConstantSpacing == true {
			_, max := MinMax(meshfree.nodalSpacing)
			meshfree.basisFunctionRadii[i] = max * meshfree.gamma[i]

		} else {
			meshfree.basisFunctionRadii[i] = meshfree.gamma[i] * meshfree.nodalSpacing[i]
		}
	}
}

// get neighbours of the point p, which was used to construct the shifted_coordinates
func (meshfree *domain.Meshfree) get_neighbours(m *mat.Dense) []int {
	num_r, _ := m.Dims()
	neighbours := make([]int, 0)
	for i := 0; i < num_r; i++ {
		row := m.RowView(i)
		var distance float64 = 0
		// find norm of row of vector
		for index := 0; index < row.Len(); index++ {
			distance += math.Pow(row.AtVec(index), 2)
		}
		distance = math.Sqrt(distance)

		if distance <= meshfree.basisFunctionRadii[i] {
			neighbours = append(neighbours, i)
		}
	}

	return neighbours
}

// Construct the prior, or 'weight functions'
func prior(weight_type string, neighbours []int, xS_c *mat.Dense, meshfree *domain.Meshfree) (*mat.VecDense, *mat.Dense) {

	num_points, dim := xS_c.Dims()
	rNorm := make([]float64, len(neighbours))
	dmI := make([]float64, len(neighbours))
	weight := mat.NewVecDense(num_points, nil)
	weightDer := mat.NewDense(num_points, dim, nil)
	// Find normalised radi rnorm = ri/dmi
	for i := 0; i < len(neighbours); i++ {
		dmI[i] = meshfree.basisFunctionRadii[neighbours[i]]
		rNorm[i] = math.Sqrt(mat.Dot(xS_c.RowView(i), xS_c.RowView(i))) / meshfree.basisFunctionRadii[neighbours[i]]
	}

	if strings.Compare(weight_type, "cubic_spline") == 0 {
		weight, weightDer = cubic_spline(dmI, xS_c, rNorm)

	}
	return weight, weightDer
}

// different weight functions
func cubic_spline(dmI []float64, xS *mat.Dense, rNorm []float64) (*mat.VecDense, *mat.Dense) {

	// initialise weight function and derivative storage
	num_points, dim := xS.Dims()
	weights := mat.NewVecDense(num_points, nil)
	derWeights := mat.NewDense(num_points, dim, nil)
	dwdr := mat.NewDiagonal(num_points, nil)

	// loop over each point in domain of influence of p, and find weight and weight
	// deritavie
	var ri float64 = 0
	// calculate weights and derWeights
	for i := 0; i < num_points; i++ {
		ri = rNorm[i]
		// Obtainign weight functions
		if ri <= 0.5 {
			weights.SetVec(i, 2.00/3.00-4*math.Pow(ri, 2)+4*math.Pow(ri, 3))
			dwdr.SetSymBand(i, i, (8-12*ri)/math.Pow(dmI[i], 2))
		} else {
			weights.SetVec(i, 4.00/3.00-4*ri+4.00*math.Pow(ri, 2)-4.00*math.Pow(ri, 3)/3.00)
			dwdr.SetSymBand(i, i, (4/ri-8+4*ri)/math.Pow(dmI[i], 2))
		}
	}
	derWeights.Mul(dwdr, xS)

	return weights, derWeights

}
func (meshfree *domain.Meshfree) get_nodal_spacing() {
	distanceMin := make([]float64, meshfree.num_nodes)
	for i, nodeI := range meshfree.Nodes {
		for j, nodeJ := range meshfree.Nodes {
			if i != j {
				distanceMin[j] = get_distance(&nodeI.coords, &nodeJ.coords)
			} else {
				distanceMin[j] = 100
			}
		}
		sort.Float64s(distanceMin)
		meshfree.nodalSpacing = append(meshfree.nodalSpacing, distanceMin[meshfree.dim-1])
	}

}

func MinMax(array []float64) (float64, float64) {
	var max float64 = array[0]
	var min float64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
