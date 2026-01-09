package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryPendingExternalUserProfileId{}

func TestNewDirectoryPendingExternalUserProfileID(t *testing.T) {
	id := NewDirectoryPendingExternalUserProfileID("pendingExternalUserProfileId")

	if id.PendingExternalUserProfileId != "pendingExternalUserProfileId" {
		t.Fatalf("Expected %q but got %q for Segment 'PendingExternalUserProfileId'", id.PendingExternalUserProfileId, "pendingExternalUserProfileId")
	}
}

func TestFormatDirectoryPendingExternalUserProfileID(t *testing.T) {
	actual := NewDirectoryPendingExternalUserProfileID("pendingExternalUserProfileId").ID()
	expected := "/directory/pendingExternalUserProfiles/pendingExternalUserProfileId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryPendingExternalUserProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryPendingExternalUserProfileId
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
			Input: "/directory/pendingExternalUserProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/pendingExternalUserProfiles/pendingExternalUserProfileId",
			Expected: &DirectoryPendingExternalUserProfileId{
				PendingExternalUserProfileId: "pendingExternalUserProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/pendingExternalUserProfiles/pendingExternalUserProfileId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryPendingExternalUserProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PendingExternalUserProfileId != v.Expected.PendingExternalUserProfileId {
			t.Fatalf("Expected %q but got %q for PendingExternalUserProfileId", v.Expected.PendingExternalUserProfileId, actual.PendingExternalUserProfileId)
		}

	}
}

func TestParseDirectoryPendingExternalUserProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryPendingExternalUserProfileId
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
			Input: "/directory/pendingExternalUserProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pEnDiNgExTeRnAlUsErPrOfIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/pendingExternalUserProfiles/pendingExternalUserProfileId",
			Expected: &DirectoryPendingExternalUserProfileId{
				PendingExternalUserProfileId: "pendingExternalUserProfileId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/pendingExternalUserProfiles/pendingExternalUserProfileId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pEnDiNgExTeRnAlUsErPrOfIlEs/pEnDiNgExTeRnAlUsErPrOfIlEiD",
			Expected: &DirectoryPendingExternalUserProfileId{
				PendingExternalUserProfileId: "pEnDiNgExTeRnAlUsErPrOfIlEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/pEnDiNgExTeRnAlUsErPrOfIlEs/pEnDiNgExTeRnAlUsErPrOfIlEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryPendingExternalUserProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PendingExternalUserProfileId != v.Expected.PendingExternalUserProfileId {
			t.Fatalf("Expected %q but got %q for PendingExternalUserProfileId", v.Expected.PendingExternalUserProfileId, actual.PendingExternalUserProfileId)
		}

	}
}

func TestSegmentsForDirectoryPendingExternalUserProfileId(t *testing.T) {
	segments := DirectoryPendingExternalUserProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryPendingExternalUserProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
