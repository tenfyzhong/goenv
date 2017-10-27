package goenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal unmarshal os.env to o
// o must be point to a struct
func Unmarshal(o interface{}) error {
	rv := reflect.ValueOf(o)
	return unmarshalValue(rv, "")
}

func unmarshalValue(rv reflect.Value, tagPrefix string) error {
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("must be a pointer")
	}

	if !rv.IsValid() {
		return errors.New("zero value")
	}

	t := rv.Type()
	t = t.Elem()
	v := rv.Elem()

	if v.Kind() != reflect.Struct {
		return errors.New("must be point to a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				value.Set(reflect.New(value.Type().Elem()))
			}
			value = value.Elem()
		}
		if !value.CanSet() {
			continue
		}
		tag := field.Tag.Get("env")
		name := tag
		if tagPrefix != "" && tag != "" {
			name = strings.Join([]string{tagPrefix, tag}, ".")
		}
		if tag == "" {
			if value.Kind() == reflect.Struct {
				err := unmarshalValue(value.Addr(), tagPrefix)
				if err != nil {
					return err
				}
			}
			continue
		}
		strValue := os.Getenv(name)
		if strValue == "" {
			strValue = field.Tag.Get("envdef")
		}
		switch value.Kind() {
		case reflect.Bool:
			value.SetBool(strValue != "")
		case reflect.String:
			value.SetString(strValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			env, err := strconv.Atoi(strValue)
			if err == nil {
				value.SetInt(int64(env))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			env, err := strconv.Atoi(strValue)
			if err == nil {
				value.SetUint(uint64(env))
			}
		case reflect.Float32, reflect.Float64:
			env, err := strconv.ParseFloat(strValue, 64)
			if err == nil {
				value.SetFloat(env)
			}
		case reflect.Struct:
			err := unmarshalValue(value.Addr(), name)
			if err != nil {
				return err
			}
		default:
		}
	}

	return nil

}
