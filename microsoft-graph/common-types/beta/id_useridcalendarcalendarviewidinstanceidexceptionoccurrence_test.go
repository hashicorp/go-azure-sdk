package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{}

func TestNewUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(t *testing.T) {
	id := NewUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}

	if id.EventId2 != "eventId2Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId2'", id.EventId2, "eventId2Value")
	}
}

func TestFormatUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(t *testing.T) {
	actual := NewUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value").ID()
	expected := "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId
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
			Input: "/users/userIdValue/calendar",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value",
			Expected: &UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{
				UserId:   "userIdValue",
				EventId:  "eventIdValue",
				EventId1: "eventId1Value",
				EventId2: "eventId2Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID(v.Input)
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

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

	}
}

func TestParseUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId
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
			Input: "/users/userIdValue/calendar",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value",
			Expected: &UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{
				UserId:   "userIdValue",
				EventId:  "eventIdValue",
				EventId1: "eventId1Value",
				EventId2: "eventId2Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendar/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS/eVeNtId2vAlUe",
			Expected: &UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{
				UserId:   "uSeRiDvAlUe",
				EventId:  "eVeNtIdVaLuE",
				EventId1: "eVeNtId1vAlUe",
				EventId2: "eVeNtId2vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS/eVeNtId2vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively(v.Input)
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

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

	}
}

func TestSegmentsForUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId(t *testing.T) {
	segments := UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
