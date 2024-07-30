package calendarcalendarviewinstance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewId{}

func TestNewMeCalendarIdCalendarViewID(t *testing.T) {
	id := NewMeCalendarIdCalendarViewID("calendarIdValue", "eventIdValue")

	if id.CalendarId != "calendarIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarId'", id.CalendarId, "calendarIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}
}

func TestFormatMeCalendarIdCalendarViewID(t *testing.T) {
	actual := NewMeCalendarIdCalendarViewID("calendarIdValue", "eventIdValue").ID()
	expected := "/me/calendars/calendarIdValue/calendarView/eventIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarIdCalendarViewID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarIdCalendarViewId
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
			Input: "/me/calendars/calendarIdValue/calendarView",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue",
			Expected: &MeCalendarIdCalendarViewId{
				CalendarId: "calendarIdValue",
				EventId:    "eventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarIdCalendarViewID(v.Input)
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

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

	}
}

func TestParseMeCalendarIdCalendarViewIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarIdCalendarViewId
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
			Input: "/me/calendars/calendarIdValue/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue",
			Expected: &MeCalendarIdCalendarViewId{
				CalendarId: "calendarIdValue",
				EventId:    "eventIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Expected: &MeCalendarIdCalendarViewId{
				CalendarId: "cAlEnDaRiDvAlUe",
				EventId:    "eVeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarIdCalendarViewIDInsensitively(v.Input)
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

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

	}
}

func TestSegmentsForMeCalendarIdCalendarViewId(t *testing.T) {
	segments := MeCalendarIdCalendarViewId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarIdCalendarViewId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
