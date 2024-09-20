package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdServiceConfigurationRecordId{}

func TestNewDomainIdServiceConfigurationRecordID(t *testing.T) {
	id := NewDomainIdServiceConfigurationRecordID("domainId", "domainDnsRecordId")

	if id.DomainId != "domainId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainId")
	}

	if id.DomainDnsRecordId != "domainDnsRecordId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainDnsRecordId'", id.DomainDnsRecordId, "domainDnsRecordId")
	}
}

func TestFormatDomainIdServiceConfigurationRecordID(t *testing.T) {
	actual := NewDomainIdServiceConfigurationRecordID("domainId", "domainDnsRecordId").ID()
	expected := "/domains/domainId/serviceConfigurationRecords/domainDnsRecordId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainIdServiceConfigurationRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdServiceConfigurationRecordId
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
			Input: "/domains/domainId/serviceConfigurationRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/serviceConfigurationRecords/domainDnsRecordId",
			Expected: &DomainIdServiceConfigurationRecordId{
				DomainId:          "domainId",
				DomainDnsRecordId: "domainDnsRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/serviceConfigurationRecords/domainDnsRecordId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdServiceConfigurationRecordID(v.Input)
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

		if actual.DomainDnsRecordId != v.Expected.DomainDnsRecordId {
			t.Fatalf("Expected %q but got %q for DomainDnsRecordId", v.Expected.DomainDnsRecordId, actual.DomainDnsRecordId)
		}

	}
}

func TestParseDomainIdServiceConfigurationRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdServiceConfigurationRecordId
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
			Input: "/domains/domainId/serviceConfigurationRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/sErViCeCoNfIgUrAtIoNrEcOrDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/serviceConfigurationRecords/domainDnsRecordId",
			Expected: &DomainIdServiceConfigurationRecordId{
				DomainId:          "domainId",
				DomainDnsRecordId: "domainDnsRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/serviceConfigurationRecords/domainDnsRecordId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/sErViCeCoNfIgUrAtIoNrEcOrDs/dOmAiNdNsReCoRdId",
			Expected: &DomainIdServiceConfigurationRecordId{
				DomainId:          "dOmAiNiD",
				DomainDnsRecordId: "dOmAiNdNsReCoRdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/sErViCeCoNfIgUrAtIoNrEcOrDs/dOmAiNdNsReCoRdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdServiceConfigurationRecordIDInsensitively(v.Input)
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

		if actual.DomainDnsRecordId != v.Expected.DomainDnsRecordId {
			t.Fatalf("Expected %q but got %q for DomainDnsRecordId", v.Expected.DomainDnsRecordId, actual.DomainDnsRecordId)
		}

	}
}

func TestSegmentsForDomainIdServiceConfigurationRecordId(t *testing.T) {
	segments := DomainIdServiceConfigurationRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainIdServiceConfigurationRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
