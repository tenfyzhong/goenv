# goenv
[![Build Status](https://travis-ci.org/tenfyzhong/goenv.svg?branch=master)](https://travis-ci.org/tenfyzhong/goenv)
[![codecov](https://codecov.io/gh/tenfyzhong/goenv/branch/master/graph/badge.svg)](https://codecov.io/gh/tenfyzhong/goenv)

Unmarshal env to struct. 

# doc
Package goenv is a package to unmarshal environments of the os to a struct
object. It must be use tag for fields. The tag name is `env`. For example:

```go
type Number struct {
    One int `env:"one"`
    Two float32 `env:"two"`
}
```

## supported type
- bool, []bool
- string, []string
- int, []int
- int8, []int8
- int16, []int16
- int32, []int32
- int64, []int64
- uint, []uint
- uint9, []uint9
- uint16, []uint16
- uint32, []uint32
- uint64, []uint64
- float32, []float32
- float64, []float64
- struct

For bool type, if the environment is not set, the value is false.
Otherwise, the value is true.

## field tag
- `env` envirnoment name
- `envdef` default value if unseted. 
- `envsep` if a slice, use this to split items. Default is `,`

The Field must be exported if want to unmarshal.

# example
```go
import "fmt"
import "github.com/tenfyzhong/goenv"

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
    Substruct *struct {
        S1 int `env:"s1" envdef:"1"`   // "substruct.s1"
        S2 []int `env:"s2" envdef:"1,2"` // "substruct.s2"
        S3 []int `env:"s3" envdef:"1:2" envsep:":"` // "substruct.s2"
    } `env:"substruct"`
}

func main() {
    n = &Number{}
    err := goenv.Unmarshal(n)
    if err == nil {
        fmt.Println(n)
    }
}
```
