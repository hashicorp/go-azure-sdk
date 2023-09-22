package policysetdefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2PolicySetDefinitionId{}

func TestNewProviders2PolicySetDefinitionID(t *testing.T) {
	id := NewProviders2PolicySetDefinitionID("managementGroupValue", "policySetDefinitionValue")

	if id.ManagementGroupName != "managementGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupName'", id.ManagementGroupName, "managementGroupValue")
	}

	if id.PolicySetDefinitionName != "policySetDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicySetDefinitionName'", id.PolicySetDefinitionName, "policySetDefinitionValue")
	}
}

func TestFormatProviders2PolicySetDefinitionID(t *testing.T) {
	actual := NewProviders2PolicySetDefinitionID("managementGroupValue", "policySetDefinitionValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2PolicySetDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicySetDefinitionId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue",
			Expected: &Providers2PolicySetDefinitionId{
				ManagementGroupName:     "managementGroupValue",
				PolicySetDefinitionName: "policySetDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicySetDefinitionID(v.Input)
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

		if actual.PolicySetDefinitionName != v.Expected.PolicySetDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicySetDefinitionName", v.Expected.PolicySetDefinitionName, actual.PolicySetDefinitionName)
		}

	}
}

func TestParseProviders2PolicySetDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicySetDefinitionId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue",
			Expected: &Providers2PolicySetDefinitionId{
				ManagementGroupName:     "managementGroupValue",
				PolicySetDefinitionName: "policySetDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE",
			Expected: &Providers2PolicySetDefinitionId{
				ManagementGroupName:     "mAnAgEmEnTgRoUpVaLuE",
				PolicySetDefinitionName: "pOlIcYsEtDeFiNiTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicySetDefinitionIDInsensitively(v.Input)
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

		if actual.PolicySetDefinitionName != v.Expected.PolicySetDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicySetDefinitionName", v.Expected.PolicySetDefinitionName, actual.PolicySetDefinitionName)
		}

	}
}

func TestSegmentsForProviders2PolicySetDefinitionId(t *testing.T) {
	segments := Providers2PolicySetDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2PolicySetDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
