package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowId", "runId", "taskProcessingResultId")

	if id.WorkflowId != "workflowId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowId")
	}

	if id.RunId != "runId" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runId")
	}

	if id.TaskProcessingResultId != "taskProcessingResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowId", "runId", "taskProcessingResultId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults/taskProcessingResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId{
				WorkflowId:             "workflowId",
				RunId:                  "runId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WorkflowId != v.Expected.WorkflowId {
			t.Fatalf("Expected %q but got %q for WorkflowId", v.Expected.WorkflowId, actual.WorkflowId)
		}

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId{
				WorkflowId:             "workflowId",
				RunId:                  "runId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId{
				WorkflowId:             "wOrKfLoWiD",
				RunId:                  "rUnId",
				TaskProcessingResultId: "tAsKpRoCeSsInGrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WorkflowId != v.Expected.WorkflowId {
			t.Fatalf("Expected %q but got %q for WorkflowId", v.Expected.WorkflowId, actual.WorkflowId)
		}

		if actual.RunId != v.Expected.RunId {
			t.Fatalf("Expected %q but got %q for RunId", v.Expected.RunId, actual.RunId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
