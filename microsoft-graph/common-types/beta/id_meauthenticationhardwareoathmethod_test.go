package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAuthenticationHardwareOathMethodId{}

func TestNewMeAuthenticationHardwareOathMethodID(t *testing.T) {
	id := NewMeAuthenticationHardwareOathMethodID("hardwareOathAuthenticationMethodId")

	if id.HardwareOathAuthenticationMethodId != "hardwareOathAuthenticationMethodId" {
		t.Fatalf("Expected %q but got %q for Segment 'HardwareOathAuthenticationMethodId'", id.HardwareOathAuthenticationMethodId, "hardwareOathAuthenticationMethodId")
	}
}

func TestFormatMeAuthenticationHardwareOathMethodID(t *testing.T) {
	actual := NewMeAuthenticationHardwareOathMethodID("hardwareOathAuthenticationMethodId").ID()
	expected := "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAuthenticationHardwareOathMethodID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationHardwareOathMethodId
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
			// Valid URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Expected: &MeAuthenticationHardwareOathMethodId{
				HardwareOathAuthenticationMethodId: "hardwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationHardwareOathMethodID(v.Input)
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

	}
}

func TestParseMeAuthenticationHardwareOathMethodIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAuthenticationHardwareOathMethodId
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
			// Valid URI
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId",
			Expected: &MeAuthenticationHardwareOathMethodId{
				HardwareOathAuthenticationMethodId: "hardwareOathAuthenticationMethodId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/authentication/hardwareOathMethods/hardwareOathAuthenticationMethodId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			Expected: &MeAuthenticationHardwareOathMethodId{
				HardwareOathAuthenticationMethodId: "hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aUtHeNtIcAtIoN/hArDwArEoAtHmEtHoDs/hArDwArEoAtHaUtHeNtIcAtIoNmEtHoDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAuthenticationHardwareOathMethodIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeAuthenticationHardwareOathMethodId(t *testing.T) {
	segments := MeAuthenticationHardwareOathMethodId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAuthenticationHardwareOathMethodId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
