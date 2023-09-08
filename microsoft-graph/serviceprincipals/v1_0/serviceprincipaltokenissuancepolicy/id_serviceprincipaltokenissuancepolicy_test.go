package serviceprincipaltokenissuancepolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalTokenIssuancePolicyId{}

func TestNewServicePrincipalTokenIssuancePolicyID(t *testing.T) {
	id := NewServicePrincipalTokenIssuancePolicyID("servicePrincipalIdValue", "tokenIssuancePolicyIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.TokenIssuancePolicyId != "tokenIssuancePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenIssuancePolicyId'", id.TokenIssuancePolicyId, "tokenIssuancePolicyIdValue")
	}
}

func TestFormatServicePrincipalTokenIssuancePolicyID(t *testing.T) {
	actual := NewServicePrincipalTokenIssuancePolicyID("servicePrincipalIdValue", "tokenIssuancePolicyIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies/tokenIssuancePolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalTokenIssuancePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalTokenIssuancePolicyId
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies/tokenIssuancePolicyIdValue",
			Expected: &ServicePrincipalTokenIssuancePolicyId{
				ServicePrincipalId:    "servicePrincipalIdValue",
				TokenIssuancePolicyId: "tokenIssuancePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies/tokenIssuancePolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalTokenIssuancePolicyID(v.Input)
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

func TestParseServicePrincipalTokenIssuancePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalTokenIssuancePolicyId
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/tOkEnIsSuAnCePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies/tokenIssuancePolicyIdValue",
			Expected: &ServicePrincipalTokenIssuancePolicyId{
				ServicePrincipalId:    "servicePrincipalIdValue",
				TokenIssuancePolicyId: "tokenIssuancePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenIssuancePolicies/tokenIssuancePolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE",
			Expected: &ServicePrincipalTokenIssuancePolicyId{
				ServicePrincipalId:    "sErViCePrInCiPaLiDvAlUe",
				TokenIssuancePolicyId: "tOkEnIsSuAnCePoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalTokenIssuancePolicyIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalTokenIssuancePolicyId(t *testing.T) {
	segments := ServicePrincipalTokenIssuancePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalTokenIssuancePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
