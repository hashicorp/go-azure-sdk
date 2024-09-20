package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSettingWindowIdInstanceId{}

func TestNewUserIdSettingWindowIdInstanceID(t *testing.T) {
	id := NewUserIdSettingWindowIdInstanceID("userId", "windowsSettingId", "windowsSettingInstanceId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.WindowsSettingId != "windowsSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsSettingId'", id.WindowsSettingId, "windowsSettingId")
	}

	if id.WindowsSettingInstanceId != "windowsSettingInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'WindowsSettingInstanceId'", id.WindowsSettingInstanceId, "windowsSettingInstanceId")
	}
}

func TestFormatUserIdSettingWindowIdInstanceID(t *testing.T) {
	actual := NewUserIdSettingWindowIdInstanceID("userId", "windowsSettingId", "windowsSettingInstanceId").ID()
	expected := "/users/userId/settings/windows/windowsSettingId/instances/windowsSettingInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdSettingWindowIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdSettingWindowIdInstanceId
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
			Input: "/users/userId/settings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/settings/windows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/settings/windows/windowsSettingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/settings/windows/windowsSettingId/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/settings/windows/windowsSettingId/instances/windowsSettingInstanceId",
			Expected: &UserIdSettingWindowIdInstanceId{
				UserId:                   "userId",
				WindowsSettingId:         "windowsSettingId",
				WindowsSettingInstanceId: "windowsSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/settings/windows/windowsSettingId/instances/windowsSettingInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdSettingWindowIdInstanceID(v.Input)
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

		if actual.WindowsSettingId != v.Expected.WindowsSettingId {
			t.Fatalf("Expected %q but got %q for WindowsSettingId", v.Expected.WindowsSettingId, actual.WindowsSettingId)
		}

		if actual.WindowsSettingInstanceId != v.Expected.WindowsSettingInstanceId {
			t.Fatalf("Expected %q but got %q for WindowsSettingInstanceId", v.Expected.WindowsSettingInstanceId, actual.WindowsSettingInstanceId)
		}

	}
}

func TestParseUserIdSettingWindowIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdSettingWindowIdInstanceId
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
			Input: "/users/userId/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEtTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/settings/windows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEtTiNgS/wInDoWs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/settings/windows/windowsSettingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/settings/windows/windowsSettingId/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/settings/windows/windowsSettingId/instances/windowsSettingInstanceId",
			Expected: &UserIdSettingWindowIdInstanceId{
				UserId:                   "userId",
				WindowsSettingId:         "windowsSettingId",
				WindowsSettingInstanceId: "windowsSettingInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/settings/windows/windowsSettingId/instances/windowsSettingInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/iNsTaNcEs/wInDoWsSeTtInGiNsTaNcEiD",
			Expected: &UserIdSettingWindowIdInstanceId{
				UserId:                   "uSeRiD",
				WindowsSettingId:         "wInDoWsSeTtInGiD",
				WindowsSettingInstanceId: "wInDoWsSeTtInGiNsTaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/sEtTiNgS/wInDoWs/wInDoWsSeTtInGiD/iNsTaNcEs/wInDoWsSeTtInGiNsTaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdSettingWindowIdInstanceIDInsensitively(v.Input)
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

		if actual.WindowsSettingId != v.Expected.WindowsSettingId {
			t.Fatalf("Expected %q but got %q for WindowsSettingId", v.Expected.WindowsSettingId, actual.WindowsSettingId)
		}

		if actual.WindowsSettingInstanceId != v.Expected.WindowsSettingInstanceId {
			t.Fatalf("Expected %q but got %q for WindowsSettingInstanceId", v.Expected.WindowsSettingInstanceId, actual.WindowsSettingInstanceId)
		}

	}
}

func TestSegmentsForUserIdSettingWindowIdInstanceId(t *testing.T) {
	segments := UserIdSettingWindowIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdSettingWindowIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
