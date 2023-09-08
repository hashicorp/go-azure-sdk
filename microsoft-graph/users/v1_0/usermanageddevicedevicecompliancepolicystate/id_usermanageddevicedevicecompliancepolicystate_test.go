package usermanageddevicedevicecompliancepolicystate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceDeviceCompliancePolicyStateId{}

func TestNewUserManagedDeviceDeviceCompliancePolicyStateID(t *testing.T) {
	id := NewUserManagedDeviceDeviceCompliancePolicyStateID("userIdValue", "managedDeviceIdValue", "deviceCompliancePolicyStateIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.DeviceCompliancePolicyStateId != "deviceCompliancePolicyStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceCompliancePolicyStateId'", id.DeviceCompliancePolicyStateId, "deviceCompliancePolicyStateIdValue")
	}
}

func TestFormatUserManagedDeviceDeviceCompliancePolicyStateID(t *testing.T) {
	actual := NewUserManagedDeviceDeviceCompliancePolicyStateID("userIdValue", "managedDeviceIdValue", "deviceCompliancePolicyStateIdValue").ID()
	expected := "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserManagedDeviceDeviceCompliancePolicyStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceDeviceCompliancePolicyStateId
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
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue",
			Expected: &UserManagedDeviceDeviceCompliancePolicyStateId{
				UserId:                        "userIdValue",
				ManagedDeviceId:               "managedDeviceIdValue",
				DeviceCompliancePolicyStateId: "deviceCompliancePolicyStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceDeviceCompliancePolicyStateID(v.Input)
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

		if actual.DeviceCompliancePolicyStateId != v.Expected.DeviceCompliancePolicyStateId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyStateId", v.Expected.DeviceCompliancePolicyStateId, actual.DeviceCompliancePolicyStateId)
		}

	}
}

func TestParseUserManagedDeviceDeviceCompliancePolicyStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserManagedDeviceDeviceCompliancePolicyStateId
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
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOmPlIaNcEpOlIcYsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue",
			Expected: &UserManagedDeviceDeviceCompliancePolicyStateId{
				UserId:                        "userIdValue",
				ManagedDeviceId:               "managedDeviceIdValue",
				DeviceCompliancePolicyStateId: "deviceCompliancePolicyStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/managedDevices/managedDeviceIdValue/deviceCompliancePolicyStates/deviceCompliancePolicyStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOmPlIaNcEpOlIcYsTaTeS/dEvIcEcOmPlIaNcEpOlIcYsTaTeIdVaLuE",
			Expected: &UserManagedDeviceDeviceCompliancePolicyStateId{
				UserId:                        "uSeRiDvAlUe",
				ManagedDeviceId:               "mAnAgEdDeViCeIdVaLuE",
				DeviceCompliancePolicyStateId: "dEvIcEcOmPlIaNcEpOlIcYsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/dEvIcEcOmPlIaNcEpOlIcYsTaTeS/dEvIcEcOmPlIaNcEpOlIcYsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserManagedDeviceDeviceCompliancePolicyStateIDInsensitively(v.Input)
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

		if actual.DeviceCompliancePolicyStateId != v.Expected.DeviceCompliancePolicyStateId {
			t.Fatalf("Expected %q but got %q for DeviceCompliancePolicyStateId", v.Expected.DeviceCompliancePolicyStateId, actual.DeviceCompliancePolicyStateId)
		}

	}
}

func TestSegmentsForUserManagedDeviceDeviceCompliancePolicyStateId(t *testing.T) {
	segments := UserManagedDeviceDeviceCompliancePolicyStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserManagedDeviceDeviceCompliancePolicyStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
