package get

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProjectId{}

func TestNewProjectID(t *testing.T) {
	id := NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "resourceGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "resourceGroupName")
	}

	if id.ServiceName != "serviceName" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceName'", id.ServiceName, "serviceName")
	}

	if id.ProjectName != "projectName" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectName'", id.ProjectName, "projectName")
	}
}

func TestFormatProjectID(t *testing.T) {
	actual := NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects/projectName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProjectID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProjectId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects/projectName",
			Expected: &ProjectId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "resourceGroupName",
				ServiceName:       "serviceName",
				ProjectName:       "projectName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects/projectName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProjectID(v.Input)
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

		if actual.ServiceName != v.Expected.ServiceName {
			t.Fatalf("Expected %q but got %q for ServiceName", v.Expected.ServiceName, actual.ServiceName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

	}
}

func TestParseProjectIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProjectId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeNaMe/pRoJeCtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects/projectName",
			Expected: &ProjectId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "resourceGroupName",
				ServiceName:       "serviceName",
				ProjectName:       "projectName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupName/providers/Microsoft.DataMigration/services/serviceName/projects/projectName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeNaMe/pRoJeCtS/pRoJeCtNaMe",
			Expected: &ProjectId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "rEsOuRcEgRoUpNaMe",
				ServiceName:       "sErViCeNaMe",
				ProjectName:       "pRoJeCtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpNaMe/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeNaMe/pRoJeCtS/pRoJeCtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProjectIDInsensitively(v.Input)
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

		if actual.ServiceName != v.Expected.ServiceName {
			t.Fatalf("Expected %q but got %q for ServiceName", v.Expected.ServiceName, actual.ServiceName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

	}
}

func TestSegmentsForProjectId(t *testing.T) {
	segments := ProjectId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProjectId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
