package meonlinemeetingrecordingcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingRecordingId{}

func TestNewMeOnlineMeetingRecordingID(t *testing.T) {
	id := NewMeOnlineMeetingRecordingID("onlineMeetingIdValue", "callRecordingIdValue")

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.CallRecordingId != "callRecordingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CallRecordingId'", id.CallRecordingId, "callRecordingIdValue")
	}
}

func TestFormatMeOnlineMeetingRecordingID(t *testing.T) {
	actual := NewMeOnlineMeetingRecordingID("onlineMeetingIdValue", "callRecordingIdValue").ID()
	expected := "/me/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingRecordingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingRecordingId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue/recordings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue",
			Expected: &MeOnlineMeetingRecordingId{
				OnlineMeetingId: "onlineMeetingIdValue",
				CallRecordingId: "callRecordingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingRecordingID(v.Input)
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

func TestParseMeOnlineMeetingRecordingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingRecordingId
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
			Input: "/me/onlineMeetings/onlineMeetingIdValue/recordings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEcOrDiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue",
			Expected: &MeOnlineMeetingRecordingId{
				OnlineMeetingId: "onlineMeetingIdValue",
				CallRecordingId: "callRecordingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/recordings/callRecordingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEcOrDiNgS/cAlLrEcOrDiNgIdVaLuE",
			Expected: &MeOnlineMeetingRecordingId{
				OnlineMeetingId: "oNlInEmEeTiNgIdVaLuE",
				CallRecordingId: "cAlLrEcOrDiNgIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/rEcOrDiNgS/cAlLrEcOrDiNgIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingRecordingIDInsensitively(v.Input)
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

func TestSegmentsForMeOnlineMeetingRecordingId(t *testing.T) {
	segments := MeOnlineMeetingRecordingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingRecordingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
