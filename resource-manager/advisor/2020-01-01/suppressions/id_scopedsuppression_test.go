package suppressions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedSuppressionId{}

func TestNewScopedSuppressionID(t *testing.T) {
	id := NewScopedSuppressionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "recommendationIdValue", "suppressionValue")

	if id.ResourceUri != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceUri'", id.ResourceUri, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.RecommendationId != "recommendationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RecommendationId'", id.RecommendationId, "recommendationIdValue")
	}

	if id.SuppressionName != "suppressionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SuppressionName'", id.SuppressionName, "suppressionValue")
	}
}

func TestFormatScopedSuppressionID(t *testing.T) {
	actual := NewScopedSuppressionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "recommendationIdValue", "suppressionValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions/suppressionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedSuppressionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedSuppressionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions/suppressionValue",
			Expected: &ScopedSuppressionId{
				ResourceUri:      "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RecommendationId: "recommendationIdValue",
				SuppressionName:  "suppressionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions/suppressionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedSuppressionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceUri != v.Expected.ResourceUri {
			t.Fatalf("Expected %q but got %q for ResourceUri", v.Expected.ResourceUri, actual.ResourceUri)
		}

		if actual.RecommendationId != v.Expected.RecommendationId {
			t.Fatalf("Expected %q but got %q for RecommendationId", v.Expected.RecommendationId, actual.RecommendationId)
		}

		if actual.SuppressionName != v.Expected.SuppressionName {
			t.Fatalf("Expected %q but got %q for SuppressionName", v.Expected.SuppressionName, actual.SuppressionName)
		}

	}
}

func TestParseScopedSuppressionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedSuppressionId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aDvIsOr",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aDvIsOr/rEcOmMeNdAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aDvIsOr/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aDvIsOr/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/sUpPrEsSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions/suppressionValue",
			Expected: &ScopedSuppressionId{
				ResourceUri:      "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				RecommendationId: "recommendationIdValue",
				SuppressionName:  "suppressionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.Advisor/recommendations/recommendationIdValue/suppressions/suppressionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aDvIsOr/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/sUpPrEsSiOnS/sUpPrEsSiOnVaLuE",
			Expected: &ScopedSuppressionId{
				ResourceUri:      "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				RecommendationId: "rEcOmMeNdAtIoNiDvAlUe",
				SuppressionName:  "sUpPrEsSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.aDvIsOr/rEcOmMeNdAtIoNs/rEcOmMeNdAtIoNiDvAlUe/sUpPrEsSiOnS/sUpPrEsSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedSuppressionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceUri != v.Expected.ResourceUri {
			t.Fatalf("Expected %q but got %q for ResourceUri", v.Expected.ResourceUri, actual.ResourceUri)
		}

		if actual.RecommendationId != v.Expected.RecommendationId {
			t.Fatalf("Expected %q but got %q for RecommendationId", v.Expected.RecommendationId, actual.RecommendationId)
		}

		if actual.SuppressionName != v.Expected.SuppressionName {
			t.Fatalf("Expected %q but got %q for SuppressionName", v.Expected.SuppressionName, actual.SuppressionName)
		}

	}
}

func TestSegmentsForScopedSuppressionId(t *testing.T) {
	segments := ScopedSuppressionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedSuppressionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
