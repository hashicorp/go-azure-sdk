package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{}

func TestNewServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	id := NewServicePrincipalIdSynchronizationJobIdSchemaDirectoryID("servicePrincipalId", "synchronizationJobId", "directoryDefinitionId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.SynchronizationJobId != "synchronizationJobId" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationJobId'", id.SynchronizationJobId, "synchronizationJobId")
	}

	if id.DirectoryDefinitionId != "directoryDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryDefinitionId'", id.DirectoryDefinitionId, "directoryDefinitionId")
	}
}

func TestFormatServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	actual := NewServicePrincipalIdSynchronizationJobIdSchemaDirectoryID("servicePrincipalId", "synchronizationJobId", "directoryDefinitionId").ID()
	expected := "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId",
			Expected: &ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{
				ServicePrincipalId:    "servicePrincipalId",
				SynchronizationJobId:  "synchronizationJobId",
				DirectoryDefinitionId: "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.SynchronizationJobId != v.Expected.SynchronizationJobId {
			t.Fatalf("Expected %q but got %q for SynchronizationJobId", v.Expected.SynchronizationJobId, actual.SynchronizationJobId)
		}

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA/dIrEcToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId",
			Expected: &ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{
				ServicePrincipalId:    "servicePrincipalId",
				SynchronizationJobId:  "synchronizationJobId",
				DirectoryDefinitionId: "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/jobs/synchronizationJobId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId",
			Expected: &ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{
				ServicePrincipalId:    "sErViCePrInCiPaLiD",
				SynchronizationJobId:  "sYnChRoNiZaTiOnJoBiD",
				DirectoryDefinitionId: "dIrEcToRyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/jObS/sYnChRoNiZaTiOnJoBiD/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationJobIdSchemaDirectoryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalId != v.Expected.ServicePrincipalId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalId", v.Expected.ServicePrincipalId, actual.ServicePrincipalId)
		}

		if actual.SynchronizationJobId != v.Expected.SynchronizationJobId {
			t.Fatalf("Expected %q but got %q for SynchronizationJobId", v.Expected.SynchronizationJobId, actual.SynchronizationJobId)
		}

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestSegmentsForServicePrincipalIdSynchronizationJobIdSchemaDirectoryId(t *testing.T) {
	segments := ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdSynchronizationJobIdSchemaDirectoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
