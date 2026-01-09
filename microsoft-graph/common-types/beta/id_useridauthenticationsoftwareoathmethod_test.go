package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationSoftwareOathMethodId{}

func TestNewUserIdAuthenticationSoftwareOathMethodID(t *testing.T) {
	id := NewUserIdAuthenticationSoftwareOathMethodID("userId", "softwareOathAuthenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SoftwareOathAuthenticationMethodId != "softwareOathAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'SoftwareOathAuthenticationMethodId'", id.SoftwareOathAuthenticationMethodId, "softwareOathAuthenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationSoftwareOathMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationSoftwareOathMethodID("userId", "softwareOathAuthenticationMethodId").ID()
	expected := "/users/userId/authentication/softwareOathMethods/softwareOathAuthenticationMethodId"
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
			Input: "/users/userId/authentication/softwareOathMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/softwareOathMethods/softwareOathAuthenticationMethodId",
			Expected: &UserIdAuthenticationSoftwareOathMethodId{
				UserId:                             "userId",
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/softwareOathMethods/softwareOathAuthenticationMethodId/extra",
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
			Input: "/users/userId/authentication/softwareOathMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/softwareOathMethods/softwareOathAuthenticationMethodId",
			Expected: &UserIdAuthenticationSoftwareOathMethodId{
				UserId:                             "userId",
				SoftwareOathAuthenticationMethodId: "softwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/softwareOathMethods/softwareOathAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &UserIdAuthenticationSoftwareOathMethodId{
				UserId:                             "uSeRiD",
				SoftwareOathAuthenticationMethodId: "sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/sOfTwArEoAtHmEtHoDs/sOfTwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/extra",
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
