package material

import "gonum.org/v1/gonum/mat"

type Material interface {
	GetStress() *mat.Dense
}

type Linear_Elastic struct {
	E  float64
	mu float64
}

func New_linear_elastic(E_in, mu_in float64) *Linear_Elastic {
	return &(Linear_Elastic{E: E_in, mu: mu_in})

}

type St_Venant_Kirchoff struct {
	lambda float64
	mu     float64
}
