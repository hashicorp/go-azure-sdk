package billingrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingRequestId{}

func TestNewBillingRequestID(t *testing.T) {
	id := NewBillingRequestID("billingRequestName")

	if id.BillingRequestName != "billingRequestName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRequestName'", id.BillingRequestName, "billingRequestName")
	}
}

func TestFormatBillingRequestID(t *testing.T) {
	actual := NewBillingRequestID("billingRequestName").ID()
	expected := "/providers/Microsoft.Billing/billingRequests/billingRequestName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingRequestId
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
			Input: "/providers/Microsoft.Billing/billingRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingRequests/billingRequestName",
			Expected: &BillingRequestId{
				BillingRequestName: "billingRequestName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingRequests/billingRequestName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingRequestName != v.Expected.BillingRequestName {
			t.Fatalf("Expected %q but got %q for BillingRequestName", v.Expected.BillingRequestName, actual.BillingRequestName)
		}

	}
}

func TestParseBillingRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingRequestId
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
			Input: "/providers/Microsoft.Billing/billingRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingRequests/billingRequestName",
			Expected: &BillingRequestId{
				BillingRequestName: "billingRequestName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingRequests/billingRequestName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgReQuEsTs/bIlLiNgReQuEsTnAmE",
			Expected: &BillingRequestId{
				BillingRequestName: "bIlLiNgReQuEsTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgReQuEsTs/bIlLiNgReQuEsTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingRequestName != v.Expected.BillingRequestName {
			t.Fatalf("Expected %q but got %q for BillingRequestName", v.Expected.BillingRequestName, actual.BillingRequestName)
		}

	}
}

func TestSegmentsForBillingRequestId(t *testing.T) {
	segments := BillingRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
