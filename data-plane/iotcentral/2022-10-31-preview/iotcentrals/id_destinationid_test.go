package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DestinationIdId{}

func TestNewDestinationIdID(t *testing.T) {
	id := NewDestinationIdID("https://endpoint-url.example.com", "destinationId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DestinationId != "destinationId" {
		t.Fatalf("Expected %q but got %q for Segment 'DestinationId'", id.DestinationId, "destinationId")
	}
}

func TestFormatDestinationIdID(t *testing.T) {
	actual := NewDestinationIdID("https://endpoint-url.example.com", "destinationId").ID()
	expected := "https://endpoint-url.example.com/dataExport/destinations/destinationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDestinationIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DestinationIdId
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
			Input: "https://endpoint-url.example.com/dataExport",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/dataExport/destinations",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dataExport/destinations/destinationId",
			Expected: &DestinationIdId{
				BaseURI:       "https://endpoint-url.example.com",
				DestinationId: "destinationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dataExport/destinations/destinationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDestinationIdID(v.Input)
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

		if actual.DestinationId != v.Expected.DestinationId {
			t.Fatalf("Expected %q but got %q for DestinationId", v.Expected.DestinationId, actual.DestinationId)
		}

	}
}

func TestParseDestinationIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DestinationIdId
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
			Input: "https://endpoint-url.example.com/dataExport",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/dataExport/destinations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT/dEsTiNaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dataExport/destinations/destinationId",
			Expected: &DestinationIdId{
				BaseURI:       "https://endpoint-url.example.com",
				DestinationId: "destinationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dataExport/destinations/destinationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT/dEsTiNaTiOnS/dEsTiNaTiOnId",
			Expected: &DestinationIdId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DestinationId: "dEsTiNaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT/dEsTiNaTiOnS/dEsTiNaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDestinationIdIDInsensitively(v.Input)
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

		if actual.DestinationId != v.Expected.DestinationId {
			t.Fatalf("Expected %q but got %q for DestinationId", v.Expected.DestinationId, actual.DestinationId)
		}

	}
}

func TestSegmentsForDestinationIdId(t *testing.T) {
	segments := DestinationIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DestinationIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
