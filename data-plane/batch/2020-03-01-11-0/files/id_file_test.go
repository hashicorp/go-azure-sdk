package files

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FileId{}

func TestNewFileID(t *testing.T) {
	id := NewFileID("https://endpoint-url.example.com", "jobId", "taskId", "fileName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.JobId != "jobId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobId'", id.JobId, "jobId")
	}

	if id.TaskId != "taskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskId")
	}

	if id.FileName != "fileName" {
		t.Fatalf("Expected %q but got %q for Segment 'FileName'", id.FileName, "fileName")
	}
}

func TestFormatFileID(t *testing.T) {
	actual := NewFileID("https://endpoint-url.example.com", "jobId", "taskId", "fileName").ID()
	expected := "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files/fileName"
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
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files/fileName",
			Expected: &FileId{
				BaseURI:  "https://endpoint-url.example.com",
				JobId:    "jobId",
				TaskId:   "taskId",
				FileName: "fileName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files/fileName/extra",
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

		if actual.FileName != v.Expected.FileName {
			t.Fatalf("Expected %q but got %q for FileName", v.Expected.FileName, actual.FileName)
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
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/tAsKs/tAsKiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/tAsKs/tAsKiD/fIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files/fileName",
			Expected: &FileId{
				BaseURI:  "https://endpoint-url.example.com",
				JobId:    "jobId",
				TaskId:   "taskId",
				FileName: "fileName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/jobs/jobId/tasks/taskId/files/fileName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/tAsKs/tAsKiD/fIlEs/fIlEnAmE",
			Expected: &FileId{
				BaseURI:  "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				JobId:    "jObId",
				TaskId:   "tAsKiD",
				FileName: "fIlEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/jObS/jObId/tAsKs/tAsKiD/fIlEs/fIlEnAmE/extra",
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

		if actual.FileName != v.Expected.FileName {
			t.Fatalf("Expected %q but got %q for FileName", v.Expected.FileName, actual.FileName)
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
