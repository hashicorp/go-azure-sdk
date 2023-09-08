package useronlinemeetingrecording

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingRecordingId{}

func TestNewUserOnlineMeetingRecordingID(t *testing.T) {
	id := NewUserOnlineMeetingRecordingID("userIdValue", "onlineMeetingIdValue", "callRecordingIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.CallRecordingId != "callRecordingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CallRecordingId'", id.CallRecordingId, "callRecordingIdValue")
	}
}

func TestFormatUserOnlineMeetingRecordingID(t *testing.T) {
	actual := NewUserOnlineMeetingRecordingID("userIdValue", "onlineMeetingIdValue", "callRecordingIdValue").ID()
	expected := "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnlineMeetingRecordingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingRecordingId
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
			Input: "/users/userIdValue/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue",
			Expected: &UserOnlineMeetingRecordingId{
				UserId:          "userIdValue",
				OnlineMeetingId: "onlineMeetingIdValue",
				CallRecordingId: "callRecordingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingRecordingID(v.Input)
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

func TestParseUserOnlineMeetingRecordingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingRecordingId
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
			Input: "/users/userIdValue/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEcOrDiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue",
			Expected: &UserOnlineMeetingRecordingId{
				UserId:          "userIdValue",
				OnlineMeetingId: "onlineMeetingIdValue",
				CallRecordingId: "callRecordingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEcOrDiNgS/cAlLrEcOrDiNgIdVaLuE",
			Expected: &UserOnlineMeetingRecordingId{
				UserId:          "uSeRiDvAlUe",
				OnlineMeetingId: "oNlInEmEeTiNgIdVaLuE",
				CallRecordingId: "cAlLrEcOrDiNgIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEcOrDiNgS/cAlLrEcOrDiNgIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingRecordingIDInsensitively(v.Input)
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

func TestSegmentsForUserOnlineMeetingRecordingId(t *testing.T) {
	segments := UserOnlineMeetingRecordingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnlineMeetingRecordingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
