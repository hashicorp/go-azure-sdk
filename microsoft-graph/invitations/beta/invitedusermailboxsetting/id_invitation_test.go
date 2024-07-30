package invitedusermailboxsetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InvitationId{}

func TestNewInvitationID(t *testing.T) {
	id := NewInvitationID("invitationIdValue")

	if id.InvitationId != "invitationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvitationId'", id.InvitationId, "invitationIdValue")
	}
}

func TestFormatInvitationID(t *testing.T) {
	actual := NewInvitationID("invitationIdValue").ID()
	expected := "/invitations/invitationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseInvitationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvitationId
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
			// Valid URI
			Input: "/invitations/invitationIdValue",
			Expected: &InvitationId{
				InvitationId: "invitationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/invitations/invitationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvitationID(v.Input)
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

	}
}

func TestParseInvitationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *InvitationId
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
			// Valid URI
			Input: "/invitations/invitationIdValue",
			Expected: &InvitationId{
				InvitationId: "invitationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/invitations/invitationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItAtIoNiDvAlUe",
			Expected: &InvitationId{
				InvitationId: "iNvItAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iNvItAtIoNs/iNvItAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseInvitationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForInvitationId(t *testing.T) {
	segments := InvitationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("InvitationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
