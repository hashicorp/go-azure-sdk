package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationJobIdSchemaDirectoryId{}

func TestNewApplicationIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	id := NewApplicationIdSynchronizationJobIdSchemaDirectoryID("applicationIdValue", "synchronizationJobIdValue", "directoryDefinitionIdValue")

	if id.ApplicationId != "applicationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationIdValue")
	}

	if id.SynchronizationJobId != "synchronizationJobIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationJobId'", id.SynchronizationJobId, "synchronizationJobIdValue")
	}

	if id.DirectoryDefinitionId != "directoryDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryDefinitionId'", id.DirectoryDefinitionId, "directoryDefinitionIdValue")
	}
}

func TestFormatApplicationIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	actual := NewApplicationIdSynchronizationJobIdSchemaDirectoryID("applicationIdValue", "synchronizationJobIdValue", "directoryDefinitionIdValue").ID()
	expected := "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories/directoryDefinitionIdValue"
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
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories/directoryDefinitionIdValue",
			Expected: &ApplicationIdSynchronizationJobIdSchemaDirectoryId{
				ApplicationId:         "applicationIdValue",
				SynchronizationJobId:  "synchronizationJobIdValue",
				DirectoryDefinitionId: "directoryDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories/directoryDefinitionIdValue/extra",
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
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe/sChEmA",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe/sChEmA/dIrEcToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories/directoryDefinitionIdValue",
			Expected: &ApplicationIdSynchronizationJobIdSchemaDirectoryId{
				ApplicationId:         "applicationIdValue",
				SynchronizationJobId:  "synchronizationJobIdValue",
				DirectoryDefinitionId: "directoryDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationIdValue/synchronization/jobs/synchronizationJobIdValue/schema/directories/directoryDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnIdVaLuE",
			Expected: &ApplicationIdSynchronizationJobIdSchemaDirectoryId{
				ApplicationId:         "aPpLiCaTiOnIdVaLuE",
				SynchronizationJobId:  "sYnChRoNiZaTiOnJoBiDvAlUe",
				DirectoryDefinitionId: "dIrEcToRyDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnIdVaLuE/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiDvAlUe/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnIdVaLuE/extra",
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
