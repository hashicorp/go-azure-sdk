// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flux

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ProviderId{}

func TestNewProviderID(t *testing.T) {
	id := NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "clusterResourceValue", "clusterValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ProviderName != "providerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProviderName'", id.ProviderName, "providerValue")
	}

	if id.ClusterResourceName != "clusterResourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ClusterResourceName'", id.ClusterResourceName, "clusterResourceValue")
	}

	if id.ClusterName != "clusterValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ClusterName'", id.ClusterName, "clusterValue")
	}
}

func TestFormatProviderID(t *testing.T) {
	actual := NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "clusterResourceValue", "clusterValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue/clusterValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue/clusterValue",
			Expected: &ProviderId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				ProviderName:        "providerValue",
				ClusterResourceName: "clusterResourceValue",
				ClusterName:         "clusterValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue/clusterValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.ProviderName != v.Expected.ProviderName {
			t.Fatalf("Expected %q but got %q for ProviderName", v.Expected.ProviderName, actual.ProviderName)
		}

		if actual.ClusterResourceName != v.Expected.ClusterResourceName {
			t.Fatalf("Expected %q but got %q for ClusterResourceName", v.Expected.ClusterResourceName, actual.ClusterResourceName)
		}

		if actual.ClusterName != v.Expected.ClusterName {
			t.Fatalf("Expected %q but got %q for ClusterName", v.Expected.ClusterName, actual.ClusterName)
		}

	}
}

func TestParseProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/cLuStErReSoUrCeVaLuE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue/clusterValue",
			Expected: &ProviderId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "example-resource-group",
				ProviderName:        "providerValue",
				ClusterResourceName: "clusterResourceValue",
				ClusterName:         "clusterValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/clusterResourceValue/clusterValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/cLuStErReSoUrCeVaLuE/cLuStErVaLuE",
			Expected: &ProviderId{
				SubscriptionId:      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:   "eXaMpLe-rEsOuRcE-GrOuP",
				ProviderName:        "pRoViDeRvAlUe",
				ClusterResourceName: "cLuStErReSoUrCeVaLuE",
				ClusterName:         "cLuStErVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/cLuStErReSoUrCeVaLuE/cLuStErVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.ProviderName != v.Expected.ProviderName {
			t.Fatalf("Expected %q but got %q for ProviderName", v.Expected.ProviderName, actual.ProviderName)
		}

		if actual.ClusterResourceName != v.Expected.ClusterResourceName {
			t.Fatalf("Expected %q but got %q for ClusterResourceName", v.Expected.ClusterResourceName, actual.ClusterResourceName)
		}

		if actual.ClusterName != v.Expected.ClusterName {
			t.Fatalf("Expected %q but got %q for ClusterName", v.Expected.ClusterName, actual.ClusterName)
		}

	}
}

func TestSegmentsForProviderId(t *testing.T) {
	segments := ProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
