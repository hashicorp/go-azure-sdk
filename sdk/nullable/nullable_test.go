package nullable

import (
	"bytes"
	"testing"
)

func TestTypeMarshalBool(t *testing.T) {
	testCases := []struct {
		value    Type[bool]
		expected []byte
	}{
		{
			value:    Value(true),
			expected: []byte(`true`),
		},
		{
			value:    Value(false),
			expected: []byte(`false`),
		},
		{
			value:    NoZero(true),
			expected: []byte(`true`),
		},
		{
			value:    NoZero(false),
			expected: []byte(`null`),
		},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.MarshalJSON()
		if err != nil {
			t.Errorf("test case %d: %v", i, err)
		}
		if !bytes.Equal(result, testCase.expected) {
			t.Errorf("test case %d: expected %q, got %q", i, string(testCase.expected), string(result))
		}
	}
}

func TestTypeMarshalFloat(t *testing.T) {
	testCases := []struct {
		value    Type[float64]
		expected []byte
	}{
		{
			value:    Value(123.45),
			expected: []byte(`123.45`),
		},
		{
			value:    Value(0.0),
			expected: []byte(`0`),
		},
		{
			value:    NoZero(123.45),
			expected: []byte(`123.45`),
		},
		{
			value:    NoZero(-123.45),
			expected: []byte(`-123.45`),
		},
		{
			value:    NoZero(0.0),
			expected: []byte(`null`),
		},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.MarshalJSON()
		if err != nil {
			t.Errorf("test case %d: %v", i, err)
		}
		if !bytes.Equal(result, testCase.expected) {
			t.Errorf("test case %d: expected %q, got %q", i, string(testCase.expected), string(result))
		}
	}
}

func TestTypeMarshalInt(t *testing.T) {
	testCases := []struct {
		value    Type[int]
		expected []byte
	}{
		{
			value:    Value(123),
			expected: []byte(`123`),
		},
		{
			value:    Value(0),
			expected: []byte(`0`),
		},
		{
			value:    NoZero(123),
			expected: []byte(`123`),
		},
		{
			value:    NoZero(-123),
			expected: []byte(`-123`),
		},
		{
			value:    NoZero(0),
			expected: []byte(`null`),
		},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.MarshalJSON()
		if err != nil {
			t.Errorf("test case %d: %v", i, err)
		}
		if !bytes.Equal(result, testCase.expected) {
			t.Errorf("test case %d: expected %q, got %q", i, string(testCase.expected), string(result))
		}
	}
}

func TestTypeMarshalString(t *testing.T) {
	testCases := []struct {
		value    Type[string]
		expected []byte
	}{
		{
			value:    Value("foo"),
			expected: []byte(`"foo"`),
		},
		{
			value:    Value(""),
			expected: []byte(`""`),
		},
		{
			value:    NoZero("foo"),
			expected: []byte(`"foo"`),
		},
		{
			value:    NoZero(""),
			expected: []byte(`null`),
		},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.MarshalJSON()
		if err != nil {
			t.Errorf("test case %d: %v", i, err)
		}
		if !bytes.Equal(result, testCase.expected) {
			t.Errorf("test case %d: expected %q, got %q", i, string(testCase.expected), string(result))
		}
	}
}
