package groupcalendarviewexceptionoccurrence

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewExceptionOccurrenceId{}

func TestNewGroupCalendarViewExceptionOccurrenceID(t *testing.T) {
	id := NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}
}

func TestFormatGroupCalendarViewExceptionOccurrenceID(t *testing.T) {
	actual := NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value").ID()
	expected := "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupCalendarViewExceptionOccurrenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupCalendarViewExceptionOccurrenceId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value",
			Expected: &GroupCalendarViewExceptionOccurrenceId{
				GroupId:  "groupIdValue",
				EventId:  "eventIdValue",
				EventId1: "eventId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupCalendarViewExceptionOccurrenceID(v.Input)
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

func TestParseGroupCalendarViewExceptionOccurrenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupCalendarViewExceptionOccurrenceId
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
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value",
			Expected: &GroupCalendarViewExceptionOccurrenceId{
				GroupId:  "groupIdValue",
				EventId:  "eventIdValue",
				EventId1: "eventId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe",
			Expected: &GroupCalendarViewExceptionOccurrenceId{
				GroupId:  "gRoUpIdVaLuE",
				EventId:  "eVeNtIdVaLuE",
				EventId1: "eVeNtId1vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupCalendarViewExceptionOccurrenceIDInsensitively(v.Input)
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

func TestSegmentsForGroupCalendarViewExceptionOccurrenceId(t *testing.T) {
	segments := GroupCalendarViewExceptionOccurrenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupCalendarViewExceptionOccurrenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
