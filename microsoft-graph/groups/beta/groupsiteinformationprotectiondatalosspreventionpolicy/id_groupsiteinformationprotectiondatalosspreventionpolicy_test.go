package groupsiteinformationprotectiondatalosspreventionpolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupSiteInformationProtectionDataLossPreventionPolicyId{}

func TestNewGroupSiteInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	id := NewGroupSiteInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.DataLossPreventionPolicyId != "dataLossPreventionPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DataLossPreventionPolicyId'", id.DataLossPreventionPolicyId, "dataLossPreventionPolicyIdValue")
	}
}

func TestFormatGroupSiteInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	actual := NewGroupSiteInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupSiteInformationProtectionDataLossPreventionPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteInformationProtectionDataLossPreventionPolicyId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue",
			Expected: &GroupSiteInformationProtectionDataLossPreventionPolicyId{
				GroupId:                    "groupIdValue",
				SiteId:                     "siteIdValue",
				DataLossPreventionPolicyId: "dataLossPreventionPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteInformationProtectionDataLossPreventionPolicyID(v.Input)
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

		if actual.DataLossPreventionPolicyId != v.Expected.DataLossPreventionPolicyId {
			t.Fatalf("Expected %q but got %q for DataLossPreventionPolicyId", v.Expected.DataLossPreventionPolicyId, actual.DataLossPreventionPolicyId)
		}

	}
}

func TestParseGroupSiteInformationProtectionDataLossPreventionPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupSiteInformationProtectionDataLossPreventionPolicyId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue",
			Expected: &GroupSiteInformationProtectionDataLossPreventionPolicyId{
				GroupId:                    "groupIdValue",
				SiteId:                     "siteIdValue",
				DataLossPreventionPolicyId: "dataLossPreventionPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/dataLossPreventionPolicies/dataLossPreventionPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe",
			Expected: &GroupSiteInformationProtectionDataLossPreventionPolicyId{
				GroupId:                    "gRoUpIdVaLuE",
				SiteId:                     "sItEiDvAlUe",
				DataLossPreventionPolicyId: "dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/dAtAlOsSpReVeNtIoNpOlIcIeS/dAtAlOsSpReVeNtIoNpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupSiteInformationProtectionDataLossPreventionPolicyIDInsensitively(v.Input)
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

		if actual.DataLossPreventionPolicyId != v.Expected.DataLossPreventionPolicyId {
			t.Fatalf("Expected %q but got %q for DataLossPreventionPolicyId", v.Expected.DataLossPreventionPolicyId, actual.DataLossPreventionPolicyId)
		}

	}
}

func TestSegmentsForGroupSiteInformationProtectionDataLossPreventionPolicyId(t *testing.T) {
	segments := GroupSiteInformationProtectionDataLossPreventionPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupSiteInformationProtectionDataLossPreventionPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
