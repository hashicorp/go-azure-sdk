package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID("workflowTemplateId")

	if id.WorkflowTemplateId != "workflowTemplateId" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowTemplateId'", id.WorkflowTemplateId, "workflowTemplateId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID("workflowTemplateId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateId
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
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
				WorkflowTemplateId: "workflowTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(v.Input)
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

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateId
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
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
				WorkflowTemplateId: "workflowTemplateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
				WorkflowTemplateId: "wOrKfLoWtEmPlAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowTemplateId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowTemplateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
