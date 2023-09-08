package meonlinemeetingtranscript

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingTranscriptId{}

func TestNewMeOnlineMeetingTranscriptID(t *testing.T) {
	id := NewMeOnlineMeetingTranscriptID("onlineMeetingIdValue", "callTranscriptIdValue")

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.CallTranscriptId != "callTranscriptIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CallTranscriptId'", id.CallTranscriptId, "callTranscriptIdValue")
	}
}

func TestFormatMeOnlineMeetingTranscriptID(t *testing.T) {
	actual := NewMeOnlineMeetingTranscriptID("onlineMeetingIdValue", "callTranscriptIdValue").ID()
	expected := "/me/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingTranscriptID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingTranscriptId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/transcripts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue",
			Expected: &MeOnlineMeetingTranscriptId{
				OnlineMeetingId:  "onlineMeetingIdValue",
				CallTranscriptId: "callTranscriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingTranscriptID(v.Input)
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

func TestParseMeOnlineMeetingTranscriptIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingTranscriptId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/transcripts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/tRaNsCrIpTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue",
			Expected: &MeOnlineMeetingTranscriptId{
				OnlineMeetingId:  "onlineMeetingIdValue",
				CallTranscriptId: "callTranscriptIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/transcripts/callTranscriptIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/tRaNsCrIpTs/cAlLtRaNsCrIpTiDvAlUe",
			Expected: &MeOnlineMeetingTranscriptId{
				OnlineMeetingId:  "oNlInEmEeTiNgIdVaLuE",
				CallTranscriptId: "cAlLtRaNsCrIpTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/tRaNsCrIpTs/cAlLtRaNsCrIpTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingTranscriptIDInsensitively(v.Input)
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

func TestSegmentsForMeOnlineMeetingTranscriptId(t *testing.T) {
	segments := MeOnlineMeetingTranscriptId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingTranscriptId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
