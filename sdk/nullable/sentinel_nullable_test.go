package nullable_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

type InnerStruct struct {
	InnerName *string `json:"inner_name,omitempty"`
	InnerID   string  `json:"inner_id"`
	Number    int64   `json:"number,omitempty"`
}

func (i InnerStruct) MarshalJSON() ([]byte, error) {
	return nullable.MarshalNullableStruct(i)
}

type TestStruct struct {
	ID      string       `json:"id"`
	OmitID  string       `json:"omit_id,omitempty"`
	Name    *string      `json:"name,omitempty"`
	Age     *float64     `json:"age,omitempty"`
	Address *string      `json:"address,omitempty"`
	Inner   *InnerStruct `json:"inner,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (t TestStruct) MarshalJSON() ([]byte, error) {
	return nullable.MarshalNullableStruct(t)
}

var _ json.Marshaler = TestStruct{}

func TestNullableValuePanic(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("Expected panic, but got nil")
		} else if !strings.Contains(e.(string), "Elem of invalid type") {
			t.Fatalf("Expected panic of invalid type but got %v", e)
		}
	}()
	nullable.NullValue[string]()
}

func TestMarshalNullableNil(t *testing.T) {
	// nullable field address is nil, it should be omitted
	name := "John Doe"
	age := 30.0
	obj := TestStruct{
		Name: &name,
		Age:  &age,
	}

	expected := map[string]interface{}{
		"age":  age,
		"id":   "",
		"name": name,
	}

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("json.Unmarshal returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMarshalNullable(t *testing.T) {
	// nullable field address is set to null value, it should be included
	name := "John Doe"
	age := 30.0
	obj := TestStruct{
		Name: &name,
		Age:  &age,
	}
	obj.Address = nullable.NullValue[*string]()

	expected := map[string]interface{}{
		"age":     age,
		"address": nil,
		"id":      "",
		"name":    name,
	}

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("json.Unmarshal returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMarshalNullableInnerStruct(t *testing.T) {
	name := "John Doe"
	age := 30.0
	obj := TestStruct{
		Name: &name,
		Age:  &age,
	}
	obj.Address = nullable.NullValue[*string]()
	obj.Inner = nullable.NullValue[*InnerStruct]()

	expected := map[string]interface{}{
		"address": nil,
		"age":     age,
		"id":      "",
		"inner":   nil,
		"name":    name,
	}
	expectedBytes, _ := json.Marshal(expected)

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	if !reflect.DeepEqual(data, expectedBytes) {
		t.Errorf("Expected %s, but got %s", expectedBytes, data)
	}
}

func TestMarshalNullableWithInnerNullale(t *testing.T) {
	name := "John Doe"
	age := 30.0
	obj := TestStruct{
		Name: &name,
		Age:  &age,
		Inner: &InnerStruct{
			InnerID:   "",
			InnerName: nullable.NullValue[*string](),
		},
	}
	obj.Address = nullable.NullValue[*string]()

	expected := map[string]interface{}{
		"address": nil,
		"age":     age,
		"id":      "",
		"inner": map[string]interface{}{
			"inner_id":   "", // inner_id is not omitempty flagged
			"inner_name": nil,
		},
		"name": name,
	}
	expectedBytes, _ := json.Marshal(expected)

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	if !reflect.DeepEqual(data, expectedBytes) {
		t.Errorf("Expected %s, but got %s", expectedBytes, data)
	}
}

type ITest interface {
	Foo() string
}

type TestImpl struct{}

func (t TestImpl) Foo() string {
	return "foo"
}

type NullableInterface struct {
	ITest ITest `json:"itest,omitempty"`
}

func (n NullableInterface) MarshalJSON() ([]byte, error) {
	return nullable.MarshalNullableStruct(n)
}

func TestMarshalNullableWithInterface(t *testing.T) {
	obj := NullableInterface{
		ITest: nil,
	}

	expected := map[string]interface{}{}
	expectedBytes, _ := json.Marshal(expected)

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	if !reflect.DeepEqual(data, expectedBytes) {
		t.Errorf("Expected %s, but got %s", expectedBytes, data)
	}
}

func TestMarshalNullableWithInterfaceNullValue(t *testing.T) {
	// for interface, set the null value of it's implementation
	obj := NullableInterface{
		ITest: nullable.NullValue[*TestImpl](),
	}

	expected := map[string]interface{}{
		"itest": nil,
	}
	expectedBytes, _ := json.Marshal(expected)

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	if !reflect.DeepEqual(data, expectedBytes) {
		t.Errorf("Expected %s, but got %s", expectedBytes, data)
	}
}
