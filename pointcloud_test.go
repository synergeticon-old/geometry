package geometry

import (
	"testing"
)

func TestSavePLY(t *testing.T) {
	pc := PointCloud{}
	pc.FillRandom(100)

	err := pc.SavePLY("./test.ply")
	if err != nil {
		t.Error(err)
	}
}
