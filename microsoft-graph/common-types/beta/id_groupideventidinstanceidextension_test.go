package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEventIdInstanceIdExtensionId{}

func TestNewGroupIdEventIdInstanceIdExtensionID(t *testing.T) {
	id := NewGroupIdEventIdInstanceIdExtensionID("groupId", "eventId", "eventId1", "extensionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.EventId != "eventId" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventId")
	}

	if id.EventId1 != "eventId1" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId1'", id.EventId1, "eventId1")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatGroupIdEventIdInstanceIdExtensionID(t *testing.T) {
	actual := NewGroupIdEventIdInstanceIdExtensionID("groupId", "eventId", "eventId1", "extensionId").ID()
	expected := "/groups/groupId/events/eventId/instances/eventId1/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdEventIdInstanceIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdEventIdInstanceIdExtensionId
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
			Input: "/groups/groupId/events",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId/instances/eventId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId/instances/eventId1/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/events/eventId/instances/eventId1/extensions/extensionId",
			Expected: &GroupIdEventIdInstanceIdExtensionId{
				GroupId:     "groupId",
				EventId:     "eventId",
				EventId1:    "eventId1",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/events/eventId/instances/eventId1/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdEventIdInstanceIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseGroupIdEventIdInstanceIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdEventIdInstanceIdExtensionId
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
			Input: "/groups/groupId/events",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS/eVeNtId/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId/instances/eventId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/events/eventId/instances/eventId1/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/events/eventId/instances/eventId1/extensions/extensionId",
			Expected: &GroupIdEventIdInstanceIdExtensionId{
				GroupId:     "groupId",
				EventId:     "eventId",
				EventId1:    "eventId1",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/events/eventId/instances/eventId1/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &GroupIdEventIdInstanceIdExtensionId{
				GroupId:     "gRoUpId",
				EventId:     "eVeNtId",
				EventId1:    "eVeNtId1",
				ExtensionId: "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eVeNtS/eVeNtId/iNsTaNcEs/eVeNtId1/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdEventIdInstanceIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForGroupIdEventIdInstanceIdExtensionId(t *testing.T) {
	segments := GroupIdEventIdInstanceIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdEventIdInstanceIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
