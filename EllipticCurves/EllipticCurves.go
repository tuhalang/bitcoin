package main

import (
	"errors"
	"fmt"
	"math"
)

/*
INF ...
*/
const INF = math.MaxInt32

/*
EllipticCurves ...
*/
type EllipticCurves struct {
	x int32
	y int32
	a int32
	b int32
}

/*
Init ...
*/
func Init(x, y, a, b int32) (error, EllipticCurves) {
	if y*y != (x*x*x + a*x + b) {
		err := "Error (x, y) is not on the curve"
		e := errors.New(err)
		return e, EllipticCurves{}
	}
	return nil, EllipticCurves{
		x: x,
		y: y,
		a: a,
		b: b,
	}
}

/*
Add ..
*/
func Add(e1, e2 EllipticCurves) (error, EllipticCurves) {
	if e1.a != e2.a || e1.b != e2.b {
		err := "Two point are not on the same curves"
		e := errors.New(err)
		return e, EllipticCurves{}
	}

	if e1.x == INF {
		return nil, e2
	}

	if e2.x == INF {
		return nil, e1
	}

	//handle the case where the two points are additive inverses
	if e1.x == e2.x && e1.y != e2.y {
		return nil, EllipticCurves{
			x: INF,
			y: INF,
			a: e1.a,
			b: e1.b,
		}
	}

	// P != Q
	if e1.x != e2.x && e1.y != e2.y {
		s := (e2.y - e1.y) / (e2.x - e1.x)
		x := s*s - e1.x - e2.x
		y := s*(e1.x-x) - e1.y
		return nil, EllipticCurves{
			x: x,
			y: y,
			a: e1.a,
			b: e2.b,
		}
	}

	// P == Q
	if e1.x == e2.x && e1.y == e2.y {
		if e1.y == 0 {
			return nil, EllipticCurves{
				x: INF,
				y: INF,
				a: e1.a,
				b: e1.b,
			}
		} else {
			s := (3*e1.x*e1.x + e1.a) / 2 * e1.y
			x := s*s - 2*e1.x
			y := s*(e1.x-x) - e1.y
			return nil, EllipticCurves{
				x: x,
				y: y,
				a: e1.a,
				b: e1.b,
			}
		}
	}

	return nil, EllipticCurves{}
}

func main() {
	//_, e1 := Init(3, 7, 5, 7)
	_, e2 := Init(-1, -1, 5, 7)
	err, e3 := Add(e2, e2)
	if err == nil {
		fmt.Println(e3)
	} else {
		fmt.Println(err)
	}
}
