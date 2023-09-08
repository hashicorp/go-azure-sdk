package directorycertificateauthoritycertificatebasedapplicationconfigurationtrustedcertificateauthority

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{}

func TestNewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(t *testing.T) {
	id := NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID("certificateBasedApplicationConfigurationIdValue", "certificateAuthorityAsEntityIdValue")

	if id.CertificateBasedApplicationConfigurationId != "certificateBasedApplicationConfigurationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateBasedApplicationConfigurationId'", id.CertificateBasedApplicationConfigurationId, "certificateBasedApplicationConfigurationIdValue")
	}

	if id.CertificateAuthorityAsEntityId != "certificateAuthorityAsEntityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateAuthorityAsEntityId'", id.CertificateAuthorityAsEntityId, "certificateAuthorityAsEntityIdValue")
	}
}

func TestFormatDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(t *testing.T) {
	actual := NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID("certificateBasedApplicationConfigurationIdValue", "certificateAuthorityAsEntityIdValue").ID()
	expected := "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities/certificateAuthorityAsEntityIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId
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
			Input: "/directory/certificateAuthorities",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities/certificateAuthorityAsEntityIdValue",
			Expected: &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{
				CertificateBasedApplicationConfigurationId: "certificateBasedApplicationConfigurationIdValue",
				CertificateAuthorityAsEntityId:             "certificateAuthorityAsEntityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities/certificateAuthorityAsEntityIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CertificateBasedApplicationConfigurationId != v.Expected.CertificateBasedApplicationConfigurationId {
			t.Fatalf("Expected %q but got %q for CertificateBasedApplicationConfigurationId", v.Expected.CertificateBasedApplicationConfigurationId, actual.CertificateBasedApplicationConfigurationId)
		}

		if actual.CertificateAuthorityAsEntityId != v.Expected.CertificateAuthorityAsEntityId {
			t.Fatalf("Expected %q but got %q for CertificateAuthorityAsEntityId", v.Expected.CertificateAuthorityAsEntityId, actual.CertificateAuthorityAsEntityId)
		}

	}
}

func TestParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId
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
			Input: "/directory/certificateAuthorities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiDvAlUe/tRuStEdCeRtIfIcAtEaUtHoRiTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities/certificateAuthorityAsEntityIdValue",
			Expected: &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{
				CertificateBasedApplicationConfigurationId: "certificateBasedApplicationConfigurationIdValue",
				CertificateAuthorityAsEntityId:             "certificateAuthorityAsEntityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationIdValue/trustedCertificateAuthorities/certificateAuthorityAsEntityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiDvAlUe/tRuStEdCeRtIfIcAtEaUtHoRiTiEs/cErTiFiCaTeAuThOrItYaSeNtItYiDvAlUe",
			Expected: &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{
				CertificateBasedApplicationConfigurationId: "cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiDvAlUe",
				CertificateAuthorityAsEntityId:             "cErTiFiCaTeAuThOrItYaSeNtItYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiDvAlUe/tRuStEdCeRtIfIcAtEaUtHoRiTiEs/cErTiFiCaTeAuThOrItYaSeNtItYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CertificateBasedApplicationConfigurationId != v.Expected.CertificateBasedApplicationConfigurationId {
			t.Fatalf("Expected %q but got %q for CertificateBasedApplicationConfigurationId", v.Expected.CertificateBasedApplicationConfigurationId, actual.CertificateBasedApplicationConfigurationId)
		}

		if actual.CertificateAuthorityAsEntityId != v.Expected.CertificateAuthorityAsEntityId {
			t.Fatalf("Expected %q but got %q for CertificateAuthorityAsEntityId", v.Expected.CertificateAuthorityAsEntityId, actual.CertificateAuthorityAsEntityId)
		}

	}
}

func TestSegmentsForDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId(t *testing.T) {
	segments := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationTrustedCertificateAuthorityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
