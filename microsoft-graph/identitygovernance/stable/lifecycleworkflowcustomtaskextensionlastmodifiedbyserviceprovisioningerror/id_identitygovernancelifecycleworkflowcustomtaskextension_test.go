package lifecycleworkflowcustomtaskextensionlastmodifiedbyserviceprovisioningerror

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{}

func TestNewIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(t *testing.T) {
	id := NewIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID("customTaskExtensionIdValue")

	if id.CustomTaskExtensionId != "customTaskExtensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomTaskExtensionId'", id.CustomTaskExtensionId, "customTaskExtensionIdValue")
	}
}

func TestFormatIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(t *testing.T) {
	actual := NewIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID("customTaskExtensionIdValue").ID()
	expected := "/identityGovernance/lifecycleWorkflows/customTaskExtensions/customTaskExtensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId
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
			Input: "/identityGovernance/lifecycleWorkflows/customTaskExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/customTaskExtensions/customTaskExtensionIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{
				CustomTaskExtensionId: "customTaskExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/customTaskExtensions/customTaskExtensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomTaskExtensionId != v.Expected.CustomTaskExtensionId {
			t.Fatalf("Expected %q but got %q for CustomTaskExtensionId", v.Expected.CustomTaskExtensionId, actual.CustomTaskExtensionId)
		}

	}
}

func TestParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId
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
			Input: "/identityGovernance/lifecycleWorkflows/customTaskExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/cUsToMtAsKeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/lifecycleWorkflows/customTaskExtensions/customTaskExtensionIdValue",
			Expected: &IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{
				CustomTaskExtensionId: "customTaskExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/lifecycleWorkflows/customTaskExtensions/customTaskExtensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/cUsToMtAsKeXtEnSiOnS/cUsToMtAsKeXtEnSiOnIdVaLuE",
			Expected: &IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{
				CustomTaskExtensionId: "cUsToMtAsKeXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/lIfEcYcLeWoRkFlOwS/cUsToMtAsKeXtEnSiOnS/cUsToMtAsKeXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceLifecycleWorkflowCustomTaskExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CustomTaskExtensionId != v.Expected.CustomTaskExtensionId {
			t.Fatalf("Expected %q but got %q for CustomTaskExtensionId", v.Expected.CustomTaskExtensionId, actual.CustomTaskExtensionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceLifecycleWorkflowCustomTaskExtensionId(t *testing.T) {
	segments := IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
