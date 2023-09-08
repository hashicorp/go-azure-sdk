package userprofileeducationalactivity

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileEducationalActivityId{}

func TestNewUserProfileEducationalActivityID(t *testing.T) {
	id := NewUserProfileEducationalActivityID("userIdValue", "educationalActivityIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.EducationalActivityId != "educationalActivityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EducationalActivityId'", id.EducationalActivityId, "educationalActivityIdValue")
	}
}

func TestFormatUserProfileEducationalActivityID(t *testing.T) {
	actual := NewUserProfileEducationalActivityID("userIdValue", "educationalActivityIdValue").ID()
	expected := "/users/userIdValue/profile/educationalActivities/educationalActivityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileEducationalActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileEducationalActivityId
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
			Input: "/users/userIdValue/profile/educationalActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/educationalActivities/educationalActivityIdValue",
			Expected: &UserProfileEducationalActivityId{
				UserId:                "userIdValue",
				EducationalActivityId: "educationalActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/educationalActivities/educationalActivityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileEducationalActivityID(v.Input)
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

func TestParseUserProfileEducationalActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileEducationalActivityId
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
			Input: "/users/userIdValue/profile/educationalActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/eDuCaTiOnAlAcTiViTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/educationalActivities/educationalActivityIdValue",
			Expected: &UserProfileEducationalActivityId{
				UserId:                "userIdValue",
				EducationalActivityId: "educationalActivityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/educationalActivities/educationalActivityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/eDuCaTiOnAlAcTiViTiEs/eDuCaTiOnAlAcTiViTyIdVaLuE",
			Expected: &UserProfileEducationalActivityId{
				UserId:                "uSeRiDvAlUe",
				EducationalActivityId: "eDuCaTiOnAlAcTiViTyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/eDuCaTiOnAlAcTiViTiEs/eDuCaTiOnAlAcTiViTyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileEducationalActivityIDInsensitively(v.Input)
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

func TestSegmentsForUserProfileEducationalActivityId(t *testing.T) {
	segments := UserProfileEducationalActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileEducationalActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
