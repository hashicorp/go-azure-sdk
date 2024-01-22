package invoices

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingSubscriptionInvoiceId{}

func TestNewBillingSubscriptionInvoiceID(t *testing.T) {
	id := NewBillingSubscriptionInvoiceID("billingAccountValue", "billingSubscriptionValue", "invoiceValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.BillingSubscriptionName != "billingSubscriptionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingSubscriptionName'", id.BillingSubscriptionName, "billingSubscriptionValue")
	}

	if id.InvoiceName != "invoiceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvoiceName'", id.InvoiceName, "invoiceValue")
	}
}

func TestFormatBillingSubscriptionInvoiceID(t *testing.T) {
	actual := NewBillingSubscriptionInvoiceID("billingAccountValue", "billingSubscriptionValue", "invoiceValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices/invoiceValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingSubscriptionInvoiceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingSubscriptionInvoiceId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices/invoiceValue",
			Expected: &BillingSubscriptionInvoiceId{
				BillingAccountName:      "billingAccountValue",
				BillingSubscriptionName: "billingSubscriptionValue",
				InvoiceName:             "invoiceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices/invoiceValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingSubscriptionInvoiceID(v.Input)
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

		if actual.BillingSubscriptionName != v.Expected.BillingSubscriptionName {
			t.Fatalf("Expected %q but got %q for BillingSubscriptionName", v.Expected.BillingSubscriptionName, actual.BillingSubscriptionName)
		}

		if actual.InvoiceName != v.Expected.InvoiceName {
			t.Fatalf("Expected %q but got %q for InvoiceName", v.Expected.InvoiceName, actual.InvoiceName)
		}

	}
}

func TestParseBillingSubscriptionInvoiceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingSubscriptionInvoiceId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgSuBsCrIpTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgSuBsCrIpTiOnS/bIlLiNgSuBsCrIpTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgSuBsCrIpTiOnS/bIlLiNgSuBsCrIpTiOnVaLuE/iNvOiCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices/invoiceValue",
			Expected: &BillingSubscriptionInvoiceId{
				BillingAccountName:      "billingAccountValue",
				BillingSubscriptionName: "billingSubscriptionValue",
				InvoiceName:             "invoiceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingSubscriptions/billingSubscriptionValue/invoices/invoiceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgSuBsCrIpTiOnS/bIlLiNgSuBsCrIpTiOnVaLuE/iNvOiCeS/iNvOiCeVaLuE",
			Expected: &BillingSubscriptionInvoiceId{
				BillingAccountName:      "bIlLiNgAcCoUnTvAlUe",
				BillingSubscriptionName: "bIlLiNgSuBsCrIpTiOnVaLuE",
				InvoiceName:             "iNvOiCeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgSuBsCrIpTiOnS/bIlLiNgSuBsCrIpTiOnVaLuE/iNvOiCeS/iNvOiCeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingSubscriptionInvoiceIDInsensitively(v.Input)
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

		if actual.BillingSubscriptionName != v.Expected.BillingSubscriptionName {
			t.Fatalf("Expected %q but got %q for BillingSubscriptionName", v.Expected.BillingSubscriptionName, actual.BillingSubscriptionName)
		}

		if actual.InvoiceName != v.Expected.InvoiceName {
			t.Fatalf("Expected %q but got %q for InvoiceName", v.Expected.InvoiceName, actual.InvoiceName)
		}

	}
}

func TestSegmentsForBillingSubscriptionInvoiceId(t *testing.T) {
	segments := BillingSubscriptionInvoiceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingSubscriptionInvoiceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
