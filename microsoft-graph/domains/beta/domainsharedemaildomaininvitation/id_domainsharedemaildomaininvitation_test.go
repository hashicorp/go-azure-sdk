package domainsharedemaildomaininvitation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DomainSharedEmailDomainInvitationId{}

func TestNewDomainSharedEmailDomainInvitationID(t *testing.T) {
	id := NewDomainSharedEmailDomainInvitationID("domainIdValue", "sharedEmailDomainInvitationIdValue")

	if id.DomainId != "domainIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DomainId'", id.DomainId, "domainIdValue")
	}

	if id.SharedEmailDomainInvitationId != "sharedEmailDomainInvitationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedEmailDomainInvitationId'", id.SharedEmailDomainInvitationId, "sharedEmailDomainInvitationIdValue")
	}
}

func TestFormatDomainSharedEmailDomainInvitationID(t *testing.T) {
	actual := NewDomainSharedEmailDomainInvitationID("domainIdValue", "sharedEmailDomainInvitationIdValue").ID()
	expected := "/domains/domainIdValue/sharedEmailDomainInvitations/sharedEmailDomainInvitationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDomainSharedEmailDomainInvitationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainSharedEmailDomainInvitationId
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
			Input: "/domains/domainIdValue/sharedEmailDomainInvitations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainIdValue/sharedEmailDomainInvitations/sharedEmailDomainInvitationIdValue",
			Expected: &DomainSharedEmailDomainInvitationId{
				DomainId:                      "domainIdValue",
				SharedEmailDomainInvitationId: "sharedEmailDomainInvitationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainIdValue/sharedEmailDomainInvitations/sharedEmailDomainInvitationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainSharedEmailDomainInvitationID(v.Input)
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

func TestParseDomainSharedEmailDomainInvitationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DomainSharedEmailDomainInvitationId
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
			Input: "/domains/domainIdValue/sharedEmailDomainInvitations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/sHaReDeMaIlDoMaInInViTaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/domains/domainIdValue/sharedEmailDomainInvitations/sharedEmailDomainInvitationIdValue",
			Expected: &DomainSharedEmailDomainInvitationId{
				DomainId:                      "domainIdValue",
				SharedEmailDomainInvitationId: "sharedEmailDomainInvitationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/domains/domainIdValue/sharedEmailDomainInvitations/sharedEmailDomainInvitationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/sHaReDeMaIlDoMaInInViTaTiOnS/sHaReDeMaIlDoMaInInViTaTiOnIdVaLuE",
			Expected: &DomainSharedEmailDomainInvitationId{
				DomainId:                      "dOmAiNiDvAlUe",
				SharedEmailDomainInvitationId: "sHaReDeMaIlDoMaInInViTaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dOmAiNs/dOmAiNiDvAlUe/sHaReDeMaIlDoMaInInViTaTiOnS/sHaReDeMaIlDoMaInInViTaTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDomainSharedEmailDomainInvitationIDInsensitively(v.Input)
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

func TestSegmentsForDomainSharedEmailDomainInvitationId(t *testing.T) {
	segments := DomainSharedEmailDomainInvitationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DomainSharedEmailDomainInvitationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
