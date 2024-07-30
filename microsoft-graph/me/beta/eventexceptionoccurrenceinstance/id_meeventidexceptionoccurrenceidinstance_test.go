package eventexceptionoccurrenceinstance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdExceptionOccurrenceIdInstanceId{}

func TestNewMeEventIdExceptionOccurrenceIdInstanceID(t *testing.T) {
	id := NewMeEventIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}

	if id.EventId2 != "eventId2Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId2'", id.EventId2, "eventId2Value")
	}
}

func TestFormatMeEventIdExceptionOccurrenceIdInstanceID(t *testing.T) {
	actual := NewMeEventIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value").ID()
	expected := "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeEventIdExceptionOccurrenceIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEventIdExceptionOccurrenceIdInstanceId
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
			Input: "/me/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value",
			Expected: &MeEventIdExceptionOccurrenceIdInstanceId{
				EventId:  "eventIdValue",
				EventId1: "eventId1Value",
				EventId2: "eventId2Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEventIdExceptionOccurrenceIdInstanceID(v.Input)
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

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

	}
}

func TestParseMeEventIdExceptionOccurrenceIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEventIdExceptionOccurrenceIdInstanceId
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
			Input: "/me/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value",
			Expected: &MeEventIdExceptionOccurrenceIdInstanceId{
				EventId:  "eventIdValue",
				EventId1: "eventId1Value",
				EventId2: "eventId2Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs/eVeNtId2vAlUe",
			Expected: &MeEventIdExceptionOccurrenceIdInstanceId{
				EventId:  "eVeNtIdVaLuE",
				EventId1: "eVeNtId1vAlUe",
				EventId2: "eVeNtId2vAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs/eVeNtId2vAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEventIdExceptionOccurrenceIdInstanceIDInsensitively(v.Input)
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

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

	}
}

func TestSegmentsForMeEventIdExceptionOccurrenceIdInstanceId(t *testing.T) {
	segments := MeEventIdExceptionOccurrenceIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeEventIdExceptionOccurrenceIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
