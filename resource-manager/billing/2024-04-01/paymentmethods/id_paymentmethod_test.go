package paymentmethods

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PaymentMethodId{}

func TestNewPaymentMethodID(t *testing.T) {
	id := NewPaymentMethodID("paymentMethodName")

	if id.PaymentMethodName != "paymentMethodName" {
		t.Fatalf("Expected %q but got %q for Segment 'PaymentMethodName'", id.PaymentMethodName, "paymentMethodName")
	}
}

func TestFormatPaymentMethodID(t *testing.T) {
	actual := NewPaymentMethodID("paymentMethodName").ID()
	expected := "/providers/Microsoft.Billing/paymentMethods/paymentMethodName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePaymentMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PaymentMethodId
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
			Input: "/providers/Microsoft.Billing/paymentMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/paymentMethods/paymentMethodName",
			Expected: &PaymentMethodId{
				PaymentMethodName: "paymentMethodName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/paymentMethods/paymentMethodName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePaymentMethodID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PaymentMethodName != v.Expected.PaymentMethodName {
			t.Fatalf("Expected %q but got %q for PaymentMethodName", v.Expected.PaymentMethodName, actual.PaymentMethodName)
		}

	}
}

func TestParsePaymentMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PaymentMethodId
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
			Input: "/providers/Microsoft.Billing/paymentMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/pAyMeNtMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/paymentMethods/paymentMethodName",
			Expected: &PaymentMethodId{
				PaymentMethodName: "paymentMethodName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/paymentMethods/paymentMethodName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/pAyMeNtMeThOdS/pAyMeNtMeThOdNaMe",
			Expected: &PaymentMethodId{
				PaymentMethodName: "pAyMeNtMeThOdNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/pAyMeNtMeThOdS/pAyMeNtMeThOdNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePaymentMethodIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PaymentMethodName != v.Expected.PaymentMethodName {
			t.Fatalf("Expected %q but got %q for PaymentMethodName", v.Expected.PaymentMethodName, actual.PaymentMethodName)
		}

	}
}

func TestSegmentsForPaymentMethodId(t *testing.T) {
	segments := PaymentMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PaymentMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
