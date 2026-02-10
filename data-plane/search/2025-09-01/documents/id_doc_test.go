package documents

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DocId{}

func TestNewDocID(t *testing.T) {
	id := NewDocID("https://endpoint-url.example.com", "docName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DocName != "docName" {
		t.Fatalf("Expected %q but got %q for Segment 'DocName'", id.DocName, "docName")
	}
}

func TestFormatDocID(t *testing.T) {
	actual := NewDocID("https://endpoint-url.example.com", "docName").ID()
	expected := "https://endpoint-url.example.com/docs/docName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDocID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DocId
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
			Input: "https://endpoint-url.example.com/docs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/docs/docName",
			Expected: &DocId{
				BaseURI: "https://endpoint-url.example.com",
				DocName: "docName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/docs/docName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDocID(v.Input)
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

		if actual.DocName != v.Expected.DocName {
			t.Fatalf("Expected %q but got %q for DocName", v.Expected.DocName, actual.DocName)
		}

	}
}

func TestParseDocIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DocId
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
			Input: "https://endpoint-url.example.com/docs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dOcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/docs/docName",
			Expected: &DocId{
				BaseURI: "https://endpoint-url.example.com",
				DocName: "docName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/docs/docName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dOcS/dOcNaMe",
			Expected: &DocId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DocName: "dOcNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dOcS/dOcNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDocIDInsensitively(v.Input)
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

		if actual.DocName != v.Expected.DocName {
			t.Fatalf("Expected %q but got %q for DocName", v.Expected.DocName, actual.DocName)
		}

	}
}

func TestSegmentsForDocId(t *testing.T) {
	segments := DocId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DocId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
