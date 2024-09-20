package publishedartifact

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &VersionArtifactScopedId{}

func TestNewVersionArtifactScopedID(t *testing.T) {
	id := NewVersionArtifactScopedID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintName", "versionId", "artifactName")

	if id.ResourceScope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceScope'", id.ResourceScope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.BlueprintName != "blueprintName" {
		t.Fatalf("Expected %q but got %q for Segment 'BlueprintName'", id.BlueprintName, "blueprintName")
	}

	if id.VersionId != "versionId" {
		t.Fatalf("Expected %q but got %q for Segment 'VersionId'", id.VersionId, "versionId")
	}

	if id.ArtifactName != "artifactName" {
		t.Fatalf("Expected %q but got %q for Segment 'ArtifactName'", id.ArtifactName, "artifactName")
	}
}

func TestFormatVersionArtifactScopedID(t *testing.T) {
	actual := NewVersionArtifactScopedID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintName", "versionId", "artifactName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts/artifactName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVersionArtifactScopedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionArtifactScopedId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts/artifactName",
			Expected: &VersionArtifactScopedId{
				ResourceScope: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				BlueprintName: "blueprintName",
				VersionId:     "versionId",
				ArtifactName:  "artifactName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts/artifactName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionArtifactScopedID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceScope != v.Expected.ResourceScope {
			t.Fatalf("Expected %q but got %q for ResourceScope", v.Expected.ResourceScope, actual.ResourceScope)
		}

		if actual.BlueprintName != v.Expected.BlueprintName {
			t.Fatalf("Expected %q but got %q for BlueprintName", v.Expected.BlueprintName, actual.BlueprintName)
		}

		if actual.VersionId != v.Expected.VersionId {
			t.Fatalf("Expected %q but got %q for VersionId", v.Expected.VersionId, actual.VersionId)
		}

		if actual.ArtifactName != v.Expected.ArtifactName {
			t.Fatalf("Expected %q but got %q for ArtifactName", v.Expected.ArtifactName, actual.ArtifactName)
		}

	}
}

func TestParseVersionArtifactScopedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionArtifactScopedId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/vErSiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/vErSiOnS/vErSiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/vErSiOnS/vErSiOnId/aRtIfAcTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts/artifactName",
			Expected: &VersionArtifactScopedId{
				ResourceScope: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				BlueprintName: "blueprintName",
				VersionId:     "versionId",
				ArtifactName:  "artifactName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/versions/versionId/artifacts/artifactName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/vErSiOnS/vErSiOnId/aRtIfAcTs/aRtIfAcTnAmE",
			Expected: &VersionArtifactScopedId{
				ResourceScope: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				BlueprintName: "bLuEpRiNtNaMe",
				VersionId:     "vErSiOnId",
				ArtifactName:  "aRtIfAcTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/vErSiOnS/vErSiOnId/aRtIfAcTs/aRtIfAcTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionArtifactScopedIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceScope != v.Expected.ResourceScope {
			t.Fatalf("Expected %q but got %q for ResourceScope", v.Expected.ResourceScope, actual.ResourceScope)
		}

		if actual.BlueprintName != v.Expected.BlueprintName {
			t.Fatalf("Expected %q but got %q for BlueprintName", v.Expected.BlueprintName, actual.BlueprintName)
		}

		if actual.VersionId != v.Expected.VersionId {
			t.Fatalf("Expected %q but got %q for VersionId", v.Expected.VersionId, actual.VersionId)
		}

		if actual.ArtifactName != v.Expected.ArtifactName {
			t.Fatalf("Expected %q but got %q for ArtifactName", v.Expected.ArtifactName, actual.ArtifactName)
		}

	}
}

func TestSegmentsForVersionArtifactScopedId(t *testing.T) {
	segments := VersionArtifactScopedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VersionArtifactScopedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
