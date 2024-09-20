package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId{}

func TestNewUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionID(t *testing.T) {
	id := NewUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionID("userId", "eventId", "eventId1", "eventId2", "extensionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
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

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionID(t *testing.T) {
	actual := NewUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionID("userId", "eventId", "eventId1", "eventId2", "extensionId").ID()
	expected := "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions/extensionId",
			Expected: &UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId{
				UserId:      "userId",
				EventId:     "eventId",
				EventId1:    "eventId1",
				EventId2:    "eventId2",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions/extensionId",
			Expected: &UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId{
				UserId:      "userId",
				EventId:     "eventId",
				EventId1:    "eventId1",
				EventId2:    "eventId2",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/events/eventId/exceptionOccurrences/eventId1/instances/eventId2/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId{
				UserId:      "uSeRiD",
				EventId:     "eVeNtId",
				EventId1:    "eVeNtId1",
				EventId2:    "eVeNtId2",
				ExtensionId: "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/eVeNtS/eVeNtId/eXcEpTiOnOcCuRrEnCeS/eVeNtId1/iNsTaNcEs/eVeNtId2/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId(t *testing.T) {
	segments := UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdEventIdExceptionOccurrenceIdInstanceIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
