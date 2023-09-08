package userappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppConsentRequestsForApprovalUserConsentRequestId{}

func TestNewUserAppConsentRequestsForApprovalUserConsentRequestID(t *testing.T) {
	id := NewUserAppConsentRequestsForApprovalUserConsentRequestID("userIdValue", "appConsentRequestIdValue", "userConsentRequestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}

	if id.UserConsentRequestId != "userConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestIdValue")
	}
}

func TestFormatUserAppConsentRequestsForApprovalUserConsentRequestID(t *testing.T) {
	actual := NewUserAppConsentRequestsForApprovalUserConsentRequestID("userIdValue", "appConsentRequestIdValue", "userConsentRequestIdValue").ID()
	expected := "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAppConsentRequestsForApprovalUserConsentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppConsentRequestsForApprovalUserConsentRequestId
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
			Input: "/users/userIdValue/appConsentRequestsForApproval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Expected: &UserAppConsentRequestsForApprovalUserConsentRequestId{
				UserId:               "userIdValue",
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppConsentRequestsForApprovalUserConsentRequestID(v.Input)
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

func TestParseUserAppConsentRequestsForApprovalUserConsentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppConsentRequestsForApprovalUserConsentRequestId
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
			Input: "/users/userIdValue/appConsentRequestsForApproval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Expected: &UserAppConsentRequestsForApprovalUserConsentRequestId{
				UserId:               "userIdValue",
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe",
			Expected: &UserAppConsentRequestsForApprovalUserConsentRequestId{
				UserId:               "uSeRiDvAlUe",
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStIdVaLuE",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppConsentRequestsForApprovalUserConsentRequestIDInsensitively(v.Input)
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

func TestSegmentsForUserAppConsentRequestsForApprovalUserConsentRequestId(t *testing.T) {
	segments := UserAppConsentRequestsForApprovalUserConsentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAppConsentRequestsForApprovalUserConsentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
