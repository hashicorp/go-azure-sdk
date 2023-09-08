package mescopedrolememberof

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeScopedRoleMemberOfId{}

func TestNewMeScopedRoleMemberOfID(t *testing.T) {
	id := NewMeScopedRoleMemberOfID("scopedRoleMembershipIdValue")

	if id.ScopedRoleMembershipId != "scopedRoleMembershipIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopedRoleMembershipId'", id.ScopedRoleMembershipId, "scopedRoleMembershipIdValue")
	}
}

func TestFormatMeScopedRoleMemberOfID(t *testing.T) {
	actual := NewMeScopedRoleMemberOfID("scopedRoleMembershipIdValue").ID()
	expected := "/me/scopedRoleMemberOf/scopedRoleMembershipIdValue"
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
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipIdValue",
			Expected: &MeScopedRoleMemberOfId{
				ScopedRoleMembershipId: "scopedRoleMembershipIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipIdValue/extra",
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
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipIdValue",
			Expected: &MeScopedRoleMemberOfId{
				ScopedRoleMembershipId: "scopedRoleMembershipIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/scopedRoleMemberOf/scopedRoleMembershipIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sCoPeDrOlEmEmBeRoF/sCoPeDrOlEmEmBeRsHiPiDvAlUe",
			Expected: &MeScopedRoleMemberOfId{
				ScopedRoleMembershipId: "sCoPeDrOlEmEmBeRsHiPiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sCoPeDrOlEmEmBeRoF/sCoPeDrOlEmEmBeRsHiPiDvAlUe/extra",
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
