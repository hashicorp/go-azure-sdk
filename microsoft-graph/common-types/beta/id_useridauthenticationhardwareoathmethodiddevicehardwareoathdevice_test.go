package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}

func TestNewUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(t *testing.T) {
	id := NewUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID("userId", "hardwareOathAuthenticationMethodId", "hardwareOathTokenAuthenticationMethodDeviceId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.HardwareOathAuthenticationMethodId != "hardwareOathAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathAuthenticationMethodId'", id.HardwareOathAuthenticationMethodId, "hardwareOathAuthenticationMethodId")
	}

	if id.HardwareOathTokenAuthenticationMethodDeviceId != "hardwareOathTokenAuthenticationMethodDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathTokenAuthenticationMethodDeviceId'", id.HardwareOathTokenAuthenticationMethodDeviceId, "hardwareOathTokenAuthenticationMethodDeviceId")
	}
}

func TestFormatUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(t *testing.T) {
	actual := NewUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID("userId", "hardwareOathAuthenticationMethodId", "hardwareOathTokenAuthenticationMethodDeviceId").ID()
	expected := "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
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
			// Incomplete URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId",
			Expected: &UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
				UserId:                             "userId",
				HardwareOathAuthenticationMethodId: "hardwareOathAuthenticationMethodId",
				HardwareOathTokenAuthenticationMethodDeviceId: "hardwareOathTokenAuthenticationMethodDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(v.Input)
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

		if actual.HardwareOathTokenAuthenticationMethodDeviceId != v.Expected.HardwareOathTokenAuthenticationMethodDeviceId {
			t.Fatalf("Expected %q but got %q for HardwareOathTokenAuthenticationMethodDeviceId", v.Expected.HardwareOathTokenAuthenticationMethodDeviceId, actual.HardwareOathTokenAuthenticationMethodDeviceId)
		}

	}
}

func TestParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
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
			// Incomplete URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE/hArDwArEoAtHdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId",
			Expected: &UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
				UserId:                             "userId",
				HardwareOathAuthenticationMethodId: "hardwareOathAuthenticationMethodId",
				HardwareOathTokenAuthenticationMethodDeviceId: "hardwareOathTokenAuthenticationMethodDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE/hArDwArEoAtHdEvIcEs/hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId",
			Expected: &UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
				UserId:                             "uSeRiD",
				HardwareOathAuthenticationMethodId: "hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
				HardwareOathTokenAuthenticationMethodDeviceId: "hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE/hArDwArEoAtHdEvIcEs/hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively(v.Input)
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

		if actual.HardwareOathTokenAuthenticationMethodDeviceId != v.Expected.HardwareOathTokenAuthenticationMethodDeviceId {
			t.Fatalf("Expected %q but got %q for HardwareOathTokenAuthenticationMethodDeviceId", v.Expected.HardwareOathTokenAuthenticationMethodDeviceId, actual.HardwareOathTokenAuthenticationMethodDeviceId)
		}

	}
}

func TestSegmentsForUserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId(t *testing.T) {
	segments := UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
