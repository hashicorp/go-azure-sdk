package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyServicePrincipalCreationPolicyIdExcludeId{}

func TestNewPolicyServicePrincipalCreationPolicyIdExcludeID(t *testing.T) {
	id := NewPolicyServicePrincipalCreationPolicyIdExcludeID("servicePrincipalCreationPolicyId", "servicePrincipalCreationConditionSetId")

	if id.ServicePrincipalCreationPolicyId != "servicePrincipalCreationPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalCreationPolicyId'", id.ServicePrincipalCreationPolicyId, "servicePrincipalCreationPolicyId")
	}

	if id.ServicePrincipalCreationConditionSetId != "servicePrincipalCreationConditionSetId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalCreationConditionSetId'", id.ServicePrincipalCreationConditionSetId, "servicePrincipalCreationConditionSetId")
	}
}

func TestFormatPolicyServicePrincipalCreationPolicyIdExcludeID(t *testing.T) {
	actual := NewPolicyServicePrincipalCreationPolicyIdExcludeID("servicePrincipalCreationPolicyId", "servicePrincipalCreationConditionSetId").ID()
	expected := "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes/servicePrincipalCreationConditionSetId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyServicePrincipalCreationPolicyIdExcludeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyServicePrincipalCreationPolicyIdExcludeId
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
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes/servicePrincipalCreationConditionSetId",
			Expected: &PolicyServicePrincipalCreationPolicyIdExcludeId{
				ServicePrincipalCreationPolicyId:       "servicePrincipalCreationPolicyId",
				ServicePrincipalCreationConditionSetId: "servicePrincipalCreationConditionSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes/servicePrincipalCreationConditionSetId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyServicePrincipalCreationPolicyIdExcludeID(v.Input)
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

func TestParsePolicyServicePrincipalCreationPolicyIdExcludeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyServicePrincipalCreationPolicyIdExcludeId
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
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiD/eXcLuDeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes/servicePrincipalCreationConditionSetId",
			Expected: &PolicyServicePrincipalCreationPolicyIdExcludeId{
				ServicePrincipalCreationPolicyId:       "servicePrincipalCreationPolicyId",
				ServicePrincipalCreationConditionSetId: "servicePrincipalCreationConditionSetId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/servicePrincipalCreationPolicies/servicePrincipalCreationPolicyId/excludes/servicePrincipalCreationConditionSetId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiD/eXcLuDeS/sErViCePrInCiPaLcReAtIoNcOnDiTiOnSeTiD",
			Expected: &PolicyServicePrincipalCreationPolicyIdExcludeId{
				ServicePrincipalCreationPolicyId:       "sErViCePrInCiPaLcReAtIoNpOlIcYiD",
				ServicePrincipalCreationConditionSetId: "sErViCePrInCiPaLcReAtIoNcOnDiTiOnSeTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcIeS/sErViCePrInCiPaLcReAtIoNpOlIcYiD/eXcLuDeS/sErViCePrInCiPaLcReAtIoNcOnDiTiOnSeTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyServicePrincipalCreationPolicyIdExcludeIDInsensitively(v.Input)
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

func TestSegmentsForPolicyServicePrincipalCreationPolicyIdExcludeId(t *testing.T) {
	segments := PolicyServicePrincipalCreationPolicyIdExcludeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyServicePrincipalCreationPolicyIdExcludeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
