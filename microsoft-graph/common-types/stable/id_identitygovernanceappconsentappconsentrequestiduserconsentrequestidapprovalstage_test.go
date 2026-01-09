package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{}

func TestNewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(t *testing.T) {
	id := NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID("appConsentRequestId", "userConsentRequestId", "approvalStageId")

	if id.AppConsentRequestId != "appConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'AppConsentRequestId'", id.AppConsentRequestId, "appConsentRequestId")
	}

	if id.UserConsentRequestId != "userConsentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserConsentRequestId'", id.UserConsentRequestId, "userConsentRequestId")
	}

	if id.ApprovalStageId != "approvalStageId" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStageId'", id.ApprovalStageId, "approvalStageId")
	}
}

func TestFormatIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID(t *testing.T) {
	actual := NewIdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageID("appConsentRequestId", "userConsentRequestId", "approvalStageId").ID()
	expected := "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages/approvalStageId"
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
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages/approvalStageId",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
				ApprovalStageId:      "approvalStageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages/approvalStageId/extra",
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
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL/sTaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages/approvalStageId",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
				AppConsentRequestId:  "appConsentRequestId",
				UserConsentRequestId: "userConsentRequestId",
				ApprovalStageId:      "approvalStageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/appConsent/appConsentRequests/appConsentRequestId/userConsentRequests/userConsentRequestId/approval/stages/approvalStageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL/sTaGeS/aPpRoVaLsTaGeId",
			Expected: &IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestIdApprovalStageId{
				AppConsentRequestId:  "aPpCoNsEnTrEqUeStId",
				UserConsentRequestId: "uSeRcOnSeNtReQuEsTiD",
				ApprovalStageId:      "aPpRoVaLsTaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/aPpCoNsEnT/aPpCoNsEnTrEqUeStS/aPpCoNsEnTrEqUeStId/uSeRcOnSeNtReQuEsTs/uSeRcOnSeNtReQuEsTiD/aPpRoVaL/sTaGeS/aPpRoVaLsTaGeId/extra",
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
