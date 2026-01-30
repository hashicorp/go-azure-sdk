package indexes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IndexId{}

func TestNewIndexID(t *testing.T) {
	id := NewIndexID("https://endpoint-url.example.com", "indexName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.IndexName != "indexName" {
		t.Fatalf("Expected %q but got %q for Segment 'IndexName'", id.IndexName, "indexName")
	}
}

func TestFormatIndexID(t *testing.T) {
	actual := NewIndexID("https://endpoint-url.example.com", "indexName").ID()
	expected := "https://endpoint-url.example.com/indexes/indexName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIndexID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IndexId
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
			Input: "https://endpoint-url.example.com/indexes",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/indexes/indexName",
			Expected: &IndexId{
				BaseURI:   "https://endpoint-url.example.com",
				IndexName: "indexName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/indexes/indexName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIndexID(v.Input)
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

		if actual.IndexName != v.Expected.IndexName {
			t.Fatalf("Expected %q but got %q for IndexName", v.Expected.IndexName, actual.IndexName)
		}

	}
}

func TestParseIndexIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IndexId
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
			Input: "https://endpoint-url.example.com/indexes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNdExEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/indexes/indexName",
			Expected: &IndexId{
				BaseURI:   "https://endpoint-url.example.com",
				IndexName: "indexName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/indexes/indexName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNdExEs/iNdExNaMe",
			Expected: &IndexId{
				BaseURI:   "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				IndexName: "iNdExNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNdExEs/iNdExNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIndexIDInsensitively(v.Input)
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

		if actual.IndexName != v.Expected.IndexName {
			t.Fatalf("Expected %q but got %q for IndexName", v.Expected.IndexName, actual.IndexName)
		}

	}
}

func TestSegmentsForIndexId(t *testing.T) {
	segments := IndexId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IndexId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
