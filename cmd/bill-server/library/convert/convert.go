package convert

import (
	"bytes"
	"reflect"
	"strings"
)

// ExtractFields extract all none zero value field value from struct using reflection
func ExtractFields(obj interface{}) (ret map[string]interface{}) {
	ret = make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}
		if !f.IsValid() {
			// in case f is accessed with nil pointer
			continue
		}
		if f.Interface() != reflect.Zero(f.Type()).Interface() {
			ret[toSnakeCase(t.Field(i).Name)] = f.Interface()
		}
	}
	return ret
}

// transform camel case to lower snake case
func toSnakeCase(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(v)
	}
	return strings.ToLower(buf.String())
}

// SetFields set struct all field value using map
func SetFields(obj interface{}, data map[string]interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}
		if f.CanSet() {
			if f.Kind() == reflect.Struct {
				if err := SetFields(f.Addr().Interface(), data); err != nil {
					return err
				}
			} else {
				if val, ok := data[toSnakeCase(t.Field(i).Name)]; ok {
					f.Set(reflect.ValueOf(val))
				}
			}
		}
	}
	return nil
}

// Transform transform struct to another struct
func Transform(obj interface{}, dst interface{}) error {
	data := ExtractFields(obj)
	return SetFields(dst, data)
}
