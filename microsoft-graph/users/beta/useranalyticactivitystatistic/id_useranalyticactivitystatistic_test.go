package useranalyticactivitystatistic

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAnalyticActivityStatisticId{}

func TestNewUserAnalyticActivityStatisticID(t *testing.T) {
	id := NewUserAnalyticActivityStatisticID("userIdValue", "activityStatisticsIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ActivityStatisticsId != "activityStatisticsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityStatisticsId'", id.ActivityStatisticsId, "activityStatisticsIdValue")
	}
}

func TestFormatUserAnalyticActivityStatisticID(t *testing.T) {
	actual := NewUserAnalyticActivityStatisticID("userIdValue", "activityStatisticsIdValue").ID()
	expected := "/users/userIdValue/analytics/activityStatistics/activityStatisticsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAnalyticActivityStatisticID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAnalyticActivityStatisticId
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
			Input: "/users/userIdValue/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/analytics/activityStatistics",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/analytics/activityStatistics/activityStatisticsIdValue",
			Expected: &UserAnalyticActivityStatisticId{
				UserId:               "userIdValue",
				ActivityStatisticsId: "activityStatisticsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/analytics/activityStatistics/activityStatisticsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAnalyticActivityStatisticID(v.Input)
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

		if actual.ActivityStatisticsId != v.Expected.ActivityStatisticsId {
			t.Fatalf("Expected %q but got %q for ActivityStatisticsId", v.Expected.ActivityStatisticsId, actual.ActivityStatisticsId)
		}

	}
}

func TestParseUserAnalyticActivityStatisticIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAnalyticActivityStatisticId
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
			Input: "/users/userIdValue/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/analytics/activityStatistics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aNaLyTiCs/aCtIvItYsTaTiStIcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/analytics/activityStatistics/activityStatisticsIdValue",
			Expected: &UserAnalyticActivityStatisticId{
				UserId:               "userIdValue",
				ActivityStatisticsId: "activityStatisticsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/analytics/activityStatistics/activityStatisticsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aNaLyTiCs/aCtIvItYsTaTiStIcS/aCtIvItYsTaTiStIcSiDvAlUe",
			Expected: &UserAnalyticActivityStatisticId{
				UserId:               "uSeRiDvAlUe",
				ActivityStatisticsId: "aCtIvItYsTaTiStIcSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aNaLyTiCs/aCtIvItYsTaTiStIcS/aCtIvItYsTaTiStIcSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAnalyticActivityStatisticIDInsensitively(v.Input)
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

		if actual.ActivityStatisticsId != v.Expected.ActivityStatisticsId {
			t.Fatalf("Expected %q but got %q for ActivityStatisticsId", v.Expected.ActivityStatisticsId, actual.ActivityStatisticsId)
		}

	}
}

func TestSegmentsForUserAnalyticActivityStatisticId(t *testing.T) {
	segments := UserAnalyticActivityStatisticId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAnalyticActivityStatisticId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
