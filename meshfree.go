/***********************************************************************/
// Example test file for the Mfree go library

// Created by : Stephen P. Smith
// Web        : N/A
// E-mail     : ssmith54@qub.ac.uk
// Version    : 1.0
// Date       : 25/09/2018

// Problem : Tip loaded cantilever beam
// F^
//  | |-------------------|//
//  | |                   |//
//  | |-------------------|//
package main

import (
	"Meshfree/domain"
	"Meshfree/geometry"
	"Meshfree/scni"
	"Meshfree/shapefunctions"
	"fmt"
)

func main() {

	// Generate a domain, from a PLSG file and then generate clipped voronoi
	//var physical_domain domain.Domain
	physical_domain := new(domain.Domain)
	physical_domain.TriGen("models/square", "pDa0.05")
	physical_domain.PrintNodesToImg("nodes")
	physical_domain.GenerateClippedVoronoi()
	physical_domain.GetVoronoi().PrintVoronoiToImg("outputs/cells.eps")
	nodesIn := domain.FindNodesIn(&physical_domain.Nodes, geometry.CreateRectangle(0, 0, .1, .1, 0, 2))

	fmt.Printf("%v\n", nodesIn)

	// create boundary conditions

	// find nodes on traction and displacement boundaries
	// given a boundary, find nodes on boundary that are in a primative shape

	// Meshfree structure

	// refers to the domain
	isConstantSpacing := true
	isVariousPoints := false
	dim := 2
	tol := 1e-8
	meshfree := shapefunctions.NewMeshfree(physical_domain, isConstantSpacing, isVariousPoints, dim, nil, tol)
	meshfree.SetConstantGamma(1.2)
	meshfree.Set_basis_function_radii()
	// p1 := geometry.NewPoint(0, 0, 0)
	// compute := 2
	// // compute shape functions at p
	// meshfree.ComputeMeshfree(&p1, compute, false)

	scni.CreateSCNI(meshfree, physical_domain.GetVoronoi())

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
