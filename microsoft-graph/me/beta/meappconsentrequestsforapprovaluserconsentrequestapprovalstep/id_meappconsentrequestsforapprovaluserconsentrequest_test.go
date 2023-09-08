package meappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAppConsentRequestsForApprovalUserConsentRequestId{}

func TestNewMeAppConsentRequestsForApprovalUserConsentRequestID(t *testing.T) {
	id := NewMeAppConsentRequestsForApprovalUserConsentRequestID("appConsentRequestIdValue", "userConsentRequestIdValue")

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}

	if id.UserConsentRequestId != "userConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestIdValue")
	}
}

func TestFormatMeAppConsentRequestsForApprovalUserConsentRequestID(t *testing.T) {
	actual := NewMeAppConsentRequestsForApprovalUserConsentRequestID("appConsentRequestIdValue", "userConsentRequestIdValue").ID()
	expected := "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAppConsentRequestsForApprovalUserConsentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalUserConsentRequestId
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
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Expected: &MeAppConsentRequestsForApprovalUserConsentRequestId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalUserConsentRequestID(v.Input)
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

	}
}

func TestParseMeAppConsentRequestsForApprovalUserConsentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalUserConsentRequestId
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
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Expected: &MeAppConsentRequestsForApprovalUserConsentRequestId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe",
			Expected: &MeAppConsentRequestsForApprovalUserConsentRequestId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStIdVaLuE",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalUserConsentRequestIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeAppConsentRequestsForApprovalUserConsentRequestId(t *testing.T) {
	segments := MeAppConsentRequestsForApprovalUserConsentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAppConsentRequestsForApprovalUserConsentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
