package mewindowsinformationprotectiondeviceregistration

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeWindowsInformationProtectionDeviceRegistrationId{}

func TestNewMeWindowsInformationProtectionDeviceRegistrationID(t *testing.T) {
	id := NewMeWindowsInformationProtectionDeviceRegistrationID("windowsInformationProtectionDeviceRegistrationIdValue")

	if id.WindowsInformationProtectionDeviceRegistrationId != "windowsInformationProtectionDeviceRegistrationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsInformationProtectionDeviceRegistrationId'", id.WindowsInformationProtectionDeviceRegistrationId, "windowsInformationProtectionDeviceRegistrationIdValue")
	}
}

func TestFormatMeWindowsInformationProtectionDeviceRegistrationID(t *testing.T) {
	actual := NewMeWindowsInformationProtectionDeviceRegistrationID("windowsInformationProtectionDeviceRegistrationIdValue").ID()
	expected := "/me/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeWindowsInformationProtectionDeviceRegistrationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeWindowsInformationProtectionDeviceRegistrationId
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
			Input: "/me/windowsInformationProtectionDeviceRegistrations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationIdValue",
			Expected: &MeWindowsInformationProtectionDeviceRegistrationId{
				WindowsInformationProtectionDeviceRegistrationId: "windowsInformationProtectionDeviceRegistrationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeWindowsInformationProtectionDeviceRegistrationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsInformationProtectionDeviceRegistrationId != v.Expected.WindowsInformationProtectionDeviceRegistrationId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionDeviceRegistrationId", v.Expected.WindowsInformationProtectionDeviceRegistrationId, actual.WindowsInformationProtectionDeviceRegistrationId)
		}

	}
}

func TestParseMeWindowsInformationProtectionDeviceRegistrationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeWindowsInformationProtectionDeviceRegistrationId
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
			Input: "/me/windowsInformationProtectionDeviceRegistrations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationIdValue",
			Expected: &MeWindowsInformationProtectionDeviceRegistrationId{
				WindowsInformationProtectionDeviceRegistrationId: "windowsInformationProtectionDeviceRegistrationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/windowsInformationProtectionDeviceRegistrations/windowsInformationProtectionDeviceRegistrationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNs/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNiDvAlUe",
			Expected: &MeWindowsInformationProtectionDeviceRegistrationId{
				WindowsInformationProtectionDeviceRegistrationId: "wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNs/wInDoWsInFoRmAtIoNpRoTeCtIoNdEvIcErEgIsTrAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeWindowsInformationProtectionDeviceRegistrationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsInformationProtectionDeviceRegistrationId != v.Expected.WindowsInformationProtectionDeviceRegistrationId {
			t.Fatalf("Expected %q but got %q for WindowsInformationProtectionDeviceRegistrationId", v.Expected.WindowsInformationProtectionDeviceRegistrationId, actual.WindowsInformationProtectionDeviceRegistrationId)
		}

	}
}

func TestSegmentsForMeWindowsInformationProtectionDeviceRegistrationId(t *testing.T) {
	segments := MeWindowsInformationProtectionDeviceRegistrationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeWindowsInformationProtectionDeviceRegistrationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
