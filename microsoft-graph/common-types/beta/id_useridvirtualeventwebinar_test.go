package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdVirtualEventWebinarId{}

func TestNewUserIdVirtualEventWebinarID(t *testing.T) {
	id := NewUserIdVirtualEventWebinarID("userId", "virtualEventWebinarId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.VirtualEventWebinarId != "virtualEventWebinarId" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualEventWebinarId'", id.VirtualEventWebinarId, "virtualEventWebinarId")
	}
}

func TestFormatUserIdVirtualEventWebinarID(t *testing.T) {
	actual := NewUserIdVirtualEventWebinarID("userId", "virtualEventWebinarId").ID()
	expected := "/users/userId/virtualEvents/webinars/virtualEventWebinarId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdVirtualEventWebinarID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdVirtualEventWebinarId
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
			Input: "/users/userId/virtualEvents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/virtualEvents/webinars",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/virtualEvents/webinars/virtualEventWebinarId",
			Expected: &UserIdVirtualEventWebinarId{
				UserId:                "userId",
				VirtualEventWebinarId: "virtualEventWebinarId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/virtualEvents/webinars/virtualEventWebinarId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdVirtualEventWebinarID(v.Input)
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

		if actual.VirtualEventWebinarId != v.Expected.VirtualEventWebinarId {
			t.Fatalf("Expected %q but got %q for VirtualEventWebinarId", v.Expected.VirtualEventWebinarId, actual.VirtualEventWebinarId)
		}

	}
}

func TestParseUserIdVirtualEventWebinarIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdVirtualEventWebinarId
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
			Input: "/users/userId/virtualEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/vIrTuAlEvEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/virtualEvents/webinars",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/vIrTuAlEvEnTs/wEbInArS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/virtualEvents/webinars/virtualEventWebinarId",
			Expected: &UserIdVirtualEventWebinarId{
				UserId:                "userId",
				VirtualEventWebinarId: "virtualEventWebinarId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/virtualEvents/webinars/virtualEventWebinarId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/vIrTuAlEvEnTs/wEbInArS/vIrTuAlEvEnTwEbInArId",
			Expected: &UserIdVirtualEventWebinarId{
				UserId:                "uSeRiD",
				VirtualEventWebinarId: "vIrTuAlEvEnTwEbInArId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/vIrTuAlEvEnTs/wEbInArS/vIrTuAlEvEnTwEbInArId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdVirtualEventWebinarIDInsensitively(v.Input)
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

		if actual.VirtualEventWebinarId != v.Expected.VirtualEventWebinarId {
			t.Fatalf("Expected %q but got %q for VirtualEventWebinarId", v.Expected.VirtualEventWebinarId, actual.VirtualEventWebinarId)
		}

	}
}

func TestSegmentsForUserIdVirtualEventWebinarId(t *testing.T) {
	segments := UserIdVirtualEventWebinarId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdVirtualEventWebinarId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
