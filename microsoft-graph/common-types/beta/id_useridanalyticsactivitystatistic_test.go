package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAnalyticsActivityStatisticId{}

func TestNewUserIdAnalyticsActivityStatisticID(t *testing.T) {
	id := NewUserIdAnalyticsActivityStatisticID("userId", "activityStatisticsId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ActivityStatisticsId != "activityStatisticsId" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityStatisticsId'", id.ActivityStatisticsId, "activityStatisticsId")
	}
}

func TestFormatUserIdAnalyticsActivityStatisticID(t *testing.T) {
	actual := NewUserIdAnalyticsActivityStatisticID("userId", "activityStatisticsId").ID()
	expected := "/users/userId/analytics/activityStatistics/activityStatisticsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAnalyticsActivityStatisticID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAnalyticsActivityStatisticId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/analytics/activityStatistics",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/analytics/activityStatistics/activityStatisticsId",
			Expected: &UserIdAnalyticsActivityStatisticId{
				UserId:               "userId",
				ActivityStatisticsId: "activityStatisticsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/analytics/activityStatistics/activityStatisticsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAnalyticsActivityStatisticID(v.Input)
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

func TestParseUserIdAnalyticsActivityStatisticIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAnalyticsActivityStatisticId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/analytics/activityStatistics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aNaLyTiCs/aCtIvItYsTaTiStIcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/analytics/activityStatistics/activityStatisticsId",
			Expected: &UserIdAnalyticsActivityStatisticId{
				UserId:               "userId",
				ActivityStatisticsId: "activityStatisticsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/analytics/activityStatistics/activityStatisticsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aNaLyTiCs/aCtIvItYsTaTiStIcS/aCtIvItYsTaTiStIcSiD",
			Expected: &UserIdAnalyticsActivityStatisticId{
				UserId:               "uSeRiD",
				ActivityStatisticsId: "aCtIvItYsTaTiStIcSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aNaLyTiCs/aCtIvItYsTaTiStIcS/aCtIvItYsTaTiStIcSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAnalyticsActivityStatisticIDInsensitively(v.Input)
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

func TestSegmentsForUserIdAnalyticsActivityStatisticId(t *testing.T) {
	segments := UserIdAnalyticsActivityStatisticId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAnalyticsActivityStatisticId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
