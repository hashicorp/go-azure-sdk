package mecalendarcalendarpermission

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarPermissionId{}

func TestNewMeCalendarCalendarPermissionID(t *testing.T) {
	id := NewMeCalendarCalendarPermissionID("calendarPermissionIdValue")

	if id.CalendarPermissionId != "calendarPermissionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarPermissionId'", id.CalendarPermissionId, "calendarPermissionIdValue")
	}
}

func TestFormatMeCalendarCalendarPermissionID(t *testing.T) {
	actual := NewMeCalendarCalendarPermissionID("calendarPermissionIdValue").ID()
	expected := "/me/calendar/calendarPermissions/calendarPermissionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarCalendarPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarCalendarPermissionId
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
			Input: "/me/calendar",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarPermissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendar/calendarPermissions/calendarPermissionIdValue",
			Expected: &MeCalendarCalendarPermissionId{
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendar/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarCalendarPermissionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestParseMeCalendarCalendarPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarCalendarPermissionId
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
			Input: "/me/calendar",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarPermissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRpErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendar/calendarPermissions/calendarPermissionIdValue",
			Expected: &MeCalendarCalendarPermissionId{
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendar/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe",
			Expected: &MeCalendarCalendarPermissionId{
				CalendarPermissionId: "cAlEnDaRpErMiSsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarCalendarPermissionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestSegmentsForMeCalendarCalendarPermissionId(t *testing.T) {
	segments := MeCalendarCalendarPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarCalendarPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
