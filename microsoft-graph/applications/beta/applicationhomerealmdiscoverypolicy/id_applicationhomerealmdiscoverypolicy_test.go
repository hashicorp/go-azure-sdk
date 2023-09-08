package applicationhomerealmdiscoverypolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationHomeRealmDiscoveryPolicyId{}

func TestNewApplicationHomeRealmDiscoveryPolicyID(t *testing.T) {
	id := NewApplicationHomeRealmDiscoveryPolicyID("applicationIdValue", "homeRealmDiscoveryPolicyIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.HomeRealmDiscoveryPolicyId != "homeRealmDiscoveryPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HomeRealmDiscoveryPolicyId'", id.HomeRealmDiscoveryPolicyId, "homeRealmDiscoveryPolicyIdValue")
	}
}

func TestFormatApplicationHomeRealmDiscoveryPolicyID(t *testing.T) {
	actual := NewApplicationHomeRealmDiscoveryPolicyID("applicationIdValue", "homeRealmDiscoveryPolicyIdValue").ID()
	expected := "/applications/applicationIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationHomeRealmDiscoveryPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationHomeRealmDiscoveryPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue",
			Expected: &ApplicationHomeRealmDiscoveryPolicyId{
				ApplicationId:              "applicationIdValue",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationHomeRealmDiscoveryPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

	}
}

func TestParseApplicationHomeRealmDiscoveryPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationHomeRealmDiscoveryPolicyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/hOmErEaLmDiScOvErYpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue",
			Expected: &ApplicationHomeRealmDiscoveryPolicyId{
				ApplicationId:              "applicationIdValue",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe",
			Expected: &ApplicationHomeRealmDiscoveryPolicyId{
				ApplicationId:              "aPpLiCaTiOnIdVaLuE",
				HomeRealmDiscoveryPolicyId: "hOmErEaLmDiScOvErYpOlIcYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationHomeRealmDiscoveryPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

	}
}

func TestSegmentsForApplicationHomeRealmDiscoveryPolicyId(t *testing.T) {
	segments := ApplicationHomeRealmDiscoveryPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationHomeRealmDiscoveryPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
