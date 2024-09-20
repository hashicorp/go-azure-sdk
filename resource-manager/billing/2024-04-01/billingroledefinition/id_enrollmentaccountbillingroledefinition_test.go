package billingroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EnrollmentAccountBillingRoleDefinitionId{}

func TestNewEnrollmentAccountBillingRoleDefinitionID(t *testing.T) {
	id := NewEnrollmentAccountBillingRoleDefinitionID("billingAccountName", "enrollmentAccountName", "roleDefinitionName")

	if id.BillingAccountName != "billingAccountName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountName")
	}

	if id.EnrollmentAccountName != "enrollmentAccountName" {
		t.Fatalf("Expected %q but got %q for Segment 'EnrollmentAccountName'", id.EnrollmentAccountName, "enrollmentAccountName")
	}

	if id.BillingRoleDefinitionName != "roleDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleDefinitionName'", id.BillingRoleDefinitionName, "roleDefinitionName")
	}
}

func TestFormatEnrollmentAccountBillingRoleDefinitionID(t *testing.T) {
	actual := NewEnrollmentAccountBillingRoleDefinitionID("billingAccountName", "enrollmentAccountName", "roleDefinitionName").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions/roleDefinitionName"
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions/roleDefinitionName",
			Expected: &EnrollmentAccountBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountName",
				EnrollmentAccountName:     "enrollmentAccountName",
				BillingRoleDefinitionName: "roleDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions/roleDefinitionName/extra",
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/eNrOlLmEnTaCcOuNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtNaMe/bIlLiNgRoLeDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions/roleDefinitionName",
			Expected: &EnrollmentAccountBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountName",
				EnrollmentAccountName:     "enrollmentAccountName",
				BillingRoleDefinitionName: "roleDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/enrollmentAccounts/enrollmentAccountName/billingRoleDefinitions/roleDefinitionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtNaMe/bIlLiNgRoLeDeFiNiTiOnS/rOlEdEfInItIoNnAmE",
			Expected: &EnrollmentAccountBillingRoleDefinitionId{
				BillingAccountName:        "bIlLiNgAcCoUnTnAmE",
				EnrollmentAccountName:     "eNrOlLmEnTaCcOuNtNaMe",
				BillingRoleDefinitionName: "rOlEdEfInItIoNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/eNrOlLmEnTaCcOuNtS/eNrOlLmEnTaCcOuNtNaMe/bIlLiNgRoLeDeFiNiTiOnS/rOlEdEfInItIoNnAmE/extra",
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
