package synchronizationtemplateschemadirectory

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{}

func TestNewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	id := NewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID("servicePrincipalIdValue", "synchronizationTemplateIdValue", "directoryDefinitionIdValue")

	if id.ServicePrincipalId != "servicePrincipalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalId'", id.ServicePrincipalId, "servicePrincipalIdValue")
	}

	if id.SynchronizationTemplateId != "synchronizationTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SynchronizationTemplateId'", id.SynchronizationTemplateId, "synchronizationTemplateIdValue")
	}

	if id.DirectoryDefinitionId != "directoryDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryDefinitionId'", id.DirectoryDefinitionId, "directoryDefinitionIdValue")
	}
}

func TestFormatServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID(t *testing.T) {
	actual := NewServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryID("servicePrincipalIdValue", "synchronizationTemplateIdValue", "directoryDefinitionIdValue").ID()
	expected := "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories/directoryDefinitionIdValue"
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories/directoryDefinitionIdValue",
			Expected: &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
				ServicePrincipalId:        "servicePrincipalIdValue",
				SynchronizationTemplateId: "synchronizationTemplateIdValue",
				DirectoryDefinitionId:     "directoryDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories/directoryDefinitionIdValue/extra",
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
			Input: "/servicePrincipals/servicePrincipalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE/sChEmA",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE/sChEmA/dIrEcToRiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories/directoryDefinitionIdValue",
			Expected: &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
				ServicePrincipalId:        "servicePrincipalIdValue",
				SynchronizationTemplateId: "synchronizationTemplateIdValue",
				DirectoryDefinitionId:     "directoryDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/servicePrincipals/servicePrincipalIdValue/synchronization/templates/synchronizationTemplateIdValue/schema/directories/directoryDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnIdVaLuE",
			Expected: &ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId{
				ServicePrincipalId:        "sErViCePrInCiPaLiDvAlUe",
				SynchronizationTemplateId: "sYnChRoNiZaTiOnTeMpLaTeIdVaLuE",
				DirectoryDefinitionId:     "dIrEcToRyDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sErViCePrInCiPaLs/sErViCePrInCiPaLiDvAlUe/sYnChRoNiZaTiOn/tEmPlAtEs/sYnChRoNiZaTiOnTeMpLaTeIdVaLuE/sChEmA/dIrEcToRiEs/dIrEcToRyDeFiNiTiOnIdVaLuE/extra",
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
