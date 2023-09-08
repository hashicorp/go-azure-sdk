package groupeventinstanceattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEventInstanceAttachmentId{}

func TestNewGroupEventInstanceAttachmentID(t *testing.T) {
	id := NewGroupEventInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.EventId != "eventIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventIdValue")
	}

	if id.EventId1 != "eventId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1Value")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatGroupEventInstanceAttachmentID(t *testing.T) {
	actual := NewGroupEventInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue").ID()
	expected := "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupEventInstanceAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupEventInstanceAttachmentId
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
			Input: "/groups/groupIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue",
			Expected: &GroupEventInstanceAttachmentId{
				GroupId:      "groupIdValue",
				EventId:      "eventIdValue",
				EventId1:     "eventId1Value",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupEventInstanceAttachmentID(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseGroupEventInstanceAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupEventInstanceAttachmentId
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
			Input: "/groups/groupIdValue/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS/eVeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue",
			Expected: &GroupEventInstanceAttachmentId{
				GroupId:      "groupIdValue",
				EventId:      "eventIdValue",
				EventId1:     "eventId1Value",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/events/eventIdValue/instances/eventId1Value/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &GroupEventInstanceAttachmentId{
				GroupId:      "gRoUpIdVaLuE",
				EventId:      "eVeNtIdVaLuE",
				EventId1:     "eVeNtId1vAlUe",
				AttachmentId: "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eVeNtS/eVeNtIdVaLuE/iNsTaNcEs/eVeNtId1vAlUe/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupEventInstanceAttachmentIDInsensitively(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForGroupEventInstanceAttachmentId(t *testing.T) {
	segments := GroupEventInstanceAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupEventInstanceAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
