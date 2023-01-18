package integrationserviceenvironmentmanagedapi

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ManagedApiId{}

func TestNewManagedApiID(t *testing.T) {
	id := NewManagedApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue", "apiValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroup != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroup'", id.ResourceGroup, "example-resource-group")
	}

	if id.IntegrationServiceEnvironmentName != "integrationServiceEnvironmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'IntegrationServiceEnvironmentName'", id.IntegrationServiceEnvironmentName, "integrationServiceEnvironmentValue")
	}

	if id.ApiName != "apiValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApiName'", id.ApiName, "apiValue")
	}
}

func TestFormatManagedApiID(t *testing.T) {
	actual := NewManagedApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue", "apiValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis/apiValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseManagedApiID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedApiId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis/apiValue",
			Expected: &ManagedApiId{
				SubscriptionId:                    "12345678-1234-9876-4563-123456789012",
				ResourceGroup:                     "example-resource-group",
				IntegrationServiceEnvironmentName: "integrationServiceEnvironmentValue",
				ApiName:                           "apiValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis/apiValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedApiID(v.Input)
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

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}

		if actual.IntegrationServiceEnvironmentName != v.Expected.IntegrationServiceEnvironmentName {
			t.Fatalf("Expected %q but got %q for IntegrationServiceEnvironmentName", v.Expected.IntegrationServiceEnvironmentName, actual.IntegrationServiceEnvironmentName)
		}

		if actual.ApiName != v.Expected.ApiName {
			t.Fatalf("Expected %q but got %q for ApiName", v.Expected.ApiName, actual.ApiName)
		}

	}
}

func TestParseManagedApiIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ManagedApiId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.lOgIc",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.lOgIc/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.lOgIc/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtS/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.lOgIc/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtS/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtVaLuE/mAnAgEdApIs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis/apiValue",
			Expected: &ManagedApiId{
				SubscriptionId:                    "12345678-1234-9876-4563-123456789012",
				ResourceGroup:                     "example-resource-group",
				IntegrationServiceEnvironmentName: "integrationServiceEnvironmentValue",
				ApiName:                           "apiValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/integrationServiceEnvironmentValue/managedApis/apiValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.lOgIc/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtS/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtVaLuE/mAnAgEdApIs/aPiVaLuE",
			Expected: &ManagedApiId{
				SubscriptionId:                    "12345678-1234-9876-4563-123456789012",
				ResourceGroup:                     "eXaMpLe-rEsOuRcE-GrOuP",
				IntegrationServiceEnvironmentName: "iNtEgRaTiOnSeRvIcEeNvIrOnMeNtVaLuE",
				ApiName:                           "aPiVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.lOgIc/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtS/iNtEgRaTiOnSeRvIcEeNvIrOnMeNtVaLuE/mAnAgEdApIs/aPiVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseManagedApiIDInsensitively(v.Input)
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

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}

		if actual.IntegrationServiceEnvironmentName != v.Expected.IntegrationServiceEnvironmentName {
			t.Fatalf("Expected %q but got %q for IntegrationServiceEnvironmentName", v.Expected.IntegrationServiceEnvironmentName, actual.IntegrationServiceEnvironmentName)
		}

		if actual.ApiName != v.Expected.ApiName {
			t.Fatalf("Expected %q but got %q for ApiName", v.Expected.ApiName, actual.ApiName)
		}

	}
}

func TestSegmentsForManagedApiId(t *testing.T) {
	segments := ManagedApiId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ManagedApiId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
