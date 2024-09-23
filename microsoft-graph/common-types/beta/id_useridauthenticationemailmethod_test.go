package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationEmailMethodId{}

func TestNewUserIdAuthenticationEmailMethodID(t *testing.T) {
	id := NewUserIdAuthenticationEmailMethodID("userId", "emailAuthenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.EmailAuthenticationMethodId != "emailAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'EmailAuthenticationMethodId'", id.EmailAuthenticationMethodId, "emailAuthenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationEmailMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationEmailMethodID("userId", "emailAuthenticationMethodId").ID()
	expected := "/users/userId/authentication/emailMethods/emailAuthenticationMethodId"
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
			Input: "/users/userId/authentication/emailMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/emailMethods/emailAuthenticationMethodId",
			Expected: &UserIdAuthenticationEmailMethodId{
				UserId:                      "userId",
				EmailAuthenticationMethodId: "emailAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/emailMethods/emailAuthenticationMethodId/extra",
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
			Input: "/users/userId/authentication/emailMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/eMaIlMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/emailMethods/emailAuthenticationMethodId",
			Expected: &UserIdAuthenticationEmailMethodId{
				UserId:                      "userId",
				EmailAuthenticationMethodId: "emailAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/emailMethods/emailAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/eMaIlMeThOdS/eMaIlAuThEnTiCaTiOnMeThOdId",
			Expected: &UserIdAuthenticationEmailMethodId{
				UserId:                      "uSeRiD",
				EmailAuthenticationMethodId: "eMaIlAuThEnTiCaTiOnMeThOdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/eMaIlMeThOdS/eMaIlAuThEnTiCaTiOnMeThOdId/extra",
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
