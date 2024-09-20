package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdVerificationDnsRecordId{}

func TestNewDomainIdVerificationDnsRecordID(t *testing.T) {
	id := NewDomainIdVerificationDnsRecordID("domainId", "domainDnsRecordId")

	if id.DomainId != "domainId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainId")
	}

	if id.DomainDnsRecordId != "domainDnsRecordId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainDnsRecordId'", id.DomainDnsRecordId, "domainDnsRecordId")
	}
}

func TestFormatDomainIdVerificationDnsRecordID(t *testing.T) {
	actual := NewDomainIdVerificationDnsRecordID("domainId", "domainDnsRecordId").ID()
	expected := "/domains/domainId/verificationDnsRecords/domainDnsRecordId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainIdVerificationDnsRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdVerificationDnsRecordId
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
			Input: "/domains/domainId/verificationDnsRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/verificationDnsRecords/domainDnsRecordId",
			Expected: &DomainIdVerificationDnsRecordId{
				DomainId:          "domainId",
				DomainDnsRecordId: "domainDnsRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/verificationDnsRecords/domainDnsRecordId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdVerificationDnsRecordID(v.Input)
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

func TestParseDomainIdVerificationDnsRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdVerificationDnsRecordId
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
			Input: "/domains/domainId/verificationDnsRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/vErIfIcAtIoNdNsReCoRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/verificationDnsRecords/domainDnsRecordId",
			Expected: &DomainIdVerificationDnsRecordId{
				DomainId:          "domainId",
				DomainDnsRecordId: "domainDnsRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/verificationDnsRecords/domainDnsRecordId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/vErIfIcAtIoNdNsReCoRdS/dOmAiNdNsReCoRdId",
			Expected: &DomainIdVerificationDnsRecordId{
				DomainId:          "dOmAiNiD",
				DomainDnsRecordId: "dOmAiNdNsReCoRdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/vErIfIcAtIoNdNsReCoRdS/dOmAiNdNsReCoRdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdVerificationDnsRecordIDInsensitively(v.Input)
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

func TestSegmentsForDomainIdVerificationDnsRecordId(t *testing.T) {
	segments := DomainIdVerificationDnsRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainIdVerificationDnsRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
