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
)

func main() {

	// Generate the domain
	//var vor voronoi.Voronoi

	// read in nodes then generate a voronoi diagram

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
	// Add nodes and set number of nodes in domain
	domain := domain.NewDomain(nodes, numnodes)

	domain.TriGen("square", "pDa0.05")
	domain.PrintNodesToImg("nodes")
	domain.GenerateClippedVoronoi()

	// // Meshfree structure
	// isConstantSpacing := true
	// isVariousPoints := false
	// gamma := []float64{1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2}
	// dim := 2
	// meshfree := shapefunctions.NewMeshfree(&domain, isConstantSpacing, isVariousPoints, dim, gamma)
	//
	// p1 := geometry.NewPoint(0.25, 0.25, 0)
	// tol := 1e-8
	// compute := 2
	//
	// // compute shape functions at p
	// meshfree.ComputeMeshfree(&p1, dim, compute, tol)

}
