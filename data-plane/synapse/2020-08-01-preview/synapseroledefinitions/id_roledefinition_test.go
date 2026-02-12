package synapseroledefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleDefinitionId{}

func TestNewRoleDefinitionID(t *testing.T) {
	id := NewRoleDefinitionID("https://endpoint-url.example.com", "roleDefinitionId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.RoleDefinitionId != "roleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleDefinitionId'", id.RoleDefinitionId, "roleDefinitionId")
	}
}

func TestFormatRoleDefinitionID(t *testing.T) {
	actual := NewRoleDefinitionID("https://endpoint-url.example.com", "roleDefinitionId").ID()
	expected := "https://endpoint-url.example.com/roleDefinitions/roleDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleDefinitionId
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
			Input: "https://endpoint-url.example.com/roleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/roleDefinitions/roleDefinitionId",
			Expected: &RoleDefinitionId{
				BaseURI:          "https://endpoint-url.example.com",
				RoleDefinitionId: "roleDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/roleDefinitions/roleDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleDefinitionID(v.Input)
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

		if actual.RoleDefinitionId != v.Expected.RoleDefinitionId {
			t.Fatalf("Expected %q but got %q for RoleDefinitionId", v.Expected.RoleDefinitionId, actual.RoleDefinitionId)
		}

	}
}

func TestParseRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleDefinitionId
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
			Input: "https://endpoint-url.example.com/roleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/roleDefinitions/roleDefinitionId",
			Expected: &RoleDefinitionId{
				BaseURI:          "https://endpoint-url.example.com",
				RoleDefinitionId: "roleDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/roleDefinitions/roleDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEdEfInItIoNs/rOlEdEfInItIoNiD",
			Expected: &RoleDefinitionId{
				BaseURI:          "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				RoleDefinitionId: "rOlEdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEdEfInItIoNs/rOlEdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleDefinitionIDInsensitively(v.Input)
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

		if actual.RoleDefinitionId != v.Expected.RoleDefinitionId {
			t.Fatalf("Expected %q but got %q for RoleDefinitionId", v.Expected.RoleDefinitionId, actual.RoleDefinitionId)
		}

	}
}

func TestSegmentsForRoleDefinitionId(t *testing.T) {
	segments := RoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
