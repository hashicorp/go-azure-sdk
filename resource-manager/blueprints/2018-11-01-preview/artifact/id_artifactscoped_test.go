package artifact

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ArtifactScopedId{}

func TestNewArtifactScopedID(t *testing.T) {
	id := NewArtifactScopedID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintName", "artifactName")

	if id.ResourceScope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceScope'", id.ResourceScope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.BlueprintName != "blueprintName" {
		t.Fatalf("Expected %q but got %q for Segment 'BlueprintName'", id.BlueprintName, "blueprintName")
	}

	if id.ArtifactName != "artifactName" {
		t.Fatalf("Expected %q but got %q for Segment 'ArtifactName'", id.ArtifactName, "artifactName")
	}
}

func TestFormatArtifactScopedID(t *testing.T) {
	actual := NewArtifactScopedID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintName", "artifactName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts/artifactName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseArtifactScopedID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ArtifactScopedId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts/artifactName",
			Expected: &ArtifactScopedId{
				ResourceScope: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				BlueprintName: "blueprintName",
				ArtifactName:  "artifactName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts/artifactName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseArtifactScopedID(v.Input)
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

		if actual.ArtifactName != v.Expected.ArtifactName {
			t.Fatalf("Expected %q but got %q for ArtifactName", v.Expected.ArtifactName, actual.ArtifactName)
		}

	}
}

func TestParseArtifactScopedIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ArtifactScopedId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/aRtIfAcTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts/artifactName",
			Expected: &ArtifactScopedId{
				ResourceScope: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				BlueprintName: "blueprintName",
				ArtifactName:  "artifactName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprints/blueprintName/artifacts/artifactName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/aRtIfAcTs/aRtIfAcTnAmE",
			Expected: &ArtifactScopedId{
				ResourceScope: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				BlueprintName: "bLuEpRiNtNaMe",
				ArtifactName:  "aRtIfAcTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtS/bLuEpRiNtNaMe/aRtIfAcTs/aRtIfAcTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseArtifactScopedIDInsensitively(v.Input)
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

		if actual.ArtifactName != v.Expected.ArtifactName {
			t.Fatalf("Expected %q but got %q for ArtifactName", v.Expected.ArtifactName, actual.ArtifactName)
		}

	}
}

func TestSegmentsForArtifactScopedId(t *testing.T) {
	segments := ArtifactScopedId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ArtifactScopedId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
