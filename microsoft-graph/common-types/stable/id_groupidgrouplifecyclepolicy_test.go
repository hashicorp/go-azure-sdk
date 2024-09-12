package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdGroupLifecyclePolicyId{}

func TestNewGroupIdGroupLifecyclePolicyID(t *testing.T) {
	id := NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.GroupLifecyclePolicyId != "groupLifecyclePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupLifecyclePolicyId'", id.GroupLifecyclePolicyId, "groupLifecyclePolicyIdValue")
	}
}

func TestFormatGroupIdGroupLifecyclePolicyID(t *testing.T) {
	actual := NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue").ID()
	expected := "/groups/groupIdValue/groupLifecyclePolicies/groupLifecyclePolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdGroupLifecyclePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdGroupLifecyclePolicyId
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
			Input: "/groups/groupIdValue/groupLifecyclePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/groupLifecyclePolicies/groupLifecyclePolicyIdValue",
			Expected: &GroupIdGroupLifecyclePolicyId{
				GroupId:                "groupIdValue",
				GroupLifecyclePolicyId: "groupLifecyclePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/groupLifecyclePolicies/groupLifecyclePolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdGroupLifecyclePolicyID(v.Input)
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

		if actual.GroupLifecyclePolicyId != v.Expected.GroupLifecyclePolicyId {
			t.Fatalf("Expected %q but got %q for GroupLifecyclePolicyId", v.Expected.GroupLifecyclePolicyId, actual.GroupLifecyclePolicyId)
		}

	}
}

func TestParseGroupIdGroupLifecyclePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdGroupLifecyclePolicyId
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
			Input: "/groups/groupIdValue/groupLifecyclePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/gRoUpLiFeCyClEpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/groupLifecyclePolicies/groupLifecyclePolicyIdValue",
			Expected: &GroupIdGroupLifecyclePolicyId{
				GroupId:                "groupIdValue",
				GroupLifecyclePolicyId: "groupLifecyclePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/groupLifecyclePolicies/groupLifecyclePolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/gRoUpLiFeCyClEpOlIcIeS/gRoUpLiFeCyClEpOlIcYiDvAlUe",
			Expected: &GroupIdGroupLifecyclePolicyId{
				GroupId:                "gRoUpIdVaLuE",
				GroupLifecyclePolicyId: "gRoUpLiFeCyClEpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/gRoUpLiFeCyClEpOlIcIeS/gRoUpLiFeCyClEpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdGroupLifecyclePolicyIDInsensitively(v.Input)
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

		if actual.GroupLifecyclePolicyId != v.Expected.GroupLifecyclePolicyId {
			t.Fatalf("Expected %q but got %q for GroupLifecyclePolicyId", v.Expected.GroupLifecyclePolicyId, actual.GroupLifecyclePolicyId)
		}

	}
}

func TestSegmentsForGroupIdGroupLifecyclePolicyId(t *testing.T) {
	segments := GroupIdGroupLifecyclePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdGroupLifecyclePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
