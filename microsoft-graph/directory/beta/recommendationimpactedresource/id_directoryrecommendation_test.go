package recommendationimpactedresource

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRecommendationId{}

func TestNewDirectoryRecommendationID(t *testing.T) {
	id := NewDirectoryRecommendationID("recommendationIdValue")

	if id.RecommendationId != "recommendationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RecommendationId'", id.RecommendationId, "recommendationIdValue")
	}
}

func TestFormatDirectoryRecommendationID(t *testing.T) {
	actual := NewDirectoryRecommendationID("recommendationIdValue").ID()
	expected := "/directory/recommendations/recommendationIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDirectoryRecommendationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRecommendationId
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
			// Valid URI
			Input: "/directory/recommendations/recommendationIdValue",
			Expected: &DirectoryRecommendationId{
				RecommendationId: "recommendationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/recommendations/recommendationIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRecommendationID(v.Input)
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

	}
}

func TestParseDirectoryRecommendationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DirectoryRecommendationId
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
			// Valid URI
			Input: "/directory/recommendations/recommendationIdValue",
			Expected: &DirectoryRecommendationId{
				RecommendationId: "recommendationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/directory/recommendations/recommendationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe",
			Expected: &DirectoryRecommendationId{
				RecommendationId: "rEcOmMeNdAtIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dIrEcToRy/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDirectoryRecommendationIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDirectoryRecommendationId(t *testing.T) {
	segments := DirectoryRecommendationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DirectoryRecommendationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
