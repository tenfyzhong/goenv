// Package goenv is a package to unmarshal environments of the os to a struct
// object. It must be use tag for fields. The tag name is `env`. For example:
// ```go
// type Obj struct {
//		One int `env:"one"`
//		Two float32 `env:"two"`
// }
// ```
//
// supported type: bool, string, int, int8, int16, int32, int64, uint, uint8,
// uint16, uint32, uint64, float32, float64
//
// For bool type, if the environment is not set, the value is false.
// Otherwise, the value is true.
//
// The Field must be exported if want to unmarshal.
package goenv
