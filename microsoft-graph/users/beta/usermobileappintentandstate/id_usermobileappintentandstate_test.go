package usermobileappintentandstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMobileAppIntentAndStateId{}

func TestNewUserMobileAppIntentAndStateID(t *testing.T) {
	id := NewUserMobileAppIntentAndStateID("userIdValue", "mobileAppIntentAndStateIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MobileAppIntentAndStateId != "mobileAppIntentAndStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppIntentAndStateId'", id.MobileAppIntentAndStateId, "mobileAppIntentAndStateIdValue")
	}
}

func TestFormatUserMobileAppIntentAndStateID(t *testing.T) {
	actual := NewUserMobileAppIntentAndStateID("userIdValue", "mobileAppIntentAndStateIdValue").ID()
	expected := "/users/userIdValue/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserMobileAppIntentAndStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMobileAppIntentAndStateId
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
			Input: "/users/userIdValue/mobileAppIntentAndStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue",
			Expected: &UserMobileAppIntentAndStateId{
				UserId:                    "userIdValue",
				MobileAppIntentAndStateId: "mobileAppIntentAndStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMobileAppIntentAndStateID(v.Input)
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

func TestParseUserMobileAppIntentAndStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMobileAppIntentAndStateId
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
			Input: "/users/userIdValue/mobileAppIntentAndStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpInTeNtAnDsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue",
			Expected: &UserMobileAppIntentAndStateId{
				UserId:                    "userIdValue",
				MobileAppIntentAndStateId: "mobileAppIntentAndStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpInTeNtAnDsTaTeS/mObIlEaPpInTeNtAnDsTaTeIdVaLuE",
			Expected: &UserMobileAppIntentAndStateId{
				UserId:                    "uSeRiDvAlUe",
				MobileAppIntentAndStateId: "mObIlEaPpInTeNtAnDsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mObIlEaPpInTeNtAnDsTaTeS/mObIlEaPpInTeNtAnDsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMobileAppIntentAndStateIDInsensitively(v.Input)
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

func TestSegmentsForUserMobileAppIntentAndStateId(t *testing.T) {
	segments := UserMobileAppIntentAndStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserMobileAppIntentAndStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
