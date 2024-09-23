package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{}

func TestNewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(t *testing.T) {
	id := NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID("certificateBasedApplicationConfigurationId")

	if id.CertificateBasedApplicationConfigurationId != "certificateBasedApplicationConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateBasedApplicationConfigurationId'", id.CertificateBasedApplicationConfigurationId, "certificateBasedApplicationConfigurationId")
	}
}

func TestFormatDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(t *testing.T) {
	actual := NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID("certificateBasedApplicationConfigurationId").ID()
	expected := "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId
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
			// Valid URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationId",
			Expected: &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{
				CertificateBasedApplicationConfigurationId: "certificateBasedApplicationConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(v.Input)
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

	}
}

func TestParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId
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
			// Valid URI
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationId",
			Expected: &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{
				CertificateBasedApplicationConfigurationId: "certificateBasedApplicationConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/certificateBasedApplicationConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiD",
			Expected: &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{
				CertificateBasedApplicationConfigurationId: "cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNs/cErTiFiCaTeBaSeDaPpLiCaTiOnCoNfIgUrAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId(t *testing.T) {
	segments := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
