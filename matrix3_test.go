package geometry

import (
	"fmt"
	"testing"

	"github.com/gonum/matrix/mat64"
)

func TestRotation(t *testing.T) {

	rZ := NewZRotation(0 * (3.1415 / 180))
	rX := NewXRotation(0 * (3.1415 / 180))
	rZ2 := NewZRotation(0 * (3.1415 / 180))
	transl := NewTranslation(0, 30, 0)

	transform := &mat64.Dense{}
	transform.Mul(rZ, rX)
	transform.Mul(transform, rZ2)
	transform.Mul(transform, transl)
	fmt.Println(transform)
}
