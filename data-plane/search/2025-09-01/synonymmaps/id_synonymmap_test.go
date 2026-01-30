package synonymmaps

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SynonymmapId{}

func TestNewSynonymmapID(t *testing.T) {
	id := NewSynonymmapID("https://endpoint-url.example.com", "synonymmapName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SynonymmapName != "synonymmapName" {
		t.Fatalf("Expected %q but got %q for Segment 'SynonymmapName'", id.SynonymmapName, "synonymmapName")
	}
}

func TestFormatSynonymmapID(t *testing.T) {
	actual := NewSynonymmapID("https://endpoint-url.example.com", "synonymmapName").ID()
	expected := "https://endpoint-url.example.com/synonymmaps/synonymmapName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSynonymmapID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SynonymmapId
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
			Input: "https://endpoint-url.example.com/synonymmaps",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/synonymmaps/synonymmapName",
			Expected: &SynonymmapId{
				BaseURI:        "https://endpoint-url.example.com",
				SynonymmapName: "synonymmapName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/synonymmaps/synonymmapName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSynonymmapID(v.Input)
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

		if actual.SynonymmapName != v.Expected.SynonymmapName {
			t.Fatalf("Expected %q but got %q for SynonymmapName", v.Expected.SynonymmapName, actual.SynonymmapName)
		}

	}
}

func TestParseSynonymmapIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SynonymmapId
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
			Input: "https://endpoint-url.example.com/synonymmaps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sYnOnYmMaPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/synonymmaps/synonymmapName",
			Expected: &SynonymmapId{
				BaseURI:        "https://endpoint-url.example.com",
				SynonymmapName: "synonymmapName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/synonymmaps/synonymmapName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sYnOnYmMaPs/sYnOnYmMaPnAmE",
			Expected: &SynonymmapId{
				BaseURI:        "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SynonymmapName: "sYnOnYmMaPnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sYnOnYmMaPs/sYnOnYmMaPnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSynonymmapIDInsensitively(v.Input)
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

		if actual.SynonymmapName != v.Expected.SynonymmapName {
			t.Fatalf("Expected %q but got %q for SynonymmapName", v.Expected.SynonymmapName, actual.SynonymmapName)
		}

	}
}

func TestSegmentsForSynonymmapId(t *testing.T) {
	segments := SynonymmapId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SynonymmapId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
