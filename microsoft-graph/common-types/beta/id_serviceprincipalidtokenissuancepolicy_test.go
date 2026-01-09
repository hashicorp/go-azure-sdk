package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdTokenIssuancePolicyId{}

func TestNewServicePrincipalIdTokenIssuancePolicyID(t *testing.T) {
	id := NewServicePrincipalIdTokenIssuancePolicyID("servicePrincipalId", "tokenIssuancePolicyId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.TokenIssuancePolicyId != "tokenIssuancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenIssuancePolicyId'", id.TokenIssuancePolicyId, "tokenIssuancePolicyId")
	}
}

func TestFormatServicePrincipalIdTokenIssuancePolicyID(t *testing.T) {
	actual := NewServicePrincipalIdTokenIssuancePolicyID("servicePrincipalId", "tokenIssuancePolicyId").ID()
	expected := "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies/tokenIssuancePolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdTokenIssuancePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdTokenIssuancePolicyId
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
			Input: "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies/tokenIssuancePolicyId",
			Expected: &ServicePrincipalIdTokenIssuancePolicyId{
				ServicePrincipalId:    "servicePrincipalId",
				TokenIssuancePolicyId: "tokenIssuancePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies/tokenIssuancePolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdTokenIssuancePolicyID(v.Input)
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

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

	}
}

func TestParseServicePrincipalIdTokenIssuancePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdTokenIssuancePolicyId
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
			Input: "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/tOkEnIsSuAnCePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies/tokenIssuancePolicyId",
			Expected: &ServicePrincipalIdTokenIssuancePolicyId{
				ServicePrincipalId:    "servicePrincipalId",
				TokenIssuancePolicyId: "tokenIssuancePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/tokenIssuancePolicies/tokenIssuancePolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId",
			Expected: &ServicePrincipalIdTokenIssuancePolicyId{
				ServicePrincipalId:    "sErViCePrInCiPaLiD",
				TokenIssuancePolicyId: "tOkEnIsSuAnCePoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdTokenIssuancePolicyIDInsensitively(v.Input)
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

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

	}
}

func TestSegmentsForServicePrincipalIdTokenIssuancePolicyId(t *testing.T) {
	segments := ServicePrincipalIdTokenIssuancePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdTokenIssuancePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
