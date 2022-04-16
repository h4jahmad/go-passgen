package main

import (
	"testing"
)


func TestCheckInput(t *testing.T) {
	_, err := isYes("y")
	if err != nil {
		t.Fatal(err)
	}

	_, err2 := isYes("ye")
	if err2 == nil {
		t.Fatalf("Non nil error: %s", "ye")
	}

	_, err3 := isYes("yes")
	if err3 != nil {
		t.Fatal(err3)
	}
	
	_, err4 := isYes("ys")
	if err4 == nil {
		t.Fatalf("Non nil error: %s", "ys")
	}

	_, err5 := isYes("Y")
	if err5 != nil {
		t.Fatal(err5)
	}

	_, err6 := isYes("YeS")
	if err6 != nil {
		t.Fatal(err6)
	}
	
	_, err7 := isYes("n")
	if err7 != nil {
		t.Fatal(err7)
	}

	_, err8 := isYes("no")
	if err8 != nil {
		t.Fatal(err8)
	}

	_, err9 := isYes("N")
	if err9 != nil {
		t.Fatal(err9)
	}

	_, err10 := isYes("nO")
	if err10 != nil {
		t.Fatal(err10)
	}
}

