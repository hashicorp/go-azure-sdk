package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{}

func TestNewUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(t *testing.T) {
	id := NewUserIdAppConsentRequestsForApprovalIdUserConsentRequestID("userId", "appConsentRequestId", "userConsentRequestId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AppConsentRequestId != "appConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestId")
	}

	if id.UserConsentRequestId != "userConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestId")
	}
}

func TestFormatUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(t *testing.T) {
	actual := NewUserIdAppConsentRequestsForApprovalIdUserConsentRequestID("userId", "appConsentRequestId", "userConsentRequestId").ID()
	expected := "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAppConsentRequestsForApprovalIdUserConsentRequestId
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
			Input: "/users/userId/appConsentRequestsForApproval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Expected: &UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{
				UserId:               "userId",
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestID(v.Input)
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

		if actual.AppConsentRequestId != v.Expected.AppConsentRequestId {
			t.Fatalf("Expected %q but got %q for AppConsentRequestId", v.Expected.AppConsentRequestId, actual.AppConsentRequestId)
		}

		if actual.UserConsentRequestId != v.Expected.UserConsentRequestId {
			t.Fatalf("Expected %q but got %q for UserConsentRequestId", v.Expected.UserConsentRequestId, actual.UserConsentRequestId)
		}

	}
}

func TestParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAppConsentRequestsForApprovalIdUserConsentRequestId
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
			Input: "/users/userId/appConsentRequestsForApproval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Expected: &UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{
				UserId:               "userId",
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD",
			Expected: &UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{
				UserId:               "uSeRiD",
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStId",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively(v.Input)
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

		if actual.AppConsentRequestId != v.Expected.AppConsentRequestId {
			t.Fatalf("Expected %q but got %q for AppConsentRequestId", v.Expected.AppConsentRequestId, actual.AppConsentRequestId)
		}

		if actual.UserConsentRequestId != v.Expected.UserConsentRequestId {
			t.Fatalf("Expected %q but got %q for UserConsentRequestId", v.Expected.UserConsentRequestId, actual.UserConsentRequestId)
		}

	}
}

func TestSegmentsForUserIdAppConsentRequestsForApprovalIdUserConsentRequestId(t *testing.T) {
	segments := UserIdAppConsentRequestsForApprovalIdUserConsentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAppConsentRequestsForApprovalIdUserConsentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
