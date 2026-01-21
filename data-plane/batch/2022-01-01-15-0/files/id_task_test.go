package files

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TaskId{}

func TestNewTaskID(t *testing.T) {
	id := NewTaskID("jobId", "taskId")

	if id.JobId != "jobId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobId'", id.JobId, "jobId")
	}

	if id.TaskId != "taskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskId")
	}
}

func TestFormatTaskID(t *testing.T) {
	actual := NewTaskID("jobId", "taskId").ID()
	expected := "/jobs/jobId/tasks/taskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TaskId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/jobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/jobs/jobId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/jobs/jobId/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/jobs/jobId/tasks/taskId",
			Expected: &TaskId{
				JobId:  "jobId",
				TaskId: "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/jobs/jobId/tasks/taskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTaskID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.JobId != v.Expected.JobId {
			t.Fatalf("Expected %q but got %q for JobId", v.Expected.JobId, actual.JobId)
		}

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

	}
}

func TestParseTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TaskId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/jObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/jobs/jobId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/jObS/jObId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/jobs/jobId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/jObS/jObId/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/jobs/jobId/tasks/taskId",
			Expected: &TaskId{
				JobId:  "jobId",
				TaskId: "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/jobs/jobId/tasks/taskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/jObS/jObId/tAsKs/tAsKiD",
			Expected: &TaskId{
				JobId:  "jObId",
				TaskId: "tAsKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/jObS/jObId/tAsKs/tAsKiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTaskIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.JobId != v.Expected.JobId {
			t.Fatalf("Expected %q but got %q for JobId", v.Expected.JobId, actual.JobId)
		}

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

	}
}

func TestSegmentsForTaskId(t *testing.T) {
	segments := TaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
