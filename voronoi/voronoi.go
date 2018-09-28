package voronoi

type Voronoi struct {
	polygon      int
	num_polygons int
}

func NewVoronoi(polygonIn int, num_polygonsIn int) *Voronoi {
	return &Voronoi{polygon: polygonIn, num_polygons: num_polygonsIn}
}

func GenerateVoronoi() {

}
