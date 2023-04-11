package assessmentsmetadata

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AssessmentMetadataId{}

func TestNewAssessmentMetadataID(t *testing.T) {
	id := NewAssessmentMetadataID("assessmentMetadataValue")

	if id.AssessmentMetadataName != "assessmentMetadataValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssessmentMetadataName'", id.AssessmentMetadataName, "assessmentMetadataValue")
	}
}

func TestFormatAssessmentMetadataID(t *testing.T) {
	actual := NewAssessmentMetadataID("assessmentMetadataValue").ID()
	expected := "/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAssessmentMetadataID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AssessmentMetadataId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Security/assessmentMetadata",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue",
			Expected: &AssessmentMetadataId{
				AssessmentMetadataName: "assessmentMetadataValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAssessmentMetadataID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AssessmentMetadataName != v.Expected.AssessmentMetadataName {
			t.Fatalf("Expected %q but got %q for AssessmentMetadataName", v.Expected.AssessmentMetadataName, actual.AssessmentMetadataName)
		}

	}
}

func TestParseAssessmentMetadataIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AssessmentMetadataId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Security/assessmentMetadata",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTmEtAdAtA",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue",
			Expected: &AssessmentMetadataId{
				AssessmentMetadataName: "assessmentMetadataValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTmEtAdAtA/aSsEsSmEnTmEtAdAtAvAlUe",
			Expected: &AssessmentMetadataId{
				AssessmentMetadataName: "aSsEsSmEnTmEtAdAtAvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTmEtAdAtA/aSsEsSmEnTmEtAdAtAvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAssessmentMetadataIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AssessmentMetadataName != v.Expected.AssessmentMetadataName {
			t.Fatalf("Expected %q but got %q for AssessmentMetadataName", v.Expected.AssessmentMetadataName, actual.AssessmentMetadataName)
		}

	}
}

func TestSegmentsForAssessmentMetadataId(t *testing.T) {
	segments := AssessmentMetadataId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AssessmentMetadataId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
