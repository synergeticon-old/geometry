package geometry_test

import (
	"fmt"
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

func TestContainsPoint(t *testing.T) {
	center := mat64.NewVector(3, []float64{0, 0, 0})
	size := mat64.NewVector(3, []float64{10, 10, 10})

	b3 := geometry.NewBox3(nil, nil)
	b3.SetFromCenterAndSize(center, size)

	outlier := mat64.NewVector(3, []float64{50, 3, 2})
	inlier := mat64.NewVector(3, []float64{1, 1, 1})

	if b3.ContainsPoint(outlier) != false {
		t.Error("outlier is not contained by box")
	}

	if b3.ContainsPoint(inlier) != true {
		t.Error("inlier is contained by box, but function says it isn't")
	}
}

func TestExpandByScalar(t *testing.T) {
	center := mat64.NewVector(3, []float64{0, 0, 0})
	size := mat64.NewVector(3, []float64{10, 10, 10})

	b3 := geometry.NewBox3(nil, nil)
	b3.SetFromCenterAndSize(center, size)

	fmt.Println(b3.Volume())

	b3.ExpandByScalar(1)
	fmt.Println(b3.Volume())
	t.Error()
}
