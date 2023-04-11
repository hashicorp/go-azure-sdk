package subassessments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedSubAssessmentId{}

func TestNewScopedSubAssessmentID(t *testing.T) {
	id := NewScopedSubAssessmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assessmentValue", "subAssessmentValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.AssessmentName != "assessmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssessmentName'", id.AssessmentName, "assessmentValue")
	}

	if id.SubAssessmentName != "subAssessmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SubAssessmentName'", id.SubAssessmentName, "subAssessmentValue")
	}
}

func TestFormatScopedSubAssessmentID(t *testing.T) {
	actual := NewScopedSubAssessmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assessmentValue", "subAssessmentValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments/subAssessmentValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedSubAssessmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedSubAssessmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments/subAssessmentValue",
			Expected: &ScopedSubAssessmentId{
				Scope:             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AssessmentName:    "assessmentValue",
				SubAssessmentName: "subAssessmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments/subAssessmentValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedSubAssessmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.AssessmentName != v.Expected.AssessmentName {
			t.Fatalf("Expected %q but got %q for AssessmentName", v.Expected.AssessmentName, actual.AssessmentName)
		}

		if actual.SubAssessmentName != v.Expected.SubAssessmentName {
			t.Fatalf("Expected %q but got %q for SubAssessmentName", v.Expected.SubAssessmentName, actual.SubAssessmentName)
		}

	}
}

func TestParseScopedSubAssessmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedSubAssessmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTvAlUe/sUbAsSeSsMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments/subAssessmentValue",
			Expected: &ScopedSubAssessmentId{
				Scope:             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AssessmentName:    "assessmentValue",
				SubAssessmentName: "subAssessmentValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentValue/subAssessments/subAssessmentValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTvAlUe/sUbAsSeSsMeNtS/sUbAsSeSsMeNtVaLuE",
			Expected: &ScopedSubAssessmentId{
				Scope:             "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				AssessmentName:    "aSsEsSmEnTvAlUe",
				SubAssessmentName: "sUbAsSeSsMeNtVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTvAlUe/sUbAsSeSsMeNtS/sUbAsSeSsMeNtVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedSubAssessmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.AssessmentName != v.Expected.AssessmentName {
			t.Fatalf("Expected %q but got %q for AssessmentName", v.Expected.AssessmentName, actual.AssessmentName)
		}

		if actual.SubAssessmentName != v.Expected.SubAssessmentName {
			t.Fatalf("Expected %q but got %q for SubAssessmentName", v.Expected.SubAssessmentName, actual.SubAssessmentName)
		}

	}
}

func TestSegmentsForScopedSubAssessmentId(t *testing.T) {
	segments := ScopedSubAssessmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedSubAssessmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
