package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPhotoId{}

func TestNewUserIdPhotoID(t *testing.T) {
	id := NewUserIdPhotoID("userId", "profilePhotoId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ProfilePhotoId != "profilePhotoId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProfilePhotoId'", id.ProfilePhotoId, "profilePhotoId")
	}
}

func TestFormatUserIdPhotoID(t *testing.T) {
	actual := NewUserIdPhotoID("userId", "profilePhotoId").ID()
	expected := "/users/userId/photos/profilePhotoId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPhotoID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPhotoId
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
			Input: "/users/userId/photos",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/photos/profilePhotoId",
			Expected: &UserIdPhotoId{
				UserId:         "userId",
				ProfilePhotoId: "profilePhotoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/photos/profilePhotoId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPhotoID(v.Input)
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

		if actual.ProfilePhotoId != v.Expected.ProfilePhotoId {
			t.Fatalf("Expected %q but got %q for ProfilePhotoId", v.Expected.ProfilePhotoId, actual.ProfilePhotoId)
		}

	}
}

func TestParseUserIdPhotoIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPhotoId
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
			Input: "/users/userId/photos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pHoToS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/photos/profilePhotoId",
			Expected: &UserIdPhotoId{
				UserId:         "userId",
				ProfilePhotoId: "profilePhotoId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/photos/profilePhotoId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pHoToS/pRoFiLePhOtOiD",
			Expected: &UserIdPhotoId{
				UserId:         "uSeRiD",
				ProfilePhotoId: "pRoFiLePhOtOiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pHoToS/pRoFiLePhOtOiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPhotoIDInsensitively(v.Input)
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

		if actual.ProfilePhotoId != v.Expected.ProfilePhotoId {
			t.Fatalf("Expected %q but got %q for ProfilePhotoId", v.Expected.ProfilePhotoId, actual.ProfilePhotoId)
		}

	}
}

func TestSegmentsForUserIdPhotoId(t *testing.T) {
	segments := UserIdPhotoId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPhotoId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
