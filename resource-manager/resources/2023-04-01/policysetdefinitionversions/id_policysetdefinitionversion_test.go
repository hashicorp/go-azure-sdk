package policysetdefinitionversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicySetDefinitionVersionId{}

func TestNewPolicySetDefinitionVersionID(t *testing.T) {
	id := NewPolicySetDefinitionVersionID("policySetDefinitionValue", "versionValue")

	if id.PolicySetDefinitionName != "policySetDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicySetDefinitionName'", id.PolicySetDefinitionName, "policySetDefinitionValue")
	}

	if id.VersionName != "versionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionName'", id.VersionName, "versionValue")
	}
}

func TestFormatPolicySetDefinitionVersionID(t *testing.T) {
	actual := NewPolicySetDefinitionVersionID("policySetDefinitionValue", "versionValue").ID()
	expected := "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicySetDefinitionVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicySetDefinitionVersionId
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
			Input: "/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue",
			Expected: &PolicySetDefinitionVersionId{
				PolicySetDefinitionName: "policySetDefinitionValue",
				VersionName:             "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicySetDefinitionVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicySetDefinitionName != v.Expected.PolicySetDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicySetDefinitionName", v.Expected.PolicySetDefinitionName, actual.PolicySetDefinitionName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestParsePolicySetDefinitionVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicySetDefinitionVersionId
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
			Input: "/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue",
			Expected: &PolicySetDefinitionVersionId{
				PolicySetDefinitionName: "policySetDefinitionValue",
				VersionName:             "versionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionValue/versions/versionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/vErSiOnS/vErSiOnVaLuE",
			Expected: &PolicySetDefinitionVersionId{
				PolicySetDefinitionName: "pOlIcYsEtDeFiNiTiOnVaLuE",
				VersionName:             "vErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnVaLuE/vErSiOnS/vErSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicySetDefinitionVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicySetDefinitionName != v.Expected.PolicySetDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicySetDefinitionName", v.Expected.PolicySetDefinitionName, actual.PolicySetDefinitionName)
		}

		if actual.VersionName != v.Expected.VersionName {
			t.Fatalf("Expected %q but got %q for VersionName", v.Expected.VersionName, actual.VersionName)
		}

	}
}

func TestSegmentsForPolicySetDefinitionVersionId(t *testing.T) {
	segments := PolicySetDefinitionVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicySetDefinitionVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
