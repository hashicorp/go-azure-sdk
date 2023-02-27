// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sourcecontrolsyncjobstreams

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SourceControlSyncJobId{}

func TestNewSourceControlSyncJobID(t *testing.T) {
	id := NewSourceControlSyncJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "sourceControlValue", "sourceControlSyncJobIdValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.AutomationAccountName != "automationAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AutomationAccountName'", id.AutomationAccountName, "automationAccountValue")
	}

	if id.SourceControlName != "sourceControlValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SourceControlName'", id.SourceControlName, "sourceControlValue")
	}

	if id.SourceControlSyncJobId != "sourceControlSyncJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SourceControlSyncJobId'", id.SourceControlSyncJobId, "sourceControlSyncJobIdValue")
	}
}

func TestFormatSourceControlSyncJobID(t *testing.T) {
	actual := NewSourceControlSyncJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "sourceControlValue", "sourceControlSyncJobIdValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs/sourceControlSyncJobIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSourceControlSyncJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SourceControlSyncJobId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs/sourceControlSyncJobIdValue",
			Expected: &SourceControlSyncJobId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				AutomationAccountName:  "automationAccountValue",
				SourceControlName:      "sourceControlValue",
				SourceControlSyncJobId: "sourceControlSyncJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs/sourceControlSyncJobIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSourceControlSyncJobID(v.Input)
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

		if actual.SourceControlName != v.Expected.SourceControlName {
			t.Fatalf("Expected %q but got %q for SourceControlName", v.Expected.SourceControlName, actual.SourceControlName)
		}

		if actual.SourceControlSyncJobId != v.Expected.SourceControlSyncJobId {
			t.Fatalf("Expected %q but got %q for SourceControlSyncJobId", v.Expected.SourceControlSyncJobId, actual.SourceControlSyncJobId)
		}

	}
}

func TestParseSourceControlSyncJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SourceControlSyncJobId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOuRcEcOnTrOlS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOuRcEcOnTrOlS/sOuRcEcOnTrOlVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOuRcEcOnTrOlS/sOuRcEcOnTrOlVaLuE/sOuRcEcOnTrOlSyNcJoBs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs/sourceControlSyncJobIdValue",
			Expected: &SourceControlSyncJobId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				AutomationAccountName:  "automationAccountValue",
				SourceControlName:      "sourceControlValue",
				SourceControlSyncJobId: "sourceControlSyncJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Automation/automationAccounts/automationAccountValue/sourceControls/sourceControlValue/sourceControlSyncJobs/sourceControlSyncJobIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOuRcEcOnTrOlS/sOuRcEcOnTrOlVaLuE/sOuRcEcOnTrOlSyNcJoBs/sOuRcEcOnTrOlSyNcJoBiDvAlUe",
			Expected: &SourceControlSyncJobId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "eXaMpLe-rEsOuRcE-GrOuP",
				AutomationAccountName:  "aUtOmAtIoNaCcOuNtVaLuE",
				SourceControlName:      "sOuRcEcOnTrOlVaLuE",
				SourceControlSyncJobId: "sOuRcEcOnTrOlSyNcJoBiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aUtOmAtIoN/aUtOmAtIoNaCcOuNtS/aUtOmAtIoNaCcOuNtVaLuE/sOuRcEcOnTrOlS/sOuRcEcOnTrOlVaLuE/sOuRcEcOnTrOlSyNcJoBs/sOuRcEcOnTrOlSyNcJoBiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSourceControlSyncJobIDInsensitively(v.Input)
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

		if actual.SourceControlName != v.Expected.SourceControlName {
			t.Fatalf("Expected %q but got %q for SourceControlName", v.Expected.SourceControlName, actual.SourceControlName)
		}

		if actual.SourceControlSyncJobId != v.Expected.SourceControlSyncJobId {
			t.Fatalf("Expected %q but got %q for SourceControlSyncJobId", v.Expected.SourceControlSyncJobId, actual.SourceControlSyncJobId)
		}

	}
}

func TestSegmentsForSourceControlSyncJobId(t *testing.T) {
	segments := SourceControlSyncJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SourceControlSyncJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
