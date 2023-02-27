package consumerinvitation

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ConsumerInvitationId{}

func TestNewConsumerInvitationID(t *testing.T) {
	id := NewConsumerInvitationID("locationValue", "invitationIdValue")

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.InvitationId != "invitationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InvitationId'", id.InvitationId, "invitationIdValue")
	}
}

func TestFormatConsumerInvitationID(t *testing.T) {
	actual := NewConsumerInvitationID("locationValue", "invitationIdValue").ID()
	expected := "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations/invitationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseConsumerInvitationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ConsumerInvitationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations/invitationIdValue",
			Expected: &ConsumerInvitationId{
				LocationName: "locationValue",
				InvitationId: "invitationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations/invitationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseConsumerInvitationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.InvitationId != v.Expected.InvitationId {
			t.Fatalf("Expected %q but got %q for InvitationId", v.Expected.InvitationId, actual.InvitationId)
		}

	}
}

func TestParseConsumerInvitationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ConsumerInvitationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs/lOcAtIoNvAlUe/cOnSuMeRiNvItAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations/invitationIdValue",
			Expected: &ConsumerInvitationId{
				LocationName: "locationValue",
				InvitationId: "invitationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DataShare/locations/locationValue/consumerInvitations/invitationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs/lOcAtIoNvAlUe/cOnSuMeRiNvItAtIoNs/iNvItAtIoNiDvAlUe",
			Expected: &ConsumerInvitationId{
				LocationName: "lOcAtIoNvAlUe",
				InvitationId: "iNvItAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dAtAsHaRe/lOcAtIoNs/lOcAtIoNvAlUe/cOnSuMeRiNvItAtIoNs/iNvItAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseConsumerInvitationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.InvitationId != v.Expected.InvitationId {
			t.Fatalf("Expected %q but got %q for InvitationId", v.Expected.InvitationId, actual.InvitationId)
		}

	}
}

func TestSegmentsForConsumerInvitationId(t *testing.T) {
	segments := ConsumerInvitationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ConsumerInvitationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
