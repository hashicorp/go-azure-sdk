package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryFederationConfigurationId{}

func TestNewDirectoryFederationConfigurationID(t *testing.T) {
	id := NewDirectoryFederationConfigurationID("identityProviderBaseId")

	if id.IdentityProviderBaseId != "identityProviderBaseId" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityProviderBaseId'", id.IdentityProviderBaseId, "identityProviderBaseId")
	}
}

func TestFormatDirectoryFederationConfigurationID(t *testing.T) {
	actual := NewDirectoryFederationConfigurationID("identityProviderBaseId").ID()
	expected := "/directory/federationConfigurations/identityProviderBaseId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryFederationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryFederationConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/federationConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/federationConfigurations/identityProviderBaseId",
			Expected: &DirectoryFederationConfigurationId{
				IdentityProviderBaseId: "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/federationConfigurations/identityProviderBaseId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryFederationConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestParseDirectoryFederationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryFederationConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/federationConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEdErAtIoNcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/federationConfigurations/identityProviderBaseId",
			Expected: &DirectoryFederationConfigurationId{
				IdentityProviderBaseId: "identityProviderBaseId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/federationConfigurations/identityProviderBaseId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEdErAtIoNcOnFiGuRaTiOnS/iDeNtItYpRoViDeRbAsEiD",
			Expected: &DirectoryFederationConfigurationId{
				IdentityProviderBaseId: "iDeNtItYpRoViDeRbAsEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEdErAtIoNcOnFiGuRaTiOnS/iDeNtItYpRoViDeRbAsEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryFederationConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.IdentityProviderBaseId != v.Expected.IdentityProviderBaseId {
			t.Fatalf("Expected %q but got %q for IdentityProviderBaseId", v.Expected.IdentityProviderBaseId, actual.IdentityProviderBaseId)
		}

	}
}

func TestSegmentsForDirectoryFederationConfigurationId(t *testing.T) {
	segments := DirectoryFederationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryFederationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
