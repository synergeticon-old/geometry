package geometry

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

// TransMat is the Transformation Matrix representation.
type TransMat struct {
	mat64.Dense
}

// NewTransMat creates a new Transformation Matrix which an 4x4 Idendity.
func NewTransMat() *TransMat {
	data := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	mat := mat64.NewDense(4, 4, data)
	tmat := &TransMat{}
	tmat.Clone(mat)
	return tmat
}

// XRotation applies a X-Axis rotation.
func (tmat *TransMat) XRotation(angle float64) {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	data := []float64{
		1, 0, 0, 0,
		0, cosine, -sine, 0,
		0, sine, cosine, 0,
		0, 0, 0, 1,
	}

	mat := mat64.NewDense(4, 4, data)

	m2 := &TransMat{}
	m2.Clone(tmat)
	tmat.Mul(m2, mat)
}

// YRotation applies a Y-Axis rotation.
func (tmat *TransMat) YRotation(angle float64) {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	data := []float64{
		cosine, 0, sine, 0,
		0, 1, 0, 0,
		-sine, 0, cosine, 0,
		0, 0, 0, 1,
	}

	mat := mat64.NewDense(4, 4, data)

	m2 := &TransMat{}
	m2.Clone(tmat)
	tmat.Mul(m2, mat)
}

// ZRotation applies a Z-Axis rotation.
func (tmat *TransMat) ZRotation(angle float64) {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	data := []float64{
		cosine, -sine, 0, 0,
		sine, cosine, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	mat := mat64.NewDense(4, 4, data)

	m2 := &TransMat{}
	m2.Clone(tmat)
	tmat.Mul(m2, mat)
}

// Translation applies a Translation.
func (tmat *TransMat) Translation(dX, dY, dZ float64) {
	data := []float64{
		1, 0, 0, dX,
		0, 1, 0, dY,
		0, 0, 1, dZ,
		0, 0, 0, 1,
	}
	mat := mat64.NewDense(4, 4, data)

	m2 := &TransMat{}
	m2.Clone(tmat)
	tmat.Mul(m2, mat)
}

// Transform applis Transformation to a Vector
func (tmat *TransMat) Transform(vec *mat64.Vector) *mat64.Vector {
	v := &mat64.Vector{}
	v.MulVec(tmat, vec)
	return v
}
