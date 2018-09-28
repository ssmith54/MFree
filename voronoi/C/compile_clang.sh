#!/usr/bin/env bash

gcc -DMKL_ILP64 -DMKL_DIRECT_CALL_SEQ -m64 -Wall -fopenmp -Wno-float-equal -Wno-implicit-int \
 main.c gpc.c trigen.c connect_segments.c generate_scni.c set_domain.c \
-I. -I/home/stephen/Documents/voronoi/src \
-I${MKLROOT}/include \
-I/home/stephen/Documents/software/Meschach \
-L/home/stephen/Documents/triangle -ltriangle -lm \
-L/home/stephen/Documents/software/Meschach -lmeschach \
 -Wl,--start-group ${MKLROOT}/lib/intel64/libmkl_intel_ilp64.a ${MKLROOT}/lib/intel64/libmkl_sequential.a ${MKLROOT}/lib/intel64/libmkl_core.a -Wl,--end-group -lpthread -lm -ldl \
 -o simple 

