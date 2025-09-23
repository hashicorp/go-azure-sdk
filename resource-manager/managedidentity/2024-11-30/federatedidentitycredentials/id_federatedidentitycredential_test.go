package federatedidentitycredentials

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FederatedIdentityCredentialId{}

func TestNewFederatedIdentityCredentialID(t *testing.T) {
	id := NewFederatedIdentityCredentialID("12345678-1234-9876-4563-123456789012", "example-resource-group", "userAssignedIdentityName", "federatedIdentityCredentialName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.UserAssignedIdentityName != "userAssignedIdentityName" {
		t.Fatalf("Expected %q but got %q for Segment 'UserAssignedIdentityName'", id.UserAssignedIdentityName, "userAssignedIdentityName")
	}

	if id.FederatedIdentityCredentialName != "federatedIdentityCredentialName" {
		t.Fatalf("Expected %q but got %q for Segment 'FederatedIdentityCredentialName'", id.FederatedIdentityCredentialName, "federatedIdentityCredentialName")
	}
}

func TestFormatFederatedIdentityCredentialID(t *testing.T) {
	actual := NewFederatedIdentityCredentialID("12345678-1234-9876-4563-123456789012", "example-resource-group", "userAssignedIdentityName", "federatedIdentityCredentialName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials/federatedIdentityCredentialName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseFederatedIdentityCredentialID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FederatedIdentityCredentialId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials/federatedIdentityCredentialName",
			Expected: &FederatedIdentityCredentialId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "example-resource-group",
				UserAssignedIdentityName:        "userAssignedIdentityName",
				FederatedIdentityCredentialName: "federatedIdentityCredentialName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials/federatedIdentityCredentialName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFederatedIdentityCredentialID(v.Input)
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

		if actual.UserAssignedIdentityName != v.Expected.UserAssignedIdentityName {
			t.Fatalf("Expected %q but got %q for UserAssignedIdentityName", v.Expected.UserAssignedIdentityName, actual.UserAssignedIdentityName)
		}

		if actual.FederatedIdentityCredentialName != v.Expected.FederatedIdentityCredentialName {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialName", v.Expected.FederatedIdentityCredentialName, actual.FederatedIdentityCredentialName)
		}

	}
}

func TestParseFederatedIdentityCredentialIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FederatedIdentityCredentialId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/uSeRaSsIgNeDiDeNtItYnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/uSeRaSsIgNeDiDeNtItYnAmE/fEdErAtEdIdEnTiTyCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials/federatedIdentityCredentialName",
			Expected: &FederatedIdentityCredentialId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "example-resource-group",
				UserAssignedIdentityName:        "userAssignedIdentityName",
				FederatedIdentityCredentialName: "federatedIdentityCredentialName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userAssignedIdentityName/federatedIdentityCredentials/federatedIdentityCredentialName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/uSeRaSsIgNeDiDeNtItYnAmE/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlNaMe",
			Expected: &FederatedIdentityCredentialId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "eXaMpLe-rEsOuRcE-GrOuP",
				UserAssignedIdentityName:        "uSeRaSsIgNeDiDeNtItYnAmE",
				FederatedIdentityCredentialName: "fEdErAtEdIdEnTiTyCrEdEnTiAlNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/uSeRaSsIgNeDiDeNtItYnAmE/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFederatedIdentityCredentialIDInsensitively(v.Input)
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

		if actual.UserAssignedIdentityName != v.Expected.UserAssignedIdentityName {
			t.Fatalf("Expected %q but got %q for UserAssignedIdentityName", v.Expected.UserAssignedIdentityName, actual.UserAssignedIdentityName)
		}

		if actual.FederatedIdentityCredentialName != v.Expected.FederatedIdentityCredentialName {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialName", v.Expected.FederatedIdentityCredentialName, actual.FederatedIdentityCredentialName)
		}

	}
}

func TestSegmentsForFederatedIdentityCredentialId(t *testing.T) {
	segments := FederatedIdentityCredentialId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("FederatedIdentityCredentialId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
