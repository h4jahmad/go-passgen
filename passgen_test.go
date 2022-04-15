package main

import (
	"testing"
)

func TestCheckInput(t *testing.T) {
	_, err := checkInput("y")
	assertNilError(err, t)
	_, err2 := checkInput("ye")
	assertNonNilError(err2, t)
	_, err3 := checkInput("yes")
	assertNilError(err3, t)
	_, err4 := checkInput("ys")
	assertNonNilError(err4, t)
	_, err5 := checkInput("Y")
	assertNilError(err5, t)
	_, err6 := checkInput("YeS")
	assertNilError(err6, t)
	
	_, err7 := checkInput("n")
	assertNilError(err7, t)
	_, err8 := checkInput("no")
	assertNilError(err8, t)
	_, err9 := checkInput("N")
	assertNilError(err9, t)
	_, err10 := checkInput("nO")
	assertNilError(err10, t)
}

func assertNilError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertNonNilError(err error, t *testing.T) {
	if err == nil {
		t.Fatal("Non nil error")
	}
}
