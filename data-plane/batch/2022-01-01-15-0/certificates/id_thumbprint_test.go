package certificates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ThumbprintId{}

func TestNewThumbprintID(t *testing.T) {
	id := NewThumbprintID("thumbprintAlgorithmName", "thumbprintName")

	if id.ThumbprintAlgorithmName != "thumbprintAlgorithmName" {
		t.Fatalf("Expected %q but got %q for Segment 'ThumbprintAlgorithmName'", id.ThumbprintAlgorithmName, "thumbprintAlgorithmName")
	}

	if id.ThumbprintName != "thumbprintName" {
		t.Fatalf("Expected %q but got %q for Segment 'ThumbprintName'", id.ThumbprintName, "thumbprintName")
	}
}

func TestFormatThumbprintID(t *testing.T) {
	actual := NewThumbprintID("thumbprintAlgorithmName", "thumbprintName").ID()
	expected := "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint/thumbprintName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseThumbprintID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ThumbprintId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/thumbprintAlgorithm",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint",
			Error: true,
		},
		{
			// Valid URI
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint/thumbprintName",
			Expected: &ThumbprintId{
				ThumbprintAlgorithmName: "thumbprintAlgorithmName",
				ThumbprintName:          "thumbprintName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint/thumbprintName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseThumbprintID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ThumbprintAlgorithmName != v.Expected.ThumbprintAlgorithmName {
			t.Fatalf("Expected %q but got %q for ThumbprintAlgorithmName", v.Expected.ThumbprintAlgorithmName, actual.ThumbprintAlgorithmName)
		}

		if actual.ThumbprintName != v.Expected.ThumbprintName {
			t.Fatalf("Expected %q but got %q for ThumbprintName", v.Expected.ThumbprintName, actual.ThumbprintName)
		}

	}
}

func TestParseThumbprintIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ThumbprintId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/thumbprintAlgorithm",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/tHuMbPrInTaLgOrItHm",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/tHuMbPrInTaLgOrItHm/tHuMbPrInTaLgOrItHmNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/tHuMbPrInTaLgOrItHm/tHuMbPrInTaLgOrItHmNaMe/tHuMbPrInT",
			Error: true,
		},
		{
			// Valid URI
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint/thumbprintName",
			Expected: &ThumbprintId{
				ThumbprintAlgorithmName: "thumbprintAlgorithmName",
				ThumbprintName:          "thumbprintName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/thumbprintAlgorithm/thumbprintAlgorithmName/thumbprint/thumbprintName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/tHuMbPrInTaLgOrItHm/tHuMbPrInTaLgOrItHmNaMe/tHuMbPrInT/tHuMbPrInTnAmE",
			Expected: &ThumbprintId{
				ThumbprintAlgorithmName: "tHuMbPrInTaLgOrItHmNaMe",
				ThumbprintName:          "tHuMbPrInTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/tHuMbPrInTaLgOrItHm/tHuMbPrInTaLgOrItHmNaMe/tHuMbPrInT/tHuMbPrInTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseThumbprintIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ThumbprintAlgorithmName != v.Expected.ThumbprintAlgorithmName {
			t.Fatalf("Expected %q but got %q for ThumbprintAlgorithmName", v.Expected.ThumbprintAlgorithmName, actual.ThumbprintAlgorithmName)
		}

		if actual.ThumbprintName != v.Expected.ThumbprintName {
			t.Fatalf("Expected %q but got %q for ThumbprintName", v.Expected.ThumbprintName, actual.ThumbprintName)
		}

	}
}

func TestSegmentsForThumbprintId(t *testing.T) {
	segments := ThumbprintId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ThumbprintId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
