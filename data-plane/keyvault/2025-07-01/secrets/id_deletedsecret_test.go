package secrets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedsecretId{}

func TestNewDeletedsecretID(t *testing.T) {
	id := NewDeletedsecretID("https://endpoint-url.example.com", "deletedsecretName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeletedsecretName != "deletedsecretName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedsecretName'", id.DeletedsecretName, "deletedsecretName")
	}
}

func TestFormatDeletedsecretID(t *testing.T) {
	actual := NewDeletedsecretID("https://endpoint-url.example.com", "deletedsecretName").ID()
	expected := "https://endpoint-url.example.com/deletedsecrets/deletedsecretName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedsecretID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedsecretId
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
			Input: "https://endpoint-url.example.com/deletedsecrets",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedsecrets/deletedsecretName",
			Expected: &DeletedsecretId{
				BaseURI:           "https://endpoint-url.example.com",
				DeletedsecretName: "deletedsecretName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedsecrets/deletedsecretName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedsecretID(v.Input)
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

		if actual.DeletedsecretName != v.Expected.DeletedsecretName {
			t.Fatalf("Expected %q but got %q for DeletedsecretName", v.Expected.DeletedsecretName, actual.DeletedsecretName)
		}

	}
}

func TestParseDeletedsecretIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedsecretId
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
			Input: "https://endpoint-url.example.com/deletedsecrets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdSeCrEtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedsecrets/deletedsecretName",
			Expected: &DeletedsecretId{
				BaseURI:           "https://endpoint-url.example.com",
				DeletedsecretName: "deletedsecretName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedsecrets/deletedsecretName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdSeCrEtS/dElEtEdSeCrEtNaMe",
			Expected: &DeletedsecretId{
				BaseURI:           "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeletedsecretName: "dElEtEdSeCrEtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdSeCrEtS/dElEtEdSeCrEtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedsecretIDInsensitively(v.Input)
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

		if actual.DeletedsecretName != v.Expected.DeletedsecretName {
			t.Fatalf("Expected %q but got %q for DeletedsecretName", v.Expected.DeletedsecretName, actual.DeletedsecretName)
		}

	}
}

func TestSegmentsForDeletedsecretId(t *testing.T) {
	segments := DeletedsecretId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedsecretId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
