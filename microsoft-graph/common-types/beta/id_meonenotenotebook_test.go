package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteNotebookId{}

func TestNewMeOnenoteNotebookID(t *testing.T) {
	id := NewMeOnenoteNotebookID("notebookId")

	if id.NotebookId != "notebookId" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookId")
	}
}

func TestFormatMeOnenoteNotebookID(t *testing.T) {
	actual := NewMeOnenoteNotebookID("notebookId").ID()
	expected := "/me/onenote/notebooks/notebookId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteNotebookID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/notebooks/notebookId",
			Expected: &MeOnenoteNotebookId{
				NotebookId: "notebookId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

	}
}

func TestParseMeOnenoteNotebookIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteNotebookId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/notebooks/notebookId",
			Expected: &MeOnenoteNotebookId{
				NotebookId: "notebookId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/notebooks/notebookId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD",
			Expected: &MeOnenoteNotebookId{
				NotebookId: "nOtEbOoKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/nOtEbOoKs/nOtEbOoKiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteNotebookIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

	}
}

func TestSegmentsForMeOnenoteNotebookId(t *testing.T) {
	segments := MeOnenoteNotebookId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteNotebookId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
