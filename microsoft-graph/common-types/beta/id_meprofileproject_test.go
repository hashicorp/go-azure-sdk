package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeProfileProjectId{}

func TestNewMeProfileProjectID(t *testing.T) {
	id := NewMeProfileProjectID("projectParticipationId")

	if id.ProjectParticipationId != "projectParticipationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectParticipationId'", id.ProjectParticipationId, "projectParticipationId")
	}
}

func TestFormatMeProfileProjectID(t *testing.T) {
	actual := NewMeProfileProjectID("projectParticipationId").ID()
	expected := "/me/profile/projects/projectParticipationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeProfileProjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileProjectId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/projects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/projects/projectParticipationId",
			Expected: &MeProfileProjectId{
				ProjectParticipationId: "projectParticipationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/projects/projectParticipationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileProjectID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProjectParticipationId != v.Expected.ProjectParticipationId {
			t.Fatalf("Expected %q but got %q for ProjectParticipationId", v.Expected.ProjectParticipationId, actual.ProjectParticipationId)
		}

	}
}

func TestParseMeProfileProjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeProfileProjectId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/profile/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pRoJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/profile/projects/projectParticipationId",
			Expected: &MeProfileProjectId{
				ProjectParticipationId: "projectParticipationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/profile/projects/projectParticipationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pRoJeCtS/pRoJeCtPaRtIcIpAtIoNiD",
			Expected: &MeProfileProjectId{
				ProjectParticipationId: "pRoJeCtPaRtIcIpAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pRoFiLe/pRoJeCtS/pRoJeCtPaRtIcIpAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeProfileProjectIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProjectParticipationId != v.Expected.ProjectParticipationId {
			t.Fatalf("Expected %q but got %q for ProjectParticipationId", v.Expected.ProjectParticipationId, actual.ProjectParticipationId)
		}

	}
}

func TestSegmentsForMeProfileProjectId(t *testing.T) {
	segments := MeProfileProjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeProfileProjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
