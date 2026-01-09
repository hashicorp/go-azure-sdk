package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteResourceId{}

func TestNewMeOnenoteResourceID(t *testing.T) {
	id := NewMeOnenoteResourceID("onenoteResourceId")

	if id.OnenoteResourceId != "onenoteResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteResourceId'", id.OnenoteResourceId, "onenoteResourceId")
	}
}

func TestFormatMeOnenoteResourceID(t *testing.T) {
	actual := NewMeOnenoteResourceID("onenoteResourceId").ID()
	expected := "/me/onenote/resources/onenoteResourceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/resources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/resources/onenoteResourceId",
			Expected: &MeOnenoteResourceId{
				OnenoteResourceId: "onenoteResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/resources/onenoteResourceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteResourceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnenoteResourceId != v.Expected.OnenoteResourceId {
			t.Fatalf("Expected %q but got %q for OnenoteResourceId", v.Expected.OnenoteResourceId, actual.OnenoteResourceId)
		}

	}
}

func TestParseMeOnenoteResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onenote/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/rEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/resources/onenoteResourceId",
			Expected: &MeOnenoteResourceId{
				OnenoteResourceId: "onenoteResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/resources/onenoteResourceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeId",
			Expected: &MeOnenoteResourceId{
				OnenoteResourceId: "oNeNoTeReSoUrCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteResourceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnenoteResourceId != v.Expected.OnenoteResourceId {
			t.Fatalf("Expected %q but got %q for OnenoteResourceId", v.Expected.OnenoteResourceId, actual.OnenoteResourceId)
		}

	}
}

func TestSegmentsForMeOnenoteResourceId(t *testing.T) {
	segments := MeOnenoteResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
