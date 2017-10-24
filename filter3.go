package geometry

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

type PassThroughFilter struct {
	LimitLow    float64
	LimitHigh   float64
	filterField int
	Normal      *mat64.Vector
}

func (ptf *PassThroughFilter) SetFilterFieldName(field string) {
	switch field {
	case "x":
		ptf.filterField = 0
	case "y":
		ptf.filterField = 1
	case "z":
		ptf.filterField = 2
	default:
		panic("must be x, y or z")
	}
}

func (ptf *PassThroughFilter) Filter(pc PointCloud) PointCloud {
	filteredCloud := PointCloud{}
	if ptf.Normal == nil {
		for _, point := range pc.Vectors {
			value := point.At(ptf.filterField, 0)
			fmt.Println(value)
			if value < ptf.LimitHigh && value > ptf.LimitLow {
				filteredCloud.Vectors = append(filteredCloud.Vectors, point)
			}
		}
	}
	return filteredCloud
}
