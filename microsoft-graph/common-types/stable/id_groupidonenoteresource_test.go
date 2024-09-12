package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteResourceId{}

func TestNewGroupIdOnenoteResourceID(t *testing.T) {
	id := NewGroupIdOnenoteResourceID("groupIdValue", "onenoteResourceIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.OnenoteResourceId != "onenoteResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnenoteResourceId'", id.OnenoteResourceId, "onenoteResourceIdValue")
	}
}

func TestFormatGroupIdOnenoteResourceID(t *testing.T) {
	actual := NewGroupIdOnenoteResourceID("groupIdValue", "onenoteResourceIdValue").ID()
	expected := "/groups/groupIdValue/onenote/resources/onenoteResourceIdValue"
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
			Input: "/groups/groupIdValue/onenote/resources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/resources/onenoteResourceIdValue",
			Expected: &GroupIdOnenoteResourceId{
				GroupId:           "groupIdValue",
				OnenoteResourceId: "onenoteResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/resources/onenoteResourceIdValue/extra",
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
			Input: "/groups/groupIdValue/onenote/resources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/rEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/resources/onenoteResourceIdValue",
			Expected: &GroupIdOnenoteResourceId{
				GroupId:           "groupIdValue",
				OnenoteResourceId: "onenoteResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/resources/onenoteResourceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeIdVaLuE",
			Expected: &GroupIdOnenoteResourceId{
				GroupId:           "gRoUpIdVaLuE",
				OnenoteResourceId: "oNeNoTeReSoUrCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/rEsOuRcEs/oNeNoTeReSoUrCeIdVaLuE/extra",
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
