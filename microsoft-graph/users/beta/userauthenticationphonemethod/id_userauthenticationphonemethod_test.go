package userauthenticationphonemethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationPhoneMethodId{}

func TestNewUserAuthenticationPhoneMethodID(t *testing.T) {
	id := NewUserAuthenticationPhoneMethodID("userIdValue", "phoneAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PhoneAuthenticationMethodId != "phoneAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PhoneAuthenticationMethodId'", id.PhoneAuthenticationMethodId, "phoneAuthenticationMethodIdValue")
	}
}

func TestFormatUserAuthenticationPhoneMethodID(t *testing.T) {
	actual := NewUserAuthenticationPhoneMethodID("userIdValue", "phoneAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/phoneMethods/phoneAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAuthenticationPhoneMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAuthenticationPhoneMethodId
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
			Input: "/users/userIdValue/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/authentication/phoneMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/phoneMethods/phoneAuthenticationMethodIdValue",
			Expected: &UserAuthenticationPhoneMethodId{
				UserId:                      "userIdValue",
				PhoneAuthenticationMethodId: "phoneAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/phoneMethods/phoneAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAuthenticationPhoneMethodID(v.Input)
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

func TestParseUserAuthenticationPhoneMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAuthenticationPhoneMethodId
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
			Input: "/users/userIdValue/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/authentication/phoneMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pHoNeMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/phoneMethods/phoneAuthenticationMethodIdValue",
			Expected: &UserAuthenticationPhoneMethodId{
				UserId:                      "userIdValue",
				PhoneAuthenticationMethodId: "phoneAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/phoneMethods/phoneAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pHoNeMeThOdS/pHoNeAuThEnTiCaTiOnMeThOdIdVaLuE",
			Expected: &UserAuthenticationPhoneMethodId{
				UserId:                      "uSeRiDvAlUe",
				PhoneAuthenticationMethodId: "pHoNeAuThEnTiCaTiOnMeThOdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pHoNeMeThOdS/pHoNeAuThEnTiCaTiOnMeThOdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAuthenticationPhoneMethodIDInsensitively(v.Input)
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

func TestSegmentsForUserAuthenticationPhoneMethodId(t *testing.T) {
	segments := UserAuthenticationPhoneMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAuthenticationPhoneMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
