package keys

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeletedkeyId{}

func TestNewDeletedkeyID(t *testing.T) {
	id := NewDeletedkeyID("https://endpoint-url.example.com", "deletedkeyName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeletedkeyName != "deletedkeyName" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedkeyName'", id.DeletedkeyName, "deletedkeyName")
	}
}

func TestFormatDeletedkeyID(t *testing.T) {
	actual := NewDeletedkeyID("https://endpoint-url.example.com", "deletedkeyName").ID()
	expected := "https://endpoint-url.example.com/deletedkeys/deletedkeyName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedkeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedkeyId
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
			Input: "https://endpoint-url.example.com/deletedkeys",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedkeys/deletedkeyName",
			Expected: &DeletedkeyId{
				BaseURI:        "https://endpoint-url.example.com",
				DeletedkeyName: "deletedkeyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedkeys/deletedkeyName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedkeyID(v.Input)
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

		if actual.DeletedkeyName != v.Expected.DeletedkeyName {
			t.Fatalf("Expected %q but got %q for DeletedkeyName", v.Expected.DeletedkeyName, actual.DeletedkeyName)
		}

	}
}

func TestParseDeletedkeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedkeyId
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
			Input: "https://endpoint-url.example.com/deletedkeys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdKeYs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/deletedkeys/deletedkeyName",
			Expected: &DeletedkeyId{
				BaseURI:        "https://endpoint-url.example.com",
				DeletedkeyName: "deletedkeyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/deletedkeys/deletedkeyName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdKeYs/dElEtEdKeYnAmE",
			Expected: &DeletedkeyId{
				BaseURI:        "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeletedkeyName: "dElEtEdKeYnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dElEtEdKeYs/dElEtEdKeYnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedkeyIDInsensitively(v.Input)
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

		if actual.DeletedkeyName != v.Expected.DeletedkeyName {
			t.Fatalf("Expected %q but got %q for DeletedkeyName", v.Expected.DeletedkeyName, actual.DeletedkeyName)
		}

	}
}

func TestSegmentsForDeletedkeyId(t *testing.T) {
	segments := DeletedkeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedkeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
