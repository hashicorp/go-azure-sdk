package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdHomeRealmDiscoveryPolicyId{}

func TestNewServicePrincipalIdHomeRealmDiscoveryPolicyID(t *testing.T) {
	id := NewServicePrincipalIdHomeRealmDiscoveryPolicyID("servicePrincipalId", "homeRealmDiscoveryPolicyId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.HomeRealmDiscoveryPolicyId != "homeRealmDiscoveryPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'HomeRealmDiscoveryPolicyId'", id.HomeRealmDiscoveryPolicyId, "homeRealmDiscoveryPolicyId")
	}
}

func TestFormatServicePrincipalIdHomeRealmDiscoveryPolicyID(t *testing.T) {
	actual := NewServicePrincipalIdHomeRealmDiscoveryPolicyID("servicePrincipalId", "homeRealmDiscoveryPolicyId").ID()
	expected := "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdHomeRealmDiscoveryPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdHomeRealmDiscoveryPolicyId
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
			Input: "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId",
			Expected: &ServicePrincipalIdHomeRealmDiscoveryPolicyId{
				ServicePrincipalId:         "servicePrincipalId",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdHomeRealmDiscoveryPolicyID(v.Input)
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

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

	}
}

func TestParseServicePrincipalIdHomeRealmDiscoveryPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdHomeRealmDiscoveryPolicyId
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
			Input: "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/hOmErEaLmDiScOvErYpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId",
			Expected: &ServicePrincipalIdHomeRealmDiscoveryPolicyId{
				ServicePrincipalId:         "servicePrincipalId",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiD",
			Expected: &ServicePrincipalIdHomeRealmDiscoveryPolicyId{
				ServicePrincipalId:         "sErViCePrInCiPaLiD",
				HomeRealmDiscoveryPolicyId: "hOmErEaLmDiScOvErYpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdHomeRealmDiscoveryPolicyIDInsensitively(v.Input)
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

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

	}
}

func TestSegmentsForServicePrincipalIdHomeRealmDiscoveryPolicyId(t *testing.T) {
	segments := ServicePrincipalIdHomeRealmDiscoveryPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdHomeRealmDiscoveryPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
