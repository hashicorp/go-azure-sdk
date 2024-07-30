package siteinformationprotectionpolicylabel

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionPolicyLabelId{}

func TestNewGroupIdSiteIdInformationProtectionPolicyLabelID(t *testing.T) {
	id := NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.InformationProtectionLabelId != "informationProtectionLabelIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InformationProtectionLabelId'", id.InformationProtectionLabelId, "informationProtectionLabelIdValue")
	}
}

func TestFormatGroupIdSiteIdInformationProtectionPolicyLabelID(t *testing.T) {
	actual := NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels/informationProtectionLabelIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdInformationProtectionPolicyLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionPolicyLabelId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels/informationProtectionLabelIdValue",
			Expected: &GroupIdSiteIdInformationProtectionPolicyLabelId{
				GroupId:                      "groupIdValue",
				SiteId:                       "siteIdValue",
				InformationProtectionLabelId: "informationProtectionLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels/informationProtectionLabelIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionPolicyLabelID(v.Input)
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

		if actual.InformationProtectionLabelId != v.Expected.InformationProtectionLabelId {
			t.Fatalf("Expected %q but got %q for InformationProtectionLabelId", v.Expected.InformationProtectionLabelId, actual.InformationProtectionLabelId)
		}

	}
}

func TestParseGroupIdSiteIdInformationProtectionPolicyLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionPolicyLabelId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/pOlIcY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels/informationProtectionLabelIdValue",
			Expected: &GroupIdSiteIdInformationProtectionPolicyLabelId{
				GroupId:                      "groupIdValue",
				SiteId:                       "siteIdValue",
				InformationProtectionLabelId: "informationProtectionLabelIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/policy/labels/informationProtectionLabelIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS/iNfOrMaTiOnPrOtEcTiOnLaBeLiDvAlUe",
			Expected: &GroupIdSiteIdInformationProtectionPolicyLabelId{
				GroupId:                      "gRoUpIdVaLuE",
				SiteId:                       "sItEiDvAlUe",
				InformationProtectionLabelId: "iNfOrMaTiOnPrOtEcTiOnLaBeLiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/pOlIcY/lAbElS/iNfOrMaTiOnPrOtEcTiOnLaBeLiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionPolicyLabelIDInsensitively(v.Input)
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

		if actual.InformationProtectionLabelId != v.Expected.InformationProtectionLabelId {
			t.Fatalf("Expected %q but got %q for InformationProtectionLabelId", v.Expected.InformationProtectionLabelId, actual.InformationProtectionLabelId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdInformationProtectionPolicyLabelId(t *testing.T) {
	segments := GroupIdSiteIdInformationProtectionPolicyLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdInformationProtectionPolicyLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
