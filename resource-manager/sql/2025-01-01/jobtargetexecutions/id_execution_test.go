package jobtargetexecutions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExecutionId{}

func TestNewExecutionID(t *testing.T) {
	id := NewExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "jobName", "jobExecutionId")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ServerName != "serverName" {
		t.Fatalf("Expected %q but got %q for Segment 'ServerName'", id.ServerName, "serverName")
	}

	if id.JobAgentName != "jobAgentName" {
		t.Fatalf("Expected %q but got %q for Segment 'JobAgentName'", id.JobAgentName, "jobAgentName")
	}

	if id.JobName != "jobName" {
		t.Fatalf("Expected %q but got %q for Segment 'JobName'", id.JobName, "jobName")
	}

	if id.JobExecutionId != "jobExecutionId" {
		t.Fatalf("Expected %q but got %q for Segment 'JobExecutionId'", id.JobExecutionId, "jobExecutionId")
	}
}

func TestFormatExecutionID(t *testing.T) {
	actual := NewExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "jobName", "jobExecutionId").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions/jobExecutionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseExecutionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExecutionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions/jobExecutionId",
			Expected: &ExecutionId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				ServerName:        "serverName",
				JobAgentName:      "jobAgentName",
				JobName:           "jobName",
				JobExecutionId:    "jobExecutionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions/jobExecutionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExecutionID(v.Input)
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

		if actual.ServerName != v.Expected.ServerName {
			t.Fatalf("Expected %q but got %q for ServerName", v.Expected.ServerName, actual.ServerName)
		}

		if actual.JobAgentName != v.Expected.JobAgentName {
			t.Fatalf("Expected %q but got %q for JobAgentName", v.Expected.JobAgentName, actual.JobAgentName)
		}

		if actual.JobName != v.Expected.JobName {
			t.Fatalf("Expected %q but got %q for JobName", v.Expected.JobName, actual.JobName)
		}

		if actual.JobExecutionId != v.Expected.JobExecutionId {
			t.Fatalf("Expected %q but got %q for JobExecutionId", v.Expected.JobExecutionId, actual.JobExecutionId)
		}

	}
}

func TestParseExecutionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ExecutionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs/jObAgEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs/jObAgEnTnAmE/jObS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs/jObAgEnTnAmE/jObS/jObNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs/jObAgEnTnAmE/jObS/jObNaMe/eXeCuTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions/jobExecutionId",
			Expected: &ExecutionId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				ServerName:        "serverName",
				JobAgentName:      "jobAgentName",
				JobName:           "jobName",
				JobExecutionId:    "jobExecutionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Sql/servers/serverName/jobAgents/jobAgentName/jobs/jobName/executions/jobExecutionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs/jObAgEnTnAmE/jObS/jObNaMe/eXeCuTiOnS/jObExEcUtIoNiD",
			Expected: &ExecutionId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-rEsOuRcE-GrOuP",
				ServerName:        "sErVeRnAmE",
				JobAgentName:      "jObAgEnTnAmE",
				JobName:           "jObNaMe",
				JobExecutionId:    "jObExEcUtIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sQl/sErVeRs/sErVeRnAmE/jObAgEnTs/jObAgEnTnAmE/jObS/jObNaMe/eXeCuTiOnS/jObExEcUtIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseExecutionIDInsensitively(v.Input)
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

		if actual.ServerName != v.Expected.ServerName {
			t.Fatalf("Expected %q but got %q for ServerName", v.Expected.ServerName, actual.ServerName)
		}

		if actual.JobAgentName != v.Expected.JobAgentName {
			t.Fatalf("Expected %q but got %q for JobAgentName", v.Expected.JobAgentName, actual.JobAgentName)
		}

		if actual.JobName != v.Expected.JobName {
			t.Fatalf("Expected %q but got %q for JobName", v.Expected.JobName, actual.JobName)
		}

		if actual.JobExecutionId != v.Expected.JobExecutionId {
			t.Fatalf("Expected %q but got %q for JobExecutionId", v.Expected.JobExecutionId, actual.JobExecutionId)
		}

	}
}

func TestSegmentsForExecutionId(t *testing.T) {
	segments := ExecutionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ExecutionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
