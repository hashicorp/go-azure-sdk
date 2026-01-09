package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRecommendationIdImpactedResourceId{}

func TestNewDirectoryRecommendationIdImpactedResourceID(t *testing.T) {
	id := NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

	if id.RecommendationId != "recommendationId" {
		t.Fatalf("Expected %q but got %q for Segment 'RecommendationId'", id.RecommendationId, "recommendationId")
	}

	if id.ImpactedResourceId != "impactedResourceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ImpactedResourceId'", id.ImpactedResourceId, "impactedResourceId")
	}
}

func TestFormatDirectoryRecommendationIdImpactedResourceID(t *testing.T) {
	actual := NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId").ID()
	expected := "/directory/recommendations/recommendationId/impactedResources/impactedResourceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRecommendationIdImpactedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRecommendationIdImpactedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations/recommendationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations/recommendationId/impactedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/recommendations/recommendationId/impactedResources/impactedResourceId",
			Expected: &DirectoryRecommendationIdImpactedResourceId{
				RecommendationId:   "recommendationId",
				ImpactedResourceId: "impactedResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/recommendations/recommendationId/impactedResources/impactedResourceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRecommendationIdImpactedResourceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RecommendationId != v.Expected.RecommendationId {
			t.Fatalf("Expected %q but got %q for RecommendationId", v.Expected.RecommendationId, actual.RecommendationId)
		}

		if actual.ImpactedResourceId != v.Expected.ImpactedResourceId {
			t.Fatalf("Expected %q but got %q for ImpactedResourceId", v.Expected.ImpactedResourceId, actual.ImpactedResourceId)
		}

	}
}

func TestParseDirectoryRecommendationIdImpactedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRecommendationIdImpactedResourceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations/recommendationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations/recommendationId/impactedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiD/iMpAcTeDrEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/recommendations/recommendationId/impactedResources/impactedResourceId",
			Expected: &DirectoryRecommendationIdImpactedResourceId{
				RecommendationId:   "recommendationId",
				ImpactedResourceId: "impactedResourceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/recommendations/recommendationId/impactedResources/impactedResourceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiD/iMpAcTeDrEsOuRcEs/iMpAcTeDrEsOuRcEiD",
			Expected: &DirectoryRecommendationIdImpactedResourceId{
				RecommendationId:   "rEcOmMeNdAtIoNiD",
				ImpactedResourceId: "iMpAcTeDrEsOuRcEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiD/iMpAcTeDrEsOuRcEs/iMpAcTeDrEsOuRcEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRecommendationIdImpactedResourceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.RecommendationId != v.Expected.RecommendationId {
			t.Fatalf("Expected %q but got %q for RecommendationId", v.Expected.RecommendationId, actual.RecommendationId)
		}

		if actual.ImpactedResourceId != v.Expected.ImpactedResourceId {
			t.Fatalf("Expected %q but got %q for ImpactedResourceId", v.Expected.ImpactedResourceId, actual.ImpactedResourceId)
		}

	}
}

func TestSegmentsForDirectoryRecommendationIdImpactedResourceId(t *testing.T) {
	segments := DirectoryRecommendationIdImpactedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRecommendationIdImpactedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
