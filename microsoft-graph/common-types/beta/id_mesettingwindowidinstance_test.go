package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeSettingWindowIdInstanceId{}

func TestNewMeSettingWindowIdInstanceID(t *testing.T) {
	id := NewMeSettingWindowIdInstanceID("windowsSettingId", "windowsSettingInstanceId")

	if id.WindowsSettingId != "windowsSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsSettingId'", id.WindowsSettingId, "windowsSettingId")
	}

	if id.WindowsSettingInstanceId != "windowsSettingInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsSettingInstanceId'", id.WindowsSettingInstanceId, "windowsSettingInstanceId")
	}
}

func TestFormatMeSettingWindowIdInstanceID(t *testing.T) {
	actual := NewMeSettingWindowIdInstanceID("windowsSettingId", "windowsSettingInstanceId").ID()
	expected := "/me/settings/windows/windowsSettingId/instances/windowsSettingInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeSettingWindowIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSettingWindowIdInstanceId
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
			// Incomplete URI
			Input: "/me/settings/windows/windowsSettingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/windows/windowsSettingId/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/settings/windows/windowsSettingId/instances/windowsSettingInstanceId",
			Expected: &MeSettingWindowIdInstanceId{
				WindowsSettingId:         "windowsSettingId",
				WindowsSettingInstanceId: "windowsSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/windows/windowsSettingId/instances/windowsSettingInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSettingWindowIdInstanceID(v.Input)
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

		if actual.WindowsSettingInstanceId != v.Expected.WindowsSettingInstanceId {
			t.Fatalf("Expected %q but got %q for WindowsSettingInstanceId", v.Expected.WindowsSettingInstanceId, actual.WindowsSettingInstanceId)
		}

	}
}

func TestParseMeSettingWindowIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeSettingWindowIdInstanceId
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
			// Incomplete URI
			Input: "/me/settings/windows/windowsSettingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/settings/windows/windowsSettingId/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/settings/windows/windowsSettingId/instances/windowsSettingInstanceId",
			Expected: &MeSettingWindowIdInstanceId{
				WindowsSettingId:         "windowsSettingId",
				WindowsSettingInstanceId: "windowsSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/settings/windows/windowsSettingId/instances/windowsSettingInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/iNsTaNcEs/wInDoWsSeTtInGiNsTaNcEiD",
			Expected: &MeSettingWindowIdInstanceId{
				WindowsSettingId:         "wInDoWsSeTtInGiD",
				WindowsSettingInstanceId: "wInDoWsSeTtInGiNsTaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/iNsTaNcEs/wInDoWsSeTtInGiNsTaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeSettingWindowIdInstanceIDInsensitively(v.Input)
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

		if actual.WindowsSettingInstanceId != v.Expected.WindowsSettingInstanceId {
			t.Fatalf("Expected %q but got %q for WindowsSettingInstanceId", v.Expected.WindowsSettingInstanceId, actual.WindowsSettingInstanceId)
		}

	}
}

func TestSegmentsForMeSettingWindowIdInstanceId(t *testing.T) {
	segments := MeSettingWindowIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeSettingWindowIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
