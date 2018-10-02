#ifndef CLIPVORONOI_H_
#define CLIPVORONOI_H_

#include "./gpc/gpc.h"
#include "./jc_voronoi/jc_voronoi.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>




int clipVoronoi(gpc_polygon *** voronoi, double points[], size_t boundary[],int numPoints, int numBoundary);










#endif
