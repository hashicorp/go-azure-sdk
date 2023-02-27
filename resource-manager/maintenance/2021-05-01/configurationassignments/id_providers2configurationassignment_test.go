// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configurationassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2ConfigurationAssignmentId{}

func TestNewProviders2ConfigurationAssignmentID(t *testing.T) {
	id := NewProviders2ConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceParentTypeValue", "resourceParentValue", "resourceTypeValue", "resourceValue", "configurationAssignmentValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ProviderName != "providerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProviderName'", id.ProviderName, "providerValue")
	}

	if id.ResourceParentType != "resourceParentTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceParentType'", id.ResourceParentType, "resourceParentTypeValue")
	}

	if id.ResourceParentName != "resourceParentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceParentName'", id.ResourceParentName, "resourceParentValue")
	}

	if id.ResourceType != "resourceTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceType'", id.ResourceType, "resourceTypeValue")
	}

	if id.ResourceName != "resourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceName'", id.ResourceName, "resourceValue")
	}

	if id.ConfigurationAssignmentName != "configurationAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ConfigurationAssignmentName'", id.ConfigurationAssignmentName, "configurationAssignmentValue")
	}
}

func TestFormatProviders2ConfigurationAssignmentID(t *testing.T) {
	actual := NewProviders2ConfigurationAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceParentTypeValue", "resourceParentValue", "resourceTypeValue", "resourceValue", "configurationAssignmentValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments/configurationAssignmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2ConfigurationAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2ConfigurationAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments/configurationAssignmentValue",
			Expected: &Providers2ConfigurationAssignmentId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				ProviderName:                "providerValue",
				ResourceParentType:          "resourceParentTypeValue",
				ResourceParentName:          "resourceParentValue",
				ResourceType:                "resourceTypeValue",
				ResourceName:                "resourceValue",
				ConfigurationAssignmentName: "configurationAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments/configurationAssignmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2ConfigurationAssignmentID(v.Input)
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

		if actual.ResourceParentType != v.Expected.ResourceParentType {
			t.Fatalf("Expected %q but got %q for ResourceParentType", v.Expected.ResourceParentType, actual.ResourceParentType)
		}

		if actual.ResourceParentName != v.Expected.ResourceParentName {
			t.Fatalf("Expected %q but got %q for ResourceParentName", v.Expected.ResourceParentName, actual.ResourceParentName)
		}

		if actual.ResourceType != v.Expected.ResourceType {
			t.Fatalf("Expected %q but got %q for ResourceType", v.Expected.ResourceType, actual.ResourceType)
		}

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.ConfigurationAssignmentName != v.Expected.ConfigurationAssignmentName {
			t.Fatalf("Expected %q but got %q for ConfigurationAssignmentName", v.Expected.ConfigurationAssignmentName, actual.ConfigurationAssignmentName)
		}

	}
}

func TestParseProviders2ConfigurationAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2ConfigurationAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/cOnFiGuRaTiOnAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments/configurationAssignmentValue",
			Expected: &Providers2ConfigurationAssignmentId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "example-resource-group",
				ProviderName:                "providerValue",
				ResourceParentType:          "resourceParentTypeValue",
				ResourceParentName:          "resourceParentValue",
				ResourceType:                "resourceTypeValue",
				ResourceName:                "resourceValue",
				ConfigurationAssignmentName: "configurationAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/configurationAssignments/configurationAssignmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/cOnFiGuRaTiOnAsSiGnMeNtS/cOnFiGuRaTiOnAsSiGnMeNtVaLuE",
			Expected: &Providers2ConfigurationAssignmentId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:           "eXaMpLe-rEsOuRcE-GrOuP",
				ProviderName:                "pRoViDeRvAlUe",
				ResourceParentType:          "rEsOuRcEpArEnTtYpEvAlUe",
				ResourceParentName:          "rEsOuRcEpArEnTvAlUe",
				ResourceType:                "rEsOuRcEtYpEvAlUe",
				ResourceName:                "rEsOuRcEvAlUe",
				ConfigurationAssignmentName: "cOnFiGuRaTiOnAsSiGnMeNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/cOnFiGuRaTiOnAsSiGnMeNtS/cOnFiGuRaTiOnAsSiGnMeNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2ConfigurationAssignmentIDInsensitively(v.Input)
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

		if actual.ResourceParentType != v.Expected.ResourceParentType {
			t.Fatalf("Expected %q but got %q for ResourceParentType", v.Expected.ResourceParentType, actual.ResourceParentType)
		}

		if actual.ResourceParentName != v.Expected.ResourceParentName {
			t.Fatalf("Expected %q but got %q for ResourceParentName", v.Expected.ResourceParentName, actual.ResourceParentName)
		}

		if actual.ResourceType != v.Expected.ResourceType {
			t.Fatalf("Expected %q but got %q for ResourceType", v.Expected.ResourceType, actual.ResourceType)
		}

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.ConfigurationAssignmentName != v.Expected.ConfigurationAssignmentName {
			t.Fatalf("Expected %q but got %q for ConfigurationAssignmentName", v.Expected.ConfigurationAssignmentName, actual.ConfigurationAssignmentName)
		}

	}
}

func TestSegmentsForProviders2ConfigurationAssignmentId(t *testing.T) {
	segments := Providers2ConfigurationAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2ConfigurationAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
