package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{}

func TestNewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID("workflowId", "runId")

	if id.WorkflowId != "workflowId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowId")
	}

	if id.RunId != "runId" {
		t.Fatalf("Expected %q but got %q for Segment 'RunId'", id.RunId, "runId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID("workflowId", "runId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs/runId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId
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
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs/runId",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{
				WorkflowId: "workflowId",
				RunId:      "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs/runId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID(v.Input)
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

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId
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
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/rUnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs/runId",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{
				WorkflowId: "workflowId",
				RunId:      "runId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowId/runs/runId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{
				WorkflowId: "wOrKfLoWiD",
				RunId:      "rUnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiD/rUnS/rUnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
