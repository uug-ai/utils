package ustrings

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

const (
	DefaultPrecision      = 2
	DefaultBase           = 10
	DefaultFormat         = 'f'
	DefaultNilMapFormat   = "{}"
	DefaultNilSliceFormat = "[]"
	DefaultNilFormat      = "<nil>"
)

// Options defines stringification behavior.
type Options struct {
	Base           int
	Format         byte
	Precision      int
	NilFormat      string
	NilMapFormat   string
	NilSliceFormat string
}

func parseOptions(opts ...Options) Options {
	if len(opts) == 0 {
		return Options{
			Base:           DefaultBase,
			Format:         DefaultFormat,
			Precision:      DefaultPrecision,
			NilFormat:      DefaultNilFormat,
			NilMapFormat:   DefaultNilMapFormat,
			NilSliceFormat: DefaultNilSliceFormat,
		}
	}
	options := opts[0]
	if options.Base == 0 {
		options.Base = DefaultBase
	}
	if options.Format == 0 {
		options.Format = DefaultFormat
	}
	if options.Precision == 0 {
		options.Precision = DefaultPrecision
	}
	if options.NilFormat == "" {
		options.NilFormat = DefaultNilFormat
	}
	if options.NilMapFormat == "" {
		options.NilMapFormat = DefaultNilMapFormat
	}
	if options.NilSliceFormat == "" {
		options.NilSliceFormat = DefaultNilSliceFormat
	}
	return options
}

// Stringify converts a value to a string with optional formatting options.
func Stringify(value any, opts ...Options) string {
	options := parseOptions(opts...)
	return stringifyValue(value, options)
}

func stringifyValue(value any, options Options) string {
	if value == nil {
		return options.NilFormat
	}
	switch cast := value.(type) {
	case string:
		return cast
	case []byte:
		return string(cast)
	case fmt.Stringer:
		return cast.String()
	case error:
		return cast.Error()
	case bool:
		return strconv.FormatBool(cast)
	case int:
		return strconv.FormatInt(int64(cast), options.Base)
	case int8:
		return strconv.FormatInt(int64(cast), options.Base)
	case int16:
		return strconv.FormatInt(int64(cast), options.Base)
	case int32:
		return strconv.FormatInt(int64(cast), options.Base)
	case int64:
		return strconv.FormatInt(cast, options.Base)
	case uint:
		return strconv.FormatUint(uint64(cast), options.Base)
	case uint8:
		return strconv.FormatUint(uint64(cast), options.Base)
	case uint16:
		return strconv.FormatUint(uint64(cast), options.Base)
	case uint32:
		return strconv.FormatUint(uint64(cast), options.Base)
	case uint64:
		return strconv.FormatUint(cast, options.Base)
	case uintptr:
		return strconv.FormatUint(uint64(cast), options.Base)
	case float32:
		return strconv.FormatFloat(float64(cast), options.Format, options.Precision, 32)
	case float64:
		return strconv.FormatFloat(cast, options.Format, options.Precision, 64)
	case complex64:
		return strconv.FormatComplex(complex128(cast), options.Format, options.Precision, 64)
	case complex128:
		return strconv.FormatComplex(cast, options.Format, options.Precision, 128)
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Pointer:
		if rv.IsNil() {
			return options.NilFormat
		}
		return stringifyValue(rv.Elem().Interface(), options)
	case reflect.Slice:
		if rv.IsNil() {
			return options.NilSliceFormat
		}
		return stringifySlice(rv, options)
	case reflect.Array:
		return stringifySlice(rv, options)
	case reflect.Map:
		if rv.IsNil() {
			return options.NilMapFormat
		}
		return stringifyMap(rv, options)
	case reflect.Struct:
		return stringifyStruct(rv, options)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func stringifySlice(rv reflect.Value, options Options) string {
	length := rv.Len()
	items := make([]string, 0, length)
	for i := 0; i < length; i++ {
		items = append(items, stringifyValue(rv.Index(i).Interface(), options))
	}
	return "[" + joinWithComma(items) + "]"
}

func stringifyMap(rv reflect.Value, options Options) string {
	keys := rv.MapKeys()
	stringKeys := make([]string, 0, len(keys))
	keyByString := make(map[string]reflect.Value, len(keys))
	for _, key := range keys {
		keyString := fmt.Sprint(key.Interface())
		stringKeys = append(stringKeys, keyString)
		keyByString[keyString] = key
	}
	sort.Strings(stringKeys)
	items := make([]string, 0, len(stringKeys))
	for _, keyString := range stringKeys {
		value := rv.MapIndex(keyByString[keyString])
		entry := fmt.Sprintf("%s: %s", keyString, stringifyValue(value.Interface(), options))
		items = append(items, entry)
	}
	return "{" + joinWithComma(items) + "}"
}

func stringifyStruct(rv reflect.Value, options Options) string {
	rt := rv.Type()
	items := make([]string, 0, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.PkgPath != "" {
			continue
		}
		value := rv.Field(i)
		entry := fmt.Sprintf("%s: %s", field.Name, stringifyValue(value.Interface(), options))
		items = append(items, entry)
	}
	return "{" + joinWithComma(items) + "}"
}

func joinWithComma(items []string) string {
	if len(items) == 0 {
		return ""
	}
	return stringsJoin(items, ", ")
}

func stringsJoin(items []string, sep string) string {
	if len(items) == 0 {
		return ""
	}
	length := 0
	for _, item := range items {
		length += len(item)
	}
	length += len(sep) * (len(items) - 1)
	builder := make([]byte, 0, length)
	for i, item := range items {
		if i > 0 {
			builder = append(builder, sep...)
		}
		builder = append(builder, item...)
	}
	return string(builder)
}
