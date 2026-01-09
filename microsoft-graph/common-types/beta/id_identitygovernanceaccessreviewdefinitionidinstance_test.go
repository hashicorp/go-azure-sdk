package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceId{}

func TestNewIdentityGovernanceAccessReviewDefinitionIdInstanceID(t *testing.T) {
	id := NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId")

	if id.AccessReviewScheduleDefinitionId != "accessReviewScheduleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewScheduleDefinitionId'", id.AccessReviewScheduleDefinitionId, "accessReviewScheduleDefinitionId")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}
}

func TestFormatIdentityGovernanceAccessReviewDefinitionIdInstanceID(t *testing.T) {
	actual := NewIdentityGovernanceAccessReviewDefinitionIdInstanceID("accessReviewScheduleDefinitionId", "accessReviewInstanceId").ID()
	expected := "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAccessReviewDefinitionIdInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewDefinitionIdInstanceId
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
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceId{
				AccessReviewScheduleDefinitionId: "accessReviewScheduleDefinitionId",
				AccessReviewInstanceId:           "accessReviewInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceID(v.Input)
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

	}
}

func TestParseIdentityGovernanceAccessReviewDefinitionIdInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewDefinitionIdInstanceId
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
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceId{
				AccessReviewScheduleDefinitionId: "accessReviewScheduleDefinitionId",
				AccessReviewInstanceId:           "accessReviewInstanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiD",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceId{
				AccessReviewScheduleDefinitionId: "aCcEsSrEvIeWsChEdUlEdEfInItIoNiD",
				AccessReviewInstanceId:           "aCcEsSrEvIeWiNsTaNcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForIdentityGovernanceAccessReviewDefinitionIdInstanceId(t *testing.T) {
	segments := IdentityGovernanceAccessReviewDefinitionIdInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAccessReviewDefinitionIdInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
