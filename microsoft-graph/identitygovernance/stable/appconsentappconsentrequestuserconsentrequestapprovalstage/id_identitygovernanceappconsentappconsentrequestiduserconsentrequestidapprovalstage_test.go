package appconsentappconsentrequestuserconsentrequestapprovalstage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{}

func TestNewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(t *testing.T) {
	id := NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID("appConsentRequestIdValue", "userConsentRequestIdValue", "approvalStageIdValue")

	if id.AppConsentRequestId != "appConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestIdValue")
	}

	if id.UserConsentRequestId != "userConsentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestIdValue")
	}

	if id.ApprovalStageId != "approvalStageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStageId'", id.ApprovalStageId, "approvalStageIdValue")
	}
}

func TestFormatIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(t *testing.T) {
	actual := NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID("appConsentRequestIdValue", "userConsentRequestIdValue", "approvalStageIdValue").ID()
	expected := "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages/approvalStageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages/approvalStageIdValue",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
				ApprovalStageId:      "approvalStageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages/approvalStageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(v.Input)
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

		if actual.ApprovalStageId != v.Expected.ApprovalStageId {
			t.Fatalf("Expected %q but got %q for ApprovalStageId", v.Expected.ApprovalStageId, actual.ApprovalStageId)
		}

	}
}

func TestParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL/sTaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages/approvalStageIdValue",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
				AppConsentRequestId:  "appConsentRequestIdValue",
				UserConsentRequestId: "userConsentRequestIdValue",
				ApprovalStageId:      "approvalStageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestIdValue/userConsentRequests/userConsentRequestIdValue/approval/stages/approvalStageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL/sTaGeS/aPpRoVaLsTaGeIdVaLuE",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStIdVaLuE",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiDvAlUe",
				ApprovalStageId:      "aPpRoVaLsTaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStIdVaLuE/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiDvAlUe/aPpRoVaL/sTaGeS/aPpRoVaLsTaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageIDInsensitively(v.Input)
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

		if actual.ApprovalStageId != v.Expected.ApprovalStageId {
			t.Fatalf("Expected %q but got %q for ApprovalStageId", v.Expected.ApprovalStageId, actual.ApprovalStageId)
		}

	}
}

func TestSegmentsForIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId(t *testing.T) {
	segments := IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
