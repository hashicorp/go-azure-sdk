package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAnalyticsActivityStatisticId{}

func TestNewMeAnalyticsActivityStatisticID(t *testing.T) {
	id := NewMeAnalyticsActivityStatisticID("activityStatisticsIdValue")

	if id.ActivityStatisticsId != "activityStatisticsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityStatisticsId'", id.ActivityStatisticsId, "activityStatisticsIdValue")
	}
}

func TestFormatMeAnalyticsActivityStatisticID(t *testing.T) {
	actual := NewMeAnalyticsActivityStatisticID("activityStatisticsIdValue").ID()
	expected := "/me/analytics/activityStatistics/activityStatisticsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAnalyticsActivityStatisticID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAnalyticsActivityStatisticId
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
			Input: "/me/analytics",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/analytics/activityStatistics",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/analytics/activityStatistics/activityStatisticsIdValue",
			Expected: &MeAnalyticsActivityStatisticId{
				ActivityStatisticsId: "activityStatisticsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/analytics/activityStatistics/activityStatisticsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAnalyticsActivityStatisticID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActivityStatisticsId != v.Expected.ActivityStatisticsId {
			t.Fatalf("Expected %q but got %q for ActivityStatisticsId", v.Expected.ActivityStatisticsId, actual.ActivityStatisticsId)
		}

	}
}

func TestParseMeAnalyticsActivityStatisticIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAnalyticsActivityStatisticId
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
			Input: "/me/analytics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aNaLyTiCs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/analytics/activityStatistics",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aNaLyTiCs/aCtIvItYsTaTiStIcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/analytics/activityStatistics/activityStatisticsIdValue",
			Expected: &MeAnalyticsActivityStatisticId{
				ActivityStatisticsId: "activityStatisticsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/analytics/activityStatistics/activityStatisticsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aNaLyTiCs/aCtIvItYsTaTiStIcS/aCtIvItYsTaTiStIcSiDvAlUe",
			Expected: &MeAnalyticsActivityStatisticId{
				ActivityStatisticsId: "aCtIvItYsTaTiStIcSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aNaLyTiCs/aCtIvItYsTaTiStIcS/aCtIvItYsTaTiStIcSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAnalyticsActivityStatisticIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActivityStatisticsId != v.Expected.ActivityStatisticsId {
			t.Fatalf("Expected %q but got %q for ActivityStatisticsId", v.Expected.ActivityStatisticsId, actual.ActivityStatisticsId)
		}

	}
}

func TestSegmentsForMeAnalyticsActivityStatisticId(t *testing.T) {
	segments := MeAnalyticsActivityStatisticId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAnalyticsActivityStatisticId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
