/***********************************************************************/
// Example test file for the Mfree go library

// Created by : Stephen P. Smith
// Web        : N/A
// E-mail     : ssmith54@qub.ac.uk
// Version    : 1.0
// Date       : 25/09/2018

package main

import (
	"Meshfree/domain"
	"Meshfree/geometry"
	"Meshfree/shapefunctions"
)

func main() {

	// Generate a domain, from a PLSG file and then generate clipped voronoi
	var domain domain.Domain
	domain.TriGen("square", "pDa0.05")
	domain.PrintNodesToImg("nodes")
	domain.GenerateClippedVoronoi()

	// Meshfree structure
	isConstantSpacing := true
	isVariousPoints := false
	dim := 2
	meshfree := shapefunctions.NewMeshfree(&domain, isConstantSpacing, isVariousPoints, dim, nil)
	meshfree.SetConstantGamma(1.2)
	meshfree.Set_basis_function_radii()
	p1 := geometry.NewPoint(0, 0, 0)
	tol := 1e-8
	compute := 2

	// compute shape functions at p
	meshfree.ComputeMeshfree(&p1, dim, compute, tol)

}
