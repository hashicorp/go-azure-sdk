package domainverificationdnsrecord

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainVerificationDnsRecordId{}

func TestNewDomainVerificationDnsRecordID(t *testing.T) {
	id := NewDomainVerificationDnsRecordID("domainIdValue", "domainDnsRecordIdValue")

	if id.DomainId != "domainIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainIdValue")
	}

	if id.DomainDnsRecordId != "domainDnsRecordIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainDnsRecordId'", id.DomainDnsRecordId, "domainDnsRecordIdValue")
	}
}

func TestFormatDomainVerificationDnsRecordID(t *testing.T) {
	actual := NewDomainVerificationDnsRecordID("domainIdValue", "domainDnsRecordIdValue").ID()
	expected := "/domains/domainIdValue/verificationDnsRecords/domainDnsRecordIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainVerificationDnsRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainVerificationDnsRecordId
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
			Input: "/domains/domainIdValue/verificationDnsRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainIdValue/verificationDnsRecords/domainDnsRecordIdValue",
			Expected: &DomainVerificationDnsRecordId{
				DomainId:          "domainIdValue",
				DomainDnsRecordId: "domainDnsRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainIdValue/verificationDnsRecords/domainDnsRecordIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainVerificationDnsRecordID(v.Input)
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

func TestParseDomainVerificationDnsRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainVerificationDnsRecordId
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
			Input: "/domains/domainIdValue/verificationDnsRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/vErIfIcAtIoNdNsReCoRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainIdValue/verificationDnsRecords/domainDnsRecordIdValue",
			Expected: &DomainVerificationDnsRecordId{
				DomainId:          "domainIdValue",
				DomainDnsRecordId: "domainDnsRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainIdValue/verificationDnsRecords/domainDnsRecordIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/vErIfIcAtIoNdNsReCoRdS/dOmAiNdNsReCoRdIdVaLuE",
			Expected: &DomainVerificationDnsRecordId{
				DomainId:          "dOmAiNiDvAlUe",
				DomainDnsRecordId: "dOmAiNdNsReCoRdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/vErIfIcAtIoNdNsReCoRdS/dOmAiNdNsReCoRdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainVerificationDnsRecordIDInsensitively(v.Input)
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

func TestSegmentsForDomainVerificationDnsRecordId(t *testing.T) {
	segments := DomainVerificationDnsRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainVerificationDnsRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
