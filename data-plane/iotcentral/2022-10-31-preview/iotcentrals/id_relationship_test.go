package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RelationshipId{}

func TestNewRelationshipID(t *testing.T) {
	id := NewRelationshipID("https://endpoint-url.example.com", "deviceId", "relationshipId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DeviceId != "deviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceId'", id.DeviceId, "deviceId")
	}

	if id.RelationshipId != "relationshipId" {
		t.Fatalf("Expected %q but got %q for Segment 'RelationshipId'", id.RelationshipId, "relationshipId")
	}
}

func TestFormatRelationshipID(t *testing.T) {
	actual := NewRelationshipID("https://endpoint-url.example.com", "deviceId", "relationshipId").ID()
	expected := "https://endpoint-url.example.com/devices/deviceId/relationships/relationshipId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRelationshipID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RelationshipId
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
			Input: "https://endpoint-url.example.com/devices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/relationships",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/devices/deviceId/relationships/relationshipId",
			Expected: &RelationshipId{
				BaseURI:        "https://endpoint-url.example.com",
				DeviceId:       "deviceId",
				RelationshipId: "relationshipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/devices/deviceId/relationships/relationshipId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRelationshipID(v.Input)
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

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.RelationshipId != v.Expected.RelationshipId {
			t.Fatalf("Expected %q but got %q for RelationshipId", v.Expected.RelationshipId, actual.RelationshipId)
		}

	}
}

func TestParseRelationshipIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RelationshipId
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
			Input: "https://endpoint-url.example.com/devices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/devices/deviceId/relationships",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/rElAtIoNsHiPs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/devices/deviceId/relationships/relationshipId",
			Expected: &RelationshipId{
				BaseURI:        "https://endpoint-url.example.com",
				DeviceId:       "deviceId",
				RelationshipId: "relationshipId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/devices/deviceId/relationships/relationshipId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/rElAtIoNsHiPs/rElAtIoNsHiPiD",
			Expected: &RelationshipId{
				BaseURI:        "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DeviceId:       "dEvIcEiD",
				RelationshipId: "rElAtIoNsHiPiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dEvIcEs/dEvIcEiD/rElAtIoNsHiPs/rElAtIoNsHiPiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRelationshipIDInsensitively(v.Input)
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

		if actual.DeviceId != v.Expected.DeviceId {
			t.Fatalf("Expected %q but got %q for DeviceId", v.Expected.DeviceId, actual.DeviceId)
		}

		if actual.RelationshipId != v.Expected.RelationshipId {
			t.Fatalf("Expected %q but got %q for RelationshipId", v.Expected.RelationshipId, actual.RelationshipId)
		}

	}
}

func TestSegmentsForRelationshipId(t *testing.T) {
	segments := RelationshipId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RelationshipId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
