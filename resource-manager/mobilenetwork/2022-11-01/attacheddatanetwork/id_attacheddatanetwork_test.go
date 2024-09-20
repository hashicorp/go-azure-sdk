package attacheddatanetwork

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AttachedDataNetworkId{}

func TestNewAttachedDataNetworkID(t *testing.T) {
	id := NewAttachedDataNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName", "packetCoreDataPlaneName", "attachedDataNetworkName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.PacketCoreControlPlaneName != "packetCoreControlPlaneName" {
		t.Fatalf("Expected %q but got %q for Segment 'PacketCoreControlPlaneName'", id.PacketCoreControlPlaneName, "packetCoreControlPlaneName")
	}

	if id.PacketCoreDataPlaneName != "packetCoreDataPlaneName" {
		t.Fatalf("Expected %q but got %q for Segment 'PacketCoreDataPlaneName'", id.PacketCoreDataPlaneName, "packetCoreDataPlaneName")
	}

	if id.AttachedDataNetworkName != "attachedDataNetworkName" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachedDataNetworkName'", id.AttachedDataNetworkName, "attachedDataNetworkName")
	}
}

func TestFormatAttachedDataNetworkID(t *testing.T) {
	actual := NewAttachedDataNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneName", "packetCoreDataPlaneName", "attachedDataNetworkName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks/attachedDataNetworkName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAttachedDataNetworkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AttachedDataNetworkId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks/attachedDataNetworkName",
			Expected: &AttachedDataNetworkId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				PacketCoreControlPlaneName: "packetCoreControlPlaneName",
				PacketCoreDataPlaneName:    "packetCoreDataPlaneName",
				AttachedDataNetworkName:    "attachedDataNetworkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks/attachedDataNetworkName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAttachedDataNetworkID(v.Input)
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

		if actual.PacketCoreControlPlaneName != v.Expected.PacketCoreControlPlaneName {
			t.Fatalf("Expected %q but got %q for PacketCoreControlPlaneName", v.Expected.PacketCoreControlPlaneName, actual.PacketCoreControlPlaneName)
		}

		if actual.PacketCoreDataPlaneName != v.Expected.PacketCoreDataPlaneName {
			t.Fatalf("Expected %q but got %q for PacketCoreDataPlaneName", v.Expected.PacketCoreDataPlaneName, actual.PacketCoreDataPlaneName)
		}

		if actual.AttachedDataNetworkName != v.Expected.AttachedDataNetworkName {
			t.Fatalf("Expected %q but got %q for AttachedDataNetworkName", v.Expected.AttachedDataNetworkName, actual.AttachedDataNetworkName)
		}

	}
}

func TestParseAttachedDataNetworkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AttachedDataNetworkId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEnAmE/pAcKeTcOrEdAtApLaNeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEnAmE/pAcKeTcOrEdAtApLaNeS/pAcKeTcOrEdAtApLaNeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEnAmE/pAcKeTcOrEdAtApLaNeS/pAcKeTcOrEdAtApLaNeNaMe/aTtAcHeDdAtAnEtWoRkS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks/attachedDataNetworkName",
			Expected: &AttachedDataNetworkId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				PacketCoreControlPlaneName: "packetCoreControlPlaneName",
				PacketCoreDataPlaneName:    "packetCoreDataPlaneName",
				AttachedDataNetworkName:    "attachedDataNetworkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneName/packetCoreDataPlanes/packetCoreDataPlaneName/attachedDataNetworks/attachedDataNetworkName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEnAmE/pAcKeTcOrEdAtApLaNeS/pAcKeTcOrEdAtApLaNeNaMe/aTtAcHeDdAtAnEtWoRkS/aTtAcHeDdAtAnEtWoRkNaMe",
			Expected: &AttachedDataNetworkId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				PacketCoreControlPlaneName: "pAcKeTcOrEcOnTrOlPlAnEnAmE",
				PacketCoreDataPlaneName:    "pAcKeTcOrEdAtApLaNeNaMe",
				AttachedDataNetworkName:    "aTtAcHeDdAtAnEtWoRkNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEnAmE/pAcKeTcOrEdAtApLaNeS/pAcKeTcOrEdAtApLaNeNaMe/aTtAcHeDdAtAnEtWoRkS/aTtAcHeDdAtAnEtWoRkNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAttachedDataNetworkIDInsensitively(v.Input)
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

		if actual.PacketCoreControlPlaneName != v.Expected.PacketCoreControlPlaneName {
			t.Fatalf("Expected %q but got %q for PacketCoreControlPlaneName", v.Expected.PacketCoreControlPlaneName, actual.PacketCoreControlPlaneName)
		}

		if actual.PacketCoreDataPlaneName != v.Expected.PacketCoreDataPlaneName {
			t.Fatalf("Expected %q but got %q for PacketCoreDataPlaneName", v.Expected.PacketCoreDataPlaneName, actual.PacketCoreDataPlaneName)
		}

		if actual.AttachedDataNetworkName != v.Expected.AttachedDataNetworkName {
			t.Fatalf("Expected %q but got %q for AttachedDataNetworkName", v.Expected.AttachedDataNetworkName, actual.AttachedDataNetworkName)
		}

	}
}

func TestSegmentsForAttachedDataNetworkId(t *testing.T) {
	segments := AttachedDataNetworkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AttachedDataNetworkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
