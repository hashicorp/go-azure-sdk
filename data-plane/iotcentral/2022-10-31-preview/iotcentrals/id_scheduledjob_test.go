package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScheduledJobId{}

func TestNewScheduledJobID(t *testing.T) {
	id := NewScheduledJobID("https://endpoint-url.example.com", "scheduledJobId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.ScheduledJobId != "scheduledJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'ScheduledJobId'", id.ScheduledJobId, "scheduledJobId")
	}
}

func TestFormatScheduledJobID(t *testing.T) {
	actual := NewScheduledJobID("https://endpoint-url.example.com", "scheduledJobId").ID()
	expected := "https://endpoint-url.example.com/scheduledJobs/scheduledJobId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScheduledJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScheduledJobId
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
			Input: "https://endpoint-url.example.com/scheduledJobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/scheduledJobs/scheduledJobId",
			Expected: &ScheduledJobId{
				BaseURI:        "https://endpoint-url.example.com",
				ScheduledJobId: "scheduledJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/scheduledJobs/scheduledJobId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScheduledJobID(v.Input)
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

		if actual.ScheduledJobId != v.Expected.ScheduledJobId {
			t.Fatalf("Expected %q but got %q for ScheduledJobId", v.Expected.ScheduledJobId, actual.ScheduledJobId)
		}

	}
}

func TestParseScheduledJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScheduledJobId
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
			Input: "https://endpoint-url.example.com/scheduledJobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sChEdUlEdJoBs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/scheduledJobs/scheduledJobId",
			Expected: &ScheduledJobId{
				BaseURI:        "https://endpoint-url.example.com",
				ScheduledJobId: "scheduledJobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/scheduledJobs/scheduledJobId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sChEdUlEdJoBs/sChEdUlEdJoBiD",
			Expected: &ScheduledJobId{
				BaseURI:        "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				ScheduledJobId: "sChEdUlEdJoBiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sChEdUlEdJoBs/sChEdUlEdJoBiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScheduledJobIDInsensitively(v.Input)
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

		if actual.ScheduledJobId != v.Expected.ScheduledJobId {
			t.Fatalf("Expected %q but got %q for ScheduledJobId", v.Expected.ScheduledJobId, actual.ScheduledJobId)
		}

	}
}

func TestSegmentsForScheduledJobId(t *testing.T) {
	segments := ScheduledJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScheduledJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
