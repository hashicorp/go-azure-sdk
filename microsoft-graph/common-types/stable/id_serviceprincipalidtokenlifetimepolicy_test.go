package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdTokenLifetimePolicyId{}

func TestNewServicePrincipalIdTokenLifetimePolicyID(t *testing.T) {
	id := NewServicePrincipalIdTokenLifetimePolicyID("servicePrincipalId", "tokenLifetimePolicyId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.TokenLifetimePolicyId != "tokenLifetimePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenLifetimePolicyId'", id.TokenLifetimePolicyId, "tokenLifetimePolicyId")
	}
}

func TestFormatServicePrincipalIdTokenLifetimePolicyID(t *testing.T) {
	actual := NewServicePrincipalIdTokenLifetimePolicyID("servicePrincipalId", "tokenLifetimePolicyId").ID()
	expected := "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies/tokenLifetimePolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdTokenLifetimePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdTokenLifetimePolicyId
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
			Input: "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies/tokenLifetimePolicyId",
			Expected: &ServicePrincipalIdTokenLifetimePolicyId{
				ServicePrincipalId:    "servicePrincipalId",
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies/tokenLifetimePolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdTokenLifetimePolicyID(v.Input)
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

		if actual.TokenLifetimePolicyId != v.Expected.TokenLifetimePolicyId {
			t.Fatalf("Expected %q but got %q for TokenLifetimePolicyId", v.Expected.TokenLifetimePolicyId, actual.TokenLifetimePolicyId)
		}

	}
}

func TestParseServicePrincipalIdTokenLifetimePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdTokenLifetimePolicyId
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
			Input: "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/tOkEnLiFeTiMePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies/tokenLifetimePolicyId",
			Expected: &ServicePrincipalIdTokenLifetimePolicyId{
				ServicePrincipalId:    "servicePrincipalId",
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/tokenLifetimePolicies/tokenLifetimePolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId",
			Expected: &ServicePrincipalIdTokenLifetimePolicyId{
				ServicePrincipalId:    "sErViCePrInCiPaLiD",
				TokenLifetimePolicyId: "tOkEnLiFeTiMePoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdTokenLifetimePolicyIDInsensitively(v.Input)
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

		if actual.TokenLifetimePolicyId != v.Expected.TokenLifetimePolicyId {
			t.Fatalf("Expected %q but got %q for TokenLifetimePolicyId", v.Expected.TokenLifetimePolicyId, actual.TokenLifetimePolicyId)
		}

	}
}

func TestSegmentsForServicePrincipalIdTokenLifetimePolicyId(t *testing.T) {
	segments := ServicePrincipalIdTokenLifetimePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdTokenLifetimePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
