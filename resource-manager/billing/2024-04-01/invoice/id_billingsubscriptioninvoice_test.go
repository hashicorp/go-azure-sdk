package invoice

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingSubscriptionInvoiceId{}

func TestNewBillingSubscriptionInvoiceID(t *testing.T) {
	id := NewBillingSubscriptionInvoiceID("subscriptionIdValue", "invoiceValue")

	if id.SubscriptionId != "subscriptionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "subscriptionIdValue")
	}

	if id.InvoiceName != "invoiceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvoiceName'", id.InvoiceName, "invoiceValue")
	}
}

func TestFormatBillingSubscriptionInvoiceID(t *testing.T) {
	actual := NewBillingSubscriptionInvoiceID("subscriptionIdValue", "invoiceValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices/invoiceValue"
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
			Input: "/providers/Microsoft.Billing/billingAccounts/default",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices/invoiceValue",
			Expected: &BillingSubscriptionInvoiceId{
				SubscriptionId: "subscriptionIdValue",
				InvoiceName:    "invoiceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices/invoiceValue/extra",
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

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
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
			Input: "/providers/Microsoft.Billing/billingAccounts/default",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/dEfAuLt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/dEfAuLt/bIlLiNgSuBsCrIpTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/dEfAuLt/bIlLiNgSuBsCrIpTiOnS/sUbScRiPtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/dEfAuLt/bIlLiNgSuBsCrIpTiOnS/sUbScRiPtIoNiDvAlUe/iNvOiCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices/invoiceValue",
			Expected: &BillingSubscriptionInvoiceId{
				SubscriptionId: "subscriptionIdValue",
				InvoiceName:    "invoiceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/subscriptionIdValue/invoices/invoiceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/dEfAuLt/bIlLiNgSuBsCrIpTiOnS/sUbScRiPtIoNiDvAlUe/iNvOiCeS/iNvOiCeVaLuE",
			Expected: &BillingSubscriptionInvoiceId{
				SubscriptionId: "sUbScRiPtIoNiDvAlUe",
				InvoiceName:    "iNvOiCeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/dEfAuLt/bIlLiNgSuBsCrIpTiOnS/sUbScRiPtIoNiDvAlUe/iNvOiCeS/iNvOiCeVaLuE/extra",
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

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
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
