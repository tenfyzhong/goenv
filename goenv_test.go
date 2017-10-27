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
	One    int    `env:"one" envdef:"101"`
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

type DefautltTest struct {
	D0 int    `env:"d0" envdef:"100"`
	D1 string `env:"d1" envdef:"d1"`
}

func TestDefault(t *testing.T) {
	o := &DefautltTest{}
	err := Unmarshal(o)
	if err != nil {
		t.Fatalf("Unmarshal Number should be success, err: %v", err)
	}
	if o.D0 != 100 {
		t.Fatalf("d0 default is 100, but: %d", o.D0)
	}
	if o.D1 != "d1" {
		t.Fatalf("d1 default is 'd1', but: %s", o.D1)
	}
}

type SliceTest struct {
	Bools    []bool    `env:"bools" envsep:","`
	Strings  []string  `env:"strings" envsep:","`
	Ints     []int     `env:"ints" envsep:","`
	Int8s    []int8    `env:"int8s" envsep:","`
	Int16s   []int16   `env:"int16s" envsep:","`
	Int32s   []int32   `env:"int32s" envsep:","`
	Int64s   []int64   `env:"int64s" envsep:","`
	Uints    []uint    `env:"uints" envsep:","`
	Uint8s   []uint8   `env:"uint8s" envsep:","`
	Uint16s  []uint16  `env:"uint16s" envsep:","`
	Uint32s  []uint32  `env:"uint32s" envsep:","`
	Uint64s  []uint64  `env:"uint64s" envsep:","`
	Float32s []float32 `env:"float32s" envsep:","`
	Float64s []float64 `env:"float64s"`
}

func TestSlice(t *testing.T) {
	o := &SliceTest{}
	os.Setenv("bools", ",1")
	os.Setenv("strings", "hello,world")
	os.Setenv("ints", "1, 2, 3")
	os.Setenv("int8s", " 1, 2, 3 ")
	os.Setenv("int16s", " 1, 2,3 ")
	os.Setenv("int32s", " 1, 2 ,3")
	os.Setenv("int64s", " 1, 2 ,3")
	os.Setenv("uints", "1, 2,3")
	os.Setenv("uint8s", " 1, 2,3 ")
	os.Setenv("uint16s", " 1, 2,3 ")
	os.Setenv("uint32s", " 1, 2 ,3")
	os.Setenv("uint64s", " 1, 2 ,3")
	os.Setenv("float32s", "1.0,2.0,3")
	os.Setenv("float64s", "1.0,2.0,3")

	err := Unmarshal(o)
	if err != nil {
		t.Fatalf("Unmarshal SliceTest should be success, err: %v", err)
	}

	if len(o.Bools) != 2 || o.Bools[0] != false || o.Bools[1] != true {
		t.Fatalf("o.Bools failed, %v", o.Bools)
	}
	if len(o.Strings) != 2 || o.Strings[0] != "hello" || o.Strings[1] != "world" {
		t.Fatalf("o.String failed, %v", o.Strings)
	}
	if len(o.Ints) != 3 || o.Ints[0] != 1 || o.Ints[1] != 2 || o.Ints[2] != 3 {
		t.Fatalf("o.Ints failed, %v", o.Ints)
	}
	if len(o.Int8s) != 3 || o.Int8s[0] != 1 || o.Int8s[1] != 2 || o.Int8s[2] != 3 {
		t.Fatalf("o.Int8s failed, %v", o.Int8s)
	}
	if len(o.Int16s) != 3 || o.Int16s[0] != 1 || o.Int16s[1] != 2 || o.Int16s[2] != 3 {
		t.Fatalf("o.Int16s failed, %v", o.Int16s)
	}
	if len(o.Int32s) != 3 || o.Int32s[0] != 1 || o.Int32s[1] != 2 || o.Int32s[2] != 3 {
		t.Fatalf("o.Int32s failed, %v", o.Int32s)
	}
	if len(o.Int64s) != 3 || o.Int64s[0] != 1 || o.Int64s[1] != 2 || o.Int64s[2] != 3 {
		t.Fatalf("o.Int64s failed, %v", o.Int64s)
	}
	if len(o.Uints) != 3 || o.Uints[0] != 1 || o.Uints[1] != 2 || o.Uints[2] != 3 {
		t.Fatalf("o.Uints failed, %v", o.Uints)
	}
	if len(o.Uint8s) != 3 || o.Uint8s[0] != 1 || o.Uint8s[1] != 2 || o.Uint8s[2] != 3 {
		t.Fatalf("o.Uint8s failed, %v", o.Uint8s)
	}
	if len(o.Uint16s) != 3 || o.Uint16s[0] != 1 || o.Uint16s[1] != 2 || o.Uint16s[2] != 3 {
		t.Fatalf("o.Uint16s failed, %v", o.Uint16s)
	}
	if len(o.Uint32s) != 3 || o.Uint32s[0] != 1 || o.Uint32s[1] != 2 || o.Uint32s[2] != 3 {
		t.Fatalf("o.Uint32s failed, %v", o.Uint32s)
	}
	if len(o.Uint64s) != 3 || o.Uint64s[0] != 1 || o.Uint64s[1] != 2 || o.Uint64s[2] != 3 {
		t.Fatalf("o.Uint64s failed, %v", o.Uint64s)
	}
	if len(o.Float32s) != 3 || o.Float32s[0] != 1.0 || o.Float32s[1] != 2.0 || o.Float32s[2] != 3.0 {
		t.Fatalf("o.Floats failed, %v", o.Float32s)
	}
	if len(o.Float64s) != 3 || o.Float64s[0] != 1.0 || o.Float64s[1] != 2.0 || o.Float64s[2] != 3.0 {
		t.Fatalf("o.Floats failed, %v", o.Float64s)
	}
}
