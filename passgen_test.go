package main

import (
	"testing"
)

func TestCheckInput(t *testing.T) {
	_, err := checkInput("y")
	if err != nil {
		t.Fatal(err)
	}

	val2, err2 := checkInput("ye")
	if err2 == nil {
		t.Fatalf("Non nil error: %s", val2)
	}

	_, err3 := checkInput("yes")
	if err3 != nil {
		t.Fatal(err3)
	}
	
	val4, err4 := checkInput("ys")
	if err4 == nil {
		t.Fatalf("Non nil error: %s", val4)
	}

	_, err5 := checkInput("Y")
	if err5 != nil {
		t.Fatal(err5)
	}

	_, err6 := checkInput("YeS")
	if err6 != nil {
		t.Fatal(err6)
	}
	
	_, err7 := checkInput("n")
	if err7 != nil {
		t.Fatal(err7)
	}

	_, err8 := checkInput("no")
	if err8 != nil {
		t.Fatal(err8)
	}

	_, err9 := checkInput("N")
	if err9 != nil {
		t.Fatal(err9)
	}

	_, err10 := checkInput("nO")
	if err10 != nil {
		t.Fatal(err10)
	}
}

