package inboundshareduserprofile

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryInboundSharedUserProfileId{}

func TestNewDirectoryInboundSharedUserProfileID(t *testing.T) {
	id := NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

	if id.InboundSharedUserProfileUserId != "inboundSharedUserProfileUserIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InboundSharedUserProfileUserId'", id.InboundSharedUserProfileUserId, "inboundSharedUserProfileUserIdValue")
	}
}

func TestFormatDirectoryInboundSharedUserProfileID(t *testing.T) {
	actual := NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue").ID()
	expected := "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserIdValue"
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
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserIdValue",
			Expected: &DirectoryInboundSharedUserProfileId{
				InboundSharedUserProfileUserId: "inboundSharedUserProfileUserIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserIdValue/extra",
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
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserIdValue",
			Expected: &DirectoryInboundSharedUserProfileId{
				InboundSharedUserProfileUserId: "inboundSharedUserProfileUserIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/inboundSharedUserProfiles/inboundSharedUserProfileUserIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iNbOuNdShArEdUsErPrOfIlEs/iNbOuNdShArEdUsErPrOfIlEuSeRiDvAlUe",
			Expected: &DirectoryInboundSharedUserProfileId{
				InboundSharedUserProfileUserId: "iNbOuNdShArEdUsErPrOfIlEuSeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/iNbOuNdShArEdUsErPrOfIlEs/iNbOuNdShArEdUsErPrOfIlEuSeRiDvAlUe/extra",
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
