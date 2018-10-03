#ifndef CLIPVORONOI_H_
#define CLIPVORONOI_H_

#include "./gpc/gpc.h"
#include "./jc_voronoi/jc_voronoi.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>




int clipVoronoi(double ** output_vertex, int ** output_num_vertex, int * total_vertex,  gpc_polygon*** voronoi, double points[], size_t boundary[],int numPoints, int numBoundary);









#endif
