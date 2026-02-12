package library

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LibraryOperationResultId{}

func TestNewLibraryOperationResultID(t *testing.T) {
	id := NewLibraryOperationResultID("https://endpoint-url.example.com", "operationId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.OperationId != "operationId" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationId'", id.OperationId, "operationId")
	}
}

func TestFormatLibraryOperationResultID(t *testing.T) {
	actual := NewLibraryOperationResultID("https://endpoint-url.example.com", "operationId").ID()
	expected := "https://endpoint-url.example.com/libraryOperationResults/operationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLibraryOperationResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LibraryOperationResultId
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
			Input: "https://endpoint-url.example.com/libraryOperationResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/libraryOperationResults/operationId",
			Expected: &LibraryOperationResultId{
				BaseURI:     "https://endpoint-url.example.com",
				OperationId: "operationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/libraryOperationResults/operationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLibraryOperationResultID(v.Input)
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

func TestParseLibraryOperationResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LibraryOperationResultId
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
			Input: "https://endpoint-url.example.com/libraryOperationResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lIbRaRyOpErAtIoNrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/libraryOperationResults/operationId",
			Expected: &LibraryOperationResultId{
				BaseURI:     "https://endpoint-url.example.com",
				OperationId: "operationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/libraryOperationResults/operationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lIbRaRyOpErAtIoNrEsUlTs/oPeRaTiOnId",
			Expected: &LibraryOperationResultId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				OperationId: "oPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lIbRaRyOpErAtIoNrEsUlTs/oPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLibraryOperationResultIDInsensitively(v.Input)
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

func TestSegmentsForLibraryOperationResultId(t *testing.T) {
	segments := LibraryOperationResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LibraryOperationResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
