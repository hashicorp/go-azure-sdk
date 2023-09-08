package meappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}

func TestNewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(t *testing.T) {
	id := NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID("appConsentRequestIdValue", "userConsentRequestIdValue", "approvalStepIdValue")

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}

	if id.UserConsentRequestId != "userConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestIdValue")
	}

	if id.ApprovalStepId != "approvalStepIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepIdValue")
	}
}

func TestFormatMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(t *testing.T) {
	actual := NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID("appConsentRequestIdValue", "userConsentRequestIdValue", "approvalStepIdValue").ID()
	expected := "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps/approvalStepIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId
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
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps/approvalStepIdValue",
			Expected: &MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
				ApprovalStepId:       "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps/approvalStepIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepID(v.Input)
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

func TestParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId
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
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps/approvalStepIdValue",
			Expected: &MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
				ApprovalStepId:       "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/steps/approvalStepIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL/sTePs/aPpRoVaLsTePiDvAlUe",
			Expected: &MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStIdVaLuE",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiDvAlUe",
				ApprovalStepId:       "aPpRoVaLsTePiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL/sTePs/aPpRoVaLsTePiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepIDInsensitively(v.Input)
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

func TestSegmentsForMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId(t *testing.T) {
	segments := MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
