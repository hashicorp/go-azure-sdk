package files

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FileId{}

func TestNewFileID(t *testing.T) {
	id := NewFileID("https://endpoint_url", "jobId", "taskId")

	if id.BaseURI != "https://endpoint_url" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint_url")
	}

	if id.JobId != "jobId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobId'", id.JobId, "jobId")
	}

	if id.TaskId != "taskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskId")
	}
}

func TestFormatFileID(t *testing.T) {
	actual := NewFileID("https://endpoint_url", "jobId", "taskId").ID()
	expected := "/https://endpoint_url/jobs/jobId/tasks/taskId/files"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseFileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FileId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs/jobId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs/jobId/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs/jobId/tasks/taskId",
			Error: true,
		},
		{
			// Valid URI
			Input: "/https://endpoint_url/jobs/jobId/tasks/taskId/files",
			Expected: &FileId{
				BaseURI: "https://endpoint_url",
				JobId:   "jobId",
				TaskId:  "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/jobs/jobId/tasks/taskId/files/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFileID(v.Input)
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

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

	}
}

func TestParseFileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FileId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/jObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs/jobId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/jObS/jObId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs/jobId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/jObS/jObId/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/jobs/jobId/tasks/taskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/jObS/jObId/tAsKs/tAsKiD",
			Error: true,
		},
		{
			// Valid URI
			Input: "/https://endpoint_url/jobs/jobId/tasks/taskId/files",
			Expected: &FileId{
				BaseURI: "https://endpoint_url",
				JobId:   "jobId",
				TaskId:  "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/jobs/jobId/tasks/taskId/files/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/jObS/jObId/tAsKs/tAsKiD/fIlEs",
			Expected: &FileId{
				BaseURI: "hTtPs://eNdPoInT_UrL",
				JobId:   "jObId",
				TaskId:  "tAsKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/jObS/jObId/tAsKs/tAsKiD/fIlEs/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseFileIDInsensitively(v.Input)
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

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

	}
}

func TestSegmentsForFileId(t *testing.T) {
	segments := FileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("FileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
