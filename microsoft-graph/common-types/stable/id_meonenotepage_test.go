package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenotePageId{}

func TestNewMeOnenotePageID(t *testing.T) {
	id := NewMeOnenotePageID("onenotePageId")

	if id.OnenotePageId != "onenotePageId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenotePageId'", id.OnenotePageId, "onenotePageId")
	}
}

func TestFormatMeOnenotePageID(t *testing.T) {
	actual := NewMeOnenotePageID("onenotePageId").ID()
	expected := "/me/onenote/pages/onenotePageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenotePageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenotePageId
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
			Input: "/me/onenote/pages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/pages/onenotePageId",
			Expected: &MeOnenotePageId{
				OnenotePageId: "onenotePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/pages/onenotePageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenotePageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestParseMeOnenotePageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenotePageId
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
			Input: "/me/onenote/pages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/pAgEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/pages/onenotePageId",
			Expected: &MeOnenotePageId{
				OnenotePageId: "onenotePageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/pages/onenotePageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/pAgEs/oNeNoTePaGeId",
			Expected: &MeOnenotePageId{
				OnenotePageId: "oNeNoTePaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/pAgEs/oNeNoTePaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenotePageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnenotePageId != v.Expected.OnenotePageId {
			t.Fatalf("Expected %q but got %q for OnenotePageId", v.Expected.OnenotePageId, actual.OnenotePageId)
		}

	}
}

func TestSegmentsForMeOnenotePageId(t *testing.T) {
	segments := MeOnenotePageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenotePageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
