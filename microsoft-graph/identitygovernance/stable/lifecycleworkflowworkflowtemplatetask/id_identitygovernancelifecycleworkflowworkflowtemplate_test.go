package lifecycleworkflowworkflowtemplatetask

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID("workflowTemplateIdValue")

	if id.WorkflowTemplateId != "workflowTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowTemplateId'", id.WorkflowTemplateId, "workflowTemplateIdValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowTemplateID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateID("workflowTemplateIdValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue"
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
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
				WorkflowTemplateId: "workflowTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/extra",
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
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
				WorkflowTemplateId: "workflowTemplateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateId{
				WorkflowTemplateId: "wOrKfLoWtEmPlAtEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe/extra",
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
