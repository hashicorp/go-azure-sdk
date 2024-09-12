package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryFeatureRolloutPolicyIdAppliesToId{}

func TestNewDirectoryFeatureRolloutPolicyIdAppliesToID(t *testing.T) {
	id := NewDirectoryFeatureRolloutPolicyIdAppliesToID("featureRolloutPolicyIdValue", "directoryObjectIdValue")

	if id.FeatureRolloutPolicyId != "featureRolloutPolicyIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'FeatureRolloutPolicyId'", id.FeatureRolloutPolicyId, "featureRolloutPolicyIdValue")
	}

	if id.DirectoryObjectId != "directoryObjectIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectoryObjectId'", id.DirectoryObjectId, "directoryObjectIdValue")
	}
}

func TestFormatDirectoryFeatureRolloutPolicyIdAppliesToID(t *testing.T) {
	actual := NewDirectoryFeatureRolloutPolicyIdAppliesToID("featureRolloutPolicyIdValue", "directoryObjectIdValue").ID()
	expected := "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryFeatureRolloutPolicyIdAppliesToID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryFeatureRolloutPolicyIdAppliesToId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/featureRolloutPolicies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &DirectoryFeatureRolloutPolicyIdAppliesToId{
				FeatureRolloutPolicyId: "featureRolloutPolicyIdValue",
				DirectoryObjectId:      "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryFeatureRolloutPolicyIdAppliesToID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.FeatureRolloutPolicyId != v.Expected.FeatureRolloutPolicyId {
			t.Fatalf("Expected %q but got %q for FeatureRolloutPolicyId", v.Expected.FeatureRolloutPolicyId, actual.FeatureRolloutPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestParseDirectoryFeatureRolloutPolicyIdAppliesToIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryFeatureRolloutPolicyIdAppliesToId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/featureRolloutPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEaTuReRoLlOuTpOlIcIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/aPpLiEsTo",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue",
			Expected: &DirectoryFeatureRolloutPolicyIdAppliesToId{
				FeatureRolloutPolicyId: "featureRolloutPolicyIdValue",
				DirectoryObjectId:      "directoryObjectIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/featureRolloutPolicies/featureRolloutPolicyIdValue/appliesTo/directoryObjectIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE",
			Expected: &DirectoryFeatureRolloutPolicyIdAppliesToId{
				FeatureRolloutPolicyId: "fEaTuReRoLlOuTpOlIcYiDvAlUe",
				DirectoryObjectId:      "dIrEcToRyObJeCtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/fEaTuReRoLlOuTpOlIcIeS/fEaTuReRoLlOuTpOlIcYiDvAlUe/aPpLiEsTo/dIrEcToRyObJeCtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryFeatureRolloutPolicyIdAppliesToIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.FeatureRolloutPolicyId != v.Expected.FeatureRolloutPolicyId {
			t.Fatalf("Expected %q but got %q for FeatureRolloutPolicyId", v.Expected.FeatureRolloutPolicyId, actual.FeatureRolloutPolicyId)
		}

		if actual.DirectoryObjectId != v.Expected.DirectoryObjectId {
			t.Fatalf("Expected %q but got %q for DirectoryObjectId", v.Expected.DirectoryObjectId, actual.DirectoryObjectId)
		}

	}
}

func TestSegmentsForDirectoryFeatureRolloutPolicyIdAppliesToId(t *testing.T) {
	segments := DirectoryFeatureRolloutPolicyIdAppliesToId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryFeatureRolloutPolicyIdAppliesToId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
