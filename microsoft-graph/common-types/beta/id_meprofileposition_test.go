package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfilePositionId{}

func TestNewMeProfilePositionID(t *testing.T) {
	id := NewMeProfilePositionID("workPositionIdValue")

	if id.WorkPositionId != "workPositionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkPositionId'", id.WorkPositionId, "workPositionIdValue")
	}
}

func TestFormatMeProfilePositionID(t *testing.T) {
	actual := NewMeProfilePositionID("workPositionIdValue").ID()
	expected := "/me/profile/positions/workPositionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfilePositionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePositionId
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
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/positions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/positions/workPositionIdValue",
			Expected: &MeProfilePositionId{
				WorkPositionId: "workPositionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/positions/workPositionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePositionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WorkPositionId != v.Expected.WorkPositionId {
			t.Fatalf("Expected %q but got %q for WorkPositionId", v.Expected.WorkPositionId, actual.WorkPositionId)
		}

	}
}

func TestParseMeProfilePositionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfilePositionId
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
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/positions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pOsItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/positions/workPositionIdValue",
			Expected: &MeProfilePositionId{
				WorkPositionId: "workPositionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/positions/workPositionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pOsItIoNs/wOrKpOsItIoNiDvAlUe",
			Expected: &MeProfilePositionId{
				WorkPositionId: "wOrKpOsItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pOsItIoNs/wOrKpOsItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfilePositionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WorkPositionId != v.Expected.WorkPositionId {
			t.Fatalf("Expected %q but got %q for WorkPositionId", v.Expected.WorkPositionId, actual.WorkPositionId)
		}

	}
}

func TestSegmentsForMeProfilePositionId(t *testing.T) {
	segments := MeProfilePositionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfilePositionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
