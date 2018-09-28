
/*****************************************************************************/
/*                                                                           */
/*  (trigen.c)                                                              */
/*                                                                           */
/*  Program that calls triangle              */
/*                                                                           */                                                  
/*                                                                           */
/*  Takes in a boundary profile 
	along with meshing options                             */
/*                                                                           */
/*****************************************************************************/
#ifndef TRIGEN_H_
#define TRIGEN_H_

#include "triangle.h"
#include <stdlib.h>
#include <stdio.h>
#include <gpc.h>
#include <string.h>
#include <jc_voronoi.h>


typedef struct triangulateio * tri ;



int trigen(gpc_polygon *** voronoi, jcv_point ** points, int ** boundary, char * options, char * fileName, int * numPoints );



#endif // end def TRIGEN_H_