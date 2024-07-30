package calendarcalendarpermission

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarPermissionId{}

func TestNewMeCalendarIdCalendarPermissionID(t *testing.T) {
	id := NewMeCalendarIdCalendarPermissionID("calendarIdValue", "calendarPermissionIdValue")

	if id.CalendarId != "calendarIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarId'", id.CalendarId, "calendarIdValue")
	}

	if id.CalendarPermissionId != "calendarPermissionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarPermissionId'", id.CalendarPermissionId, "calendarPermissionIdValue")
	}
}

func TestFormatMeCalendarIdCalendarPermissionID(t *testing.T) {
	actual := NewMeCalendarIdCalendarPermissionID("calendarIdValue", "calendarPermissionIdValue").ID()
	expected := "/me/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarIdCalendarPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarIdCalendarPermissionId
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
			Input: "/me/calendars",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue/calendarPermissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue",
			Expected: &MeCalendarIdCalendarPermissionId{
				CalendarId:           "calendarIdValue",
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarIdCalendarPermissionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CalendarId != v.Expected.CalendarId {
			t.Fatalf("Expected %q but got %q for CalendarId", v.Expected.CalendarId, actual.CalendarId)
		}

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestParseMeCalendarIdCalendarPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarIdCalendarPermissionId
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
			Input: "/me/calendars",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue/calendarPermissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRpErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue",
			Expected: &MeCalendarIdCalendarPermissionId{
				CalendarId:           "calendarIdValue",
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe",
			Expected: &MeCalendarIdCalendarPermissionId{
				CalendarId:           "cAlEnDaRiDvAlUe",
				CalendarPermissionId: "cAlEnDaRpErMiSsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarIdCalendarPermissionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CalendarId != v.Expected.CalendarId {
			t.Fatalf("Expected %q but got %q for CalendarId", v.Expected.CalendarId, actual.CalendarId)
		}

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestSegmentsForMeCalendarIdCalendarPermissionId(t *testing.T) {
	segments := MeCalendarIdCalendarPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarIdCalendarPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
