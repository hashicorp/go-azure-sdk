package accessreviewdefinitioninstancestagedecision

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{}

func TestNewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(t *testing.T) {
	id := NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue")

	if id.AccessReviewScheduleDefinitionId != "accessReviewScheduleDefinitionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewScheduleDefinitionId'", id.AccessReviewScheduleDefinitionId, "accessReviewScheduleDefinitionIdValue")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewStageId != "accessReviewStageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewStageId'", id.AccessReviewStageId, "accessReviewStageIdValue")
	}
}

func TestFormatIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(t *testing.T) {
	actual := NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID("accessReviewScheduleDefinitionIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue").ID()
	expected := "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId
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
			Input: "/identityGovernance/accessReviews/definitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{
				AccessReviewScheduleDefinitionId: "accessReviewScheduleDefinitionIdValue",
				AccessReviewInstanceId:           "accessReviewInstanceIdValue",
				AccessReviewStageId:              "accessReviewStageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessReviewScheduleDefinitionId != v.Expected.AccessReviewScheduleDefinitionId {
			t.Fatalf("Expected %q but got %q for AccessReviewScheduleDefinitionId", v.Expected.AccessReviewScheduleDefinitionId, actual.AccessReviewScheduleDefinitionId)
		}

		if actual.AccessReviewInstanceId != v.Expected.AccessReviewInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceId", v.Expected.AccessReviewInstanceId, actual.AccessReviewInstanceId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

	}
}

func TestParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId
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
			Input: "/identityGovernance/accessReviews/definitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe/iNsTaNcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{
				AccessReviewScheduleDefinitionId: "accessReviewScheduleDefinitionIdValue",
				AccessReviewInstanceId:           "accessReviewInstanceIdValue",
				AccessReviewStageId:              "accessReviewStageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionIdValue/instances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{
				AccessReviewScheduleDefinitionId: "aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe",
				AccessReviewInstanceId:           "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewStageId:              "aCcEsSrEvIeWsTaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiDvAlUe/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessReviewScheduleDefinitionId != v.Expected.AccessReviewScheduleDefinitionId {
			t.Fatalf("Expected %q but got %q for AccessReviewScheduleDefinitionId", v.Expected.AccessReviewScheduleDefinitionId, actual.AccessReviewScheduleDefinitionId)
		}

		if actual.AccessReviewInstanceId != v.Expected.AccessReviewInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceId", v.Expected.AccessReviewInstanceId, actual.AccessReviewInstanceId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId(t *testing.T) {
	segments := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
