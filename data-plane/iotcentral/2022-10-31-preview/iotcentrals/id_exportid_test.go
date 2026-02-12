package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExportIdId{}

func TestNewExportIdID(t *testing.T) {
	id := NewExportIdID("https://endpoint-url.example.com", "exportId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.ExportId != "exportId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExportId'", id.ExportId, "exportId")
	}
}

func TestFormatExportIdID(t *testing.T) {
	actual := NewExportIdID("https://endpoint-url.example.com", "exportId").ID()
	expected := "https://endpoint-url.example.com/dataExport/exports/exportId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExportIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExportIdId
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
			Input: "https://endpoint-url.example.com/dataExport/exports",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dataExport/exports/exportId",
			Expected: &ExportIdId{
				BaseURI:  "https://endpoint-url.example.com",
				ExportId: "exportId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dataExport/exports/exportId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExportIdID(v.Input)
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

		if actual.ExportId != v.Expected.ExportId {
			t.Fatalf("Expected %q but got %q for ExportId", v.Expected.ExportId, actual.ExportId)
		}

	}
}

func TestParseExportIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExportIdId
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
			Input: "https://endpoint-url.example.com/dataExport/exports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT/eXpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dataExport/exports/exportId",
			Expected: &ExportIdId{
				BaseURI:  "https://endpoint-url.example.com",
				ExportId: "exportId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dataExport/exports/exportId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT/eXpOrTs/eXpOrTiD",
			Expected: &ExportIdId{
				BaseURI:  "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				ExportId: "eXpOrTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAeXpOrT/eXpOrTs/eXpOrTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExportIdIDInsensitively(v.Input)
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

		if actual.ExportId != v.Expected.ExportId {
			t.Fatalf("Expected %q but got %q for ExportId", v.Expected.ExportId, actual.ExportId)
		}

	}
}

func TestSegmentsForExportIdId(t *testing.T) {
	segments := ExportIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExportIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
