package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMobileAppIntentAndStateId{}

func TestNewUserIdMobileAppIntentAndStateID(t *testing.T) {
	id := NewUserIdMobileAppIntentAndStateID("userId", "mobileAppIntentAndStateId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MobileAppIntentAndStateId != "mobileAppIntentAndStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppIntentAndStateId'", id.MobileAppIntentAndStateId, "mobileAppIntentAndStateId")
	}
}

func TestFormatUserIdMobileAppIntentAndStateID(t *testing.T) {
	actual := NewUserIdMobileAppIntentAndStateID("userId", "mobileAppIntentAndStateId").ID()
	expected := "/users/userId/mobileAppIntentAndStates/mobileAppIntentAndStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMobileAppIntentAndStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMobileAppIntentAndStateId
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
			Input: "/users/userId/mobileAppIntentAndStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mobileAppIntentAndStates/mobileAppIntentAndStateId",
			Expected: &UserIdMobileAppIntentAndStateId{
				UserId:                    "userId",
				MobileAppIntentAndStateId: "mobileAppIntentAndStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mobileAppIntentAndStates/mobileAppIntentAndStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMobileAppIntentAndStateID(v.Input)
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

		if actual.MobileAppIntentAndStateId != v.Expected.MobileAppIntentAndStateId {
			t.Fatalf("Expected %q but got %q for MobileAppIntentAndStateId", v.Expected.MobileAppIntentAndStateId, actual.MobileAppIntentAndStateId)
		}

	}
}

func TestParseUserIdMobileAppIntentAndStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMobileAppIntentAndStateId
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
			Input: "/users/userId/mobileAppIntentAndStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpInTeNtAnDsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mobileAppIntentAndStates/mobileAppIntentAndStateId",
			Expected: &UserIdMobileAppIntentAndStateId{
				UserId:                    "userId",
				MobileAppIntentAndStateId: "mobileAppIntentAndStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mobileAppIntentAndStates/mobileAppIntentAndStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpInTeNtAnDsTaTeS/mObIlEaPpInTeNtAnDsTaTeId",
			Expected: &UserIdMobileAppIntentAndStateId{
				UserId:                    "uSeRiD",
				MobileAppIntentAndStateId: "mObIlEaPpInTeNtAnDsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mObIlEaPpInTeNtAnDsTaTeS/mObIlEaPpInTeNtAnDsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMobileAppIntentAndStateIDInsensitively(v.Input)
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

		if actual.MobileAppIntentAndStateId != v.Expected.MobileAppIntentAndStateId {
			t.Fatalf("Expected %q but got %q for MobileAppIntentAndStateId", v.Expected.MobileAppIntentAndStateId, actual.MobileAppIntentAndStateId)
		}

	}
}

func TestSegmentsForUserIdMobileAppIntentAndStateId(t *testing.T) {
	segments := UserIdMobileAppIntentAndStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMobileAppIntentAndStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
