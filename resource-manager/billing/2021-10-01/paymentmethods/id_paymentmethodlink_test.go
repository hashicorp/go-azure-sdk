package paymentmethods

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PaymentMethodLinkId{}

func TestNewPaymentMethodLinkID(t *testing.T) {
	id := NewPaymentMethodLinkID("billingAccountName", "billingProfileName", "paymentMethodLinkName")

	if id.BillingAccountName != "billingAccountName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountName")
	}

	if id.BillingProfileName != "billingProfileName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileName")
	}

	if id.PaymentMethodLinkName != "paymentMethodLinkName" {
		t.Fatalf("Expected %q but got %q for Segment 'PaymentMethodLinkName'", id.PaymentMethodLinkName, "paymentMethodLinkName")
	}
}

func TestFormatPaymentMethodLinkID(t *testing.T) {
	actual := NewPaymentMethodLinkID("billingAccountName", "billingProfileName", "paymentMethodLinkName").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks/paymentMethodLinkName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePaymentMethodLinkID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PaymentMethodLinkId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks/paymentMethodLinkName",
			Expected: &PaymentMethodLinkId{
				BillingAccountName:    "billingAccountName",
				BillingProfileName:    "billingProfileName",
				PaymentMethodLinkName: "paymentMethodLinkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks/paymentMethodLinkName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePaymentMethodLinkID(v.Input)
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

		if actual.BillingProfileName != v.Expected.BillingProfileName {
			t.Fatalf("Expected %q but got %q for BillingProfileName", v.Expected.BillingProfileName, actual.BillingProfileName)
		}

		if actual.PaymentMethodLinkName != v.Expected.PaymentMethodLinkName {
			t.Fatalf("Expected %q but got %q for PaymentMethodLinkName", v.Expected.PaymentMethodLinkName, actual.PaymentMethodLinkName)
		}

	}
}

func TestParsePaymentMethodLinkIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PaymentMethodLinkId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE/pAyMeNtMeThOdLiNkS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks/paymentMethodLinkName",
			Expected: &PaymentMethodLinkId{
				BillingAccountName:    "billingAccountName",
				BillingProfileName:    "billingProfileName",
				PaymentMethodLinkName: "paymentMethodLinkName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/paymentMethodLinks/paymentMethodLinkName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE/pAyMeNtMeThOdLiNkS/pAyMeNtMeThOdLiNkNaMe",
			Expected: &PaymentMethodLinkId{
				BillingAccountName:    "bIlLiNgAcCoUnTnAmE",
				BillingProfileName:    "bIlLiNgPrOfIlEnAmE",
				PaymentMethodLinkName: "pAyMeNtMeThOdLiNkNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE/pAyMeNtMeThOdLiNkS/pAyMeNtMeThOdLiNkNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePaymentMethodLinkIDInsensitively(v.Input)
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

		if actual.BillingProfileName != v.Expected.BillingProfileName {
			t.Fatalf("Expected %q but got %q for BillingProfileName", v.Expected.BillingProfileName, actual.BillingProfileName)
		}

		if actual.PaymentMethodLinkName != v.Expected.PaymentMethodLinkName {
			t.Fatalf("Expected %q but got %q for PaymentMethodLinkName", v.Expected.PaymentMethodLinkName, actual.PaymentMethodLinkName)
		}

	}
}

func TestSegmentsForPaymentMethodLinkId(t *testing.T) {
	segments := PaymentMethodLinkId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PaymentMethodLinkId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
