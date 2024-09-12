package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

	if id.WorkflowId != "workflowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowIdValue")
	}

	if id.RunId != "runIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runIdValue")
	}

	if id.UserProcessingResultId != "userProcessingResultIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserProcessingResultId'", id.UserProcessingResultId, "userProcessingResultIdValue")
	}

	if id.TaskProcessingResultId != "taskProcessingResultIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultIdValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults/taskProcessingResultIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults/taskProcessingResultIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{
				WorkflowId:             "workflowIdValue",
				RunId:                  "runIdValue",
				UserProcessingResultId: "userProcessingResultIdValue",
				TaskProcessingResultId: "taskProcessingResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults/taskProcessingResultIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID(v.Input)
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

		if actual.UserProcessingResultId != v.Expected.UserProcessingResultId {
			t.Fatalf("Expected %q but got %q for UserProcessingResultId", v.Expected.UserProcessingResultId, actual.UserProcessingResultId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS/rUnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS/rUnIdVaLuE/uSeRpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS/rUnIdVaLuE/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS/rUnIdVaLuE/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiDvAlUe/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults/taskProcessingResultIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{
				WorkflowId:             "workflowIdValue",
				RunId:                  "runIdValue",
				UserProcessingResultId: "userProcessingResultIdValue",
				TaskProcessingResultId: "taskProcessingResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/runs/runIdValue/userProcessingResults/userProcessingResultIdValue/taskProcessingResults/taskProcessingResultIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS/rUnIdVaLuE/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiDvAlUe/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiDvAlUe",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{
				WorkflowId:             "wOrKfLoWiDvAlUe",
				RunId:                  "rUnIdVaLuE",
				UserProcessingResultId: "uSeRpRoCeSsInGrEsUlTiDvAlUe",
				TaskProcessingResultId: "tAsKpRoCeSsInGrEsUlTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/rUnS/rUnIdVaLuE/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiDvAlUe/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultIDInsensitively(v.Input)
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

		if actual.UserProcessingResultId != v.Expected.UserProcessingResultId {
			t.Fatalf("Expected %q but got %q for UserProcessingResultId", v.Expected.UserProcessingResultId, actual.UserProcessingResultId)
		}

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
