package nullable

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"
)

// holds sentinel values used to send nulls
var nullSentinels map[reflect.Type]any = map[reflect.Type]any{}
var nullablesLock sync.RWMutex

// NullValue is used to send an explicit 'null' within a request.
// This is typically used in JSON-MERGE-PATCH operations to delete a value.
// Type arugment `T` MUST be a pointer type (pointer, map, or slice)
// for interface type's null value, a pointer to implementor type is required
func NullValue[T any]() T {
	t := reflect.TypeFor[T]()

	nullablesLock.RLock()
	v, found := nullSentinels[t]
	nullablesLock.RUnlock()

	if found {
		// return the sentinel object
		if t.Kind() == reflect.Interface {
			var zero T
			return zero
		}
		return v.(T)
	}

	// promote to exclusive lock and check again (double-checked locking pattern)
	nullablesLock.Lock()
	defer nullablesLock.Unlock()

	v, found = nullSentinels[t]
	if !found {
		var o reflect.Value
		switch k := t.Kind(); k {
		case reflect.Map:
			o = reflect.MakeMap(t)
		case reflect.Slice:
			o = reflect.MakeSlice(t, 1, 1)
		default:
			// let it panic here if non-pointer type is passed
			o = reflect.New(t.Elem())
		}
		v = o.Interface()
		nullSentinels[t] = v
	}
	// return the sentinel object
	return v.(T)
}

func IsNullValue[T any](v T) bool {
	t := reflect.TypeOf(v)
	nullablesLock.RLock()
	defer nullablesLock.RUnlock()

	// if found, it MUST be a pointer, so never panic here
	if o, found := nullSentinels[t]; found {
		o1 := reflect.ValueOf(o)
		v1 := reflect.ValueOf(v)
		return o1.Pointer() == v1.Pointer()
	}
	return false
}

func MarshalNullableStruct(obj interface{}) ([]byte, error) {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.Struct:
		return marshalStruct(v)
	}
	return json.Marshal(obj)
}

func marshalStruct(v reflect.Value) ([]byte, error) {
	m := make(map[string]any)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		jsonName := field.Name
		omitEmtpy := false
		if tag := field.Tag.Get("json"); tag != "" {
			opts := strings.Split(tag, ",")
			jsonName = opts[0]
			for _, opt := range opts[1:] {
				if opt == "omitempty" {
					omitEmtpy = true
					break
				}
			}
		}

		rval := v.Field(i)
		val := rval.Interface()

		if val == nil || (omitEmtpy && rval.IsZero()) {
			continue
		} else if IsNullValue(val) {
			m[jsonName] = nil
		} else {
			m[jsonName] = val
		}
	}
	return json.Marshal(m)
}
