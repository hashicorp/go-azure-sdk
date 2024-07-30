package informationprotectionbitlockerrecoverykey

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionBitlockerRecoveryKeyId{}

func TestNewMeInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	id := NewMeInformationProtectionBitlockerRecoveryKeyID("bitlockerRecoveryKeyIdValue")

	if id.BitlockerRecoveryKeyId != "bitlockerRecoveryKeyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BitlockerRecoveryKeyId'", id.BitlockerRecoveryKeyId, "bitlockerRecoveryKeyIdValue")
	}
}

func TestFormatMeInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	actual := NewMeInformationProtectionBitlockerRecoveryKeyID("bitlockerRecoveryKeyIdValue").ID()
	expected := "/me/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInformationProtectionBitlockerRecoveryKeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionBitlockerRecoveryKeyId
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
			Input: "/me/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue",
			Expected: &MeInformationProtectionBitlockerRecoveryKeyId{
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionBitlockerRecoveryKeyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BitlockerRecoveryKeyId != v.Expected.BitlockerRecoveryKeyId {
			t.Fatalf("Expected %q but got %q for BitlockerRecoveryKeyId", v.Expected.BitlockerRecoveryKeyId, actual.BitlockerRecoveryKeyId)
		}

	}
}

func TestParseMeInformationProtectionBitlockerRecoveryKeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionBitlockerRecoveryKeyId
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
			Input: "/me/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/bitlocker",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/bitlocker/recoveryKeys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue",
			Expected: &MeInformationProtectionBitlockerRecoveryKeyId{
				BitlockerRecoveryKeyId: "bitlockerRecoveryKeyIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/bitlocker/recoveryKeys/bitlockerRecoveryKeyIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiDvAlUe",
			Expected: &MeInformationProtectionBitlockerRecoveryKeyId{
				BitlockerRecoveryKeyId: "bItLoCkErReCoVeRyKeYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/bItLoCkEr/rEcOvErYkEyS/bItLoCkErReCoVeRyKeYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionBitlockerRecoveryKeyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BitlockerRecoveryKeyId != v.Expected.BitlockerRecoveryKeyId {
			t.Fatalf("Expected %q but got %q for BitlockerRecoveryKeyId", v.Expected.BitlockerRecoveryKeyId, actual.BitlockerRecoveryKeyId)
		}

	}
}

func TestSegmentsForMeInformationProtectionBitlockerRecoveryKeyId(t *testing.T) {
	segments := MeInformationProtectionBitlockerRecoveryKeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInformationProtectionBitlockerRecoveryKeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
