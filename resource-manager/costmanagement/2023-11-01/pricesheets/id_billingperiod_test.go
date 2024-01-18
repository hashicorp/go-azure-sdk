package pricesheets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingPeriodId{}

func TestNewBillingPeriodID(t *testing.T) {
	id := NewBillingPeriodID("billingAccountIdValue", "billingPeriodValue")

	if id.BillingAccountId != "billingAccountIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountId'", id.BillingAccountId, "billingAccountIdValue")
	}

	if id.BillingPeriodName != "billingPeriodValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingPeriodName'", id.BillingPeriodName, "billingPeriodValue")
	}
}

func TestFormatBillingPeriodID(t *testing.T) {
	actual := NewBillingPeriodID("billingAccountIdValue", "billingPeriodValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods/billingPeriodValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingPeriodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingPeriodId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods/billingPeriodValue",
			Expected: &BillingPeriodId{
				BillingAccountId:  "billingAccountIdValue",
				BillingPeriodName: "billingPeriodValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods/billingPeriodValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingPeriodID(v.Input)
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

		if actual.BillingPeriodName != v.Expected.BillingPeriodName {
			t.Fatalf("Expected %q but got %q for BillingPeriodName", v.Expected.BillingPeriodName, actual.BillingPeriodName)
		}

	}
}

func TestParseBillingPeriodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingPeriodId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/bIlLiNgPeRiOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods/billingPeriodValue",
			Expected: &BillingPeriodId{
				BillingAccountId:  "billingAccountIdValue",
				BillingPeriodName: "billingPeriodValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingPeriods/billingPeriodValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/bIlLiNgPeRiOdS/bIlLiNgPeRiOdVaLuE",
			Expected: &BillingPeriodId{
				BillingAccountId:  "bIlLiNgAcCoUnTiDvAlUe",
				BillingPeriodName: "bIlLiNgPeRiOdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/bIlLiNgPeRiOdS/bIlLiNgPeRiOdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingPeriodIDInsensitively(v.Input)
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

		if actual.BillingPeriodName != v.Expected.BillingPeriodName {
			t.Fatalf("Expected %q but got %q for BillingPeriodName", v.Expected.BillingPeriodName, actual.BillingPeriodName)
		}

	}
}

func TestSegmentsForBillingPeriodId(t *testing.T) {
	segments := BillingPeriodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingPeriodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
