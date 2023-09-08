package meowneddevice

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOwnedDeviceId{}

func TestNewMeOwnedDeviceID(t *testing.T) {
	id := NewMeOwnedDeviceID("directoryObjectIdValue")

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatMeOwnedDeviceID(t *testing.T) {
	actual := NewMeOwnedDeviceID("directoryObjectIdValue").ID()
	expected := "/me/ownedDevices/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOwnedDeviceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOwnedDeviceId
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
			Input: "/me/ownedDevices",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/ownedDevices/directoryObjectIdValue",
			Expected: &MeOwnedDeviceId{
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/ownedDevices/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOwnedDeviceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseMeOwnedDeviceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOwnedDeviceId
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
			Input: "/me/ownedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oWnEdDeViCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/ownedDevices/directoryObjectIdValue",
			Expected: &MeOwnedDeviceId{
				DirectoryObjectId: "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/ownedDevices/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oWnEdDeViCeS/dIrEcToRyObJeCtIdVaLuE",
			Expected: &MeOwnedDeviceId{
				DirectoryObjectId: "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oWnEdDeViCeS/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOwnedDeviceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForMeOwnedDeviceId(t *testing.T) {
	segments := MeOwnedDeviceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOwnedDeviceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
