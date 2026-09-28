package goldengatedeployments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AssignedConnectionId{}

func TestNewAssignedConnectionID(t *testing.T) {
	id := NewAssignedConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "goldenGateDeploymentName", "assignmentId")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.GoldenGateDeploymentName != "goldenGateDeploymentName" {
		t.Fatalf("Expected %q but got %q for Segment 'GoldenGateDeploymentName'", id.GoldenGateDeploymentName, "goldenGateDeploymentName")
	}

	if id.AssignmentId != "assignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentId'", id.AssignmentId, "assignmentId")
	}
}

func TestFormatAssignedConnectionID(t *testing.T) {
	actual := NewAssignedConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "goldenGateDeploymentName", "assignmentId").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections/assignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAssignedConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AssignedConnectionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections/assignmentId",
			Expected: &AssignedConnectionId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				GoldenGateDeploymentName: "goldenGateDeploymentName",
				AssignmentId:             "assignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections/assignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAssignedConnectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.GoldenGateDeploymentName != v.Expected.GoldenGateDeploymentName {
			t.Fatalf("Expected %q but got %q for GoldenGateDeploymentName", v.Expected.GoldenGateDeploymentName, actual.GoldenGateDeploymentName)
		}

		if actual.AssignmentId != v.Expected.AssignmentId {
			t.Fatalf("Expected %q but got %q for AssignmentId", v.Expected.AssignmentId, actual.AssignmentId)
		}

	}
}

func TestParseAssignedConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AssignedConnectionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/gOlDeNgAtEdEpLoYmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/gOlDeNgAtEdEpLoYmEnTs/gOlDeNgAtEdEpLoYmEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/gOlDeNgAtEdEpLoYmEnTs/gOlDeNgAtEdEpLoYmEnTnAmE/aSsIgNeDcOnNeCtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections/assignmentId",
			Expected: &AssignedConnectionId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "example-resource-group",
				GoldenGateDeploymentName: "goldenGateDeploymentName",
				AssignmentId:             "assignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Oracle.Database/goldenGateDeployments/goldenGateDeploymentName/assignedConnections/assignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/gOlDeNgAtEdEpLoYmEnTs/gOlDeNgAtEdEpLoYmEnTnAmE/aSsIgNeDcOnNeCtIoNs/aSsIgNmEnTiD",
			Expected: &AssignedConnectionId{
				SubscriptionId:           "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:        "eXaMpLe-rEsOuRcE-GrOuP",
				GoldenGateDeploymentName: "gOlDeNgAtEdEpLoYmEnTnAmE",
				AssignmentId:             "aSsIgNmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/oRaClE.DaTaBaSe/gOlDeNgAtEdEpLoYmEnTs/gOlDeNgAtEdEpLoYmEnTnAmE/aSsIgNeDcOnNeCtIoNs/aSsIgNmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAssignedConnectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.GoldenGateDeploymentName != v.Expected.GoldenGateDeploymentName {
			t.Fatalf("Expected %q but got %q for GoldenGateDeploymentName", v.Expected.GoldenGateDeploymentName, actual.GoldenGateDeploymentName)
		}

		if actual.AssignmentId != v.Expected.AssignmentId {
			t.Fatalf("Expected %q but got %q for AssignmentId", v.Expected.AssignmentId, actual.AssignmentId)
		}

	}
}

func TestSegmentsForAssignedConnectionId(t *testing.T) {
	segments := AssignedConnectionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AssignedConnectionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %d unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
