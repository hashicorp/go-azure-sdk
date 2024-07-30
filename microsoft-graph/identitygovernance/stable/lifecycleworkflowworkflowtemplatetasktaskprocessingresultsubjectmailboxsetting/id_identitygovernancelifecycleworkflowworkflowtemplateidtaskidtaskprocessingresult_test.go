package lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectmailboxsetting

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

	if id.WorkflowTemplateId != "workflowTemplateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowTemplateId'", id.WorkflowTemplateId, "workflowTemplateIdValue")
	}

	if id.TaskId != "taskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskIdValue")
	}

	if id.TaskProcessingResultId != "taskProcessingResultIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultIdValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{
				WorkflowTemplateId:     "workflowTemplateIdValue",
				TaskId:                 "taskIdValue",
				TaskProcessingResultId: "taskProcessingResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID(v.Input)
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

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId
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
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe/tAsKs/tAsKiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe/tAsKs/tAsKiDvAlUe/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{
				WorkflowTemplateId:     "workflowTemplateIdValue",
				TaskId:                 "taskIdValue",
				TaskProcessingResultId: "taskProcessingResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflowTemplates/workflowTemplateIdValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe/tAsKs/tAsKiDvAlUe/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiDvAlUe",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{
				WorkflowTemplateId:     "wOrKfLoWtEmPlAtEiDvAlUe",
				TaskId:                 "tAsKiDvAlUe",
				TaskProcessingResultId: "tAsKpRoCeSsInGrEsUlTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWtEmPlAtEs/wOrKfLoWtEmPlAtEiDvAlUe/tAsKs/tAsKiDvAlUe/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultIDInsensitively(v.Input)
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

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
