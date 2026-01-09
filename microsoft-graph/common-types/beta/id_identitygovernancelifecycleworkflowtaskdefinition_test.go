package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowTaskDefinitionId{}

func TestNewIdentityGovernanceLifecycleWorkflowTaskDefinitionID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowTaskDefinitionID("taskDefinitionId")

	if id.TaskDefinitionId != "taskDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'TaskDefinitionId'", id.TaskDefinitionId, "taskDefinitionId")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowTaskDefinitionID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowTaskDefinitionID("taskDefinitionId").ID()
	expected := "/identityGovernance/lifecycleWorkflows/taskDefinitions/taskDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowTaskDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowTaskDefinitionId
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
			Input: "/identityGovernance/lifecycleWorkflows/taskDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/taskDefinitions/taskDefinitionId",
			Expected: &IdentityGovernanceLifecycleWorkflowTaskDefinitionId{
				TaskDefinitionId: "taskDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/taskDefinitions/taskDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TaskDefinitionId != v.Expected.TaskDefinitionId {
			t.Fatalf("Expected %q but got %q for TaskDefinitionId", v.Expected.TaskDefinitionId, actual.TaskDefinitionId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowTaskDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowTaskDefinitionId
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
			Input: "/identityGovernance/lifecycleWorkflows/taskDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/tAsKdEfInItIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/taskDefinitions/taskDefinitionId",
			Expected: &IdentityGovernanceLifecycleWorkflowTaskDefinitionId{
				TaskDefinitionId: "taskDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/taskDefinitions/taskDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/tAsKdEfInItIoNs/tAsKdEfInItIoNiD",
			Expected: &IdentityGovernanceLifecycleWorkflowTaskDefinitionId{
				TaskDefinitionId: "tAsKdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/tAsKdEfInItIoNs/tAsKdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowTaskDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TaskDefinitionId != v.Expected.TaskDefinitionId {
			t.Fatalf("Expected %q but got %q for TaskDefinitionId", v.Expected.TaskDefinitionId, actual.TaskDefinitionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowTaskDefinitionId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowTaskDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowTaskDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
