package accessreviewhistorydefinitioninstance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{}

func TestNewIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(t *testing.T) {
	id := NewIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID("accessReviewHistoryDefinitionIdValue", "accessReviewHistoryInstanceIdValue")

	if id.AccessReviewHistoryDefinitionId != "accessReviewHistoryDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewHistoryDefinitionId'", id.AccessReviewHistoryDefinitionId, "accessReviewHistoryDefinitionIdValue")
	}

	if id.AccessReviewHistoryInstanceId != "accessReviewHistoryInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewHistoryInstanceId'", id.AccessReviewHistoryInstanceId, "accessReviewHistoryInstanceIdValue")
	}
}

func TestFormatIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(t *testing.T) {
	actual := NewIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID("accessReviewHistoryDefinitionIdValue", "accessReviewHistoryInstanceIdValue").ID()
	expected := "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances/accessReviewHistoryInstanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId
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
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances/accessReviewHistoryInstanceIdValue",
			Expected: &IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{
				AccessReviewHistoryDefinitionId: "accessReviewHistoryDefinitionIdValue",
				AccessReviewHistoryInstanceId:   "accessReviewHistoryInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances/accessReviewHistoryInstanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(v.Input)
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

		if actual.AccessReviewHistoryInstanceId != v.Expected.AccessReviewHistoryInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewHistoryInstanceId", v.Expected.AccessReviewHistoryInstanceId, actual.AccessReviewHistoryInstanceId)
		}

	}
}

func TestParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId
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
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS/aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS/aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances/accessReviewHistoryInstanceIdValue",
			Expected: &IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{
				AccessReviewHistoryDefinitionId: "accessReviewHistoryDefinitionIdValue",
				AccessReviewHistoryInstanceId:   "accessReviewHistoryInstanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/historyDefinitions/accessReviewHistoryDefinitionIdValue/instances/accessReviewHistoryInstanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS/aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE/iNsTaNcEs/aCcEsSrEvIeWhIsToRyInStAnCeIdVaLuE",
			Expected: &IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{
				AccessReviewHistoryDefinitionId: "aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE",
				AccessReviewHistoryInstanceId:   "aCcEsSrEvIeWhIsToRyInStAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/hIsToRyDeFiNiTiOnS/aCcEsSrEvIeWhIsToRyDeFiNiTiOnIdVaLuE/iNsTaNcEs/aCcEsSrEvIeWhIsToRyInStAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceIDInsensitively(v.Input)
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

		if actual.AccessReviewHistoryInstanceId != v.Expected.AccessReviewHistoryInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewHistoryInstanceId", v.Expected.AccessReviewHistoryInstanceId, actual.AccessReviewHistoryInstanceId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId(t *testing.T) {
	segments := IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
