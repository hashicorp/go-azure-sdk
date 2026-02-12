package notebooks

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NotebookId{}

func TestNewNotebookID(t *testing.T) {
	id := NewNotebookID("https://endpoint-url.example.com", "notebookName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.NotebookName != "notebookName" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookName'", id.NotebookName, "notebookName")
	}
}

func TestFormatNotebookID(t *testing.T) {
	actual := NewNotebookID("https://endpoint-url.example.com", "notebookName").ID()
	expected := "https://endpoint-url.example.com/notebooks/notebookName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNotebookID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NotebookId
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
			Input: "https://endpoint-url.example.com/notebooks",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/notebooks/notebookName",
			Expected: &NotebookId{
				BaseURI:      "https://endpoint-url.example.com",
				NotebookName: "notebookName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/notebooks/notebookName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNotebookID(v.Input)
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

		if actual.NotebookName != v.Expected.NotebookName {
			t.Fatalf("Expected %q but got %q for NotebookName", v.Expected.NotebookName, actual.NotebookName)
		}

	}
}

func TestParseNotebookIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NotebookId
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
			Input: "https://endpoint-url.example.com/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/nOtEbOoKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/notebooks/notebookName",
			Expected: &NotebookId{
				BaseURI:      "https://endpoint-url.example.com",
				NotebookName: "notebookName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/notebooks/notebookName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/nOtEbOoKs/nOtEbOoKnAmE",
			Expected: &NotebookId{
				BaseURI:      "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				NotebookName: "nOtEbOoKnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/nOtEbOoKs/nOtEbOoKnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNotebookIDInsensitively(v.Input)
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

		if actual.NotebookName != v.Expected.NotebookName {
			t.Fatalf("Expected %q but got %q for NotebookName", v.Expected.NotebookName, actual.NotebookName)
		}

	}
}

func TestSegmentsForNotebookId(t *testing.T) {
	segments := NotebookId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NotebookId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
