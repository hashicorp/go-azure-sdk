package usercalendargroupcalendarcalendarviewattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarCalendarViewId{}

func TestNewUserCalendarGroupCalendarCalendarViewID(t *testing.T) {
	id := NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.CalendarGroupId != "calendarGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarGroupId'", id.CalendarGroupId, "calendarGroupIdValue")
	}

	if id.CalendarId != "calendarIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarId'", id.CalendarId, "calendarIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}
}

func TestFormatUserCalendarGroupCalendarCalendarViewID(t *testing.T) {
	actual := NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue").ID()
	expected := "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserCalendarGroupCalendarCalendarViewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarGroupCalendarCalendarViewId
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
			Input: "/users/userIdValue/calendarGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue",
			Expected: &UserCalendarGroupCalendarCalendarViewId{
				UserId:          "userIdValue",
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarGroupCalendarCalendarViewID(v.Input)
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

		if actual.CalendarGroupId != v.Expected.CalendarGroupId {
			t.Fatalf("Expected %q but got %q for CalendarGroupId", v.Expected.CalendarGroupId, actual.CalendarGroupId)
		}

		if actual.CalendarId != v.Expected.CalendarId {
			t.Fatalf("Expected %q but got %q for CalendarId", v.Expected.CalendarId, actual.CalendarId)
		}

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

	}
}

func TestParseUserCalendarGroupCalendarCalendarViewIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarGroupCalendarCalendarViewId
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
			Input: "/users/userIdValue/calendarGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue",
			Expected: &UserCalendarGroupCalendarCalendarViewId{
				UserId:          "userIdValue",
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Expected: &UserCalendarGroupCalendarCalendarViewId{
				UserId:          "uSeRiDvAlUe",
				CalendarGroupId: "cAlEnDaRgRoUpIdVaLuE",
				CalendarId:      "cAlEnDaRiDvAlUe",
				EventId:         "eVeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarGroupCalendarCalendarViewIDInsensitively(v.Input)
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

		if actual.CalendarGroupId != v.Expected.CalendarGroupId {
			t.Fatalf("Expected %q but got %q for CalendarGroupId", v.Expected.CalendarGroupId, actual.CalendarGroupId)
		}

		if actual.CalendarId != v.Expected.CalendarId {
			t.Fatalf("Expected %q but got %q for CalendarId", v.Expected.CalendarId, actual.CalendarId)
		}

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

	}
}

func TestSegmentsForUserCalendarGroupCalendarCalendarViewId(t *testing.T) {
	segments := UserCalendarGroupCalendarCalendarViewId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserCalendarGroupCalendarCalendarViewId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
