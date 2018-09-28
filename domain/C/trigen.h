
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
#include <string.h>

#define REAL double

typedef struct triangulateio * tri ;

int trigen(double **points, int ** boundary, char * options, char * fileName, int * numPoints );



#endif // end def TRIGEN_H_
