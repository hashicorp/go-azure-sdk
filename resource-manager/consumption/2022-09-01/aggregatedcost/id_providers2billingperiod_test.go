package aggregatedcost

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &Providers2BillingPeriodId{}

func TestNewProviders2BillingPeriodID(t *testing.T) {
	id := NewProviders2BillingPeriodID("managementGroupIdValue", "billingPeriodValue")

	if id.ManagementGroupId != "managementGroupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupId'", id.ManagementGroupId, "managementGroupIdValue")
	}

	if id.BillingPeriodName != "billingPeriodValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingPeriodName'", id.BillingPeriodName, "billingPeriodValue")
	}
}

func TestFormatProviders2BillingPeriodID(t *testing.T) {
	actual := NewProviders2BillingPeriodID("managementGroupIdValue", "billingPeriodValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods/billingPeriodValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2BillingPeriodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2BillingPeriodId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods/billingPeriodValue",
			Expected: &Providers2BillingPeriodId{
				ManagementGroupId: "managementGroupIdValue",
				BillingPeriodName: "billingPeriodValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods/billingPeriodValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2BillingPeriodID(v.Input)
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

		if actual.BillingPeriodName != v.Expected.BillingPeriodName {
			t.Fatalf("Expected %q but got %q for BillingPeriodName", v.Expected.BillingPeriodName, actual.BillingPeriodName)
		}

	}
}

func TestParseProviders2BillingPeriodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2BillingPeriodId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.bIlLiNg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgPeRiOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods/billingPeriodValue",
			Expected: &Providers2BillingPeriodId{
				ManagementGroupId: "managementGroupIdValue",
				BillingPeriodName: "billingPeriodValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupIdValue/providers/Microsoft.Billing/billingPeriods/billingPeriodValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgPeRiOdS/bIlLiNgPeRiOdVaLuE",
			Expected: &Providers2BillingPeriodId{
				ManagementGroupId: "mAnAgEmEnTgRoUpIdVaLuE",
				BillingPeriodName: "bIlLiNgPeRiOdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpIdVaLuE/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgPeRiOdS/bIlLiNgPeRiOdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2BillingPeriodIDInsensitively(v.Input)
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

		if actual.BillingPeriodName != v.Expected.BillingPeriodName {
			t.Fatalf("Expected %q but got %q for BillingPeriodName", v.Expected.BillingPeriodName, actual.BillingPeriodName)
		}

	}
}

func TestSegmentsForProviders2BillingPeriodId(t *testing.T) {
	segments := Providers2BillingPeriodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2BillingPeriodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
