package indexers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IndexerId{}

func TestNewIndexerID(t *testing.T) {
	id := NewIndexerID("https://endpoint-url.example.com", "indexerName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.IndexerName != "indexerName" {
		t.Fatalf("Expected %q but got %q for Segment 'IndexerName'", id.IndexerName, "indexerName")
	}
}

func TestFormatIndexerID(t *testing.T) {
	actual := NewIndexerID("https://endpoint-url.example.com", "indexerName").ID()
	expected := "https://endpoint-url.example.com/indexers/indexerName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIndexerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IndexerId
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
			Input: "https://endpoint-url.example.com/indexers",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/indexers/indexerName",
			Expected: &IndexerId{
				BaseURI:     "https://endpoint-url.example.com",
				IndexerName: "indexerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/indexers/indexerName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIndexerID(v.Input)
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

		if actual.IndexerName != v.Expected.IndexerName {
			t.Fatalf("Expected %q but got %q for IndexerName", v.Expected.IndexerName, actual.IndexerName)
		}

	}
}

func TestParseIndexerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IndexerId
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
			Input: "https://endpoint-url.example.com/indexers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNdExErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/indexers/indexerName",
			Expected: &IndexerId{
				BaseURI:     "https://endpoint-url.example.com",
				IndexerName: "indexerName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/indexers/indexerName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNdExErS/iNdExErNaMe",
			Expected: &IndexerId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				IndexerName: "iNdExErNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/iNdExErS/iNdExErNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIndexerIDInsensitively(v.Input)
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

		if actual.IndexerName != v.Expected.IndexerName {
			t.Fatalf("Expected %q but got %q for IndexerName", v.Expected.IndexerName, actual.IndexerName)
		}

	}
}

func TestSegmentsForIndexerId(t *testing.T) {
	segments := IndexerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IndexerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
