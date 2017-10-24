package geometry_test

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/synergeticon/geometry"
)

func TestPassThroughFilter(t *testing.T) {
	pc := geometry.PointCloud{}

	v1 := mat64.NewVector(3, []float64{
		0,
		2,
		0,
	})
	v2 := mat64.NewVector(3, []float64{
		0,
		1,
		0,
	})
	v3 := mat64.NewVector(3, []float64{
		0,
		1,
		1,
	})
	pc.Add(v1, v2, v3)

	flt := geometry.PassThroughFilter{}
	flt.LimitHigh = 1.5
	flt.LimitLow = 0
	flt.SetFilterFieldName("y")

	filtered := flt.Filter(pc)

	if filtered.Length() != 2 {
		t.Error("Filtered pointcloud must have 2 vertext but has", filtered.Length())
	}

}
