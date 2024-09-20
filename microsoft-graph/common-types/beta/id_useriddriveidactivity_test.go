package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDriveIdActivityId{}

func TestNewUserIdDriveIdActivityID(t *testing.T) {
	id := NewUserIdDriveIdActivityID("userId", "driveId", "itemActivityOLDId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.DriveId != "driveId" {
		t.Fatalf("Expected %q but got %q for Segment 'DriveId'", id.DriveId, "driveId")
	}

	if id.ItemActivityOLDId != "itemActivityOLDId" {
		t.Fatalf("Expected %q but got %q for Segment 'ItemActivityOLDId'", id.ItemActivityOLDId, "itemActivityOLDId")
	}
}

func TestFormatUserIdDriveIdActivityID(t *testing.T) {
	actual := NewUserIdDriveIdActivityID("userId", "driveId", "itemActivityOLDId").ID()
	expected := "/users/userId/drives/driveId/activities/itemActivityOLDId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDriveIdActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdActivityId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/activities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/activities/itemActivityOLDId",
			Expected: &UserIdDriveIdActivityId{
				UserId:            "userId",
				DriveId:           "driveId",
				ItemActivityOLDId: "itemActivityOLDId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/activities/itemActivityOLDId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdActivityID(v.Input)
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

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestParseUserIdDriveIdActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDriveIdActivityId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/drives/driveId/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/aCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/drives/driveId/activities/itemActivityOLDId",
			Expected: &UserIdDriveIdActivityId{
				UserId:            "userId",
				DriveId:           "driveId",
				ItemActivityOLDId: "itemActivityOLDId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/drives/driveId/activities/itemActivityOLDId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/aCtIvItIeS/iTeMaCtIvItYoLdId",
			Expected: &UserIdDriveIdActivityId{
				UserId:            "uSeRiD",
				DriveId:           "dRiVeId",
				ItemActivityOLDId: "iTeMaCtIvItYoLdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dRiVeS/dRiVeId/aCtIvItIeS/iTeMaCtIvItYoLdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDriveIdActivityIDInsensitively(v.Input)
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

		if actual.ItemActivityOLDId != v.Expected.ItemActivityOLDId {
			t.Fatalf("Expected %q but got %q for ItemActivityOLDId", v.Expected.ItemActivityOLDId, actual.ItemActivityOLDId)
		}

	}
}

func TestSegmentsForUserIdDriveIdActivityId(t *testing.T) {
	segments := UserIdDriveIdActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDriveIdActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
