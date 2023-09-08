package mecalendargroupcalendareventinstanceattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarEventInstanceAttachmentId{}

func TestNewMeCalendarGroupCalendarEventInstanceAttachmentID(t *testing.T) {
	id := NewMeCalendarGroupCalendarEventInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

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

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatMeCalendarGroupCalendarEventInstanceAttachmentID(t *testing.T) {
	actual := NewMeCalendarGroupCalendarEventInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue").ID()
	expected := "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarGroupCalendarEventInstanceAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarGroupCalendarEventInstanceAttachmentId
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
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue",
			Expected: &MeCalendarGroupCalendarEventInstanceAttachmentId{
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
				EventId1:        "eventId1Value",
				AttachmentId:    "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarGroupCalendarEventInstanceAttachmentID(v.Input)
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

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseMeCalendarGroupCalendarEventInstanceAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarGroupCalendarEventInstanceAttachmentId
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
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue",
			Expected: &MeCalendarGroupCalendarEventInstanceAttachmentId{
				CalendarGroupId: "calendarGroupIdValue",
				CalendarId:      "calendarIdValue",
				EventId:         "eventIdValue",
				EventId1:        "eventId1Value",
				AttachmentId:    "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendarGroups/calendarGroupIdValue/calendars/calendarIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &MeCalendarGroupCalendarEventInstanceAttachmentId{
				CalendarGroupId: "cAlEnDaRgRoUpIdVaLuE",
				CalendarId:      "cAlEnDaRiDvAlUe",
				EventId:         "eVeNtIdVaLuE",
				EventId1:        "eVeNtId1vAlUe",
				AttachmentId:    "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRgRoUpS/cAlEnDaRgRoUpIdVaLuE/cAlEnDaRs/cAlEnDaRiDvAlUe/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarGroupCalendarEventInstanceAttachmentIDInsensitively(v.Input)
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

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForMeCalendarGroupCalendarEventInstanceAttachmentId(t *testing.T) {
	segments := MeCalendarGroupCalendarEventInstanceAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarGroupCalendarEventInstanceAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
