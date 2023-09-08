package userinformationprotectionthreatassessmentrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserInformationProtectionThreatAssessmentRequestId{}

func TestNewUserInformationProtectionThreatAssessmentRequestID(t *testing.T) {
	id := NewUserInformationProtectionThreatAssessmentRequestID("userIdValue", "threatAssessmentRequestIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ThreatAssessmentRequestId != "threatAssessmentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentRequestId'", id.ThreatAssessmentRequestId, "threatAssessmentRequestIdValue")
	}
}

func TestFormatUserInformationProtectionThreatAssessmentRequestID(t *testing.T) {
	actual := NewUserInformationProtectionThreatAssessmentRequestID("userIdValue", "threatAssessmentRequestIdValue").ID()
	expected := "/users/userIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserInformationProtectionThreatAssessmentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserInformationProtectionThreatAssessmentRequestId
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
			Input: "/users/userIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue",
			Expected: &UserInformationProtectionThreatAssessmentRequestId{
				UserId:                    "userIdValue",
				ThreatAssessmentRequestId: "threatAssessmentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserInformationProtectionThreatAssessmentRequestID(v.Input)
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

	}
}

func TestParseUserInformationProtectionThreatAssessmentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserInformationProtectionThreatAssessmentRequestId
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
			Input: "/users/userIdValue/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue",
			Expected: &UserInformationProtectionThreatAssessmentRequestId{
				UserId:                    "userIdValue",
				ThreatAssessmentRequestId: "threatAssessmentRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE",
			Expected: &UserInformationProtectionThreatAssessmentRequestId{
				UserId:                    "uSeRiDvAlUe",
				ThreatAssessmentRequestId: "tHrEaTaSsEsSmEnTrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserInformationProtectionThreatAssessmentRequestIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserInformationProtectionThreatAssessmentRequestId(t *testing.T) {
	segments := UserInformationProtectionThreatAssessmentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserInformationProtectionThreatAssessmentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
