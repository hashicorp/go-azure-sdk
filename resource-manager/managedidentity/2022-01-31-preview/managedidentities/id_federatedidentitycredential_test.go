package managedidentities

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = FederatedIdentityCredentialId{}

func TestNewFederatedIdentityCredentialID(t *testing.T) {
	id := NewFederatedIdentityCredentialID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "federatedIdentityCredentialResourceValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ResourceName != "resourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceName'", id.ResourceName, "resourceValue")
	}

	if id.FederatedIdentityCredentialResourceName != "federatedIdentityCredentialResourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FederatedIdentityCredentialResourceName'", id.FederatedIdentityCredentialResourceName, "federatedIdentityCredentialResourceValue")
	}
}

func TestFormatFederatedIdentityCredentialID(t *testing.T) {
	actual := NewFederatedIdentityCredentialID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "federatedIdentityCredentialResourceValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials/federatedIdentityCredentialResourceValue"
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials/federatedIdentityCredentialResourceValue",
			Expected: &FederatedIdentityCredentialId{
				SubscriptionId:                          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                       "example-resource-group",
				ResourceName:                            "resourceValue",
				FederatedIdentityCredentialResourceName: "federatedIdentityCredentialResourceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials/federatedIdentityCredentialResourceValue/extra",
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

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.FederatedIdentityCredentialResourceName != v.Expected.FederatedIdentityCredentialResourceName {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialResourceName", v.Expected.FederatedIdentityCredentialResourceName, actual.FederatedIdentityCredentialResourceName)
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/rEsOuRcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/rEsOuRcEvAlUe/fEdErAtEdIdEnTiTyCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials/federatedIdentityCredentialResourceValue",
			Expected: &FederatedIdentityCredentialId{
				SubscriptionId:                          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                       "example-resource-group",
				ResourceName:                            "resourceValue",
				FederatedIdentityCredentialResourceName: "federatedIdentityCredentialResourceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/resourceValue/federatedIdentityCredentials/federatedIdentityCredentialResourceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/rEsOuRcEvAlUe/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlReSoUrCeVaLuE",
			Expected: &FederatedIdentityCredentialId{
				SubscriptionId:                          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                       "eXaMpLe-rEsOuRcE-GrOuP",
				ResourceName:                            "rEsOuRcEvAlUe",
				FederatedIdentityCredentialResourceName: "fEdErAtEdIdEnTiTyCrEdEnTiAlReSoUrCeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mAnAgEdIdEnTiTy/uSeRaSsIgNeDiDeNtItIeS/rEsOuRcEvAlUe/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlReSoUrCeVaLuE/extra",
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

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.FederatedIdentityCredentialResourceName != v.Expected.FederatedIdentityCredentialResourceName {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialResourceName", v.Expected.FederatedIdentityCredentialResourceName, actual.FederatedIdentityCredentialResourceName)
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
