package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdHomeRealmDiscoveryPolicyId{}

func TestNewApplicationIdHomeRealmDiscoveryPolicyID(t *testing.T) {
	id := NewApplicationIdHomeRealmDiscoveryPolicyID("applicationId", "homeRealmDiscoveryPolicyId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.HomeRealmDiscoveryPolicyId != "homeRealmDiscoveryPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'HomeRealmDiscoveryPolicyId'", id.HomeRealmDiscoveryPolicyId, "homeRealmDiscoveryPolicyId")
	}
}

func TestFormatApplicationIdHomeRealmDiscoveryPolicyID(t *testing.T) {
	actual := NewApplicationIdHomeRealmDiscoveryPolicyID("applicationId", "homeRealmDiscoveryPolicyId").ID()
	expected := "/applications/applicationId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdHomeRealmDiscoveryPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdHomeRealmDiscoveryPolicyId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId",
			Expected: &ApplicationIdHomeRealmDiscoveryPolicyId{
				ApplicationId:              "applicationId",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdHomeRealmDiscoveryPolicyID(v.Input)
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

func TestParseApplicationIdHomeRealmDiscoveryPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdHomeRealmDiscoveryPolicyId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/hOmErEaLmDiScOvErYpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId",
			Expected: &ApplicationIdHomeRealmDiscoveryPolicyId{
				ApplicationId:              "applicationId",
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiD",
			Expected: &ApplicationIdHomeRealmDiscoveryPolicyId{
				ApplicationId:              "aPpLiCaTiOnId",
				HomeRealmDiscoveryPolicyId: "hOmErEaLmDiScOvErYpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdHomeRealmDiscoveryPolicyIDInsensitively(v.Input)
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

func TestSegmentsForApplicationIdHomeRealmDiscoveryPolicyId(t *testing.T) {
	segments := ApplicationIdHomeRealmDiscoveryPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdHomeRealmDiscoveryPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
