package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdFederatedIdentityCredentialId{}

func TestNewServicePrincipalIdFederatedIdentityCredentialID(t *testing.T) {
	id := NewServicePrincipalIdFederatedIdentityCredentialID("servicePrincipalId", "federatedIdentityCredentialId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.FederatedIdentityCredentialId != "federatedIdentityCredentialId" {
		t.Fatalf("Expected %q but got %q for Segment 'FederatedIdentityCredentialId'", id.FederatedIdentityCredentialId, "federatedIdentityCredentialId")
	}
}

func TestFormatServicePrincipalIdFederatedIdentityCredentialID(t *testing.T) {
	actual := NewServicePrincipalIdFederatedIdentityCredentialID("servicePrincipalId", "federatedIdentityCredentialId").ID()
	expected := "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials/federatedIdentityCredentialId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdFederatedIdentityCredentialID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdFederatedIdentityCredentialId
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
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials/federatedIdentityCredentialId",
			Expected: &ServicePrincipalIdFederatedIdentityCredentialId{
				ServicePrincipalId:            "servicePrincipalId",
				FederatedIdentityCredentialId: "federatedIdentityCredentialId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials/federatedIdentityCredentialId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdFederatedIdentityCredentialID(v.Input)
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

		if actual.FederatedIdentityCredentialId != v.Expected.FederatedIdentityCredentialId {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialId", v.Expected.FederatedIdentityCredentialId, actual.FederatedIdentityCredentialId)
		}

	}
}

func TestParseServicePrincipalIdFederatedIdentityCredentialIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdFederatedIdentityCredentialId
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
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/fEdErAtEdIdEnTiTyCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials/federatedIdentityCredentialId",
			Expected: &ServicePrincipalIdFederatedIdentityCredentialId{
				ServicePrincipalId:            "servicePrincipalId",
				FederatedIdentityCredentialId: "federatedIdentityCredentialId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/federatedIdentityCredentials/federatedIdentityCredentialId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlId",
			Expected: &ServicePrincipalIdFederatedIdentityCredentialId{
				ServicePrincipalId:            "sErViCePrInCiPaLiD",
				FederatedIdentityCredentialId: "fEdErAtEdIdEnTiTyCrEdEnTiAlId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdFederatedIdentityCredentialIDInsensitively(v.Input)
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

		if actual.FederatedIdentityCredentialId != v.Expected.FederatedIdentityCredentialId {
			t.Fatalf("Expected %q but got %q for FederatedIdentityCredentialId", v.Expected.FederatedIdentityCredentialId, actual.FederatedIdentityCredentialId)
		}

	}
}

func TestSegmentsForServicePrincipalIdFederatedIdentityCredentialId(t *testing.T) {
	segments := ServicePrincipalIdFederatedIdentityCredentialId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdFederatedIdentityCredentialId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
