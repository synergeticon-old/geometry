package geometry_test

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/synergeticon/geometry"
)

func TestSavePLY(t *testing.T) {
	pc := geometry.PointCloud{}
	pc.FillRandom(100)

	err := pc.SavePLY("./test.ply")
	if err != nil {
		t.Error(err)
	}
}

func TestTransform(t *testing.T) {
	pc := geometry.PointCloud{}

	v1 := mat64.NewVector(3, []float64{1, 2, 3})
	v2 := mat64.NewVector(3, []float64{1, 0, 0})
	v3 := mat64.NewVector(3, []float64{1, 1, 1})
	pc.Vectors = append(pc.Vectors, v1, v2, v3)

	tmat := geometry.NewTransMat()
	tmat.Translation(10, 0, 0)

	pc.Transform(tmat)
	if v1.At(0, 0)+10 != pc.Vectors[0].At(0, 0) {
		t.Error("translation by 10 failed")
	}

	if v2.At(0, 0)+10 != pc.Vectors[1].At(0, 0) {
		t.Error("translation by 10 failed")
	}

	if v3.At(0, 0)+10 != pc.Vectors[2].At(0, 0) {
		t.Error("translation by 10 failed")
	}

}

func TestViewer(t *testing.T) {
	pc := geometry.PointCloud{}
	pc.FillRandom(100)
	err := pc.ShowInMeshlab()
	if err != nil {
		t.Error(err)
	}
}
