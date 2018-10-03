#include "./C/clipVoronoi.h"




int clipVoronoi(double ** output_vertex, int ** output_num_vertex, int * total_vertex, gpc_polygon*** voronoi, double points[], size_t boundary[],int numPoints, int numBoundary)
{


  // create the clipping polygon
  gpc_vertex clipVertex[numBoundary];

  for (int i = 0; i < numBoundary; i++) {
    clipVertex[i].x = points[2*boundary[i]];
    clipVertex[i].y = points[2*boundary[i]+1];
  }

  gpc_vertex_list clip_list = {.num_vertices = numBoundary, .vertex = clipVertex};
  gpc_polygon clip_polygon = {.num_contours=1, .hole = NULL, .contour = &clip_list};


  printf("CLIPPING VORONOI\n");

  // clip boundary cells
  for (size_t i = 0; i < numBoundary; i++) {

    int indx = boundary[i];

    gpc_polygon * clipped_polygon = malloc(1*sizeof(gpc_polygon));

    gpc_polygon_clip(GPC_INT,&clip_polygon,(*voronoi)[indx],clipped_polygon);

    gpc_free_polygon((*voronoi)[indx]);
    (*voronoi)[indx] = clipped_polygon;
    /* code */
  }

  printf("FINISHED CLIPPING VORONOI\n");

  // write cells to file

  //FILE * fp;
  //char fileName[256;]
  // for (int i = 0; i < numPoints; i++) {
  //   // write cell to file
  //   	snprintf(fileName,sizeof(fileName),"./outputs/Cells/%d.txt",i);
  //   	// write clipped cell to files
  //   	fp = fopen(fileName,"w");
  //   	if ( fp != NULL)
  //   		gpc_write_polygon(fp, 0, (*voronoi)[i]);
  //   	fclose(fp);
  //
  // }
  //
  //
  // fp = fopen("cells.txt","w");
	// gpc_write_polygon(fp, 0, &clip_polygon);
  // fclose(fp);

  // put voronoi into a vertex point list and
  //create an array to store how many verticies each polygon has

  int * list_num_vertices = malloc(numPoints * sizeof(int));
  int sum_vertices = 0;
  for (size_t i = 0; i < numPoints; i++) {

    list_num_vertices[i] = (*voronoi)[i]->contour->num_vertices;

    /* codei */
    sum_vertices += list_num_vertices[i];

  }

  double * vertex_list = malloc (2* sum_vertices * sizeof(double));
  int count = 0;
  for (size_t i = 0; i < numPoints; i++) {

    for (size_t k = 0; k < list_num_vertices[i]; k++) {
      vertex_list[2*count] = (*voronoi)[i]->contour->vertex[k].x;
      vertex_list[2*count+1] = (*voronoi)[i]->contour->vertex[k].y;
      count = count + 1;
    }

    /* code */
  }

  *output_vertex = vertex_list;
  *output_num_vertex = list_num_vertices;
  *total_vertex = sum_vertices;
  // clear voronoi memory



  for (size_t i = 0; i < numPoints; i++) {
    /* code */
    gpc_free_polygon((*voronoi)[i]);
  }

  return 0;
}
