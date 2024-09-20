package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{}

func TestNewApplicationIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	id := NewApplicationIdSynchronizationTemplateIdSchemaDirectoryID("applicationId", "synchronizationTemplateId", "directoryDefinitionId")

	if id.ApplicationId != "applicationId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationId'", id.ApplicationId, "applicationId")
	}

	if id.SynchronizationTemplateId != "synchronizationTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationTemplateId'", id.SynchronizationTemplateId, "synchronizationTemplateId")
	}

	if id.DirectoryDefinitionId != "directoryDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryDefinitionId'", id.DirectoryDefinitionId, "directoryDefinitionId")
	}
}

func TestFormatApplicationIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	actual := NewApplicationIdSynchronizationTemplateIdSchemaDirectoryID("applicationId", "synchronizationTemplateId", "directoryDefinitionId").ID()
	expected := "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApplicationIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdSynchronizationTemplateIdSchemaDirectoryId
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
			Input: "/applications/applicationId/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId",
			Expected: &ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{
				ApplicationId:             "applicationId",
				SynchronizationTemplateId: "synchronizationTemplateId",
				DirectoryDefinitionId:     "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryID(v.Input)
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

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestParseApplicationIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApplicationIdSynchronizationTemplateIdSchemaDirectoryId
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
			Input: "/applications/applicationId/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA/dIrEcToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId",
			Expected: &ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{
				ApplicationId:             "applicationId",
				SynchronizationTemplateId: "synchronizationTemplateId",
				DirectoryDefinitionId:     "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/applications/applicationId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId",
			Expected: &ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{
				ApplicationId:             "aPpLiCaTiOnId",
				SynchronizationTemplateId: "sYnChRoNiZaTiOnTeMpLaTeId",
				DirectoryDefinitionId:     "dIrEcToRyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/aPpLiCaTiOnS/aPpLiCaTiOnId/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApplicationIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively(v.Input)
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

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestSegmentsForApplicationIdSynchronizationTemplateIdSchemaDirectoryId(t *testing.T) {
	segments := ApplicationIdSynchronizationTemplateIdSchemaDirectoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApplicationIdSynchronizationTemplateIdSchemaDirectoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
