package authenticationplatformcredentialmethod

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationPlatformCredentialMethodId{}

func TestNewUserIdAuthenticationPlatformCredentialMethodID(t *testing.T) {
	id := NewUserIdAuthenticationPlatformCredentialMethodID("userIdValue", "platformCredentialAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.PlatformCredentialAuthenticationMethodId != "platformCredentialAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PlatformCredentialAuthenticationMethodId'", id.PlatformCredentialAuthenticationMethodId, "platformCredentialAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationPlatformCredentialMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationPlatformCredentialMethodID("userIdValue", "platformCredentialAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationPlatformCredentialMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPlatformCredentialMethodId
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
			Input: "/users/userIdValue/authentication/platformCredentialMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationPlatformCredentialMethodId{
				UserId:                                   "userIdValue",
				PlatformCredentialAuthenticationMethodId: "platformCredentialAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPlatformCredentialMethodID(v.Input)
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

		if actual.PlatformCredentialAuthenticationMethodId != v.Expected.PlatformCredentialAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PlatformCredentialAuthenticationMethodId", v.Expected.PlatformCredentialAuthenticationMethodId, actual.PlatformCredentialAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationPlatformCredentialMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationPlatformCredentialMethodId
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
			Input: "/users/userIdValue/authentication/platformCredentialMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pLaTfOrMcReDeNtIaLmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationPlatformCredentialMethodId{
				UserId:                                   "userIdValue",
				PlatformCredentialAuthenticationMethodId: "platformCredentialAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/platformCredentialMethods/platformCredentialAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pLaTfOrMcReDeNtIaLmEtHoDs/pLaTfOrMcReDeNtIaLaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &UserIdAuthenticationPlatformCredentialMethodId{
				UserId:                                   "uSeRiDvAlUe",
				PlatformCredentialAuthenticationMethodId: "pLaTfOrMcReDeNtIaLaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/pLaTfOrMcReDeNtIaLmEtHoDs/pLaTfOrMcReDeNtIaLaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationPlatformCredentialMethodIDInsensitively(v.Input)
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

		if actual.PlatformCredentialAuthenticationMethodId != v.Expected.PlatformCredentialAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for PlatformCredentialAuthenticationMethodId", v.Expected.PlatformCredentialAuthenticationMethodId, actual.PlatformCredentialAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationPlatformCredentialMethodId(t *testing.T) {
	segments := UserIdAuthenticationPlatformCredentialMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationPlatformCredentialMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
