package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestId{}

func TestNewIdentityGovernanceAppConsentAppConsentRequestID(t *testing.T) {
	id := NewIdentityGovernanceAppConsentAppConsentRequestID("appConsentRequestIdValue")

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}
}

func TestFormatIdentityGovernanceAppConsentAppConsentRequestID(t *testing.T) {
	actual := NewIdentityGovernanceAppConsentAppConsentRequestID("appConsentRequestIdValue").ID()
	expected := "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAppConsentAppConsentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAppConsentAppConsentRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestId{
				AppConsentRequestId: "appConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAppConsentAppConsentRequestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppConsentRequestId != v.Expected.AppConsentRequestId {
			t.Fatalf("Expected %q but got %q for AppConsentRequestId", v.Expected.AppConsentRequestId, actual.AppConsentRequestId)
		}

	}
}

func TestParseIdentityGovernanceAppConsentAppConsentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAppConsentAppConsentRequestId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestId{
				AppConsentRequestId: "appConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestId{
				AppConsentRequestId: "aPpCoNsEnTrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAppConsentAppConsentRequestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppConsentRequestId != v.Expected.AppConsentRequestId {
			t.Fatalf("Expected %q but got %q for AppConsentRequestId", v.Expected.AppConsentRequestId, actual.AppConsentRequestId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAppConsentAppConsentRequestId(t *testing.T) {
	segments := IdentityGovernanceAppConsentAppConsentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAppConsentAppConsentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
