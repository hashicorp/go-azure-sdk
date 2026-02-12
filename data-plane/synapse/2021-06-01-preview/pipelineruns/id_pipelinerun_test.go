package pipelineruns

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PipelineRunId{}

func TestNewPipelineRunID(t *testing.T) {
	id := NewPipelineRunID("https://endpoint-url.example.com", "runId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.RunId != "runId" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runId")
	}
}

func TestFormatPipelineRunID(t *testing.T) {
	actual := NewPipelineRunID("https://endpoint-url.example.com", "runId").ID()
	expected := "https://endpoint-url.example.com/pipelineRuns/runId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePipelineRunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PipelineRunId
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
			Input: "https://endpoint-url.example.com/pipelineRuns",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/pipelineRuns/runId",
			Expected: &PipelineRunId{
				BaseURI: "https://endpoint-url.example.com",
				RunId:   "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pipelineRuns/runId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePipelineRunID(v.Input)
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

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

	}
}

func TestParsePipelineRunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PipelineRunId
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
			Input: "https://endpoint-url.example.com/pipelineRuns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInErUnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/pipelineRuns/runId",
			Expected: &PipelineRunId{
				BaseURI: "https://endpoint-url.example.com",
				RunId:   "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pipelineRuns/runId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInErUnS/rUnId",
			Expected: &PipelineRunId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				RunId:   "rUnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInErUnS/rUnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePipelineRunIDInsensitively(v.Input)
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

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

	}
}

func TestSegmentsForPipelineRunId(t *testing.T) {
	segments := PipelineRunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PipelineRunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
