package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId{}

func TestNewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID(t *testing.T) {
	id := NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

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
}

func TestFormatGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID(t *testing.T) {
	actual := NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2").ID()
	expected := "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances/eventId2"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId
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
			Input: "/groups/groupId/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances/eventId2",
			Expected: &GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId{
				GroupId:  "groupId",
				EventId:  "eventId",
				EventId1: "eventId1",
				EventId2: "eventId2",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances/eventId2/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID(v.Input)
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

	}
}

func TestParseGroupIdCalendarViewIdExceptionOccurrenceIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId
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
			Input: "/groups/groupId/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances/eventId2",
			Expected: &GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId{
				GroupId:  "groupId",
				EventId:  "eventId",
				EventId1: "eventId1",
				EventId2: "eventId2",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendarView/eventId/exceptionOccurrences/eventId1/instances/eventId2/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2",
			Expected: &GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId{
				GroupId:  "gRoUpId",
				EventId:  "eVeNtId",
				EventId1: "eVeNtId1",
				EventId2: "eVeNtId2",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarViewIdExceptionOccurrenceIdInstanceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupIdCalendarViewIdExceptionOccurrenceIdInstanceId(t *testing.T) {
	segments := GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdCalendarViewIdExceptionOccurrenceIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
