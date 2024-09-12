package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationSoftwareOathMethodId{}

func TestNewUserIdAuthenticationSoftwareOathMethodID(t *testing.T) {
	id := NewUserIdAuthenticationSoftwareOathMethodID("userIdValue", "softwareOathAuthenticationMethodIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.SoftwareOathAuthenticationMethodId != "softwareOathAuthenticationMethodIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SoftwareOathAuthenticationMethodId'", id.SoftwareOathAuthenticationMethodId, "softwareOathAuthenticationMethodIdValue")
	}
}

func TestFormatUserIdAuthenticationSoftwareOathMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationSoftwareOathMethodID("userIdValue", "softwareOathAuthenticationMethodIdValue").ID()
	expected := "/users/userIdValue/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationSoftwareOathMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationSoftwareOathMethodId
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
			Input: "/users/userIdValue/authentication/softwareOathMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationSoftwareOathMethodId{
				UserId:                             "userIdValue",
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationSoftwareOathMethodID(v.Input)
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

		if actual.SoftwareOathAuthenticationMethodId != v.Expected.SoftwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for SoftwareOathAuthenticationMethodId", v.Expected.SoftwareOathAuthenticationMethodId, actual.SoftwareOathAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationSoftwareOathMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationSoftwareOathMethodId
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
			Input: "/users/userIdValue/authentication/softwareOathMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue",
			Expected: &UserIdAuthenticationSoftwareOathMethodId{
				UserId:                             "userIdValue",
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/authentication/softwareOathMethods/softwareOathAuthenticationMethodIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			Expected: &UserIdAuthenticationSoftwareOathMethodId{
				UserId:                             "uSeRiDvAlUe",
				SoftwareOathAuthenticationMethodId: "sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationSoftwareOathMethodIDInsensitively(v.Input)
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

		if actual.SoftwareOathAuthenticationMethodId != v.Expected.SoftwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for SoftwareOathAuthenticationMethodId", v.Expected.SoftwareOathAuthenticationMethodId, actual.SoftwareOathAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationSoftwareOathMethodId(t *testing.T) {
	segments := UserIdAuthenticationSoftwareOathMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationSoftwareOathMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
