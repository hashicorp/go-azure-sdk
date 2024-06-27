package billingroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DepartmentBillingRoleDefinitionId{}

func TestNewDepartmentBillingRoleDefinitionID(t *testing.T) {
	id := NewDepartmentBillingRoleDefinitionID("billingAccountValue", "departmentValue", "billingRoleDefinitionValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.DepartmentName != "departmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DepartmentName'", id.DepartmentName, "departmentValue")
	}

	if id.BillingRoleDefinitionName != "billingRoleDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleDefinitionName'", id.BillingRoleDefinitionName, "billingRoleDefinitionValue")
	}
}

func TestFormatDepartmentBillingRoleDefinitionID(t *testing.T) {
	actual := NewDepartmentBillingRoleDefinitionID("billingAccountValue", "departmentValue", "billingRoleDefinitionValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions/billingRoleDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDepartmentBillingRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DepartmentBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions/billingRoleDefinitionValue",
			Expected: &DepartmentBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountValue",
				DepartmentName:            "departmentValue",
				BillingRoleDefinitionName: "billingRoleDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions/billingRoleDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDepartmentBillingRoleDefinitionID(v.Input)
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

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestParseDepartmentBillingRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DepartmentBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe/bIlLiNgRoLeDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions/billingRoleDefinitionValue",
			Expected: &DepartmentBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountValue",
				DepartmentName:            "departmentValue",
				BillingRoleDefinitionName: "billingRoleDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/departments/departmentValue/billingRoleDefinitions/billingRoleDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnVaLuE",
			Expected: &DepartmentBillingRoleDefinitionId{
				BillingAccountName:        "bIlLiNgAcCoUnTvAlUe",
				DepartmentName:            "dEpArTmEnTvAlUe",
				BillingRoleDefinitionName: "bIlLiNgRoLeDeFiNiTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/dEpArTmEnTs/dEpArTmEnTvAlUe/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDepartmentBillingRoleDefinitionIDInsensitively(v.Input)
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

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestSegmentsForDepartmentBillingRoleDefinitionId(t *testing.T) {
	segments := DepartmentBillingRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DepartmentBillingRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
