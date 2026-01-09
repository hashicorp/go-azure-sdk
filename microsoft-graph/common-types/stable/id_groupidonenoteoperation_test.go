package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteOperationId{}

func TestNewGroupIdOnenoteOperationID(t *testing.T) {
	id := NewGroupIdOnenoteOperationID("groupId", "onenoteOperationId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.OnenoteOperationId != "onenoteOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteOperationId'", id.OnenoteOperationId, "onenoteOperationId")
	}
}

func TestFormatGroupIdOnenoteOperationID(t *testing.T) {
	actual := NewGroupIdOnenoteOperationID("groupId", "onenoteOperationId").ID()
	expected := "/groups/groupId/onenote/operations/onenoteOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdOnenoteOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdOnenoteOperationId
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
			Input: "/groups/groupId/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/onenote/operations/onenoteOperationId",
			Expected: &GroupIdOnenoteOperationId{
				GroupId:            "groupId",
				OnenoteOperationId: "onenoteOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/onenote/operations/onenoteOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdOnenoteOperationID(v.Input)
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

		if actual.OnenoteOperationId != v.Expected.OnenoteOperationId {
			t.Fatalf("Expected %q but got %q for OnenoteOperationId", v.Expected.OnenoteOperationId, actual.OnenoteOperationId)
		}

	}
}

func TestParseGroupIdOnenoteOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdOnenoteOperationId
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
			Input: "/groups/groupId/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/onenote/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/onenote/operations/onenoteOperationId",
			Expected: &GroupIdOnenoteOperationId{
				GroupId:            "groupId",
				OnenoteOperationId: "onenoteOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/onenote/operations/onenoteOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiD",
			Expected: &GroupIdOnenoteOperationId{
				GroupId:            "gRoUpId",
				OnenoteOperationId: "oNeNoTeOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdOnenoteOperationIDInsensitively(v.Input)
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

		if actual.OnenoteOperationId != v.Expected.OnenoteOperationId {
			t.Fatalf("Expected %q but got %q for OnenoteOperationId", v.Expected.OnenoteOperationId, actual.OnenoteOperationId)
		}

	}
}

func TestSegmentsForGroupIdOnenoteOperationId(t *testing.T) {
	segments := GroupIdOnenoteOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdOnenoteOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
