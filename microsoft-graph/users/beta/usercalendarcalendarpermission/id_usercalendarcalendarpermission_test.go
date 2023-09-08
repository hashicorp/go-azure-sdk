package usercalendarcalendarpermission

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarPermissionId{}

func TestNewUserCalendarCalendarPermissionID(t *testing.T) {
	id := NewUserCalendarCalendarPermissionID("userIdValue", "calendarIdValue", "calendarPermissionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.CalendarId != "calendarIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarId'", id.CalendarId, "calendarIdValue")
	}

	if id.CalendarPermissionId != "calendarPermissionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarPermissionId'", id.CalendarPermissionId, "calendarPermissionIdValue")
	}
}

func TestFormatUserCalendarCalendarPermissionID(t *testing.T) {
	actual := NewUserCalendarCalendarPermissionID("userIdValue", "calendarIdValue", "calendarPermissionIdValue").ID()
	expected := "/users/userIdValue/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserCalendarCalendarPermissionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarCalendarPermissionId
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
			Input: "/users/userIdValue/calendars",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue/calendarPermissions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue",
			Expected: &UserCalendarCalendarPermissionId{
				UserId:               "userIdValue",
				CalendarId:           "calendarIdValue",
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarCalendarPermissionID(v.Input)
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

		if actual.CalendarId != v.Expected.CalendarId {
			t.Fatalf("Expected %q but got %q for CalendarId", v.Expected.CalendarId, actual.CalendarId)
		}

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestParseUserCalendarCalendarPermissionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarCalendarPermissionId
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
			Input: "/users/userIdValue/calendars",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue/calendarPermissions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRpErMiSsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue",
			Expected: &UserCalendarCalendarPermissionId{
				UserId:               "userIdValue",
				CalendarId:           "calendarIdValue",
				CalendarPermissionId: "calendarPermissionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendars/calendarIdValue/calendarPermissions/calendarPermissionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe",
			Expected: &UserCalendarCalendarPermissionId{
				UserId:               "uSeRiDvAlUe",
				CalendarId:           "cAlEnDaRiDvAlUe",
				CalendarPermissionId: "cAlEnDaRpErMiSsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRpErMiSsIoNs/cAlEnDaRpErMiSsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarCalendarPermissionIDInsensitively(v.Input)
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

		if actual.CalendarId != v.Expected.CalendarId {
			t.Fatalf("Expected %q but got %q for CalendarId", v.Expected.CalendarId, actual.CalendarId)
		}

		if actual.CalendarPermissionId != v.Expected.CalendarPermissionId {
			t.Fatalf("Expected %q but got %q for CalendarPermissionId", v.Expected.CalendarPermissionId, actual.CalendarPermissionId)
		}

	}
}

func TestSegmentsForUserCalendarCalendarPermissionId(t *testing.T) {
	segments := UserCalendarCalendarPermissionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserCalendarCalendarPermissionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
