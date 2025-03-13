package policysetdefinitionversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicySetDefinitionId{}

func TestNewPolicySetDefinitionID(t *testing.T) {
	id := NewPolicySetDefinitionID("policySetDefinitionName")

	if id.PolicySetDefinitionName != "policySetDefinitionName" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicySetDefinitionName'", id.PolicySetDefinitionName, "policySetDefinitionName")
	}
}

func TestFormatPolicySetDefinitionID(t *testing.T) {
	actual := NewPolicySetDefinitionID("policySetDefinitionName").ID()
	expected := "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicySetDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicySetDefinitionId
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
			// Valid URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName",
			Expected: &PolicySetDefinitionId{
				PolicySetDefinitionName: "policySetDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicySetDefinitionID(v.Input)
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

	}
}

func TestParsePolicySetDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicySetDefinitionId
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
			// Valid URI
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName",
			Expected: &PolicySetDefinitionId{
				PolicySetDefinitionName: "policySetDefinitionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/policySetDefinitions/policySetDefinitionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe",
			Expected: &PolicySetDefinitionId{
				PolicySetDefinitionName: "pOlIcYsEtDeFiNiTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYsEtDeFiNiTiOnS/pOlIcYsEtDeFiNiTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicySetDefinitionIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPolicySetDefinitionId(t *testing.T) {
	segments := PolicySetDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicySetDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
