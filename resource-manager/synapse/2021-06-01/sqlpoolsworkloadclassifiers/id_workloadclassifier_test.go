package sqlpoolsworkloadclassifiers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &WorkloadClassifierId{}

func TestNewWorkloadClassifierID(t *testing.T) {
	id := NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue", "workloadClassifierValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.WorkspaceName != "workspaceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkspaceName'", id.WorkspaceName, "workspaceValue")
	}

	if id.SqlPoolName != "sqlPoolValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SqlPoolName'", id.SqlPoolName, "sqlPoolValue")
	}

	if id.WorkloadGroupName != "workloadGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkloadGroupName'", id.WorkloadGroupName, "workloadGroupValue")
	}

	if id.WorkloadClassifierName != "workloadClassifierValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkloadClassifierName'", id.WorkloadClassifierName, "workloadClassifierValue")
	}
}

func TestFormatWorkloadClassifierID(t *testing.T) {
	actual := NewWorkloadClassifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "workloadGroupValue", "workloadClassifierValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers/workloadClassifierValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseWorkloadClassifierID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *WorkloadClassifierId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers/workloadClassifierValue",
			Expected: &WorkloadClassifierId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				WorkspaceName:          "workspaceValue",
				SqlPoolName:            "sqlPoolValue",
				WorkloadGroupName:      "workloadGroupValue",
				WorkloadClassifierName: "workloadClassifierValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers/workloadClassifierValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseWorkloadClassifierID(v.Input)
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

		if actual.WorkspaceName != v.Expected.WorkspaceName {
			t.Fatalf("Expected %q but got %q for WorkspaceName", v.Expected.WorkspaceName, actual.WorkspaceName)
		}

		if actual.SqlPoolName != v.Expected.SqlPoolName {
			t.Fatalf("Expected %q but got %q for SqlPoolName", v.Expected.SqlPoolName, actual.SqlPoolName)
		}

		if actual.WorkloadGroupName != v.Expected.WorkloadGroupName {
			t.Fatalf("Expected %q but got %q for WorkloadGroupName", v.Expected.WorkloadGroupName, actual.WorkloadGroupName)
		}

		if actual.WorkloadClassifierName != v.Expected.WorkloadClassifierName {
			t.Fatalf("Expected %q but got %q for WorkloadClassifierName", v.Expected.WorkloadClassifierName, actual.WorkloadClassifierName)
		}

	}
}

func TestParseWorkloadClassifierIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *WorkloadClassifierId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS/sQlPoOlVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS/sQlPoOlVaLuE/wOrKlOaDgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS/sQlPoOlVaLuE/wOrKlOaDgRoUpS/wOrKlOaDgRoUpVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS/sQlPoOlVaLuE/wOrKlOaDgRoUpS/wOrKlOaDgRoUpVaLuE/wOrKlOaDcLaSsIfIeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers/workloadClassifierValue",
			Expected: &WorkloadClassifierId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				WorkspaceName:          "workspaceValue",
				SqlPoolName:            "sqlPoolValue",
				WorkloadGroupName:      "workloadGroupValue",
				WorkloadClassifierName: "workloadClassifierValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Synapse/workspaces/workspaceValue/sqlPools/sqlPoolValue/workloadGroups/workloadGroupValue/workloadClassifiers/workloadClassifierValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS/sQlPoOlVaLuE/wOrKlOaDgRoUpS/wOrKlOaDgRoUpVaLuE/wOrKlOaDcLaSsIfIeRs/wOrKlOaDcLaSsIfIeRvAlUe",
			Expected: &WorkloadClassifierId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "eXaMpLe-rEsOuRcE-GrOuP",
				WorkspaceName:          "wOrKsPaCeVaLuE",
				SqlPoolName:            "sQlPoOlVaLuE",
				WorkloadGroupName:      "wOrKlOaDgRoUpVaLuE",
				WorkloadClassifierName: "wOrKlOaDcLaSsIfIeRvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sYnApSe/wOrKsPaCeS/wOrKsPaCeVaLuE/sQlPoOlS/sQlPoOlVaLuE/wOrKlOaDgRoUpS/wOrKlOaDgRoUpVaLuE/wOrKlOaDcLaSsIfIeRs/wOrKlOaDcLaSsIfIeRvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseWorkloadClassifierIDInsensitively(v.Input)
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

		if actual.WorkspaceName != v.Expected.WorkspaceName {
			t.Fatalf("Expected %q but got %q for WorkspaceName", v.Expected.WorkspaceName, actual.WorkspaceName)
		}

		if actual.SqlPoolName != v.Expected.SqlPoolName {
			t.Fatalf("Expected %q but got %q for SqlPoolName", v.Expected.SqlPoolName, actual.SqlPoolName)
		}

		if actual.WorkloadGroupName != v.Expected.WorkloadGroupName {
			t.Fatalf("Expected %q but got %q for WorkloadGroupName", v.Expected.WorkloadGroupName, actual.WorkloadGroupName)
		}

		if actual.WorkloadClassifierName != v.Expected.WorkloadClassifierName {
			t.Fatalf("Expected %q but got %q for WorkloadClassifierName", v.Expected.WorkloadClassifierName, actual.WorkloadClassifierName)
		}

	}
}

func TestSegmentsForWorkloadClassifierId(t *testing.T) {
	segments := WorkloadClassifierId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("WorkloadClassifierId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
