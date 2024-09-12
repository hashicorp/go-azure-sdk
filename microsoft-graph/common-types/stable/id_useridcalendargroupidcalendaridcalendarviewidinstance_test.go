package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId{}

func TestNewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID(t *testing.T) {
	id := NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

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

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}
}

func TestFormatUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID(t *testing.T) {
	actual := NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value").ID()
	expected := "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances/eventId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId
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
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances/eventId1Value",
			Expected: &UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId{
				UserId:          "userIdValue",
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
				EventId1:        "eventId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances/eventId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID(v.Input)
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

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

	}
}

func TestParseUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId
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
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances/eventId1Value",
			Expected: &UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId{
				UserId:          "userIdValue",
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
				EventId1:        "eventId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/calendarView/eventIdValue/instances/eventId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe",
			Expected: &UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId{
				UserId:          "uSeRiDvAlUe",
				CalendarGroupId: "cAlEnDaRgRoUpIdVaLuE",
				CalendarId:      "cAlEnDaRiDvAlUe",
				EventId:         "eVeNtIdVaLuE",
				EventId1:        "eVeNtId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIDInsensitively(v.Input)
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

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

	}
}

func TestSegmentsForUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId(t *testing.T) {
	segments := UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
