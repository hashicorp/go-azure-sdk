package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowId", "taskReportId", "taskProcessingResultId")

	if id.WorkflowId != "workflowId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowId")
	}

	if id.TaskReportId != "taskReportId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskReportId'", id.TaskReportId, "taskReportId")
	}

	if id.TaskProcessingResultId != "taskProcessingResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowId", "taskReportId", "taskProcessingResultId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults/taskProcessingResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{
				WorkflowId:             "workflowId",
				TaskReportId:           "taskReportId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID(v.Input)
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

		if actual.TaskReportId != v.Expected.TaskReportId {
			t.Fatalf("Expected %q but got %q for TaskReportId", v.Expected.TaskReportId, actual.TaskReportId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/tAsKrEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/tAsKrEpOrTs/tAsKrEpOrTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/tAsKrEpOrTs/tAsKrEpOrTiD/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults/taskProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{
				WorkflowId:             "workflowId",
				TaskReportId:           "taskReportId",
				TaskProcessingResultId: "taskProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/taskReports/taskReportId/taskProcessingResults/taskProcessingResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/tAsKrEpOrTs/tAsKrEpOrTiD/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{
				WorkflowId:             "wOrKfLoWiD",
				TaskReportId:           "tAsKrEpOrTiD",
				TaskProcessingResultId: "tAsKpRoCeSsInGrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/tAsKrEpOrTs/tAsKrEpOrTiD/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultIDInsensitively(v.Input)
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

		if actual.TaskReportId != v.Expected.TaskReportId {
			t.Fatalf("Expected %q but got %q for TaskReportId", v.Expected.TaskReportId, actual.TaskReportId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
