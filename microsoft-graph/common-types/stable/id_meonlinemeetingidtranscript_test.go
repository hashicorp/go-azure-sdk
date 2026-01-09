package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdTranscriptId{}

func TestNewMeOnlineMeetingIdTranscriptID(t *testing.T) {
	id := NewMeOnlineMeetingIdTranscriptID("onlineMeetingId", "callTranscriptId")

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.CallTranscriptId != "callTranscriptId" {
		t.Fatalf("Expected %q but got %q for Segment 'CallTranscriptId'", id.CallTranscriptId, "callTranscriptId")
	}
}

func TestFormatMeOnlineMeetingIdTranscriptID(t *testing.T) {
	actual := NewMeOnlineMeetingIdTranscriptID("onlineMeetingId", "callTranscriptId").ID()
	expected := "/me/onlineMeetings/onlineMeetingId/transcripts/callTranscriptId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingIdTranscriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdTranscriptId
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
			Input: "/me/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/transcripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/transcripts/callTranscriptId",
			Expected: &MeOnlineMeetingIdTranscriptId{
				OnlineMeetingId:  "onlineMeetingId",
				CallTranscriptId: "callTranscriptId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/transcripts/callTranscriptId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdTranscriptID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.CallTranscriptId != v.Expected.CallTranscriptId {
			t.Fatalf("Expected %q but got %q for CallTranscriptId", v.Expected.CallTranscriptId, actual.CallTranscriptId)
		}

	}
}

func TestParseMeOnlineMeetingIdTranscriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdTranscriptId
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
			Input: "/me/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/transcripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/tRaNsCrIpTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/transcripts/callTranscriptId",
			Expected: &MeOnlineMeetingIdTranscriptId{
				OnlineMeetingId:  "onlineMeetingId",
				CallTranscriptId: "callTranscriptId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/transcripts/callTranscriptId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/tRaNsCrIpTs/cAlLtRaNsCrIpTiD",
			Expected: &MeOnlineMeetingIdTranscriptId{
				OnlineMeetingId:  "oNlInEmEeTiNgId",
				CallTranscriptId: "cAlLtRaNsCrIpTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/tRaNsCrIpTs/cAlLtRaNsCrIpTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdTranscriptIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.CallTranscriptId != v.Expected.CallTranscriptId {
			t.Fatalf("Expected %q but got %q for CallTranscriptId", v.Expected.CallTranscriptId, actual.CallTranscriptId)
		}

	}
}

func TestSegmentsForMeOnlineMeetingIdTranscriptId(t *testing.T) {
	segments := MeOnlineMeetingIdTranscriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingIdTranscriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
