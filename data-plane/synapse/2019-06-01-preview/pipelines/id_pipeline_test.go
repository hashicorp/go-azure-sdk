package pipelines

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PipelineId{}

func TestNewPipelineID(t *testing.T) {
	id := NewPipelineID("https://endpoint-url.example.com", "pipelineName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.PipelineName != "pipelineName" {
		t.Fatalf("Expected %q but got %q for Segment 'PipelineName'", id.PipelineName, "pipelineName")
	}
}

func TestFormatPipelineID(t *testing.T) {
	actual := NewPipelineID("https://endpoint-url.example.com", "pipelineName").ID()
	expected := "https://endpoint-url.example.com/pipelines/pipelineName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePipelineID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PipelineId
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
			// Valid URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName",
			Expected: &PipelineId{
				BaseURI:      "https://endpoint-url.example.com",
				PipelineName: "pipelineName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePipelineID(v.Input)
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

	}
}

func TestParsePipelineIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PipelineId
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
			// Valid URI
			Input: "https://endpoint-url.example.com/pipelines/pipelineName",
			Expected: &PipelineId{
				BaseURI:      "https://endpoint-url.example.com",
				PipelineName: "pipelineName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/pipelines/pipelineName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs/pIpElInEnAmE",
			Expected: &PipelineId{
				BaseURI:      "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				PipelineName: "pIpElInEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/pIpElInEs/pIpElInEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePipelineIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForPipelineId(t *testing.T) {
	segments := PipelineId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PipelineId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
