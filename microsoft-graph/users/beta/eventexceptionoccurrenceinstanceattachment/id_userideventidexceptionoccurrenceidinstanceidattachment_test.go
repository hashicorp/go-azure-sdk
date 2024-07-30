package eventexceptionoccurrenceinstanceattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId{}

func TestNewUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID(t *testing.T) {
	id := NewUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}

	if id.EventId2 != "eventId2Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId2'", id.EventId2, "eventId2Value")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID(t *testing.T) {
	actual := NewUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue").ID()
	expected := "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId
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
			Input: "/users/userIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments/attachmentIdValue",
			Expected: &UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId{
				UserId:       "userIdValue",
				EventId:      "eventIdValue",
				EventId1:     "eventId1Value",
				EventId2:     "eventId2Value",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID(v.Input)
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

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId
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
			Input: "/users/userIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs/eVeNtId2vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs/eVeNtId2vAlUe/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments/attachmentIdValue",
			Expected: &UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId{
				UserId:       "userIdValue",
				EventId:      "eventIdValue",
				EventId1:     "eventId1Value",
				EventId2:     "eventId2Value",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/events/eventIdValue/exceptionOccurrences/eventId1Value/instances/eventId2Value/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs/eVeNtId2vAlUe/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId{
				UserId:       "uSeRiDvAlUe",
				EventId:      "eVeNtIdVaLuE",
				EventId1:     "eVeNtId1vAlUe",
				EventId2:     "eVeNtId2vAlUe",
				AttachmentId: "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/eVeNtS/eVeNtIdVaLuE/eXcEpTiOnOcCuRrEnCeS/eVeNtId1vAlUe/iNsTaNcEs/eVeNtId2vAlUe/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively(v.Input)
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

		if actual.EventId2 != v.Expected.EventId2 {
			t.Fatalf("Expected %q but got %q for EventId2", v.Expected.EventId2, actual.EventId2)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForUserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId(t *testing.T) {
	segments := UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdEventIdExceptionOccurrenceIdInstanceIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
