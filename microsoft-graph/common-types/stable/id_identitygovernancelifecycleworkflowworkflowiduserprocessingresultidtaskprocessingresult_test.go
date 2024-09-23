package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowId", "userProcessingResultId", "taskProcessingResultId")

	if id.WorkflowId != "workflowId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowId")
	}

	if id.UserProcessingResultId != "userProcessingResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserProcessingResultId'", id.UserProcessingResultId, "userProcessingResultId")
	}

	if id.TaskProcessingResultId != "taskProcessingResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowId", "userProcessingResultId", "taskProcessingResultId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults/taskProcessingResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId{
				WorkflowId:             "workflowId",
				UserProcessingResultId: "userProcessingResultId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID(v.Input)
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

		if actual.UserProcessingResultId != v.Expected.UserProcessingResultId {
			t.Fatalf("Expected %q but got %q for UserProcessingResultId", v.Expected.UserProcessingResultId, actual.UserProcessingResultId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/uSeRpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiD/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId{
				WorkflowId:             "workflowId",
				UserProcessingResultId: "userProcessingResultId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/userProcessingResults/userProcessingResultId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiD/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId{
				WorkflowId:             "wOrKfLoWiD",
				UserProcessingResultId: "uSeRpRoCeSsInGrEsUlTiD",
				TaskProcessingResultId: "tAsKpRoCeSsInGrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiD/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultIDInsensitively(v.Input)
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

		if actual.UserProcessingResultId != v.Expected.UserProcessingResultId {
			t.Fatalf("Expected %q but got %q for UserProcessingResultId", v.Expected.UserProcessingResultId, actual.UserProcessingResultId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
