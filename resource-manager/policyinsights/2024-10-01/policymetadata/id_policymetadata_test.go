package policymetadata

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicyMetadataId{}

func TestNewPolicyMetadataID(t *testing.T) {
	id := NewPolicyMetadataID("policyMetadataName")

	if id.PolicyMetadataName != "policyMetadataName" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyMetadataName'", id.PolicyMetadataName, "policyMetadataName")
	}
}

func TestFormatPolicyMetadataID(t *testing.T) {
	actual := NewPolicyMetadataID("policyMetadataName").ID()
	expected := "/providers/Microsoft.PolicyInsights/policyMetadata/policyMetadataName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyMetadataID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMetadataId
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
			Input: "/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.PolicyInsights/policyMetadata",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.PolicyInsights/policyMetadata/policyMetadataName",
			Expected: &PolicyMetadataId{
				PolicyMetadataName: "policyMetadataName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.PolicyInsights/policyMetadata/policyMetadataName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMetadataID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicyMetadataName != v.Expected.PolicyMetadataName {
			t.Fatalf("Expected %q but got %q for PolicyMetadataName", v.Expected.PolicyMetadataName, actual.PolicyMetadataName)
		}

	}
}

func TestParsePolicyMetadataIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyMetadataId
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
			Input: "/providers/Microsoft.PolicyInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.PolicyInsights/policyMetadata",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYmEtAdAtA",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.PolicyInsights/policyMetadata/policyMetadataName",
			Expected: &PolicyMetadataId{
				PolicyMetadataName: "policyMetadataName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.PolicyInsights/policyMetadata/policyMetadataName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYmEtAdAtA/pOlIcYmEtAdAtAnAmE",
			Expected: &PolicyMetadataId{
				PolicyMetadataName: "pOlIcYmEtAdAtAnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.pOlIcYiNsIgHtS/pOlIcYmEtAdAtA/pOlIcYmEtAdAtAnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyMetadataIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicyMetadataName != v.Expected.PolicyMetadataName {
			t.Fatalf("Expected %q but got %q for PolicyMetadataName", v.Expected.PolicyMetadataName, actual.PolicyMetadataName)
		}

	}
}

func TestSegmentsForPolicyMetadataId(t *testing.T) {
	segments := PolicyMetadataId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyMetadataId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
