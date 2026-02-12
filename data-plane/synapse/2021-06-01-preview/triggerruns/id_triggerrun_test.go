package triggerruns

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TriggerRunId{}

func TestNewTriggerRunID(t *testing.T) {
	id := NewTriggerRunID("https://endpoint-url.example.com", "triggerName", "runId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.TriggerName != "triggerName" {
		t.Fatalf("Expected %q but got %q for Segment 'TriggerName'", id.TriggerName, "triggerName")
	}

	if id.RunId != "runId" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runId")
	}
}

func TestFormatTriggerRunID(t *testing.T) {
	actual := NewTriggerRunID("https://endpoint-url.example.com", "triggerName", "runId").ID()
	expected := "https://endpoint-url.example.com/triggers/triggerName/triggerRuns/runId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTriggerRunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TriggerRunId
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
			Input: "https://endpoint-url.example.com/triggers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/triggers/triggerName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/triggers/triggerName/triggerRuns",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/triggers/triggerName/triggerRuns/runId",
			Expected: &TriggerRunId{
				BaseURI:     "https://endpoint-url.example.com",
				TriggerName: "triggerName",
				RunId:       "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/triggers/triggerName/triggerRuns/runId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTriggerRunID(v.Input)
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

		if actual.TriggerName != v.Expected.TriggerName {
			t.Fatalf("Expected %q but got %q for TriggerName", v.Expected.TriggerName, actual.TriggerName)
		}

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

	}
}

func TestParseTriggerRunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TriggerRunId
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
			Input: "https://endpoint-url.example.com/triggers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/tRiGgErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/triggers/triggerName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/tRiGgErS/tRiGgErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/triggers/triggerName/triggerRuns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/tRiGgErS/tRiGgErNaMe/tRiGgErRuNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/triggers/triggerName/triggerRuns/runId",
			Expected: &TriggerRunId{
				BaseURI:     "https://endpoint-url.example.com",
				TriggerName: "triggerName",
				RunId:       "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/triggers/triggerName/triggerRuns/runId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/tRiGgErS/tRiGgErNaMe/tRiGgErRuNs/rUnId",
			Expected: &TriggerRunId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				TriggerName: "tRiGgErNaMe",
				RunId:       "rUnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/tRiGgErS/tRiGgErNaMe/tRiGgErRuNs/rUnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTriggerRunIDInsensitively(v.Input)
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

		if actual.TriggerName != v.Expected.TriggerName {
			t.Fatalf("Expected %q but got %q for TriggerName", v.Expected.TriggerName, actual.TriggerName)
		}

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

	}
}

func TestSegmentsForTriggerRunId(t *testing.T) {
	segments := TriggerRunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TriggerRunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
