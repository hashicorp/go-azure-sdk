package policydefinitionversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &Providers2PolicyDefinitionVersionId{}

func TestNewProviders2PolicyDefinitionVersionID(t *testing.T) {
	id := NewProviders2PolicyDefinitionVersionID("managementGroupName", "policyDefinitionName", "versionName")

	if id.ManagementGroupName != "managementGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupName'", id.ManagementGroupName, "managementGroupName")
	}

	if id.PolicyDefinitionName != "policyDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyDefinitionName'", id.PolicyDefinitionName, "policyDefinitionName")
	}

	if id.VersionName != "versionName" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionName")
	}
}

func TestFormatProviders2PolicyDefinitionVersionID(t *testing.T) {
	actual := NewProviders2PolicyDefinitionVersionID("managementGroupName", "policyDefinitionName", "versionName").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2PolicyDefinitionVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicyDefinitionVersionId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName",
			Expected: &Providers2PolicyDefinitionVersionId{
				ManagementGroupName:  "managementGroupName",
				PolicyDefinitionName: "policyDefinitionName",
				VersionName:          "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicyDefinitionVersionID(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParseProviders2PolicyDefinitionVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicyDefinitionVersionId
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
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName",
			Expected: &Providers2PolicyDefinitionVersionId{
				ManagementGroupName:  "managementGroupName",
				PolicyDefinitionName: "policyDefinitionName",
				VersionName:          "versionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionName/versions/versionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE/vErSiOnS/vErSiOnNaMe",
			Expected: &Providers2PolicyDefinitionVersionId{
				ManagementGroupName:  "mAnAgEmEnTgRoUpNaMe",
				PolicyDefinitionName: "pOlIcYdEfInItIoNnAmE",
				VersionName:          "vErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNnAmE/vErSiOnS/vErSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicyDefinitionVersionIDInsensitively(v.Input)
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

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForProviders2PolicyDefinitionVersionId(t *testing.T) {
	segments := Providers2PolicyDefinitionVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2PolicyDefinitionVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
