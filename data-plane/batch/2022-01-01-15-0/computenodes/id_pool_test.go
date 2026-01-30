package computenodes

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PoolId{}

func TestNewPoolID(t *testing.T) {
	id := NewPoolID("https://endpoint-url.example.com", "poolId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.PoolId != "poolId" {
		t.Fatalf("Expected %q but got %q for Segment 'PoolId'", id.PoolId, "poolId")
	}
}

func TestFormatPoolID(t *testing.T) {
	actual := NewPoolID("https://endpoint-url.example.com", "poolId").ID()
	expected := "https://endpoint-url.example.com/pools/poolId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePoolID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PoolId
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
			// Valid URI
			Input: "https://endpoint-url.example.com/pools/poolId",
			Expected: &PoolId{
				BaseURI: "https://endpoint-url.example.com",
				PoolId:  "poolId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pools/poolId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePoolID(v.Input)
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

	}
}

func TestParsePoolIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PoolId
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
			// Valid URI
			Input: "https://endpoint-url.example.com/pools/poolId",
			Expected: &PoolId{
				BaseURI: "https://endpoint-url.example.com",
				PoolId:  "poolId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pools/poolId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD",
			Expected: &PoolId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				PoolId:  "pOoLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pOoLs/pOoLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePoolIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPoolId(t *testing.T) {
	segments := PoolId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PoolId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
