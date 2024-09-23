package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionThreatAssessmentRequestId{}

func TestNewMeInformationProtectionThreatAssessmentRequestID(t *testing.T) {
	id := NewMeInformationProtectionThreatAssessmentRequestID("threatAssessmentRequestId")

	if id.ThreatAssessmentRequestId != "threatAssessmentRequestId" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentRequestId'", id.ThreatAssessmentRequestId, "threatAssessmentRequestId")
	}
}

func TestFormatMeInformationProtectionThreatAssessmentRequestID(t *testing.T) {
	actual := NewMeInformationProtectionThreatAssessmentRequestID("threatAssessmentRequestId").ID()
	expected := "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeInformationProtectionThreatAssessmentRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionThreatAssessmentRequestId
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
			// Valid URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId",
			Expected: &MeInformationProtectionThreatAssessmentRequestId{
				ThreatAssessmentRequestId: "threatAssessmentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionThreatAssessmentRequestID(v.Input)
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

	}
}

func TestParseMeInformationProtectionThreatAssessmentRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeInformationProtectionThreatAssessmentRequestId
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
			// Valid URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId",
			Expected: &MeInformationProtectionThreatAssessmentRequestId{
				ThreatAssessmentRequestId: "threatAssessmentRequestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId",
			Expected: &MeInformationProtectionThreatAssessmentRequestId{
				ThreatAssessmentRequestId: "tHrEaTaSsEsSmEnTrEqUeStId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeInformationProtectionThreatAssessmentRequestIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeInformationProtectionThreatAssessmentRequestId(t *testing.T) {
	segments := MeInformationProtectionThreatAssessmentRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeInformationProtectionThreatAssessmentRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
