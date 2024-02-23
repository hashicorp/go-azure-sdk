package packetcorecontrolplaneversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PacketCoreControlPlaneVersionId{}

func TestNewPacketCoreControlPlaneVersionID(t *testing.T) {
	id := NewPacketCoreControlPlaneVersionID("packetCoreControlPlaneVersionValue")

	if id.PacketCoreControlPlaneVersionName != "packetCoreControlPlaneVersionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PacketCoreControlPlaneVersionName'", id.PacketCoreControlPlaneVersionName, "packetCoreControlPlaneVersionValue")
	}
}

func TestFormatPacketCoreControlPlaneVersionID(t *testing.T) {
	actual := NewPacketCoreControlPlaneVersionID("packetCoreControlPlaneVersionValue").ID()
	expected := "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePacketCoreControlPlaneVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PacketCoreControlPlaneVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.MobileNetwork",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionValue",
			Expected: &PacketCoreControlPlaneVersionId{
				PacketCoreControlPlaneVersionName: "packetCoreControlPlaneVersionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePacketCoreControlPlaneVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PacketCoreControlPlaneVersionName != v.Expected.PacketCoreControlPlaneVersionName {
			t.Fatalf("Expected %q but got %q for PacketCoreControlPlaneVersionName", v.Expected.PacketCoreControlPlaneVersionName, actual.PacketCoreControlPlaneVersionName)
		}

	}
}

func TestParsePacketCoreControlPlaneVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PacketCoreControlPlaneVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.MobileNetwork",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionValue",
			Expected: &PacketCoreControlPlaneVersionId{
				PacketCoreControlPlaneVersionName: "packetCoreControlPlaneVersionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/packetCoreControlPlaneVersionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS/pAcKeTcOrEcOnTrOlPlAnEvErSiOnVaLuE",
			Expected: &PacketCoreControlPlaneVersionId{
				PacketCoreControlPlaneVersionName: "pAcKeTcOrEcOnTrOlPlAnEvErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS/pAcKeTcOrEcOnTrOlPlAnEvErSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePacketCoreControlPlaneVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PacketCoreControlPlaneVersionName != v.Expected.PacketCoreControlPlaneVersionName {
			t.Fatalf("Expected %q but got %q for PacketCoreControlPlaneVersionName", v.Expected.PacketCoreControlPlaneVersionName, actual.PacketCoreControlPlaneVersionName)
		}

	}
}

func TestSegmentsForPacketCoreControlPlaneVersionId(t *testing.T) {
	segments := PacketCoreControlPlaneVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PacketCoreControlPlaneVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
