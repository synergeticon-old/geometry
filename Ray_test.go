package geometry_test

import (
	"fmt"
	"testing"

	"github.com/gonum/matrix/mat64"

	"github.com/synergeticon/geometry"
)

func TestRayDistanceToPoint(t *testing.T) {

	v1 := mat64.NewVector(3, []float64{0, 0, 0})
	v2 := mat64.NewVector(3, []float64{1, 0, 0})
	point := mat64.NewVector(3, []float64{0, 1, 0})

	ray := geometry.NewRayFromPoints(v1, v2)

	fmt.Println(point, ray)

	distance := ray.DistanceToPoint(point)
	if distance != 1 {
		t.Log("Should be 1, but is", distance)
		t.Fail()
	}
}
