package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarEventIdExceptionOccurrenceId{}

func TestNewGroupIdCalendarEventIdExceptionOccurrenceID(t *testing.T) {
	id := NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.EventId != "eventId" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventId")
	}

	if id.EventId1 != "eventId1" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1")
	}
}

func TestFormatGroupIdCalendarEventIdExceptionOccurrenceID(t *testing.T) {
	actual := NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1").ID()
	expected := "/groups/groupId/calendar/events/eventId/exceptionOccurrences/eventId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdCalendarEventIdExceptionOccurrenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarEventIdExceptionOccurrenceId
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
			Input: "/groups/groupId/calendar/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/events/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendar/events/eventId/exceptionOccurrences/eventId1",
			Expected: &GroupIdCalendarEventIdExceptionOccurrenceId{
				GroupId:  "groupId",
				EventId:  "eventId",
				EventId1: "eventId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendar/events/eventId/exceptionOccurrences/eventId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarEventIdExceptionOccurrenceID(v.Input)
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

	}
}

func TestParseGroupIdCalendarEventIdExceptionOccurrenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarEventIdExceptionOccurrenceId
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
			Input: "/groups/groupId/calendar/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/eVeNtS/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendar/events/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendar/events/eventId/exceptionOccurrences/eventId1",
			Expected: &GroupIdCalendarEventIdExceptionOccurrenceId{
				GroupId:  "groupId",
				EventId:  "eventId",
				EventId1: "eventId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendar/events/eventId/exceptionOccurrences/eventId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1",
			Expected: &GroupIdCalendarEventIdExceptionOccurrenceId{
				GroupId:  "gRoUpId",
				EventId:  "eVeNtId",
				EventId1: "eVeNtId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaR/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarEventIdExceptionOccurrenceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupIdCalendarEventIdExceptionOccurrenceId(t *testing.T) {
	segments := GroupIdCalendarEventIdExceptionOccurrenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdCalendarEventIdExceptionOccurrenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
