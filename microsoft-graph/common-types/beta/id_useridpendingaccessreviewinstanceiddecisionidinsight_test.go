package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{}

func TestNewUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(t *testing.T) {
	id := NewUserIdPendingAccessReviewInstanceIdDecisionIdInsightID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "governanceInsightId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemId")
	}

	if id.GovernanceInsightId != "governanceInsightId" {
		t.Fatalf("Expected %q but got %q for Segment 'GovernanceInsightId'", id.GovernanceInsightId, "governanceInsightId")
	}
}

func TestFormatUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(t *testing.T) {
	actual := NewUserIdPendingAccessReviewInstanceIdDecisionIdInsightID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "governanceInsightId").ID()
	expected := "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights/governanceInsightId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPendingAccessReviewInstanceIdDecisionIdInsightId
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
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights/governanceInsightId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{
				UserId:                             "userId",
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				GovernanceInsightId:                "governanceInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights/governanceInsightId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(v.Input)
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

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPendingAccessReviewInstanceIdDecisionIdInsightId
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
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights/governanceInsightId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{
				UserId:                             "userId",
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				GovernanceInsightId:                "governanceInsightId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/insights/governanceInsightId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsIgHtS/gOvErNaNcEiNsIgHtId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{
				UserId:                             "uSeRiD",
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiD",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
				GovernanceInsightId:                "gOvErNaNcEiNsIgHtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsIgHtS/gOvErNaNcEiNsIgHtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightIDInsensitively(v.Input)
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

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestSegmentsForUserIdPendingAccessReviewInstanceIdDecisionIdInsightId(t *testing.T) {
	segments := UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPendingAccessReviewInstanceIdDecisionIdInsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
