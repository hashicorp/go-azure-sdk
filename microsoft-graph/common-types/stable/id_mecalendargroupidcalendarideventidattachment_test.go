package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdEventIdAttachmentId{}

func TestNewMeCalendarGroupIdCalendarIdEventIdAttachmentID(t *testing.T) {
	id := NewMeCalendarGroupIdCalendarIdEventIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

	if id.CalendarGroupId != "calendarGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarGroupId'", id.CalendarGroupId, "calendarGroupIdValue")
	}

	if id.CalendarId != "calendarIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CalendarId'", id.CalendarId, "calendarIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatMeCalendarGroupIdCalendarIdEventIdAttachmentID(t *testing.T) {
	actual := NewMeCalendarGroupIdCalendarIdEventIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue").ID()
	expected := "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarGroupIdCalendarIdEventIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarGroupIdCalendarIdEventIdAttachmentId
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
			Input: "/me/calendarGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments/attachmentIdValue",
			Expected: &MeCalendarGroupIdCalendarIdEventIdAttachmentId{
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
				AttachmentId:    "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarGroupIdCalendarIdEventIdAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseMeCalendarGroupIdCalendarIdEventIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarGroupIdCalendarIdEventIdAttachmentId
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
			Input: "/me/calendarGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments/attachmentIdValue",
			Expected: &MeCalendarGroupIdCalendarIdEventIdAttachmentId{
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
				AttachmentId:    "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &MeCalendarGroupIdCalendarIdEventIdAttachmentId{
				CalendarGroupId: "cAlEnDaRgRoUpIdVaLuE",
				CalendarId:      "cAlEnDaRiDvAlUe",
				EventId:         "eVeNtIdVaLuE",
				AttachmentId:    "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarGroupIdCalendarIdEventIdAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForMeCalendarGroupIdCalendarIdEventIdAttachmentId(t *testing.T) {
	segments := MeCalendarGroupIdCalendarIdEventIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarGroupIdCalendarIdEventIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
