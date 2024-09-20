package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryExternalUserProfileId{}

func TestNewDirectoryExternalUserProfileID(t *testing.T) {
	id := NewDirectoryExternalUserProfileID("externalUserProfileId")

	if id.ExternalUserProfileId != "externalUserProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExternalUserProfileId'", id.ExternalUserProfileId, "externalUserProfileId")
	}
}

func TestFormatDirectoryExternalUserProfileID(t *testing.T) {
	actual := NewDirectoryExternalUserProfileID("externalUserProfileId").ID()
	expected := "/directory/externalUserProfiles/externalUserProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryExternalUserProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryExternalUserProfileId
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
			Input: "/directory/externalUserProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/externalUserProfiles/externalUserProfileId",
			Expected: &DirectoryExternalUserProfileId{
				ExternalUserProfileId: "externalUserProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/externalUserProfiles/externalUserProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryExternalUserProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ExternalUserProfileId != v.Expected.ExternalUserProfileId {
			t.Fatalf("Expected %q but got %q for ExternalUserProfileId", v.Expected.ExternalUserProfileId, actual.ExternalUserProfileId)
		}

	}
}

func TestParseDirectoryExternalUserProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryExternalUserProfileId
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
			Input: "/directory/externalUserProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/eXtErNaLuSeRpRoFiLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/externalUserProfiles/externalUserProfileId",
			Expected: &DirectoryExternalUserProfileId{
				ExternalUserProfileId: "externalUserProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/externalUserProfiles/externalUserProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/eXtErNaLuSeRpRoFiLeS/eXtErNaLuSeRpRoFiLeId",
			Expected: &DirectoryExternalUserProfileId{
				ExternalUserProfileId: "eXtErNaLuSeRpRoFiLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/eXtErNaLuSeRpRoFiLeS/eXtErNaLuSeRpRoFiLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryExternalUserProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ExternalUserProfileId != v.Expected.ExternalUserProfileId {
			t.Fatalf("Expected %q but got %q for ExternalUserProfileId", v.Expected.ExternalUserProfileId, actual.ExternalUserProfileId)
		}

	}
}

func TestSegmentsForDirectoryExternalUserProfileId(t *testing.T) {
	segments := DirectoryExternalUserProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryExternalUserProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
