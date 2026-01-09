package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{}

func TestNewGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	id := NewGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID("groupId", "siteId", "bitlockerRecoveryKeyId")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.BitlockerRecoveryKeyId != "bitlockerRecoveryKeyId" {
		t.Fatalf("Expected %q but got %q for Segment 'BitlockerRecoveryKeyId'", id.BitlockerRecoveryKeyId, "bitlockerRecoveryKeyId")
	}
}

func TestFormatGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	actual := NewGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID("groupId", "siteId", "bitlockerRecoveryKeyId").ID()
	expected := "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId
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
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyId",
			Expected: &GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{
				GroupId:                "groupId",
				SiteId:                 "siteId",
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyID(v.Input)
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

func TestParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId
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
			Input: "/groups/groupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyId",
			Expected: &GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{
				GroupId:                "groupId",
				SiteId:                 "siteId",
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiD",
			Expected: &GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{
				GroupId:                "gRoUpId",
				SiteId:                 "sItEiD",
				BitlockerRecoveryKeyId: "bItLoCkErReCoVeRyKeYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyIDInsensitively(v.Input)
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

func TestSegmentsForGroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId(t *testing.T) {
	segments := GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdInformationProtectionBitlockerRecoveryKeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
