package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewHistoryDefinitionId{}

func TestNewIdentityGovernanceAccessReviewHistoryDefinitionID(t *testing.T) {
	id := NewIdentityGovernanceAccessReviewHistoryDefinitionID("accessReviewHistoryDefinitionIdValue")

	if id.AccessReviewHistoryDefinitionId != "accessReviewHistoryDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewHistoryDefinitionId'", id.AccessReviewHistoryDefinitionId, "accessReviewHistoryDefinitionIdValue")
	}
}

func TestFormatIdentityGovernanceAccessReviewHistoryDefinitionID(t *testing.T) {
	actual := NewIdentityGovernanceAccessReviewHistoryDefinitionID("accessReviewHistoryDefinitionIdValue").ID()
	expected := "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAccessReviewHistoryDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewHistoryDefinitionId
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
			Input: "/identityGovernance/accessReviews",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/historyDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue",
			Expected: &IdentityGovernanceAccessReviewHistoryDefinitionId{
				AccessReviewHistoryDefinitionId: "accessReviewHistoryDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewHistoryDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessReviewHistoryDefinitionId != v.Expected.AccessReviewHistoryDefinitionId {
			t.Fatalf("Expected %q but got %q for AccessReviewHistoryDefinitionId", v.Expected.AccessReviewHistoryDefinitionId, actual.AccessReviewHistoryDefinitionId)
		}

	}
}

func TestParseIdentityGovernanceAccessReviewHistoryDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewHistoryDefinitionId
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
			Input: "/identityGovernance/accessReviews",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/historyDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue",
			Expected: &IdentityGovernanceAccessReviewHistoryDefinitionId{
				AccessReviewHistoryDefinitionId: "accessReviewHistoryDefinitionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS/aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE",
			Expected: &IdentityGovernanceAccessReviewHistoryDefinitionId{
				AccessReviewHistoryDefinitionId: "aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS/aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewHistoryDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessReviewHistoryDefinitionId != v.Expected.AccessReviewHistoryDefinitionId {
			t.Fatalf("Expected %q but got %q for AccessReviewHistoryDefinitionId", v.Expected.AccessReviewHistoryDefinitionId, actual.AccessReviewHistoryDefinitionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAccessReviewHistoryDefinitionId(t *testing.T) {
	segments := IdentityGovernanceAccessReviewHistoryDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAccessReviewHistoryDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
