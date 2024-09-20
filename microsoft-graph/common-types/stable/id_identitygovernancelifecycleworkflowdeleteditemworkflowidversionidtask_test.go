package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId{}

func TestNewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowId", "workflowVersionVersionNumber", "taskId")

	if id.WorkflowId != "workflowId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowId")
	}

	if id.WorkflowVersionVersionNumber != "workflowVersionVersionNumber" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowVersionVersionNumber'", id.WorkflowVersionVersionNumber, "workflowVersionVersionNumber")
	}

	if id.TaskId != "taskId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowId", "workflowVersionVersionNumber", "taskId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks/taskId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId
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
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks/taskId",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId{
				WorkflowId:                   "workflowId",
				WorkflowVersionVersionNumber: "workflowVersionVersionNumber",
				TaskId:                       "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks/taskId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID(v.Input)
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

		if actual.WorkflowVersionVersionNumber != v.Expected.WorkflowVersionVersionNumber {
			t.Fatalf("Expected %q but got %q for WorkflowVersionVersionNumber", v.Expected.WorkflowVersionVersionNumber, actual.WorkflowVersionVersionNumber)
		}

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId
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
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/vErSiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeR/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks/taskId",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId{
				WorkflowId:                   "workflowId",
				WorkflowVersionVersionNumber: "workflowVersionVersionNumber",
				TaskId:                       "taskId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/versions/workflowVersionVersionNumber/tasks/taskId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeR/tAsKs/tAsKiD",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId{
				WorkflowId:                   "wOrKfLoWiD",
				WorkflowVersionVersionNumber: "wOrKfLoWvErSiOnVeRsIoNnUmBeR",
				TaskId:                       "tAsKiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeR/tAsKs/tAsKiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIDInsensitively(v.Input)
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

		if actual.WorkflowVersionVersionNumber != v.Expected.WorkflowVersionVersionNumber {
			t.Fatalf("Expected %q but got %q for WorkflowVersionVersionNumber", v.Expected.WorkflowVersionVersionNumber, actual.WorkflowVersionVersionNumber)
		}

		if actual.TaskId != v.Expected.TaskId {
			t.Fatalf("Expected %q but got %q for TaskId", v.Expected.TaskId, actual.TaskId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
