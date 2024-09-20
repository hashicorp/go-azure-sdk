package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvitationInvitedUserSponsorId{}

func TestNewInvitationInvitedUserSponsorID(t *testing.T) {
	id := NewInvitationInvitedUserSponsorID("directoryObjectId")

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatInvitationInvitedUserSponsorID(t *testing.T) {
	actual := NewInvitationInvitedUserSponsorID("directoryObjectId").ID()
	expected := "/invitations/invitedUserSponsors/directoryObjectId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseInvitationInvitedUserSponsorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvitationInvitedUserSponsorId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/invitations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/invitations/invitedUserSponsors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/invitations/invitedUserSponsors/directoryObjectId",
			Expected: &InvitationInvitedUserSponsorId{
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/invitations/invitedUserSponsors/directoryObjectId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvitationInvitedUserSponsorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseInvitationInvitedUserSponsorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvitationInvitedUserSponsorId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/invitations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/invitations/invitedUserSponsors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItEdUsErSpOnSoRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/invitations/invitedUserSponsors/directoryObjectId",
			Expected: &InvitationInvitedUserSponsorId{
				DirectoryObjectId: "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/invitations/invitedUserSponsors/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItEdUsErSpOnSoRs/dIrEcToRyObJeCtId",
			Expected: &InvitationInvitedUserSponsorId{
				DirectoryObjectId: "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItEdUsErSpOnSoRs/dIrEcToRyObJeCtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvitationInvitedUserSponsorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForInvitationInvitedUserSponsorId(t *testing.T) {
	segments := InvitationInvitedUserSponsorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("InvitationInvitedUserSponsorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
