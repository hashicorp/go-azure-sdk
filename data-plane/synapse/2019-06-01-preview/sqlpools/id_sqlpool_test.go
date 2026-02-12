package sqlpools

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SqlPoolId{}

func TestNewSqlPoolID(t *testing.T) {
	id := NewSqlPoolID("https://endpoint-url.example.com", "sqlPoolName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SqlPoolName != "sqlPoolName" {
		t.Fatalf("Expected %q but got %q for Segment 'SqlPoolName'", id.SqlPoolName, "sqlPoolName")
	}
}

func TestFormatSqlPoolID(t *testing.T) {
	actual := NewSqlPoolID("https://endpoint-url.example.com", "sqlPoolName").ID()
	expected := "https://endpoint-url.example.com/sqlPools/sqlPoolName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSqlPoolID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlPoolId
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
			Input: "https://endpoint-url.example.com/sqlPools",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sqlPools/sqlPoolName",
			Expected: &SqlPoolId{
				BaseURI:     "https://endpoint-url.example.com",
				SqlPoolName: "sqlPoolName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sqlPools/sqlPoolName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlPoolID(v.Input)
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

		if actual.SqlPoolName != v.Expected.SqlPoolName {
			t.Fatalf("Expected %q but got %q for SqlPoolName", v.Expected.SqlPoolName, actual.SqlPoolName)
		}

	}
}

func TestParseSqlPoolIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlPoolId
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
			Input: "https://endpoint-url.example.com/sqlPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sQlPoOlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sqlPools/sqlPoolName",
			Expected: &SqlPoolId{
				BaseURI:     "https://endpoint-url.example.com",
				SqlPoolName: "sqlPoolName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sqlPools/sqlPoolName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sQlPoOlS/sQlPoOlNaMe",
			Expected: &SqlPoolId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SqlPoolName: "sQlPoOlNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sQlPoOlS/sQlPoOlNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSqlPoolIDInsensitively(v.Input)
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

		if actual.SqlPoolName != v.Expected.SqlPoolName {
			t.Fatalf("Expected %q but got %q for SqlPoolName", v.Expected.SqlPoolName, actual.SqlPoolName)
		}

	}
}

func TestSegmentsForSqlPoolId(t *testing.T) {
	segments := SqlPoolId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SqlPoolId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
