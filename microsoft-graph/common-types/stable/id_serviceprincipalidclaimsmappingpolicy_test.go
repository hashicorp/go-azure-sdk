package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdClaimsMappingPolicyId{}

func TestNewServicePrincipalIdClaimsMappingPolicyID(t *testing.T) {
	id := NewServicePrincipalIdClaimsMappingPolicyID("servicePrincipalId", "claimsMappingPolicyId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.ClaimsMappingPolicyId != "claimsMappingPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'ClaimsMappingPolicyId'", id.ClaimsMappingPolicyId, "claimsMappingPolicyId")
	}
}

func TestFormatServicePrincipalIdClaimsMappingPolicyID(t *testing.T) {
	actual := NewServicePrincipalIdClaimsMappingPolicyID("servicePrincipalId", "claimsMappingPolicyId").ID()
	expected := "/servicePrincipals/servicePrincipalId/claimsMappingPolicies/claimsMappingPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdClaimsMappingPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdClaimsMappingPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/claimsMappingPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/claimsMappingPolicies/claimsMappingPolicyId",
			Expected: &ServicePrincipalIdClaimsMappingPolicyId{
				ServicePrincipalId:    "servicePrincipalId",
				ClaimsMappingPolicyId: "claimsMappingPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/claimsMappingPolicies/claimsMappingPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdClaimsMappingPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.ClaimsMappingPolicyId != v.Expected.ClaimsMappingPolicyId {
			t.Fatalf("Expected %q but got %q for ClaimsMappingPolicyId", v.Expected.ClaimsMappingPolicyId, actual.ClaimsMappingPolicyId)
		}

	}
}

func TestParseServicePrincipalIdClaimsMappingPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdClaimsMappingPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/claimsMappingPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/cLaImSmApPiNgPoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/claimsMappingPolicies/claimsMappingPolicyId",
			Expected: &ServicePrincipalIdClaimsMappingPolicyId{
				ServicePrincipalId:    "servicePrincipalId",
				ClaimsMappingPolicyId: "claimsMappingPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/claimsMappingPolicies/claimsMappingPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId",
			Expected: &ServicePrincipalIdClaimsMappingPolicyId{
				ServicePrincipalId:    "sErViCePrInCiPaLiD",
				ClaimsMappingPolicyId: "cLaImSmApPiNgPoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/cLaImSmApPiNgPoLiCiEs/cLaImSmApPiNgPoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdClaimsMappingPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.ClaimsMappingPolicyId != v.Expected.ClaimsMappingPolicyId {
			t.Fatalf("Expected %q but got %q for ClaimsMappingPolicyId", v.Expected.ClaimsMappingPolicyId, actual.ClaimsMappingPolicyId)
		}

	}
}

func TestSegmentsForServicePrincipalIdClaimsMappingPolicyId(t *testing.T) {
	segments := ServicePrincipalIdClaimsMappingPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdClaimsMappingPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
