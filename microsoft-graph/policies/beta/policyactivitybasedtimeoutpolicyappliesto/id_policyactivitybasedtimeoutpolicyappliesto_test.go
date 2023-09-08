package policyactivitybasedtimeoutpolicyappliesto

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyActivityBasedTimeoutPolicyAppliesToId{}

func TestNewPolicyActivityBasedTimeoutPolicyAppliesToID(t *testing.T) {
	id := NewPolicyActivityBasedTimeoutPolicyAppliesToID("activityBasedTimeoutPolicyIdValue", "directoryObjectIdValue")

	if id.ActivityBasedTimeoutPolicyId != "activityBasedTimeoutPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ActivityBasedTimeoutPolicyId'", id.ActivityBasedTimeoutPolicyId, "activityBasedTimeoutPolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatPolicyActivityBasedTimeoutPolicyAppliesToID(t *testing.T) {
	actual := NewPolicyActivityBasedTimeoutPolicyAppliesToID("activityBasedTimeoutPolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyActivityBasedTimeoutPolicyAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyActivityBasedTimeoutPolicyAppliesToId
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
			Input: "/policies/activityBasedTimeoutPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyActivityBasedTimeoutPolicyAppliesToId{
				ActivityBasedTimeoutPolicyId: "activityBasedTimeoutPolicyIdValue",
				DirectoryObjectId:            "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyActivityBasedTimeoutPolicyAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActivityBasedTimeoutPolicyId != v.Expected.ActivityBasedTimeoutPolicyId {
			t.Fatalf("Expected %q but got %q for ActivityBasedTimeoutPolicyId", v.Expected.ActivityBasedTimeoutPolicyId, actual.ActivityBasedTimeoutPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParsePolicyActivityBasedTimeoutPolicyAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyActivityBasedTimeoutPolicyAppliesToId
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
			Input: "/policies/activityBasedTimeoutPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &PolicyActivityBasedTimeoutPolicyAppliesToId{
				ActivityBasedTimeoutPolicyId: "activityBasedTimeoutPolicyIdValue",
				DirectoryObjectId:            "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policies/activityBasedTimeoutPolicies/activityBasedTimeoutPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &PolicyActivityBasedTimeoutPolicyAppliesToId{
				ActivityBasedTimeoutPolicyId: "aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe",
				DirectoryObjectId:            "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcIeS/aCtIvItYbAsEdTiMeOuTpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyActivityBasedTimeoutPolicyAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ActivityBasedTimeoutPolicyId != v.Expected.ActivityBasedTimeoutPolicyId {
			t.Fatalf("Expected %q but got %q for ActivityBasedTimeoutPolicyId", v.Expected.ActivityBasedTimeoutPolicyId, actual.ActivityBasedTimeoutPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForPolicyActivityBasedTimeoutPolicyAppliesToId(t *testing.T) {
	segments := PolicyActivityBasedTimeoutPolicyAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyActivityBasedTimeoutPolicyAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
