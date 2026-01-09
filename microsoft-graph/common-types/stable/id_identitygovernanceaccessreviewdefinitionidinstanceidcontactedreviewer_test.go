package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{}

func TestNewIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(t *testing.T) {
	id := NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewReviewerId")

	if id.AccessReviewScheduleDefinitionId != "accessReviewScheduleDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewScheduleDefinitionId'", id.AccessReviewScheduleDefinitionId, "accessReviewScheduleDefinitionId")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}

	if id.AccessReviewReviewerId != "accessReviewReviewerId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewReviewerId'", id.AccessReviewReviewerId, "accessReviewReviewerId")
	}
}

func TestFormatIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(t *testing.T) {
	actual := NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID("accessReviewScheduleDefinitionId", "accessReviewInstanceId", "accessReviewReviewerId").ID()
	expected := "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers/accessReviewReviewerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId
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
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers/accessReviewReviewerId",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{
				AccessReviewScheduleDefinitionId: "accessReviewScheduleDefinitionId",
				AccessReviewInstanceId:           "accessReviewInstanceId",
				AccessReviewReviewerId:           "accessReviewReviewerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers/accessReviewReviewerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId
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
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiD/cOnTaCtEdReViEwErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers/accessReviewReviewerId",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{
				AccessReviewScheduleDefinitionId: "accessReviewScheduleDefinitionId",
				AccessReviewInstanceId:           "accessReviewInstanceId",
				AccessReviewReviewerId:           "accessReviewReviewerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/accessReviews/definitions/accessReviewScheduleDefinitionId/instances/accessReviewInstanceId/contactedReviewers/accessReviewReviewerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiD/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiD",
			Expected: &IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{
				AccessReviewScheduleDefinitionId: "aCcEsSrEvIeWsChEdUlEdEfInItIoNiD",
				AccessReviewInstanceId:           "aCcEsSrEvIeWiNsTaNcEiD",
				AccessReviewReviewerId:           "aCcEsSrEvIeWrEvIeWeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aCcEsSrEvIeWs/dEfInItIoNs/aCcEsSrEvIeWsChEdUlEdEfInItIoNiD/iNsTaNcEs/aCcEsSrEvIeWiNsTaNcEiD/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerIDInsensitively(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId(t *testing.T) {
	segments := IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
