package calendarviewinstanceexceptionoccurrenceextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}

func TestNewMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(t *testing.T) {
	id := NewMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID("eventIdValue", "eventId1Value", "eventId2Value", "extensionIdValue")

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}

	if id.EventId2 != "eventId2Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId2'", id.EventId2, "eventId2Value")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(t *testing.T) {
	actual := NewMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID("eventIdValue", "eventId1Value", "eventId2Value", "extensionIdValue").ID()
	expected := "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId
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
			Input: "/me/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions/extensionIdValue",
			Expected: &MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{
				EventId:     "eventIdValue",
				EventId1:    "eventId1Value",
				EventId2:    "eventId2Value",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId
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
			Input: "/me/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS/eVeNtId2vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS/eVeNtId2vAlUe/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions/extensionIdValue",
			Expected: &MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{
				EventId:     "eventIdValue",
				EventId1:    "eventId1Value",
				EventId2:    "eventId2Value",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/calendarView/eventIdValue/instances/eventId1Value/exceptionOccurrences/eventId2Value/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS/eVeNtId2vAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{
				EventId:     "eVeNtIdVaLuE",
				EventId1:    "eVeNtId1vAlUe",
				EventId2:    "eVeNtId2vAlUe",
				ExtensionId: "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cAlEnDaRvIeW/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/eXcEpTiOnOcCuRrEnCeS/eVeNtId2vAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForMeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId(t *testing.T) {
	segments := MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
