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

	// create a coordinate system
	globalCS := coordinatesystem.CreateCartesian()
	// set dimension of the problem
	DIM := 2
	//Generate a domain, from a PLSG file
	modelname := "Beam"
	PLSG_file := "models/square"
	mesh_options := "pDa0.05"
	// creates the points and the voronoi diagram
	domain_ := domain.NewDomain(modelname, PLSG_file, mesh_options, DIM, globalCS)

	// modify dof for essential boundaries
	nodesEB1 := node.FindNodesIn(&domain_.Nodes, geometry.CreateRectangle(0, 2, 0, 0, 0.1, 0))
	node.PrintNodes(nodesEB1)
	node.SetNodalDofs(nodesEB1, globalCS.X, dof.DOF_FIXED, domain_.GetDim())
	node.SetNodalDofs(nodesEB1, globalCS.Y, dof.DOF_FIXED, domain_.GetDim())
	domain_.GetNodesIn(geometry.CreateRectangle(0, 2, 0, 0, 0.1, 0))

	// add material to the domain
	domain_.Add_material()

	// find nodes on right side of beam (traction)

	// Meshfree structure
	// refers to the domain
	isConstantSpacing := true
	isVariousPoints := false
	dim := 2
	tol := 1e-8
	meshfree := shapefunctions.NewMeshfree(domain_, isConstantSpacing, isVariousPoints, dim, nil, tol)
	meshfree.SetConstantGamma(1.2)
	meshfree.Set_basis_function_radii()
	// p1 := geometry.NewPoint(0, 0, 0)
	// compute := 2
	// // compute shape functions at p
	// meshfree.ComputeMeshfree(&p1, compute, false)

	scni.CreateSCNI(meshfree, domain_.GetVoronoi())

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
