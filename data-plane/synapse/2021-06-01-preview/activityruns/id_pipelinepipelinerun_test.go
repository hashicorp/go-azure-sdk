package activityruns

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PipelinePipelineRunId{}

func TestNewPipelinePipelineRunID(t *testing.T) {
	id := NewPipelinePipelineRunID("https://endpoint-url.example.com", "pipelineName", "runId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.PipelineName != "pipelineName" {
		t.Fatalf("Expected %q but got %q for Segment 'PipelineName'", id.PipelineName, "pipelineName")
	}

	if id.RunId != "runId" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runId")
	}
}

func TestFormatPipelinePipelineRunID(t *testing.T) {
	actual := NewPipelinePipelineRunID("https://endpoint-url.example.com", "pipelineName", "runId").ID()
	expected := "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns/runId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePipelinePipelineRunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PipelinePipelineRunId
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
			Input: "https://endpoint-url.example.com/pipelines",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns/runId",
			Expected: &PipelinePipelineRunId{
				BaseURI:      "https://endpoint-url.example.com",
				PipelineName: "pipelineName",
				RunId:        "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns/runId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePipelinePipelineRunID(v.Input)
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

		if actual.PipelineName != v.Expected.PipelineName {
			t.Fatalf("Expected %q but got %q for PipelineName", v.Expected.PipelineName, actual.PipelineName)
		}

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

	}
}

func TestParsePipelinePipelineRunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PipelinePipelineRunId
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
			Input: "https://endpoint-url.example.com/pipelines",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs/pIpElInEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs/pIpElInEnAmE/pIpElInErUnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns/runId",
			Expected: &PipelinePipelineRunId{
				BaseURI:      "https://endpoint-url.example.com",
				PipelineName: "pipelineName",
				RunId:        "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/pipelineRuns/runId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs/pIpElInEnAmE/pIpElInErUnS/rUnId",
			Expected: &PipelinePipelineRunId{
				BaseURI:      "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				PipelineName: "pIpElInEnAmE",
				RunId:        "rUnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs/pIpElInEnAmE/pIpElInErUnS/rUnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePipelinePipelineRunIDInsensitively(v.Input)
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

		if actual.PipelineName != v.Expected.PipelineName {
			t.Fatalf("Expected %q but got %q for PipelineName", v.Expected.PipelineName, actual.PipelineName)
		}

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

	}
}

func TestSegmentsForPipelinePipelineRunId(t *testing.T) {
	segments := PipelinePipelineRunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PipelinePipelineRunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
