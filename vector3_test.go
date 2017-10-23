package geometry_test

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/synergeticon/geometry"
)

func TestLength(t *testing.T) {
	v := mat64.NewVector(3, []float64{0, 1, 0})
	length := geometry.Length(v)
	if length != 1 {
		t.Error("Length is", length, "should be 1")
	}
}
func TestNormalize(t *testing.T) {
	v := mat64.NewVector(3, []float64{2, 2, 2})
	length := geometry.Length(v)
	t.Log("Length before:", length)

	geometry.Normalize(v)
	lengthAfter := geometry.Length(v)
	if lengthAfter != 1 {
		t.Error("Length is ", lengthAfter, "should be 1")
	}
}

func TestDistance(t *testing.T) {

	// Two vectors of which we want to calc. the distance
	v1 := mat64.NewVector(3, []float64{
		0,
		1,
		0,
	})

	v2 := mat64.NewVector(3, []float64{
		0,
		1,
		1,
	})

	// Calculating the Distance
	distance := geometry.Distance(v1, v2)

	if distance != 1 {
		t.Error("Distance should be 1, but is", distance)
	}
}
