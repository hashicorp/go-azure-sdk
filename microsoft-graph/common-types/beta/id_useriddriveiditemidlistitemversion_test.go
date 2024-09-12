package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdItemIdListItemVersionId{}

func TestNewUserIdDriveIdItemIdListItemVersionID(t *testing.T) {
	id := NewUserIdDriveIdItemIdListItemVersionID("userIdValue", "driveIdValue", "driveItemIdValue", "listItemVersionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.DriveId != "driveIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveIdValue")
	}

	if id.DriveItemId != "driveItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveItemId'", id.DriveItemId, "driveItemIdValue")
	}

	if id.ListItemVersionId != "listItemVersionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ListItemVersionId'", id.ListItemVersionId, "listItemVersionIdValue")
	}
}

func TestFormatUserIdDriveIdItemIdListItemVersionID(t *testing.T) {
	actual := NewUserIdDriveIdItemIdListItemVersionID("userIdValue", "driveIdValue", "driveItemIdValue", "listItemVersionIdValue").ID()
	expected := "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions/listItemVersionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdItemIdListItemVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions/listItemVersionIdValue",
			Expected: &UserIdDriveIdItemIdListItemVersionId{
				UserId:            "userIdValue",
				DriveId:           "driveIdValue",
				DriveItemId:       "driveItemIdValue",
				ListItemVersionId: "listItemVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions/listItemVersionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestParseUserIdDriveIdItemIdListItemVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdItemIdListItemVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/iTeMs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/lIsTiTeM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/lIsTiTeM/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions/listItemVersionIdValue",
			Expected: &UserIdDriveIdItemIdListItemVersionId{
				UserId:            "userIdValue",
				DriveId:           "driveIdValue",
				DriveItemId:       "driveItemIdValue",
				ListItemVersionId: "listItemVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/drives/driveIdValue/items/driveItemIdValue/listItem/versions/listItemVersionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/lIsTiTeM/vErSiOnS/lIsTiTeMvErSiOnIdVaLuE",
			Expected: &UserIdDriveIdItemIdListItemVersionId{
				UserId:            "uSeRiDvAlUe",
				DriveId:           "dRiVeIdVaLuE",
				DriveItemId:       "dRiVeItEmIdVaLuE",
				ListItemVersionId: "lIsTiTeMvErSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/dRiVeS/dRiVeIdVaLuE/iTeMs/dRiVeItEmIdVaLuE/lIsTiTeM/vErSiOnS/lIsTiTeMvErSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdItemIdListItemVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.DriveId != v.Expected.DriveId {
			t.Fatalf("Expected %q but got %q for DriveId", v.Expected.DriveId, actual.DriveId)
		}

		if actual.DriveItemId != v.Expected.DriveItemId {
			t.Fatalf("Expected %q but got %q for DriveItemId", v.Expected.DriveItemId, actual.DriveItemId)
		}

		if actual.ListItemVersionId != v.Expected.ListItemVersionId {
			t.Fatalf("Expected %q but got %q for ListItemVersionId", v.Expected.ListItemVersionId, actual.ListItemVersionId)
		}

	}
}

func TestSegmentsForUserIdDriveIdItemIdListItemVersionId(t *testing.T) {
	segments := UserIdDriveIdItemIdListItemVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdItemIdListItemVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
