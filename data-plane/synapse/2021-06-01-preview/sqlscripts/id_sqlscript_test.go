package sqlscripts

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SqlScriptId{}

func TestNewSqlScriptID(t *testing.T) {
	id := NewSqlScriptID("https://endpoint-url.example.com", "sqlScriptName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SqlScriptName != "sqlScriptName" {
		t.Fatalf("Expected %q but got %q for Segment 'SqlScriptName'", id.SqlScriptName, "sqlScriptName")
	}
}

func TestFormatSqlScriptID(t *testing.T) {
	actual := NewSqlScriptID("https://endpoint-url.example.com", "sqlScriptName").ID()
	expected := "https://endpoint-url.example.com/sqlScripts/sqlScriptName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSqlScriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlScriptId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/sqlScripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sqlScripts/sqlScriptName",
			Expected: &SqlScriptId{
				BaseURI:       "https://endpoint-url.example.com",
				SqlScriptName: "sqlScriptName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sqlScripts/sqlScriptName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlScriptID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.SqlScriptName != v.Expected.SqlScriptName {
			t.Fatalf("Expected %q but got %q for SqlScriptName", v.Expected.SqlScriptName, actual.SqlScriptName)
		}

	}
}

func TestParseSqlScriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlScriptId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/sqlScripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sQlScRiPtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sqlScripts/sqlScriptName",
			Expected: &SqlScriptId{
				BaseURI:       "https://endpoint-url.example.com",
				SqlScriptName: "sqlScriptName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sqlScripts/sqlScriptName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sQlScRiPtS/sQlScRiPtNaMe",
			Expected: &SqlScriptId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SqlScriptName: "sQlScRiPtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sQlScRiPtS/sQlScRiPtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlScriptIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.SqlScriptName != v.Expected.SqlScriptName {
			t.Fatalf("Expected %q but got %q for SqlScriptName", v.Expected.SqlScriptName, actual.SqlScriptName)
		}

	}
}

func TestSegmentsForSqlScriptId(t *testing.T) {
	segments := SqlScriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SqlScriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
