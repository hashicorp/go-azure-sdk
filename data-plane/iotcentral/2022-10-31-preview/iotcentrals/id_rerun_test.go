package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RerunId{}

func TestNewRerunID(t *testing.T) {
	id := NewRerunID("https://endpoint-url.example.com", "jobId", "rerunId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.JobId != "jobId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobId'", id.JobId, "jobId")
	}

	if id.RerunId != "rerunId" {
		t.Fatalf("Expected %q but got %q for Segment 'RerunId'", id.RerunId, "rerunId")
	}
}

func TestFormatRerunID(t *testing.T) {
	actual := NewRerunID("https://endpoint-url.example.com", "jobId", "rerunId").ID()
	expected := "https://endpoint-url.example.com/jobs/jobId/rerun/rerunId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRerunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RerunId
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
			Input: "https://endpoint-url.example.com/jobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId/rerun",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/jobs/jobId/rerun/rerunId",
			Expected: &RerunId{
				BaseURI: "https://endpoint-url.example.com",
				JobId:   "jobId",
				RerunId: "rerunId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/jobs/jobId/rerun/rerunId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRerunID(v.Input)
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

		if actual.JobId != v.Expected.JobId {
			t.Fatalf("Expected %q but got %q for JobId", v.Expected.JobId, actual.JobId)
		}

		if actual.RerunId != v.Expected.RerunId {
			t.Fatalf("Expected %q but got %q for RerunId", v.Expected.RerunId, actual.RerunId)
		}

	}
}

func TestParseRerunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RerunId
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
			Input: "https://endpoint-url.example.com/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId/rerun",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/rErUn",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/jobs/jobId/rerun/rerunId",
			Expected: &RerunId{
				BaseURI: "https://endpoint-url.example.com",
				JobId:   "jobId",
				RerunId: "rerunId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/jobs/jobId/rerun/rerunId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/rErUn/rErUnId",
			Expected: &RerunId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				JobId:   "jObId",
				RerunId: "rErUnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/rErUn/rErUnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRerunIDInsensitively(v.Input)
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

		if actual.JobId != v.Expected.JobId {
			t.Fatalf("Expected %q but got %q for JobId", v.Expected.JobId, actual.JobId)
		}

		if actual.RerunId != v.Expected.RerunId {
			t.Fatalf("Expected %q but got %q for RerunId", v.Expected.RerunId, actual.RerunId)
		}

	}
}

func TestSegmentsForRerunId(t *testing.T) {
	segments := RerunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RerunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
