package geometry_test

import "testing"
import "github.com/synergeticon/geometry"
import "github.com/gonum/matrix/mat64"

func TestDistanceToPoint(t *testing.T) {

	// Normal Vector for the Plane
	normal := mat64.NewVector(3, []float64{
		0,
		1,
		0,
	})

	// Definition of new Plane
	p := geometry.NewPlane3(normal, 0)

	// Point where we want to know the distance to plane
	point := mat64.NewVector(3, []float64{
		0,
		10,
		0,
	})

	// Calculation of the distance
	distance := p.DistanceToPoint(point)

	if distance != 10 {
		t.Error("distance should be 10, but is", distance)
	}
}
