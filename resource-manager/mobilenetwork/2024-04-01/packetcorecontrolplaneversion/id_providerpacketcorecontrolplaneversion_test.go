package packetcorecontrolplaneversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProviderPacketCoreControlPlaneVersionId{}

func TestNewProviderPacketCoreControlPlaneVersionID(t *testing.T) {
	id := NewProviderPacketCoreControlPlaneVersionID("12345678-1234-9876-4563-123456789012", "packetCoreControlPlaneVersionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.PacketCoreControlPlaneVersionName != "packetCoreControlPlaneVersionName" {
		t.Fatalf("Expected %q but got %q for Segment 'PacketCoreControlPlaneVersionName'", id.PacketCoreControlPlaneVersionName, "packetCoreControlPlaneVersionName")
	}
}

func TestFormatProviderPacketCoreControlPlaneVersionID(t *testing.T) {
	actual := NewProviderPacketCoreControlPlaneVersionID("12345678-1234-9876-4563-123456789012", "packetCoreControlPlaneVersionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderPacketCoreControlPlaneVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderPacketCoreControlPlaneVersionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionName",
			Expected: &ProviderPacketCoreControlPlaneVersionId{
				SubscriptionId:                    "12345678-1234-9876-4563-123456789012",
				PacketCoreControlPlaneVersionName: "packetCoreControlPlaneVersionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderPacketCoreControlPlaneVersionID(v.Input)
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

		if actual.PacketCoreControlPlaneVersionName != v.Expected.PacketCoreControlPlaneVersionName {
			t.Fatalf("Expected %q but got %q for PacketCoreControlPlaneVersionName", v.Expected.PacketCoreControlPlaneVersionName, actual.PacketCoreControlPlaneVersionName)
		}

	}
}

func TestParseProviderPacketCoreControlPlaneVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderPacketCoreControlPlaneVersionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionName",
			Expected: &ProviderPacketCoreControlPlaneVersionId{
				SubscriptionId:                    "12345678-1234-9876-4563-123456789012",
				PacketCoreControlPlaneVersionName: "packetCoreControlPlaneVersionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS/pAcKeTcOrEcOnTrOlPlAnEvErSiOnNaMe",
			Expected: &ProviderPacketCoreControlPlaneVersionId{
				SubscriptionId:                    "12345678-1234-9876-4563-123456789012",
				PacketCoreControlPlaneVersionName: "pAcKeTcOrEcOnTrOlPlAnEvErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS/pAcKeTcOrEcOnTrOlPlAnEvErSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderPacketCoreControlPlaneVersionIDInsensitively(v.Input)
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

		if actual.PacketCoreControlPlaneVersionName != v.Expected.PacketCoreControlPlaneVersionName {
			t.Fatalf("Expected %q but got %q for PacketCoreControlPlaneVersionName", v.Expected.PacketCoreControlPlaneVersionName, actual.PacketCoreControlPlaneVersionName)
		}

	}
}

func TestSegmentsForProviderPacketCoreControlPlaneVersionId(t *testing.T) {
	segments := ProviderPacketCoreControlPlaneVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderPacketCoreControlPlaneVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
