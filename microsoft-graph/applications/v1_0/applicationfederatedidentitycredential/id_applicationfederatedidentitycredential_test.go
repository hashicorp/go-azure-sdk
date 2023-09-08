package applicationfederatedidentitycredential

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationFederatedIdentityCredentialId{}

func TestNewApplicationFederatedIdentityCredentialID(t *testing.T) {
	id := NewApplicationFederatedIdentityCredentialID("applicationIdValue", "federatedIdentityCredentialIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.FederatedIdentityCredentialId != "federatedIdentityCredentialIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FederatedIdentityCredentialId'", id.FederatedIdentityCredentialId, "federatedIdentityCredentialIdValue")
	}
}

func TestFormatApplicationFederatedIdentityCredentialID(t *testing.T) {
	actual := NewApplicationFederatedIdentityCredentialID("applicationIdValue", "federatedIdentityCredentialIdValue").ID()
	expected := "/applications/applicationIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationFederatedIdentityCredentialID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationFederatedIdentityCredentialId
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
			Input: "/applications/applicationIdValue/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue",
			Expected: &ApplicationFederatedIdentityCredentialId{
				ApplicationId:                 "applicationIdValue",
				FederatedIdentityCredentialId: "federatedIdentityCredentialIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationFederatedIdentityCredentialID(v.Input)
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

		if actual.FederatedIdentityCredentialId != v.Expected.FederatedIdentityCredentialId {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialId", v.Expected.FederatedIdentityCredentialId, actual.FederatedIdentityCredentialId)
		}

	}
}

func TestParseApplicationFederatedIdentityCredentialIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationFederatedIdentityCredentialId
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
			Input: "/applications/applicationIdValue/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/fEdErAtEdIdEnTiTyCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue",
			Expected: &ApplicationFederatedIdentityCredentialId{
				ApplicationId:                 "applicationIdValue",
				FederatedIdentityCredentialId: "federatedIdentityCredentialIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlIdVaLuE",
			Expected: &ApplicationFederatedIdentityCredentialId{
				ApplicationId:                 "aPpLiCaTiOnIdVaLuE",
				FederatedIdentityCredentialId: "fEdErAtEdIdEnTiTyCrEdEnTiAlIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationFederatedIdentityCredentialIDInsensitively(v.Input)
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

		if actual.FederatedIdentityCredentialId != v.Expected.FederatedIdentityCredentialId {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialId", v.Expected.FederatedIdentityCredentialId, actual.FederatedIdentityCredentialId)
		}

	}
}

func TestSegmentsForApplicationFederatedIdentityCredentialId(t *testing.T) {
	segments := ApplicationFederatedIdentityCredentialId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationFederatedIdentityCredentialId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
