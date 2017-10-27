package geometry_test

import (
	"testing"

	"github.com/gonum/matrix/mat64"

	"github.com/synergeticon/geometry"
)

func TestSetFromCenterAndSize(t *testing.T) {

	center := mat64.NewVector(3, []float64{0, 0, 0})
	size := mat64.NewVector(3, []float64{10, 10, 10})

	b3 := geometry.NewBox3(nil, nil)
	b3.SetFromCenterAndSize(center, size)

	// TO-DO: Check Results
}
