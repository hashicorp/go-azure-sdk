package dataflows

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DataflowId{}

func TestNewDataflowID(t *testing.T) {
	id := NewDataflowID("https://endpoint-url.example.com", "dataflowName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DataflowName != "dataflowName" {
		t.Fatalf("Expected %q but got %q for Segment 'DataflowName'", id.DataflowName, "dataflowName")
	}
}

func TestFormatDataflowID(t *testing.T) {
	actual := NewDataflowID("https://endpoint-url.example.com", "dataflowName").ID()
	expected := "https://endpoint-url.example.com/dataflows/dataflowName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDataflowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DataflowId
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
			Input: "https://endpoint-url.example.com/dataflows",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dataflows/dataflowName",
			Expected: &DataflowId{
				BaseURI:      "https://endpoint-url.example.com",
				DataflowName: "dataflowName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dataflows/dataflowName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDataflowID(v.Input)
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

		if actual.DataflowName != v.Expected.DataflowName {
			t.Fatalf("Expected %q but got %q for DataflowName", v.Expected.DataflowName, actual.DataflowName)
		}

	}
}

func TestParseDataflowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DataflowId
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
			Input: "https://endpoint-url.example.com/dataflows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAfLoWs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dataflows/dataflowName",
			Expected: &DataflowId{
				BaseURI:      "https://endpoint-url.example.com",
				DataflowName: "dataflowName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dataflows/dataflowName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAfLoWs/dAtAfLoWnAmE",
			Expected: &DataflowId{
				BaseURI:      "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DataflowName: "dAtAfLoWnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAfLoWs/dAtAfLoWnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDataflowIDInsensitively(v.Input)
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

		if actual.DataflowName != v.Expected.DataflowName {
			t.Fatalf("Expected %q but got %q for DataflowName", v.Expected.DataflowName, actual.DataflowName)
		}

	}
}

func TestSegmentsForDataflowId(t *testing.T) {
	segments := DataflowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DataflowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
