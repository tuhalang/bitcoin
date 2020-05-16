package main

import (
	"errors"
	"fmt"
	"math"
)

/*
FiniteField is a struct declare finite field
*/
type FiniteField struct {
	num   int32
	prime int32
}

func powMode(a, x, p int32) int32 {
	return int32(math.Pow(float64(a), float64(x))) % p
}

/*
Create function is used to create new FiniteField
*/
func Create(num, prime int32) (error, FiniteField) {
	if num < 0 || num >= prime {
		err := "Input wrong"
		e := errors.New(err)
		return e, FiniteField{}
	}
	return nil, FiniteField{
		num:   num,
		prime: prime,
	}
}

/*
Add function is used to add two FiniteField
*/
func Add(f1, f2 FiniteField) (error, FiniteField) {

	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := (f1.num + f2.num) % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}

/*
Sub function is used to substract two finiteField
*/
func Sub(f1, f2 FiniteField) (error, FiniteField) {
	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := (f1.num - f2.num) % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}

/*
Mul function is used to product two finiteField
*/
func Mul(f1, f2 FiniteField) (error, FiniteField) {
	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := (f1.num * f2.num) % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}

/*
Pow function is used to  two finiteField
*/
func Pow(f FiniteField, exp int32) FiniteField {
	n := exp % (f.prime - 1)
	num := powMode(f.num, n, f.prime)
	return FiniteField{
		num:   num,
		prime: f.prime,
	}
}

/*
Div function is used to div two finiteField
*/
func Div(f1, f2 FiniteField) (error, FiniteField) {
	if f1.prime != f2.prime {
		err := "Cannot add two numbers in different Fields"
		e := errors.New(err)
		return e, FiniteField{}
	}
	num := powMode(f2.num, f1.prime-2, f1.prime) * f1.num % f1.prime
	return nil, FiniteField{
		num:   num,
		prime: f1.prime,
	}
}

func main() {
	_, f1 := Create(6, 17)

	_, f2 := Create(2, 17)

	_, f3 := Div(f1, f2)

	fmt.Println(f3)
}
