package geometry

import (
	"testing"

	"github.com/gonum/matrix/mat64"
)

func TestLength(t *testing.T) {
	v := mat64.NewVector(3, []float64{0, 1, 0})
	length := Length(v)
	if length != 1 {
		t.Error("Length is", length, "should be 1")
	}
}
func TestNormalize(t *testing.T) {
	v := mat64.NewVector(3, []float64{2, 2, 2})
	length := Length(v)
	t.Log("Length before:", length)

	Normalize(v)
	lengthAfter := Length(v)
	if lengthAfter != 1 {
		t.Error("Length is ", lengthAfter, "should be 1")
	}
}
