package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdRecordingId{}

func TestNewMeOnlineMeetingIdRecordingID(t *testing.T) {
	id := NewMeOnlineMeetingIdRecordingID("onlineMeetingId", "callRecordingId")

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.CallRecordingId != "callRecordingId" {
		t.Fatalf("Expected %q but got %q for Segment 'CallRecordingId'", id.CallRecordingId, "callRecordingId")
	}
}

func TestFormatMeOnlineMeetingIdRecordingID(t *testing.T) {
	actual := NewMeOnlineMeetingIdRecordingID("onlineMeetingId", "callRecordingId").ID()
	expected := "/me/onlineMeetings/onlineMeetingId/recordings/callRecordingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingIdRecordingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdRecordingId
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
			Input: "/me/onlineMeetings/onlineMeetingId/recordings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/recordings/callRecordingId",
			Expected: &MeOnlineMeetingIdRecordingId{
				OnlineMeetingId: "onlineMeetingId",
				CallRecordingId: "callRecordingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/recordings/callRecordingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdRecordingID(v.Input)
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

		if actual.CallRecordingId != v.Expected.CallRecordingId {
			t.Fatalf("Expected %q but got %q for CallRecordingId", v.Expected.CallRecordingId, actual.CallRecordingId)
		}

	}
}

func TestParseMeOnlineMeetingIdRecordingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdRecordingId
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
			Input: "/me/onlineMeetings/onlineMeetingId/recordings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEcOrDiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/recordings/callRecordingId",
			Expected: &MeOnlineMeetingIdRecordingId{
				OnlineMeetingId: "onlineMeetingId",
				CallRecordingId: "callRecordingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/recordings/callRecordingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEcOrDiNgS/cAlLrEcOrDiNgId",
			Expected: &MeOnlineMeetingIdRecordingId{
				OnlineMeetingId: "oNlInEmEeTiNgId",
				CallRecordingId: "cAlLrEcOrDiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/rEcOrDiNgS/cAlLrEcOrDiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdRecordingIDInsensitively(v.Input)
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

		if actual.CallRecordingId != v.Expected.CallRecordingId {
			t.Fatalf("Expected %q but got %q for CallRecordingId", v.Expected.CallRecordingId, actual.CallRecordingId)
		}

	}
}

func TestSegmentsForMeOnlineMeetingIdRecordingId(t *testing.T) {
	segments := MeOnlineMeetingIdRecordingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingIdRecordingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
