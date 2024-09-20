package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}

func TestNewMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(t *testing.T) {
	id := NewMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID("appConsentRequestId", "userConsentRequestId", "approvalStepId")

	if id.AppConsentRequestId != "appConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestId")
	}

	if id.UserConsentRequestId != "userConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestId")
	}

	if id.ApprovalStepId != "approvalStepId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepId")
	}
}

func TestFormatMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(t *testing.T) {
	actual := NewMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID("appConsentRequestId", "userConsentRequestId", "approvalStepId").ID()
	expected := "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps/approvalStepId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps/approvalStepId",
			Expected: &MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
				ApprovalStepId:       "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps/approvalStepId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppConsentRequestId != v.Expected.AppConsentRequestId {
			t.Fatalf("Expected %q but got %q for AppConsentRequestId", v.Expected.AppConsentRequestId, actual.AppConsentRequestId)
		}

		if actual.UserConsentRequestId != v.Expected.UserConsentRequestId {
			t.Fatalf("Expected %q but got %q for UserConsentRequestId", v.Expected.UserConsentRequestId, actual.UserConsentRequestId)
		}

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps/approvalStepId",
			Expected: &MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
				ApprovalStepId:       "approvalStepId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/steps/approvalStepId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL/sTePs/aPpRoVaLsTePiD",
			Expected: &MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStId",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiD",
				ApprovalStepId:       "aPpRoVaLsTePiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL/sTePs/aPpRoVaLsTePiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AppConsentRequestId != v.Expected.AppConsentRequestId {
			t.Fatalf("Expected %q but got %q for AppConsentRequestId", v.Expected.AppConsentRequestId, actual.AppConsentRequestId)
		}

		if actual.UserConsentRequestId != v.Expected.UserConsentRequestId {
			t.Fatalf("Expected %q but got %q for UserConsentRequestId", v.Expected.UserConsentRequestId, actual.UserConsentRequestId)
		}

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestSegmentsForMeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId(t *testing.T) {
	segments := MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAppConsentRequestsForApprovalIdUserConsentRequestIdApprovalStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
