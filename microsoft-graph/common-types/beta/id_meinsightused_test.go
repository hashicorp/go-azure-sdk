package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInsightUsedId{}

func TestNewMeInsightUsedID(t *testing.T) {
	id := NewMeInsightUsedID("usedInsightId")

	if id.UsedInsightId != "usedInsightId" {
		t.Fatalf("Expected %q but got %q for Segment 'UsedInsightId'", id.UsedInsightId, "usedInsightId")
	}
}

func TestFormatMeInsightUsedID(t *testing.T) {
	actual := NewMeInsightUsedID("usedInsightId").ID()
	expected := "/me/insights/used/usedInsightId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInsightUsedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInsightUsedId
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
			Input: "/me/insights/used",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/insights/used/usedInsightId",
			Expected: &MeInsightUsedId{
				UsedInsightId: "usedInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/insights/used/usedInsightId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInsightUsedID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UsedInsightId != v.Expected.UsedInsightId {
			t.Fatalf("Expected %q but got %q for UsedInsightId", v.Expected.UsedInsightId, actual.UsedInsightId)
		}

	}
}

func TestParseMeInsightUsedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInsightUsedId
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
			Input: "/me/insights/used",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/uSeD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/insights/used/usedInsightId",
			Expected: &MeInsightUsedId{
				UsedInsightId: "usedInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/insights/used/usedInsightId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/uSeD/uSeDiNsIgHtId",
			Expected: &MeInsightUsedId{
				UsedInsightId: "uSeDiNsIgHtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/uSeD/uSeDiNsIgHtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInsightUsedIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UsedInsightId != v.Expected.UsedInsightId {
			t.Fatalf("Expected %q but got %q for UsedInsightId", v.Expected.UsedInsightId, actual.UsedInsightId)
		}

	}
}

func TestSegmentsForMeInsightUsedId(t *testing.T) {
	segments := MeInsightUsedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInsightUsedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
