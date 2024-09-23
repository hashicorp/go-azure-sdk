package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedAppRegistrationId{}

func TestNewMeManagedAppRegistrationID(t *testing.T) {
	id := NewMeManagedAppRegistrationID("managedAppRegistrationId")

	if id.ManagedAppRegistrationId != "managedAppRegistrationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedAppRegistrationId'", id.ManagedAppRegistrationId, "managedAppRegistrationId")
	}
}

func TestFormatMeManagedAppRegistrationID(t *testing.T) {
	actual := NewMeManagedAppRegistrationID("managedAppRegistrationId").ID()
	expected := "/me/managedAppRegistrations/managedAppRegistrationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedAppRegistrationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedAppRegistrationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedAppRegistrations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedAppRegistrations/managedAppRegistrationId",
			Expected: &MeManagedAppRegistrationId{
				ManagedAppRegistrationId: "managedAppRegistrationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedAppRegistrations/managedAppRegistrationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedAppRegistrationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedAppRegistrationId != v.Expected.ManagedAppRegistrationId {
			t.Fatalf("Expected %q but got %q for ManagedAppRegistrationId", v.Expected.ManagedAppRegistrationId, actual.ManagedAppRegistrationId)
		}

	}
}

func TestParseMeManagedAppRegistrationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedAppRegistrationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedAppRegistrations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdApPrEgIsTrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedAppRegistrations/managedAppRegistrationId",
			Expected: &MeManagedAppRegistrationId{
				ManagedAppRegistrationId: "managedAppRegistrationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedAppRegistrations/managedAppRegistrationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdApPrEgIsTrAtIoNs/mAnAgEdApPrEgIsTrAtIoNiD",
			Expected: &MeManagedAppRegistrationId{
				ManagedAppRegistrationId: "mAnAgEdApPrEgIsTrAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdApPrEgIsTrAtIoNs/mAnAgEdApPrEgIsTrAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedAppRegistrationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedAppRegistrationId != v.Expected.ManagedAppRegistrationId {
			t.Fatalf("Expected %q but got %q for ManagedAppRegistrationId", v.Expected.ManagedAppRegistrationId, actual.ManagedAppRegistrationId)
		}

	}
}

func TestSegmentsForMeManagedAppRegistrationId(t *testing.T) {
	segments := MeManagedAppRegistrationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedAppRegistrationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
