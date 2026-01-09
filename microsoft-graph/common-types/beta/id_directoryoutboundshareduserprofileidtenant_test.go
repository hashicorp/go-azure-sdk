package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOutboundSharedUserProfileIdTenantId{}

func TestNewDirectoryOutboundSharedUserProfileIdTenantID(t *testing.T) {
	id := NewDirectoryOutboundSharedUserProfileIdTenantID("outboundSharedUserProfileUserId", "tenantReferenceTenantId")

	if id.OutboundSharedUserProfileUserId != "outboundSharedUserProfileUserId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutboundSharedUserProfileUserId'", id.OutboundSharedUserProfileUserId, "outboundSharedUserProfileUserId")
	}

	if id.TenantReferenceTenantId != "tenantReferenceTenantId" {
		t.Fatalf("Expected %q but got %q for Segment 'TenantReferenceTenantId'", id.TenantReferenceTenantId, "tenantReferenceTenantId")
	}
}

func TestFormatDirectoryOutboundSharedUserProfileIdTenantID(t *testing.T) {
	actual := NewDirectoryOutboundSharedUserProfileIdTenantID("outboundSharedUserProfileUserId", "tenantReferenceTenantId").ID()
	expected := "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants/tenantReferenceTenantId"
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
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants/tenantReferenceTenantId",
			Expected: &DirectoryOutboundSharedUserProfileIdTenantId{
				OutboundSharedUserProfileUserId: "outboundSharedUserProfileUserId",
				TenantReferenceTenantId:         "tenantReferenceTenantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants/tenantReferenceTenantId/extra",
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
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErId/tEnAnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants/tenantReferenceTenantId",
			Expected: &DirectoryOutboundSharedUserProfileIdTenantId{
				OutboundSharedUserProfileUserId: "outboundSharedUserProfileUserId",
				TenantReferenceTenantId:         "tenantReferenceTenantId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/tenants/tenantReferenceTenantId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErId/tEnAnTs/tEnAnTrEfErEnCeTeNaNtId",
			Expected: &DirectoryOutboundSharedUserProfileIdTenantId{
				OutboundSharedUserProfileUserId: "oUtBoUnDsHaReDuSeRpRoFiLeUsErId",
				TenantReferenceTenantId:         "tEnAnTrEfErEnCeTeNaNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErId/tEnAnTs/tEnAnTrEfErEnCeTeNaNtId/extra",
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
