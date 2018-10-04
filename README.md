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
<ul>
<li> Create a domain using a PLSG, examples are in the models folder </li>
<li> Create meshfree domain</li>
<li> Create voronoi diagram</li>
<li> Create SCNI cells</li>
<li> Create boundary conditions</li>
<li> Choose material</li>
<li> Set up time stepping</li>
<li> Choose outputs to save</li>
<li> From root</li>
</ul>

<h4> Finally in the root folder</h4>
<ul>
<li> go build -a</li>
<li> ./meshfree </li>
</ul>
