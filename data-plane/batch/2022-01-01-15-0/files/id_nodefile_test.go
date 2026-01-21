package files

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NodeFileId{}

func TestNewNodeFileID(t *testing.T) {
	id := NewNodeFileID("https://endpoint_url", "poolId", "nodeId")

	if id.BaseURI != "https://endpoint_url" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint_url")
	}

	if id.PoolId != "poolId" {
		t.Fatalf("Expected %q but got %q for Segment 'PoolId'", id.PoolId, "poolId")
	}

	if id.NodeId != "nodeId" {
		t.Fatalf("Expected %q but got %q for Segment 'NodeId'", id.NodeId, "nodeId")
	}
}

func TestFormatNodeFileID(t *testing.T) {
	actual := NewNodeFileID("https://endpoint_url", "poolId", "nodeId").ID()
	expected := "/https://endpoint_url/pools/poolId/nodes/nodeId/files"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseNodeFileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NodeFileId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools/poolId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools/poolId/nodes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId",
			Error: true,
		},
		{
			// Valid URI
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/files",
			Expected: &NodeFileId{
				BaseURI: "https://endpoint_url",
				PoolId:  "poolId",
				NodeId:  "nodeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/files/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNodeFileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.PoolId != v.Expected.PoolId {
			t.Fatalf("Expected %q but got %q for PoolId", v.Expected.PoolId, actual.PoolId)
		}

		if actual.NodeId != v.Expected.NodeId {
			t.Fatalf("Expected %q but got %q for NodeId", v.Expected.NodeId, actual.NodeId)
		}

	}
}

func TestParseNodeFileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *NodeFileId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools/poolId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools/poolId/nodes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD/nOdEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD/nOdEs/nOdEiD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/files",
			Expected: &NodeFileId{
				BaseURI: "https://endpoint_url",
				PoolId:  "poolId",
				NodeId:  "nodeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/files/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD/nOdEs/nOdEiD/fIlEs",
			Expected: &NodeFileId{
				BaseURI: "hTtPs://eNdPoInT_UrL",
				PoolId:  "pOoLiD",
				NodeId:  "nOdEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD/nOdEs/nOdEiD/fIlEs/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseNodeFileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.PoolId != v.Expected.PoolId {
			t.Fatalf("Expected %q but got %q for PoolId", v.Expected.PoolId, actual.PoolId)
		}

		if actual.NodeId != v.Expected.NodeId {
			t.Fatalf("Expected %q but got %q for NodeId", v.Expected.NodeId, actual.NodeId)
		}

	}
}

func TestSegmentsForNodeFileId(t *testing.T) {
	segments := NodeFileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("NodeFileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
