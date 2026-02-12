package datasets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DatasetId{}

func TestNewDatasetID(t *testing.T) {
	id := NewDatasetID("https://endpoint-url.example.com", "datasetName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DatasetName != "datasetName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatasetName'", id.DatasetName, "datasetName")
	}
}

func TestFormatDatasetID(t *testing.T) {
	actual := NewDatasetID("https://endpoint-url.example.com", "datasetName").ID()
	expected := "https://endpoint-url.example.com/datasets/datasetName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDatasetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DatasetId
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
			Input: "https://endpoint-url.example.com/datasets",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/datasets/datasetName",
			Expected: &DatasetId{
				BaseURI:     "https://endpoint-url.example.com",
				DatasetName: "datasetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/datasets/datasetName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDatasetID(v.Input)
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

		if actual.DatasetName != v.Expected.DatasetName {
			t.Fatalf("Expected %q but got %q for DatasetName", v.Expected.DatasetName, actual.DatasetName)
		}

	}
}

func TestParseDatasetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DatasetId
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
			Input: "https://endpoint-url.example.com/datasets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAsEtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/datasets/datasetName",
			Expected: &DatasetId{
				BaseURI:     "https://endpoint-url.example.com",
				DatasetName: "datasetName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/datasets/datasetName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAsEtS/dAtAsEtNaMe",
			Expected: &DatasetId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DatasetName: "dAtAsEtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAtAsEtS/dAtAsEtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDatasetIDInsensitively(v.Input)
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

		if actual.DatasetName != v.Expected.DatasetName {
			t.Fatalf("Expected %q but got %q for DatasetName", v.Expected.DatasetName, actual.DatasetName)
		}

	}
}

func TestSegmentsForDatasetId(t *testing.T) {
	segments := DatasetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DatasetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
