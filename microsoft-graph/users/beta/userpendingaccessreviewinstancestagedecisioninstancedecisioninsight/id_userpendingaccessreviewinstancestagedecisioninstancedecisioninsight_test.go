package userpendingaccessreviewinstancestagedecisioninstancedecisioninsight

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{}

func TestNewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(t *testing.T) {
	id := NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value", "governanceInsightIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewStageId != "accessReviewStageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewStageId'", id.AccessReviewStageId, "accessReviewStageIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId1 != "accessReviewInstanceDecisionItemId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId1'", id.AccessReviewInstanceDecisionItemId1, "accessReviewInstanceDecisionItemId1Value")
	}

	if id.GovernanceInsightId != "governanceInsightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GovernanceInsightId'", id.GovernanceInsightId, "governanceInsightIdValue")
	}
}

func TestFormatUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(t *testing.T) {
	actual := NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID("userIdValue", "accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewInstanceDecisionItemId1Value", "governanceInsightIdValue").ID()
	expected := "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue",
			Expected: &UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{
				UserId:                              "userIdValue",
				AccessReviewInstanceId:              "accessReviewInstanceIdValue",
				AccessReviewStageId:                 "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId:  "accessReviewInstanceDecisionItemIdValue",
				AccessReviewInstanceDecisionItemId1: "accessReviewInstanceDecisionItemId1Value",
				GovernanceInsightId:                 "governanceInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.AccessReviewInstanceId != v.Expected.AccessReviewInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceId", v.Expected.AccessReviewInstanceId, actual.AccessReviewInstanceId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.AccessReviewInstanceDecisionItemId1 != v.Expected.AccessReviewInstanceDecisionItemId1 {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId1", v.Expected.AccessReviewInstanceDecisionItemId1, actual.AccessReviewInstanceDecisionItemId1)
		}

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/iNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue",
			Expected: &UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{
				UserId:                              "userIdValue",
				AccessReviewInstanceId:              "accessReviewInstanceIdValue",
				AccessReviewStageId:                 "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId:  "accessReviewInstanceDecisionItemIdValue",
				AccessReviewInstanceDecisionItemId1: "accessReviewInstanceDecisionItemId1Value",
				GovernanceInsightId:                 "governanceInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/iNsIgHtS/gOvErNaNcEiNsIgHtIdVaLuE",
			Expected: &UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{
				UserId:                              "uSeRiDvAlUe",
				AccessReviewInstanceId:              "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewStageId:                 "aCcEsSrEvIeWsTaGeIdVaLuE",
				AccessReviewInstanceDecisionItemId:  "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
				AccessReviewInstanceDecisionItemId1: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE",
				GovernanceInsightId:                 "gOvErNaNcEiNsIgHtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/iNsIgHtS/gOvErNaNcEiNsIgHtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.AccessReviewInstanceId != v.Expected.AccessReviewInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceId", v.Expected.AccessReviewInstanceId, actual.AccessReviewInstanceId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.AccessReviewInstanceDecisionItemId1 != v.Expected.AccessReviewInstanceDecisionItemId1 {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId1", v.Expected.AccessReviewInstanceDecisionItemId1, actual.AccessReviewInstanceDecisionItemId1)
		}

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestSegmentsForUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId(t *testing.T) {
	segments := UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
