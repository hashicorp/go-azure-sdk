package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionBitlockerRecoveryKeyId{}

func TestNewUserIdInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	id := NewUserIdInformationProtectionBitlockerRecoveryKeyID("userIdValue", "bitlockerRecoveryKeyIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.BitlockerRecoveryKeyId != "bitlockerRecoveryKeyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BitlockerRecoveryKeyId'", id.BitlockerRecoveryKeyId, "bitlockerRecoveryKeyIdValue")
	}
}

func TestFormatUserIdInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	actual := NewUserIdInformationProtectionBitlockerRecoveryKeyID("userIdValue", "bitlockerRecoveryKeyIdValue").ID()
	expected := "/users/userIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionBitlockerRecoveryKeyId
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
			Input: "/users/userIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue",
			Expected: &UserIdInformationProtectionBitlockerRecoveryKeyId{
				UserId:                 "userIdValue",
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionBitlockerRecoveryKeyID(v.Input)
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

		if actual.BitlockerRecoveryKeyId != v.Expected.BitlockerRecoveryKeyId {
			t.Fatalf("Expected %q but got %q for BitlockerRecoveryKeyId", v.Expected.BitlockerRecoveryKeyId, actual.BitlockerRecoveryKeyId)
		}

	}
}

func TestParseUserIdInformationProtectionBitlockerRecoveryKeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionBitlockerRecoveryKeyId
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
			Input: "/users/userIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue",
			Expected: &UserIdInformationProtectionBitlockerRecoveryKeyId{
				UserId:                 "userIdValue",
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiDvAlUe",
			Expected: &UserIdInformationProtectionBitlockerRecoveryKeyId{
				UserId:                 "uSeRiDvAlUe",
				BitlockerRecoveryKeyId: "bItLoCkErReCoVeRyKeYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionBitlockerRecoveryKeyIDInsensitively(v.Input)
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

		if actual.BitlockerRecoveryKeyId != v.Expected.BitlockerRecoveryKeyId {
			t.Fatalf("Expected %q but got %q for BitlockerRecoveryKeyId", v.Expected.BitlockerRecoveryKeyId, actual.BitlockerRecoveryKeyId)
		}

	}
}

func TestSegmentsForUserIdInformationProtectionBitlockerRecoveryKeyId(t *testing.T) {
	segments := UserIdInformationProtectionBitlockerRecoveryKeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInformationProtectionBitlockerRecoveryKeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
