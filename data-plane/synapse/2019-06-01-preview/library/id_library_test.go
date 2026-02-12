package library

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LibraryId{}

func TestNewLibraryID(t *testing.T) {
	id := NewLibraryID("https://endpoint-url.example.com", "libraryName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.LibraryName != "libraryName" {
		t.Fatalf("Expected %q but got %q for Segment 'LibraryName'", id.LibraryName, "libraryName")
	}
}

func TestFormatLibraryID(t *testing.T) {
	actual := NewLibraryID("https://endpoint-url.example.com", "libraryName").ID()
	expected := "https://endpoint-url.example.com/libraries/libraryName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLibraryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LibraryId
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
			Input: "https://endpoint-url.example.com/libraries",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/libraries/libraryName",
			Expected: &LibraryId{
				BaseURI:     "https://endpoint-url.example.com",
				LibraryName: "libraryName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/libraries/libraryName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLibraryID(v.Input)
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

		if actual.LibraryName != v.Expected.LibraryName {
			t.Fatalf("Expected %q but got %q for LibraryName", v.Expected.LibraryName, actual.LibraryName)
		}

	}
}

func TestParseLibraryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LibraryId
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
			Input: "https://endpoint-url.example.com/libraries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lIbRaRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/libraries/libraryName",
			Expected: &LibraryId{
				BaseURI:     "https://endpoint-url.example.com",
				LibraryName: "libraryName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/libraries/libraryName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lIbRaRiEs/lIbRaRyNaMe",
			Expected: &LibraryId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				LibraryName: "lIbRaRyNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lIbRaRiEs/lIbRaRyNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLibraryIDInsensitively(v.Input)
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

		if actual.LibraryName != v.Expected.LibraryName {
			t.Fatalf("Expected %q but got %q for LibraryName", v.Expected.LibraryName, actual.LibraryName)
		}

	}
}

func TestSegmentsForLibraryId(t *testing.T) {
	segments := LibraryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LibraryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
