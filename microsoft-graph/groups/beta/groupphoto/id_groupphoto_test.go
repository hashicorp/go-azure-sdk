package groupphoto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupPhotoId{}

func TestNewGroupPhotoID(t *testing.T) {
	id := NewGroupPhotoID("groupIdValue", "profilePhotoIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ProfilePhotoId != "profilePhotoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProfilePhotoId'", id.ProfilePhotoId, "profilePhotoIdValue")
	}
}

func TestFormatGroupPhotoID(t *testing.T) {
	actual := NewGroupPhotoID("groupIdValue", "profilePhotoIdValue").ID()
	expected := "/groups/groupIdValue/photos/profilePhotoIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupPhotoID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupPhotoId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/photos",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/photos/profilePhotoIdValue",
			Expected: &GroupPhotoId{
				GroupId:        "groupIdValue",
				ProfilePhotoId: "profilePhotoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/photos/profilePhotoIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupPhotoID(v.Input)
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

func TestParseGroupPhotoIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupPhotoId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/photos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pHoToS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/photos/profilePhotoIdValue",
			Expected: &GroupPhotoId{
				GroupId:        "groupIdValue",
				ProfilePhotoId: "profilePhotoIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/photos/profilePhotoIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pHoToS/pRoFiLePhOtOiDvAlUe",
			Expected: &GroupPhotoId{
				GroupId:        "gRoUpIdVaLuE",
				ProfilePhotoId: "pRoFiLePhOtOiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/pHoToS/pRoFiLePhOtOiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupPhotoIDInsensitively(v.Input)
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

func TestSegmentsForGroupPhotoId(t *testing.T) {
	segments := GroupPhotoId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupPhotoId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
