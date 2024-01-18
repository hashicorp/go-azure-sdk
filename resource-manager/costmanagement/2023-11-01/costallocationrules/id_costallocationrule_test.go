package costallocationrules

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CostAllocationRuleId{}

func TestNewCostAllocationRuleID(t *testing.T) {
	id := NewCostAllocationRuleID("billingAccountIdValue", "costAllocationRuleValue")

	if id.BillingAccountId != "billingAccountIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountId'", id.BillingAccountId, "billingAccountIdValue")
	}

	if id.CostAllocationRuleName != "costAllocationRuleValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CostAllocationRuleName'", id.CostAllocationRuleName, "costAllocationRuleValue")
	}
}

func TestFormatCostAllocationRuleID(t *testing.T) {
	actual := NewCostAllocationRuleID("billingAccountIdValue", "costAllocationRuleValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCostAllocationRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CostAllocationRuleId
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
			Input: "/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleValue",
			Expected: &CostAllocationRuleId{
				BillingAccountId:       "billingAccountIdValue",
				CostAllocationRuleName: "costAllocationRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCostAllocationRuleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountId != v.Expected.BillingAccountId {
			t.Fatalf("Expected %q but got %q for BillingAccountId", v.Expected.BillingAccountId, actual.BillingAccountId)
		}

		if actual.CostAllocationRuleName != v.Expected.CostAllocationRuleName {
			t.Fatalf("Expected %q but got %q for CostAllocationRuleName", v.Expected.CostAllocationRuleName, actual.CostAllocationRuleName)
		}

	}
}

func TestParseCostAllocationRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CostAllocationRuleId
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
			Input: "/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/cOsTaLlOcAtIoNrUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleValue",
			Expected: &CostAllocationRuleId{
				BillingAccountId:       "billingAccountIdValue",
				CostAllocationRuleName: "costAllocationRuleValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/cOsTaLlOcAtIoNrUlEs/cOsTaLlOcAtIoNrUlEvAlUe",
			Expected: &CostAllocationRuleId{
				BillingAccountId:       "bIlLiNgAcCoUnTiDvAlUe",
				CostAllocationRuleName: "cOsTaLlOcAtIoNrUlEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/cOsTaLlOcAtIoNrUlEs/cOsTaLlOcAtIoNrUlEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCostAllocationRuleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountId != v.Expected.BillingAccountId {
			t.Fatalf("Expected %q but got %q for BillingAccountId", v.Expected.BillingAccountId, actual.BillingAccountId)
		}

		if actual.CostAllocationRuleName != v.Expected.CostAllocationRuleName {
			t.Fatalf("Expected %q but got %q for CostAllocationRuleName", v.Expected.CostAllocationRuleName, actual.CostAllocationRuleName)
		}

	}
}

func TestSegmentsForCostAllocationRuleId(t *testing.T) {
	segments := CostAllocationRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CostAllocationRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
