package ustrings

import "testing"

type sampleStruct struct {
	A int
	B string
}

func TestStringifyPrimitive(t *testing.T) {
	if Stringify(42) != "42" {
		t.Errorf("Stringify int did not match")
	}
	if Stringify(42, Options{Base: 16}) != "2a" {
		t.Errorf("Stringify base16 did not match")
	}
	if Stringify(3.14159, Options{Precision: 2}) != "3.14" {
		t.Errorf("Stringify float did not match")
	}
	if Stringify(true) != "true" {
		t.Errorf("Stringify bool did not match")
	}
}

func TestStringifyNil(t *testing.T) {
	var s []string
	var m map[string]int
	var p *int
	if Stringify(s) != "[]" {
		t.Errorf("Stringify nil slice did not match")
	}
	if Stringify(m) != "{}" {
		t.Errorf("Stringify nil map did not match")
	}
	if Stringify(p) != "<nil>" {
		t.Errorf("Stringify nil pointer did not match")
	}
}

func TestStringifySlice(t *testing.T) {
	if Stringify([]int{1, 2, 3}) != "[1, 2, 3]" {
		t.Errorf("Stringify slice did not match")
	}
}

func TestStringifyMap(t *testing.T) {
	input := map[string]int{"b": 2, "a": 1}
	if Stringify(input) != "{a: 1, b: 2}" {
		t.Errorf("Stringify map did not match")
	}
}

func TestStringifyStruct(t *testing.T) {
	input := sampleStruct{A: 1, B: "x"}
	if Stringify(input) != "{A: 1, B: x}" {
		t.Errorf("Stringify struct did not match")
	}
}
