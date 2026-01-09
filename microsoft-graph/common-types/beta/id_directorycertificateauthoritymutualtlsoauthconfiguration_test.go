package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{}

func TestNewDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(t *testing.T) {
	id := NewDirectoryCertificateAuthorityMutualTlsOauthConfigurationID("mutualTlsOauthConfigurationId")

	if id.MutualTlsOauthConfigurationId != "mutualTlsOauthConfigurationId" {
		t.Fatalf("Expected %q but got %q for Segment 'MutualTlsOauthConfigurationId'", id.MutualTlsOauthConfigurationId, "mutualTlsOauthConfigurationId")
	}
}

func TestFormatDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(t *testing.T) {
	actual := NewDirectoryCertificateAuthorityMutualTlsOauthConfigurationID("mutualTlsOauthConfigurationId").ID()
	expected := "/directory/certificateAuthorities/mutualTlsOauthConfigurations/mutualTlsOauthConfigurationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCertificateAuthorityMutualTlsOauthConfigurationId
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
			Input: "/directory/certificateAuthorities/mutualTlsOauthConfigurations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/certificateAuthorities/mutualTlsOauthConfigurations/mutualTlsOauthConfigurationId",
			Expected: &DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{
				MutualTlsOauthConfigurationId: "mutualTlsOauthConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/certificateAuthorities/mutualTlsOauthConfigurations/mutualTlsOauthConfigurationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MutualTlsOauthConfigurationId != v.Expected.MutualTlsOauthConfigurationId {
			t.Fatalf("Expected %q but got %q for MutualTlsOauthConfigurationId", v.Expected.MutualTlsOauthConfigurationId, actual.MutualTlsOauthConfigurationId)
		}

	}
}

func TestParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryCertificateAuthorityMutualTlsOauthConfigurationId
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
			Input: "/directory/certificateAuthorities/mutualTlsOauthConfigurations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/mUtUaLtLsOaUtHcOnFiGuRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/certificateAuthorities/mutualTlsOauthConfigurations/mutualTlsOauthConfigurationId",
			Expected: &DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{
				MutualTlsOauthConfigurationId: "mutualTlsOauthConfigurationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/certificateAuthorities/mutualTlsOauthConfigurations/mutualTlsOauthConfigurationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/mUtUaLtLsOaUtHcOnFiGuRaTiOnS/mUtUaLtLsOaUtHcOnFiGuRaTiOnId",
			Expected: &DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{
				MutualTlsOauthConfigurationId: "mUtUaLtLsOaUtHcOnFiGuRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/cErTiFiCaTeAuThOrItIeS/mUtUaLtLsOaUtHcOnFiGuRaTiOnS/mUtUaLtLsOaUtHcOnFiGuRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MutualTlsOauthConfigurationId != v.Expected.MutualTlsOauthConfigurationId {
			t.Fatalf("Expected %q but got %q for MutualTlsOauthConfigurationId", v.Expected.MutualTlsOauthConfigurationId, actual.MutualTlsOauthConfigurationId)
		}

	}
}

func TestSegmentsForDirectoryCertificateAuthorityMutualTlsOauthConfigurationId(t *testing.T) {
	segments := DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryCertificateAuthorityMutualTlsOauthConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
