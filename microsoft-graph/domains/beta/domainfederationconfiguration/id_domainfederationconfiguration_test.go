package domainfederationconfiguration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainFederationConfigurationId{}

func TestNewDomainFederationConfigurationID(t *testing.T) {
	id := NewDomainFederationConfigurationID("domainIdValue", "internalDomainFederationIdValue")

	if id.DomainId != "domainIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainIdValue")
	}

	if id.InternalDomainFederationId != "internalDomainFederationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InternalDomainFederationId'", id.InternalDomainFederationId, "internalDomainFederationIdValue")
	}
}

func TestFormatDomainFederationConfigurationID(t *testing.T) {
	actual := NewDomainFederationConfigurationID("domainIdValue", "internalDomainFederationIdValue").ID()
	expected := "/domains/domainIdValue/federationConfiguration/internalDomainFederationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainFederationConfigurationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainFederationConfigurationId
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
			Input: "/domains/domainIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains/domainIdValue/federationConfiguration",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainIdValue/federationConfiguration/internalDomainFederationIdValue",
			Expected: &DomainFederationConfigurationId{
				DomainId:                   "domainIdValue",
				InternalDomainFederationId: "internalDomainFederationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainIdValue/federationConfiguration/internalDomainFederationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainFederationConfigurationID(v.Input)
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

func TestParseDomainFederationConfigurationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainFederationConfigurationId
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
			Input: "/domains/domainIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/domains/domainIdValue/federationConfiguration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/fEdErAtIoNcOnFiGuRaTiOn",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainIdValue/federationConfiguration/internalDomainFederationIdValue",
			Expected: &DomainFederationConfigurationId{
				DomainId:                   "domainIdValue",
				InternalDomainFederationId: "internalDomainFederationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainIdValue/federationConfiguration/internalDomainFederationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/fEdErAtIoNcOnFiGuRaTiOn/iNtErNaLdOmAiNfEdErAtIoNiDvAlUe",
			Expected: &DomainFederationConfigurationId{
				DomainId:                   "dOmAiNiDvAlUe",
				InternalDomainFederationId: "iNtErNaLdOmAiNfEdErAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/fEdErAtIoNcOnFiGuRaTiOn/iNtErNaLdOmAiNfEdErAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainFederationConfigurationIDInsensitively(v.Input)
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

func TestSegmentsForDomainFederationConfigurationId(t *testing.T) {
	segments := DomainFederationConfigurationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainFederationConfigurationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
