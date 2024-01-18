package policydefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &Providers2PolicyDefinitionId{}

func TestNewProviders2PolicyDefinitionID(t *testing.T) {
	id := NewProviders2PolicyDefinitionID("managementGroupValue", "policyDefinitionValue")

	if id.ManagementGroupName != "managementGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupName'", id.ManagementGroupName, "managementGroupValue")
	}

	if id.PolicyDefinitionName != "policyDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyDefinitionName'", id.PolicyDefinitionName, "policyDefinitionValue")
	}
}

func TestFormatProviders2PolicyDefinitionID(t *testing.T) {
	actual := NewProviders2PolicyDefinitionID("managementGroupValue", "policyDefinitionValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2PolicyDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicyDefinitionId
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
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue",
			Expected: &Providers2PolicyDefinitionId{
				ManagementGroupName:  "managementGroupValue",
				PolicyDefinitionName: "policyDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicyDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupName != v.Expected.ManagementGroupName {
			t.Fatalf("Expected %q but got %q for ManagementGroupName", v.Expected.ManagementGroupName, actual.ManagementGroupName)
		}

		if actual.PolicyDefinitionName != v.Expected.PolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicyDefinitionName", v.Expected.PolicyDefinitionName, actual.PolicyDefinitionName)
		}

	}
}

func TestParseProviders2PolicyDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicyDefinitionId
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
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue",
			Expected: &Providers2PolicyDefinitionId{
				ManagementGroupName:  "managementGroupValue",
				PolicyDefinitionName: "policyDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNvAlUe",
			Expected: &Providers2PolicyDefinitionId{
				ManagementGroupName:  "mAnAgEmEnTgRoUpVaLuE",
				PolicyDefinitionName: "pOlIcYdEfInItIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicyDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupName != v.Expected.ManagementGroupName {
			t.Fatalf("Expected %q but got %q for ManagementGroupName", v.Expected.ManagementGroupName, actual.ManagementGroupName)
		}

		if actual.PolicyDefinitionName != v.Expected.PolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicyDefinitionName", v.Expected.PolicyDefinitionName, actual.PolicyDefinitionName)
		}

	}
}

func TestSegmentsForProviders2PolicyDefinitionId(t *testing.T) {
	segments := Providers2PolicyDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2PolicyDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
