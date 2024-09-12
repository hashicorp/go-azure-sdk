package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{}

func TestNewGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	id := NewGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID("groupIdValue", "siteIdValue", "threatAssessmentRequestIdValue", "threatAssessmentResultIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.SiteId != "siteIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SiteId'", id.SiteId, "siteIdValue")
	}

	if id.ThreatAssessmentRequestId != "threatAssessmentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentRequestId'", id.ThreatAssessmentRequestId, "threatAssessmentRequestIdValue")
	}

	if id.ThreatAssessmentResultId != "threatAssessmentResultIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentResultId'", id.ThreatAssessmentResultId, "threatAssessmentResultIdValue")
	}
}

func TestFormatGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	actual := NewGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID("groupIdValue", "siteIdValue", "threatAssessmentRequestIdValue", "threatAssessmentResultIdValue").ID()
	expected := "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue",
			Expected: &GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{
				GroupId:                   "groupIdValue",
				SiteId:                    "siteIdValue",
				ThreatAssessmentRequestId: "threatAssessmentRequestIdValue",
				ThreatAssessmentResultId:  "threatAssessmentResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(v.Input)
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

		if actual.ThreatAssessmentRequestId != v.Expected.ThreatAssessmentRequestId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentRequestId", v.Expected.ThreatAssessmentRequestId, actual.ThreatAssessmentRequestId)
		}

		if actual.ThreatAssessmentResultId != v.Expected.ThreatAssessmentResultId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentResultId", v.Expected.ThreatAssessmentResultId, actual.ThreatAssessmentResultId)
		}

	}
}

func TestParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId
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
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/rEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue",
			Expected: &GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{
				GroupId:                   "groupIdValue",
				SiteId:                    "siteIdValue",
				ThreatAssessmentRequestId: "threatAssessmentRequestIdValue",
				ThreatAssessmentResultId:  "threatAssessmentResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/sites/siteIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiDvAlUe",
			Expected: &GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{
				GroupId:                   "gRoUpIdVaLuE",
				SiteId:                    "sItEiDvAlUe",
				ThreatAssessmentRequestId: "tHrEaTaSsEsSmEnTrEqUeStIdVaLuE",
				ThreatAssessmentResultId:  "tHrEaTaSsEsSmEnTrEsUlTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sItEs/sItEiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(v.Input)
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

		if actual.ThreatAssessmentRequestId != v.Expected.ThreatAssessmentRequestId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentRequestId", v.Expected.ThreatAssessmentRequestId, actual.ThreatAssessmentRequestId)
		}

		if actual.ThreatAssessmentResultId != v.Expected.ThreatAssessmentResultId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentResultId", v.Expected.ThreatAssessmentResultId, actual.ThreatAssessmentResultId)
		}

	}
}

func TestSegmentsForGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId(t *testing.T) {
	segments := GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
