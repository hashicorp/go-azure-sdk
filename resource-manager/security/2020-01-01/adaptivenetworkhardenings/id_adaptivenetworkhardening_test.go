package adaptivenetworkhardenings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdaptiveNetworkHardeningId{}

func TestNewAdaptiveNetworkHardeningID(t *testing.T) {
	id := NewAdaptiveNetworkHardeningID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceTypeValue", "resourceValue", "adaptiveNetworkHardeningValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ProviderName != "providerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProviderName'", id.ProviderName, "providerValue")
	}

	if id.ResourceType != "resourceTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceType'", id.ResourceType, "resourceTypeValue")
	}

	if id.ResourceName != "resourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceName'", id.ResourceName, "resourceValue")
	}

	if id.AdaptiveNetworkHardeningName != "adaptiveNetworkHardeningValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AdaptiveNetworkHardeningName'", id.AdaptiveNetworkHardeningName, "adaptiveNetworkHardeningValue")
	}
}

func TestFormatAdaptiveNetworkHardeningID(t *testing.T) {
	actual := NewAdaptiveNetworkHardeningID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceTypeValue", "resourceValue", "adaptiveNetworkHardeningValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAdaptiveNetworkHardeningID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdaptiveNetworkHardeningId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue",
			Expected: &AdaptiveNetworkHardeningId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				ProviderName:                 "providerValue",
				ResourceType:                 "resourceTypeValue",
				ResourceName:                 "resourceValue",
				AdaptiveNetworkHardeningName: "adaptiveNetworkHardeningValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdaptiveNetworkHardeningID(v.Input)
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

		if actual.ResourceType != v.Expected.ResourceType {
			t.Fatalf("Expected %q but got %q for ResourceType", v.Expected.ResourceType, actual.ResourceType)
		}

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.AdaptiveNetworkHardeningName != v.Expected.AdaptiveNetworkHardeningName {
			t.Fatalf("Expected %q but got %q for AdaptiveNetworkHardeningName", v.Expected.AdaptiveNetworkHardeningName, actual.AdaptiveNetworkHardeningName)
		}

	}
}

func TestParseAdaptiveNetworkHardeningIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AdaptiveNetworkHardeningId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.sEcUrItY/aDaPtIvEnEtWoRkHaRdEnInGs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue",
			Expected: &AdaptiveNetworkHardeningId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "example-resource-group",
				ProviderName:                 "providerValue",
				ResourceType:                 "resourceTypeValue",
				ResourceName:                 "resourceValue",
				AdaptiveNetworkHardeningName: "adaptiveNetworkHardeningValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceTypeValue/resourceValue/providers/Microsoft.Security/adaptiveNetworkHardenings/adaptiveNetworkHardeningValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.sEcUrItY/aDaPtIvEnEtWoRkHaRdEnInGs/aDaPtIvEnEtWoRkHaRdEnInGvAlUe",
			Expected: &AdaptiveNetworkHardeningId{
				SubscriptionId:               "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:            "eXaMpLe-rEsOuRcE-GrOuP",
				ProviderName:                 "pRoViDeRvAlUe",
				ResourceType:                 "rEsOuRcEtYpEvAlUe",
				ResourceName:                 "rEsOuRcEvAlUe",
				AdaptiveNetworkHardeningName: "aDaPtIvEnEtWoRkHaRdEnInGvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.sEcUrItY/aDaPtIvEnEtWoRkHaRdEnInGs/aDaPtIvEnEtWoRkHaRdEnInGvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAdaptiveNetworkHardeningIDInsensitively(v.Input)
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

		if actual.ResourceType != v.Expected.ResourceType {
			t.Fatalf("Expected %q but got %q for ResourceType", v.Expected.ResourceType, actual.ResourceType)
		}

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.AdaptiveNetworkHardeningName != v.Expected.AdaptiveNetworkHardeningName {
			t.Fatalf("Expected %q but got %q for AdaptiveNetworkHardeningName", v.Expected.AdaptiveNetworkHardeningName, actual.AdaptiveNetworkHardeningName)
		}

	}
}

func TestSegmentsForAdaptiveNetworkHardeningId(t *testing.T) {
	segments := AdaptiveNetworkHardeningId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AdaptiveNetworkHardeningId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
