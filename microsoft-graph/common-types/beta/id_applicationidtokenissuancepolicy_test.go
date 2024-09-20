package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdTokenIssuancePolicyId{}

func TestNewApplicationIdTokenIssuancePolicyID(t *testing.T) {
	id := NewApplicationIdTokenIssuancePolicyID("applicationId", "tokenIssuancePolicyId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.TokenIssuancePolicyId != "tokenIssuancePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenIssuancePolicyId'", id.TokenIssuancePolicyId, "tokenIssuancePolicyId")
	}
}

func TestFormatApplicationIdTokenIssuancePolicyID(t *testing.T) {
	actual := NewApplicationIdTokenIssuancePolicyID("applicationId", "tokenIssuancePolicyId").ID()
	expected := "/applications/applicationId/tokenIssuancePolicies/tokenIssuancePolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdTokenIssuancePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdTokenIssuancePolicyId
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
			Input: "/applications/applicationId/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/tokenIssuancePolicies/tokenIssuancePolicyId",
			Expected: &ApplicationIdTokenIssuancePolicyId{
				ApplicationId:         "applicationId",
				TokenIssuancePolicyId: "tokenIssuancePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/tokenIssuancePolicies/tokenIssuancePolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdTokenIssuancePolicyID(v.Input)
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

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

	}
}

func TestParseApplicationIdTokenIssuancePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdTokenIssuancePolicyId
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
			Input: "/applications/applicationId/tokenIssuancePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/tOkEnIsSuAnCePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/tokenIssuancePolicies/tokenIssuancePolicyId",
			Expected: &ApplicationIdTokenIssuancePolicyId{
				ApplicationId:         "applicationId",
				TokenIssuancePolicyId: "tokenIssuancePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/tokenIssuancePolicies/tokenIssuancePolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId",
			Expected: &ApplicationIdTokenIssuancePolicyId{
				ApplicationId:         "aPpLiCaTiOnId",
				TokenIssuancePolicyId: "tOkEnIsSuAnCePoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/tOkEnIsSuAnCePoLiCiEs/tOkEnIsSuAnCePoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdTokenIssuancePolicyIDInsensitively(v.Input)
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

		if actual.TokenIssuancePolicyId != v.Expected.TokenIssuancePolicyId {
			t.Fatalf("Expected %q but got %q for TokenIssuancePolicyId", v.Expected.TokenIssuancePolicyId, actual.TokenIssuancePolicyId)
		}

	}
}

func TestSegmentsForApplicationIdTokenIssuancePolicyId(t *testing.T) {
	segments := ApplicationIdTokenIssuancePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdTokenIssuancePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
