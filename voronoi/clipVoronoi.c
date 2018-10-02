#include "./C/clipVoronoi.h"




int clipVoronoi(gpc_polygon*** voronoi, double points[], size_t boundary[],int numPoints, int numBoundary)
{

  // intial variables
  FILE * fp;
  char fileName[256];

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
  for (int i = 0; i < numPoints; i++) {
    // write cell to file
    	snprintf(fileName,sizeof(fileName),"./outputs/Cells/%d.txt",i);
    	// write clipped cell to files
    	fp = fopen(fileName,"w");
    	if ( fp != NULL)
    		gpc_write_polygon(fp, 0, (*voronoi)[i]);
    	fclose(fp);

  }


  fp = fopen("cells.txt","w");
	gpc_write_polygon(fp, 0, &clip_polygon);
  fclose(fp);



  return 0;
}
