package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeScopedRoleMemberOfId{}

func TestNewMeScopedRoleMemberOfID(t *testing.T) {
	id := NewMeScopedRoleMemberOfID("scopedRoleMembershipId")

	if id.ScopedRoleMembershipId != "scopedRoleMembershipId" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipId")
	}
}

func TestFormatMeScopedRoleMemberOfID(t *testing.T) {
	actual := NewMeScopedRoleMemberOfID("scopedRoleMembershipId").ID()
	expected := "/me/scopedRoleMemberOf/scopedRoleMembershipId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeScopedRoleMemberOfID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeScopedRoleMemberOfId
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
			Input: "/me/scopedRoleMemberOf",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipId",
			Expected: &MeScopedRoleMemberOfId{
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeScopedRoleMemberOfID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestParseMeScopedRoleMemberOfIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeScopedRoleMemberOfId
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
			Input: "/me/scopedRoleMemberOf",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sCoPeDrOlEmEmBeRoF",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipId",
			Expected: &MeScopedRoleMemberOfId{
				ScopedRoleMembershipId: "scopedRoleMembershipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sCoPeDrOlEmEmBeRoF/sCoPeDrOlEmEmBeRsHiPiD",
			Expected: &MeScopedRoleMemberOfId{
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sCoPeDrOlEmEmBeRoF/sCoPeDrOlEmEmBeRsHiPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeScopedRoleMemberOfIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScopedRoleMembershipId != v.Expected.ScopedRoleMembershipId {
			t.Fatalf("Expected %q but got %q for ScopedRoleMembershipId", v.Expected.ScopedRoleMembershipId, actual.ScopedRoleMembershipId)
		}

	}
}

func TestSegmentsForMeScopedRoleMemberOfId(t *testing.T) {
	segments := MeScopedRoleMemberOfId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeScopedRoleMemberOfId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
