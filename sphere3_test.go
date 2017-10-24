package geometry_test

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/synergeticon/geometry"
)

func TestSphereFromPoints(t *testing.T) {
	sph := geometry.Sphere3{}

	points := []*mat64.Vector{}

	v1 := mat64.NewVector(3, []float64{
		0,
		-2,
		0,
	})
	v2 := mat64.NewVector(3, []float64{
		0,
		2,
		0,
	})
	v3 := mat64.NewVector(3, []float64{
		0,
		0,
		2,
	})
	v4 := mat64.NewVector(3, []float64{
		0,
		0,
		-2,
	})
	points = append(points, v1, v2, v3, v4)

	sph.SetFromPoints(points)

	if sph.Radius != 2 {
		t.Error("Radius should be 2 but is", sph.Radius)
	}

	// Test for Center
}
