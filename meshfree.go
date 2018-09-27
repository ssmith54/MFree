/***********************************************************************/
// This program computes the Maximum Entropy and MLS shape shape functions
// in 1D, 2D and 3D

// Created by : Stephen P. Smith
// Web        : N/A
// E-mail     : ssmith54@qub.ac.uk
// Version    : 1.0
// Date       : 25/09/2018

//
//				Department of Mechanical and Aerospace Engineering
//									Queen's Unversity Belfast
//

// Based on the implementation by Alejandro A. Ortiz (www.cec.uchile.cl/~aortizb)

// __________
// References
// __________

// Sukumar N. Construction of polygonal interpolants: a maximum entropy
// approach. Int. J. Numer. Meth. Engng 2004; 61(12):2159-2181.

package main

import "Meshfree/domain"

func main() {

	// Construct meshfree shape functions
	nodes := make([]domain.Node, 0)
	// define the nodes
	n1 := domain.NewNode(0, 0, 0, 1)
	n2 := domain.NewNode(1, 0, 0, 2)
	n3 := domain.NewNode(1, 1, 0, 3)
	n4 := domain.NewNode(0, 1, 0, 4)
	n5 := domain.NewNode(0.5, 0.5, 0, 5)
	n6 := domain.NewNode(0.25, 0.25, 0, 6)
	n7 := domain.NewNode(0.25, 0.75, 0, 7)
	n8 := domain.NewNode(0.75, 0.25, 0, 8)
	n9 := domain.NewNode(0.75, 0.75, 0, 9)
	// add nodes to meshfree domain
	nodes = append(nodes, *n1, *n2, *n3, *n4, *n5, *n6, *n7, *n8, *n9)
	numnodes := int(9)
	// set up Domain of nodes
	domain := domain.NewDomain(nodes, numnodes)

	// Meshfree structure
	meshfree.Add_nodes(n1, n2, n3, n4, n5, n6, n7, n8, n9)

	// set number of nodes
	meshfree.num_nodes = 9

	// Meshfree stuff

	// // constant domain domainSize
	// meshfree.isConstantSpacing = true
	// // Find shape function at various points?
	// meshfree.isVariousPoints = false
	// // nodal support size paramter gamma
	// meshfree.gamma = []float64{1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2}
	// meshfree.dim = 2
	// // get nodal spacing
	// meshfree.get_nodal_spacing()
	// // set basis function radii
	// meshfree.basisFunctionRadii = make([]float64, meshfree.num_nodes)
	// meshfree.set_basis_function_radii()
	// // find meshfree shape functions at point p1
	// p1 := Point{0.25, 0.25, 0}
	// tol := 1e-8
	// dim := 2
	// compute := 2
	//
	// // should return phi, phiDer, and the neighbours
	// meshfree.compute_meshfree(&p1, dim, compute, tol)

}
