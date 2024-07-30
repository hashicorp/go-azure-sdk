package authenticationemailmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationEmailMethodId{}

func TestNewUserIdAuthenticationEmailMethodID(t *testing.T) {
	id := NewUserIdAuthenticationEmailMethodID("userIdValue", "emailAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.EmailAuthenticationMethodId != "emailAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'EmailAuthenticationMethodId'", id.EmailAuthenticationMethodId, "emailAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationEmailMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationEmailMethodID("userIdValue", "emailAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/emailMethods/emailAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationEmailMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationEmailMethodId
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
			Input: "/users/userIdValue/authentication/emailMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/emailMethods/emailAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationEmailMethodId{
				UserId:                      "userIdValue",
				EmailAuthenticationMethodId: "emailAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/emailMethods/emailAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationEmailMethodID(v.Input)
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

		if actual.EmailAuthenticationMethodId != v.Expected.EmailAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for EmailAuthenticationMethodId", v.Expected.EmailAuthenticationMethodId, actual.EmailAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationEmailMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationEmailMethodId
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
			Input: "/users/userIdValue/authentication/emailMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/eMaIlMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/emailMethods/emailAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationEmailMethodId{
				UserId:                      "userIdValue",
				EmailAuthenticationMethodId: "emailAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/emailMethods/emailAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/eMaIlMeThOdS/eMaIlAuThEnTiCaTiOnMeThOdIdVaLuE",
			Expected: &UserIdAuthenticationEmailMethodId{
				UserId:                      "uSeRiDvAlUe",
				EmailAuthenticationMethodId: "eMaIlAuThEnTiCaTiOnMeThOdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/eMaIlMeThOdS/eMaIlAuThEnTiCaTiOnMeThOdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationEmailMethodIDInsensitively(v.Input)
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

		if actual.EmailAuthenticationMethodId != v.Expected.EmailAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for EmailAuthenticationMethodId", v.Expected.EmailAuthenticationMethodId, actual.EmailAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationEmailMethodId(t *testing.T) {
	segments := UserIdAuthenticationEmailMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationEmailMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
