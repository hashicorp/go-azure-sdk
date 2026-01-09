package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingWindowId{}

func TestNewMeSettingWindowID(t *testing.T) {
	id := NewMeSettingWindowID("windowsSettingId")

	if id.WindowsSettingId != "windowsSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsSettingId'", id.WindowsSettingId, "windowsSettingId")
	}
}

func TestFormatMeSettingWindowID(t *testing.T) {
	actual := NewMeSettingWindowID("windowsSettingId").ID()
	expected := "/me/settings/windows/windowsSettingId"
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
			Input: "/me/settings/windows/windowsSettingId",
			Expected: &MeSettingWindowId{
				WindowsSettingId: "windowsSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/windows/windowsSettingId/extra",
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
			Input: "/me/settings/windows/windowsSettingId",
			Expected: &MeSettingWindowId{
				WindowsSettingId: "windowsSettingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/windows/windowsSettingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD",
			Expected: &MeSettingWindowId{
				WindowsSettingId: "wInDoWsSeTtInGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/extra",
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
