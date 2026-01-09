package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAppConsentRequestsForApprovalId{}

func TestNewUserIdAppConsentRequestsForApprovalID(t *testing.T) {
	id := NewUserIdAppConsentRequestsForApprovalID("userId", "appConsentRequestId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AppConsentRequestId != "appConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestId")
	}
}

func TestFormatUserIdAppConsentRequestsForApprovalID(t *testing.T) {
	actual := NewUserIdAppConsentRequestsForApprovalID("userId", "appConsentRequestId").ID()
	expected := "/users/userId/appConsentRequestsForApproval/appConsentRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAppConsentRequestsForApprovalID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAppConsentRequestsForApprovalId
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
			// Valid URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId",
			Expected: &UserIdAppConsentRequestsForApprovalId{
				UserId:              "userId",
				AppConsentRequestId: "appConsentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAppConsentRequestsForApprovalID(v.Input)
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

	}
}

func TestParseUserIdAppConsentRequestsForApprovalIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAppConsentRequestsForApprovalId
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
			// Valid URI
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId",
			Expected: &UserIdAppConsentRequestsForApprovalId{
				UserId:              "userId",
				AppConsentRequestId: "appConsentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/appConsentRequestsForApproval/appConsentRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId",
			Expected: &UserIdAppConsentRequestsForApprovalId{
				UserId:              "uSeRiD",
				AppConsentRequestId: "aPpCoNsEnTrEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAppConsentRequestsForApprovalIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserIdAppConsentRequestsForApprovalId(t *testing.T) {
	segments := UserIdAppConsentRequestsForApprovalId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAppConsentRequestsForApprovalId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
