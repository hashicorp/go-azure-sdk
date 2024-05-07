package packetcaptures

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PacketCaptureId{}

func TestNewPacketCaptureID(t *testing.T) {
	id := NewPacketCaptureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "packetCaptureValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.PacketCoreControlPlaneName != "packetCoreControlPlaneValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PacketCoreControlPlaneName'", id.PacketCoreControlPlaneName, "packetCoreControlPlaneValue")
	}

	if id.PacketCaptureName != "packetCaptureValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PacketCaptureName'", id.PacketCaptureName, "packetCaptureValue")
	}
}

func TestFormatPacketCaptureID(t *testing.T) {
	actual := NewPacketCaptureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "packetCoreControlPlaneValue", "packetCaptureValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures/packetCaptureValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePacketCaptureID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PacketCaptureId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures/packetCaptureValue",
			Expected: &PacketCaptureId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				PacketCoreControlPlaneName: "packetCoreControlPlaneValue",
				PacketCaptureName:          "packetCaptureValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures/packetCaptureValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePacketCaptureID(v.Input)
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

		if actual.PacketCaptureName != v.Expected.PacketCaptureName {
			t.Fatalf("Expected %q but got %q for PacketCaptureName", v.Expected.PacketCaptureName, actual.PacketCaptureName)
		}

	}
}

func TestParsePacketCaptureIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PacketCaptureId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEvAlUe/pAcKeTcApTuReS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures/packetCaptureValue",
			Expected: &PacketCaptureId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "example-resource-group",
				PacketCoreControlPlaneName: "packetCoreControlPlaneValue",
				PacketCaptureName:          "packetCaptureValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/packetCoreControlPlaneValue/packetCaptures/packetCaptureValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEvAlUe/pAcKeTcApTuReS/pAcKeTcApTuReVaLuE",
			Expected: &PacketCaptureId{
				SubscriptionId:             "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:          "eXaMpLe-rEsOuRcE-GrOuP",
				PacketCoreControlPlaneName: "pAcKeTcOrEcOnTrOlPlAnEvAlUe",
				PacketCaptureName:          "pAcKeTcApTuReVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEs/pAcKeTcOrEcOnTrOlPlAnEvAlUe/pAcKeTcApTuReS/pAcKeTcApTuReVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePacketCaptureIDInsensitively(v.Input)
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

		if actual.PacketCaptureName != v.Expected.PacketCaptureName {
			t.Fatalf("Expected %q but got %q for PacketCaptureName", v.Expected.PacketCaptureName, actual.PacketCaptureName)
		}

	}
}

func TestSegmentsForPacketCaptureId(t *testing.T) {
	segments := PacketCaptureId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PacketCaptureId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
