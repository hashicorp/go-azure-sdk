package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionThreatAssessmentRequestIdResultId{}

func TestNewMeInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	id := NewMeInformationProtectionThreatAssessmentRequestIdResultID("threatAssessmentRequestId", "threatAssessmentResultId")

	if id.ThreatAssessmentRequestId != "threatAssessmentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentRequestId'", id.ThreatAssessmentRequestId, "threatAssessmentRequestId")
	}

	if id.ThreatAssessmentResultId != "threatAssessmentResultId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentResultId'", id.ThreatAssessmentResultId, "threatAssessmentResultId")
	}
}

func TestFormatMeInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	actual := NewMeInformationProtectionThreatAssessmentRequestIdResultID("threatAssessmentRequestId", "threatAssessmentResultId").ID()
	expected := "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionThreatAssessmentRequestIdResultId
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
			Input: "/me/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId",
			Expected: &MeInformationProtectionThreatAssessmentRequestIdResultId{
				ThreatAssessmentRequestId: "threatAssessmentRequestId",
				ThreatAssessmentResultId:  "threatAssessmentResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionThreatAssessmentRequestIdResultID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ThreatAssessmentRequestId != v.Expected.ThreatAssessmentRequestId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentRequestId", v.Expected.ThreatAssessmentRequestId, actual.ThreatAssessmentRequestId)
		}

		if actual.ThreatAssessmentResultId != v.Expected.ThreatAssessmentResultId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentResultId", v.Expected.ThreatAssessmentResultId, actual.ThreatAssessmentResultId)
		}

	}
}

func TestParseMeInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionThreatAssessmentRequestIdResultId
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
			Input: "/me/informationProtection",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/rEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId",
			Expected: &MeInformationProtectionThreatAssessmentRequestIdResultId{
				ThreatAssessmentRequestId: "threatAssessmentRequestId",
				ThreatAssessmentResultId:  "threatAssessmentResultId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/results/threatAssessmentResultId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiD",
			Expected: &MeInformationProtectionThreatAssessmentRequestIdResultId{
				ThreatAssessmentRequestId: "tHrEaTaSsEsSmEnTrEqUeStId",
				ThreatAssessmentResultId:  "tHrEaTaSsEsSmEnTrEsUlTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ThreatAssessmentRequestId != v.Expected.ThreatAssessmentRequestId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentRequestId", v.Expected.ThreatAssessmentRequestId, actual.ThreatAssessmentRequestId)
		}

		if actual.ThreatAssessmentResultId != v.Expected.ThreatAssessmentResultId {
			t.Fatalf("Expected %q but got %q for ThreatAssessmentResultId", v.Expected.ThreatAssessmentResultId, actual.ThreatAssessmentResultId)
		}

	}
}

func TestSegmentsForMeInformationProtectionThreatAssessmentRequestIdResultId(t *testing.T) {
	segments := MeInformationProtectionThreatAssessmentRequestIdResultId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInformationProtectionThreatAssessmentRequestIdResultId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
