package geometry

import (
	"fmt"
	"testing"

	"github.com/gonum/matrix/mat64"
)

func TestRotation(t *testing.T) {
	tmat := NewTransMat()
	tmat.Translation(10, 5, 0)

	tmat.ZRotation(30 * (3.1415 / 180)) // dA
	tmat.XRotation(40 * (3.1415 / 180)) //dA
	tmat.ZRotation(50 * (3.1415 / 180))

	// tmat.Translation(1, 0, 0)

	// fmt.Println(tmat)
	vec := mat64.NewVector(3, []float64{-4, 98, 0})
	erg := tmat.Transform(vec)
	fmt.Println(erg)
}
