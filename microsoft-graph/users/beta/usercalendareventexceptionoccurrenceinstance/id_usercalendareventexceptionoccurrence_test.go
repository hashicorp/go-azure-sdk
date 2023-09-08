package usercalendareventexceptionoccurrenceinstance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventExceptionOccurrenceId{}

func TestNewUserCalendarEventExceptionOccurrenceID(t *testing.T) {
	id := NewUserCalendarEventExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
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

func TestFormatUserCalendarEventExceptionOccurrenceID(t *testing.T) {
	actual := NewUserCalendarEventExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value").ID()
	expected := "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserCalendarEventExceptionOccurrenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarEventExceptionOccurrenceId
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
			Input: "/users/userIdValue/calendars/calendarIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value",
			Expected: &UserCalendarEventExceptionOccurrenceId{
				UserId:     "userIdValue",
				CalendarId: "calendarIdValue",
				EventId:    "eventIdValue",
				EventId1:   "eventId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarEventExceptionOccurrenceID(v.Input)
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

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

	}
}

func TestParseUserCalendarEventExceptionOccurrenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarEventExceptionOccurrenceId
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
			Input: "/users/userIdValue/calendars/calendarIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value",
			Expected: &UserCalendarEventExceptionOccurrenceId{
				UserId:     "userIdValue",
				CalendarId: "calendarIdValue",
				EventId:    "eventIdValue",
				EventId1:   "eventId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendars/calendarIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe",
			Expected: &UserCalendarEventExceptionOccurrenceId{
				UserId:     "uSeRiDvAlUe",
				CalendarId: "cAlEnDaRiDvAlUe",
				EventId:    "eVeNtIdVaLuE",
				EventId1:   "eVeNtId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarEventExceptionOccurrenceIDInsensitively(v.Input)
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

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

	}
}

func TestSegmentsForUserCalendarEventExceptionOccurrenceId(t *testing.T) {
	segments := UserCalendarEventExceptionOccurrenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserCalendarEventExceptionOccurrenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
