package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationFido2MethodId{}

func TestNewUserIdAuthenticationFido2MethodID(t *testing.T) {
	id := NewUserIdAuthenticationFido2MethodID("userIdValue", "fido2AuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.Fido2AuthenticationMethodId != "fido2AuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'Fido2AuthenticationMethodId'", id.Fido2AuthenticationMethodId, "fido2AuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationFido2MethodID(t *testing.T) {
	actual := NewUserIdAuthenticationFido2MethodID("userIdValue", "fido2AuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/fido2Methods/fido2AuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationFido2MethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationFido2MethodId
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
			Input: "/users/userIdValue/authentication/fido2Methods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/fido2Methods/fido2AuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationFido2MethodId{
				UserId:                      "userIdValue",
				Fido2AuthenticationMethodId: "fido2AuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/fido2Methods/fido2AuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationFido2MethodID(v.Input)
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

		if actual.Fido2AuthenticationMethodId != v.Expected.Fido2AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for Fido2AuthenticationMethodId", v.Expected.Fido2AuthenticationMethodId, actual.Fido2AuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationFido2MethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationFido2MethodId
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
			Input: "/users/userIdValue/authentication/fido2Methods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/fIdO2MeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/fido2Methods/fido2AuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationFido2MethodId{
				UserId:                      "userIdValue",
				Fido2AuthenticationMethodId: "fido2AuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/fido2Methods/fido2AuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/fIdO2MeThOdS/fIdO2AuThEnTiCaTiOnMeThOdIdVaLuE",
			Expected: &UserIdAuthenticationFido2MethodId{
				UserId:                      "uSeRiDvAlUe",
				Fido2AuthenticationMethodId: "fIdO2AuThEnTiCaTiOnMeThOdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/fIdO2MeThOdS/fIdO2AuThEnTiCaTiOnMeThOdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationFido2MethodIDInsensitively(v.Input)
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

		if actual.Fido2AuthenticationMethodId != v.Expected.Fido2AuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for Fido2AuthenticationMethodId", v.Expected.Fido2AuthenticationMethodId, actual.Fido2AuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationFido2MethodId(t *testing.T) {
	segments := UserIdAuthenticationFido2MethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationFido2MethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
