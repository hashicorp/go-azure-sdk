package policystates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &Providers2PolicyStatePolicyStatesResourceId{}

func TestNewProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	id := NewProviders2PolicyStatePolicyStatesResourceID("managementGroupName", "default")

	if id.ManagementGroupName != "managementGroupName" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagementGroupName'", id.ManagementGroupName, "managementGroupName")
	}

	if id.PolicyStatesResource != "default" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyStatesResource'", id.PolicyStatesResource, "default")
	}
}

func TestFormatProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	actual := NewProviders2PolicyStatePolicyStatesResourceID("managementGroupName", "default").ID()
	expected := "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates/default"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2PolicyStatePolicyStatesResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicyStatePolicyStatesResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &Providers2PolicyStatePolicyStatesResourceId{
				ManagementGroupName:  "managementGroupName",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicyStatePolicyStatesResourceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupName != v.Expected.ManagementGroupName {
			t.Fatalf("Expected %q but got %q for ManagementGroupName", v.Expected.ManagementGroupName, actual.ManagementGroupName)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestParseProviders2PolicyStatePolicyStatesResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2PolicyStatePolicyStatesResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates/default",
			Expected: &Providers2PolicyStatePolicyStatesResourceId{
				ManagementGroupName:  "managementGroupName",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Management/managementGroups/managementGroupName/providers/Microsoft.PolicyInsights/policyStates/default/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt",
			Expected: &Providers2PolicyStatePolicyStatesResourceId{
				ManagementGroupName:  "mAnAgEmEnTgRoUpNaMe",
				PolicyStatesResource: "default",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEmEnT/mAnAgEmEnTgRoUpS/mAnAgEmEnTgRoUpNaMe/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYsTaTeS/dEfAuLt/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2PolicyStatePolicyStatesResourceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagementGroupName != v.Expected.ManagementGroupName {
			t.Fatalf("Expected %q but got %q for ManagementGroupName", v.Expected.ManagementGroupName, actual.ManagementGroupName)
		}

		if actual.PolicyStatesResource != v.Expected.PolicyStatesResource {
			t.Fatalf("Expected %q but got %q for PolicyStatesResource", v.Expected.PolicyStatesResource, actual.PolicyStatesResource)
		}

	}
}

func TestSegmentsForProviders2PolicyStatePolicyStatesResourceId(t *testing.T) {
	segments := Providers2PolicyStatePolicyStatesResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2PolicyStatePolicyStatesResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
