package jobs

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &JobscheduleId{}

func TestNewJobscheduleID(t *testing.T) {
	id := NewJobscheduleID("https://endpoint-url.example.com", "jobScheduleId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.JobScheduleId != "jobScheduleId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobScheduleId'", id.JobScheduleId, "jobScheduleId")
	}
}

func TestFormatJobscheduleID(t *testing.T) {
	actual := NewJobscheduleID("https://endpoint-url.example.com", "jobScheduleId").ID()
	expected := "https://endpoint-url.example.com/jobschedules/jobScheduleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseJobscheduleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *JobscheduleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobschedules",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/jobschedules/jobScheduleId",
			Expected: &JobscheduleId{
				BaseURI:       "https://endpoint-url.example.com",
				JobScheduleId: "jobScheduleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/jobschedules/jobScheduleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseJobscheduleID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.JobScheduleId != v.Expected.JobScheduleId {
			t.Fatalf("Expected %q but got %q for JobScheduleId", v.Expected.JobScheduleId, actual.JobScheduleId)
		}

	}
}

func TestParseJobscheduleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *JobscheduleId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobschedules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObScHeDuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/jobschedules/jobScheduleId",
			Expected: &JobscheduleId{
				BaseURI:       "https://endpoint-url.example.com",
				JobScheduleId: "jobScheduleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/jobschedules/jobScheduleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObScHeDuLeS/jObScHeDuLeId",
			Expected: &JobscheduleId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				JobScheduleId: "jObScHeDuLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObScHeDuLeS/jObScHeDuLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseJobscheduleIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.JobScheduleId != v.Expected.JobScheduleId {
			t.Fatalf("Expected %q but got %q for JobScheduleId", v.Expected.JobScheduleId, actual.JobScheduleId)
		}

	}
}

func TestSegmentsForJobscheduleId(t *testing.T) {
	segments := JobscheduleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("JobscheduleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
