package userappconsentrequestsforapproval

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppConsentRequestsForApprovalId{}

func TestNewUserAppConsentRequestsForApprovalID(t *testing.T) {
	id := NewUserAppConsentRequestsForApprovalID("userIdValue", "appConsentRequestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}
}

func TestFormatUserAppConsentRequestsForApprovalID(t *testing.T) {
	actual := NewUserAppConsentRequestsForApprovalID("userIdValue", "appConsentRequestIdValue").ID()
	expected := "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAppConsentRequestsForApprovalID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppConsentRequestsForApprovalId
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
			// Valid URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue",
			Expected: &UserAppConsentRequestsForApprovalId{
				UserId:              "userIdValue",
				AppConsentRequestId: "appConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppConsentRequestsForApprovalID(v.Input)
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

func TestParseUserAppConsentRequestsForApprovalIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppConsentRequestsForApprovalId
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
			// Valid URI
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue",
			Expected: &UserAppConsentRequestsForApprovalId{
				UserId:              "userIdValue",
				AppConsentRequestId: "appConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appConsentRequestsForApproval/appConsentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE",
			Expected: &UserAppConsentRequestsForApprovalId{
				UserId:              "uSeRiDvAlUe",
				AppConsentRequestId: "aPpCoNsEnTrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppConsentRequestsForApprovalIDInsensitively(v.Input)
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

func TestSegmentsForUserAppConsentRequestsForApprovalId(t *testing.T) {
	segments := UserAppConsentRequestsForApprovalId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAppConsentRequestsForApprovalId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
