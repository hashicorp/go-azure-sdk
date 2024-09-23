package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId{}

func TestNewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID(t *testing.T) {
	id := NewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID("eventId", "eventId1", "eventId2", "attachmentId")

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

func TestFormatMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID(t *testing.T) {
	actual := NewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID("eventId", "eventId1", "eventId2", "attachmentId").ID()
	expected := "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId
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
			Input: "/me/events/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments/attachmentId",
			Expected: &MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId{
				EventId:      "eventId",
				EventId1:     "eventId1",
				EventId2:     "eventId2",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseMeEventIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId
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
			Input: "/me/events/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments/attachmentId",
			Expected: &MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId{
				EventId:      "eventId",
				EventId1:     "eventId1",
				EventId2:     "eventId2",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId{
				EventId:      "eVeNtId",
				EventId1:     "eVeNtId1",
				EventId2:     "eVeNtId2",
				AttachmentId: "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeEventIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForMeEventIdExceptionOccurrenceIdInstanceIdAttachmentId(t *testing.T) {
	segments := MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeEventIdExceptionOccurrenceIdInstanceIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
