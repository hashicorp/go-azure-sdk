package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyAppManagementPolicyIdAppliesToId{}

func TestNewPolicyAppManagementPolicyIdAppliesToID(t *testing.T) {
	id := NewPolicyAppManagementPolicyIdAppliesToID("appManagementPolicyId", "directoryObjectId")

	if id.AppManagementPolicyId != "appManagementPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppManagementPolicyId'", id.AppManagementPolicyId, "appManagementPolicyId")
	}

	if id.DirectoryObjectId != "directoryObjectId" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectId")
	}
}

func TestFormatPolicyAppManagementPolicyIdAppliesToID(t *testing.T) {
	actual := NewPolicyAppManagementPolicyIdAppliesToID("appManagementPolicyId", "directoryObjectId").ID()
	expected := "/policies/appManagementPolicies/appManagementPolicyId/appliesTo/directoryObjectId"
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
			Input: "/policies/appManagementPolicies/appManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/appManagementPolicies/appManagementPolicyId/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/appManagementPolicies/appManagementPolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyAppManagementPolicyIdAppliesToId{
				AppManagementPolicyId: "appManagementPolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/appManagementPolicies/appManagementPolicyId/appliesTo/directoryObjectId/extra",
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
			Input: "/policies/appManagementPolicies/appManagementPolicyId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/appManagementPolicies/appManagementPolicyId/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/appManagementPolicies/appManagementPolicyId/appliesTo/directoryObjectId",
			Expected: &PolicyAppManagementPolicyIdAppliesToId{
				AppManagementPolicyId: "appManagementPolicyId",
				DirectoryObjectId:     "directoryObjectId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/appManagementPolicies/appManagementPolicyId/appliesTo/directoryObjectId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId",
			Expected: &PolicyAppManagementPolicyIdAppliesToId{
				AppManagementPolicyId: "aPpMaNaGeMeNtPoLiCyId",
				DirectoryObjectId:     "dIrEcToRyObJeCtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aPpMaNaGeMeNtPoLiCiEs/aPpMaNaGeMeNtPoLiCyId/aPpLiEsTo/dIrEcToRyObJeCtId/extra",
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
