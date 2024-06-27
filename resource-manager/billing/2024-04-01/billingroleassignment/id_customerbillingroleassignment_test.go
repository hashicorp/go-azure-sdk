package billingroleassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CustomerBillingRoleAssignmentId{}

func TestNewCustomerBillingRoleAssignmentID(t *testing.T) {
	id := NewCustomerBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "customerValue", "billingRoleAssignmentValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.BillingProfileName != "billingProfileValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileValue")
	}

	if id.CustomerName != "customerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomerName'", id.CustomerName, "customerValue")
	}

	if id.BillingRoleAssignmentName != "billingRoleAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleAssignmentName'", id.BillingRoleAssignmentName, "billingRoleAssignmentValue")
	}
}

func TestFormatCustomerBillingRoleAssignmentID(t *testing.T) {
	actual := NewCustomerBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "customerValue", "billingRoleAssignmentValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments/billingRoleAssignmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCustomerBillingRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CustomerBillingRoleAssignmentId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments/billingRoleAssignmentValue",
			Expected: &CustomerBillingRoleAssignmentId{
				BillingAccountName:        "billingAccountValue",
				BillingProfileName:        "billingProfileValue",
				CustomerName:              "customerValue",
				BillingRoleAssignmentName: "billingRoleAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments/billingRoleAssignmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCustomerBillingRoleAssignmentID(v.Input)
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

		if actual.CustomerName != v.Expected.CustomerName {
			t.Fatalf("Expected %q but got %q for CustomerName", v.Expected.CustomerName, actual.CustomerName)
		}

		if actual.BillingRoleAssignmentName != v.Expected.BillingRoleAssignmentName {
			t.Fatalf("Expected %q but got %q for BillingRoleAssignmentName", v.Expected.BillingRoleAssignmentName, actual.BillingRoleAssignmentName)
		}

	}
}

func TestParseCustomerBillingRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CustomerBillingRoleAssignmentId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/cUsToMeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/cUsToMeRs/cUsToMeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/cUsToMeRs/cUsToMeRvAlUe/bIlLiNgRoLeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments/billingRoleAssignmentValue",
			Expected: &CustomerBillingRoleAssignmentId{
				BillingAccountName:        "billingAccountValue",
				BillingProfileName:        "billingProfileValue",
				CustomerName:              "customerValue",
				BillingRoleAssignmentName: "billingRoleAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/customers/customerValue/billingRoleAssignments/billingRoleAssignmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/cUsToMeRs/cUsToMeRvAlUe/bIlLiNgRoLeAsSiGnMeNtS/bIlLiNgRoLeAsSiGnMeNtVaLuE",
			Expected: &CustomerBillingRoleAssignmentId{
				BillingAccountName:        "bIlLiNgAcCoUnTvAlUe",
				BillingProfileName:        "bIlLiNgPrOfIlEvAlUe",
				CustomerName:              "cUsToMeRvAlUe",
				BillingRoleAssignmentName: "bIlLiNgRoLeAsSiGnMeNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/cUsToMeRs/cUsToMeRvAlUe/bIlLiNgRoLeAsSiGnMeNtS/bIlLiNgRoLeAsSiGnMeNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCustomerBillingRoleAssignmentIDInsensitively(v.Input)
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

		if actual.CustomerName != v.Expected.CustomerName {
			t.Fatalf("Expected %q but got %q for CustomerName", v.Expected.CustomerName, actual.CustomerName)
		}

		if actual.BillingRoleAssignmentName != v.Expected.BillingRoleAssignmentName {
			t.Fatalf("Expected %q but got %q for BillingRoleAssignmentName", v.Expected.BillingRoleAssignmentName, actual.BillingRoleAssignmentName)
		}

	}
}

func TestSegmentsForCustomerBillingRoleAssignmentId(t *testing.T) {
	segments := CustomerBillingRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CustomerBillingRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
