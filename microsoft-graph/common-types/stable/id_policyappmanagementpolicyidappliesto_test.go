package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAppManagementPolicyIdAppliesToId{}

func TestNewPolicyAppManagementPolicyIdAppliesToID(t *testing.T) {
	id := NewPolicyAppManagementPolicyIdAppliesToID("appManagementPolicyIdValue", "directoryObjectIdValue")

	if id.AppManagementPolicyId != "appManagementPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppManagementPolicyId'", id.AppManagementPolicyId, "appManagementPolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatPolicyAppManagementPolicyIdAppliesToID(t *testing.T) {
	actual := NewPolicyAppManagementPolicyIdAppliesToID("appManagementPolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAppManagementPolicyIdAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAppManagementPolicyIdAppliesToId
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
			Input: "/policies/appManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyAppManagementPolicyIdAppliesToId{
				AppManagementPolicyId: "appManagementPolicyIdValue",
				DirectoryObjectId:     "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAppManagementPolicyIdAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppManagementPolicyId != v.Expected.AppManagementPolicyId {
			t.Fatalf("Expected %q but got %q for AppManagementPolicyId", v.Expected.AppManagementPolicyId, actual.AppManagementPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyAppManagementPolicyIdAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAppManagementPolicyIdAppliesToId
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
			Input: "/policies/appManagementPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyIdVaLuE/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyAppManagementPolicyIdAppliesToId{
				AppManagementPolicyId: "appManagementPolicyIdValue",
				DirectoryObjectId:     "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/appManagementPolicies/appManagementPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyIdVaLuE/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &PolicyAppManagementPolicyIdAppliesToId{
				AppManagementPolicyId: "aPpMaNaGeMeNtPoLiCyIdVaLuE",
				DirectoryObjectId:     "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyIdVaLuE/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAppManagementPolicyIdAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppManagementPolicyId != v.Expected.AppManagementPolicyId {
			t.Fatalf("Expected %q but got %q for AppManagementPolicyId", v.Expected.AppManagementPolicyId, actual.AppManagementPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyAppManagementPolicyIdAppliesToId(t *testing.T) {
	segments := PolicyAppManagementPolicyIdAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAppManagementPolicyIdAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
