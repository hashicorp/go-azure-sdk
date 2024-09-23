package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionThreatAssessmentRequestIdResultId{}

func TestNewUserIdInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	id := NewUserIdInformationProtectionThreatAssessmentRequestIdResultID("userId", "threatAssessmentRequestId", "threatAssessmentResultId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ThreatAssessmentRequestId != "threatAssessmentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentRequestId'", id.ThreatAssessmentRequestId, "threatAssessmentRequestId")
	}

	if id.ThreatAssessmentResultId != "threatAssessmentResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentResultId'", id.ThreatAssessmentResultId, "threatAssessmentResultId")
	}
}

func TestFormatUserIdInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	actual := NewUserIdInformationProtectionThreatAssessmentRequestIdResultID("userId", "threatAssessmentRequestId", "threatAssessmentResultId").ID()
	expected := "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionThreatAssessmentRequestIdResultId
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
			Input: "/users/userId/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId",
			Expected: &UserIdInformationProtectionThreatAssessmentRequestIdResultId{
				UserId:                    "userId",
				ThreatAssessmentRequestId: "threatAssessmentRequestId",
				ThreatAssessmentResultId:  "threatAssessmentResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionThreatAssessmentRequestIdResultID(v.Input)
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

		if actual.ThreatAssessmentRequestId != v.Expected.ThreatAssessmentRequestId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentRequestId", v.Expected.ThreatAssessmentRequestId, actual.ThreatAssessmentRequestId)
		}

		if actual.ThreatAssessmentResultId != v.Expected.ThreatAssessmentResultId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentResultId", v.Expected.ThreatAssessmentResultId, actual.ThreatAssessmentResultId)
		}

	}
}

func TestParseUserIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdInformationProtectionThreatAssessmentRequestIdResultId
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
			Input: "/users/userId/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/rEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId",
			Expected: &UserIdInformationProtectionThreatAssessmentRequestIdResultId{
				UserId:                    "userId",
				ThreatAssessmentRequestId: "threatAssessmentRequestId",
				ThreatAssessmentResultId:  "threatAssessmentResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiD",
			Expected: &UserIdInformationProtectionThreatAssessmentRequestIdResultId{
				UserId:                    "uSeRiD",
				ThreatAssessmentRequestId: "tHrEaTaSsEsSmEnTrEqUeStId",
				ThreatAssessmentResultId:  "tHrEaTaSsEsSmEnTrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(v.Input)
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

		if actual.ThreatAssessmentRequestId != v.Expected.ThreatAssessmentRequestId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentRequestId", v.Expected.ThreatAssessmentRequestId, actual.ThreatAssessmentRequestId)
		}

		if actual.ThreatAssessmentResultId != v.Expected.ThreatAssessmentResultId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentResultId", v.Expected.ThreatAssessmentResultId, actual.ThreatAssessmentResultId)
		}

	}
}

func TestSegmentsForUserIdInformationProtectionThreatAssessmentRequestIdResultId(t *testing.T) {
	segments := UserIdInformationProtectionThreatAssessmentRequestIdResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdInformationProtectionThreatAssessmentRequestIdResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
