// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package put

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = TaskId{}

func TestNewTaskID(t *testing.T) {
	id := NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "resourceGroupValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "resourceGroupValue")
	}

	if id.ServiceName != "serviceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ServiceName'", id.ServiceName, "serviceValue")
	}

	if id.ProjectName != "projectValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectName'", id.ProjectName, "projectValue")
	}

	if id.TaskName != "taskValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskName'", id.TaskName, "taskValue")
	}
}

func TestFormatTaskID(t *testing.T) {
	actual := NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks/taskValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TaskId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks/taskValue",
			Expected: &TaskId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "resourceGroupValue",
				ServiceName:       "serviceValue",
				ProjectName:       "projectValue",
				TaskName:          "taskValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks/taskValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTaskID(v.Input)
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

		if actual.TaskName != v.Expected.TaskName {
			t.Fatalf("Expected %q but got %q for TaskName", v.Expected.TaskName, actual.TaskName)
		}

	}
}

func TestParseTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *TaskId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeVaLuE/pRoJeCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeVaLuE/pRoJeCtS/pRoJeCtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeVaLuE/pRoJeCtS/pRoJeCtVaLuE/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks/taskValue",
			Expected: &TaskId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "resourceGroupValue",
				ServiceName:       "serviceValue",
				ProjectName:       "projectValue",
				TaskName:          "taskValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroupValue/providers/Microsoft.DataMigration/services/serviceValue/projects/projectValue/tasks/taskValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeVaLuE/pRoJeCtS/pRoJeCtVaLuE/tAsKs/tAsKvAlUe",
			Expected: &TaskId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "rEsOuRcEgRoUpVaLuE",
				ServiceName:       "sErViCeVaLuE",
				ProjectName:       "pRoJeCtVaLuE",
				TaskName:          "tAsKvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/rEsOuRcEgRoUpVaLuE/pRoViDeRs/mIcRoSoFt.dAtAmIgRaTiOn/sErViCeS/sErViCeVaLuE/pRoJeCtS/pRoJeCtVaLuE/tAsKs/tAsKvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseTaskIDInsensitively(v.Input)
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

		if actual.TaskName != v.Expected.TaskName {
			t.Fatalf("Expected %q but got %q for TaskName", v.Expected.TaskName, actual.TaskName)
		}

	}
}

func TestSegmentsForTaskId(t *testing.T) {
	segments := TaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("TaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
