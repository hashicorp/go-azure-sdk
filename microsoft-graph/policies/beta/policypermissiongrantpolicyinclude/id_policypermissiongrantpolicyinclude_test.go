package policypermissiongrantpolicyinclude

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyPermissionGrantPolicyIncludeId{}

func TestNewPolicyPermissionGrantPolicyIncludeID(t *testing.T) {
	id := NewPolicyPermissionGrantPolicyIncludeID("permissionGrantPolicyIdValue", "permissionGrantConditionSetIdValue")

	if id.PermissionGrantPolicyId != "permissionGrantPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantPolicyId'", id.PermissionGrantPolicyId, "permissionGrantPolicyIdValue")
	}

	if id.PermissionGrantConditionSetId != "permissionGrantConditionSetIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PermissionGrantConditionSetId'", id.PermissionGrantConditionSetId, "permissionGrantConditionSetIdValue")
	}
}

func TestFormatPolicyPermissionGrantPolicyIncludeID(t *testing.T) {
	actual := NewPolicyPermissionGrantPolicyIncludeID("permissionGrantPolicyIdValue", "permissionGrantConditionSetIdValue").ID()
	expected := "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes/permissionGrantConditionSetIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyPermissionGrantPolicyIncludeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyIncludeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes/permissionGrantConditionSetIdValue",
			Expected: &PolicyPermissionGrantPolicyIncludeId{
				PermissionGrantPolicyId:       "permissionGrantPolicyIdValue",
				PermissionGrantConditionSetId: "permissionGrantConditionSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes/permissionGrantConditionSetIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIncludeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionGrantPolicyId != v.Expected.PermissionGrantPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPolicyId", v.Expected.PermissionGrantPolicyId, actual.PermissionGrantPolicyId)
		}

		if actual.PermissionGrantConditionSetId != v.Expected.PermissionGrantConditionSetId {
			t.Fatalf("Expected %q but got %q for PermissionGrantConditionSetId", v.Expected.PermissionGrantConditionSetId, actual.PermissionGrantConditionSetId)
		}

	}
}

func TestParsePolicyPermissionGrantPolicyIncludeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyPermissionGrantPolicyIncludeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyIdVaLuE/iNcLuDeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes/permissionGrantConditionSetIdValue",
			Expected: &PolicyPermissionGrantPolicyIncludeId{
				PermissionGrantPolicyId:       "permissionGrantPolicyIdValue",
				PermissionGrantConditionSetId: "permissionGrantConditionSetIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/permissionGrantPolicies/permissionGrantPolicyIdValue/includes/permissionGrantConditionSetIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyIdVaLuE/iNcLuDeS/pErMiSsIoNgRaNtCoNdItIoNsEtIdVaLuE",
			Expected: &PolicyPermissionGrantPolicyIncludeId{
				PermissionGrantPolicyId:       "pErMiSsIoNgRaNtPoLiCyIdVaLuE",
				PermissionGrantConditionSetId: "pErMiSsIoNgRaNtCoNdItIoNsEtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/pErMiSsIoNgRaNtPoLiCiEs/pErMiSsIoNgRaNtPoLiCyIdVaLuE/iNcLuDeS/pErMiSsIoNgRaNtCoNdItIoNsEtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyPermissionGrantPolicyIncludeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PermissionGrantPolicyId != v.Expected.PermissionGrantPolicyId {
			t.Fatalf("Expected %q but got %q for PermissionGrantPolicyId", v.Expected.PermissionGrantPolicyId, actual.PermissionGrantPolicyId)
		}

		if actual.PermissionGrantConditionSetId != v.Expected.PermissionGrantConditionSetId {
			t.Fatalf("Expected %q but got %q for PermissionGrantConditionSetId", v.Expected.PermissionGrantConditionSetId, actual.PermissionGrantConditionSetId)
		}

	}
}

func TestSegmentsForPolicyPermissionGrantPolicyIncludeId(t *testing.T) {
	segments := PolicyPermissionGrantPolicyIncludeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyPermissionGrantPolicyIncludeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
