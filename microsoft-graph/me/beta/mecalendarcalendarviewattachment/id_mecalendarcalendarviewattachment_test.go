package mecalendarcalendarviewattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewAttachmentId{}

func TestNewMeCalendarCalendarViewAttachmentID(t *testing.T) {
	id := NewMeCalendarCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatMeCalendarCalendarViewAttachmentID(t *testing.T) {
	actual := NewMeCalendarCalendarViewAttachmentID("eventIdValue", "attachmentIdValue").ID()
	expected := "/me/calendar/calendarView/eventIdValue/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarCalendarViewAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarCalendarViewAttachmentId
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
			Input: "/me/calendar",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarView/eventIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendar/calendarView/eventIdValue/attachments/attachmentIdValue",
			Expected: &MeCalendarCalendarViewAttachmentId{
				EventId:      "eventIdValue",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendar/calendarView/eventIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarCalendarViewAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseMeCalendarCalendarViewAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarCalendarViewAttachmentId
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
			Input: "/me/calendar",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendar/calendarView/eventIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendar/calendarView/eventIdValue/attachments/attachmentIdValue",
			Expected: &MeCalendarCalendarViewAttachmentId{
				EventId:      "eventIdValue",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendar/calendarView/eventIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &MeCalendarCalendarViewAttachmentId{
				EventId:      "eVeNtIdVaLuE",
				AttachmentId: "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaR/cAlEnDaRvIeW/eVeNtIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarCalendarViewAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForMeCalendarCalendarViewAttachmentId(t *testing.T) {
	segments := MeCalendarCalendarViewAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarCalendarViewAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
