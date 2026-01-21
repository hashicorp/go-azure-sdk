package computenodes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NodeId{}

func TestNewNodeID(t *testing.T) {
	id := NewNodeID("poolId", "nodeId")

	if id.PoolId != "poolId" {
		t.Fatalf("Expected %q but got %q for Segment 'PoolId'", id.PoolId, "poolId")
	}

	if id.NodeId != "nodeId" {
		t.Fatalf("Expected %q but got %q for Segment 'NodeId'", id.NodeId, "nodeId")
	}
}

func TestFormatNodeID(t *testing.T) {
	actual := NewNodeID("poolId", "nodeId").ID()
	expected := "/pools/poolId/nodes/nodeId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNodeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NodeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/pools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/pools/poolId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/pools/poolId/nodes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/pools/poolId/nodes/nodeId",
			Expected: &NodeId{
				PoolId: "poolId",
				NodeId: "nodeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/pools/poolId/nodes/nodeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNodeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PoolId != v.Expected.PoolId {
			t.Fatalf("Expected %q but got %q for PoolId", v.Expected.PoolId, actual.PoolId)
		}

		if actual.NodeId != v.Expected.NodeId {
			t.Fatalf("Expected %q but got %q for NodeId", v.Expected.NodeId, actual.NodeId)
		}

	}
}

func TestParseNodeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NodeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/pools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOoLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/pools/poolId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOoLs/pOoLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/pools/poolId/nodes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOoLs/pOoLiD/nOdEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/pools/poolId/nodes/nodeId",
			Expected: &NodeId{
				PoolId: "poolId",
				NodeId: "nodeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/pools/poolId/nodes/nodeId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOoLs/pOoLiD/nOdEs/nOdEiD",
			Expected: &NodeId{
				PoolId: "pOoLiD",
				NodeId: "nOdEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOoLs/pOoLiD/nOdEs/nOdEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNodeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PoolId != v.Expected.PoolId {
			t.Fatalf("Expected %q but got %q for PoolId", v.Expected.PoolId, actual.PoolId)
		}

		if actual.NodeId != v.Expected.NodeId {
			t.Fatalf("Expected %q but got %q for NodeId", v.Expected.NodeId, actual.NodeId)
		}

	}
}

func TestSegmentsForNodeId(t *testing.T) {
	segments := NodeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NodeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
