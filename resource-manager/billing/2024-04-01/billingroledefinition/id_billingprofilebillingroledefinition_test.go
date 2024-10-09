package billingroledefinition

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingProfileBillingRoleDefinitionId{}

func TestNewBillingProfileBillingRoleDefinitionID(t *testing.T) {
	id := NewBillingProfileBillingRoleDefinitionID("billingAccountName", "billingProfileName", "billingRoleDefinitionName")

	if id.BillingAccountName != "billingAccountName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountName'", id.BillingAccountName, "billingAccountName")
	}

	if id.BillingProfileName != "billingProfileName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileName'", id.BillingProfileName, "billingProfileName")
	}

	if id.BillingRoleDefinitionName != "billingRoleDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingRoleDefinitionName'", id.BillingRoleDefinitionName, "billingRoleDefinitionName")
	}
}

func TestFormatBillingProfileBillingRoleDefinitionID(t *testing.T) {
	actual := NewBillingProfileBillingRoleDefinitionID("billingAccountName", "billingProfileName", "billingRoleDefinitionName").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions/billingRoleDefinitionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingProfileBillingRoleDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions/billingRoleDefinitionName",
			Expected: &BillingProfileBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountName",
				BillingProfileName:        "billingProfileName",
				BillingRoleDefinitionName: "billingRoleDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions/billingRoleDefinitionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileBillingRoleDefinitionID(v.Input)
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

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestParseBillingProfileBillingRoleDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileBillingRoleDefinitionId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE/bIlLiNgRoLeDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions/billingRoleDefinitionName",
			Expected: &BillingProfileBillingRoleDefinitionId{
				BillingAccountName:        "billingAccountName",
				BillingProfileName:        "billingProfileName",
				BillingRoleDefinitionName: "billingRoleDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountName/billingProfiles/billingProfileName/billingRoleDefinitions/billingRoleDefinitionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnNaMe",
			Expected: &BillingProfileBillingRoleDefinitionId{
				BillingAccountName:        "bIlLiNgAcCoUnTnAmE",
				BillingProfileName:        "bIlLiNgPrOfIlEnAmE",
				BillingRoleDefinitionName: "bIlLiNgRoLeDeFiNiTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTnAmE/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEnAmE/bIlLiNgRoLeDeFiNiTiOnS/bIlLiNgRoLeDeFiNiTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileBillingRoleDefinitionIDInsensitively(v.Input)
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

		if actual.BillingRoleDefinitionName != v.Expected.BillingRoleDefinitionName {
			t.Fatalf("Expected %q but got %q for BillingRoleDefinitionName", v.Expected.BillingRoleDefinitionName, actual.BillingRoleDefinitionName)
		}

	}
}

func TestSegmentsForBillingProfileBillingRoleDefinitionId(t *testing.T) {
	segments := BillingProfileBillingRoleDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingProfileBillingRoleDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
