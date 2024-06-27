package billingroleassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingProfileBillingRoleAssignmentId{}

func TestNewBillingProfileBillingRoleAssignmentID(t *testing.T) {
	id := NewBillingProfileBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "billingRoleAssignmentValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.BillingProfileName != "billingProfileValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileValue")
	}

	if id.BillingRoleAssignmentName != "billingRoleAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleAssignmentName'", id.BillingRoleAssignmentName, "billingRoleAssignmentValue")
	}
}

func TestFormatBillingProfileBillingRoleAssignmentID(t *testing.T) {
	actual := NewBillingProfileBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "billingRoleAssignmentValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments/billingRoleAssignmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingProfileBillingRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileBillingRoleAssignmentId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments/billingRoleAssignmentValue",
			Expected: &BillingProfileBillingRoleAssignmentId{
				BillingAccountName:        "billingAccountValue",
				BillingProfileName:        "billingProfileValue",
				BillingRoleAssignmentName: "billingRoleAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments/billingRoleAssignmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileBillingRoleAssignmentID(v.Input)
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

		if actual.BillingRoleAssignmentName != v.Expected.BillingRoleAssignmentName {
			t.Fatalf("Expected %q but got %q for BillingRoleAssignmentName", v.Expected.BillingRoleAssignmentName, actual.BillingRoleAssignmentName)
		}

	}
}

func TestParseBillingProfileBillingRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileBillingRoleAssignmentId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/bIlLiNgRoLeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments/billingRoleAssignmentValue",
			Expected: &BillingProfileBillingRoleAssignmentId{
				BillingAccountName:        "billingAccountValue",
				BillingProfileName:        "billingProfileValue",
				BillingRoleAssignmentName: "billingRoleAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/billingRoleAssignments/billingRoleAssignmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/bIlLiNgRoLeAsSiGnMeNtS/bIlLiNgRoLeAsSiGnMeNtVaLuE",
			Expected: &BillingProfileBillingRoleAssignmentId{
				BillingAccountName:        "bIlLiNgAcCoUnTvAlUe",
				BillingProfileName:        "bIlLiNgPrOfIlEvAlUe",
				BillingRoleAssignmentName: "bIlLiNgRoLeAsSiGnMeNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/bIlLiNgRoLeAsSiGnMeNtS/bIlLiNgRoLeAsSiGnMeNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileBillingRoleAssignmentIDInsensitively(v.Input)
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

		if actual.BillingRoleAssignmentName != v.Expected.BillingRoleAssignmentName {
			t.Fatalf("Expected %q but got %q for BillingRoleAssignmentName", v.Expected.BillingRoleAssignmentName, actual.BillingRoleAssignmentName)
		}

	}
}

func TestSegmentsForBillingProfileBillingRoleAssignmentId(t *testing.T) {
	segments := BillingProfileBillingRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingProfileBillingRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
