package calendarcalendarviewextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewIdExtensionId{}

func TestNewMeCalendarIdCalendarViewIdExtensionID(t *testing.T) {
	id := NewMeCalendarIdCalendarViewIdExtensionID("calendarIdValue", "eventIdValue", "extensionIdValue")

	if id.CalendarId != "calendarIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarId'", id.CalendarId, "calendarIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatMeCalendarIdCalendarViewIdExtensionID(t *testing.T) {
	actual := NewMeCalendarIdCalendarViewIdExtensionID("calendarIdValue", "eventIdValue", "extensionIdValue").ID()
	expected := "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarIdCalendarViewIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarIdCalendarViewIdExtensionId
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
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions/extensionIdValue",
			Expected: &MeCalendarIdCalendarViewIdExtensionId{
				CalendarId:  "calendarIdValue",
				EventId:     "eventIdValue",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarIdCalendarViewIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseMeCalendarIdCalendarViewIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarIdCalendarViewIdExtensionId
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
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions/extensionIdValue",
			Expected: &MeCalendarIdCalendarViewIdExtensionId{
				CalendarId:  "calendarIdValue",
				EventId:     "eventIdValue",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendars/calendarIdValue/calendarView/eventIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &MeCalendarIdCalendarViewIdExtensionId{
				CalendarId:  "cAlEnDaRiDvAlUe",
				EventId:     "eVeNtIdVaLuE",
				ExtensionId: "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRs/cAlEnDaRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarIdCalendarViewIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForMeCalendarIdCalendarViewIdExtensionId(t *testing.T) {
	segments := MeCalendarIdCalendarViewIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarIdCalendarViewIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
