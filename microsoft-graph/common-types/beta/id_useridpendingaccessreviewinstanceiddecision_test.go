package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionId{}

func TestNewUserIdPendingAccessReviewInstanceIdDecisionID(t *testing.T) {
	id := NewUserIdPendingAccessReviewInstanceIdDecisionID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemId")
	}
}

func TestFormatUserIdPendingAccessReviewInstanceIdDecisionID(t *testing.T) {
	actual := NewUserIdPendingAccessReviewInstanceIdDecisionID("userId", "accessReviewInstanceId", "accessReviewInstanceDecisionItemId").ID()
	expected := "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdPendingAccessReviewInstanceIdDecisionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPendingAccessReviewInstanceIdDecisionId
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
			// Valid URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionId{
				UserId:                             "userId",
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPendingAccessReviewInstanceIdDecisionID(v.Input)
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

	}
}

func TestParseUserIdPendingAccessReviewInstanceIdDecisionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdPendingAccessReviewInstanceIdDecisionId
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
			// Valid URI
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionId{
				UserId:                             "userId",
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
			Expected: &UserIdPendingAccessReviewInstanceIdDecisionId{
				UserId:                             "uSeRiD",
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiD",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserIdPendingAccessReviewInstanceIdDecisionId(t *testing.T) {
	segments := UserIdPendingAccessReviewInstanceIdDecisionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdPendingAccessReviewInstanceIdDecisionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
