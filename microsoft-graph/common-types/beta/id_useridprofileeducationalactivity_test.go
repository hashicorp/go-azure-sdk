package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileEducationalActivityId{}

func TestNewUserIdProfileEducationalActivityID(t *testing.T) {
	id := NewUserIdProfileEducationalActivityID("userId", "educationalActivityId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.EducationalActivityId != "educationalActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'EducationalActivityId'", id.EducationalActivityId, "educationalActivityId")
	}
}

func TestFormatUserIdProfileEducationalActivityID(t *testing.T) {
	actual := NewUserIdProfileEducationalActivityID("userId", "educationalActivityId").ID()
	expected := "/users/userId/profile/educationalActivities/educationalActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileEducationalActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileEducationalActivityId
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
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/educationalActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/educationalActivities/educationalActivityId",
			Expected: &UserIdProfileEducationalActivityId{
				UserId:                "userId",
				EducationalActivityId: "educationalActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/educationalActivities/educationalActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileEducationalActivityID(v.Input)
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

		if actual.EducationalActivityId != v.Expected.EducationalActivityId {
			t.Fatalf("Expected %q but got %q for EducationalActivityId", v.Expected.EducationalActivityId, actual.EducationalActivityId)
		}

	}
}

func TestParseUserIdProfileEducationalActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileEducationalActivityId
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
			Input: "/users/userId/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/profile/educationalActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/eDuCaTiOnAlAcTiViTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/educationalActivities/educationalActivityId",
			Expected: &UserIdProfileEducationalActivityId{
				UserId:                "userId",
				EducationalActivityId: "educationalActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/educationalActivities/educationalActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/eDuCaTiOnAlAcTiViTiEs/eDuCaTiOnAlAcTiViTyId",
			Expected: &UserIdProfileEducationalActivityId{
				UserId:                "uSeRiD",
				EducationalActivityId: "eDuCaTiOnAlAcTiViTyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/eDuCaTiOnAlAcTiViTiEs/eDuCaTiOnAlAcTiViTyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileEducationalActivityIDInsensitively(v.Input)
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

		if actual.EducationalActivityId != v.Expected.EducationalActivityId {
			t.Fatalf("Expected %q but got %q for EducationalActivityId", v.Expected.EducationalActivityId, actual.EducationalActivityId)
		}

	}
}

func TestSegmentsForUserIdProfileEducationalActivityId(t *testing.T) {
	segments := UserIdProfileEducationalActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileEducationalActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
