package geometry

import (
	"fmt"
	"math"

	"github.com/gonum/matrix/mat64"
)

type CatmullRome3 struct {
	pc PointCloud
}

func NewCatmullRome3(pc PointCloud) *CatmullRome3 {
	cr := &CatmullRome3{pc}
	return cr
}

func (cr *CatmullRome3) GetPoint(t float64) *mat64.Vector {
	point := mat64.NewVector(3, []float64{0, 0, 0})

	// Vector for calculations
	tmp := mat64.NewVector(3, []float64{0, 0, 0})

	points := cr.pc.Vectors
	l := len(points)

	p := (float64(l) - 1) * t

	intPoint := int(math.Floor(p))

	weight := p - float64(intPoint)

	if weight == 0 && intPoint == l-1 {
		intPoint = l - 2
		weight = 1
	}

	p0 := mat64.NewVector(3, []float64{0, 0, 0})
	p1 := mat64.NewVector(3, []float64{0, 0, 0})
	p2 := mat64.NewVector(3, []float64{0, 0, 0})
	p3 := mat64.NewVector(3, []float64{0, 0, 0})

	if intPoint > 0 {
		p0 = points[(intPoint-1)%l]
	} else {
		// extrapolate first point
		tmp.SubVec(points[0], points[1])
		tmp.AddVec(tmp, points[0])
		p0.CloneVec(tmp)
	}

	p1 = points[intPoint%l]
	p2 = points[(intPoint+1)%l]

	if intPoint+2 < l {
		p3 = points[(intPoint+2)%l]
	} else {
		tmp.SubVec(points[l-1], points[l-2])
		tmp.AddVec(tmp, points[l-1])
		p3.CloneVec(tmp)
	}

	px := NewCubicPoly()
	py := NewCubicPoly()
	pz := NewCubicPoly()

	// init Centripetal / Chordal Catmull-Rom
	pow := 0.25
	var dt0 = math.Pow(DistanceSquared(p0, p1), pow)
	var dt1 = math.Pow(DistanceSquared(p1, p2), pow)
	var dt2 = math.Pow(DistanceSquared(p2, p3), pow)

	// safety check for repeated points
	if dt1 < 1e-4 {
		dt1 = 1.0
	}
	if dt0 < 1e-4 {
		dt0 = dt1
	}
	if dt2 < 1e-4 {
		dt2 = dt1
	}

	px.initNonuniformCatmullRom(p0.At(0, 0), p1.At(0, 0), p2.At(0, 0), p3.At(0, 0), dt0, dt1, dt2)
	py.initNonuniformCatmullRom(p0.At(1, 0), p1.At(1, 0), p2.At(1, 0), p3.At(1, 0), dt0, dt1, dt2)
	pz.initNonuniformCatmullRom(p0.At(2, 0), p1.At(2, 0), p2.At(2, 0), p3.At(2, 0), dt0, dt1, dt2)

	point.SetVec(0, px.calc(weight))
	point.SetVec(1, py.calc(weight))
	point.SetVec(2, pz.calc(weight))

	// 	return point;
	return point
}

func (cr *CatmullRome3) IntersectPlane(plane *Plane3) (*mat64.Vector, bool) {
	intersection := mat64.NewVector(3, []float64{0, 0, 0})

	threshold := 10e-10
	maxIterations := 10000

	errorDistance := math.Inf(1)

	t := 0.0
	i := 0

	for errorDistance > threshold && i < maxIterations {
		intersection = cr.GetPoint(t)
		fmt.Println(intersection)
		errorDistance = math.Abs(plane.DistanceToPoint(intersection))
		fmt.Println(errorDistance)
		t = t + 0.001*errorDistance
		i++
	}

	if errorDistance > threshold {
		return nil, false
	}

	return intersection, true
}

type CubicPoly struct {
	c0 float64
	c1 float64
	c2 float64
	c3 float64
}

func NewCubicPoly() *CubicPoly {
	cp := &CubicPoly{}
	return cp
}

func (cp *CubicPoly) init(x0, x1, t0, t1 float64) {
	cp.c0 = x0
	cp.c1 = t0
	cp.c2 = -3*x0 + 3*x1 - 2*t0 - t1
	cp.c3 = 2*x0 - 2*x1 + t0 + t1
}

func (cp *CubicPoly) initCatmullRom(x0, x1, x2, x3, tension float64) {
	cp.init(x1, x2, tension*(x2-x0), tension*(x3-x1))
}

func (cp *CubicPoly) initNonuniformCatmullRom(x0, x1, x2, x3, dt0, dt1, dt2 float64) {
	// compute tangents when parameterized in [t1,t2]
	t1 := (x1-x0)/dt0 - (x2-x0)/(dt0+dt1) + (x2-x1)/dt1
	t2 := (x2-x1)/dt1 - (x3-x1)/(dt1+dt2) + (x3-x2)/dt2

	// rescale tangents for parametrization in [0,1]
	t1 *= dt1
	t2 *= dt1

	cp.init(x1, x2, t1, t2)
}

func (cp *CubicPoly) calc(t float64) float64 {
	t2 := t * t
	t3 := t2 * t
	return cp.c0 + cp.c1*t + cp.c2*t2 + cp.c3*t3
}
