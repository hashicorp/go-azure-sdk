package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}

func TestNewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(t *testing.T) {
	id := NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewReviewerId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemId")
	}

	if id.AccessReviewReviewerId != "accessReviewReviewerId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewReviewerId'", id.AccessReviewReviewerId, "accessReviewReviewerId")
	}
}

func TestFormatUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(t *testing.T) {
	actual := NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewReviewerId").ID()
	expected := "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
				UserId:                             "userId",
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				AccessReviewReviewerId:             "accessReviewReviewerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
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
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/cOnTaCtEdReViEwErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
				UserId:                             "userId",
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				AccessReviewReviewerId:             "accessReviewReviewerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiD",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
				UserId:                             "uSeRiD",
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiD",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
				AccessReviewReviewerId:             "aCcEsSrEvIeWrEvIeWeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestSegmentsForUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId(t *testing.T) {
	segments := UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
