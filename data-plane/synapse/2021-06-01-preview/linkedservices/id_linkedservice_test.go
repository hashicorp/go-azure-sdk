package linkedservices

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LinkedServiceId{}

func TestNewLinkedServiceID(t *testing.T) {
	id := NewLinkedServiceID("https://endpoint-url.example.com", "linkedServiceName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.LinkedServiceName != "linkedServiceName" {
		t.Fatalf("Expected %q but got %q for Segment 'LinkedServiceName'", id.LinkedServiceName, "linkedServiceName")
	}
}

func TestFormatLinkedServiceID(t *testing.T) {
	actual := NewLinkedServiceID("https://endpoint-url.example.com", "linkedServiceName").ID()
	expected := "https://endpoint-url.example.com/linkedServices/linkedServiceName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLinkedServiceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LinkedServiceId
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
			Input: "https://endpoint-url.example.com/linkedServices",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/linkedServices/linkedServiceName",
			Expected: &LinkedServiceId{
				BaseURI:           "https://endpoint-url.example.com",
				LinkedServiceName: "linkedServiceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/linkedServices/linkedServiceName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLinkedServiceID(v.Input)
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

		if actual.LinkedServiceName != v.Expected.LinkedServiceName {
			t.Fatalf("Expected %q but got %q for LinkedServiceName", v.Expected.LinkedServiceName, actual.LinkedServiceName)
		}

	}
}

func TestParseLinkedServiceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LinkedServiceId
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
			Input: "https://endpoint-url.example.com/linkedServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lInKeDsErViCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/linkedServices/linkedServiceName",
			Expected: &LinkedServiceId{
				BaseURI:           "https://endpoint-url.example.com",
				LinkedServiceName: "linkedServiceName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/linkedServices/linkedServiceName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lInKeDsErViCeS/lInKeDsErViCeNaMe",
			Expected: &LinkedServiceId{
				BaseURI:           "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				LinkedServiceName: "lInKeDsErViCeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lInKeDsErViCeS/lInKeDsErViCeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLinkedServiceIDInsensitively(v.Input)
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

		if actual.LinkedServiceName != v.Expected.LinkedServiceName {
			t.Fatalf("Expected %q but got %q for LinkedServiceName", v.Expected.LinkedServiceName, actual.LinkedServiceName)
		}

	}
}

func TestSegmentsForLinkedServiceId(t *testing.T) {
	segments := LinkedServiceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LinkedServiceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
