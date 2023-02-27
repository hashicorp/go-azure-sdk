// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syncsets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SyncSetId{}

func TestNewSyncSetID(t *testing.T) {
	id := NewSyncSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncSetValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.OpenShiftClusterName != "openShiftClusterValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OpenShiftClusterName'", id.OpenShiftClusterName, "openShiftClusterValue")
	}

	if id.SyncSetName != "syncSetValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SyncSetName'", id.SyncSetName, "syncSetValue")
	}
}

func TestFormatSyncSetID(t *testing.T) {
	actual := NewSyncSetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncSetValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet/syncSetValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSyncSetID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SyncSetId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet/syncSetValue",
			Expected: &SyncSetId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				OpenShiftClusterName: "openShiftClusterValue",
				SyncSetName:          "syncSetValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet/syncSetValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSyncSetID(v.Input)
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

		if actual.OpenShiftClusterName != v.Expected.OpenShiftClusterName {
			t.Fatalf("Expected %q but got %q for OpenShiftClusterName", v.Expected.OpenShiftClusterName, actual.OpenShiftClusterName)
		}

		if actual.SyncSetName != v.Expected.SyncSetName {
			t.Fatalf("Expected %q but got %q for SyncSetName", v.Expected.SyncSetName, actual.SyncSetName)
		}

	}
}

func TestParseSyncSetIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SyncSetId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/oPeNsHiFtClUsTeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/oPeNsHiFtClUsTeRs/oPeNsHiFtClUsTeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/oPeNsHiFtClUsTeRs/oPeNsHiFtClUsTeRvAlUe/sYnCsEt",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet/syncSetValue",
			Expected: &SyncSetId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "example-resource-group",
				OpenShiftClusterName: "openShiftClusterValue",
				SyncSetName:          "syncSetValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.RedHatOpenShift/openShiftClusters/openShiftClusterValue/syncSet/syncSetValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/oPeNsHiFtClUsTeRs/oPeNsHiFtClUsTeRvAlUe/sYnCsEt/sYnCsEtVaLuE",
			Expected: &SyncSetId{
				SubscriptionId:       "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:    "eXaMpLe-rEsOuRcE-GrOuP",
				OpenShiftClusterName: "oPeNsHiFtClUsTeRvAlUe",
				SyncSetName:          "sYnCsEtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/oPeNsHiFtClUsTeRs/oPeNsHiFtClUsTeRvAlUe/sYnCsEt/sYnCsEtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSyncSetIDInsensitively(v.Input)
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

		if actual.OpenShiftClusterName != v.Expected.OpenShiftClusterName {
			t.Fatalf("Expected %q but got %q for OpenShiftClusterName", v.Expected.OpenShiftClusterName, actual.OpenShiftClusterName)
		}

		if actual.SyncSetName != v.Expected.SyncSetName {
			t.Fatalf("Expected %q but got %q for SyncSetName", v.Expected.SyncSetName, actual.SyncSetName)
		}

	}
}

func TestSegmentsForSyncSetId(t *testing.T) {
	segments := SyncSetId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SyncSetId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
