package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePermissionGrantId{}

func TestNewMePermissionGrantID(t *testing.T) {
	id := NewMePermissionGrantID("resourceSpecificPermissionGrantId")

	if id.ResourceSpecificPermissionGrantId != "resourceSpecificPermissionGrantId" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceSpecificPermissionGrantId'", id.ResourceSpecificPermissionGrantId, "resourceSpecificPermissionGrantId")
	}
}

func TestFormatMePermissionGrantID(t *testing.T) {
	actual := NewMePermissionGrantID("resourceSpecificPermissionGrantId").ID()
	expected := "/me/permissionGrants/resourceSpecificPermissionGrantId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePermissionGrantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePermissionGrantId
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
			Input: "/me/permissionGrants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/permissionGrants/resourceSpecificPermissionGrantId",
			Expected: &MePermissionGrantId{
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/permissionGrants/resourceSpecificPermissionGrantId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePermissionGrantID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestParseMePermissionGrantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePermissionGrantId
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
			Input: "/me/permissionGrants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pErMiSsIoNgRaNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/permissionGrants/resourceSpecificPermissionGrantId",
			Expected: &MePermissionGrantId{
				ResourceSpecificPermissionGrantId: "resourceSpecificPermissionGrantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/permissionGrants/resourceSpecificPermissionGrantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId",
			Expected: &MePermissionGrantId{
				ResourceSpecificPermissionGrantId: "rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pErMiSsIoNgRaNtS/rEsOuRcEsPeCiFiCpErMiSsIoNgRaNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePermissionGrantIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceSpecificPermissionGrantId != v.Expected.ResourceSpecificPermissionGrantId {
			t.Fatalf("Expected %q but got %q for ResourceSpecificPermissionGrantId", v.Expected.ResourceSpecificPermissionGrantId, actual.ResourceSpecificPermissionGrantId)
		}

	}
}

func TestSegmentsForMePermissionGrantId(t *testing.T) {
	segments := MePermissionGrantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePermissionGrantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
