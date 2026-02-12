package sparkconfigurations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SparkConfigurationId{}

func TestNewSparkConfigurationID(t *testing.T) {
	id := NewSparkConfigurationID("https://endpoint-url.example.com", "sparkConfigurationName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SparkConfigurationName != "sparkConfigurationName" {
		t.Fatalf("Expected %q but got %q for Segment 'SparkConfigurationName'", id.SparkConfigurationName, "sparkConfigurationName")
	}
}

func TestFormatSparkConfigurationID(t *testing.T) {
	actual := NewSparkConfigurationID("https://endpoint-url.example.com", "sparkConfigurationName").ID()
	expected := "https://endpoint-url.example.com/sparkConfigurations/sparkConfigurationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSparkConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SparkConfigurationId
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
			Input: "https://endpoint-url.example.com/sparkConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sparkConfigurations/sparkConfigurationName",
			Expected: &SparkConfigurationId{
				BaseURI:                "https://endpoint-url.example.com",
				SparkConfigurationName: "sparkConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sparkConfigurations/sparkConfigurationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSparkConfigurationID(v.Input)
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

		if actual.SparkConfigurationName != v.Expected.SparkConfigurationName {
			t.Fatalf("Expected %q but got %q for SparkConfigurationName", v.Expected.SparkConfigurationName, actual.SparkConfigurationName)
		}

	}
}

func TestParseSparkConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SparkConfigurationId
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
			Input: "https://endpoint-url.example.com/sparkConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sPaRkCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/sparkConfigurations/sparkConfigurationName",
			Expected: &SparkConfigurationId{
				BaseURI:                "https://endpoint-url.example.com",
				SparkConfigurationName: "sparkConfigurationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/sparkConfigurations/sparkConfigurationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sPaRkCoNfIgUrAtIoNs/sPaRkCoNfIgUrAtIoNnAmE",
			Expected: &SparkConfigurationId{
				BaseURI:                "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SparkConfigurationName: "sPaRkCoNfIgUrAtIoNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sPaRkCoNfIgUrAtIoNs/sPaRkCoNfIgUrAtIoNnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSparkConfigurationIDInsensitively(v.Input)
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

		if actual.SparkConfigurationName != v.Expected.SparkConfigurationName {
			t.Fatalf("Expected %q but got %q for SparkConfigurationName", v.Expected.SparkConfigurationName, actual.SparkConfigurationName)
		}

	}
}

func TestSegmentsForSparkConfigurationId(t *testing.T) {
	segments := SparkConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SparkConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
