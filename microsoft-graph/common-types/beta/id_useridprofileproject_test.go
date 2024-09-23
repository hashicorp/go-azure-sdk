package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileProjectId{}

func TestNewUserIdProfileProjectID(t *testing.T) {
	id := NewUserIdProfileProjectID("userId", "projectParticipationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ProjectParticipationId != "projectParticipationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectParticipationId'", id.ProjectParticipationId, "projectParticipationId")
	}
}

func TestFormatUserIdProfileProjectID(t *testing.T) {
	actual := NewUserIdProfileProjectID("userId", "projectParticipationId").ID()
	expected := "/users/userId/profile/projects/projectParticipationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdProfileProjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileProjectId
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
			Input: "/users/userId/profile/projects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/projects/projectParticipationId",
			Expected: &UserIdProfileProjectId{
				UserId:                 "userId",
				ProjectParticipationId: "projectParticipationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/projects/projectParticipationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileProjectID(v.Input)
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

func TestParseUserIdProfileProjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdProfileProjectId
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
			Input: "/users/userId/profile/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pRoJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/profile/projects/projectParticipationId",
			Expected: &UserIdProfileProjectId{
				UserId:                 "userId",
				ProjectParticipationId: "projectParticipationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/profile/projects/projectParticipationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pRoJeCtS/pRoJeCtPaRtIcIpAtIoNiD",
			Expected: &UserIdProfileProjectId{
				UserId:                 "uSeRiD",
				ProjectParticipationId: "pRoJeCtPaRtIcIpAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pRoFiLe/pRoJeCtS/pRoJeCtPaRtIcIpAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdProfileProjectIDInsensitively(v.Input)
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

func TestSegmentsForUserIdProfileProjectId(t *testing.T) {
	segments := UserIdProfileProjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdProfileProjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
