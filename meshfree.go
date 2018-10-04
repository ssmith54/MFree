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
	"Meshfree/scni"
	"Meshfree/shapefunctions"
)

func main() {

	// Generate a domain, from a PLSG file and then generate clipped voronoi
	var domain domain.Domain
	domain.TriGen("models/square", "pDa0.01")
	domain.PrintNodesToImg("nodes")
	domain.GenerateClippedVoronoi()
	domain.GetVoronoi().PrintVoronoiToImg("outputs/cells.eps")

	// Meshfree structure

	// refers to the domain
	isConstantSpacing := true
	isVariousPoints := false
	dim := 2
	tol := 1e-8
	meshfree := shapefunctions.NewMeshfree(&domain, isConstantSpacing, isVariousPoints, dim, nil, tol)
	meshfree.SetConstantGamma(1.2)
	meshfree.Set_basis_function_radii()
	p1 := geometry.NewPoint(0, 0, 0)
	compute := 2

	// compute shape functions at p
	meshfree.ComputeMeshfree(&p1, compute, true)

	scni.CreateSCNI(meshfree, domain.GetVoronoi())

	// set up stabalised conforming nodal integration SCNI
	// each cell just has a Bmatrix (axi, Pstress,Pstrain, 3D), and a volume. Simple?

	// set up material

	// set up boundary conditions

	// store point marker list in domain

	// set up timestepping

	// do time steps
	//-o-// Make a time step
	//-o-// Find Fint, Fcontat, Fext
	//-o-// Balance forces
	//-o-// Energy balance

	// postprocess

}
