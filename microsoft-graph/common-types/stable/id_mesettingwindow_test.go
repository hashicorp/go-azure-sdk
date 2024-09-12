package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingWindowId{}

func TestNewMeSettingWindowID(t *testing.T) {
	id := NewMeSettingWindowID("windowsSettingIdValue")

	if id.WindowsSettingId != "windowsSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsSettingId'", id.WindowsSettingId, "windowsSettingIdValue")
	}
}

func TestFormatMeSettingWindowID(t *testing.T) {
	actual := NewMeSettingWindowID("windowsSettingIdValue").ID()
	expected := "/me/settings/windows/windowsSettingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeSettingWindowID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSettingWindowId
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
			Input: "/me/settings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/windows",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/settings/windows/windowsSettingIdValue",
			Expected: &MeSettingWindowId{
				WindowsSettingId: "windowsSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/windows/windowsSettingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSettingWindowID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsSettingId != v.Expected.WindowsSettingId {
			t.Fatalf("Expected %q but got %q for WindowsSettingId", v.Expected.WindowsSettingId, actual.WindowsSettingId)
		}

	}
}

func TestParseMeSettingWindowIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSettingWindowId
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
			Input: "/me/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/windows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/settings/windows/windowsSettingIdValue",
			Expected: &MeSettingWindowId{
				WindowsSettingId: "windowsSettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/windows/windowsSettingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiDvAlUe",
			Expected: &MeSettingWindowId{
				WindowsSettingId: "wInDoWsSeTtInGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSettingWindowIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WindowsSettingId != v.Expected.WindowsSettingId {
			t.Fatalf("Expected %q but got %q for WindowsSettingId", v.Expected.WindowsSettingId, actual.WindowsSettingId)
		}

	}
}

func TestSegmentsForMeSettingWindowId(t *testing.T) {
	segments := MeSettingWindowId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeSettingWindowId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
