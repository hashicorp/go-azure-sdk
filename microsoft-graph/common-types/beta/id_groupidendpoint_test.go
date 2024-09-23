package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdEndpointId{}

func TestNewGroupIdEndpointID(t *testing.T) {
	id := NewGroupIdEndpointID("groupId", "endpointId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.EndpointId != "endpointId" {
		t.Fatalf("Expected %q but got %q for Segment 'EndpointId'", id.EndpointId, "endpointId")
	}
}

func TestFormatGroupIdEndpointID(t *testing.T) {
	actual := NewGroupIdEndpointID("groupId", "endpointId").ID()
	expected := "/groups/groupId/endpoints/endpointId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdEndpointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdEndpointId
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
			Input: "/groups/groupId/endpoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/endpoints/endpointId",
			Expected: &GroupIdEndpointId{
				GroupId:    "groupId",
				EndpointId: "endpointId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/endpoints/endpointId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdEndpointID(v.Input)
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

		if actual.EndpointId != v.Expected.EndpointId {
			t.Fatalf("Expected %q but got %q for EndpointId", v.Expected.EndpointId, actual.EndpointId)
		}

	}
}

func TestParseGroupIdEndpointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdEndpointId
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
			Input: "/groups/groupId/endpoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eNdPoInTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/endpoints/endpointId",
			Expected: &GroupIdEndpointId{
				GroupId:    "groupId",
				EndpointId: "endpointId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/endpoints/endpointId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eNdPoInTs/eNdPoInTiD",
			Expected: &GroupIdEndpointId{
				GroupId:    "gRoUpId",
				EndpointId: "eNdPoInTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/eNdPoInTs/eNdPoInTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdEndpointIDInsensitively(v.Input)
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

		if actual.EndpointId != v.Expected.EndpointId {
			t.Fatalf("Expected %q but got %q for EndpointId", v.Expected.EndpointId, actual.EndpointId)
		}

	}
}

func TestSegmentsForGroupIdEndpointId(t *testing.T) {
	segments := GroupIdEndpointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdEndpointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
