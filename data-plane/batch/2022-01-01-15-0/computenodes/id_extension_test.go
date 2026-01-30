package computenodes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExtensionId{}

func TestNewExtensionID(t *testing.T) {
	id := NewExtensionID("https://endpoint-url.example.com", "poolId", "nodeId", "extensionName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.PoolId != "poolId" {
		t.Fatalf("Expected %q but got %q for Segment 'PoolId'", id.PoolId, "poolId")
	}

	if id.NodeId != "nodeId" {
		t.Fatalf("Expected %q but got %q for Segment 'NodeId'", id.NodeId, "nodeId")
	}

	if id.ExtensionName != "extensionName" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionName'", id.ExtensionName, "extensionName")
	}
}

func TestFormatExtensionID(t *testing.T) {
	actual := NewExtensionID("https://endpoint-url.example.com", "poolId", "nodeId", "extensionName").ID()
	expected := "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions/extensionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions/extensionName",
			Expected: &ExtensionId{
				BaseURI:       "https://endpoint-url.example.com",
				PoolId:        "poolId",
				NodeId:        "nodeId",
				ExtensionName: "extensionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions/extensionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExtensionID(v.Input)
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

		if actual.ExtensionName != v.Expected.ExtensionName {
			t.Fatalf("Expected %q but got %q for ExtensionName", v.Expected.ExtensionName, actual.ExtensionName)
		}

	}
}

func TestParseExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExtensionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD/nOdEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD/nOdEs/nOdEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD/nOdEs/nOdEiD/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions/extensionName",
			Expected: &ExtensionId{
				BaseURI:       "https://endpoint-url.example.com",
				PoolId:        "poolId",
				NodeId:        "nodeId",
				ExtensionName: "extensionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pools/poolId/nodes/nodeId/extensions/extensionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD/nOdEs/nOdEiD/eXtEnSiOnS/eXtEnSiOnNaMe",
			Expected: &ExtensionId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				PoolId:        "pOoLiD",
				NodeId:        "nOdEiD",
				ExtensionName: "eXtEnSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD/nOdEs/nOdEiD/eXtEnSiOnS/eXtEnSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionName != v.Expected.ExtensionName {
			t.Fatalf("Expected %q but got %q for ExtensionName", v.Expected.ExtensionName, actual.ExtensionName)
		}

	}
}

func TestSegmentsForExtensionId(t *testing.T) {
	segments := ExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
