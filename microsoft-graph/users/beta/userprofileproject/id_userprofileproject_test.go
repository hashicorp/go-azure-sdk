package userprofileproject

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileProjectId{}

func TestNewUserProfileProjectID(t *testing.T) {
	id := NewUserProfileProjectID("userIdValue", "projectParticipationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ProjectParticipationId != "projectParticipationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectParticipationId'", id.ProjectParticipationId, "projectParticipationIdValue")
	}
}

func TestFormatUserProfileProjectID(t *testing.T) {
	actual := NewUserProfileProjectID("userIdValue", "projectParticipationIdValue").ID()
	expected := "/users/userIdValue/profile/projects/projectParticipationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserProfileProjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileProjectId
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
			Input: "/users/userIdValue/profile/projects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/projects/projectParticipationIdValue",
			Expected: &UserProfileProjectId{
				UserId:                 "userIdValue",
				ProjectParticipationId: "projectParticipationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/projects/projectParticipationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileProjectID(v.Input)
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

		if actual.ProjectParticipationId != v.Expected.ProjectParticipationId {
			t.Fatalf("Expected %q but got %q for ProjectParticipationId", v.Expected.ProjectParticipationId, actual.ProjectParticipationId)
		}

	}
}

func TestParseUserProfileProjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserProfileProjectId
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
			Input: "/users/userIdValue/profile/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pRoJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/profile/projects/projectParticipationIdValue",
			Expected: &UserProfileProjectId{
				UserId:                 "userIdValue",
				ProjectParticipationId: "projectParticipationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/profile/projects/projectParticipationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pRoJeCtS/pRoJeCtPaRtIcIpAtIoNiDvAlUe",
			Expected: &UserProfileProjectId{
				UserId:                 "uSeRiDvAlUe",
				ProjectParticipationId: "pRoJeCtPaRtIcIpAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pRoFiLe/pRoJeCtS/pRoJeCtPaRtIcIpAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserProfileProjectIDInsensitively(v.Input)
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

		if actual.ProjectParticipationId != v.Expected.ProjectParticipationId {
			t.Fatalf("Expected %q but got %q for ProjectParticipationId", v.Expected.ProjectParticipationId, actual.ProjectParticipationId)
		}

	}
}

func TestSegmentsForUserProfileProjectId(t *testing.T) {
	segments := UserProfileProjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserProfileProjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
