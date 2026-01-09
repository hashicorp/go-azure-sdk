package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdSharedEmailDomainInvitationId{}

func TestNewDomainIdSharedEmailDomainInvitationID(t *testing.T) {
	id := NewDomainIdSharedEmailDomainInvitationID("domainId", "sharedEmailDomainInvitationId")

	if id.DomainId != "domainId" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainId")
	}

	if id.SharedEmailDomainInvitationId != "sharedEmailDomainInvitationId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedEmailDomainInvitationId'", id.SharedEmailDomainInvitationId, "sharedEmailDomainInvitationId")
	}
}

func TestFormatDomainIdSharedEmailDomainInvitationID(t *testing.T) {
	actual := NewDomainIdSharedEmailDomainInvitationID("domainId", "sharedEmailDomainInvitationId").ID()
	expected := "/domains/domainId/sharedEmailDomainInvitations/sharedEmailDomainInvitationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainIdSharedEmailDomainInvitationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdSharedEmailDomainInvitationId
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
			Input: "/domains/domainId/sharedEmailDomainInvitations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/sharedEmailDomainInvitations/sharedEmailDomainInvitationId",
			Expected: &DomainIdSharedEmailDomainInvitationId{
				DomainId:                      "domainId",
				SharedEmailDomainInvitationId: "sharedEmailDomainInvitationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/sharedEmailDomainInvitations/sharedEmailDomainInvitationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdSharedEmailDomainInvitationID(v.Input)
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

		if actual.SharedEmailDomainInvitationId != v.Expected.SharedEmailDomainInvitationId {
			t.Fatalf("Expected %q but got %q for SharedEmailDomainInvitationId", v.Expected.SharedEmailDomainInvitationId, actual.SharedEmailDomainInvitationId)
		}

	}
}

func TestParseDomainIdSharedEmailDomainInvitationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainIdSharedEmailDomainInvitationId
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
			Input: "/domains/domainId/sharedEmailDomainInvitations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/sHaReDeMaIlDoMaInInViTaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainId/sharedEmailDomainInvitations/sharedEmailDomainInvitationId",
			Expected: &DomainIdSharedEmailDomainInvitationId{
				DomainId:                      "domainId",
				SharedEmailDomainInvitationId: "sharedEmailDomainInvitationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainId/sharedEmailDomainInvitations/sharedEmailDomainInvitationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/sHaReDeMaIlDoMaInInViTaTiOnS/sHaReDeMaIlDoMaInInViTaTiOnId",
			Expected: &DomainIdSharedEmailDomainInvitationId{
				DomainId:                      "dOmAiNiD",
				SharedEmailDomainInvitationId: "sHaReDeMaIlDoMaInInViTaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiD/sHaReDeMaIlDoMaInInViTaTiOnS/sHaReDeMaIlDoMaInInViTaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainIdSharedEmailDomainInvitationIDInsensitively(v.Input)
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

		if actual.SharedEmailDomainInvitationId != v.Expected.SharedEmailDomainInvitationId {
			t.Fatalf("Expected %q but got %q for SharedEmailDomainInvitationId", v.Expected.SharedEmailDomainInvitationId, actual.SharedEmailDomainInvitationId)
		}

	}
}

func TestSegmentsForDomainIdSharedEmailDomainInvitationId(t *testing.T) {
	segments := DomainIdSharedEmailDomainInvitationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainIdSharedEmailDomainInvitationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
