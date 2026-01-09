package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateId", "taskId", "taskProcessingResultId")

	if id.WorkflowTemplateId != "workflowTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowTemplateId'", id.WorkflowTemplateId, "workflowTemplateId")
	}

	if id.TaskId != "taskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskId")
	}

	if id.TaskProcessingResultId != "taskProcessingResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateId", "taskId", "taskProcessingResultId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults/taskProcessingResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{
				WorkflowTemplateId:     "workflowTemplateId",
				TaskId:                 "taskId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WorkflowTemplateId != v.Expected.WorkflowTemplateId {
			t.Fatalf("Expected %q but got %q for WorkflowTemplateId", v.Expected.WorkflowTemplateId, actual.WorkflowTemplateId)
		}

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs/tAsKiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs/tAsKiD/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{
				WorkflowTemplateId:     "workflowTemplateId",
				TaskId:                 "taskId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs/tAsKiD/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{
				WorkflowTemplateId:     "wOrKfLoWtEmPlAtEiD",
				TaskId:                 "tAsKiD",
				TaskProcessingResultId: "tAsKpRoCeSsInGrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs/tAsKiD/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.WorkflowTemplateId != v.Expected.WorkflowTemplateId {
			t.Fatalf("Expected %q but got %q for WorkflowTemplateId", v.Expected.WorkflowTemplateId, actual.WorkflowTemplateId)
		}

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
