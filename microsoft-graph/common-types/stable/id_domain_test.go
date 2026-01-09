package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainId{}

func TestNewDomainID(t *testing.T) {
	id := NewDomainID("domainId")

	if id.DomainId != "domainId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainId")
	}
}

func TestFormatDomainID(t *testing.T) {
	actual := NewDomainID("domainId").ID()
	expected := "/domains/domainId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainId
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
			// Valid URI
			Input: "/domains/domainId",
			Expected: &DomainId{
				DomainId: "domainId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainID(v.Input)
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

	}
}

func TestParseDomainIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainId
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
			// Valid URI
			Input: "/domains/domainId",
			Expected: &DomainId{
				DomainId: "domainId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD",
			Expected: &DomainId{
				DomainId: "dOmAiNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDomainId(t *testing.T) {
	segments := DomainId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
