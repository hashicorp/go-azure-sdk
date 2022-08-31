package odata_test

import (
	odata2 "github.com/hashicorp/go-azure-sdk/sdk/odata"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func TestQueryHeaders(t *testing.T) {
	type testCase struct {
		query    odata2.Query
		expected http.Header
	}
	testCases := []testCase{
		{
			query: odata2.Query{},
			expected: http.Header{
				"Accept":           []string{"application/json; charset=utf-8; IEEE754Compatible=false"},
				"Odata-Maxversion": []string{odata2.ODataVersion},
				"Odata-Version":    []string{odata2.ODataVersion},
			},
		},
		{
			query: odata2.Query{
				ConsistencyLevel: odata2.ConsistencyLevelEventual,
				Metadata:         odata2.MetadataMinimal,
			},
			expected: http.Header{
				"Accept":           []string{"application/json; charset=utf-8; IEEE754Compatible=false; odata.metadata=minimal"},
				"Consistencylevel": []string{"eventual"},
				"Odata-Maxversion": []string{odata2.ODataVersion},
				"Odata-Version":    []string{odata2.ODataVersion},
			},
		},
	}
	for n, c := range testCases {
		v := c.query.Headers()
		if !reflect.DeepEqual(v, c.expected) {
			t.Errorf("test case %d for %s: expected %#v, got %#v", n, "Headers()", c.expected, v)
		}
		c.expected.Set("Testing", "foo")
		v = c.query.AppendHeaders(http.Header{"Testing": []string{"foo"}})
		if !reflect.DeepEqual(v, c.expected) {
			t.Errorf("test case %d for %s: expected %#v, got %#v", n, "AppendHeaders()", c.expected, v)
		}
	}
}

func TestQueryValues(t *testing.T) {
	type testCase struct {
		query    odata2.Query
		expected url.Values
	}
	testCases := []testCase{
		{
			query:    odata2.Query{},
			expected: url.Values{},
		},
		{
			query: odata2.Query{
				Count:  true,
				Format: odata2.FormatAtom,
				Skip:   20,
				Top:    10,
			},
			expected: url.Values{
				"$count":  []string{"true"},
				"$format": []string{"atom"},
				"$skip":   []string{"20"},
				"$top":    []string{"10"},
			},
		},
		{
			query: odata2.Query{
				OrderBy: odata2.OrderBy{
					Field:     "displayName",
					Direction: "desc",
				},
			},
			expected: url.Values{
				"$orderby": []string{"displayName desc"},
			},
		},
		{
			query: odata2.Query{
				Expand: odata2.Expand{
					Relationship: "children",
					Select:       []string{"id", "childName"},
				},
			},
			expected: url.Values{
				"$expand": []string{"children($select=id,childName)"},
			},
		},
		{
			query: odata2.Query{
				Filter: "startsWith(displayName,'Widgets')",
			},
			expected: url.Values{
				"$filter": []string{"startsWith(displayName,'Widgets')"},
			},
		},
		{
			query: odata2.Query{
				Search: "displayName:Astley",
			},
			expected: url.Values{
				"$search": []string{`"displayName:Astley"`},
			},
		},
		{
			query: odata2.Query{
				Select: []string{"id", "userPrincipalName"},
			},
			expected: url.Values{
				"$select": []string{"id,userPrincipalName"},
			},
		},
	}
	for n, c := range testCases {
		v := c.query.Values()
		if !reflect.DeepEqual(v, c.expected) {
			t.Errorf("test case %d: expected %#v, got %#v", n, c.expected, v)
		}
	}
}
