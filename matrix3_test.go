package geometry

import (
	"fmt"
	"testing"
)

func TestRotation(t *testing.T) {
	tmat := NewTransMat()
	tmat.XRotation(0)
	tmat.YRotation(0)
	tmat.ZRotation(0)
	tmat.Translation(30, 0, 0)
	fmt.Println(tmat)
}
