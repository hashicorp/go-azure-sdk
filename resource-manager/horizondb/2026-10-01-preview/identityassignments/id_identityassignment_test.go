package identityassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityAssignmentId{}

func TestNewIdentityAssignmentID(t *testing.T) {
	id := NewIdentityAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName", "identityAssignmentName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ClusterName != "clusterName" {
		t.Fatalf("Expected %q but got %q for Segment 'ClusterName'", id.ClusterName, "clusterName")
	}

	if id.ChangeDataCaptureDestinationName != "changeDataCaptureDestinationName" {
		t.Fatalf("Expected %q but got %q for Segment 'ChangeDataCaptureDestinationName'", id.ChangeDataCaptureDestinationName, "changeDataCaptureDestinationName")
	}

	if id.IdentityAssignmentName != "identityAssignmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'IdentityAssignmentName'", id.IdentityAssignmentName, "identityAssignmentName")
	}
}

func TestFormatIdentityAssignmentID(t *testing.T) {
	actual := NewIdentityAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "changeDataCaptureDestinationName", "identityAssignmentName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments/identityAssignmentName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments/identityAssignmentName",
			Expected: &IdentityAssignmentId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				ClusterName:                      "clusterName",
				ChangeDataCaptureDestinationName: "changeDataCaptureDestinationName",
				IdentityAssignmentName:           "identityAssignmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments/identityAssignmentName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAssignmentID(v.Input)
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

		if actual.ClusterName != v.Expected.ClusterName {
			t.Fatalf("Expected %q but got %q for ClusterName", v.Expected.ClusterName, actual.ClusterName)
		}

		if actual.ChangeDataCaptureDestinationName != v.Expected.ChangeDataCaptureDestinationName {
			t.Fatalf("Expected %q but got %q for ChangeDataCaptureDestinationName", v.Expected.ChangeDataCaptureDestinationName, actual.ChangeDataCaptureDestinationName)
		}

		if actual.IdentityAssignmentName != v.Expected.IdentityAssignmentName {
			t.Fatalf("Expected %q but got %q for IdentityAssignmentName", v.Expected.IdentityAssignmentName, actual.IdentityAssignmentName)
		}

	}
}

func TestParseIdentityAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityAssignmentId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS/cLuStErNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS/cLuStErNaMe/cHaNgEdAtAcApTuReDeStInAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS/cLuStErNaMe/cHaNgEdAtAcApTuReDeStInAtIoNs/cHaNgEdAtAcApTuReDeStInAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS/cLuStErNaMe/cHaNgEdAtAcApTuReDeStInAtIoNs/cHaNgEdAtAcApTuReDeStInAtIoNnAmE/iDeNtItYaSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments/identityAssignmentName",
			Expected: &IdentityAssignmentId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "example-resource-group",
				ClusterName:                      "clusterName",
				ChangeDataCaptureDestinationName: "changeDataCaptureDestinationName",
				IdentityAssignmentName:           "identityAssignmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/clusterName/changeDataCaptureDestinations/changeDataCaptureDestinationName/identityAssignments/identityAssignmentName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS/cLuStErNaMe/cHaNgEdAtAcApTuReDeStInAtIoNs/cHaNgEdAtAcApTuReDeStInAtIoNnAmE/iDeNtItYaSsIgNmEnTs/iDeNtItYaSsIgNmEnTnAmE",
			Expected: &IdentityAssignmentId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:                "eXaMpLe-rEsOuRcE-GrOuP",
				ClusterName:                      "cLuStErNaMe",
				ChangeDataCaptureDestinationName: "cHaNgEdAtAcApTuReDeStInAtIoNnAmE",
				IdentityAssignmentName:           "iDeNtItYaSsIgNmEnTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.hOrIzOnDb/cLuStErS/cLuStErNaMe/cHaNgEdAtAcApTuReDeStInAtIoNs/cHaNgEdAtAcApTuReDeStInAtIoNnAmE/iDeNtItYaSsIgNmEnTs/iDeNtItYaSsIgNmEnTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityAssignmentIDInsensitively(v.Input)
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

		if actual.ClusterName != v.Expected.ClusterName {
			t.Fatalf("Expected %q but got %q for ClusterName", v.Expected.ClusterName, actual.ClusterName)
		}

		if actual.ChangeDataCaptureDestinationName != v.Expected.ChangeDataCaptureDestinationName {
			t.Fatalf("Expected %q but got %q for ChangeDataCaptureDestinationName", v.Expected.ChangeDataCaptureDestinationName, actual.ChangeDataCaptureDestinationName)
		}

		if actual.IdentityAssignmentName != v.Expected.IdentityAssignmentName {
			t.Fatalf("Expected %q but got %q for IdentityAssignmentName", v.Expected.IdentityAssignmentName, actual.IdentityAssignmentName)
		}

	}
}

func TestSegmentsForIdentityAssignmentId(t *testing.T) {
	segments := IdentityAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %d unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
