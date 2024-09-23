package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdDomainNameReferenceId{}

func TestNewDomainIdDomainNameReferenceID(t *testing.T) {
	id := NewDomainIdDomainNameReferenceID("domainId", "directoryObjectId")

	if id.DomainId != "domainId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatDomainIdDomainNameReferenceID(t *testing.T) {
	actual := NewDomainIdDomainNameReferenceID("domainId", "directoryObjectId").ID()
	expected := "/domains/domainId/domainNameReferences/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainIdDomainNameReferenceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdDomainNameReferenceId
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
			Input: "/domains/domainId/domainNameReferences",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/domainNameReferences/directoryObjectId",
			Expected: &DomainIdDomainNameReferenceId{
				DomainId:          "domainId",
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/domainNameReferences/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdDomainNameReferenceID(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseDomainIdDomainNameReferenceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdDomainNameReferenceId
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
			Input: "/domains/domainId/domainNameReferences",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/dOmAiNnAmErEfErEnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/domainNameReferences/directoryObjectId",
			Expected: &DomainIdDomainNameReferenceId{
				DomainId:          "domainId",
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/domainNameReferences/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/dOmAiNnAmErEfErEnCeS/dIrEcToRyObJeCtId",
			Expected: &DomainIdDomainNameReferenceId{
				DomainId:          "dOmAiNiD",
				DirectoryObjectId: "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/dOmAiNnAmErEfErEnCeS/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdDomainNameReferenceIDInsensitively(v.Input)
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

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForDomainIdDomainNameReferenceId(t *testing.T) {
	segments := DomainIdDomainNameReferenceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainIdDomainNameReferenceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
