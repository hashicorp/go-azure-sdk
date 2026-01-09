package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPhoneMethodId{}

func TestNewUserIdAuthenticationPhoneMethodID(t *testing.T) {
	id := NewUserIdAuthenticationPhoneMethodID("userId", "phoneAuthenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.PhoneAuthenticationMethodId != "phoneAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'PhoneAuthenticationMethodId'", id.PhoneAuthenticationMethodId, "phoneAuthenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationPhoneMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationPhoneMethodID("userId", "phoneAuthenticationMethodId").ID()
	expected := "/users/userId/authentication/phoneMethods/phoneAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationPhoneMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPhoneMethodId
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
			Input: "/users/userId/authentication/phoneMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/phoneMethods/phoneAuthenticationMethodId",
			Expected: &UserIdAuthenticationPhoneMethodId{
				UserId:                      "userId",
				PhoneAuthenticationMethodId: "phoneAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/phoneMethods/phoneAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPhoneMethodID(v.Input)
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

		if actual.PhoneAuthenticationMethodId != v.Expected.PhoneAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PhoneAuthenticationMethodId", v.Expected.PhoneAuthenticationMethodId, actual.PhoneAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationPhoneMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPhoneMethodId
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
			Input: "/users/userId/authentication/phoneMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/pHoNeMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/phoneMethods/phoneAuthenticationMethodId",
			Expected: &UserIdAuthenticationPhoneMethodId{
				UserId:                      "userId",
				PhoneAuthenticationMethodId: "phoneAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/phoneMethods/phoneAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/pHoNeMeThOdS/pHoNeAuThEnTiCaTiOnMeThOdId",
			Expected: &UserIdAuthenticationPhoneMethodId{
				UserId:                      "uSeRiD",
				PhoneAuthenticationMethodId: "pHoNeAuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/pHoNeMeThOdS/pHoNeAuThEnTiCaTiOnMeThOdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPhoneMethodIDInsensitively(v.Input)
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

		if actual.PhoneAuthenticationMethodId != v.Expected.PhoneAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PhoneAuthenticationMethodId", v.Expected.PhoneAuthenticationMethodId, actual.PhoneAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationPhoneMethodId(t *testing.T) {
	segments := UserIdAuthenticationPhoneMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationPhoneMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
