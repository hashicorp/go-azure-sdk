// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accept

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	testData := []struct {
		Input    string
		Expected Header
	}{
		{
			Input: "text/html;charset=utf-8;q=0.8, application/xhtml+xml;charset=utf-8;q=0.9, application/xml;charset=utf-8, */*;q=0.1",
			Expected: Header{
				types: []PreferredType{
					{
						ContentType: "application/xml",
						Parameters: map[string]string{
							"charset": "utf-8",
						},
						Weight: 1000,
					},
					{
						ContentType: "application/xhtml+xml",
						Parameters: map[string]string{
							"charset": "utf-8",
						},
						Weight: 900,
					},
					{
						ContentType: "text/html",
						Parameters: map[string]string{
							"charset": "utf-8",
						},
						Weight: 800,
					},
					{
						ContentType: "*/*",
						Parameters:  map[string]string{},
						Weight:      100,
					},
				},
			},
		},
	}
	for i, v := range testData {
		t.Run(fmt.Sprintf("Iteration %d..", i), func(t *testing.T) {
			h, err := FromString(v.Input)
			if err != nil {
				t.Fatalf("error received: %+v", err)
			} else if !reflect.DeepEqual(h, v.Expected) {
				t.Fatalf("unexpected Header{} value.\nexpected:\n\n%#v\n\nreceived:\n\n%#v\n", v.Expected, h)
			}
		})
	}
}

func TestFirstChoice(t *testing.T) {
	testData := []struct {
		Input    Header
		Expected PreferredType
	}{
		{
			Input: Header{
				types: []PreferredType{
					{
						ContentType: "application/xml",
						Parameters: map[string]string{
							"charset": "utf-8",
						},
						Weight: 1000,
					},
					{
						ContentType: "application/xhtml+xml",
						Parameters: map[string]string{
							"charset": "utf-8",
						},
						Weight: 900,
					},
					{
						ContentType: "text/html",
						Parameters: map[string]string{
							"charset": "utf-8",
						},
						Weight: 800,
					},
					{
						ContentType: "*/*",
						Parameters:  map[string]string{},
						Weight:      100,
					},
				},
			},
			Expected: PreferredType{
				ContentType: "application/xml",
				Parameters: map[string]string{
					"charset": "utf-8",
				},
				Weight: 1000,
			},
		},
	}
	for i, v := range testData {
		t.Run(fmt.Sprintf("Iteration %d..", i), func(t *testing.T) {
			if r := v.Input.FirstChoice(); reflect.DeepEqual(r, v.Expected) {
				t.Fatalf("unexpected result.\nexpected:\n\n%#v\n\nreceived:\n\n%#v\n", v.Expected, r)
			}
		})
	}
}
