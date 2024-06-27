package billingroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EnrollmentAccountBillingRoleDefinitionId{}

func TestNewEnrollmentAccountBillingRoleDefinitionID(t *testing.T) {
	id := NewEnrollmentAccountBillingRoleDefinitionID("billingAccountValue", "enrollmentAccountValue", "billingRoleDefinitionValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.EnrollmentAccountName != "enrollmentAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentAccountName'", id.EnrollmentAccountName, "enrollmentAccountValue")
	}

	if id.BillingRoleDefinitionName != "billingRoleDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleDefinitionName'", id.BillingRoleDefinitionName, "billingRoleDefinitionValue")
	}
}

func TestFormatEnrollmentAccountBillingRoleDefinitionID(t *testing.T) {
	actual := NewEnrollmentAccountBillingRoleDefinitionID("billingAccountValue", "enrollmentAccountValue", "billingRoleDefinitionValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions/billingRoleDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseEnrollmentAccountBillingRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnrollmentAccountBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions/billingRoleDefinitionValue",
			Expected: &EnrollmentAccountBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountValue",
				EnrollmentAccountName:     "enrollmentAccountValue",
				BillingRoleDefinitionName: "billingRoleDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions/billingRoleDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnrollmentAccountBillingRoleDefinitionID(v.Input)
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

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestParseEnrollmentAccountBillingRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EnrollmentAccountBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/bIlLiNgRoLeDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions/billingRoleDefinitionValue",
			Expected: &EnrollmentAccountBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountValue",
				EnrollmentAccountName:     "enrollmentAccountValue",
				BillingRoleDefinitionName: "billingRoleDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/enrollmentAccounts/enrollmentAccountValue/billingRoleDefinitions/billingRoleDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnVaLuE",
			Expected: &EnrollmentAccountBillingRoleDefinitionId{
				BillingAccountName:        "bIlLiNgAcCoUnTvAlUe",
				EnrollmentAccountName:     "eNrOlLmEnTaCcOuNtVaLuE",
				BillingRoleDefinitionName: "bIlLiNgRoLeDeFiNiTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtVaLuE/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseEnrollmentAccountBillingRoleDefinitionIDInsensitively(v.Input)
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

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestSegmentsForEnrollmentAccountBillingRoleDefinitionId(t *testing.T) {
	segments := EnrollmentAccountBillingRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("EnrollmentAccountBillingRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
