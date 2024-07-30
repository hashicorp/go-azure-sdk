package lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{}

func TestNewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

	if id.WorkflowId != "workflowIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkflowId'", id.WorkflowId, "workflowIdValue")
	}

	if id.TaskReportId != "taskReportIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskReportId'", id.TaskReportId, "taskReportIdValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports/taskReportIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId
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
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports/taskReportIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{
				WorkflowId:   "workflowIdValue",
				TaskReportId: "taskReportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports/taskReportIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID(v.Input)
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

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId
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
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiDvAlUe/tAsKrEpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports/taskReportIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{
				WorkflowId:   "workflowIdValue",
				TaskReportId: "taskReportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/deletedItems/workflows/workflowIdValue/taskReports/taskReportIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiDvAlUe/tAsKrEpOrTs/tAsKrEpOrTiDvAlUe",
			Expected: &IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{
				WorkflowId:   "wOrKfLoWiDvAlUe",
				TaskReportId: "tAsKrEpOrTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/dElEtEdItEmS/wOrKfLoWs/wOrKfLoWiDvAlUe/tAsKrEpOrTs/tAsKrEpOrTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
