package ransomwarereports

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RansomwareReportId{}

func TestNewRansomwareReportID(t *testing.T) {
	id := NewRansomwareReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "ransomwareReportName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.NetAppAccountName != "netAppAccountName" {
		t.Fatalf("Expected %q but got %q for Segment 'NetAppAccountName'", id.NetAppAccountName, "netAppAccountName")
	}

	if id.CapacityPoolName != "capacityPoolName" {
		t.Fatalf("Expected %q but got %q for Segment 'CapacityPoolName'", id.CapacityPoolName, "capacityPoolName")
	}

	if id.VolumeName != "volumeName" {
		t.Fatalf("Expected %q but got %q for Segment 'VolumeName'", id.VolumeName, "volumeName")
	}

	if id.RansomwareReportName != "ransomwareReportName" {
		t.Fatalf("Expected %q but got %q for Segment 'RansomwareReportName'", id.RansomwareReportName, "ransomwareReportName")
	}
}

func TestFormatRansomwareReportID(t *testing.T) {
	actual := NewRansomwareReportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "ransomwareReportName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports/ransomwareReportName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRansomwareReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RansomwareReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports/ransomwareReportName",
			Expected: &RansomwareReportId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				NetAppAccountName:    "netAppAccountName",
				CapacityPoolName:     "capacityPoolName",
				VolumeName:           "volumeName",
				RansomwareReportName: "ransomwareReportName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports/ransomwareReportName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRansomwareReportID(v.Input)
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

		if actual.NetAppAccountName != v.Expected.NetAppAccountName {
			t.Fatalf("Expected %q but got %q for NetAppAccountName", v.Expected.NetAppAccountName, actual.NetAppAccountName)
		}

		if actual.CapacityPoolName != v.Expected.CapacityPoolName {
			t.Fatalf("Expected %q but got %q for CapacityPoolName", v.Expected.CapacityPoolName, actual.CapacityPoolName)
		}

		if actual.VolumeName != v.Expected.VolumeName {
			t.Fatalf("Expected %q but got %q for VolumeName", v.Expected.VolumeName, actual.VolumeName)
		}

		if actual.RansomwareReportName != v.Expected.RansomwareReportName {
			t.Fatalf("Expected %q but got %q for RansomwareReportName", v.Expected.RansomwareReportName, actual.RansomwareReportName)
		}

	}
}

func TestParseRansomwareReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RansomwareReportId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs/cApAcItYpOoLnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs/cApAcItYpOoLnAmE/vOlUmEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs/cApAcItYpOoLnAmE/vOlUmEs/vOlUmEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs/cApAcItYpOoLnAmE/vOlUmEs/vOlUmEnAmE/rAnSoMwArErEpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports/ransomwareReportName",
			Expected: &RansomwareReportId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				NetAppAccountName:    "netAppAccountName",
				CapacityPoolName:     "capacityPoolName",
				VolumeName:           "volumeName",
				RansomwareReportName: "ransomwareReportName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.NetApp/netAppAccounts/netAppAccountName/capacityPools/capacityPoolName/volumes/volumeName/ransomwareReports/ransomwareReportName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs/cApAcItYpOoLnAmE/vOlUmEs/vOlUmEnAmE/rAnSoMwArErEpOrTs/rAnSoMwArErEpOrTnAmE",
			Expected: &RansomwareReportId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "eXaMpLe-rEsOuRcE-GrOuP",
				NetAppAccountName:    "nEtApPaCcOuNtNaMe",
				CapacityPoolName:     "cApAcItYpOoLnAmE",
				VolumeName:           "vOlUmEnAmE",
				RansomwareReportName: "rAnSoMwArErEpOrTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.nEtApP/nEtApPaCcOuNtS/nEtApPaCcOuNtNaMe/cApAcItYpOoLs/cApAcItYpOoLnAmE/vOlUmEs/vOlUmEnAmE/rAnSoMwArErEpOrTs/rAnSoMwArErEpOrTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRansomwareReportIDInsensitively(v.Input)
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

		if actual.NetAppAccountName != v.Expected.NetAppAccountName {
			t.Fatalf("Expected %q but got %q for NetAppAccountName", v.Expected.NetAppAccountName, actual.NetAppAccountName)
		}

		if actual.CapacityPoolName != v.Expected.CapacityPoolName {
			t.Fatalf("Expected %q but got %q for CapacityPoolName", v.Expected.CapacityPoolName, actual.CapacityPoolName)
		}

		if actual.VolumeName != v.Expected.VolumeName {
			t.Fatalf("Expected %q but got %q for VolumeName", v.Expected.VolumeName, actual.VolumeName)
		}

		if actual.RansomwareReportName != v.Expected.RansomwareReportName {
			t.Fatalf("Expected %q but got %q for RansomwareReportName", v.Expected.RansomwareReportName, actual.RansomwareReportName)
		}

	}
}

func TestSegmentsForRansomwareReportId(t *testing.T) {
	segments := RansomwareReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RansomwareReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
