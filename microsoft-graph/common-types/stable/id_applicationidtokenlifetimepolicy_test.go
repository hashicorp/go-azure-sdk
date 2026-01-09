package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdTokenLifetimePolicyId{}

func TestNewApplicationIdTokenLifetimePolicyID(t *testing.T) {
	id := NewApplicationIdTokenLifetimePolicyID("applicationId", "tokenLifetimePolicyId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.TokenLifetimePolicyId != "tokenLifetimePolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenLifetimePolicyId'", id.TokenLifetimePolicyId, "tokenLifetimePolicyId")
	}
}

func TestFormatApplicationIdTokenLifetimePolicyID(t *testing.T) {
	actual := NewApplicationIdTokenLifetimePolicyID("applicationId", "tokenLifetimePolicyId").ID()
	expected := "/applications/applicationId/tokenLifetimePolicies/tokenLifetimePolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdTokenLifetimePolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdTokenLifetimePolicyId
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
			Input: "/applications/applicationId/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/tokenLifetimePolicies/tokenLifetimePolicyId",
			Expected: &ApplicationIdTokenLifetimePolicyId{
				ApplicationId:         "applicationId",
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/tokenLifetimePolicies/tokenLifetimePolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdTokenLifetimePolicyID(v.Input)
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

		if actual.TokenLifetimePolicyId != v.Expected.TokenLifetimePolicyId {
			t.Fatalf("Expected %q but got %q for TokenLifetimePolicyId", v.Expected.TokenLifetimePolicyId, actual.TokenLifetimePolicyId)
		}

	}
}

func TestParseApplicationIdTokenLifetimePolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdTokenLifetimePolicyId
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
			Input: "/applications/applicationId/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/tOkEnLiFeTiMePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/tokenLifetimePolicies/tokenLifetimePolicyId",
			Expected: &ApplicationIdTokenLifetimePolicyId{
				ApplicationId:         "applicationId",
				TokenLifetimePolicyId: "tokenLifetimePolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/tokenLifetimePolicies/tokenLifetimePolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId",
			Expected: &ApplicationIdTokenLifetimePolicyId{
				ApplicationId:         "aPpLiCaTiOnId",
				TokenLifetimePolicyId: "tOkEnLiFeTiMePoLiCyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdTokenLifetimePolicyIDInsensitively(v.Input)
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

		if actual.TokenLifetimePolicyId != v.Expected.TokenLifetimePolicyId {
			t.Fatalf("Expected %q but got %q for TokenLifetimePolicyId", v.Expected.TokenLifetimePolicyId, actual.TokenLifetimePolicyId)
		}

	}
}

func TestSegmentsForApplicationIdTokenLifetimePolicyId(t *testing.T) {
	segments := ApplicationIdTokenLifetimePolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdTokenLifetimePolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
