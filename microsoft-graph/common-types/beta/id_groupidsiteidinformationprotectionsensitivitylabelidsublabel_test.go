package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{}

func TestNewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	id := NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupId", "siteId", "sensitivityLabelId", "sensitivityLabelId1")

	if id.GroupId != "groupId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupId")
	}

	if id.SiteId != "siteId" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteId")
	}

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}

	if id.SensitivityLabelId1 != "sensitivityLabelId1" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId1'", id.SensitivityLabelId1, "sensitivityLabelId1")
	}
}

func TestFormatGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	actual := NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID("groupId", "siteId", "sensitivityLabelId", "sensitivityLabelId1").ID()
	expected := "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId
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
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{
				GroupId:             "groupId",
				SiteId:              "siteId",
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(v.Input)
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

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId
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
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{
				GroupId:             "groupId",
				SiteId:              "siteId",
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupId/sites/siteId/informationProtection/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1",
			Expected: &GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{
				GroupId:             "gRoUpId",
				SiteId:              "sItEiD",
				SensitivityLabelId:  "sEnSiTiViTyLaBeLiD",
				SensitivityLabelId1: "sEnSiTiViTyLaBeLiD1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpId/sItEs/sItEiD/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively(v.Input)
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

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestSegmentsForGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId(t *testing.T) {
	segments := GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
