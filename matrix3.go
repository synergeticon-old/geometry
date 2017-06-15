package geometry

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

func NewXRotation(angle float64) *mat64.Dense {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	data := []float64{
		1, 0, 0, 0,
		0, cosine, -sine, 0,
		0, sine, cosine, 0,
		0, 0, 0, 1,
	}

	mat := mat64.NewDense(4, 4, data)
	return mat
}

func NewYRotation(angle float64) *mat64.Dense {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	data := []float64{
		cosine, 0, sine, 0,
		0, 1, 0, 0,
		-sine, 0, cosine, 0,
		0, 0, 0, 1,
	}

	mat := mat64.NewDense(4, 4, data)
	return mat
}

func NewZRotation(angle float64) *mat64.Dense {
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	data := []float64{
		cosine, -sine, 0, 0,
		sine, cosine, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	mat := mat64.NewDense(4, 4, data)
	return mat
}

func NewTranslation(dX, dY, dZ float64) *mat64.Dense {
	data := []float64{
		1, 0, 0, dX,
		0, 1, 0, dY,
		0, 0, 1, dZ,
		0, 0, 0, 1,
	}
	mat := mat64.NewDense(4, 4, data)
	return mat
}
