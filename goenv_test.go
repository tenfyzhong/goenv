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

type Complex struct {
	S1 struct {
		S11 int `env:"s11"`
	} `env:"s1"`

	S2 *struct {
		S21 int `env:"s21"`
	} `env:"s2"`

	S3 struct {
		S31 int `env:"s31"`
	}

	S4 *struct {
		S41 int `env:"s41"`
	}

	S5 struct {
		S51 struct {
			S511 int `env:"s511"`
		}
		S52 struct {
			S521 int `env:"s521"`
		} `env:"s52"`
	} `env:"s5"`
}

func TestComplex(t *testing.T) {
	c := &Complex{}
	os.Setenv("s1.s11", "11")
	os.Setenv("s2.s21", "21")
	os.Setenv("s31", "31")
	os.Setenv("s41", "41")
	os.Setenv("s5.s511", "511")
	os.Setenv("s5.s52.s521", "521")

	err := Unmarshal(c)
	if err != nil {
		t.Fatalf("Unmarshal failed, err: %v", err)
	}

	if c.S1.S11 != 11 {
		t.Fatalf("c.S1.S11 should be 11, but: %d", c.S1.S11)
	}

	if c.S2.S21 != 21 {
		t.Fatalf("c.S2.S21 should be 21, but: %d", c.S2.S21)
	}

	if c.S3.S31 != 31 {
		t.Fatalf("c.S3.S31 should be 31; but: %d", c.S3.S31)
	}

	if c.S4.S41 != 41 {
		t.Fatalf("c.S4.S41 should be 41; but: %d", c.S4.S41)
	}

	if c.S5.S51.S511 != 511 {
		t.Fatalf("c.S5.S51.S511 should be 511, but: %d", c.S5.S51.S511)
	}

	if c.S5.S52.S521 != 521 {
		t.Fatalf("c.S5.S52.S521 should be 521, but: %d", c.S5.S52.S521)
	}
}
