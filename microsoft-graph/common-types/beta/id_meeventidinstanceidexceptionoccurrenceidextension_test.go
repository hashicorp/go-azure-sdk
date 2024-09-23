package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdInstanceIdExceptionOccurrenceIdExtensionId{}

func TestNewMeEventIdInstanceIdExceptionOccurrenceIdExtensionID(t *testing.T) {
	id := NewMeEventIdInstanceIdExceptionOccurrenceIdExtensionID("eventId", "eventId1", "eventId2", "extensionId")

	if id.EventId != "eventId" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventId")
	}

	if id.EventId1 != "eventId1" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1")
	}

	if id.EventId2 != "eventId2" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId2'", id.EventId2, "eventId2")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatMeEventIdInstanceIdExceptionOccurrenceIdExtensionID(t *testing.T) {
	actual := NewMeEventIdInstanceIdExceptionOccurrenceIdExtensionID("eventId", "eventId1", "eventId2", "extensionId").ID()
	expected := "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeEventIdInstanceIdExceptionOccurrenceIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEventIdInstanceIdExceptionOccurrenceIdExtensionId
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
			Input: "/me/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions/extensionId",
			Expected: &MeEventIdInstanceIdExceptionOccurrenceIdExtensionId{
				EventId:     "eventId",
				EventId1:    "eventId1",
				EventId2:    "eventId2",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEventIdInstanceIdExceptionOccurrenceIdExtensionID(v.Input)
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

func TestParseMeEventIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEventIdInstanceIdExceptionOccurrenceIdExtensionId
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
			Input: "/me/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions/extensionId",
			Expected: &MeEventIdInstanceIdExceptionOccurrenceIdExtensionId{
				EventId:     "eventId",
				EventId1:    "eventId1",
				EventId2:    "eventId2",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/events/eventId/instances/eventId1/exceptionOccurrences/eventId2/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &MeEventIdInstanceIdExceptionOccurrenceIdExtensionId{
				EventId:     "eVeNtId",
				EventId1:    "eVeNtId1",
				EventId2:    "eVeNtId2",
				ExtensionId: "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXcEpTiOnOcCuRrEnCeS/eVeNtId2/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEventIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(v.Input)
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

func TestSegmentsForMeEventIdInstanceIdExceptionOccurrenceIdExtensionId(t *testing.T) {
	segments := MeEventIdInstanceIdExceptionOccurrenceIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeEventIdInstanceIdExceptionOccurrenceIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
