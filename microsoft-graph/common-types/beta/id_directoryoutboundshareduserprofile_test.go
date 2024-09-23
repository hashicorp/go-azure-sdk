package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOutboundSharedUserProfileId{}

func TestNewDirectoryOutboundSharedUserProfileID(t *testing.T) {
	id := NewDirectoryOutboundSharedUserProfileID("outboundSharedUserProfileUserId")

	if id.OutboundSharedUserProfileUserId != "outboundSharedUserProfileUserId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutboundSharedUserProfileUserId'", id.OutboundSharedUserProfileUserId, "outboundSharedUserProfileUserId")
	}
}

func TestFormatDirectoryOutboundSharedUserProfileID(t *testing.T) {
	actual := NewDirectoryOutboundSharedUserProfileID("outboundSharedUserProfileUserId").ID()
	expected := "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryOutboundSharedUserProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryOutboundSharedUserProfileId
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
			// Valid URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId",
			Expected: &DirectoryOutboundSharedUserProfileId{
				OutboundSharedUserProfileUserId: "outboundSharedUserProfileUserId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryOutboundSharedUserProfileID(v.Input)
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

	}
}

func TestParseDirectoryOutboundSharedUserProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryOutboundSharedUserProfileId
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
			// Valid URI
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId",
			Expected: &DirectoryOutboundSharedUserProfileId{
				OutboundSharedUserProfileUserId: "outboundSharedUserProfileUserId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/outboundSharedUserProfiles/outboundSharedUserProfileUserId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErId",
			Expected: &DirectoryOutboundSharedUserProfileId{
				OutboundSharedUserProfileUserId: "oUtBoUnDsHaReDuSeRpRoFiLeUsErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/oUtBoUnDsHaReDuSeRpRoFiLeS/oUtBoUnDsHaReDuSeRpRoFiLeUsErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryOutboundSharedUserProfileIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDirectoryOutboundSharedUserProfileId(t *testing.T) {
	segments := DirectoryOutboundSharedUserProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryOutboundSharedUserProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
