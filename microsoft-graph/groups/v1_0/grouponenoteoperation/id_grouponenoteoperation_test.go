package grouponenoteoperation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteOperationId{}

func TestNewGroupOnenoteOperationID(t *testing.T) {
	id := NewGroupOnenoteOperationID("groupIdValue", "onenoteOperationIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.OnenoteOperationId != "onenoteOperationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteOperationId'", id.OnenoteOperationId, "onenoteOperationIdValue")
	}
}

func TestFormatGroupOnenoteOperationID(t *testing.T) {
	actual := NewGroupOnenoteOperationID("groupIdValue", "onenoteOperationIdValue").ID()
	expected := "/groups/groupIdValue/onenote/operations/onenoteOperationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupOnenoteOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupOnenoteOperationId
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
			Input: "/groups/groupIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/operations/onenoteOperationIdValue",
			Expected: &GroupOnenoteOperationId{
				GroupId:            "groupIdValue",
				OnenoteOperationId: "onenoteOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/operations/onenoteOperationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupOnenoteOperationID(v.Input)
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

func TestParseGroupOnenoteOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupOnenoteOperationId
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
			Input: "/groups/groupIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/operations/onenoteOperationIdValue",
			Expected: &GroupOnenoteOperationId{
				GroupId:            "groupIdValue",
				OnenoteOperationId: "onenoteOperationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/operations/onenoteOperationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiDvAlUe",
			Expected: &GroupOnenoteOperationId{
				GroupId:            "gRoUpIdVaLuE",
				OnenoteOperationId: "oNeNoTeOpErAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/oPeRaTiOnS/oNeNoTeOpErAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupOnenoteOperationIDInsensitively(v.Input)
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

func TestSegmentsForGroupOnenoteOperationId(t *testing.T) {
	segments := GroupOnenoteOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupOnenoteOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
