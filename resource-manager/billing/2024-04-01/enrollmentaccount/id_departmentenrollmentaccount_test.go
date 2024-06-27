package enrollmentaccount

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DepartmentEnrollmentAccountId{}

func TestNewDepartmentEnrollmentAccountID(t *testing.T) {
	id := NewDepartmentEnrollmentAccountID("billingAccountValue", "departmentValue", "enrollmentAccountValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.DepartmentName != "departmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DepartmentName'", id.DepartmentName, "departmentValue")
	}

	if id.EnrollmentAccountName != "enrollmentAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentAccountName'", id.EnrollmentAccountName, "enrollmentAccountValue")
	}
}

func TestFormatDepartmentEnrollmentAccountID(t *testing.T) {
	actual := NewDepartmentEnrollmentAccountID("billingAccountValue", "departmentValue", "enrollmentAccountValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts/enrollmentAccountValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDepartmentEnrollmentAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DepartmentEnrollmentAccountId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts/enrollmentAccountValue",
			Expected: &DepartmentEnrollmentAccountId{
				BillingAccountName:    "billingAccountValue",
				DepartmentName:        "departmentValue",
				EnrollmentAccountName: "enrollmentAccountValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts/enrollmentAccountValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDepartmentEnrollmentAccountID(v.Input)
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

		if actual.DepartmentName != v.Expected.DepartmentName {
			t.Fatalf("Expected %q but got %q for DepartmentName", v.Expected.DepartmentName, actual.DepartmentName)
		}

		if actual.EnrollmentAccountName != v.Expected.EnrollmentAccountName {
			t.Fatalf("Expected %q but got %q for EnrollmentAccountName", v.Expected.EnrollmentAccountName, actual.EnrollmentAccountName)
		}

	}
}

func TestParseDepartmentEnrollmentAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DepartmentEnrollmentAccountId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe/eNrOlLmEnTaCcOuNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts/enrollmentAccountValue",
			Expected: &DepartmentEnrollmentAccountId{
				BillingAccountName:    "billingAccountValue",
				DepartmentName:        "departmentValue",
				EnrollmentAccountName: "enrollmentAccountValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/enrollmentAccounts/enrollmentAccountValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE",
			Expected: &DepartmentEnrollmentAccountId{
				BillingAccountName:    "bIlLiNgAcCoUnTvAlUe",
				DepartmentName:        "dEpArTmEnTvAlUe",
				EnrollmentAccountName: "eNrOlLmEnTaCcOuNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDepartmentEnrollmentAccountIDInsensitively(v.Input)
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

		if actual.DepartmentName != v.Expected.DepartmentName {
			t.Fatalf("Expected %q but got %q for DepartmentName", v.Expected.DepartmentName, actual.DepartmentName)
		}

		if actual.EnrollmentAccountName != v.Expected.EnrollmentAccountName {
			t.Fatalf("Expected %q but got %q for EnrollmentAccountName", v.Expected.EnrollmentAccountName, actual.EnrollmentAccountName)
		}

	}
}

func TestSegmentsForDepartmentEnrollmentAccountId(t *testing.T) {
	segments := DepartmentEnrollmentAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DepartmentEnrollmentAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
