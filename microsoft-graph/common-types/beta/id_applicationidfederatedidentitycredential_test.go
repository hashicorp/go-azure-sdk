package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdFederatedIdentityCredentialId{}

func TestNewApplicationIdFederatedIdentityCredentialID(t *testing.T) {
	id := NewApplicationIdFederatedIdentityCredentialID("applicationId", "federatedIdentityCredentialId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.FederatedIdentityCredentialId != "federatedIdentityCredentialId" {
		t.Fatalf("Expected %q but got %q for Segment 'FederatedIdentityCredentialId'", id.FederatedIdentityCredentialId, "federatedIdentityCredentialId")
	}
}

func TestFormatApplicationIdFederatedIdentityCredentialID(t *testing.T) {
	actual := NewApplicationIdFederatedIdentityCredentialID("applicationId", "federatedIdentityCredentialId").ID()
	expected := "/applications/applicationId/federatedIdentityCredentials/federatedIdentityCredentialId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdFederatedIdentityCredentialID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdFederatedIdentityCredentialId
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
			Input: "/applications/applicationId/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/federatedIdentityCredentials/federatedIdentityCredentialId",
			Expected: &ApplicationIdFederatedIdentityCredentialId{
				ApplicationId:                 "applicationId",
				FederatedIdentityCredentialId: "federatedIdentityCredentialId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/federatedIdentityCredentials/federatedIdentityCredentialId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdFederatedIdentityCredentialID(v.Input)
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

func TestParseApplicationIdFederatedIdentityCredentialIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdFederatedIdentityCredentialId
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
			Input: "/applications/applicationId/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/fEdErAtEdIdEnTiTyCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/federatedIdentityCredentials/federatedIdentityCredentialId",
			Expected: &ApplicationIdFederatedIdentityCredentialId{
				ApplicationId:                 "applicationId",
				FederatedIdentityCredentialId: "federatedIdentityCredentialId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/federatedIdentityCredentials/federatedIdentityCredentialId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlId",
			Expected: &ApplicationIdFederatedIdentityCredentialId{
				ApplicationId:                 "aPpLiCaTiOnId",
				FederatedIdentityCredentialId: "fEdErAtEdIdEnTiTyCrEdEnTiAlId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdFederatedIdentityCredentialIDInsensitively(v.Input)
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

func TestSegmentsForApplicationIdFederatedIdentityCredentialId(t *testing.T) {
	segments := ApplicationIdFederatedIdentityCredentialId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdFederatedIdentityCredentialId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
