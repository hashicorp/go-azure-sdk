package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{}

func TestNewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(t *testing.T) {
	id := NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID("certificateBasedAuthPkiId", "certificateAuthorityDetailId")

	if id.CertificateBasedAuthPkiId != "certificateBasedAuthPkiId" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateBasedAuthPkiId'", id.CertificateBasedAuthPkiId, "certificateBasedAuthPkiId")
	}

	if id.CertificateAuthorityDetailId != "certificateAuthorityDetailId" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateAuthorityDetailId'", id.CertificateAuthorityDetailId, "certificateAuthorityDetailId")
	}
}

func TestFormatDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(t *testing.T) {
	actual := NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID("certificateBasedAuthPkiId", "certificateAuthorityDetailId").ID()
	expected := "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities/certificateAuthorityDetailId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId
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
			// Incomplete URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities/certificateAuthorityDetailId",
			Expected: &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{
				CertificateBasedAuthPkiId:    "certificateBasedAuthPkiId",
				CertificateAuthorityDetailId: "certificateAuthorityDetailId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities/certificateAuthorityDetailId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(v.Input)
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

		if actual.CertificateAuthorityDetailId != v.Expected.CertificateAuthorityDetailId {
			t.Fatalf("Expected %q but got %q for CertificateAuthorityDetailId", v.Expected.CertificateAuthorityDetailId, actual.CertificateAuthorityDetailId)
		}

	}
}

func TestParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId
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
			// Incomplete URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS/cErTiFiCaTeBaSeDaUtHpKiId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS/cErTiFiCaTeBaSeDaUtHpKiId/cErTiFiCaTeAuThOrItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities/certificateAuthorityDetailId",
			Expected: &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{
				CertificateBasedAuthPkiId:    "certificateBasedAuthPkiId",
				CertificateAuthorityDetailId: "certificateAuthorityDetailId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/certificateBasedAuthPkiId/certificateAuthorities/certificateAuthorityDetailId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS/cErTiFiCaTeBaSeDaUtHpKiId/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeAuThOrItYdEtAiLiD",
			Expected: &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{
				CertificateBasedAuthPkiId:    "cErTiFiCaTeBaSeDaUtHpKiId",
				CertificateAuthorityDetailId: "cErTiFiCaTeAuThOrItYdEtAiLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pUbLiCkEyInFrAsTrUcTuRe/cErTiFiCaTeBaSeDaUtHcOnFiGuRaTiOnS/cErTiFiCaTeBaSeDaUtHpKiId/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeAuThOrItYdEtAiLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityIDInsensitively(v.Input)
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

		if actual.CertificateAuthorityDetailId != v.Expected.CertificateAuthorityDetailId {
			t.Fatalf("Expected %q but got %q for CertificateAuthorityDetailId", v.Expected.CertificateAuthorityDetailId, actual.CertificateAuthorityDetailId)
		}

	}
}

func TestSegmentsForDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId(t *testing.T) {
	segments := DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
