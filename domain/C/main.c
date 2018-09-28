#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <gpc.h>
#include <memory.h>
#define JC_VORONOI_IMPLEMENTATION
#include "jc_voronoi.h"
#include "trigen.h"
#include <omp.h>
#include "structure.h"
#include <sys/time.h>
#include "mkl.h"
#include <set_domain.h>




static inline void printOut(char * a)
{
	printf("      ---------------------------------------------------------\n");
	printf("---------------          %s             -----------------\n", a );
	printf("      ---------------------------------------------------------\n");

}




int main(int argc, char const *argv[])
{

	(void)argc;
	(void)argv;

	// create a jc voronoi structure
	FILE * fp;

	// file name
	char * filename = "preform";


	// generate clipped voronoi diagram
	gpc_polygon ** voronoi = NULL;
	jcv_point * tri_points = NULL;
	int * boundary = NULL;
	int numPoints = 0;
	trigen(&voronoi, &tri_points, &boundary, "pDq0a0.5", filename, &numPoints);


	// set nodes
	Node * nodes = malloc(numPoints*sizeof(Node));
	for (int i = 0; i < numPoints; ++i)
	{
		nodes[i].coords.x = tri_points[i].x;
		nodes[i].coords.y = tri_points[i].y;

		/* code */
	}
	Meshfree * meshfree = malloc(1*sizeof(Meshfree));
	meshfree->num_nodes = numPoints;
	meshfree->nodes = nodes;

	//  set up domain size
	// only supports radial domains
	boolean isConstantSpacing = 1;
	set_domain(meshfree,isConstantSpacing);
	printOut("setting domain");


	// SET UP SCNI CELLS, SO FAR JUST NEED VORONOI?
	// loop over and find each area?
	double * area_list = malloc(numPoints * sizeof(double));
	double area_temp = 0 ;
	// looping over each cell




	gpc_vertex a, b, c;
	int vectorFlip = 1;
	for ( int i = 0 ; i <  numPoints ; ++i)
	{
		// polygon of interest
		gpc_polygon * poly_i = voronoi[i];
		area_temp = 0;


		// check orientation of this polygon,
		// take points A-B-C
		a = poly_i->contour->vertex[0];
		b = poly_i->contour->vertex[1];
		c = poly_i->contour->vertex[2];

		double detO = b.x*c.y - b.y*c.x - a.x*c.y +
		 a.x*b.y + a.y*c.x - a.y*b.x;

		if ( detO < 0)
		{
			// Normal vector will need to be flipped, i.e vectorFlip = -1
			vectorFlip = -1;
		}else
		{
			// Vector does not need to be flipped i.e vectorFlip = 1
			vectorFlip = 1;
		}

		for ( int k = 0 ; k < poly_i->contour->num_vertices ; ++k)
		{
			if ( k == poly_i->contour->num_vertices - 1)
			{


				// shape functions

			// find length of segment
			area_temp += 0.5*(poly_i->contour->vertex[k].x*poly_i->contour->vertex[0].y -
				poly_i->contour->vertex[0].x*poly_i->contour->vertex[k].y);
			}else
			{
				area_temp += 0.5*(poly_i->contour->vertex[k].x*poly_i->contour->vertex[k+1].y -
				poly_i->contour->vertex[k+1].x*poly_i->contour->vertex[k].y);
			}

		}
		area_list[i] = fabs(area_temp);

	}

	double area_sum = 0;
	for (int i = 0; i < numPoints; ++i)
	{
		area_sum += area_list[i];
	}
	printf("Area = %lf \n", area_sum );


	// set up stabalising cells.
	SCNI ** scni = malloc(numPoints*sizeof(scni));

	for (int i = 0; i < numPoints; ++i)
	{
		scni[i] = malloc(1*sizeof(scni));
		scni[i]->poly = voronoi[i];
	}

	// generate B, the strain displacement matrix




	// test an open mp routine
	int nthreads, tid;
	#pragma omp parallel private(nthreads, tid)
	{

	 /* Obtain thread number */
		tid = omp_get_thread_num();
		printf("Hello World from thread = %d\n", tid);

 	/* Only master thread does this */
		if (tid == 0)
		{
			nthreads = omp_get_num_threads();
			printf("Number of threads = %d\n", nthreads);
		}

	 } /* All threads join master thread and disband */


	// timer
	struct timeval tv1, tv2;
	gettimeofday(&tv1,NULL);

	// time code

	gettimeofday(&tv2,NULL);

	printf ("Total time = %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec) / 1000000 +
         (double) (tv2.tv_sec - tv1.tv_sec));

	return 0;

}
