package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdRecordingId{}

func TestNewUserIdOnlineMeetingIdRecordingID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdRecordingID("userId", "onlineMeetingId", "callRecordingId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.CallRecordingId != "callRecordingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CallRecordingId'", id.CallRecordingId, "callRecordingId")
	}
}

func TestFormatUserIdOnlineMeetingIdRecordingID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdRecordingID("userId", "onlineMeetingId", "callRecordingId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId/recordings/callRecordingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdRecordingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdRecordingId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/recordings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/recordings/callRecordingId",
			Expected: &UserIdOnlineMeetingIdRecordingId{
				UserId:          "userId",
				OnlineMeetingId: "onlineMeetingId",
				CallRecordingId: "callRecordingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/recordings/callRecordingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdRecordingID(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.CallRecordingId != v.Expected.CallRecordingId {
			t.Fatalf("Expected %q but got %q for CallRecordingId", v.Expected.CallRecordingId, actual.CallRecordingId)
		}

	}
}

func TestParseUserIdOnlineMeetingIdRecordingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdRecordingId
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
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/recordings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEcOrDiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/recordings/callRecordingId",
			Expected: &UserIdOnlineMeetingIdRecordingId{
				UserId:          "userId",
				OnlineMeetingId: "onlineMeetingId",
				CallRecordingId: "callRecordingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/recordings/callRecordingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEcOrDiNgS/cAlLrEcOrDiNgId",
			Expected: &UserIdOnlineMeetingIdRecordingId{
				UserId:          "uSeRiD",
				OnlineMeetingId: "oNlInEmEeTiNgId",
				CallRecordingId: "cAlLrEcOrDiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEcOrDiNgS/cAlLrEcOrDiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdRecordingIDInsensitively(v.Input)
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

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.CallRecordingId != v.Expected.CallRecordingId {
			t.Fatalf("Expected %q but got %q for CallRecordingId", v.Expected.CallRecordingId, actual.CallRecordingId)
		}

	}
}

func TestSegmentsForUserIdOnlineMeetingIdRecordingId(t *testing.T) {
	segments := UserIdOnlineMeetingIdRecordingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdRecordingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
