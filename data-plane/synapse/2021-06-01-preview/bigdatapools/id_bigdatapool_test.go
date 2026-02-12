package bigdatapools

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BigDataPoolId{}

func TestNewBigDataPoolID(t *testing.T) {
	id := NewBigDataPoolID("https://endpoint-url.example.com", "bigDataPoolName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.BigDataPoolName != "bigDataPoolName" {
		t.Fatalf("Expected %q but got %q for Segment 'BigDataPoolName'", id.BigDataPoolName, "bigDataPoolName")
	}
}

func TestFormatBigDataPoolID(t *testing.T) {
	actual := NewBigDataPoolID("https://endpoint-url.example.com", "bigDataPoolName").ID()
	expected := "https://endpoint-url.example.com/bigDataPools/bigDataPoolName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBigDataPoolID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BigDataPoolId
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
			Input: "https://endpoint-url.example.com/bigDataPools",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/bigDataPools/bigDataPoolName",
			Expected: &BigDataPoolId{
				BaseURI:         "https://endpoint-url.example.com",
				BigDataPoolName: "bigDataPoolName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/bigDataPools/bigDataPoolName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBigDataPoolID(v.Input)
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

		if actual.BigDataPoolName != v.Expected.BigDataPoolName {
			t.Fatalf("Expected %q but got %q for BigDataPoolName", v.Expected.BigDataPoolName, actual.BigDataPoolName)
		}

	}
}

func TestParseBigDataPoolIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BigDataPoolId
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
			Input: "https://endpoint-url.example.com/bigDataPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/bIgDaTaPoOlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/bigDataPools/bigDataPoolName",
			Expected: &BigDataPoolId{
				BaseURI:         "https://endpoint-url.example.com",
				BigDataPoolName: "bigDataPoolName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/bigDataPools/bigDataPoolName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/bIgDaTaPoOlS/bIgDaTaPoOlNaMe",
			Expected: &BigDataPoolId{
				BaseURI:         "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				BigDataPoolName: "bIgDaTaPoOlNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/bIgDaTaPoOlS/bIgDaTaPoOlNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBigDataPoolIDInsensitively(v.Input)
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

		if actual.BigDataPoolName != v.Expected.BigDataPoolName {
			t.Fatalf("Expected %q but got %q for BigDataPoolName", v.Expected.BigDataPoolName, actual.BigDataPoolName)
		}

	}
}

func TestSegmentsForBigDataPoolId(t *testing.T) {
	segments := BigDataPoolId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BigDataPoolId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
