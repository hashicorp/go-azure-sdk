package savingsplan

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SavingsPlanOrderId{}

func TestNewSavingsPlanOrderID(t *testing.T) {
	id := NewSavingsPlanOrderID("billingAccountValue", "savingsPlanOrderIdValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.SavingsPlanOrderId != "savingsPlanOrderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SavingsPlanOrderId'", id.SavingsPlanOrderId, "savingsPlanOrderIdValue")
	}
}

func TestFormatSavingsPlanOrderID(t *testing.T) {
	actual := NewSavingsPlanOrderID("billingAccountValue", "savingsPlanOrderIdValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSavingsPlanOrderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SavingsPlanOrderId
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
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue",
			Expected: &SavingsPlanOrderId{
				BillingAccountName: "billingAccountValue",
				SavingsPlanOrderId: "savingsPlanOrderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSavingsPlanOrderID(v.Input)
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

	}
}

func TestParseSavingsPlanOrderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SavingsPlanOrderId
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
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue",
			Expected: &SavingsPlanOrderId{
				BillingAccountName: "billingAccountValue",
				SavingsPlanOrderId: "savingsPlanOrderIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/savingsPlanOrders/savingsPlanOrderIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe",
			Expected: &SavingsPlanOrderId{
				BillingAccountName: "bIlLiNgAcCoUnTvAlUe",
				SavingsPlanOrderId: "sAvInGsPlAnOrDeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSavingsPlanOrderIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForSavingsPlanOrderId(t *testing.T) {
	segments := SavingsPlanOrderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SavingsPlanOrderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
