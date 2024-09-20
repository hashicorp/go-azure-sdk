package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationOperationId{}

func TestNewUserIdAuthenticationOperationID(t *testing.T) {
	id := NewUserIdAuthenticationOperationID("userId", "longRunningOperationId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.LongRunningOperationId != "longRunningOperationId" {
		t.Fatalf("Expected %q but got %q for Segment 'LongRunningOperationId'", id.LongRunningOperationId, "longRunningOperationId")
	}
}

func TestFormatUserIdAuthenticationOperationID(t *testing.T) {
	actual := NewUserIdAuthenticationOperationID("userId", "longRunningOperationId").ID()
	expected := "/users/userId/authentication/operations/longRunningOperationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationOperationId
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
			Input: "/users/userId/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/operations/longRunningOperationId",
			Expected: &UserIdAuthenticationOperationId{
				UserId:                 "userId",
				LongRunningOperationId: "longRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/operations/longRunningOperationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationOperationID(v.Input)
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

		if actual.LongRunningOperationId != v.Expected.LongRunningOperationId {
			t.Fatalf("Expected %q but got %q for LongRunningOperationId", v.Expected.LongRunningOperationId, actual.LongRunningOperationId)
		}

	}
}

func TestParseUserIdAuthenticationOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationOperationId
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
			Input: "/users/userId/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/operations/longRunningOperationId",
			Expected: &UserIdAuthenticationOperationId{
				UserId:                 "userId",
				LongRunningOperationId: "longRunningOperationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/operations/longRunningOperationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/oPeRaTiOnS/lOnGrUnNiNgOpErAtIoNiD",
			Expected: &UserIdAuthenticationOperationId{
				UserId:                 "uSeRiD",
				LongRunningOperationId: "lOnGrUnNiNgOpErAtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/oPeRaTiOnS/lOnGrUnNiNgOpErAtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationOperationIDInsensitively(v.Input)
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

		if actual.LongRunningOperationId != v.Expected.LongRunningOperationId {
			t.Fatalf("Expected %q but got %q for LongRunningOperationId", v.Expected.LongRunningOperationId, actual.LongRunningOperationId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationOperationId(t *testing.T) {
	segments := UserIdAuthenticationOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
