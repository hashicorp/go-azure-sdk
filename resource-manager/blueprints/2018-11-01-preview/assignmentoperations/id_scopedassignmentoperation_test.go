package assignmentoperations

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedAssignmentOperationId{}

func TestNewScopedAssignmentOperationID(t *testing.T) {
	id := NewScopedAssignmentOperationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assignmentName", "assignmentOperationName")

	if id.ResourceScope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceScope'", id.ResourceScope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.BlueprintAssignmentName != "assignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'BlueprintAssignmentName'", id.BlueprintAssignmentName, "assignmentName")
	}

	if id.AssignmentOperationName != "assignmentOperationName" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentOperationName'", id.AssignmentOperationName, "assignmentOperationName")
	}
}

func TestFormatScopedAssignmentOperationID(t *testing.T) {
	actual := NewScopedAssignmentOperationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assignmentName", "assignmentOperationName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations/assignmentOperationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedAssignmentOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedAssignmentOperationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations/assignmentOperationName",
			Expected: &ScopedAssignmentOperationId{
				ResourceScope:           "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				BlueprintAssignmentName: "assignmentName",
				AssignmentOperationName: "assignmentOperationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations/assignmentOperationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedAssignmentOperationID(v.Input)
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

		if actual.BlueprintAssignmentName != v.Expected.BlueprintAssignmentName {
			t.Fatalf("Expected %q but got %q for BlueprintAssignmentName", v.Expected.BlueprintAssignmentName, actual.BlueprintAssignmentName)
		}

		if actual.AssignmentOperationName != v.Expected.AssignmentOperationName {
			t.Fatalf("Expected %q but got %q for AssignmentOperationName", v.Expected.AssignmentOperationName, actual.AssignmentOperationName)
		}

	}
}

func TestParseScopedAssignmentOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedAssignmentOperationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtAsSiGnMeNtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtAsSiGnMeNtS/aSsIgNmEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtAsSiGnMeNtS/aSsIgNmEnTnAmE/aSsIgNmEnToPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations/assignmentOperationName",
			Expected: &ScopedAssignmentOperationId{
				ResourceScope:           "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				BlueprintAssignmentName: "assignmentName",
				AssignmentOperationName: "assignmentOperationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Blueprint/blueprintAssignments/assignmentName/assignmentOperations/assignmentOperationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtAsSiGnMeNtS/aSsIgNmEnTnAmE/aSsIgNmEnToPeRaTiOnS/aSsIgNmEnToPeRaTiOnNaMe",
			Expected: &ScopedAssignmentOperationId{
				ResourceScope:           "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				BlueprintAssignmentName: "aSsIgNmEnTnAmE",
				AssignmentOperationName: "aSsIgNmEnToPeRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.bLuEpRiNt/bLuEpRiNtAsSiGnMeNtS/aSsIgNmEnTnAmE/aSsIgNmEnToPeRaTiOnS/aSsIgNmEnToPeRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedAssignmentOperationIDInsensitively(v.Input)
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

		if actual.BlueprintAssignmentName != v.Expected.BlueprintAssignmentName {
			t.Fatalf("Expected %q but got %q for BlueprintAssignmentName", v.Expected.BlueprintAssignmentName, actual.BlueprintAssignmentName)
		}

		if actual.AssignmentOperationName != v.Expected.AssignmentOperationName {
			t.Fatalf("Expected %q but got %q for AssignmentOperationName", v.Expected.AssignmentOperationName, actual.AssignmentOperationName)
		}

	}
}

func TestSegmentsForScopedAssignmentOperationId(t *testing.T) {
	segments := ScopedAssignmentOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedAssignmentOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
