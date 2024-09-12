package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOutboundSharedUserProfileIdTenantId{}

func TestNewDirectoryOutboundSharedUserProfileIdTenantID(t *testing.T) {
	id := NewDirectoryOutboundSharedUserProfileIdTenantID("outboundSharedUserProfileUserIdValue", "tenantReferenceTenantIdValue")

	if id.OutboundSharedUserProfileUserId != "outboundSharedUserProfileUserIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutboundSharedUserProfileUserId'", id.OutboundSharedUserProfileUserId, "outboundSharedUserProfileUserIdValue")
	}

	if id.TenantReferenceTenantId != "tenantReferenceTenantIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TenantReferenceTenantId'", id.TenantReferenceTenantId, "tenantReferenceTenantIdValue")
	}
}

func TestFormatDirectoryOutboundSharedUserProfileIdTenantID(t *testing.T) {
	actual := NewDirectoryOutboundSharedUserProfileIdTenantID("outboundSharedUserProfileUserIdValue", "tenantReferenceTenantIdValue").ID()
	expected := "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants/tenantReferenceTenantIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryOutboundSharedUserProfileIdTenantID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryOutboundSharedUserProfileIdTenantId
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
			Input: "/directory/outboundSharedUserProfiles",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants/tenantReferenceTenantIdValue",
			Expected: &DirectoryOutboundSharedUserProfileIdTenantId{
				OutboundSharedUserProfileUserId: "outboundSharedUserProfileUserIdValue",
				TenantReferenceTenantId:         "tenantReferenceTenantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants/tenantReferenceTenantIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryOutboundSharedUserProfileIdTenantID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OutboundSharedUserProfileUserId != v.Expected.OutboundSharedUserProfileUserId {
			t.Fatalf("Expected %q but got %q for OutboundSharedUserProfileUserId", v.Expected.OutboundSharedUserProfileUserId, actual.OutboundSharedUserProfileUserId)
		}

		if actual.TenantReferenceTenantId != v.Expected.TenantReferenceTenantId {
			t.Fatalf("Expected %q but got %q for TenantReferenceTenantId", v.Expected.TenantReferenceTenantId, actual.TenantReferenceTenantId)
		}

	}
}

func TestParseDirectoryOutboundSharedUserProfileIdTenantIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryOutboundSharedUserProfileIdTenantId
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
			Input: "/directory/outboundSharedUserProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErIdVaLuE/tEnAnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants/tenantReferenceTenantIdValue",
			Expected: &DirectoryOutboundSharedUserProfileIdTenantId{
				OutboundSharedUserProfileUserId: "outboundSharedUserProfileUserIdValue",
				TenantReferenceTenantId:         "tenantReferenceTenantIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserIdValue/tenants/tenantReferenceTenantIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErIdVaLuE/tEnAnTs/tEnAnTrEfErEnCeTeNaNtIdVaLuE",
			Expected: &DirectoryOutboundSharedUserProfileIdTenantId{
				OutboundSharedUserProfileUserId: "oUtBoUnDsHaReDuSeRpRoFiLeUsErIdVaLuE",
				TenantReferenceTenantId:         "tEnAnTrEfErEnCeTeNaNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErIdVaLuE/tEnAnTs/tEnAnTrEfErEnCeTeNaNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryOutboundSharedUserProfileIdTenantIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OutboundSharedUserProfileUserId != v.Expected.OutboundSharedUserProfileUserId {
			t.Fatalf("Expected %q but got %q for OutboundSharedUserProfileUserId", v.Expected.OutboundSharedUserProfileUserId, actual.OutboundSharedUserProfileUserId)
		}

		if actual.TenantReferenceTenantId != v.Expected.TenantReferenceTenantId {
			t.Fatalf("Expected %q but got %q for TenantReferenceTenantId", v.Expected.TenantReferenceTenantId, actual.TenantReferenceTenantId)
		}

	}
}

func TestSegmentsForDirectoryOutboundSharedUserProfileIdTenantId(t *testing.T) {
	segments := DirectoryOutboundSharedUserProfileIdTenantId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryOutboundSharedUserProfileIdTenantId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
