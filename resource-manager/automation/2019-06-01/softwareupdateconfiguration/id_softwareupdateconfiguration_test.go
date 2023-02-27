// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package softwareupdateconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SoftwareUpdateConfigurationId{}

func TestNewSoftwareUpdateConfigurationID(t *testing.T) {
	id := NewSoftwareUpdateConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "softwareUpdateConfigurationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.AutomationAccountName != "automationAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AutomationAccountName'", id.AutomationAccountName, "automationAccountValue")
	}

	if id.SoftwareUpdateConfigurationName != "softwareUpdateConfigurationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SoftwareUpdateConfigurationName'", id.SoftwareUpdateConfigurationName, "softwareUpdateConfigurationValue")
	}
}

func TestFormatSoftwareUpdateConfigurationID(t *testing.T) {
	actual := NewSoftwareUpdateConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "softwareUpdateConfigurationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations/softwareUpdateConfigurationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSoftwareUpdateConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SoftwareUpdateConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations/softwareUpdateConfigurationValue",
			Expected: &SoftwareUpdateConfigurationId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "example-resource-group",
				AutomationAccountName:           "automationAccountValue",
				SoftwareUpdateConfigurationName: "softwareUpdateConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations/softwareUpdateConfigurationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSoftwareUpdateConfigurationID(v.Input)
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

		if actual.AutomationAccountName != v.Expected.AutomationAccountName {
			t.Fatalf("Expected %q but got %q for AutomationAccountName", v.Expected.AutomationAccountName, actual.AutomationAccountName)
		}

		if actual.SoftwareUpdateConfigurationName != v.Expected.SoftwareUpdateConfigurationName {
			t.Fatalf("Expected %q but got %q for SoftwareUpdateConfigurationName", v.Expected.SoftwareUpdateConfigurationName, actual.SoftwareUpdateConfigurationName)
		}

	}
}

func TestParseSoftwareUpdateConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SoftwareUpdateConfigurationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOfTwArEuPdAtEcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations/softwareUpdateConfigurationValue",
			Expected: &SoftwareUpdateConfigurationId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "example-resource-group",
				AutomationAccountName:           "automationAccountValue",
				SoftwareUpdateConfigurationName: "softwareUpdateConfigurationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/softwareUpdateConfigurations/softwareUpdateConfigurationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOfTwArEuPdAtEcOnFiGuRaTiOnS/sOfTwArEuPdAtEcOnFiGuRaTiOnVaLuE",
			Expected: &SoftwareUpdateConfigurationId{
				SubscriptionId:                  "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:               "eXaMpLe-rEsOuRcE-GrOuP",
				AutomationAccountName:           "aUtOmAtIoNaCcOuNtVaLuE",
				SoftwareUpdateConfigurationName: "sOfTwArEuPdAtEcOnFiGuRaTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOfTwArEuPdAtEcOnFiGuRaTiOnS/sOfTwArEuPdAtEcOnFiGuRaTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSoftwareUpdateConfigurationIDInsensitively(v.Input)
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

		if actual.AutomationAccountName != v.Expected.AutomationAccountName {
			t.Fatalf("Expected %q but got %q for AutomationAccountName", v.Expected.AutomationAccountName, actual.AutomationAccountName)
		}

		if actual.SoftwareUpdateConfigurationName != v.Expected.SoftwareUpdateConfigurationName {
			t.Fatalf("Expected %q but got %q for SoftwareUpdateConfigurationName", v.Expected.SoftwareUpdateConfigurationName, actual.SoftwareUpdateConfigurationName)
		}

	}
}

func TestSegmentsForSoftwareUpdateConfigurationId(t *testing.T) {
	segments := SoftwareUpdateConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SoftwareUpdateConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
