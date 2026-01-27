package computenodes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserId{}

func TestNewUserID(t *testing.T) {
	id := NewUserID("https://endpoint_url", "poolId", "nodeId")

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

func TestFormatUserID(t *testing.T) {
	actual := NewUserID("https://endpoint_url", "poolId", "nodeId").ID()
	expected := "/https://endpoint_url/pools/poolId/nodes/nodeId/users"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserId
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
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/users",
			Expected: &UserId{
				BaseURI: "https://endpoint_url",
				PoolId:  "poolId",
				NodeId:  "nodeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/users/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserID(v.Input)
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

func TestParseUserIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserId
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
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/users",
			Expected: &UserId{
				BaseURI: "https://endpoint_url",
				PoolId:  "poolId",
				NodeId:  "nodeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/pools/poolId/nodes/nodeId/users/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD/nOdEs/nOdEiD/uSeRs",
			Expected: &UserId{
				BaseURI: "hTtPs://eNdPoInT_UrL",
				PoolId:  "pOoLiD",
				NodeId:  "nOdEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/pOoLs/pOoLiD/nOdEs/nOdEiD/uSeRs/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIDInsensitively(v.Input)
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

func TestSegmentsForUserId(t *testing.T) {
	segments := UserId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
