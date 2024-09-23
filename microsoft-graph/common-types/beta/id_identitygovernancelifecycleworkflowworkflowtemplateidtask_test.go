package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateId", "taskId")

	if id.WorkflowTemplateId != "workflowTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowTemplateId'", id.WorkflowTemplateId, "workflowTemplateId")
	}

	if id.TaskId != "taskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateId", "taskId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId
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
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{
				WorkflowTemplateId: "workflowTemplateId",
				TaskId:             "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID(v.Input)
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

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId
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
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{
				WorkflowTemplateId: "workflowTemplateId",
				TaskId:             "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/tasks/taskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs/tAsKiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{
				WorkflowTemplateId: "wOrKfLoWtEmPlAtEiD",
				TaskId:             "tAsKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/tAsKs/tAsKiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
