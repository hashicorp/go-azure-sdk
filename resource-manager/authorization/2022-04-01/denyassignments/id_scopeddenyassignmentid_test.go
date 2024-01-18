package denyassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedDenyAssignmentIdId{}

func TestNewScopedDenyAssignmentIdID(t *testing.T) {
	id := NewScopedDenyAssignmentIdID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.DenyAssignmentId != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'DenyAssignmentId'", id.DenyAssignmentId, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}
}

func TestFormatScopedDenyAssignmentIdID(t *testing.T) {
	actual := NewScopedDenyAssignmentIdID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/denyAssignments/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedDenyAssignmentIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedDenyAssignmentIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/denyAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/denyAssignments/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Expected: &ScopedDenyAssignmentIdId{
				Scope:            "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				DenyAssignmentId: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedDenyAssignmentIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.DenyAssignmentId != v.Expected.DenyAssignmentId {
			t.Fatalf("Expected %q but got %q for DenyAssignmentId", v.Expected.DenyAssignmentId, actual.DenyAssignmentId)
		}

	}
}

func TestParseScopedDenyAssignmentIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedDenyAssignmentIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/denyAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/dEnYaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Authorization/denyAssignments/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Expected: &ScopedDenyAssignmentIdId{
				Scope:            "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				DenyAssignmentId: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			},
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/dEnYaSsIgNmEnTs/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Expected: &ScopedDenyAssignmentIdId{
				Scope:            "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				DenyAssignmentId: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedDenyAssignmentIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.DenyAssignmentId != v.Expected.DenyAssignmentId {
			t.Fatalf("Expected %q but got %q for DenyAssignmentId", v.Expected.DenyAssignmentId, actual.DenyAssignmentId)
		}

	}
}

func TestSegmentsForScopedDenyAssignmentIdId(t *testing.T) {
	segments := ScopedDenyAssignmentIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedDenyAssignmentIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
