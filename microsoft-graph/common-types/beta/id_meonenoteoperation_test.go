package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnenoteOperationId{}

func TestNewMeOnenoteOperationID(t *testing.T) {
	id := NewMeOnenoteOperationID("onenoteOperationId")

	if id.OnenoteOperationId != "onenoteOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteOperationId'", id.OnenoteOperationId, "onenoteOperationId")
	}
}

func TestFormatMeOnenoteOperationID(t *testing.T) {
	actual := NewMeOnenoteOperationID("onenoteOperationId").ID()
	expected := "/me/onenote/operations/onenoteOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnenoteOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteOperationId
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
			Input: "/me/onenote/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/operations/onenoteOperationId",
			Expected: &MeOnenoteOperationId{
				OnenoteOperationId: "onenoteOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/operations/onenoteOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnenoteOperationId != v.Expected.OnenoteOperationId {
			t.Fatalf("Expected %q but got %q for OnenoteOperationId", v.Expected.OnenoteOperationId, actual.OnenoteOperationId)
		}

	}
}

func TestParseMeOnenoteOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnenoteOperationId
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
			Input: "/me/onenote/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onenote/operations/onenoteOperationId",
			Expected: &MeOnenoteOperationId{
				OnenoteOperationId: "onenoteOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onenote/operations/onenoteOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiD",
			Expected: &MeOnenoteOperationId{
				OnenoteOperationId: "oNeNoTeOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnenoteOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnenoteOperationId != v.Expected.OnenoteOperationId {
			t.Fatalf("Expected %q but got %q for OnenoteOperationId", v.Expected.OnenoteOperationId, actual.OnenoteOperationId)
		}

	}
}

func TestSegmentsForMeOnenoteOperationId(t *testing.T) {
	segments := MeOnenoteOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnenoteOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
