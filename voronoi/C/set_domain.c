#include <set_domain.h>
#include "structure.h"

static inline double distance(jcv_point * a, jcv_point * b)
{
	return sqrt(pow(b->x - a->x,2) + pow(b->y-a->y,2));
}


int set_domain(Meshfree * meshfree, boolean constant_domain_size)
{

	// set up domain size array
	meshfree->di = malloc(meshfree->num_nodes*sizeof(double));

	// compare distance of each meshfree node with eachother
	jcv_point * n1; // node 1
	jcv_point * n2; // node 2
	int numnodes = meshfree->num_nodes;
	// 
	double distanceTemp = 0;
	double maxNodalSpacing = 0;
	double minDistance = 1e5;


	if ( constant_domain_size == true)
	// Find maximum nodal spacing, and set di as that
	{
		for (int i = 0; i < numnodes; ++i)
		{
			distanceTemp = 0;
			n1 = &meshfree->nodes[i].coords;
			for (int k = 0; k < numnodes; ++k)
			{
				if ( i != k){
					n2 = &meshfree->nodes[k].coords;

					// Find distance between node 1 and node 2
					distanceTemp = distance(n1, n2);
					// Check if
					if ( distanceTemp < minDistance )
					{
						minDistance = distanceTemp;
					}

				}// end of if
			}// end of loop over nodes
			
			if ( minDistance > maxNodalSpacing)
			{
				maxNodalSpacing = minDistance;
			}
		}

	}else
	{

	}

	for (int i = 0; i < numnodes; ++i)
	{
		meshfree->di[i] = maxNodalSpacing; 
	}


	return 0;
}