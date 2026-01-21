package backuprestores

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BackupId{}

func TestNewBackupID(t *testing.T) {
	id := NewBackupID("jobId")

	if id.JobId != "jobId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobId'", id.JobId, "jobId")
	}
}

func TestFormatBackupID(t *testing.T) {
	actual := NewBackupID("jobId").ID()
	expected := "/backup/jobId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBackupID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BackupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/backup",
			Error: true,
		},
		{
			// Valid URI
			Input: "/backup/jobId",
			Expected: &BackupId{
				JobId: "jobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/backup/jobId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBackupID(v.Input)
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

	}
}

func TestParseBackupIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BackupId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/backup",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/bAcKuP",
			Error: true,
		},
		{
			// Valid URI
			Input: "/backup/jobId",
			Expected: &BackupId{
				JobId: "jobId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/backup/jobId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/bAcKuP/jObId",
			Expected: &BackupId{
				JobId: "jObId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/bAcKuP/jObId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBackupIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForBackupId(t *testing.T) {
	segments := BackupId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BackupId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
