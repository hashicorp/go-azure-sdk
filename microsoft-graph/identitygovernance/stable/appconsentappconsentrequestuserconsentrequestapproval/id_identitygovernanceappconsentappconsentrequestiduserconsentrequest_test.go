package appconsentappconsentrequestuserconsentrequestapproval

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{}

func TestNewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(t *testing.T) {
	id := NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID("appConsentRequestIdValue", "userConsentRequestIdValue")

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}

	if id.UserConsentRequestId != "userConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestIdValue")
	}
}

func TestFormatIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(t *testing.T) {
	actual := NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID("appConsentRequestIdValue", "userConsentRequestIdValue").ID()
	expected := "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId
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
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestID(v.Input)
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

		if actual.UserConsentRequestId != v.Expected.UserConsentRequestId {
			t.Fatalf("Expected %q but got %q for UserConsentRequestId", v.Expected.UserConsentRequestId, actual.UserConsentRequestId)
		}

	}
}

func TestParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId
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
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStIdVaLuE",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIDInsensitively(v.Input)
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

		if actual.UserConsentRequestId != v.Expected.UserConsentRequestId {
			t.Fatalf("Expected %q but got %q for UserConsentRequestId", v.Expected.UserConsentRequestId, actual.UserConsentRequestId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId(t *testing.T) {
	segments := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
