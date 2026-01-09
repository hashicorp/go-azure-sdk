package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{}

func TestNewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(t *testing.T) {
	id := NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID("certificateBasedAuthPkiId")

	if id.CertificateBasedAuthPkiId != "certificateBasedAuthPkiId" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateBasedAuthPkiId'", id.CertificateBasedAuthPkiId, "certificateBasedAuthPkiId")
	}
}

func TestFormatDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(t *testing.T) {
	actual := NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID("certificateBasedAuthPkiId").ID()
	expected := "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId
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
			Input: "/directory/publicKeyInfrastructure",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId",
			Expected: &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{
				CertificateBasedAuthPkiId: "certificateBasedAuthPkiId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CertificateBasedAuthPkiId != v.Expected.CertificateBasedAuthPkiId {
			t.Fatalf("Expected %q but got %q for CertificateBasedAuthPkiId", v.Expected.CertificateBasedAuthPkiId, actual.CertificateBasedAuthPkiId)
		}

	}
}

func TestParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId
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
			Input: "/directory/publicKeyInfrastructure",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId",
			Expected: &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{
				CertificateBasedAuthPkiId: "certificateBasedAuthPkiId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS/cErTiFiCaTeBaSeDaUtHpKiId",
			Expected: &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{
				CertificateBasedAuthPkiId: "cErTiFiCaTeBaSeDaUtHpKiId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS/cErTiFiCaTeBaSeDaUtHpKiId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CertificateBasedAuthPkiId != v.Expected.CertificateBasedAuthPkiId {
			t.Fatalf("Expected %q but got %q for CertificateBasedAuthPkiId", v.Expected.CertificateBasedAuthPkiId, actual.CertificateBasedAuthPkiId)
		}

	}
}

func TestSegmentsForDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId(t *testing.T) {
	segments := DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
