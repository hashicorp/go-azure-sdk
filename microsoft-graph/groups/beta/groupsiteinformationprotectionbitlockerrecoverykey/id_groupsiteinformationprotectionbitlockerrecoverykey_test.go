package groupsiteinformationprotectionbitlockerrecoverykey

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionBitlockerRecoveryKeyId{}

func TestNewGroupSiteInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	id := NewGroupSiteInformationProtectionBitlockerRecoveryKeyID("groupIdValue", "siteIdValue", "bitlockerRecoveryKeyIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.BitlockerRecoveryKeyId != "bitlockerRecoveryKeyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BitlockerRecoveryKeyId'", id.BitlockerRecoveryKeyId, "bitlockerRecoveryKeyIdValue")
	}
}

func TestFormatGroupSiteInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	actual := NewGroupSiteInformationProtectionBitlockerRecoveryKeyID("groupIdValue", "siteIdValue", "bitlockerRecoveryKeyIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteInformationProtectionBitlockerRecoveryKeyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue",
			Expected: &GroupSiteInformationProtectionBitlockerRecoveryKeyId{
				GroupId:                "groupIdValue",
				SiteId:                 "siteIdValue",
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteInformationProtectionBitlockerRecoveryKeyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

		if actual.BitlockerRecoveryKeyId != v.Expected.BitlockerRecoveryKeyId {
			t.Fatalf("Expected %q but got %q for BitlockerRecoveryKeyId", v.Expected.BitlockerRecoveryKeyId, actual.BitlockerRecoveryKeyId)
		}

	}
}

func TestParseGroupSiteInformationProtectionBitlockerRecoveryKeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteInformationProtectionBitlockerRecoveryKeyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue",
			Expected: &GroupSiteInformationProtectionBitlockerRecoveryKeyId{
				GroupId:                "groupIdValue",
				SiteId:                 "siteIdValue",
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiDvAlUe",
			Expected: &GroupSiteInformationProtectionBitlockerRecoveryKeyId{
				GroupId:                "gRoUpIdVaLuE",
				SiteId:                 "sItEiDvAlUe",
				BitlockerRecoveryKeyId: "bItLoCkErReCoVeRyKeYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteInformationProtectionBitlockerRecoveryKeyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.SiteId != v.Expected.SiteId {
			t.Fatalf("Expected %q but got %q for SiteId", v.Expected.SiteId, actual.SiteId)
		}

		if actual.BitlockerRecoveryKeyId != v.Expected.BitlockerRecoveryKeyId {
			t.Fatalf("Expected %q but got %q for BitlockerRecoveryKeyId", v.Expected.BitlockerRecoveryKeyId, actual.BitlockerRecoveryKeyId)
		}

	}
}

func TestSegmentsForGroupSiteInformationProtectionBitlockerRecoveryKeyId(t *testing.T) {
	segments := GroupSiteInformationProtectionBitlockerRecoveryKeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteInformationProtectionBitlockerRecoveryKeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
