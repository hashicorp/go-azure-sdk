package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdTranscriptId{}

func TestNewUserIdOnlineMeetingIdTranscriptID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdTranscriptID("userIdValue", "onlineMeetingIdValue", "callTranscriptIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.CallTranscriptId != "callTranscriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CallTranscriptId'", id.CallTranscriptId, "callTranscriptIdValue")
	}
}

func TestFormatUserIdOnlineMeetingIdTranscriptID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdTranscriptID("userIdValue", "onlineMeetingIdValue", "callTranscriptIdValue").ID()
	expected := "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdTranscriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdTranscriptId
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
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue",
			Expected: &UserIdOnlineMeetingIdTranscriptId{
				UserId:           "userIdValue",
				OnlineMeetingId:  "onlineMeetingIdValue",
				CallTranscriptId: "callTranscriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdTranscriptID(v.Input)
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

		if actual.CallTranscriptId != v.Expected.CallTranscriptId {
			t.Fatalf("Expected %q but got %q for CallTranscriptId", v.Expected.CallTranscriptId, actual.CallTranscriptId)
		}

	}
}

func TestParseUserIdOnlineMeetingIdTranscriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdTranscriptId
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
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/tRaNsCrIpTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue",
			Expected: &UserIdOnlineMeetingIdTranscriptId{
				UserId:           "userIdValue",
				OnlineMeetingId:  "onlineMeetingIdValue",
				CallTranscriptId: "callTranscriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/tRaNsCrIpTs/cAlLtRaNsCrIpTiDvAlUe",
			Expected: &UserIdOnlineMeetingIdTranscriptId{
				UserId:           "uSeRiDvAlUe",
				OnlineMeetingId:  "oNlInEmEeTiNgIdVaLuE",
				CallTranscriptId: "cAlLtRaNsCrIpTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/tRaNsCrIpTs/cAlLtRaNsCrIpTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdTranscriptIDInsensitively(v.Input)
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

		if actual.CallTranscriptId != v.Expected.CallTranscriptId {
			t.Fatalf("Expected %q but got %q for CallTranscriptId", v.Expected.CallTranscriptId, actual.CallTranscriptId)
		}

	}
}

func TestSegmentsForUserIdOnlineMeetingIdTranscriptId(t *testing.T) {
	segments := UserIdOnlineMeetingIdTranscriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdTranscriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
