package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingId{}

func TestNewMeOnlineMeetingID(t *testing.T) {
	id := NewMeOnlineMeetingID("onlineMeetingId")

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}
}

func TestFormatMeOnlineMeetingID(t *testing.T) {
	actual := NewMeOnlineMeetingID("onlineMeetingId").ID()
	expected := "/me/onlineMeetings/onlineMeetingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingId
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
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId",
			Expected: &MeOnlineMeetingId{
				OnlineMeetingId: "onlineMeetingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingID(v.Input)
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

	}
}

func TestParseMeOnlineMeetingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingId
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
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId",
			Expected: &MeOnlineMeetingId{
				OnlineMeetingId: "onlineMeetingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Expected: &MeOnlineMeetingId{
				OnlineMeetingId: "oNlInEmEeTiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeOnlineMeetingId(t *testing.T) {
	segments := MeOnlineMeetingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
