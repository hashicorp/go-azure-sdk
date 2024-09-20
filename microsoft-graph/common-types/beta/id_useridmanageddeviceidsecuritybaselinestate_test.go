package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdSecurityBaselineStateId{}

func TestNewUserIdManagedDeviceIdSecurityBaselineStateID(t *testing.T) {
	id := NewUserIdManagedDeviceIdSecurityBaselineStateID("userId", "managedDeviceId", "securityBaselineStateId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.SecurityBaselineStateId != "securityBaselineStateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineStateId'", id.SecurityBaselineStateId, "securityBaselineStateId")
	}
}

func TestFormatUserIdManagedDeviceIdSecurityBaselineStateID(t *testing.T) {
	actual := NewUserIdManagedDeviceIdSecurityBaselineStateID("userId", "managedDeviceId", "securityBaselineStateId").ID()
	expected := "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedDeviceIdSecurityBaselineStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdSecurityBaselineStateId
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
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Expected: &UserIdManagedDeviceIdSecurityBaselineStateId{
				UserId:                  "userId",
				ManagedDeviceId:         "managedDeviceId",
				SecurityBaselineStateId: "securityBaselineStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdSecurityBaselineStateID(v.Input)
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

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

	}
}

func TestParseUserIdManagedDeviceIdSecurityBaselineStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdSecurityBaselineStateId
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
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId",
			Expected: &UserIdManagedDeviceIdSecurityBaselineStateId{
				UserId:                  "userId",
				ManagedDeviceId:         "managedDeviceId",
				SecurityBaselineStateId: "securityBaselineStateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/securityBaselineStates/securityBaselineStateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId",
			Expected: &UserIdManagedDeviceIdSecurityBaselineStateId{
				UserId:                  "uSeRiD",
				ManagedDeviceId:         "mAnAgEdDeViCeId",
				SecurityBaselineStateId: "sEcUrItYbAsElInEsTaTeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdSecurityBaselineStateIDInsensitively(v.Input)
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

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

	}
}

func TestSegmentsForUserIdManagedDeviceIdSecurityBaselineStateId(t *testing.T) {
	segments := UserIdManagedDeviceIdSecurityBaselineStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedDeviceIdSecurityBaselineStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
