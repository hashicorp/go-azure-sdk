package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{}

func TestNewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	id := NewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID("servicePrincipalId", "synchronizationTemplateId", "directoryDefinitionId")

	if id.ServicePrincipalId != "servicePrincipalId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalId")
	}

	if id.SynchronizationTemplateId != "synchronizationTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationTemplateId'", id.SynchronizationTemplateId, "synchronizationTemplateId")
	}

	if id.DirectoryDefinitionId != "directoryDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryDefinitionId'", id.DirectoryDefinitionId, "directoryDefinitionId")
	}
}

func TestFormatServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	actual := NewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID("servicePrincipalId", "synchronizationTemplateId", "directoryDefinitionId").ID()
	expected := "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId
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
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId",
			Expected: &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
				ServicePrincipalId:        "servicePrincipalId",
				SynchronizationTemplateId: "synchronizationTemplateId",
				DirectoryDefinitionId:     "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(v.Input)
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

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId
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
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA/dIrEcToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId",
			Expected: &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
				ServicePrincipalId:        "servicePrincipalId",
				SynchronizationTemplateId: "synchronizationTemplateId",
				DirectoryDefinitionId:     "directoryDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalId/synchronization/templates/synchronizationTemplateId/schema/directories/directoryDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId",
			Expected: &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
				ServicePrincipalId:        "sErViCePrInCiPaLiD",
				SynchronizationTemplateId: "sYnChRoNiZaTiOnTeMpLaTeId",
				DirectoryDefinitionId:     "dIrEcToRyDeFiNiTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiD/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeId/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryIDInsensitively(v.Input)
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

		if actual.SynchronizationTemplateId != v.Expected.SynchronizationTemplateId {
			t.Fatalf("Expected %q but got %q for SynchronizationTemplateId", v.Expected.SynchronizationTemplateId, actual.SynchronizationTemplateId)
		}

		if actual.DirectoryDefinitionId != v.Expected.DirectoryDefinitionId {
			t.Fatalf("Expected %q but got %q for DirectoryDefinitionId", v.Expected.DirectoryDefinitionId, actual.DirectoryDefinitionId)
		}

	}
}

func TestSegmentsForServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId(t *testing.T) {
	segments := ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
