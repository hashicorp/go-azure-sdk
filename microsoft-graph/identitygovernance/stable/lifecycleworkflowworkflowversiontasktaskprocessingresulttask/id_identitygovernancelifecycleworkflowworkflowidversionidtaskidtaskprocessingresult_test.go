package lifecycleworkflowworkflowversiontasktaskprocessingresulttask

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{}

func TestNewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

	if id.WorkflowId != "workflowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowIdValue")
	}

	if id.WorkflowVersionVersionNumber != "workflowVersionVersionNumberValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowVersionVersionNumber'", id.WorkflowVersionVersionNumber, "workflowVersionVersionNumberValue")
	}

	if id.TaskId != "taskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskId'", id.TaskId, "taskIdValue")
	}

	if id.TaskProcessingResultId != "taskProcessingResultIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskProcessingResultId'", id.TaskProcessingResultId, "taskProcessingResultIdValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId
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
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{
				WorkflowId:                   "workflowIdValue",
				WorkflowVersionVersionNumber: "workflowVersionVersionNumberValue",
				TaskId:                       "taskIdValue",
				TaskProcessingResultId:       "taskProcessingResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID(v.Input)
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

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId
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
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe/tAsKs/tAsKiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe/tAsKs/tAsKiDvAlUe/tAsKpRoCeSsInGrEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{
				WorkflowId:                   "workflowIdValue",
				WorkflowVersionVersionNumber: "workflowVersionVersionNumberValue",
				TaskId:                       "taskIdValue",
				TaskProcessingResultId:       "taskProcessingResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/workflows/workflowIdValue/versions/workflowVersionVersionNumberValue/tasks/taskIdValue/taskProcessingResults/taskProcessingResultIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe/tAsKs/tAsKiDvAlUe/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiDvAlUe",
			Expected: &IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{
				WorkflowId:                   "wOrKfLoWiDvAlUe",
				WorkflowVersionVersionNumber: "wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe",
				TaskId:                       "tAsKiDvAlUe",
				TaskProcessingResultId:       "tAsKpRoCeSsInGrEsUlTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/wOrKfLoWs/wOrKfLoWiDvAlUe/vErSiOnS/wOrKfLoWvErSiOnVeRsIoNnUmBeRvAlUe/tAsKs/tAsKiDvAlUe/tAsKpRoCeSsInGrEsUlTs/tAsKpRoCeSsInGrEsUlTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultIDInsensitively(v.Input)
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

		if actual.TaskProcessingResultId != v.Expected.TaskProcessingResultId {
			t.Fatalf("Expected %q but got %q for TaskProcessingResultId", v.Expected.TaskProcessingResultId, actual.TaskProcessingResultId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
