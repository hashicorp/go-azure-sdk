package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteResourceId{}

func TestNewGroupIdOnenoteResourceID(t *testing.T) {
	id := NewGroupIdOnenoteResourceID("groupId", "onenoteResourceId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.OnenoteResourceId != "onenoteResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteResourceId'", id.OnenoteResourceId, "onenoteResourceId")
	}
}

func TestFormatGroupIdOnenoteResourceID(t *testing.T) {
	actual := NewGroupIdOnenoteResourceID("groupId", "onenoteResourceId").ID()
	expected := "/groups/groupId/onenote/resources/onenoteResourceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdOnenoteResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdOnenoteResourceId
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
			Input: "/groups/groupId/onenote/resources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/onenote/resources/onenoteResourceId",
			Expected: &GroupIdOnenoteResourceId{
				GroupId:           "groupId",
				OnenoteResourceId: "onenoteResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/onenote/resources/onenoteResourceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdOnenoteResourceID(v.Input)
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

		if actual.OnenoteResourceId != v.Expected.OnenoteResourceId {
			t.Fatalf("Expected %q but got %q for OnenoteResourceId", v.Expected.OnenoteResourceId, actual.OnenoteResourceId)
		}

	}
}

func TestParseGroupIdOnenoteResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdOnenoteResourceId
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
			Input: "/groups/groupId/onenote/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/rEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/onenote/resources/onenoteResourceId",
			Expected: &GroupIdOnenoteResourceId{
				GroupId:           "groupId",
				OnenoteResourceId: "onenoteResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/onenote/resources/onenoteResourceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeId",
			Expected: &GroupIdOnenoteResourceId{
				GroupId:           "gRoUpId",
				OnenoteResourceId: "oNeNoTeReSoUrCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdOnenoteResourceIDInsensitively(v.Input)
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

		if actual.OnenoteResourceId != v.Expected.OnenoteResourceId {
			t.Fatalf("Expected %q but got %q for OnenoteResourceId", v.Expected.OnenoteResourceId, actual.OnenoteResourceId)
		}

	}
}

func TestSegmentsForGroupIdOnenoteResourceId(t *testing.T) {
	segments := GroupIdOnenoteResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdOnenoteResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
