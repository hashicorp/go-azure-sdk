package savingsplan

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SavingsPlanId{}

func TestNewSavingsPlanID(t *testing.T) {
	id := NewSavingsPlanID("billingAccountValue", "savingsPlanOrderIdValue", "savingsPlanIdValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.SavingsPlanOrderId != "savingsPlanOrderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SavingsPlanOrderId'", id.SavingsPlanOrderId, "savingsPlanOrderIdValue")
	}

	if id.SavingsPlanId != "savingsPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SavingsPlanId'", id.SavingsPlanId, "savingsPlanIdValue")
	}
}

func TestFormatSavingsPlanID(t *testing.T) {
	actual := NewSavingsPlanID("billingAccountValue", "savingsPlanOrderIdValue", "savingsPlanIdValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSavingsPlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SavingsPlanId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue",
			Expected: &SavingsPlanId{
				BillingAccountName: "billingAccountValue",
				SavingsPlanOrderId: "savingsPlanOrderIdValue",
				SavingsPlanId:      "savingsPlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSavingsPlanID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountName != v.Expected.BillingAccountName {
			t.Fatalf("Expected %q but got %q for BillingAccountName", v.Expected.BillingAccountName, actual.BillingAccountName)
		}

		if actual.SavingsPlanOrderId != v.Expected.SavingsPlanOrderId {
			t.Fatalf("Expected %q but got %q for SavingsPlanOrderId", v.Expected.SavingsPlanOrderId, actual.SavingsPlanOrderId)
		}

		if actual.SavingsPlanId != v.Expected.SavingsPlanId {
			t.Fatalf("Expected %q but got %q for SavingsPlanId", v.Expected.SavingsPlanId, actual.SavingsPlanId)
		}

	}
}

func TestParseSavingsPlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SavingsPlanId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/sAvInGsPlAnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue",
			Expected: &SavingsPlanId{
				BillingAccountName: "billingAccountValue",
				SavingsPlanOrderId: "savingsPlanOrderIdValue",
				SavingsPlanId:      "savingsPlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/sAvInGsPlAnS/sAvInGsPlAnIdVaLuE",
			Expected: &SavingsPlanId{
				BillingAccountName: "bIlLiNgAcCoUnTvAlUe",
				SavingsPlanOrderId: "sAvInGsPlAnOrDeRiDvAlUe",
				SavingsPlanId:      "sAvInGsPlAnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/sAvInGsPlAnS/sAvInGsPlAnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSavingsPlanIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountName != v.Expected.BillingAccountName {
			t.Fatalf("Expected %q but got %q for BillingAccountName", v.Expected.BillingAccountName, actual.BillingAccountName)
		}

		if actual.SavingsPlanOrderId != v.Expected.SavingsPlanOrderId {
			t.Fatalf("Expected %q but got %q for SavingsPlanOrderId", v.Expected.SavingsPlanOrderId, actual.SavingsPlanOrderId)
		}

		if actual.SavingsPlanId != v.Expected.SavingsPlanId {
			t.Fatalf("Expected %q but got %q for SavingsPlanId", v.Expected.SavingsPlanId, actual.SavingsPlanId)
		}

	}
}

func TestSegmentsForSavingsPlanId(t *testing.T) {
	segments := SavingsPlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SavingsPlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
