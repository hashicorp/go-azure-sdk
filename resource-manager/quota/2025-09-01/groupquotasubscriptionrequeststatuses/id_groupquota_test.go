package groupquotasubscriptionrequeststatuses

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupQuotaId{}

func TestNewGroupQuotaID(t *testing.T) {
	id := NewGroupQuotaID("managementGroupId", "groupQuotaName")

	if id.ManagementGroupId != "managementGroupId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupId")
	}

	if id.GroupQuotaName != "groupQuotaName" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupQuotaName'", id.GroupQuotaName, "groupQuotaName")
	}
}

func TestFormatGroupQuotaID(t *testing.T) {
	actual := NewGroupQuotaID("managementGroupId", "groupQuotaName").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupQuotaID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupQuotaId
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
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Expected: &GroupQuotaId{
				ManagementGroupId: "managementGroupId",
				GroupQuotaName:    "groupQuotaName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupQuotaID(v.Input)
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

	}
}

func TestParseGroupQuotaIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupQuotaId
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
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName",
			Expected: &GroupQuotaId{
				ManagementGroupId: "managementGroupId",
				GroupQuotaName:    "groupQuotaName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupId/providers/Microsoft.Quota/groupQuotas/groupQuotaName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE",
			Expected: &GroupQuotaId{
				ManagementGroupId: "mAnAgEmEnTgRoUpId",
				GroupQuotaName:    "gRoUpQuOtAnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpId/pRoViDeRs/mIcRoSoFt.qUoTa/gRoUpQuOtAs/gRoUpQuOtAnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupQuotaIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForGroupQuotaId(t *testing.T) {
	segments := GroupQuotaId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupQuotaId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
