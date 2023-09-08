package groupsiteinformationprotectionsensitivitylabel

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionSensitivityLabelId{}

func TestNewGroupSiteInformationProtectionSensitivityLabelID(t *testing.T) {
	id := NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.SensitivityLabelId != "sensitivityLabelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelIdValue")
	}
}

func TestFormatGroupSiteInformationProtectionSensitivityLabelID(t *testing.T) {
	actual := NewGroupSiteInformationProtectionSensitivityLabelID("groupIdValue", "siteIdValue", "sensitivityLabelIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteInformationProtectionSensitivityLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteInformationProtectionSensitivityLabelId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Expected: &GroupSiteInformationProtectionSensitivityLabelId{
				GroupId:            "groupIdValue",
				SiteId:             "siteIdValue",
				SensitivityLabelId: "sensitivityLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteInformationProtectionSensitivityLabelID(v.Input)
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

	}
}

func TestParseGroupSiteInformationProtectionSensitivityLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteInformationProtectionSensitivityLabelId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue",
			Expected: &GroupSiteInformationProtectionSensitivityLabelId{
				GroupId:            "groupIdValue",
				SiteId:             "siteIdValue",
				SensitivityLabelId: "sensitivityLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/sensitivityLabels/sensitivityLabelIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe",
			Expected: &GroupSiteInformationProtectionSensitivityLabelId{
				GroupId:            "gRoUpIdVaLuE",
				SiteId:             "sItEiDvAlUe",
				SensitivityLabelId: "sEnSiTiViTyLaBeLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteInformationProtectionSensitivityLabelIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupSiteInformationProtectionSensitivityLabelId(t *testing.T) {
	segments := GroupSiteInformationProtectionSensitivityLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteInformationProtectionSensitivityLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
