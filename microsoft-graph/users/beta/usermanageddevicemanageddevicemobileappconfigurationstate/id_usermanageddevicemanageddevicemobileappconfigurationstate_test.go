package usermanageddevicemanageddevicemobileappconfigurationstate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{}

func TestNewUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	id := NewUserManagedDeviceManagedDeviceMobileAppConfigurationStateID("userIdValue", "managedDeviceIdValue", "managedDeviceMobileAppConfigurationStateIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.ManagedDeviceMobileAppConfigurationStateId != "managedDeviceMobileAppConfigurationStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceMobileAppConfigurationStateId'", id.ManagedDeviceMobileAppConfigurationStateId, "managedDeviceMobileAppConfigurationStateIdValue")
	}
}

func TestFormatUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	actual := NewUserManagedDeviceManagedDeviceMobileAppConfigurationStateID("userIdValue", "managedDeviceIdValue", "managedDeviceMobileAppConfigurationStateIdValue").ID()
	expected := "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceManagedDeviceMobileAppConfigurationStateId
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
			Input: "/users/userIdValue/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue",
			Expected: &UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{
				UserId:          "userIdValue",
				ManagedDeviceId: "managedDeviceIdValue",
				ManagedDeviceMobileAppConfigurationStateId: "managedDeviceMobileAppConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateID(v.Input)
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

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.ManagedDeviceMobileAppConfigurationStateId != v.Expected.ManagedDeviceMobileAppConfigurationStateId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceMobileAppConfigurationStateId", v.Expected.ManagedDeviceMobileAppConfigurationStateId, actual.ManagedDeviceMobileAppConfigurationStateId)
		}

	}
}

func TestParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceManagedDeviceMobileAppConfigurationStateId
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
			Input: "/users/userIdValue/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue",
			Expected: &UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{
				UserId:          "userIdValue",
				ManagedDeviceId: "managedDeviceIdValue",
				ManagedDeviceMobileAppConfigurationStateId: "managedDeviceMobileAppConfigurationStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/managedDeviceMobileAppConfigurationStates/managedDeviceMobileAppConfigurationStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe",
			Expected: &UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{
				UserId:          "uSeRiDvAlUe",
				ManagedDeviceId: "mAnAgEdDeViCeIdVaLuE",
				ManagedDeviceMobileAppConfigurationStateId: "mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEs/mAnAgEdDeViCeMoBiLeApPcOnFiGuRaTiOnStAtEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceManagedDeviceMobileAppConfigurationStateIDInsensitively(v.Input)
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

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.ManagedDeviceMobileAppConfigurationStateId != v.Expected.ManagedDeviceMobileAppConfigurationStateId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceMobileAppConfigurationStateId", v.Expected.ManagedDeviceMobileAppConfigurationStateId, actual.ManagedDeviceMobileAppConfigurationStateId)
		}

	}
}

func TestSegmentsForUserManagedDeviceManagedDeviceMobileAppConfigurationStateId(t *testing.T) {
	segments := UserManagedDeviceManagedDeviceMobileAppConfigurationStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserManagedDeviceManagedDeviceMobileAppConfigurationStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
