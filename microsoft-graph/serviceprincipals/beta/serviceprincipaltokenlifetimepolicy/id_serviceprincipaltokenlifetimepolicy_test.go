package serviceprincipaltokenlifetimepolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalTokenLifetimePolicyId{}

func TestNewServicePrincipalTokenLifetimePolicyID(t *testing.T) {
	id := NewServicePrincipalTokenLifetimePolicyID("servicePrincipalIdValue", "tokenLifetimePolicyIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.TokenLifetimePolicyId != "tokenLifetimePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenLifetimePolicyId'", id.TokenLifetimePolicyId, "tokenLifetimePolicyIdValue")
	}
}

func TestFormatServicePrincipalTokenLifetimePolicyID(t *testing.T) {
	actual := NewServicePrincipalTokenLifetimePolicyID("servicePrincipalIdValue", "tokenLifetimePolicyIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalTokenLifetimePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalTokenLifetimePolicyId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue",
			Expected: &ServicePrincipalTokenLifetimePolicyId{
				ServicePrincipalId:    "servicePrincipalIdValue",
				TokenLifetimePolicyId: "tokenLifetimePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalTokenLifetimePolicyID(v.Input)
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

func TestParseServicePrincipalTokenLifetimePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalTokenLifetimePolicyId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/tOkEnLiFeTiMePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue",
			Expected: &ServicePrincipalTokenLifetimePolicyId{
				ServicePrincipalId:    "servicePrincipalIdValue",
				TokenLifetimePolicyId: "tokenLifetimePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyIdVaLuE",
			Expected: &ServicePrincipalTokenLifetimePolicyId{
				ServicePrincipalId:    "sErViCePrInCiPaLiDvAlUe",
				TokenLifetimePolicyId: "tOkEnLiFeTiMePoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalTokenLifetimePolicyIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalTokenLifetimePolicyId(t *testing.T) {
	segments := ServicePrincipalTokenLifetimePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalTokenLifetimePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
