package packetcorecontrolplaneversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PacketCoreControlPlaneVersionId{}

func TestNewPacketCoreControlPlaneVersionID(t *testing.T) {
	id := NewPacketCoreControlPlaneVersionID("versionValue")

	if id.VersionName != "versionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionValue")
	}
}

func TestFormatPacketCoreControlPlaneVersionID(t *testing.T) {
	actual := NewPacketCoreControlPlaneVersionID("versionValue").ID()
	expected := "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/versionValue"
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
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/versionValue",
			Expected: &PacketCoreControlPlaneVersionId{
				VersionName: "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/versionValue/extra",
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
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
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/versionValue",
			Expected: &PacketCoreControlPlaneVersionId{
				VersionName: "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/versionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS/vErSiOnVaLuE",
			Expected: &PacketCoreControlPlaneVersionId{
				VersionName: "vErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mObIlEnEtWoRk/pAcKeTcOrEcOnTrOlPlAnEvErSiOnS/vErSiOnVaLuE/extra",
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
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
