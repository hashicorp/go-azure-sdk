package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleId{}

func TestNewRoleID(t *testing.T) {
	id := NewRoleID("https://endpoint-url.example.com", "roleId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.RoleId != "roleId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleId'", id.RoleId, "roleId")
	}
}

func TestFormatRoleID(t *testing.T) {
	actual := NewRoleID("https://endpoint-url.example.com", "roleId").ID()
	expected := "https://endpoint-url.example.com/roles/roleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleId
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
			Input: "https://endpoint-url.example.com/roles",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/roles/roleId",
			Expected: &RoleId{
				BaseURI: "https://endpoint-url.example.com",
				RoleId:  "roleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/roles/roleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleID(v.Input)
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

		if actual.RoleId != v.Expected.RoleId {
			t.Fatalf("Expected %q but got %q for RoleId", v.Expected.RoleId, actual.RoleId)
		}

	}
}

func TestParseRoleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleId
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
			Input: "https://endpoint-url.example.com/roles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/roles/roleId",
			Expected: &RoleId{
				BaseURI: "https://endpoint-url.example.com",
				RoleId:  "roleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/roles/roleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEs/rOlEiD",
			Expected: &RoleId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				RoleId:  "rOlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEs/rOlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleIDInsensitively(v.Input)
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

		if actual.RoleId != v.Expected.RoleId {
			t.Fatalf("Expected %q but got %q for RoleId", v.Expected.RoleId, actual.RoleId)
		}

	}
}

func TestSegmentsForRoleId(t *testing.T) {
	segments := RoleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
