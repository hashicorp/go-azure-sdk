package environments

import (
	"reflect"
	"testing"
)

func TestFromName(t *testing.T) {
	testData := []struct {
		name     string
		expected *Environment
	}{
		{
			name:     "canary",
			expected: AzurePublicCanary(),
		},
		{
			name:     "china",
			expected: AzureChina(),
		},
		{
			name:     "dod",
			expected: AzureUSGovernmentL5(),
		},
		{
			name:     "global",
			expected: AzurePublic(),
		},
		{
			name:     "public",
			expected: AzurePublic(),
		},
		{
			name:     "usgovernment",
			expected: AzureUSGovernment(),
		},
		{
			name:     "usgovernmentl4",
			expected: AzureUSGovernment(),
		},
		{
			name:     "usgovernmentl5",
			expected: AzureUSGovernmentL5(),
		},
		{
			name:     "mars",
			expected: nil,
		},
	}
	for _, v := range testData {
		t.Logf("Printing %q", v.name)
		actual, err := FromName(v.name)
		if err != nil {
			if v.expected == nil {
				continue
			}
			t.Fatalf("unexpected error: %+v", err)
		}
		if v.expected == nil {
			t.Fatalf("expected an error but got the environment: %+v", *v.expected)
		}
		if actual == nil {
			t.Fatal("expected an environment but got nil")
		}
		if !reflect.DeepEqual(*actual, *v.expected) {
			t.Fatalf("expected and actual environment didn't match.\n\nExpected: %+v.\n\nActual: %+v.", *v.expected, *actual)
		}
	}
}
