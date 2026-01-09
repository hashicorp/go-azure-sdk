package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationJobIdSchemaDirectoryId{}

func TestNewApplicationIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	id := NewApplicationIdSynchronizationJobIdSchemaDirectoryID("applicationId", "synchronizationJobId", "directoryDefinitionId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.SynchronizationJobId != "synchronizationJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationJobId'", id.SynchronizationJobId, "synchronizationJobId")
	}

	if id.DirectoryDefinitionId != "directoryDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryDefinitionId'", id.DirectoryDefinitionId, "directoryDefinitionId")
	}
}

func TestFormatApplicationIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	actual := NewApplicationIdSynchronizationJobIdSchemaDirectoryID("applicationId", "synchronizationJobId", "directoryDefinitionId").ID()
	expected := "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdSynchronizationJobIdSchemaDirectoryId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId",
			Expected: &ApplicationIdSynchronizationJobIdSchemaDirectoryId{
				ApplicationId:         "applicationId",
				SynchronizationJobId:  "synchronizationJobId",
				DirectoryDefinitionId: "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdSynchronizationJobIdSchemaDirectoryID(v.Input)
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

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestParseApplicationIdSynchronizationJobIdSchemaDirectoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdSynchronizationJobIdSchemaDirectoryId
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
			Input: "/applications/applicationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/jObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA/dIrEcToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId",
			Expected: &ApplicationIdSynchronizationJobIdSchemaDirectoryId{
				ApplicationId:         "applicationId",
				SynchronizationJobId:  "synchronizationJobId",
				DirectoryDefinitionId: "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId",
			Expected: &ApplicationIdSynchronizationJobIdSchemaDirectoryId{
				ApplicationId:         "aPpLiCaTiOnId",
				SynchronizationJobId:  "sYnChRoNiZaTiOnJoBiD",
				DirectoryDefinitionId: "dIrEcToRyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdSynchronizationJobIdSchemaDirectoryIDInsensitively(v.Input)
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

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestSegmentsForApplicationIdSynchronizationJobIdSchemaDirectoryId(t *testing.T) {
	segments := ApplicationIdSynchronizationJobIdSchemaDirectoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdSynchronizationJobIdSchemaDirectoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
