package applicationsynchronizationjob

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationSynchronizationJobId{}

func TestNewApplicationSynchronizationJobID(t *testing.T) {
	id := NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.SynchronizationJobId != "synchronizationJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationJobId'", id.SynchronizationJobId, "synchronizationJobIdValue")
	}
}

func TestFormatApplicationSynchronizationJobID(t *testing.T) {
	actual := NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue").ID()
	expected := "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationSynchronizationJobID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationSynchronizationJobId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue",
			Expected: &ApplicationSynchronizationJobId{
				ApplicationId:        "applicationIdValue",
				SynchronizationJobId: "synchronizationJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationSynchronizationJobID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.SynchronizationJobId != v.Expected.SynchronizationJobId {
			t.Fatalf("Expected %q but got %q for SynchronizationJobId", v.Expected.SynchronizationJobId, actual.SynchronizationJobId)
		}

	}
}

func TestParseApplicationSynchronizationJobIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationSynchronizationJobId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue",
			Expected: &ApplicationSynchronizationJobId{
				ApplicationId:        "applicationIdValue",
				SynchronizationJobId: "synchronizationJobIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe",
			Expected: &ApplicationSynchronizationJobId{
				ApplicationId:        "aPpLiCaTiOnIdVaLuE",
				SynchronizationJobId: "sYnChRoNiZaTiOnJoBiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationSynchronizationJobIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationId != v.Expected.ApplicationId {
			t.Fatalf("Expected %q but got %q for ApplicationId", v.Expected.ApplicationId, actual.ApplicationId)
		}

		if actual.SynchronizationJobId != v.Expected.SynchronizationJobId {
			t.Fatalf("Expected %q but got %q for SynchronizationJobId", v.Expected.SynchronizationJobId, actual.SynchronizationJobId)
		}

	}
}

func TestSegmentsForApplicationSynchronizationJobId(t *testing.T) {
	segments := ApplicationSynchronizationJobId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationSynchronizationJobId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
