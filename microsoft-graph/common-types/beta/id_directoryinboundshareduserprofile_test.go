package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryInboundSharedUserProfileId{}

func TestNewDirectoryInboundSharedUserProfileID(t *testing.T) {
	id := NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId")

	if id.InboundSharedUserProfileUserId != "inboundSharedUserProfileUserId" {
		t.Fatalf("Expected %q but got %q for Segment 'InboundSharedUserProfileUserId'", id.InboundSharedUserProfileUserId, "inboundSharedUserProfileUserId")
	}
}

func TestFormatDirectoryInboundSharedUserProfileID(t *testing.T) {
	actual := NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId").ID()
	expected := "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryInboundSharedUserProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryInboundSharedUserProfileId
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
			Input: "/directory/inboundSharedUserProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserId",
			Expected: &DirectoryInboundSharedUserProfileId{
				InboundSharedUserProfileUserId: "inboundSharedUserProfileUserId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryInboundSharedUserProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InboundSharedUserProfileUserId != v.Expected.InboundSharedUserProfileUserId {
			t.Fatalf("Expected %q but got %q for InboundSharedUserProfileUserId", v.Expected.InboundSharedUserProfileUserId, actual.InboundSharedUserProfileUserId)
		}

	}
}

func TestParseDirectoryInboundSharedUserProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryInboundSharedUserProfileId
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
			Input: "/directory/inboundSharedUserProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iNbOuNdShArEdUsErPrOfIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserId",
			Expected: &DirectoryInboundSharedUserProfileId{
				InboundSharedUserProfileUserId: "inboundSharedUserProfileUserId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iNbOuNdShArEdUsErPrOfIlEs/iNbOuNdShArEdUsErPrOfIlEuSeRiD",
			Expected: &DirectoryInboundSharedUserProfileId{
				InboundSharedUserProfileUserId: "iNbOuNdShArEdUsErPrOfIlEuSeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iNbOuNdShArEdUsErPrOfIlEs/iNbOuNdShArEdUsErPrOfIlEuSeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryInboundSharedUserProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InboundSharedUserProfileUserId != v.Expected.InboundSharedUserProfileUserId {
			t.Fatalf("Expected %q but got %q for InboundSharedUserProfileUserId", v.Expected.InboundSharedUserProfileUserId, actual.InboundSharedUserProfileUserId)
		}

	}
}

func TestSegmentsForDirectoryInboundSharedUserProfileId(t *testing.T) {
	segments := DirectoryInboundSharedUserProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryInboundSharedUserProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
