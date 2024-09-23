package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{}

func TestNewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(t *testing.T) {
	id := NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.EventId != "eventId" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventId")
	}

	if id.EventId1 != "eventId1" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1")
	}

	if id.AttachmentId != "attachmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentId")
	}
}

func TestFormatGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(t *testing.T) {
	actual := NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId").ID()
	expected := "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments/attachmentId",
			Expected: &GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{
				GroupId:      "groupId",
				EventId:      "eventId",
				EventId1:     "eventId1",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
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

func TestParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments/attachmentId",
			Expected: &GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{
				GroupId:      "groupId",
				EventId:      "eventId",
				EventId1:     "eventId1",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendar/calendarView/eventId/exceptionOccurrences/eventId1/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{
				GroupId:      "gRoUpId",
				EventId:      "eVeNtId",
				EventId1:     "eVeNtId1",
				AttachmentId: "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
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

func TestSegmentsForGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId(t *testing.T) {
	segments := GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
