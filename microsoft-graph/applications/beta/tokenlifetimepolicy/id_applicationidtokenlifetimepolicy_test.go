package tokenlifetimepolicy

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdTokenLifetimePolicyId{}

func TestNewApplicationIdTokenLifetimePolicyID(t *testing.T) {
	id := NewApplicationIdTokenLifetimePolicyID("applicationIdValue", "tokenLifetimePolicyIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.TokenLifetimePolicyId != "tokenLifetimePolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenLifetimePolicyId'", id.TokenLifetimePolicyId, "tokenLifetimePolicyIdValue")
	}
}

func TestFormatApplicationIdTokenLifetimePolicyID(t *testing.T) {
	actual := NewApplicationIdTokenLifetimePolicyID("applicationIdValue", "tokenLifetimePolicyIdValue").ID()
	expected := "/applications/applicationIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue"
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
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue",
			Expected: &ApplicationIdTokenLifetimePolicyId{
				ApplicationId:         "applicationIdValue",
				TokenLifetimePolicyId: "tokenLifetimePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue/extra",
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
			Input: "/applications/applicationIdValue/tokenLifetimePolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/tOkEnLiFeTiMePoLiCiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue",
			Expected: &ApplicationIdTokenLifetimePolicyId{
				ApplicationId:         "applicationIdValue",
				TokenLifetimePolicyId: "tokenLifetimePolicyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/tokenLifetimePolicies/tokenLifetimePolicyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyIdVaLuE",
			Expected: &ApplicationIdTokenLifetimePolicyId{
				ApplicationId:         "aPpLiCaTiOnIdVaLuE",
				TokenLifetimePolicyId: "tOkEnLiFeTiMePoLiCyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/tOkEnLiFeTiMePoLiCiEs/tOkEnLiFeTiMePoLiCyIdVaLuE/extra",
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
