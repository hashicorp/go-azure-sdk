package assessmentsmetadata

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderAssessmentMetadataId{}

func TestNewProviderAssessmentMetadataID(t *testing.T) {
	id := NewProviderAssessmentMetadataID("12345678-1234-9876-4563-123456789012", "assessmentMetadataValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.AssessmentMetadataName != "assessmentMetadataValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssessmentMetadataName'", id.AssessmentMetadataName, "assessmentMetadataValue")
	}
}

func TestFormatProviderAssessmentMetadataID(t *testing.T) {
	actual := NewProviderAssessmentMetadataID("12345678-1234-9876-4563-123456789012", "assessmentMetadataValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderAssessmentMetadataID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderAssessmentMetadataId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue",
			Expected: &ProviderAssessmentMetadataId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				AssessmentMetadataName: "assessmentMetadataValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderAssessmentMetadataID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.AssessmentMetadataName != v.Expected.AssessmentMetadataName {
			t.Fatalf("Expected %q but got %q for AssessmentMetadataName", v.Expected.AssessmentMetadataName, actual.AssessmentMetadataName)
		}

	}
}

func TestParseProviderAssessmentMetadataIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderAssessmentMetadataId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTmEtAdAtA",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue",
			Expected: &ProviderAssessmentMetadataId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				AssessmentMetadataName: "assessmentMetadataValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/assessmentMetadata/assessmentMetadataValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTmEtAdAtA/aSsEsSmEnTmEtAdAtAvAlUe",
			Expected: &ProviderAssessmentMetadataId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				AssessmentMetadataName: "aSsEsSmEnTmEtAdAtAvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/aSsEsSmEnTmEtAdAtA/aSsEsSmEnTmEtAdAtAvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderAssessmentMetadataIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.AssessmentMetadataName != v.Expected.AssessmentMetadataName {
			t.Fatalf("Expected %q but got %q for AssessmentMetadataName", v.Expected.AssessmentMetadataName, actual.AssessmentMetadataName)
		}

	}
}

func TestSegmentsForProviderAssessmentMetadataId(t *testing.T) {
	segments := ProviderAssessmentMetadataId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderAssessmentMetadataId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
