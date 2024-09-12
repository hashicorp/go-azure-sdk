package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleShiftsRoleDefinitionId{}

func TestNewGroupIdTeamScheduleShiftsRoleDefinitionID(t *testing.T) {
	id := NewGroupIdTeamScheduleShiftsRoleDefinitionID("groupIdValue", "shiftsRoleDefinitionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.ShiftsRoleDefinitionId != "shiftsRoleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ShiftsRoleDefinitionId'", id.ShiftsRoleDefinitionId, "shiftsRoleDefinitionIdValue")
	}
}

func TestFormatGroupIdTeamScheduleShiftsRoleDefinitionID(t *testing.T) {
	actual := NewGroupIdTeamScheduleShiftsRoleDefinitionID("groupIdValue", "shiftsRoleDefinitionIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleShiftsRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleShiftsRoleDefinitionId
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
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionIdValue",
			Expected: &GroupIdTeamScheduleShiftsRoleDefinitionId{
				GroupId:                "groupIdValue",
				ShiftsRoleDefinitionId: "shiftsRoleDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleShiftsRoleDefinitionID(v.Input)
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

		if actual.ShiftsRoleDefinitionId != v.Expected.ShiftsRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for ShiftsRoleDefinitionId", v.Expected.ShiftsRoleDefinitionId, actual.ShiftsRoleDefinitionId)
		}

	}
}

func TestParseGroupIdTeamScheduleShiftsRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleShiftsRoleDefinitionId
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
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sHiFtSrOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionIdValue",
			Expected: &GroupIdTeamScheduleShiftsRoleDefinitionId{
				GroupId:                "groupIdValue",
				ShiftsRoleDefinitionId: "shiftsRoleDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/shiftsRoleDefinitions/shiftsRoleDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sHiFtSrOlEdEfInItIoNs/sHiFtSrOlEdEfInItIoNiDvAlUe",
			Expected: &GroupIdTeamScheduleShiftsRoleDefinitionId{
				GroupId:                "gRoUpIdVaLuE",
				ShiftsRoleDefinitionId: "sHiFtSrOlEdEfInItIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/sHiFtSrOlEdEfInItIoNs/sHiFtSrOlEdEfInItIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleShiftsRoleDefinitionIDInsensitively(v.Input)
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

		if actual.ShiftsRoleDefinitionId != v.Expected.ShiftsRoleDefinitionId {
			t.Fatalf("Expected %q but got %q for ShiftsRoleDefinitionId", v.Expected.ShiftsRoleDefinitionId, actual.ShiftsRoleDefinitionId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleShiftsRoleDefinitionId(t *testing.T) {
	segments := GroupIdTeamScheduleShiftsRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleShiftsRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
