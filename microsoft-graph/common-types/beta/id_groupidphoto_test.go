package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPhotoId{}

func TestNewGroupIdPhotoID(t *testing.T) {
	id := NewGroupIdPhotoID("groupId", "profilePhotoId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.ProfilePhotoId != "profilePhotoId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProfilePhotoId'", id.ProfilePhotoId, "profilePhotoId")
	}
}

func TestFormatGroupIdPhotoID(t *testing.T) {
	actual := NewGroupIdPhotoID("groupId", "profilePhotoId").ID()
	expected := "/groups/groupId/photos/profilePhotoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdPhotoID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPhotoId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/photos",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/photos/profilePhotoId",
			Expected: &GroupIdPhotoId{
				GroupId:        "groupId",
				ProfilePhotoId: "profilePhotoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/photos/profilePhotoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPhotoID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ProfilePhotoId != v.Expected.ProfilePhotoId {
			t.Fatalf("Expected %q but got %q for ProfilePhotoId", v.Expected.ProfilePhotoId, actual.ProfilePhotoId)
		}

	}
}

func TestParseGroupIdPhotoIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdPhotoId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/photos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pHoToS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/photos/profilePhotoId",
			Expected: &GroupIdPhotoId{
				GroupId:        "groupId",
				ProfilePhotoId: "profilePhotoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/photos/profilePhotoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pHoToS/pRoFiLePhOtOiD",
			Expected: &GroupIdPhotoId{
				GroupId:        "gRoUpId",
				ProfilePhotoId: "pRoFiLePhOtOiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/pHoToS/pRoFiLePhOtOiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdPhotoIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.ProfilePhotoId != v.Expected.ProfilePhotoId {
			t.Fatalf("Expected %q but got %q for ProfilePhotoId", v.Expected.ProfilePhotoId, actual.ProfilePhotoId)
		}

	}
}

func TestSegmentsForGroupIdPhotoId(t *testing.T) {
	segments := GroupIdPhotoId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdPhotoId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
