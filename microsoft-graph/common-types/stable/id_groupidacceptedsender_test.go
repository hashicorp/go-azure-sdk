package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdAcceptedSenderId{}

func TestNewGroupIdAcceptedSenderID(t *testing.T) {
	id := NewGroupIdAcceptedSenderID("groupId", "directoryObjectId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatGroupIdAcceptedSenderID(t *testing.T) {
	actual := NewGroupIdAcceptedSenderID("groupId", "directoryObjectId").ID()
	expected := "/groups/groupId/acceptedSenders/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdAcceptedSenderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdAcceptedSenderId
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
			Input: "/groups/groupId/acceptedSenders",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/acceptedSenders/directoryObjectId",
			Expected: &GroupIdAcceptedSenderId{
				GroupId:           "groupId",
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/acceptedSenders/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdAcceptedSenderID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseGroupIdAcceptedSenderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdAcceptedSenderId
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
			Input: "/groups/groupId/acceptedSenders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/aCcEpTeDsEnDeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/acceptedSenders/directoryObjectId",
			Expected: &GroupIdAcceptedSenderId{
				GroupId:           "groupId",
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/acceptedSenders/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/aCcEpTeDsEnDeRs/dIrEcToRyObJeCtId",
			Expected: &GroupIdAcceptedSenderId{
				GroupId:           "gRoUpId",
				DirectoryObjectId: "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/aCcEpTeDsEnDeRs/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdAcceptedSenderIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForGroupIdAcceptedSenderId(t *testing.T) {
	segments := GroupIdAcceptedSenderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdAcceptedSenderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
