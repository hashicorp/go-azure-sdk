package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInsightTrendingId{}

func TestNewUserIdInsightTrendingID(t *testing.T) {
	id := NewUserIdInsightTrendingID("userIdValue", "trendingIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.TrendingId != "trendingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TrendingId'", id.TrendingId, "trendingIdValue")
	}
}

func TestFormatUserIdInsightTrendingID(t *testing.T) {
	actual := NewUserIdInsightTrendingID("userIdValue", "trendingIdValue").ID()
	expected := "/users/userIdValue/insights/trending/trendingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInsightTrendingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInsightTrendingId
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
			Input: "/users/userIdValue/insights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/insights/trending",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/insights/trending/trendingIdValue",
			Expected: &UserIdInsightTrendingId{
				UserId:     "userIdValue",
				TrendingId: "trendingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/insights/trending/trendingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInsightTrendingID(v.Input)
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

		if actual.TrendingId != v.Expected.TrendingId {
			t.Fatalf("Expected %q but got %q for TrendingId", v.Expected.TrendingId, actual.TrendingId)
		}

	}
}

func TestParseUserIdInsightTrendingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInsightTrendingId
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
			Input: "/users/userIdValue/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/insights/trending",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/tReNdInG",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/insights/trending/trendingIdValue",
			Expected: &UserIdInsightTrendingId{
				UserId:     "userIdValue",
				TrendingId: "trendingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/insights/trending/trendingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/tReNdInG/tReNdInGiDvAlUe",
			Expected: &UserIdInsightTrendingId{
				UserId:     "uSeRiDvAlUe",
				TrendingId: "tReNdInGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNsIgHtS/tReNdInG/tReNdInGiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInsightTrendingIDInsensitively(v.Input)
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

		if actual.TrendingId != v.Expected.TrendingId {
			t.Fatalf("Expected %q but got %q for TrendingId", v.Expected.TrendingId, actual.TrendingId)
		}

	}
}

func TestSegmentsForUserIdInsightTrendingId(t *testing.T) {
	segments := UserIdInsightTrendingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInsightTrendingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
