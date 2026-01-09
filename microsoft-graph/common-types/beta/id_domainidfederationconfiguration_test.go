package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdFederationConfigurationId{}

func TestNewDomainIdFederationConfigurationID(t *testing.T) {
	id := NewDomainIdFederationConfigurationID("domainId", "internalDomainFederationId")

	if id.DomainId != "domainId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainId")
	}

	if id.InternalDomainFederationId != "internalDomainFederationId" {
		t.Fatalf("Expected %q but got %q for Segment 'InternalDomainFederationId'", id.InternalDomainFederationId, "internalDomainFederationId")
	}
}

func TestFormatDomainIdFederationConfigurationID(t *testing.T) {
	actual := NewDomainIdFederationConfigurationID("domainId", "internalDomainFederationId").ID()
	expected := "/domains/domainId/federationConfiguration/internalDomainFederationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainIdFederationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdFederationConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains/domainId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains/domainId/federationConfiguration",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/federationConfiguration/internalDomainFederationId",
			Expected: &DomainIdFederationConfigurationId{
				DomainId:                   "domainId",
				InternalDomainFederationId: "internalDomainFederationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/federationConfiguration/internalDomainFederationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdFederationConfigurationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DomainId != v.Expected.DomainId {
			t.Fatalf("Expected %q but got %q for DomainId", v.Expected.DomainId, actual.DomainId)
		}

		if actual.InternalDomainFederationId != v.Expected.InternalDomainFederationId {
			t.Fatalf("Expected %q but got %q for InternalDomainFederationId", v.Expected.InternalDomainFederationId, actual.InternalDomainFederationId)
		}

	}
}

func TestParseDomainIdFederationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdFederationConfigurationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains/domainId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains/domainId/federationConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/fEdErAtIoNcOnFiGuRaTiOn",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/federationConfiguration/internalDomainFederationId",
			Expected: &DomainIdFederationConfigurationId{
				DomainId:                   "domainId",
				InternalDomainFederationId: "internalDomainFederationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/federationConfiguration/internalDomainFederationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/fEdErAtIoNcOnFiGuRaTiOn/iNtErNaLdOmAiNfEdErAtIoNiD",
			Expected: &DomainIdFederationConfigurationId{
				DomainId:                   "dOmAiNiD",
				InternalDomainFederationId: "iNtErNaLdOmAiNfEdErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/fEdErAtIoNcOnFiGuRaTiOn/iNtErNaLdOmAiNfEdErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdFederationConfigurationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DomainId != v.Expected.DomainId {
			t.Fatalf("Expected %q but got %q for DomainId", v.Expected.DomainId, actual.DomainId)
		}

		if actual.InternalDomainFederationId != v.Expected.InternalDomainFederationId {
			t.Fatalf("Expected %q but got %q for InternalDomainFederationId", v.Expected.InternalDomainFederationId, actual.InternalDomainFederationId)
		}

	}
}

func TestSegmentsForDomainIdFederationConfigurationId(t *testing.T) {
	segments := DomainIdFederationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainIdFederationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
