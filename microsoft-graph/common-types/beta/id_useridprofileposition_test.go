package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePositionId{}

func TestNewUserIdProfilePositionID(t *testing.T) {
	id := NewUserIdProfilePositionID("userIdValue", "workPositionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.WorkPositionId != "workPositionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkPositionId'", id.WorkPositionId, "workPositionIdValue")
	}
}

func TestFormatUserIdProfilePositionID(t *testing.T) {
	actual := NewUserIdProfilePositionID("userIdValue", "workPositionIdValue").ID()
	expected := "/users/userIdValue/profile/positions/workPositionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfilePositionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePositionId
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
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/positions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/positions/workPositionIdValue",
			Expected: &UserIdProfilePositionId{
				UserId:         "userIdValue",
				WorkPositionId: "workPositionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/positions/workPositionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePositionID(v.Input)
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

		if actual.WorkPositionId != v.Expected.WorkPositionId {
			t.Fatalf("Expected %q but got %q for WorkPositionId", v.Expected.WorkPositionId, actual.WorkPositionId)
		}

	}
}

func TestParseUserIdProfilePositionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfilePositionId
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
			Input: "/users/userIdValue/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/profile/positions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pOsItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/positions/workPositionIdValue",
			Expected: &UserIdProfilePositionId{
				UserId:         "userIdValue",
				WorkPositionId: "workPositionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/positions/workPositionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pOsItIoNs/wOrKpOsItIoNiDvAlUe",
			Expected: &UserIdProfilePositionId{
				UserId:         "uSeRiDvAlUe",
				WorkPositionId: "wOrKpOsItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pOsItIoNs/wOrKpOsItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfilePositionIDInsensitively(v.Input)
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

		if actual.WorkPositionId != v.Expected.WorkPositionId {
			t.Fatalf("Expected %q but got %q for WorkPositionId", v.Expected.WorkPositionId, actual.WorkPositionId)
		}

	}
}

func TestSegmentsForUserIdProfilePositionId(t *testing.T) {
	segments := UserIdProfilePositionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfilePositionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
