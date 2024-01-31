package provisionedclusterinstances

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedAgentPoolId{}

func TestNewScopedAgentPoolID(t *testing.T) {
	id := NewScopedAgentPoolID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "agentPoolValue")

	if id.ConnectedClusterResourceUri != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ConnectedClusterResourceUri'", id.ConnectedClusterResourceUri, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.AgentPoolName != "agentPoolValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgentPoolName'", id.AgentPoolName, "agentPoolValue")
	}
}

func TestFormatScopedAgentPoolID(t *testing.T) {
	actual := NewScopedAgentPoolID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "agentPoolValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/agentPoolValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedAgentPoolID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedAgentPoolId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/agentPoolValue",
			Expected: &ScopedAgentPoolId{
				ConnectedClusterResourceUri: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AgentPoolName:               "agentPoolValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/agentPoolValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedAgentPoolID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConnectedClusterResourceUri != v.Expected.ConnectedClusterResourceUri {
			t.Fatalf("Expected %q but got %q for ConnectedClusterResourceUri", v.Expected.ConnectedClusterResourceUri, actual.ConnectedClusterResourceUri)
		}

		if actual.AgentPoolName != v.Expected.AgentPoolName {
			t.Fatalf("Expected %q but got %q for AgentPoolName", v.Expected.AgentPoolName, actual.AgentPoolName)
		}

	}
}

func TestParseScopedAgentPoolIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedAgentPoolId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.hYbRiDcOnTaInErSeRvIcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.hYbRiDcOnTaInErSeRvIcE/pRoViSiOnEdClUsTeRiNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.hYbRiDcOnTaInErSeRvIcE/pRoViSiOnEdClUsTeRiNsTaNcEs/dEfAuLt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.hYbRiDcOnTaInErSeRvIcE/pRoViSiOnEdClUsTeRiNsTaNcEs/dEfAuLt/aGeNtPoOlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/agentPoolValue",
			Expected: &ScopedAgentPoolId{
				ConnectedClusterResourceUri: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AgentPoolName:               "agentPoolValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/agentPools/agentPoolValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.hYbRiDcOnTaInErSeRvIcE/pRoViSiOnEdClUsTeRiNsTaNcEs/dEfAuLt/aGeNtPoOlS/aGeNtPoOlVaLuE",
			Expected: &ScopedAgentPoolId{
				ConnectedClusterResourceUri: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				AgentPoolName:               "aGeNtPoOlVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.hYbRiDcOnTaInErSeRvIcE/pRoViSiOnEdClUsTeRiNsTaNcEs/dEfAuLt/aGeNtPoOlS/aGeNtPoOlVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedAgentPoolIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ConnectedClusterResourceUri != v.Expected.ConnectedClusterResourceUri {
			t.Fatalf("Expected %q but got %q for ConnectedClusterResourceUri", v.Expected.ConnectedClusterResourceUri, actual.ConnectedClusterResourceUri)
		}

		if actual.AgentPoolName != v.Expected.AgentPoolName {
			t.Fatalf("Expected %q but got %q for AgentPoolName", v.Expected.AgentPoolName, actual.AgentPoolName)
		}

	}
}

func TestSegmentsForScopedAgentPoolId(t *testing.T) {
	segments := ScopedAgentPoolId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedAgentPoolId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
