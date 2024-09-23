package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarViewIdExtensionId{}

func TestNewGroupIdCalendarViewIdExtensionID(t *testing.T) {
	id := NewGroupIdCalendarViewIdExtensionID("groupId", "eventId", "extensionId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.EventId != "eventId" {
		t.Fatalf("Expected %q but got %q for Segment 'EventId'", id.EventId, "eventId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatGroupIdCalendarViewIdExtensionID(t *testing.T) {
	actual := NewGroupIdCalendarViewIdExtensionID("groupId", "eventId", "extensionId").ID()
	expected := "/groups/groupId/calendarView/eventId/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdCalendarViewIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarViewIdExtensionId
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
			Input: "/groups/groupId/calendarView",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendarView/eventId/extensions/extensionId",
			Expected: &GroupIdCalendarViewIdExtensionId{
				GroupId:     "groupId",
				EventId:     "eventId",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendarView/eventId/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarViewIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseGroupIdCalendarViewIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdCalendarViewIdExtensionId
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
			Input: "/groups/groupId/calendarView",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/calendarView/eventId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/calendarView/eventId/extensions/extensionId",
			Expected: &GroupIdCalendarViewIdExtensionId{
				GroupId:     "groupId",
				EventId:     "eventId",
				ExtensionId: "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/calendarView/eventId/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &GroupIdCalendarViewIdExtensionId{
				GroupId:     "gRoUpId",
				EventId:     "eVeNtId",
				ExtensionId: "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/cAlEnDaRvIeW/eVeNtId/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdCalendarViewIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForGroupIdCalendarViewIdExtensionId(t *testing.T) {
	segments := GroupIdCalendarViewIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdCalendarViewIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
