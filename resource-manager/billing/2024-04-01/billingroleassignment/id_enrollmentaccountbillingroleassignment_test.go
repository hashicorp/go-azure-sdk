package billingroleassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EnrollmentAccountBillingRoleAssignmentId{}

func TestNewEnrollmentAccountBillingRoleAssignmentID(t *testing.T) {
	id := NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.EnrollmentAccountName != "enrollmentAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentAccountName'", id.EnrollmentAccountName, "enrollmentAccountValue")
	}

	if id.BillingRoleAssignmentName != "billingRoleAssignmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleAssignmentName'", id.BillingRoleAssignmentName, "billingRoleAssignmentValue")
	}
}

func TestFormatEnrollmentAccountBillingRoleAssignmentID(t *testing.T) {
	actual := NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments/billingRoleAssignmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseEnrollmentAccountBillingRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnrollmentAccountBillingRoleAssignmentId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments/billingRoleAssignmentValue",
			Expected: &EnrollmentAccountBillingRoleAssignmentId{
				BillingAccountName:        "billingAccountValue",
				EnrollmentAccountName:     "enrollmentAccountValue",
				BillingRoleAssignmentName: "billingRoleAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments/billingRoleAssignmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnrollmentAccountBillingRoleAssignmentID(v.Input)
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

		if actual.EnrollmentAccountName != v.Expected.EnrollmentAccountName {
			t.Fatalf("Expected %q but got %q for EnrollmentAccountName", v.Expected.EnrollmentAccountName, actual.EnrollmentAccountName)
		}

		if actual.BillingRoleAssignmentName != v.Expected.BillingRoleAssignmentName {
			t.Fatalf("Expected %q but got %q for BillingRoleAssignmentName", v.Expected.BillingRoleAssignmentName, actual.BillingRoleAssignmentName)
		}

	}
}

func TestParseEnrollmentAccountBillingRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnrollmentAccountBillingRoleAssignmentId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/bIlLiNgRoLeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments/billingRoleAssignmentValue",
			Expected: &EnrollmentAccountBillingRoleAssignmentId{
				BillingAccountName:        "billingAccountValue",
				EnrollmentAccountName:     "enrollmentAccountValue",
				BillingRoleAssignmentName: "billingRoleAssignmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleAssignments/billingRoleAssignmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/bIlLiNgRoLeAsSiGnMeNtS/bIlLiNgRoLeAsSiGnMeNtVaLuE",
			Expected: &EnrollmentAccountBillingRoleAssignmentId{
				BillingAccountName:        "bIlLiNgAcCoUnTvAlUe",
				EnrollmentAccountName:     "eNrOlLmEnTaCcOuNtVaLuE",
				BillingRoleAssignmentName: "bIlLiNgRoLeAsSiGnMeNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/bIlLiNgRoLeAsSiGnMeNtS/bIlLiNgRoLeAsSiGnMeNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnrollmentAccountBillingRoleAssignmentIDInsensitively(v.Input)
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

		if actual.EnrollmentAccountName != v.Expected.EnrollmentAccountName {
			t.Fatalf("Expected %q but got %q for EnrollmentAccountName", v.Expected.EnrollmentAccountName, actual.EnrollmentAccountName)
		}

		if actual.BillingRoleAssignmentName != v.Expected.BillingRoleAssignmentName {
			t.Fatalf("Expected %q but got %q for BillingRoleAssignmentName", v.Expected.BillingRoleAssignmentName, actual.BillingRoleAssignmentName)
		}

	}
}

func TestSegmentsForEnrollmentAccountBillingRoleAssignmentId(t *testing.T) {
	segments := EnrollmentAccountBillingRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("EnrollmentAccountBillingRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
