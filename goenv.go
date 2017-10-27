package goenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	envTag     = "env"
	sepTag     = "envsep"
	defaultTag = "envdef"
)

var (
	sliceOfBool    = reflect.TypeOf([]bool(nil))
	sliceOfString  = reflect.TypeOf([]string(nil))
	sliceOfInt     = reflect.TypeOf([]int(nil))
	sliceOfInt8    = reflect.TypeOf([]int8(nil))
	sliceOfInt16   = reflect.TypeOf([]int16(nil))
	sliceOfInt32   = reflect.TypeOf([]int32(nil))
	sliceOfInt64   = reflect.TypeOf([]int64(nil))
	sliceOfUInt    = reflect.TypeOf([]uint(nil))
	sliceOfUInt8   = reflect.TypeOf([]uint8(nil))
	sliceOfUInt16  = reflect.TypeOf([]uint16(nil))
	sliceOfUInt32  = reflect.TypeOf([]uint32(nil))
	sliceOfUInt64  = reflect.TypeOf([]uint64(nil))
	sliceOfFloat32 = reflect.TypeOf([]float32(nil))
	sliceOfFloat64 = reflect.TypeOf([]float64(nil))
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
		tag := field.Tag.Get(envTag)
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
			strValue = field.Tag.Get(defaultTag)
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
		case reflect.Slice:
			sep := field.Tag.Get(sepTag)
			if sep == "" {
				sep = ","
			}
			handleSlice(value, strValue, sep)
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

func handleSlice(value reflect.Value, strValue, sep string) {
	switch value.Type() {
	case sliceOfBool:
		result := parseBoolSlice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfString:
		result := parseStringSlice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfInt:
		result := parseIntSlice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfInt8:
		result := parseInt8Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfInt16:
		result := parseInt16Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfInt32:
		result := parseInt32Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfInt64:
		result := parseInt64Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfUInt:
		result := parseUIntSlice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfUInt8:
		result := parseUInt8Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfUInt16:
		result := parseUInt16Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfUInt32:
		result := parseUInt32Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfUInt64:
		result := parseUInt64Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfFloat32:
		result := parseFloat32Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	case sliceOfFloat64:
		result := parseFloat64Slice(strValue, sep)
		value.Set(reflect.ValueOf(result))
	default:
	}
}

func parseBoolSlice(strValue, sep string) []bool {
	if sep == "" {
		return nil
	}
	items := strings.Split(strValue, sep)
	result := make([]bool, 0, len(items))
	for _, i := range items {
		if i != "" {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func parseStringSlice(strValue, sep string) []string {
	if sep == "" {
		return nil
	}
	items := strings.Split(strValue, sep)
	return items
}

func parseInt64Slice(strValue, sep string) []int64 {
	if sep == "" {
		return nil
	}
	items := strings.Split(strValue, sep)
	result := make([]int64, 0, len(items))
	for _, i := range items {
		v, err := strconv.ParseInt(strings.TrimSpace(i), 10, 64)
		if err == nil {
			result = append(result, v)
		}
	}
	return result
}

func parseIntSlice(strValue, sep string) []int {
	result := parseInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	ints := make([]int, 0, len(result))
	for _, r := range result {
		ints = append(ints, int(r))
	}
	return ints
}

func parseInt8Slice(strValue, sep string) []int8 {
	result := parseInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	int8s := make([]int8, 0, len(result))
	for _, r := range result {
		int8s = append(int8s, int8(r))
	}
	return int8s
}

func parseInt16Slice(strValue, sep string) []int16 {
	result := parseInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	int16s := make([]int16, 0, len(result))
	for _, r := range result {
		int16s = append(int16s, int16(r))
	}
	return int16s
}

func parseInt32Slice(strValue, sep string) []int32 {
	result := parseInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	int32s := make([]int32, 0, len(result))
	for _, r := range result {
		int32s = append(int32s, int32(r))
	}
	return int32s
}

func parseUInt64Slice(strValue, sep string) []uint64 {
	if sep == "" {
		return nil
	}
	items := strings.Split(strValue, sep)
	result := make([]uint64, 0, len(items))
	for _, i := range items {
		v, err := strconv.ParseUint(strings.TrimSpace(i), 10, 64)
		if err == nil {
			result = append(result, v)
		}
	}
	return result
}

func parseUIntSlice(strValue, sep string) []uint {
	result := parseUInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	uints := make([]uint, 0, len(result))
	for _, r := range result {
		uints = append(uints, uint(r))
	}
	return uints
}

func parseUInt8Slice(strValue, sep string) []uint8 {
	result := parseUInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	uint8s := make([]uint8, 0, len(result))
	for _, r := range result {
		uint8s = append(uint8s, uint8(r))
	}
	return uint8s
}

func parseUInt16Slice(strValue, sep string) []uint16 {
	result := parseUInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	uint16s := make([]uint16, 0, len(result))
	for _, r := range result {
		uint16s = append(uint16s, uint16(r))
	}
	return uint16s
}

func parseUInt32Slice(strValue, sep string) []uint32 {
	result := parseUInt64Slice(strValue, sep)
	if result == nil {
		return nil
	}
	uint32s := make([]uint32, 0, len(result))
	for _, r := range result {
		uint32s = append(uint32s, uint32(r))
	}
	return uint32s
}

func parseFloat64Slice(strValue, sep string) []float64 {
	if sep == "" {
		return nil
	}
	items := strings.Split(strValue, sep)
	result := make([]float64, 0, len(items))
	for _, i := range items {
		v, err := strconv.ParseFloat(strings.TrimSpace(i), 64)
		if err == nil {
			result = append(result, v)
		}
	}
	return result
}

func parseFloat32Slice(strValue, sep string) []float32 {
	if sep == "" {
		return nil
	}
	items := strings.Split(strValue, sep)
	result := make([]float32, 0, len(items))
	for _, i := range items {
		v, err := strconv.ParseFloat(strings.TrimSpace(i), 32)
		if err == nil {
			result = append(result, float32(v))
		}
	}
	return result
}
