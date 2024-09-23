package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}

func TestNewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(t *testing.T) {
	id := NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.EventId != "eventId" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventId")
	}

	if id.EventId1 != "eventId1" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1")
	}

	if id.EventId2 != "eventId2" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId2'", id.EventId2, "eventId2")
	}

	if id.AttachmentId != "attachmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentId")
	}
}

func TestFormatGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(t *testing.T) {
	actual := NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId").ID()
	expected := "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId
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
			Input: "/groups/groupId/calendar/calendarView/eventId/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments/attachmentId",
			Expected: &GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{
				GroupId:      "groupId",
				EventId:      "eventId",
				EventId1:     "eventId1",
				EventId2:     "eventId2",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(v.Input)
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

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId
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
			Input: "/groups/groupId/calendar/calendarView/eventId/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments/attachmentId",
			Expected: &GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{
				GroupId:      "groupId",
				EventId:      "eventId",
				EventId1:     "eventId1",
				EventId2:     "eventId2",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendar/calendarView/eventId/instances/eventId1/exceptionOccurrences/eventId2/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{
				GroupId:      "gRoUpId",
				EventId:      "eVeNtId",
				EventId1:     "eVeNtId1",
				EventId2:     "eVeNtId2",
				AttachmentId: "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/cAlEnDaRvIeW/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively(v.Input)
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

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId(t *testing.T) {
	segments := GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
