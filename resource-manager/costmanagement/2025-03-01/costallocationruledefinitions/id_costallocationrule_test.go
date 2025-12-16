package costallocationruledefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CostAllocationRuleId{}

func TestNewCostAllocationRuleID(t *testing.T) {
	id := NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

	if id.BillingAccountId != "billingAccountId" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountId'", id.BillingAccountId, "billingAccountId")
	}

	if id.CostAllocationRuleName != "costAllocationRuleName" {
		t.Fatalf("Expected %q but got %q for Segment 'CostAllocationRuleName'", id.CostAllocationRuleName, "costAllocationRuleName")
	}
}

func TestFormatCostAllocationRuleID(t *testing.T) {
	actual := NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleName"
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleName",
			Expected: &CostAllocationRuleId{
				BillingAccountId:       "billingAccountId",
				CostAllocationRuleName: "costAllocationRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleName/extra",
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiD/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiD/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiD/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/cOsTaLlOcAtIoNrUlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleName",
			Expected: &CostAllocationRuleId{
				BillingAccountId:       "billingAccountId",
				CostAllocationRuleName: "costAllocationRuleName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountId/providers/Microsoft.CostManagement/costAllocationRules/costAllocationRuleName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiD/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/cOsTaLlOcAtIoNrUlEs/cOsTaLlOcAtIoNrUlEnAmE",
			Expected: &CostAllocationRuleId{
				BillingAccountId:       "bIlLiNgAcCoUnTiD",
				CostAllocationRuleName: "cOsTaLlOcAtIoNrUlEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiD/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/cOsTaLlOcAtIoNrUlEs/cOsTaLlOcAtIoNrUlEnAmE/extra",
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
