package billingroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvoiceSectionBillingRoleDefinitionId{}

func TestNewInvoiceSectionBillingRoleDefinitionID(t *testing.T) {
	id := NewInvoiceSectionBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleDefinitionValue")

	if id.BillingAccountName != "billingAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountValue")
	}

	if id.BillingProfileName != "billingProfileValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileValue")
	}

	if id.InvoiceSectionName != "invoiceSectionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvoiceSectionName'", id.InvoiceSectionName, "invoiceSectionValue")
	}

	if id.BillingRoleDefinitionName != "billingRoleDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleDefinitionName'", id.BillingRoleDefinitionName, "billingRoleDefinitionValue")
	}
}

func TestFormatInvoiceSectionBillingRoleDefinitionID(t *testing.T) {
	actual := NewInvoiceSectionBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleDefinitionValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions/billingRoleDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseInvoiceSectionBillingRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvoiceSectionBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions/billingRoleDefinitionValue",
			Expected: &InvoiceSectionBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountValue",
				BillingProfileName:        "billingProfileValue",
				InvoiceSectionName:        "invoiceSectionValue",
				BillingRoleDefinitionName: "billingRoleDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions/billingRoleDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvoiceSectionBillingRoleDefinitionID(v.Input)
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

		if actual.InvoiceSectionName != v.Expected.InvoiceSectionName {
			t.Fatalf("Expected %q but got %q for InvoiceSectionName", v.Expected.InvoiceSectionName, actual.InvoiceSectionName)
		}

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestParseInvoiceSectionBillingRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvoiceSectionBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/iNvOiCeSeCtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/iNvOiCeSeCtIoNs/iNvOiCeSeCtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/iNvOiCeSeCtIoNs/iNvOiCeSeCtIoNvAlUe/bIlLiNgRoLeDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions/billingRoleDefinitionValue",
			Expected: &InvoiceSectionBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountValue",
				BillingProfileName:        "billingProfileValue",
				InvoiceSectionName:        "invoiceSectionValue",
				BillingRoleDefinitionName: "billingRoleDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountValue/billingProfiles/billingProfileValue/invoiceSections/invoiceSectionValue/billingRoleDefinitions/billingRoleDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/iNvOiCeSeCtIoNs/iNvOiCeSeCtIoNvAlUe/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnVaLuE",
			Expected: &InvoiceSectionBillingRoleDefinitionId{
				BillingAccountName:        "bIlLiNgAcCoUnTvAlUe",
				BillingProfileName:        "bIlLiNgPrOfIlEvAlUe",
				InvoiceSectionName:        "iNvOiCeSeCtIoNvAlUe",
				BillingRoleDefinitionName: "bIlLiNgRoLeDeFiNiTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEvAlUe/iNvOiCeSeCtIoNs/iNvOiCeSeCtIoNvAlUe/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvoiceSectionBillingRoleDefinitionIDInsensitively(v.Input)
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

		if actual.InvoiceSectionName != v.Expected.InvoiceSectionName {
			t.Fatalf("Expected %q but got %q for InvoiceSectionName", v.Expected.InvoiceSectionName, actual.InvoiceSectionName)
		}

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestSegmentsForInvoiceSectionBillingRoleDefinitionId(t *testing.T) {
	segments := InvoiceSectionBillingRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("InvoiceSectionBillingRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
