package policyserviceprincipalcreationpolicyinclude

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyServicePrincipalCreationPolicyIncludeId{}

func TestNewPolicyServicePrincipalCreationPolicyIncludeID(t *testing.T) {
	id := NewPolicyServicePrincipalCreationPolicyIncludeID("servicePrincipalCreationPolicyIdValue", "servicePrincipalCreationConditionSetIdValue")

	if id.ServicePrincipalCreationPolicyId != "servicePrincipalCreationPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalCreationPolicyId'", id.ServicePrincipalCreationPolicyId, "servicePrincipalCreationPolicyIdValue")
	}

	if id.ServicePrincipalCreationConditionSetId != "servicePrincipalCreationConditionSetIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalCreationConditionSetId'", id.ServicePrincipalCreationConditionSetId, "servicePrincipalCreationConditionSetIdValue")
	}
}

func TestFormatPolicyServicePrincipalCreationPolicyIncludeID(t *testing.T) {
	actual := NewPolicyServicePrincipalCreationPolicyIncludeID("servicePrincipalCreationPolicyIdValue", "servicePrincipalCreationConditionSetIdValue").ID()
	expected := "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes/servicePrincipalCreationConditionSetIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyServicePrincipalCreationPolicyIncludeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyServicePrincipalCreationPolicyIncludeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes/servicePrincipalCreationConditionSetIdValue",
			Expected: &PolicyServicePrincipalCreationPolicyIncludeId{
				ServicePrincipalCreationPolicyId:       "servicePrincipalCreationPolicyIdValue",
				ServicePrincipalCreationConditionSetId: "servicePrincipalCreationConditionSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes/servicePrincipalCreationConditionSetIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyServicePrincipalCreationPolicyIncludeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalCreationPolicyId != v.Expected.ServicePrincipalCreationPolicyId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalCreationPolicyId", v.Expected.ServicePrincipalCreationPolicyId, actual.ServicePrincipalCreationPolicyId)
		}

		if actual.ServicePrincipalCreationConditionSetId != v.Expected.ServicePrincipalCreationConditionSetId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalCreationConditionSetId", v.Expected.ServicePrincipalCreationConditionSetId, actual.ServicePrincipalCreationConditionSetId)
		}

	}
}

func TestParsePolicyServicePrincipalCreationPolicyIncludeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyServicePrincipalCreationPolicyIncludeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiDvAlUe/iNcLuDeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes/servicePrincipalCreationConditionSetIdValue",
			Expected: &PolicyServicePrincipalCreationPolicyIncludeId{
				ServicePrincipalCreationPolicyId:       "servicePrincipalCreationPolicyIdValue",
				ServicePrincipalCreationConditionSetId: "servicePrincipalCreationConditionSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyIdValue/includes/servicePrincipalCreationConditionSetIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiDvAlUe/iNcLuDeS/sErViCePrInCiPaLcReAtIoNcOnDiTiOnSeTiDvAlUe",
			Expected: &PolicyServicePrincipalCreationPolicyIncludeId{
				ServicePrincipalCreationPolicyId:       "sErViCePrInCiPaLcReAtIoNpOlIcYiDvAlUe",
				ServicePrincipalCreationConditionSetId: "sErViCePrInCiPaLcReAtIoNcOnDiTiOnSeTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiDvAlUe/iNcLuDeS/sErViCePrInCiPaLcReAtIoNcOnDiTiOnSeTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyServicePrincipalCreationPolicyIncludeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalCreationPolicyId != v.Expected.ServicePrincipalCreationPolicyId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalCreationPolicyId", v.Expected.ServicePrincipalCreationPolicyId, actual.ServicePrincipalCreationPolicyId)
		}

		if actual.ServicePrincipalCreationConditionSetId != v.Expected.ServicePrincipalCreationConditionSetId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalCreationConditionSetId", v.Expected.ServicePrincipalCreationConditionSetId, actual.ServicePrincipalCreationConditionSetId)
		}

	}
}

func TestSegmentsForPolicyServicePrincipalCreationPolicyIncludeId(t *testing.T) {
	segments := PolicyServicePrincipalCreationPolicyIncludeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyServicePrincipalCreationPolicyIncludeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
