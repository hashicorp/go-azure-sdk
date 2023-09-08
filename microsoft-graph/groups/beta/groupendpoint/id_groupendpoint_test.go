package groupendpoint

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEndpointId{}

func TestNewGroupEndpointID(t *testing.T) {
	id := NewGroupEndpointID("groupIdValue", "endpointIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.EndpointId != "endpointIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EndpointId'", id.EndpointId, "endpointIdValue")
	}
}

func TestFormatGroupEndpointID(t *testing.T) {
	actual := NewGroupEndpointID("groupIdValue", "endpointIdValue").ID()
	expected := "/groups/groupIdValue/endpoints/endpointIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupEndpointID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupEndpointId
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
			Input: "/groups/groupIdValue/endpoints",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/endpoints/endpointIdValue",
			Expected: &GroupEndpointId{
				GroupId:    "groupIdValue",
				EndpointId: "endpointIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/endpoints/endpointIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupEndpointID(v.Input)
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

func TestParseGroupEndpointIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupEndpointId
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
			Input: "/groups/groupIdValue/endpoints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eNdPoInTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/endpoints/endpointIdValue",
			Expected: &GroupEndpointId{
				GroupId:    "groupIdValue",
				EndpointId: "endpointIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/endpoints/endpointIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eNdPoInTs/eNdPoInTiDvAlUe",
			Expected: &GroupEndpointId{
				GroupId:    "gRoUpIdVaLuE",
				EndpointId: "eNdPoInTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/eNdPoInTs/eNdPoInTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupEndpointIDInsensitively(v.Input)
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

func TestSegmentsForGroupEndpointId(t *testing.T) {
	segments := GroupEndpointId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupEndpointId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
