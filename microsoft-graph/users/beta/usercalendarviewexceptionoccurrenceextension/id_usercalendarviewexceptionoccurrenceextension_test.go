package usercalendarviewexceptionoccurrenceextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewExceptionOccurrenceExtensionId{}

func TestNewUserCalendarViewExceptionOccurrenceExtensionID(t *testing.T) {
	id := NewUserCalendarViewExceptionOccurrenceExtensionID("userIdValue", "eventIdValue", "eventId1Value", "extensionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatUserCalendarViewExceptionOccurrenceExtensionID(t *testing.T) {
	actual := NewUserCalendarViewExceptionOccurrenceExtensionID("userIdValue", "eventIdValue", "eventId1Value", "extensionIdValue").ID()
	expected := "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserCalendarViewExceptionOccurrenceExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarViewExceptionOccurrenceExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions/extensionIdValue",
			Expected: &UserCalendarViewExceptionOccurrenceExtensionId{
				UserId:      "userIdValue",
				EventId:     "eventIdValue",
				EventId1:    "eventId1Value",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarViewExceptionOccurrenceExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserCalendarViewExceptionOccurrenceExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserCalendarViewExceptionOccurrenceExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions/extensionIdValue",
			Expected: &UserCalendarViewExceptionOccurrenceExtensionId{
				UserId:      "userIdValue",
				EventId:     "eventIdValue",
				EventId1:    "eventId1Value",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/calendarView/eventIdValue/exceptionOccurrences/eventId1Value/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &UserCalendarViewExceptionOccurrenceExtensionId{
				UserId:      "uSeRiDvAlUe",
				EventId:     "eVeNtIdVaLuE",
				EventId1:    "eVeNtId1vAlUe",
				ExtensionId: "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cAlEnDaRvIeW/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserCalendarViewExceptionOccurrenceExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.EventId != v.Expected.EventId {
			t.Fatalf("Expected %q but got %q for EventId", v.Expected.EventId, actual.EventId)
		}

		if actual.EventId1 != v.Expected.EventId1 {
			t.Fatalf("Expected %q but got %q for EventId1", v.Expected.EventId1, actual.EventId1)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserCalendarViewExceptionOccurrenceExtensionId(t *testing.T) {
	segments := UserCalendarViewExceptionOccurrenceExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserCalendarViewExceptionOccurrenceExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
