#ifndef structures_EFG_H
#define structures_EFG_H
#include <matrix.h>
#include <gpc.h>
#include "jc_voronoi.h"


typedef struct EFG_SF {

    int num_neighbours;
    IVEC * index ; 
    VEC * phi;
    VEC * dphidx;
    VEC * dphidy;
    VEC * phiNb;
    VEC * phiB;
    IVEC * indexB;
    IVEC * indexNb;
    VEC * Wb;
    VEC * WNb;
    
} EFG_SF; 







typedef struct efg_block{
    int numnode;
    MAT * nodes;
    VEC * di; 
    VEC * diX;
    VEC * diY;
    VEC * diZ;
    IVEC * boundaryNodes; 
    char type[20];
    VEC * dMax;
    IVEC * tNodes;
	struct EBC * eb1;
	struct EBC * eb2;
	struct EBC * eb3;
	struct _boundaryLoad * tL;
} efg_block; 

typedef struct _boundaryLoad {
	IVEC * nodes;
	MAT * coords;
	char type ; /*  p = presure, t = traction, f = follower */
	char integrationType ; /* t = trapazoidal rule, g = gauss integration */
	char formulation ; /*  l = Lagrangian, e = Eulerian */
	struct _gauss1D * gauss;
	struct _trap_block * trapz;
	VEC * force;
	double pressure; /*  Scalar pressure */
	double tMag ; /*  Scalar traction ( MAGNITUDE OF APPLIED TRACTION) */
	int numSegs;
	double alpha;
	MAT * rotation;

	

}boundaryLoad;

typedef struct _trap_block{
	int n1,n2 ; /*  Nodes */
	double normal[2] ;
	double length;
	struct EFG_SF * _phi;

}trap_block;


typedef struct _gauss1D{
	int n1,n2; /*  Nodes */
	int order; /*  Gauss order */
	MAT * gPoints; /*  Matrix of Gauss Points */
	VEC * weights; /*  Vector of Gauss Weights */
	double jacobian; /*  Jacobian of line element */
	EFG_SF * _phi;
}gauss1D;

typedef struct {
    MAT * verticies; 
    MAT * elements;
    int numElements; 
    int quad_order ;
    MAT * quad_points;
    VEC * quad_weights;
    VEC * quad_jacobian;
    int dim ;
} gauss_block;

typedef struct material_block{
    double rho;
    double nu;
    double E;
} material_block;

typedef struct options_node_select{

 char type[10];
 double inflate;
 // for circle finds points within a circle with center x,y and radius r + inflate
 VEC * center;
 double radius;

 // for box finds points within a polygon defined by 3 points, inflated in all directions
 MAT * segment;
	// box is generated from the inflated segemnt . 

}options_node_select;


typedef enum _boolean {false, true} boolean;

typedef struct deformation_tensors {
	MAT * B ; // Right C-G or Finger tensor
	MAT * C ;// Left C-G
	MAT * F ; // Deformation Gradinet
	MAT * F_T ; // Transposed Deformation Gradient 
	MAT * invF; // Inverse deformation gradient
	MAT * invF_T; // Transposed inverse Deformation Gradient
	MAT * Bstar;// Distortional component of finger tensor
	MAT * Cstar; // Distortional component of left Cauchy Green tensor 
}deformation_tensors;
typedef struct strain_disp{
	MAT * B_L;
	MAT * B_L1;
	MAT * B_L0; 
	MAT * B_L_T;
	IVEC * en;
	VEC * f_sub; 

}strain_disp;

typedef struct stress_tensors {
	MAT * Stress_C ; // 2PKF
	MAT * Stress_1PKF ; // 1PKF
	MAT * Stress_2PKF; // Cauchy
	VEC * Stress_Voigt ; // voigt stress
}stress_tensors;

typedef struct _timestep{
	double deltaT;
	double tmax;
	double Wkin;
	double Wint;
	double Wext;
	double Wcon;
	double epsilon; /* see belytsko */
}timestep;


typedef struct _kinVars{
	VEC * d; /*  Displacements */
	VEC * v; /*  Velocity */
	VEC * vh; /*  Velocity at half time point so if v_n+1/2 */
	VEC * a; /*  Acceleration */
}kinVars;

typedef struct _materialState{

	double temperature ; /*  Temperature */
	VEC * stress; /*  Stress  */
	MAT * B ; /*  Cauchy deformation Tensor */
	MAT * F ; /*  Previous deformation gradient */
	MAT * D ; /*  Rate of deformation tensor */



}materialState;

typedef struct EBC {
	MAT * V;
	MAT * phi;
	IVEC * nodes;
	MAT * coords; 
	MAT * P; 
	int dofFixed;
	VEC * uBar1;
	VEC * uBar2;
	VEC * uCorrect1;
	VEC * uCorrect2;
}EBC;

typedef struct essentialBoundary{
	// node numbers and coords 
	IVEC * nodes ;
	MAT * coords; 
	// transformation matricies - going from displacement parameters to displacements !
	MAT * lambdaNI ;// interior
	MAT * lambdaNB ;// boundary
	MAT * invLambdaNB ;// inverse matrix needed for consistent method;
	VEC * uBar; // prescribed dispalacment 
	VEC * d_NI ;// displacement due to nodal group NI
	VEC * d_NB ;// displacement due to nodal group NB 
	int dofFixed ;
	
	
	// nodal displacement
	
	VEC * disp ; 
	
}essentialBoundary;

typedef struct Node
{
	jcv_point coords;
	int node_nr;
}Node;

typedef struct Meshfree
{
	Node * nodes;
	double * di;
	int num_nodes;
	
}Meshfree;


typedef struct displacement {

	VEC * r ;
	VEC * z ;

}displacement;


typedef struct SCNI {
	gpc_polygon * poly;
	IVEC * sfIndex ;
	VEC * phi;
	MAT * B ; 
	MAT * F;
	VEC * fInt;
	double area; 
	double * center;
	double coords[2];
	int index;
}SCNI;


typedef struct _State {

	MAT * Fbar;
	MAT * Dbar;
	MAT * Wbar;
	MAT * Bbar;
	MAT * Sb;
	MAT * Sc;
	VEC * lambdaBar;
	VEC * eigValDBar;
	double mSigma; 
	double critLambdaBar; 
	double lambdaNMax;
	double temperature;
	double Jacobian;






}States;

typedef struct _State * State;


#endif 
