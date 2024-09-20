package subassessments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedSubAssessmentId{}

func TestNewScopedSubAssessmentID(t *testing.T) {
	id := NewScopedSubAssessmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assessmentName", "subAssessmentName")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.AssessmentName != "assessmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'AssessmentName'", id.AssessmentName, "assessmentName")
	}

	if id.SubAssessmentName != "subAssessmentName" {
		t.Fatalf("Expected %q but got %q for Segment 'SubAssessmentName'", id.SubAssessmentName, "subAssessmentName")
	}
}

func TestFormatScopedSubAssessmentID(t *testing.T) {
	actual := NewScopedSubAssessmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "assessmentName", "subAssessmentName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments/subAssessmentName"
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments/subAssessmentName",
			Expected: &ScopedSubAssessmentId{
				Scope:             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AssessmentName:    "assessmentName",
				SubAssessmentName: "subAssessmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments/subAssessmentName/extra",
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTnAmE/sUbAsSeSsMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments/subAssessmentName",
			Expected: &ScopedSubAssessmentId{
				Scope:             "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				AssessmentName:    "assessmentName",
				SubAssessmentName: "subAssessmentName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Security/assessments/assessmentName/subAssessments/subAssessmentName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTnAmE/sUbAsSeSsMeNtS/sUbAsSeSsMeNtNaMe",
			Expected: &ScopedSubAssessmentId{
				Scope:             "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				AssessmentName:    "aSsEsSmEnTnAmE",
				SubAssessmentName: "sUbAsSeSsMeNtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTs/aSsEsSmEnTnAmE/sUbAsSeSsMeNtS/sUbAsSeSsMeNtNaMe/extra",
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
