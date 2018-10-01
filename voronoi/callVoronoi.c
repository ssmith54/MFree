
#include "C/callVoronoi.h"


int callVoronoi(double points[], int num_points){

  // put points into jcv points structure and then call voronoi
  jcv_point voronoi_points[num_points] ;
  printf("number of points = %d\n ", num_points);

  for (int i = 0; i < num_points; i++) {
    voronoi_points[i].x = points[2*i];
    voronoi_points[i].y = points[2*i+1];
  }


  // allocate memory for voronoi diagram
  jcv_diagram diagram;
  jcv_graphedge * graph_edge;
  memset(&diagram,0,sizeof(jcv_diagram));

  printf("GENERAING VORONOI DIAGRAM \n");
  jcv_diagram_generate(num_points,voronoi_points,NULL,&diagram);
  printf("FINISHED GENERATING VORONOI DIAGRAM \n");


  // move voronoi into gpc_polygon structure, where each polygon is stored
  // as a group of verticies
  gpc_polygon ** voronoi = malloc(diagram.numsites*sizeof(gpc_polygon*));


  for (size_t i = 0; i < diagram.numsites; i++) {
    voronoi[i] = malloc(1*sizeof(gpc_polygon));
    voronoi[i]->hole = NULL;
    voronoi[i]->num_contours = 1;
    voronoi[i]->contour = malloc(1*sizeof(gpc_vertex_list));
    voronoi[i]->contour->num_vertices = 0;
    voronoi[i]->contour->vertex = malloc(voronoi[i]->contour->num_vertices*sizeof(gpc_vertex)); 

  }
  // initialise each polygon



  return 0;


}
