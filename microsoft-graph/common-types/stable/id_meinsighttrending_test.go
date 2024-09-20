package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInsightTrendingId{}

func TestNewMeInsightTrendingID(t *testing.T) {
	id := NewMeInsightTrendingID("trendingId")

	if id.TrendingId != "trendingId" {
		t.Fatalf("Expected %q but got %q for Segment 'TrendingId'", id.TrendingId, "trendingId")
	}
}

func TestFormatMeInsightTrendingID(t *testing.T) {
	actual := NewMeInsightTrendingID("trendingId").ID()
	expected := "/me/insights/trending/trendingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInsightTrendingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInsightTrendingId
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
			Input: "/me/insights/trending",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/insights/trending/trendingId",
			Expected: &MeInsightTrendingId{
				TrendingId: "trendingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/insights/trending/trendingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInsightTrendingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TrendingId != v.Expected.TrendingId {
			t.Fatalf("Expected %q but got %q for TrendingId", v.Expected.TrendingId, actual.TrendingId)
		}

	}
}

func TestParseMeInsightTrendingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInsightTrendingId
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
			Input: "/me/insights/trending",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/tReNdInG",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/insights/trending/trendingId",
			Expected: &MeInsightTrendingId{
				TrendingId: "trendingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/insights/trending/trendingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/tReNdInG/tReNdInGiD",
			Expected: &MeInsightTrendingId{
				TrendingId: "tReNdInGiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNsIgHtS/tReNdInG/tReNdInGiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInsightTrendingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TrendingId != v.Expected.TrendingId {
			t.Fatalf("Expected %q but got %q for TrendingId", v.Expected.TrendingId, actual.TrendingId)
		}

	}
}

func TestSegmentsForMeInsightTrendingId(t *testing.T) {
	segments := MeInsightTrendingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInsightTrendingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
