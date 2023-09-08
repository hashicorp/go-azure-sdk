package policyhomerealmdiscoverypolicyappliesto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyHomeRealmDiscoveryPolicyAppliesToId{}

func TestNewPolicyHomeRealmDiscoveryPolicyAppliesToID(t *testing.T) {
	id := NewPolicyHomeRealmDiscoveryPolicyAppliesToID("homeRealmDiscoveryPolicyIdValue", "directoryObjectIdValue")

	if id.HomeRealmDiscoveryPolicyId != "homeRealmDiscoveryPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HomeRealmDiscoveryPolicyId'", id.HomeRealmDiscoveryPolicyId, "homeRealmDiscoveryPolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatPolicyHomeRealmDiscoveryPolicyAppliesToID(t *testing.T) {
	actual := NewPolicyHomeRealmDiscoveryPolicyAppliesToID("homeRealmDiscoveryPolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyHomeRealmDiscoveryPolicyAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyHomeRealmDiscoveryPolicyAppliesToId
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
			Input: "/policies/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyHomeRealmDiscoveryPolicyAppliesToId{
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyIdValue",
				DirectoryObjectId:          "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyHomeRealmDiscoveryPolicyAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyHomeRealmDiscoveryPolicyAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyHomeRealmDiscoveryPolicyAppliesToId
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
			Input: "/policies/homeRealmDiscoveryPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyHomeRealmDiscoveryPolicyAppliesToId{
				HomeRealmDiscoveryPolicyId: "homeRealmDiscoveryPolicyIdValue",
				DirectoryObjectId:          "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/homeRealmDiscoveryPolicies/homeRealmDiscoveryPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &PolicyHomeRealmDiscoveryPolicyAppliesToId{
				HomeRealmDiscoveryPolicyId: "hOmErEaLmDiScOvErYpOlIcYiDvAlUe",
				DirectoryObjectId:          "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/hOmErEaLmDiScOvErYpOlIcIeS/hOmErEaLmDiScOvErYpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyHomeRealmDiscoveryPolicyAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.HomeRealmDiscoveryPolicyId != v.Expected.HomeRealmDiscoveryPolicyId {
			t.Fatalf("Expected %q but got %q for HomeRealmDiscoveryPolicyId", v.Expected.HomeRealmDiscoveryPolicyId, actual.HomeRealmDiscoveryPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyHomeRealmDiscoveryPolicyAppliesToId(t *testing.T) {
	segments := PolicyHomeRealmDiscoveryPolicyAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyHomeRealmDiscoveryPolicyAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
