package operationresult

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NotebookOperationResultId{}

func TestNewNotebookOperationResultID(t *testing.T) {
	id := NewNotebookOperationResultID("https://endpoint-url.example.com", "operationId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.OperationId != "operationId" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationId'", id.OperationId, "operationId")
	}
}

func TestFormatNotebookOperationResultID(t *testing.T) {
	actual := NewNotebookOperationResultID("https://endpoint-url.example.com", "operationId").ID()
	expected := "https://endpoint-url.example.com/notebookOperationResults/operationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNotebookOperationResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NotebookOperationResultId
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
			Input: "https://endpoint-url.example.com/notebookOperationResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/notebookOperationResults/operationId",
			Expected: &NotebookOperationResultId{
				BaseURI:     "https://endpoint-url.example.com",
				OperationId: "operationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/notebookOperationResults/operationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNotebookOperationResultID(v.Input)
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

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestParseNotebookOperationResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NotebookOperationResultId
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
			Input: "https://endpoint-url.example.com/notebookOperationResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/nOtEbOoKoPeRaTiOnReSuLtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/notebookOperationResults/operationId",
			Expected: &NotebookOperationResultId{
				BaseURI:     "https://endpoint-url.example.com",
				OperationId: "operationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/notebookOperationResults/operationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/nOtEbOoKoPeRaTiOnReSuLtS/oPeRaTiOnId",
			Expected: &NotebookOperationResultId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				OperationId: "oPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/nOtEbOoKoPeRaTiOnReSuLtS/oPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNotebookOperationResultIDInsensitively(v.Input)
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

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestSegmentsForNotebookOperationResultId(t *testing.T) {
	segments := NotebookOperationResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NotebookOperationResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
