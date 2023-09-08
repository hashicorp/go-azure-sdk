package directoryrecommendationimpactedresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryRecommendationImpactedResourceId{}

func TestNewDirectoryRecommendationImpactedResourceID(t *testing.T) {
	id := NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

	if id.RecommendationId != "recommendationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RecommendationId'", id.RecommendationId, "recommendationIdValue")
	}

	if id.ImpactedResourceId != "impactedResourceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ImpactedResourceId'", id.ImpactedResourceId, "impactedResourceIdValue")
	}
}

func TestFormatDirectoryRecommendationImpactedResourceID(t *testing.T) {
	actual := NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue").ID()
	expected := "/directory/recommendations/recommendationIdValue/impactedResources/impactedResourceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRecommendationImpactedResourceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRecommendationImpactedResourceId
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
			Input: "/directory/recommendations/recommendationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations/recommendationIdValue/impactedResources",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/recommendations/recommendationIdValue/impactedResources/impactedResourceIdValue",
			Expected: &DirectoryRecommendationImpactedResourceId{
				RecommendationId:   "recommendationIdValue",
				ImpactedResourceId: "impactedResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/recommendations/recommendationIdValue/impactedResources/impactedResourceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRecommendationImpactedResourceID(v.Input)
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

func TestParseDirectoryRecommendationImpactedResourceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRecommendationImpactedResourceId
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
			Input: "/directory/recommendations/recommendationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/directory/recommendations/recommendationIdValue/impactedResources",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/iMpAcTeDrEsOuRcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/directory/recommendations/recommendationIdValue/impactedResources/impactedResourceIdValue",
			Expected: &DirectoryRecommendationImpactedResourceId{
				RecommendationId:   "recommendationIdValue",
				ImpactedResourceId: "impactedResourceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/recommendations/recommendationIdValue/impactedResources/impactedResourceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/iMpAcTeDrEsOuRcEs/iMpAcTeDrEsOuRcEiDvAlUe",
			Expected: &DirectoryRecommendationImpactedResourceId{
				RecommendationId:   "rEcOmMeNdAtIoNiDvAlUe",
				ImpactedResourceId: "iMpAcTeDrEsOuRcEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/iMpAcTeDrEsOuRcEs/iMpAcTeDrEsOuRcEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRecommendationImpactedResourceIDInsensitively(v.Input)
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

func TestSegmentsForDirectoryRecommendationImpactedResourceId(t *testing.T) {
	segments := DirectoryRecommendationImpactedResourceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRecommendationImpactedResourceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
