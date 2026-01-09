package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInsightSharedId{}

func TestNewMeInsightSharedID(t *testing.T) {
	id := NewMeInsightSharedID("sharedInsightId")

	if id.SharedInsightId != "sharedInsightId" {
		t.Fatalf("Expected %q but got %q for Segment 'SharedInsightId'", id.SharedInsightId, "sharedInsightId")
	}
}

func TestFormatMeInsightSharedID(t *testing.T) {
	actual := NewMeInsightSharedID("sharedInsightId").ID()
	expected := "/me/insights/shared/sharedInsightId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInsightSharedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInsightSharedId
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
			Input: "/me/insights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/insights/shared",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/insights/shared/sharedInsightId",
			Expected: &MeInsightSharedId{
				SharedInsightId: "sharedInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/insights/shared/sharedInsightId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInsightSharedID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SharedInsightId != v.Expected.SharedInsightId {
			t.Fatalf("Expected %q but got %q for SharedInsightId", v.Expected.SharedInsightId, actual.SharedInsightId)
		}

	}
}

func TestParseMeInsightSharedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInsightSharedId
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
			Input: "/me/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/insights/shared",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/sHaReD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/insights/shared/sharedInsightId",
			Expected: &MeInsightSharedId{
				SharedInsightId: "sharedInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/insights/shared/sharedInsightId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/sHaReD/sHaReDiNsIgHtId",
			Expected: &MeInsightSharedId{
				SharedInsightId: "sHaReDiNsIgHtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/sHaReD/sHaReDiNsIgHtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInsightSharedIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SharedInsightId != v.Expected.SharedInsightId {
			t.Fatalf("Expected %q but got %q for SharedInsightId", v.Expected.SharedInsightId, actual.SharedInsightId)
		}

	}
}

func TestSegmentsForMeInsightSharedId(t *testing.T) {
	segments := MeInsightSharedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInsightSharedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
