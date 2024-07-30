package federatedidentitycredential

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdFederatedIdentityCredentialId{}

func TestNewServicePrincipalIdFederatedIdentityCredentialID(t *testing.T) {
	id := NewServicePrincipalIdFederatedIdentityCredentialID("servicePrincipalIdValue", "federatedIdentityCredentialIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.FederatedIdentityCredentialId != "federatedIdentityCredentialIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FederatedIdentityCredentialId'", id.FederatedIdentityCredentialId, "federatedIdentityCredentialIdValue")
	}
}

func TestFormatServicePrincipalIdFederatedIdentityCredentialID(t *testing.T) {
	actual := NewServicePrincipalIdFederatedIdentityCredentialID("servicePrincipalIdValue", "federatedIdentityCredentialIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue"
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue",
			Expected: &ServicePrincipalIdFederatedIdentityCredentialId{
				ServicePrincipalId:            "servicePrincipalIdValue",
				FederatedIdentityCredentialId: "federatedIdentityCredentialIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue/extra",
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/fEdErAtEdIdEnTiTyCrEdEnTiAlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue",
			Expected: &ServicePrincipalIdFederatedIdentityCredentialId{
				ServicePrincipalId:            "servicePrincipalIdValue",
				FederatedIdentityCredentialId: "federatedIdentityCredentialIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/federatedIdentityCredentials/federatedIdentityCredentialIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlIdVaLuE",
			Expected: &ServicePrincipalIdFederatedIdentityCredentialId{
				ServicePrincipalId:            "sErViCePrInCiPaLiDvAlUe",
				FederatedIdentityCredentialId: "fEdErAtEdIdEnTiTyCrEdEnTiAlIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/fEdErAtEdIdEnTiTyCrEdEnTiAlS/fEdErAtEdIdEnTiTyCrEdEnTiAlIdVaLuE/extra",
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
