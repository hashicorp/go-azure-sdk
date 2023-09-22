package policysetdefinitionversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2PolicySetDefinitionVersionId{}

func TestNewProviders2PolicySetDefinitionVersionID(t *testing.T) {
	id := NewProviders2PolicySetDefinitionVersionID("managementGroupValue", "policySetDefinitionValue", "versionValue")

	if id.ManagementGroupName != "managementGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupName'", id.ManagementGroupName, "managementGroupValue")
	}

	if id.PolicySetDefinitionName != "policySetDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicySetDefinitionName'", id.PolicySetDefinitionName, "policySetDefinitionValue")
	}

	if id.VersionName != "versionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionValue")
	}
}

func TestFormatProviders2PolicySetDefinitionVersionID(t *testing.T) {
	actual := NewProviders2PolicySetDefinitionVersionID("managementGroupValue", "policySetDefinitionValue", "versionValue").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2PolicySetDefinitionVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicySetDefinitionVersionId
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
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue",
			Expected: &Providers2PolicySetDefinitionVersionId{
				ManagementGroupName:     "managementGroupValue",
				PolicySetDefinitionName: "policySetDefinitionValue",
				VersionName:             "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicySetDefinitionVersionID(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseProviders2PolicySetDefinitionVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicySetDefinitionVersionId
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
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue",
			Expected: &Providers2PolicySetDefinitionVersionId{
				ManagementGroupName:     "managementGroupValue",
				PolicySetDefinitionName: "policySetDefinitionValue",
				VersionName:             "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupValue/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/vErSiOnS/vErSiOnVaLuE",
			Expected: &Providers2PolicySetDefinitionVersionId{
				ManagementGroupName:     "mAnAgEmEnTgRoUpVaLuE",
				PolicySetDefinitionName: "pOlIcYsEtDeFiNiTiOnVaLuE",
				VersionName:             "vErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/vErSiOnS/vErSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicySetDefinitionVersionIDInsensitively(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForProviders2PolicySetDefinitionVersionId(t *testing.T) {
	segments := Providers2PolicySetDefinitionVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2PolicySetDefinitionVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
