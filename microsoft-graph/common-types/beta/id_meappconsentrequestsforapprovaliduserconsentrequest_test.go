package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAppConsentRequestsForApprovalIdUserConsentRequestId{}

func TestNewMeAppConsentRequestsForApprovalIdUserConsentRequestID(t *testing.T) {
	id := NewMeAppConsentRequestsForApprovalIdUserConsentRequestID("appConsentRequestId", "userConsentRequestId")

	if id.AppConsentRequestId != "appConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestId")
	}

	if id.UserConsentRequestId != "userConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestId")
	}
}

func TestFormatMeAppConsentRequestsForApprovalIdUserConsentRequestID(t *testing.T) {
	actual := NewMeAppConsentRequestsForApprovalIdUserConsentRequestID("appConsentRequestId", "userConsentRequestId").ID()
	expected := "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAppConsentRequestsForApprovalIdUserConsentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalIdUserConsentRequestId
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
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Expected: &MeAppConsentRequestsForApprovalIdUserConsentRequestId{
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalIdUserConsentRequestID(v.Input)
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

func TestParseMeAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAppConsentRequestsForApprovalIdUserConsentRequestId
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
			// Valid URI
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Expected: &MeAppConsentRequestsForApprovalIdUserConsentRequestId{
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/appConsentRequestsForApproval/appConsentRequestId/userConsentRequests/userConsentRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD",
			Expected: &MeAppConsentRequestsForApprovalIdUserConsentRequestId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStId",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aPpCoNsEnTrEqUeStSfOrApPrOvAl/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAppConsentRequestsForApprovalIdUserConsentRequestIDInsensitively(v.Input)
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

func TestSegmentsForMeAppConsentRequestsForApprovalIdUserConsentRequestId(t *testing.T) {
	segments := MeAppConsentRequestsForApprovalIdUserConsentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAppConsentRequestsForApprovalIdUserConsentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
