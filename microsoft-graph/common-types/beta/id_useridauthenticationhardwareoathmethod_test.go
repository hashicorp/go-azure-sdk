package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationHardwareOathMethodId{}

func TestNewUserIdAuthenticationHardwareOathMethodID(t *testing.T) {
	id := NewUserIdAuthenticationHardwareOathMethodID("userId", "hardwareOathAuthenticationMethodId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.HardwareOathAuthenticationMethodId != "hardwareOathAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathAuthenticationMethodId'", id.HardwareOathAuthenticationMethodId, "hardwareOathAuthenticationMethodId")
	}
}

func TestFormatUserIdAuthenticationHardwareOathMethodID(t *testing.T) {
	actual := NewUserIdAuthenticationHardwareOathMethodID("userId", "hardwareOathAuthenticationMethodId").ID()
	expected := "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationHardwareOathMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationHardwareOathMethodId
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
			Input: "/users/userId/authentication/hardwareOathMethods",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Expected: &UserIdAuthenticationHardwareOathMethodId{
				UserId:                             "userId",
				HardwareOathAuthenticationMethodId: "hardwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationHardwareOathMethodID(v.Input)
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

		if actual.HardwareOathAuthenticationMethodId != v.Expected.HardwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for HardwareOathAuthenticationMethodId", v.Expected.HardwareOathAuthenticationMethodId, actual.HardwareOathAuthenticationMethodId)
		}

	}
}

func TestParseUserIdAuthenticationHardwareOathMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationHardwareOathMethodId
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
			Input: "/users/userId/authentication/hardwareOathMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Expected: &UserIdAuthenticationHardwareOathMethodId{
				UserId:                             "userId",
				HardwareOathAuthenticationMethodId: "hardwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &UserIdAuthenticationHardwareOathMethodId{
				UserId:                             "uSeRiD",
				HardwareOathAuthenticationMethodId: "hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationHardwareOathMethodIDInsensitively(v.Input)
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

		if actual.HardwareOathAuthenticationMethodId != v.Expected.HardwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for HardwareOathAuthenticationMethodId", v.Expected.HardwareOathAuthenticationMethodId, actual.HardwareOathAuthenticationMethodId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationHardwareOathMethodId(t *testing.T) {
	segments := UserIdAuthenticationHardwareOathMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationHardwareOathMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
