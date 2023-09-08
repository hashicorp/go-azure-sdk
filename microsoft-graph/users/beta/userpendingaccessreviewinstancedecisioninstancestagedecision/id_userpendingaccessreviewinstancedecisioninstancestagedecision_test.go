package userpendingaccessreviewinstancedecisioninstancestagedecision

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{}

func TestNewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(t *testing.T) {
	id := NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemIdValue")
	}

	if id.AccessReviewStageId != "accessReviewStageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewStageId'", id.AccessReviewStageId, "accessReviewStageIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId1 != "accessReviewInstanceDecisionItemId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId1'", id.AccessReviewInstanceDecisionItemId1, "accessReviewInstanceDecisionItemId1Value")
	}
}

func TestFormatUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(t *testing.T) {
	actual := NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID("userIdValue", "accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value").ID()
	expected := "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId
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
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value",
			Expected: &UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{
				UserId:                              "userIdValue",
				AccessReviewInstanceId:              "accessReviewInstanceIdValue",
				AccessReviewInstanceDecisionItemId:  "accessReviewInstanceDecisionItemIdValue",
				AccessReviewStageId:                 "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId1: "accessReviewInstanceDecisionItemId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionID(v.Input)
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

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId1 != v.Expected.AccessReviewInstanceDecisionItemId1 {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId1", v.Expected.AccessReviewInstanceDecisionItemId1, actual.AccessReviewInstanceDecisionItemId1)
		}

	}
}

func TestParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId
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
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value",
			Expected: &UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{
				UserId:                              "userIdValue",
				AccessReviewInstanceId:              "accessReviewInstanceIdValue",
				AccessReviewInstanceDecisionItemId:  "accessReviewInstanceDecisionItemIdValue",
				AccessReviewStageId:                 "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId1: "accessReviewInstanceDecisionItemId1Value",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE",
			Expected: &UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{
				UserId:                              "uSeRiDvAlUe",
				AccessReviewInstanceId:              "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewInstanceDecisionItemId:  "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
				AccessReviewStageId:                 "aCcEsSrEvIeWsTaGeIdVaLuE",
				AccessReviewInstanceDecisionItemId1: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserPendingAccessReviewInstanceDecisionInstanceStageDecisionIDInsensitively(v.Input)
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

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId1 != v.Expected.AccessReviewInstanceDecisionItemId1 {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId1", v.Expected.AccessReviewInstanceDecisionItemId1, actual.AccessReviewInstanceDecisionItemId1)
		}

	}
}

func TestSegmentsForUserPendingAccessReviewInstanceDecisionInstanceStageDecisionId(t *testing.T) {
	segments := UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserPendingAccessReviewInstanceDecisionInstanceStageDecisionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
