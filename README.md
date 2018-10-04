<h1> Mfree: A meshfree library written in go-lang </h1>

This repository contains the code for an open source go lang that implementents a variant on the element-free Galerkin method. The current release of this library allows for the simulation of nonlinear 2D and axisymmetric problems using an explicit solver.

Features:
-
-
-
-
-

<h2> Usage: </h2>
Edit meshfree.go:
- Create a domain using a PLSG, examples are in the models folder
- Create meshfree domain
- Create voronoi diagram
- Create SCNI cells
- Create boundary conditions
- Choose material
- Set up time stepping
- Choose outputs to save
- From root
-- go build -a
-- ./meshfree