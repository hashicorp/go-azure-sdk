package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedAppLogCollectionRequestId{}

func TestNewUserIdManagedAppLogCollectionRequestID(t *testing.T) {
	id := NewUserIdManagedAppLogCollectionRequestID("userId", "managedAppLogCollectionRequestId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ManagedAppLogCollectionRequestId != "managedAppLogCollectionRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedAppLogCollectionRequestId'", id.ManagedAppLogCollectionRequestId, "managedAppLogCollectionRequestId")
	}
}

func TestFormatUserIdManagedAppLogCollectionRequestID(t *testing.T) {
	actual := NewUserIdManagedAppLogCollectionRequestID("userId", "managedAppLogCollectionRequestId").ID()
	expected := "/users/userId/managedAppLogCollectionRequests/managedAppLogCollectionRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedAppLogCollectionRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedAppLogCollectionRequestId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedAppLogCollectionRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedAppLogCollectionRequests/managedAppLogCollectionRequestId",
			Expected: &UserIdManagedAppLogCollectionRequestId{
				UserId:                           "userId",
				ManagedAppLogCollectionRequestId: "managedAppLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedAppLogCollectionRequests/managedAppLogCollectionRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedAppLogCollectionRequestID(v.Input)
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

		if actual.ManagedAppLogCollectionRequestId != v.Expected.ManagedAppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for ManagedAppLogCollectionRequestId", v.Expected.ManagedAppLogCollectionRequestId, actual.ManagedAppLogCollectionRequestId)
		}

	}
}

func TestParseUserIdManagedAppLogCollectionRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedAppLogCollectionRequestId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedAppLogCollectionRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdApPlOgCoLlEcTiOnReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedAppLogCollectionRequests/managedAppLogCollectionRequestId",
			Expected: &UserIdManagedAppLogCollectionRequestId{
				UserId:                           "userId",
				ManagedAppLogCollectionRequestId: "managedAppLogCollectionRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedAppLogCollectionRequests/managedAppLogCollectionRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdApPlOgCoLlEcTiOnReQuEsTs/mAnAgEdApPlOgCoLlEcTiOnReQuEsTiD",
			Expected: &UserIdManagedAppLogCollectionRequestId{
				UserId:                           "uSeRiD",
				ManagedAppLogCollectionRequestId: "mAnAgEdApPlOgCoLlEcTiOnReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdApPlOgCoLlEcTiOnReQuEsTs/mAnAgEdApPlOgCoLlEcTiOnReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedAppLogCollectionRequestIDInsensitively(v.Input)
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

		if actual.ManagedAppLogCollectionRequestId != v.Expected.ManagedAppLogCollectionRequestId {
			t.Fatalf("Expected %q but got %q for ManagedAppLogCollectionRequestId", v.Expected.ManagedAppLogCollectionRequestId, actual.ManagedAppLogCollectionRequestId)
		}

	}
}

func TestSegmentsForUserIdManagedAppLogCollectionRequestId(t *testing.T) {
	segments := UserIdManagedAppLogCollectionRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedAppLogCollectionRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
