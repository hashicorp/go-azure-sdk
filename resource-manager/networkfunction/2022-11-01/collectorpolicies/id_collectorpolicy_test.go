package collectorpolicies

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CollectorPolicyId{}

func TestNewCollectorPolicyID(t *testing.T) {
	id := NewCollectorPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "azureTrafficCollectorName", "collectorPolicyName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.AzureTrafficCollectorName != "azureTrafficCollectorName" {
		t.Fatalf("Expected %q but got %q for Segment 'AzureTrafficCollectorName'", id.AzureTrafficCollectorName, "azureTrafficCollectorName")
	}

	if id.CollectorPolicyName != "collectorPolicyName" {
		t.Fatalf("Expected %q but got %q for Segment 'CollectorPolicyName'", id.CollectorPolicyName, "collectorPolicyName")
	}
}

func TestFormatCollectorPolicyID(t *testing.T) {
	actual := NewCollectorPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "azureTrafficCollectorName", "collectorPolicyName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies/collectorPolicyName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCollectorPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CollectorPolicyId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies/collectorPolicyName",
			Expected: &CollectorPolicyId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:         "example-resource-group",
				AzureTrafficCollectorName: "azureTrafficCollectorName",
				CollectorPolicyName:       "collectorPolicyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies/collectorPolicyName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCollectorPolicyID(v.Input)
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

		if actual.AzureTrafficCollectorName != v.Expected.AzureTrafficCollectorName {
			t.Fatalf("Expected %q but got %q for AzureTrafficCollectorName", v.Expected.AzureTrafficCollectorName, actual.AzureTrafficCollectorName)
		}

		if actual.CollectorPolicyName != v.Expected.CollectorPolicyName {
			t.Fatalf("Expected %q but got %q for CollectorPolicyName", v.Expected.CollectorPolicyName, actual.CollectorPolicyName)
		}

	}
}

func TestParseCollectorPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CollectorPolicyId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRkFuNcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRkFuNcTiOn/aZuReTrAfFiCcOlLeCtOrS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRkFuNcTiOn/aZuReTrAfFiCcOlLeCtOrS/aZuReTrAfFiCcOlLeCtOrNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRkFuNcTiOn/aZuReTrAfFiCcOlLeCtOrS/aZuReTrAfFiCcOlLeCtOrNaMe/cOlLeCtOrPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies/collectorPolicyName",
			Expected: &CollectorPolicyId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:         "example-resource-group",
				AzureTrafficCollectorName: "azureTrafficCollectorName",
				CollectorPolicyName:       "collectorPolicyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetworkFunction/azureTrafficCollectors/azureTrafficCollectorName/collectorPolicies/collectorPolicyName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRkFuNcTiOn/aZuReTrAfFiCcOlLeCtOrS/aZuReTrAfFiCcOlLeCtOrNaMe/cOlLeCtOrPoLiCiEs/cOlLeCtOrPoLiCyNaMe",
			Expected: &CollectorPolicyId{
				SubscriptionId:            "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:         "eXaMpLe-rEsOuRcE-GrOuP",
				AzureTrafficCollectorName: "aZuReTrAfFiCcOlLeCtOrNaMe",
				CollectorPolicyName:       "cOlLeCtOrPoLiCyNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRkFuNcTiOn/aZuReTrAfFiCcOlLeCtOrS/aZuReTrAfFiCcOlLeCtOrNaMe/cOlLeCtOrPoLiCiEs/cOlLeCtOrPoLiCyNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCollectorPolicyIDInsensitively(v.Input)
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

		if actual.AzureTrafficCollectorName != v.Expected.AzureTrafficCollectorName {
			t.Fatalf("Expected %q but got %q for AzureTrafficCollectorName", v.Expected.AzureTrafficCollectorName, actual.AzureTrafficCollectorName)
		}

		if actual.CollectorPolicyName != v.Expected.CollectorPolicyName {
			t.Fatalf("Expected %q but got %q for CollectorPolicyName", v.Expected.CollectorPolicyName, actual.CollectorPolicyName)
		}

	}
}

func TestSegmentsForCollectorPolicyId(t *testing.T) {
	segments := CollectorPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CollectorPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
