package invitedusersponsor

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvitationIdInvitedUserSponsorId{}

func TestNewInvitationIdInvitedUserSponsorID(t *testing.T) {
	id := NewInvitationIdInvitedUserSponsorID("invitationIdValue", "directoryObjectIdValue")

	if id.InvitationId != "invitationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvitationId'", id.InvitationId, "invitationIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatInvitationIdInvitedUserSponsorID(t *testing.T) {
	actual := NewInvitationIdInvitedUserSponsorID("invitationIdValue", "directoryObjectIdValue").ID()
	expected := "/invitations/invitationIdValue/invitedUserSponsors/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseInvitationIdInvitedUserSponsorID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvitationIdInvitedUserSponsorId
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
			Input: "/invitations/invitationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/invitations/invitationIdValue/invitedUserSponsors",
			Error: true,
		},
		{
			// Valid URI
			Input: "/invitations/invitationIdValue/invitedUserSponsors/directoryObjectIdValue",
			Expected: &InvitationIdInvitedUserSponsorId{
				InvitationId:      "invitationIdValue",
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/invitations/invitationIdValue/invitedUserSponsors/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvitationIdInvitedUserSponsorID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InvitationId != v.Expected.InvitationId {
			t.Fatalf("Expected %q but got %q for InvitationId", v.Expected.InvitationId, actual.InvitationId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseInvitationIdInvitedUserSponsorIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvitationIdInvitedUserSponsorId
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
			Input: "/invitations/invitationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/invitations/invitationIdValue/invitedUserSponsors",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItAtIoNiDvAlUe/iNvItEdUsErSpOnSoRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/invitations/invitationIdValue/invitedUserSponsors/directoryObjectIdValue",
			Expected: &InvitationIdInvitedUserSponsorId{
				InvitationId:      "invitationIdValue",
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/invitations/invitationIdValue/invitedUserSponsors/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItAtIoNiDvAlUe/iNvItEdUsErSpOnSoRs/dIrEcToRyObJeCtIdVaLuE",
			Expected: &InvitationIdInvitedUserSponsorId{
				InvitationId:      "iNvItAtIoNiDvAlUe",
				DirectoryObjectId: "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItAtIoNiDvAlUe/iNvItEdUsErSpOnSoRs/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvitationIdInvitedUserSponsorIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.InvitationId != v.Expected.InvitationId {
			t.Fatalf("Expected %q but got %q for InvitationId", v.Expected.InvitationId, actual.InvitationId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForInvitationIdInvitedUserSponsorId(t *testing.T) {
	segments := InvitationIdInvitedUserSponsorId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("InvitationIdInvitedUserSponsorId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
