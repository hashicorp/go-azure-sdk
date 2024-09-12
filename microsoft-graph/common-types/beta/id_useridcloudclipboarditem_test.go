package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCloudClipboardItemId{}

func TestNewUserIdCloudClipboardItemID(t *testing.T) {
	id := NewUserIdCloudClipboardItemID("userIdValue", "cloudClipboardItemIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.CloudClipboardItemId != "cloudClipboardItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudClipboardItemId'", id.CloudClipboardItemId, "cloudClipboardItemIdValue")
	}
}

func TestFormatUserIdCloudClipboardItemID(t *testing.T) {
	actual := NewUserIdCloudClipboardItemID("userIdValue", "cloudClipboardItemIdValue").ID()
	expected := "/users/userIdValue/cloudClipboard/items/cloudClipboardItemIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/cloudClipboard",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/cloudClipboard/items",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/cloudClipboard/items/cloudClipboardItemIdValue",
			Expected: &UserIdCloudClipboardItemId{
				UserId:               "userIdValue",
				CloudClipboardItemId: "cloudClipboardItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/cloudClipboard/items/cloudClipboardItemIdValue/extra",
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
			Input: "/users/userIdValue/cloudClipboard",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdClIpBoArD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/cloudClipboard/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdClIpBoArD/iTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/cloudClipboard/items/cloudClipboardItemIdValue",
			Expected: &UserIdCloudClipboardItemId{
				UserId:               "userIdValue",
				CloudClipboardItemId: "cloudClipboardItemIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/cloudClipboard/items/cloudClipboardItemIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdClIpBoArD/iTeMs/cLoUdClIpBoArDiTeMiDvAlUe",
			Expected: &UserIdCloudClipboardItemId{
				UserId:               "uSeRiDvAlUe",
				CloudClipboardItemId: "cLoUdClIpBoArDiTeMiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdClIpBoArD/iTeMs/cLoUdClIpBoArDiTeMiDvAlUe/extra",
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
