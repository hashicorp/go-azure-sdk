package informationprotectionthreatassessmentrequestresult

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionThreatAssessmentRequestIdResultId{}

func TestNewMeInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	id := NewMeInformationProtectionThreatAssessmentRequestIdResultID("threatAssessmentRequestIdValue", "threatAssessmentResultIdValue")

	if id.ThreatAssessmentRequestId != "threatAssessmentRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentRequestId'", id.ThreatAssessmentRequestId, "threatAssessmentRequestIdValue")
	}

	if id.ThreatAssessmentResultId != "threatAssessmentResultIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ThreatAssessmentResultId'", id.ThreatAssessmentResultId, "threatAssessmentResultIdValue")
	}
}

func TestFormatMeInformationProtectionThreatAssessmentRequestIdResultID(t *testing.T) {
	actual := NewMeInformationProtectionThreatAssessmentRequestIdResultID("threatAssessmentRequestIdValue", "threatAssessmentResultIdValue").ID()
	expected := "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue"
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
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue",
			Expected: &MeInformationProtectionThreatAssessmentRequestIdResultId{
				ThreatAssessmentRequestId: "threatAssessmentRequestIdValue",
				ThreatAssessmentResultId:  "threatAssessmentResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue/extra",
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
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/rEsUlTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue",
			Expected: &MeInformationProtectionThreatAssessmentRequestIdResultId{
				ThreatAssessmentRequestId: "threatAssessmentRequestIdValue",
				ThreatAssessmentResultId:  "threatAssessmentResultIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/informationProtection/threatAssessmentRequests/threatAssessmentRequestIdValue/results/threatAssessmentResultIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiDvAlUe",
			Expected: &MeInformationProtectionThreatAssessmentRequestIdResultId{
				ThreatAssessmentRequestId: "tHrEaTaSsEsSmEnTrEqUeStIdVaLuE",
				ThreatAssessmentResultId:  "tHrEaTaSsEsSmEnTrEsUlTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/iNfOrMaTiOnPrOtEcTiOn/tHrEaTaSsEsSmEnTrEqUeStS/tHrEaTaSsEsSmEnTrEqUeStIdVaLuE/rEsUlTs/tHrEaTaSsEsSmEnTrEsUlTiDvAlUe/extra",
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
