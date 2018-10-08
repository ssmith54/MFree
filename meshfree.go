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
	"Meshfree/coordinatesystem"
	"Meshfree/dof"
	"Meshfree/domain"
	"Meshfree/geometry"
	"Meshfree/node"
	"Meshfree/scni"
	"Meshfree/shapefunctions"
)

func main() {

	// create a coordinate coordinate coordinate system
	globalCS := coordinatesystem.CreateCartesian()
	// Generate a domain, from a PLSG file and then generate clipped voronoi
	//var physical_domain domain.Domain
	physical_domain := new(domain.Domain)
	physical_domain.TriGen("models/square", "pDa0.05")
	physical_domain.PrintNodesToImg("nodes")
	physical_domain.GenerateClippedVoronoi()
	physical_domain.GetVoronoi().PrintVoronoiToImg("outputs/cells.eps")

	// create BC dof and then the rest
	// find nodes on left side of beam (boundary conditions)
	nodesEB1 := node.FindNodesIn(&physical_domain.Nodes, geometry.CreateRectangle(0, 2, 0, 0, 0.1, 0))
	node.PrintNodes(nodesEB1)
	// create degrees of freedom pertaining to the essential boundary nodes
	overrideDOFs := true // whether to override preexisting dofs that match the node and direction
	node.CreateNodalDofs(nodesEB1, globalCS.X, dof.DOF_FIXED, overrideDOFs)

	// find nodes on right side of beam (traction)

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
