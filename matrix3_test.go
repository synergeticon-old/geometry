package geometry

import (
	"fmt"
	"math"
	"testing"

	"github.com/gonum/matrix/mat64"
)

func TestRotation(t *testing.T) {
	tmat := NewTransMat()
	tmat.YRotation(math.Pi / 2)
	tmat.XRotation(math.Pi)
	tmat.ZRotation(math.Pi / 3)
	tmat.Translation(9, 0, 0)
	vec := mat64.NewVector(4, []float64{1, 0, 0, 1})
	erg := tmat.Transform(vec)
	fmt.Println(erg)
}
