package sparkjobdefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SparkJobDefinitionId{}

func TestNewSparkJobDefinitionID(t *testing.T) {
	id := NewSparkJobDefinitionID("https://endpoint-url.example.com", "sparkJobDefinitionName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SparkJobDefinitionName != "sparkJobDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'SparkJobDefinitionName'", id.SparkJobDefinitionName, "sparkJobDefinitionName")
	}
}

func TestFormatSparkJobDefinitionID(t *testing.T) {
	actual := NewSparkJobDefinitionID("https://endpoint-url.example.com", "sparkJobDefinitionName").ID()
	expected := "https://endpoint-url.example.com/sparkJobDefinitions/sparkJobDefinitionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSparkJobDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SparkJobDefinitionId
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
			Input: "https://endpoint-url.example.com/sparkJobDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sparkJobDefinitions/sparkJobDefinitionName",
			Expected: &SparkJobDefinitionId{
				BaseURI:                "https://endpoint-url.example.com",
				SparkJobDefinitionName: "sparkJobDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sparkJobDefinitions/sparkJobDefinitionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSparkJobDefinitionID(v.Input)
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

		if actual.SparkJobDefinitionName != v.Expected.SparkJobDefinitionName {
			t.Fatalf("Expected %q but got %q for SparkJobDefinitionName", v.Expected.SparkJobDefinitionName, actual.SparkJobDefinitionName)
		}

	}
}

func TestParseSparkJobDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SparkJobDefinitionId
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
			Input: "https://endpoint-url.example.com/sparkJobDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sPaRkJoBdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sparkJobDefinitions/sparkJobDefinitionName",
			Expected: &SparkJobDefinitionId{
				BaseURI:                "https://endpoint-url.example.com",
				SparkJobDefinitionName: "sparkJobDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sparkJobDefinitions/sparkJobDefinitionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sPaRkJoBdEfInItIoNs/sPaRkJoBdEfInItIoNnAmE",
			Expected: &SparkJobDefinitionId{
				BaseURI:                "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SparkJobDefinitionName: "sPaRkJoBdEfInItIoNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sPaRkJoBdEfInItIoNs/sPaRkJoBdEfInItIoNnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSparkJobDefinitionIDInsensitively(v.Input)
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

		if actual.SparkJobDefinitionName != v.Expected.SparkJobDefinitionName {
			t.Fatalf("Expected %q but got %q for SparkJobDefinitionName", v.Expected.SparkJobDefinitionName, actual.SparkJobDefinitionName)
		}

	}
}

func TestSegmentsForSparkJobDefinitionId(t *testing.T) {
	segments := SparkJobDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SparkJobDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
