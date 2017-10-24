package geometry_test

import "github.com/synergeticon/geometry"
import "testing"

func TestPassThroughFilter(t *testing.T) {
	pc := geometry.PointCloud{}
	pc.FillRandom(1000)

	flt := geometry.PassThroughFilter{}
	flt.LimitHigh = 0.1
	flt.LimitLow = -0.1
	flt.SetFilterFieldName("z")

	pc.ShowInMeshlab()
	filtered := flt.Filter(pc)
	filtered.ShowInMeshlab()
	t.Error()

}
