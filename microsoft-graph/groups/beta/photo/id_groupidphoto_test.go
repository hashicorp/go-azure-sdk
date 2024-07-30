package photo

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPhotoId{}

func TestNewGroupIdPhotoID(t *testing.T) {
	id := NewGroupIdPhotoID("groupIdValue", "profilePhotoIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ProfilePhotoId != "profilePhotoIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProfilePhotoId'", id.ProfilePhotoId, "profilePhotoIdValue")
	}
}

func TestFormatGroupIdPhotoID(t *testing.T) {
	actual := NewGroupIdPhotoID("groupIdValue", "profilePhotoIdValue").ID()
	expected := "/groups/groupIdValue/photos/profilePhotoIdValue"
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
			Expected: &GroupIdPhotoId{
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
			Expected: &GroupIdPhotoId{
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
			Expected: &GroupIdPhotoId{
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
