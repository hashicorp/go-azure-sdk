package serviceendpointpolicydefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServiceEndpointPolicyDefinitionId{}

func TestNewServiceEndpointPolicyDefinitionID(t *testing.T) {
	id := NewServiceEndpointPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceEndpointPolicyValue", "serviceEndpointPolicyDefinitionValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ServiceEndpointPolicyName != "serviceEndpointPolicyValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceEndpointPolicyName'", id.ServiceEndpointPolicyName, "serviceEndpointPolicyValue")
	}

	if id.ServiceEndpointPolicyDefinitionName != "serviceEndpointPolicyDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceEndpointPolicyDefinitionName'", id.ServiceEndpointPolicyDefinitionName, "serviceEndpointPolicyDefinitionValue")
	}
}

func TestFormatServiceEndpointPolicyDefinitionID(t *testing.T) {
	actual := NewServiceEndpointPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceEndpointPolicyValue", "serviceEndpointPolicyDefinitionValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions/serviceEndpointPolicyDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServiceEndpointPolicyDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServiceEndpointPolicyDefinitionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions/serviceEndpointPolicyDefinitionValue",
			Expected: &ServiceEndpointPolicyDefinitionId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                   "example-resource-group",
				ServiceEndpointPolicyName:           "serviceEndpointPolicyValue",
				ServiceEndpointPolicyDefinitionName: "serviceEndpointPolicyDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions/serviceEndpointPolicyDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServiceEndpointPolicyDefinitionID(v.Input)
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

		if actual.ServiceEndpointPolicyName != v.Expected.ServiceEndpointPolicyName {
			t.Fatalf("Expected %q but got %q for ServiceEndpointPolicyName", v.Expected.ServiceEndpointPolicyName, actual.ServiceEndpointPolicyName)
		}

		if actual.ServiceEndpointPolicyDefinitionName != v.Expected.ServiceEndpointPolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for ServiceEndpointPolicyDefinitionName", v.Expected.ServiceEndpointPolicyDefinitionName, actual.ServiceEndpointPolicyDefinitionName)
		}

	}
}

func TestParseServiceEndpointPolicyDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServiceEndpointPolicyDefinitionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/sErViCeEnDpOiNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/sErViCeEnDpOiNtPoLiCiEs/sErViCeEnDpOiNtPoLiCyVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/sErViCeEnDpOiNtPoLiCiEs/sErViCeEnDpOiNtPoLiCyVaLuE/sErViCeEnDpOiNtPoLiCyDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions/serviceEndpointPolicyDefinitionValue",
			Expected: &ServiceEndpointPolicyDefinitionId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                   "example-resource-group",
				ServiceEndpointPolicyName:           "serviceEndpointPolicyValue",
				ServiceEndpointPolicyDefinitionName: "serviceEndpointPolicyDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/serviceEndpointPolicies/serviceEndpointPolicyValue/serviceEndpointPolicyDefinitions/serviceEndpointPolicyDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/sErViCeEnDpOiNtPoLiCiEs/sErViCeEnDpOiNtPoLiCyVaLuE/sErViCeEnDpOiNtPoLiCyDeFiNiTiOnS/sErViCeEnDpOiNtPoLiCyDeFiNiTiOnVaLuE",
			Expected: &ServiceEndpointPolicyDefinitionId{
				SubscriptionId:                      "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                   "eXaMpLe-rEsOuRcE-GrOuP",
				ServiceEndpointPolicyName:           "sErViCeEnDpOiNtPoLiCyVaLuE",
				ServiceEndpointPolicyDefinitionName: "sErViCeEnDpOiNtPoLiCyDeFiNiTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtWoRk/sErViCeEnDpOiNtPoLiCiEs/sErViCeEnDpOiNtPoLiCyVaLuE/sErViCeEnDpOiNtPoLiCyDeFiNiTiOnS/sErViCeEnDpOiNtPoLiCyDeFiNiTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServiceEndpointPolicyDefinitionIDInsensitively(v.Input)
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

		if actual.ServiceEndpointPolicyName != v.Expected.ServiceEndpointPolicyName {
			t.Fatalf("Expected %q but got %q for ServiceEndpointPolicyName", v.Expected.ServiceEndpointPolicyName, actual.ServiceEndpointPolicyName)
		}

		if actual.ServiceEndpointPolicyDefinitionName != v.Expected.ServiceEndpointPolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for ServiceEndpointPolicyDefinitionName", v.Expected.ServiceEndpointPolicyDefinitionName, actual.ServiceEndpointPolicyDefinitionName)
		}

	}
}

func TestSegmentsForServiceEndpointPolicyDefinitionId(t *testing.T) {
	segments := ServiceEndpointPolicyDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServiceEndpointPolicyDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
