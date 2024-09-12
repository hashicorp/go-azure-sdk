package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedAppRegistrationId{}

func TestNewUserIdManagedAppRegistrationID(t *testing.T) {
	id := NewUserIdManagedAppRegistrationID("userIdValue", "managedAppRegistrationIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ManagedAppRegistrationId != "managedAppRegistrationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedAppRegistrationId'", id.ManagedAppRegistrationId, "managedAppRegistrationIdValue")
	}
}

func TestFormatUserIdManagedAppRegistrationID(t *testing.T) {
	actual := NewUserIdManagedAppRegistrationID("userIdValue", "managedAppRegistrationIdValue").ID()
	expected := "/users/userIdValue/managedAppRegistrations/managedAppRegistrationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedAppRegistrationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedAppRegistrationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedAppRegistrations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedAppRegistrations/managedAppRegistrationIdValue",
			Expected: &UserIdManagedAppRegistrationId{
				UserId:                   "userIdValue",
				ManagedAppRegistrationId: "managedAppRegistrationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedAppRegistrations/managedAppRegistrationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedAppRegistrationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.ManagedAppRegistrationId != v.Expected.ManagedAppRegistrationId {
			t.Fatalf("Expected %q but got %q for ManagedAppRegistrationId", v.Expected.ManagedAppRegistrationId, actual.ManagedAppRegistrationId)
		}

	}
}

func TestParseUserIdManagedAppRegistrationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedAppRegistrationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedAppRegistrations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdApPrEgIsTrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedAppRegistrations/managedAppRegistrationIdValue",
			Expected: &UserIdManagedAppRegistrationId{
				UserId:                   "userIdValue",
				ManagedAppRegistrationId: "managedAppRegistrationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedAppRegistrations/managedAppRegistrationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdApPrEgIsTrAtIoNs/mAnAgEdApPrEgIsTrAtIoNiDvAlUe",
			Expected: &UserIdManagedAppRegistrationId{
				UserId:                   "uSeRiDvAlUe",
				ManagedAppRegistrationId: "mAnAgEdApPrEgIsTrAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdApPrEgIsTrAtIoNs/mAnAgEdApPrEgIsTrAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedAppRegistrationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.ManagedAppRegistrationId != v.Expected.ManagedAppRegistrationId {
			t.Fatalf("Expected %q but got %q for ManagedAppRegistrationId", v.Expected.ManagedAppRegistrationId, actual.ManagedAppRegistrationId)
		}

	}
}

func TestSegmentsForUserIdManagedAppRegistrationId(t *testing.T) {
	segments := UserIdManagedAppRegistrationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedAppRegistrationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
