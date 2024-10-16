package policystates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedPolicyStatesResourceId{}

func TestNewScopedPolicyStatesResourceID(t *testing.T) {
	id := NewScopedPolicyStatesResourceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "default")

	if id.ResourceId != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceId'", id.ResourceId, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.PolicyStatesResource != "default" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyStatesResource'", id.PolicyStatesResource, "default")
	}
}

func TestFormatScopedPolicyStatesResourceID(t *testing.T) {
	actual := NewScopedPolicyStatesResourceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "default").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates/default"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedPolicyStatesResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedPolicyStatesResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &ScopedPolicyStatesResourceId{
				ResourceId:           "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedPolicyStatesResourceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceId != v.Expected.ResourceId {
			t.Fatalf("Expected %q but got %q for ResourceId", v.Expected.ResourceId, actual.ResourceId)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestParseScopedPolicyStatesResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedPolicyStatesResourceId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &ScopedPolicyStatesResourceId{
				ResourceId:           "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt",
			Expected: &ScopedPolicyStatesResourceId{
				ResourceId:           "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedPolicyStatesResourceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceId != v.Expected.ResourceId {
			t.Fatalf("Expected %q but got %q for ResourceId", v.Expected.ResourceId, actual.ResourceId)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestSegmentsForScopedPolicyStatesResourceId(t *testing.T) {
	segments := ScopedPolicyStatesResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedPolicyStatesResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
