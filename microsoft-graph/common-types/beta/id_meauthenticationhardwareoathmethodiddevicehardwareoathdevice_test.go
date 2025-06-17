package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}

func TestNewMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(t *testing.T) {
	id := NewMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID("hardwareOathAuthenticationMethodId", "hardwareOathTokenAuthenticationMethodDeviceId")

	if id.HardwareOathAuthenticationMethodId != "hardwareOathAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathAuthenticationMethodId'", id.HardwareOathAuthenticationMethodId, "hardwareOathAuthenticationMethodId")
	}

	if id.HardwareOathTokenAuthenticationMethodDeviceId != "hardwareOathTokenAuthenticationMethodDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathTokenAuthenticationMethodDeviceId'", id.HardwareOathTokenAuthenticationMethodDeviceId, "hardwareOathTokenAuthenticationMethodDeviceId")
	}
}

func TestFormatMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(t *testing.T) {
	actual := NewMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID("hardwareOathAuthenticationMethodId", "hardwareOathTokenAuthenticationMethodDeviceId").ID()
	expected := "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId",
			Expected: &MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
				HardwareOathAuthenticationMethodId:            "hardwareOathAuthenticationMethodId",
				HardwareOathTokenAuthenticationMethodDeviceId: "hardwareOathTokenAuthenticationMethodDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwareOathAuthenticationMethodId != v.Expected.HardwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for HardwareOathAuthenticationMethodId", v.Expected.HardwareOathAuthenticationMethodId, actual.HardwareOathAuthenticationMethodId)
		}

		if actual.HardwareOathTokenAuthenticationMethodDeviceId != v.Expected.HardwareOathTokenAuthenticationMethodDeviceId {
			t.Fatalf("Expected %q but got %q for HardwareOathTokenAuthenticationMethodDeviceId", v.Expected.HardwareOathTokenAuthenticationMethodDeviceId, actual.HardwareOathTokenAuthenticationMethodDeviceId)
		}

	}
}

func TestParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE/hArDwArEoAtHdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId",
			Expected: &MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
				HardwareOathAuthenticationMethodId:            "hardwareOathAuthenticationMethodId",
				HardwareOathTokenAuthenticationMethodDeviceId: "hardwareOathTokenAuthenticationMethodDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/device/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE/hArDwArEoAtHdEvIcEs/hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId",
			Expected: &MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{
				HardwareOathAuthenticationMethodId:            "hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
				HardwareOathTokenAuthenticationMethodDeviceId: "hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/dEvIcE/hArDwArEoAtHdEvIcEs/hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwareOathAuthenticationMethodId != v.Expected.HardwareOathAuthenticationMethodId {
			t.Fatalf("Expected %q but got %q for HardwareOathAuthenticationMethodId", v.Expected.HardwareOathAuthenticationMethodId, actual.HardwareOathAuthenticationMethodId)
		}

		if actual.HardwareOathTokenAuthenticationMethodDeviceId != v.Expected.HardwareOathTokenAuthenticationMethodDeviceId {
			t.Fatalf("Expected %q but got %q for HardwareOathTokenAuthenticationMethodDeviceId", v.Expected.HardwareOathTokenAuthenticationMethodDeviceId, actual.HardwareOathTokenAuthenticationMethodDeviceId)
		}

	}
}

func TestSegmentsForMeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId(t *testing.T) {
	segments := MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationHardwareOathMethodIdDeviceHardwareOathDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
