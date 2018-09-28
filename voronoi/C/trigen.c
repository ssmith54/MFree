
#include <trigen.h>
#include <connect_segments.h>

void report(io, markers, reporttriangles, reportneighbors, reportsegments,
		reportedges, reportnorms)
	struct triangulateio *io;
	int markers;
	int reporttriangles;
	int reportneighbors;
	int reportsegments;
	int reportedges;
	int reportnorms;
{
	int i, j;

	for (i = 0; i < io->numberofpoints; i++) {
		printf("Point %4d:", i);
		for (j = 0; j < 2; j++) {
			printf("  %.6g", io->pointlist[i * 2 + j]);
		}
		if (io->numberofpointattributes > 0) {
			printf("   attributes");
		}
		for (j = 0; j < io->numberofpointattributes; j++) {
			printf("  %.6g",
					io->pointattributelist[i * io->numberofpointattributes + j]);
		}
		printf("\n");
	}


	if (reporttriangles || reportneighbors) {
		for (i = 0; i < io->numberoftriangles; i++) {
			if (reporttriangles) {
				printf("Triangle %4d points:", i);
				for (j = 0; j < io->numberofcorners; j++) {
					printf("  %4d", io->trianglelist[i * io->numberofcorners + j]);
				}
				if (io->numberoftriangleattributes > 0) {
					printf("   attributes");
				}
				for (j = 0; j < io->numberoftriangleattributes; j++) {
					printf("  %.6g", io->triangleattributelist[i *
							io->numberoftriangleattributes + j]);
				}
				printf("\n");
			}
			if (reportneighbors) {
				printf("Triangle %4d neighbors:", i);
				for (j = 0; j < 3; j++) {
					printf("  %4d", io->neighborlist[i * 3 + j]);
				}
				printf("\n");
			}
		}
		printf("\n");
	}

	if (reportsegments) {
		for (i = 0; i < io->numberofsegments; i++) {
			printf("Segment %4d points:", i);
			for (j = 0; j < 2; j++) {
				printf("  %4d", io->segmentlist[i * 2 + j]);
			}
			if (markers) {
				printf("   marker %d\n", io->segmentmarkerlist[i]);
			} else {
				printf("\n");
			}
		}
		printf("\n");
	}

	if (reportedges) {
		for (i = 0; i < io->numberofedges; i++) {
			printf("Edge %4d points:", i);
			for (j = 0; j < 2; j++) {
				printf("  %4d", io->edgelist[i * 2 + j]);
			}
			if (reportnorms && (io->edgelist[i * 2 + 1] == -1)) {
				for (j = 0; j < 2; j++) {
					printf("  %.6g", io->normlist[i * 2 + j]);
				}
			}
			if (markers) {
				printf("   marker %d\n", io->edgemarkerlist[i]);
			} else {
				printf("\n");
			}
		}
		printf("\n");
	}
}

int trigen(gpc_polygon *** voronoi_out, jcv_point ** points_out, int ** boundary, char * options, char * fileName, int * numPoints ){

	char nodesFile[20] ;
	char segsFile[20] ;

	strcpy(nodesFile,fileName);
	strcpy(segsFile,fileName);
	strcat(nodesFile,".nodes");
	strcat(segsFile,".segs");

	printf("%s\n",nodesFile );



	tri in = malloc(sizeof(struct triangulateio ));
	tri out = malloc(sizeof(struct triangulateio));
		/*  Input structure */
	char buf[64];
	char s[2] = " ";

	char * token;

	FILE * fp = fopen(nodesFile,"r");
	fgets(buf,sizeof(buf),fp);
	token = strtok(buf,s);
	in->numberofpoints = atoi(token);
	token = strtok(NULL,s);
	printf("got here (a)\n");
	int isAttribute = 0;
	if ( token != NULL)
	{
		isAttribute = 1;
	}
	printf("got here (b) \n");
	if ( isAttribute == 1){
		in->numberofpointattributes = 1;
		in->pointattributelist = (double* )  malloc(in->numberofpoints*sizeof(double)); 
	}else{
		in->numberofpointattributes = 0;
		in->pointattributelist = NULL; 
	}
	in->pointlist = malloc(in->numberofpoints*2*sizeof(double));
	printf("got here (c) \n");
	for ( int i = 0 ; i < in->numberofpoints ;i++){
		fgets(buf,sizeof(buf),fp);
		token = strtok(buf,s);
		in->pointlist[2*i] = atof(token);
		token = strtok(NULL,s);
		in->pointlist[2*i+1] = atof(token);

		if ( isAttribute == 1){
			token = strtok(NULL,s);
			in->pointattributelist[i] = atof(token);
		}

	}
	fclose(fp);
	fp = fopen(segsFile,"r");
	fgets(buf,sizeof(buf),fp);
	in->numberofsegments = atoi(buf);
	in->segmentlist = malloc(2*in->numberofsegments*sizeof(int));
	in->segmentmarkerlist = malloc(in->numberofsegments*sizeof(int));
	for ( int i = 0 ; i < in->numberofsegments ;i++){
		fgets(buf,sizeof(buf),fp);
		token = strtok(buf,s);
		in->segmentlist[2*i] = atoi(token);
		token = strtok(NULL,s);
		in->segmentlist[2*i+1] = atoi(token);
		token = strtok(NULL,s);
		if ( token != NULL)
			in->segmentmarkerlist[i] = atoi(token);

	}
	in->numberofregions = 0;
	in->regionlist = (double *) NULL;
		//holes 
	in->numberofholes = 0;
	in->holelist = (double *) NULL; 
	
	// set of options inserted 
	printf("Input point set:\n");
	report(in, 1, 0, 0, 1, 0, 0);

	/* Make necessary initializations so that Triangle can return a */
	/*   triangulation in `out' and a voronoi diagram in `vorout'.  */

	out->pointlist = (double *) NULL;            /* Not needed if -N switch used. */
	/* Not needed if -N switch used or number of point attributes is zero: */
	out->pointattributelist = (double *) NULL;
	out->pointmarkerlist = (int *) NULL; /* Not needed if -N or -B switch used. */
	out->trianglelist = (int *) NULL;          /* Not needed if -E switch used. */
	/* Not needed if -E switch used or number of triangle attributes is zero: */
	out->triangleattributelist = (double *) NULL;
	out->neighborlist = (int *) NULL;         /* Needed only if -n switch used. */
	/* Needed only if segments are output (-p or -c) and -P not used: */
	out->segmentlist = (int *) NULL;
	/* Needed only if segments are output (-p or -c) and -P and -B not used: */
	out->segmentmarkerlist = (int *) NULL;
	out->edgelist = (int *) NULL;             /* Needed only if -e switch used. */
	out->edgemarkerlist = (int *) NULL;   /* Needed if -e used and -B not used. */
	out->holelist = (double *) NULL;
	out->regionlist = (double * ) NULL; 
	/* Triangulate the points.  Switches are chosen to read and write a  */
	/*   PSLG (p), preserve the convex hull (c), number everything from  */
	/*   zero (z), assign a regional attribute to each element (A), and  */
	/*   produce an edge list (e), a Voronoi diagram (v), and a triangle */
	/*   neighbor list (n).                                              */

	printf("passed triangulate \n");
	/*  Failed at triangulate */
	triangulate(options,in,out, (struct trianglulateio *) NULL);

	printf("passed triangulate \n");
	report(out, 1, 0, 0, 0, 0, 0);


	// fill the jcv pointer
	jcv_point * points = malloc(out->numberofpoints * sizeof(jcv_point));

	for (int i = 0; i < out->numberofsegments; ++i)
	{
		printf("Segment: %d %d\n", out->segmentlist[2*i], out->segmentlist[2*i+1] );
				printf("segment marker: %d \n", out->segmentlist[i]);

		/* code */
	}

	// find boundary 
	int * connected_segments = NULL;
	connect_segments(out->segmentlist, &connected_segments, out->numberofsegments);
	printf("got here (b)\n");	
	for (int i = 0; i < out->numberofsegments -1 ; ++i)
	{
		printf("Segment: %d %d\n", connected_segments[2*i], connected_segments[2*i+1] );
				/* code */
	}

	// write into a polygon
	gpc_polygon * outputs = malloc(1*sizeof(gpc_polygon));
	outputs->hole = NULL;
	outputs->num_contours = 1;
	outputs->contour = malloc(1*sizeof(gpc_vertex_list));
	outputs->contour->num_vertices = (out->numberofsegments -1);
	outputs->contour->vertex = malloc(outputs->contour->num_vertices*sizeof(gpc_vertex));

	for ( int i = 0 ; i <  outputs->contour->num_vertices ; ++i )
	{
		outputs->contour->vertex[i].x = out->pointlist[2*(connected_segments[2*i]-1)];
		outputs->contour->vertex[i].y = out->pointlist[2*(connected_segments[2*i]-1) + 1];

	}
    

    // create a jc voronoi structure
	jcv_diagram diagram;
	jcv_graphedge * graph_edge;
	memset(&diagram, 0, sizeof(jcv_diagram));
	


	// generate voronoi diagram
	for (int i = 0; i < out->numberofpoints; ++i)
	{
		points[i].x = out->pointlist[2*i];
		points[i].y = out->pointlist[2*i+1];

		/* code */
	}

	// GENERATE VORONOI DIAGRAM
	printf("Generating voronoi diagram \n");
	jcv_diagram_generate(out->numberofpoints,(const jcv_point *) points, NULL, &diagram);

	//VORONOI DIAGRAM //

	gpc_polygon **voronoi = malloc(diagram.numsites*sizeof(gpc_polygon *));

	for (int i = 0; i < diagram.numsites; ++i)
	{
		// Initialise polgon
		voronoi[i] = malloc(1*sizeof(gpc_polygon));
		voronoi[i]->hole = NULL;
		voronoi[i]->num_contours = 1;
		voronoi[i]->contour = malloc(1*sizeof(gpc_vertex_list));
		voronoi[i]->contour->num_vertices = 0;
		voronoi[i]->contour->vertex = malloc(voronoi[i]->contour->num_vertices*sizeof(gpc_vertex)); 

	}



	// LOOP OVER EACH SITE
	char fileName1[256];
 	const jcv_site* sites;
  	sites = jcv_diagram_get_sites(&diagram);
    int vertCount = 0;

  	// start of loop
  	printf("Clipping edges of cells\n");
	for (int i=0; i<diagram.numsites; ++i) {

	vertCount = 0;	
    graph_edge = sites[i].edges;
    voronoi[i]->index = sites[i].index;

    if ( out->pointattributelist == NULL)
    	voronoi[i]->attribute = out->pointattributelist[voronoi[i]->index]; 
    

    // FIRST GET NUMBER OF EDGES IN THE CELL

    while (graph_edge) {
    	++vertCount;
    	graph_edge = graph_edge->next;

    }// end of while loop over edges

    // create memory for the voronoi cell
    voronoi[i]->contour->num_vertices = vertCount;
	voronoi[i]->contour->vertex = malloc(voronoi[i]->contour->num_vertices*sizeof(gpc_vertex)); 

	// loop of the edges of the site
    graph_edge = sites[i].edges;
    for (int k = 0; k < vertCount; ++k)
    {
    	// loop over each cell and put it into the voronoi polygon structure
    	voronoi[i]->contour->vertex[k].x = graph_edge->pos[0].x;
    	voronoi[i]->contour->vertex[k].y = graph_edge->pos[0].y;
    	graph_edge = graph_edge->next;
    }// end of loop over edges


	// clip the cell
	gpc_polygon * result1 = malloc(1*sizeof(gpc_polygon));

	gpc_polygon_clip(GPC_INT, outputs, voronoi[i], result1);

	gpc_free_polygon(voronoi[i]);
	voronoi[i] = result1;


	// write cell to file
	snprintf(fileName1,sizeof(fileName1),"./Cells/%d.txt",i);
	// write clipped cell to files
	fp = fopen(fileName1,"w");
	if ( fp != NULL)
		gpc_write_polygon(fp, 0, voronoi[i]);
	fclose(fp);

	}// end of loop of sites

	printf("success..\n");

	fp = fopen("cells.txt","w");
	gpc_write_polygon(fp, 0, outputs);
	// in point list, 
	FILE * ptr;
	ptr = fopen("mesh.1.ele","w");
	fprintf(ptr,"%i %i %i \n",out->numberoftriangles,3,0 );
	for(int i = 0 ; i < out->numberoftriangles ; i ++){
		fprintf(ptr,"%i %i %i %i\n",i+1,out->trianglelist[3*i],out->trianglelist[3*i +1 ], out->trianglelist[3*i+2]);	
	}
	fclose(ptr);


	ptr = fopen("mesh.1.node","w");
	fprintf(ptr,"%i %i %i %i\n",out->numberofpoints,2,0,0 );
	for (int i = 0 ; i < out->numberofpoints ; i++){
		fprintf(ptr,"%i %lf %lf \n",i+1,out->pointlist[2*i],out->pointlist[2*i+1]);

	}
	fclose(ptr);	



	// set outputs
	*voronoi_out = voronoi;
	*points_out = points;
	*numPoints = out->numberofpoints;
	*boundary = connected_segments;
	// need to free the input and output memory of tri; the only thing that needs to come out
	// is jcv_points and 




	free(in->pointlist);
 	free(in->pointattributelist);
 	free(in->pointmarkerlist);
  	free(in->segmentmarkerlist);
  	free(in->segmentlist);
  	free(in->regionlist);
  	free(in);


	free(out->pointlist);
	free(out->pointattributelist);
	free(out->trianglelist);
	free(out->triangleattributelist);
	free(out);





	return 0;


}