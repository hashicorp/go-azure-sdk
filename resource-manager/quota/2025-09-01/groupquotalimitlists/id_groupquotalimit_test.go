package groupquotalimitlists

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupQuotaLimitId{}

func TestNewGroupQuotaLimitID(t *testing.T) {
	id := NewGroupQuotaLimitID("managementGroupId", "groupQuotaName", "resourceProviderName", "groupQuotaLimitName")

	if id.ManagementGroupId != "managementGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupId")
	}

	if id.GroupQuotaName != "groupQuotaName" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupQuotaName'", id.GroupQuotaName, "groupQuotaName")
	}

	if id.ResourceProviderName != "resourceProviderName" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceProviderName'", id.ResourceProviderName, "resourceProviderName")
	}

	if id.GroupQuotaLimitName != "groupQuotaLimitName" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupQuotaLimitName'", id.GroupQuotaLimitName, "groupQuotaLimitName")
	}
}

func TestFormatGroupQuotaLimitID(t *testing.T) {
	actual := NewGroupQuotaLimitID("managementGroupId", "groupQuotaName", "resourceProviderName", "groupQuotaLimitName").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits/groupQuotaLimitName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupQuotaLimitID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupQuotaLimitId
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
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits/groupQuotaLimitName",
			Expected: &GroupQuotaLimitId{
				ManagementGroupId:    "managementGroupId",
				GroupQuotaName:       "groupQuotaName",
				ResourceProviderName: "resourceProviderName",
				GroupQuotaLimitName:  "groupQuotaLimitName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits/groupQuotaLimitName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupQuotaLimitID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.GroupQuotaName != v.Expected.GroupQuotaName {
			t.Fatalf("Expected %q but got %q for GroupQuotaName", v.Expected.GroupQuotaName, actual.GroupQuotaName)
		}

		if actual.ResourceProviderName != v.Expected.ResourceProviderName {
			t.Fatalf("Expected %q but got %q for ResourceProviderName", v.Expected.ResourceProviderName, actual.ResourceProviderName)
		}

		if actual.GroupQuotaLimitName != v.Expected.GroupQuotaLimitName {
			t.Fatalf("Expected %q but got %q for GroupQuotaLimitName", v.Expected.GroupQuotaLimitName, actual.GroupQuotaLimitName)
		}

	}
}

func TestParseGroupQuotaLimitIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupQuotaLimitId
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
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE/gRoUpQuOtAlImItS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits/groupQuotaLimitName",
			Expected: &GroupQuotaLimitId{
				ManagementGroupId:    "managementGroupId",
				GroupQuotaName:       "groupQuotaName",
				ResourceProviderName: "resourceProviderName",
				GroupQuotaLimitName:  "groupQuotaLimitName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/resourceProviders/resourceProviderName/groupQuotaLimits/groupQuotaLimitName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE/gRoUpQuOtAlImItS/gRoUpQuOtAlImItNaMe",
			Expected: &GroupQuotaLimitId{
				ManagementGroupId:    "mAnAgEmEnTgRoUpId",
				GroupQuotaName:       "gRoUpQuOtAnAmE",
				ResourceProviderName: "rEsOuRcEpRoViDeRnAmE",
				GroupQuotaLimitName:  "gRoUpQuOtAlImItNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/rEsOuRcEpRoViDeRs/rEsOuRcEpRoViDeRnAmE/gRoUpQuOtAlImItS/gRoUpQuOtAlImItNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupQuotaLimitIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupId != v.Expected.ManagementGroupId {
			t.Fatalf("Expected %q but got %q for ManagementGroupId", v.Expected.ManagementGroupId, actual.ManagementGroupId)
		}

		if actual.GroupQuotaName != v.Expected.GroupQuotaName {
			t.Fatalf("Expected %q but got %q for GroupQuotaName", v.Expected.GroupQuotaName, actual.GroupQuotaName)
		}

		if actual.ResourceProviderName != v.Expected.ResourceProviderName {
			t.Fatalf("Expected %q but got %q for ResourceProviderName", v.Expected.ResourceProviderName, actual.ResourceProviderName)
		}

		if actual.GroupQuotaLimitName != v.Expected.GroupQuotaLimitName {
			t.Fatalf("Expected %q but got %q for GroupQuotaLimitName", v.Expected.GroupQuotaLimitName, actual.GroupQuotaLimitName)
		}

	}
}

func TestSegmentsForGroupQuotaLimitId(t *testing.T) {
	segments := GroupQuotaLimitId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupQuotaLimitId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
