package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarPermissionId{}

func TestNewGroupIdCalendarCalendarPermissionID(t *testing.T) {
	id := NewGroupIdCalendarCalendarPermissionID("groupIdValue", "calendarPermissionIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.CalendarPermissionId != "calendarPermissionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarPermissionId'", id.CalendarPermissionId, "calendarPermissionIdValue")
	}
}

func TestFormatGroupIdCalendarCalendarPermissionID(t *testing.T) {
	actual := NewGroupIdCalendarCalendarPermissionID("groupIdValue", "calendarPermissionIdValue").ID()
	expected := "/groups/groupIdValue/calendar/calendarPermissions/calendarPermissionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdCalendarCalendarPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarCalendarPermissionId
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
			Input: "/groups/groupIdValue/calendar",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendar/calendarPermissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/calendar/calendarPermissions/calendarPermissionIdValue",
			Expected: &GroupIdCalendarCalendarPermissionId{
				GroupId:              "groupIdValue",
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/calendar/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarCalendarPermissionID(v.Input)
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

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestParseGroupIdCalendarCalendarPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarCalendarPermissionId
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
			Input: "/groups/groupIdValue/calendar",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendar/calendarPermissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaR/cAlEnDaRpErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/calendar/calendarPermissions/calendarPermissionIdValue",
			Expected: &GroupIdCalendarCalendarPermissionId{
				GroupId:              "groupIdValue",
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/calendar/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaR/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe",
			Expected: &GroupIdCalendarCalendarPermissionId{
				GroupId:              "gRoUpIdVaLuE",
				CalendarPermissionId: "cAlEnDaRpErMiSsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaR/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarCalendarPermissionIDInsensitively(v.Input)
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

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestSegmentsForGroupIdCalendarCalendarPermissionId(t *testing.T) {
	segments := GroupIdCalendarCalendarPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdCalendarCalendarPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
