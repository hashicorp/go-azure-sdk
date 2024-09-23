package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowId", "runId", "userProcessingResultId")

	if id.WorkflowId != "workflowId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowId")
	}

	if id.RunId != "runId" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runId")
	}

	if id.UserProcessingResultId != "userProcessingResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserProcessingResultId'", id.UserProcessingResultId, "userProcessingResultId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowId", "runId", "userProcessingResultId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults/userProcessingResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults/userProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId{
				WorkflowId:             "workflowId",
				RunId:                  "runId",
				UserProcessingResultId: "userProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults/userProcessingResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID(v.Input)
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

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/uSeRpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults/userProcessingResultId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId{
				WorkflowId:             "workflowId",
				RunId:                  "runId",
				UserProcessingResultId: "userProcessingResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowId/runs/runId/userProcessingResults/userProcessingResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId{
				WorkflowId:             "wOrKfLoWiD",
				RunId:                  "rUnId",
				UserProcessingResultId: "uSeRpRoCeSsInGrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/uSeRpRoCeSsInGrEsUlTs/uSeRpRoCeSsInGrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
