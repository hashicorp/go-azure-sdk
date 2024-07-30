package lifecycleworkflowworkflowversion

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID("workflowIdValue", "workflowVersionVersionNumberValue")

	if id.WorkflowId != "workflowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowIdValue")
	}

	if id.WorkflowVersionVersionNumber != "workflowVersionVersionNumberValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowVersionVersionNumber'", id.WorkflowVersionVersionNumber, "workflowVersionVersionNumberValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID("workflowIdValue", "workflowVersionVersionNumberValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{
				WorkflowId:                   "workflowIdValue",
				WorkflowVersionVersionNumber: "workflowVersionVersionNumberValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionID(v.Input)
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

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{
				WorkflowId:                   "workflowIdValue",
				WorkflowVersionVersionNumber: "workflowVersionVersionNumberValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{
				WorkflowId:                   "wOrKfLoWiDvAlUe",
				WorkflowVersionVersionNumber: "wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdVersionId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
