package billingsubscription

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingProfileBillingSubscriptionId{}

func TestNewBillingProfileBillingSubscriptionID(t *testing.T) {
	id := NewBillingProfileBillingSubscriptionID("billingAccountValue", "billingProfileValue", "billingSubscriptionValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.BillingProfileName != "billingProfileValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileValue")
	}

	if id.BillingSubscriptionName != "billingSubscriptionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingSubscriptionName'", id.BillingSubscriptionName, "billingSubscriptionValue")
	}
}

func TestFormatBillingProfileBillingSubscriptionID(t *testing.T) {
	actual := NewBillingProfileBillingSubscriptionID("billingAccountValue", "billingProfileValue", "billingSubscriptionValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions/billingSubscriptionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingProfileBillingSubscriptionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileBillingSubscriptionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions/billingSubscriptionValue",
			Expected: &BillingProfileBillingSubscriptionId{
				BillingAccountName:      "billingAccountValue",
				BillingProfileName:      "billingProfileValue",
				BillingSubscriptionName: "billingSubscriptionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions/billingSubscriptionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileBillingSubscriptionID(v.Input)
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

		if actual.BillingSubscriptionName != v.Expected.BillingSubscriptionName {
			t.Fatalf("Expected %q but got %q for BillingSubscriptionName", v.Expected.BillingSubscriptionName, actual.BillingSubscriptionName)
		}

	}
}

func TestParseBillingProfileBillingSubscriptionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileBillingSubscriptionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/bIlLiNgSuBsCrIpTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions/billingSubscriptionValue",
			Expected: &BillingProfileBillingSubscriptionId{
				BillingAccountName:      "billingAccountValue",
				BillingProfileName:      "billingProfileValue",
				BillingSubscriptionName: "billingSubscriptionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingSubscriptions/billingSubscriptionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/bIlLiNgSuBsCrIpTiOnS/bIlLiNgSuBsCrIpTiOnVaLuE",
			Expected: &BillingProfileBillingSubscriptionId{
				BillingAccountName:      "bIlLiNgAcCoUnTvAlUe",
				BillingProfileName:      "bIlLiNgPrOfIlEvAlUe",
				BillingSubscriptionName: "bIlLiNgSuBsCrIpTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/bIlLiNgSuBsCrIpTiOnS/bIlLiNgSuBsCrIpTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileBillingSubscriptionIDInsensitively(v.Input)
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

		if actual.BillingSubscriptionName != v.Expected.BillingSubscriptionName {
			t.Fatalf("Expected %q but got %q for BillingSubscriptionName", v.Expected.BillingSubscriptionName, actual.BillingSubscriptionName)
		}

	}
}

func TestSegmentsForBillingProfileBillingSubscriptionId(t *testing.T) {
	segments := BillingProfileBillingSubscriptionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingProfileBillingSubscriptionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
