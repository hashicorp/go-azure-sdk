package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationWindowsHelloForBusinessMethodId{}

func TestNewUserIdAuthenticationWindowsHelloForBusinessMethodID(t *testing.T) {
	id := NewUserIdAuthenticationWindowsHelloForBusinessMethodID("userIdValue", "windowsHelloForBusinessAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.WindowsHelloForBusinessAuthenticationMethodId != "windowsHelloForBusinessAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsHelloForBusinessAuthenticationMethodId'", id.WindowsHelloForBusinessAuthenticationMethodId, "windowsHelloForBusinessAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationWindowsHelloForBusinessMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationWindowsHelloForBusinessMethodID("userIdValue", "windowsHelloForBusinessAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationWindowsHelloForBusinessMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationWindowsHelloForBusinessMethodId
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
			Input: "/users/userIdValue/authentication/windowsHelloForBusinessMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationWindowsHelloForBusinessMethodId{
				UserId: "userIdValue",
				WindowsHelloForBusinessAuthenticationMethodId: "windowsHelloForBusinessAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationWindowsHelloForBusinessMethodID(v.Input)
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

		if actual.WindowsHelloForBusinessAuthenticationMethodId != v.Expected.WindowsHelloForBusinessAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for WindowsHelloForBusinessAuthenticationMethodId", v.Expected.WindowsHelloForBusinessAuthenticationMethodId, actual.WindowsHelloForBusinessAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationWindowsHelloForBusinessMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationWindowsHelloForBusinessMethodId
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
			Input: "/users/userIdValue/authentication/windowsHelloForBusinessMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/wInDoWsHeLlOfOrBuSiNeSsMeThOdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationWindowsHelloForBusinessMethodId{
				UserId: "userIdValue",
				WindowsHelloForBusinessAuthenticationMethodId: "windowsHelloForBusinessAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/windowsHelloForBusinessMethods/windowsHelloForBusinessAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/wInDoWsHeLlOfOrBuSiNeSsMeThOdS/wInDoWsHeLlOfOrBuSiNeSsAuThEnTiCaTiOnMeThOdIdVaLuE",
			Expected: &UserIdAuthenticationWindowsHelloForBusinessMethodId{
				UserId: "uSeRiDvAlUe",
				WindowsHelloForBusinessAuthenticationMethodId: "wInDoWsHeLlOfOrBuSiNeSsAuThEnTiCaTiOnMeThOdIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/wInDoWsHeLlOfOrBuSiNeSsMeThOdS/wInDoWsHeLlOfOrBuSiNeSsAuThEnTiCaTiOnMeThOdIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationWindowsHelloForBusinessMethodIDInsensitively(v.Input)
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

		if actual.WindowsHelloForBusinessAuthenticationMethodId != v.Expected.WindowsHelloForBusinessAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for WindowsHelloForBusinessAuthenticationMethodId", v.Expected.WindowsHelloForBusinessAuthenticationMethodId, actual.WindowsHelloForBusinessAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationWindowsHelloForBusinessMethodId(t *testing.T) {
	segments := UserIdAuthenticationWindowsHelloForBusinessMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationWindowsHelloForBusinessMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
