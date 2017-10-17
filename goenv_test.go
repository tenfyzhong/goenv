package goenv

import (
	"os"
	"testing"
)

func TestNil(t *testing.T) {
	err := Unmarshal(nil)
	if err == nil {
		t.Fatalf("Unmarshal nil must be error")
	}
	if err.Error() != "must be a pointer" {
		t.Fatalf("err: must be a pointer, but: %v", err)
	}
}

type Number struct {
	zero   int    `env:"zero"`
	One    int    `env:"one"`
	Two    int    `env:"two"`
	Three  bool   `env:"three"`
	Four   string `env:"four"`
	Five   string
	Six    *int8   `env:"six"`
	Sevent uint    `env:"sevent"`
	Eight  float32 `env:"eight"`
	Nine   bool    `env:"nine"`
	Ten    *bool   `env:"ten"`
}

func TestNoPointer(t *testing.T) {
	o := Number{}
	err := Unmarshal(o)
	if err == nil {
		t.Fatalf("Unmarshal no pointer must be error")
	}
	if err.Error() != "must be a pointer" {
		t.Fatalf("must be a pointer")
	}
}

func testNoStruct(t *testing.T) {
	v := ""
	err := Unmarshal(&v)
	if err == nil {
		t.Fatalf("Unmarshal no struct must be error")
	}
	if err.Error() != "must be point to a struct" {
		t.Fatalf("Unmarshal no struct must be error")
	}
}

func TestUnmarshal(t *testing.T) {
	o := &Number{}
	os.Setenv("zero", "10")
	os.Setenv("one", "1")
	os.Setenv("two", "two")
	os.Setenv("three", "1")
	os.Setenv("four", "4")
	os.Setenv("five", "5")
	os.Setenv("six", "6")
	os.Setenv("sevent", "7")
	os.Setenv("eight", "8.0")
	os.Setenv("ten", "0")

	err := Unmarshal(o)
	if err != nil {
		t.Fatalf("Unmarshal Number should be success, err: %v", err)
	}

	if o.zero != 0 {
		t.Fatalf("zero should be 0, %v", o.zero)
	}
	if o.One != 1 {
		t.Fatalf("one should be 1, %v", o.One)
	}
	if o.Two != 0 {
		t.Fatalf("two should be 0, %v", o.Two)
	}
	if o.Three != true {
		t.Fatalf("o.Three should be true")
	}
	if o.Four != "4" {
		t.Fatalf("o.Four should 4, %v", o.Four)
	}
	if o.Five != "" {
		t.Fatalf("o.Five should be empty, %v", o.Five)
	}
	if *o.Six != int8(6) {
		t.Fatalf("o.Six should be 6, %v", *o.Six)
	}
	if o.Sevent != uint(7) {
		t.Fatalf("o.Sevent should be 7, %v", o.Sevent)
	}
	if o.Eight != float32(8.0) {
		t.Fatalf("o.Eight should be 8.0, %v", o.Eight)
	}
	if o.Nine != false {
		t.Fatalf("o.Nine should be false")
	}
	if *o.Ten != true {
		t.Fatalf("o.Ten should be false")
	}
}
