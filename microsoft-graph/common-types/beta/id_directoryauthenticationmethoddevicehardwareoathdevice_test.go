package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{}

func TestNewDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(t *testing.T) {
	id := NewDirectoryAuthenticationMethodDeviceHardwareOathDeviceID("hardwareOathTokenAuthenticationMethodDeviceId")

	if id.HardwareOathTokenAuthenticationMethodDeviceId != "hardwareOathTokenAuthenticationMethodDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathTokenAuthenticationMethodDeviceId'", id.HardwareOathTokenAuthenticationMethodDeviceId, "hardwareOathTokenAuthenticationMethodDeviceId")
	}
}

func TestFormatDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(t *testing.T) {
	actual := NewDirectoryAuthenticationMethodDeviceHardwareOathDeviceID("hardwareOathTokenAuthenticationMethodDeviceId").ID()
	expected := "/directory/authenticationMethodDevices/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAuthenticationMethodDeviceHardwareOathDeviceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/authenticationMethodDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/authenticationMethodDevices/hardwareOathDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/authenticationMethodDevices/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId",
			Expected: &DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{
				HardwareOathTokenAuthenticationMethodDeviceId: "hardwareOathTokenAuthenticationMethodDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/authenticationMethodDevices/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwareOathTokenAuthenticationMethodDeviceId != v.Expected.HardwareOathTokenAuthenticationMethodDeviceId {
			t.Fatalf("Expected %q but got %q for HardwareOathTokenAuthenticationMethodDeviceId", v.Expected.HardwareOathTokenAuthenticationMethodDeviceId, actual.HardwareOathTokenAuthenticationMethodDeviceId)
		}

	}
}

func TestParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryAuthenticationMethodDeviceHardwareOathDeviceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/authenticationMethodDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aUtHeNtIcAtIoNmEtHoDdEvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/authenticationMethodDevices/hardwareOathDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aUtHeNtIcAtIoNmEtHoDdEvIcEs/hArDwArEoAtHdEvIcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/authenticationMethodDevices/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId",
			Expected: &DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{
				HardwareOathTokenAuthenticationMethodDeviceId: "hardwareOathTokenAuthenticationMethodDeviceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/authenticationMethodDevices/hardwareOathDevices/hardwareOathTokenAuthenticationMethodDeviceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aUtHeNtIcAtIoNmEtHoDdEvIcEs/hArDwArEoAtHdEvIcEs/hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId",
			Expected: &DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{
				HardwareOathTokenAuthenticationMethodDeviceId: "hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/aUtHeNtIcAtIoNmEtHoDdEvIcEs/hArDwArEoAtHdEvIcEs/hArDwArEoAtHtOkEnAuThEnTiCaTiOnMeThOdDeViCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryAuthenticationMethodDeviceHardwareOathDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HardwareOathTokenAuthenticationMethodDeviceId != v.Expected.HardwareOathTokenAuthenticationMethodDeviceId {
			t.Fatalf("Expected %q but got %q for HardwareOathTokenAuthenticationMethodDeviceId", v.Expected.HardwareOathTokenAuthenticationMethodDeviceId, actual.HardwareOathTokenAuthenticationMethodDeviceId)
		}

	}
}

func TestSegmentsForDirectoryAuthenticationMethodDeviceHardwareOathDeviceId(t *testing.T) {
	segments := DirectoryAuthenticationMethodDeviceHardwareOathDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryAuthenticationMethodDeviceHardwareOathDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
