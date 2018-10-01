#ifndef CALLVORONOI_H_
#define CALLVORONOI_H_

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#define JC_VORONOI_IMPLEMENTATION
#include "./jc_voronoi/jc_voronoi.h"
#include "./gpc/gpc.h"

int callVoronoi(gpc_polygon *** voronoi_out, double points[], int num_points);






#endif
