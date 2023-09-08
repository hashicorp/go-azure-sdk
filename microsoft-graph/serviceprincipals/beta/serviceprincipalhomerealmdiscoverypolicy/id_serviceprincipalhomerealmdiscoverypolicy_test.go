package serviceprincipalhomerealmdiscoverypolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalHomeRealmDiscoveryPolicyId{}

func TestNewServicePrincipalHomeRealmDiscoveryPolicyID(t *testing.T) {
	id := NewServicePrincipalHomeRealmDiscoveryPolicyID("servicePrincipalIdValue", "homeRealmDiscoveryPolicyIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.HomeRealmDiscoveryPolicyId != "homeRealmDiscoveryPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HomeRealmDiscoveryPolicyId'", id.HomeRealmDiscoveryPolicyId, "homeRealmDiscoveryPolicyIdValue")
	}
}

func TestFormatServicePrincipalHomeRealmDiscoveryPolicyID(t *testing.T) {
	actual := NewServicePrincipalHomeRealmDiscoveryPolicyID("servicePrincipalIdValue", "homeRealmDiscoveryPolicyIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalHomeRealmDiscoveryPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalHomeRealmDiscoveryPolicyId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue",
			Expected: &ServicePrincipalHomeRealmDiscoveryPolicyId{
				ServicePrincipalId:         "servicePrincipalIdValue",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalHomeRealmDiscoveryPolicyID(v.Input)
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

func TestParseServicePrincipalHomeRealmDiscoveryPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalHomeRealmDiscoveryPolicyId
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
			Input: "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/hOmErEaLmDiScOvErYpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue",
			Expected: &ServicePrincipalHomeRealmDiscoveryPolicyId{
				ServicePrincipalId:         "servicePrincipalIdValue",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe",
			Expected: &ServicePrincipalHomeRealmDiscoveryPolicyId{
				ServicePrincipalId:         "sErViCePrInCiPaLiDvAlUe",
				HomeRealmDiscoveryPolicyId: "hOmErEaLmDiScOvErYpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalHomeRealmDiscoveryPolicyIDInsensitively(v.Input)
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

func TestSegmentsForServicePrincipalHomeRealmDiscoveryPolicyId(t *testing.T) {
	segments := ServicePrincipalHomeRealmDiscoveryPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalHomeRealmDiscoveryPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
