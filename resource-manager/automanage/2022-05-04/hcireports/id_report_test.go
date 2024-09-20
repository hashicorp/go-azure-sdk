package hcireports

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportId{}

func TestNewReportID(t *testing.T) {
	id := NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "configurationProfileAssignmentName", "reportName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ClusterName != "clusterName" {
		t.Fatalf("Expected %q but got %q for Segment 'ClusterName'", id.ClusterName, "clusterName")
	}

	if id.ConfigurationProfileAssignmentName != "configurationProfileAssignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'ConfigurationProfileAssignmentName'", id.ConfigurationProfileAssignmentName, "configurationProfileAssignmentName")
	}

	if id.ReportName != "reportName" {
		t.Fatalf("Expected %q but got %q for Segment 'ReportName'", id.ReportName, "reportName")
	}
}

func TestFormatReportID(t *testing.T) {
	actual := NewReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "configurationProfileAssignmentName", "reportName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports/reportName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports/reportName",
			Expected: &ReportId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				ClusterName:                        "clusterName",
				ConfigurationProfileAssignmentName: "configurationProfileAssignmentName",
				ReportName:                         "reportName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports/reportName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportID(v.Input)
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

		if actual.ClusterName != v.Expected.ClusterName {
			t.Fatalf("Expected %q but got %q for ClusterName", v.Expected.ClusterName, actual.ClusterName)
		}

		if actual.ConfigurationProfileAssignmentName != v.Expected.ConfigurationProfileAssignmentName {
			t.Fatalf("Expected %q but got %q for ConfigurationProfileAssignmentName", v.Expected.ConfigurationProfileAssignmentName, actual.ConfigurationProfileAssignmentName)
		}

		if actual.ReportName != v.Expected.ReportName {
			t.Fatalf("Expected %q but got %q for ReportName", v.Expected.ReportName, actual.ReportName)
		}

	}
}

func TestParseReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTnAmE/rEpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports/reportName",
			Expected: &ReportId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "example-resource-group",
				ClusterName:                        "clusterName",
				ConfigurationProfileAssignmentName: "configurationProfileAssignmentName",
				ReportName:                         "reportName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureStackHCI/clusters/clusterName/providers/Microsoft.AutoManage/configurationProfileAssignments/configurationProfileAssignmentName/reports/reportName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTnAmE/rEpOrTs/rEpOrTnAmE",
			Expected: &ReportId{
				SubscriptionId:                     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                  "eXaMpLe-rEsOuRcE-GrOuP",
				ClusterName:                        "cLuStErNaMe",
				ConfigurationProfileAssignmentName: "cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTnAmE",
				ReportName:                         "rEpOrTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.aZuReStAcKhCi/cLuStErS/cLuStErNaMe/pRoViDeRs/mIcRoSoFt.aUtOmAnAgE/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTs/cOnFiGuRaTiOnPrOfIlEaSsIgNmEnTnAmE/rEpOrTs/rEpOrTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportIDInsensitively(v.Input)
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

		if actual.ClusterName != v.Expected.ClusterName {
			t.Fatalf("Expected %q but got %q for ClusterName", v.Expected.ClusterName, actual.ClusterName)
		}

		if actual.ConfigurationProfileAssignmentName != v.Expected.ConfigurationProfileAssignmentName {
			t.Fatalf("Expected %q but got %q for ConfigurationProfileAssignmentName", v.Expected.ConfigurationProfileAssignmentName, actual.ConfigurationProfileAssignmentName)
		}

		if actual.ReportName != v.Expected.ReportName {
			t.Fatalf("Expected %q but got %q for ReportName", v.Expected.ReportName, actual.ReportName)
		}

	}
}

func TestSegmentsForReportId(t *testing.T) {
	segments := ReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
