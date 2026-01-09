package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCloudClipboardItemId{}

func TestNewUserIdCloudClipboardItemID(t *testing.T) {
	id := NewUserIdCloudClipboardItemID("userId", "cloudClipboardItemId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.CloudClipboardItemId != "cloudClipboardItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudClipboardItemId'", id.CloudClipboardItemId, "cloudClipboardItemId")
	}
}

func TestFormatUserIdCloudClipboardItemID(t *testing.T) {
	actual := NewUserIdCloudClipboardItemID("userId", "cloudClipboardItemId").ID()
	expected := "/users/userId/cloudClipboard/items/cloudClipboardItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdCloudClipboardItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCloudClipboardItemId
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
			Input: "/users/userId/cloudClipboard",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/cloudClipboard/items",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/cloudClipboard/items/cloudClipboardItemId",
			Expected: &UserIdCloudClipboardItemId{
				UserId:               "userId",
				CloudClipboardItemId: "cloudClipboardItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/cloudClipboard/items/cloudClipboardItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCloudClipboardItemID(v.Input)
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

		if actual.CloudClipboardItemId != v.Expected.CloudClipboardItemId {
			t.Fatalf("Expected %q but got %q for CloudClipboardItemId", v.Expected.CloudClipboardItemId, actual.CloudClipboardItemId)
		}

	}
}

func TestParseUserIdCloudClipboardItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCloudClipboardItemId
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
			Input: "/users/userId/cloudClipboard",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdClIpBoArD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/cloudClipboard/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdClIpBoArD/iTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/cloudClipboard/items/cloudClipboardItemId",
			Expected: &UserIdCloudClipboardItemId{
				UserId:               "userId",
				CloudClipboardItemId: "cloudClipboardItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/cloudClipboard/items/cloudClipboardItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdClIpBoArD/iTeMs/cLoUdClIpBoArDiTeMiD",
			Expected: &UserIdCloudClipboardItemId{
				UserId:               "uSeRiD",
				CloudClipboardItemId: "cLoUdClIpBoArDiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdClIpBoArD/iTeMs/cLoUdClIpBoArDiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCloudClipboardItemIDInsensitively(v.Input)
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

		if actual.CloudClipboardItemId != v.Expected.CloudClipboardItemId {
			t.Fatalf("Expected %q but got %q for CloudClipboardItemId", v.Expected.CloudClipboardItemId, actual.CloudClipboardItemId)
		}

	}
}

func TestSegmentsForUserIdCloudClipboardItemId(t *testing.T) {
	segments := UserIdCloudClipboardItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdCloudClipboardItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
