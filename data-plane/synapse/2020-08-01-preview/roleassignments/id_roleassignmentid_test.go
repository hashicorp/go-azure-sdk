package roleassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleAssignmentIdId{}

func TestNewRoleAssignmentIdID(t *testing.T) {
	id := NewRoleAssignmentIdID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.RoleAssignmentId != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'RoleAssignmentId'", id.RoleAssignmentId, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}
}

func TestFormatRoleAssignmentIdID(t *testing.T) {
	actual := NewRoleAssignmentIdID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group").ID()
	expected := "https://endpoint-url.example.com/roleAssignments/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRoleAssignmentIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleAssignmentIdId
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
			Input: "https://endpoint-url.example.com/roleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/roleAssignments/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Expected: &RoleAssignmentIdId{
				BaseURI:          "https://endpoint-url.example.com",
				RoleAssignmentId: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleAssignmentIdID(v.Input)
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

		if actual.RoleAssignmentId != v.Expected.RoleAssignmentId {
			t.Fatalf("Expected %q but got %q for RoleAssignmentId", v.Expected.RoleAssignmentId, actual.RoleAssignmentId)
		}

	}
}

func TestParseRoleAssignmentIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RoleAssignmentIdId
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
			Input: "https://endpoint-url.example.com/roleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/roleAssignments/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Expected: &RoleAssignmentIdId{
				BaseURI:          "https://endpoint-url.example.com",
				RoleAssignmentId: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			},
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/rOlEaSsIgNmEnTs/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Expected: &RoleAssignmentIdId{
				BaseURI:          "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				RoleAssignmentId: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRoleAssignmentIdIDInsensitively(v.Input)
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

		if actual.RoleAssignmentId != v.Expected.RoleAssignmentId {
			t.Fatalf("Expected %q but got %q for RoleAssignmentId", v.Expected.RoleAssignmentId, actual.RoleAssignmentId)
		}

	}
}

func TestSegmentsForRoleAssignmentIdId(t *testing.T) {
	segments := RoleAssignmentIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RoleAssignmentIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
