package policydefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyDefinitionId{}

func TestNewPolicyDefinitionID(t *testing.T) {
	id := NewPolicyDefinitionID("policyDefinitionValue")

	if id.PolicyDefinitionName != "policyDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyDefinitionName'", id.PolicyDefinitionName, "policyDefinitionValue")
	}
}

func TestFormatPolicyDefinitionID(t *testing.T) {
	actual := NewPolicyDefinitionID("policyDefinitionValue").ID()
	expected := "/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyDefinitionId
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
			Input: "/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue",
			Expected: &PolicyDefinitionId{
				PolicyDefinitionName: "policyDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicyDefinitionName != v.Expected.PolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicyDefinitionName", v.Expected.PolicyDefinitionName, actual.PolicyDefinitionName)
		}

	}
}

func TestParsePolicyDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyDefinitionId
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
			Input: "/providers/Microsoft.Authorization/policyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue",
			Expected: &PolicyDefinitionId{
				PolicyDefinitionName: "policyDefinitionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/policyDefinitions/policyDefinitionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNvAlUe",
			Expected: &PolicyDefinitionId{
				PolicyDefinitionName: "pOlIcYdEfInItIoNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pOlIcYdEfInItIoNs/pOlIcYdEfInItIoNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicyDefinitionName != v.Expected.PolicyDefinitionName {
			t.Fatalf("Expected %q but got %q for PolicyDefinitionName", v.Expected.PolicyDefinitionName, actual.PolicyDefinitionName)
		}

	}
}

func TestSegmentsForPolicyDefinitionId(t *testing.T) {
	segments := PolicyDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
