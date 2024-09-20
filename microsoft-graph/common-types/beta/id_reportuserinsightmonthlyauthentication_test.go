package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserInsightMonthlyAuthenticationId{}

func TestNewReportUserInsightMonthlyAuthenticationID(t *testing.T) {
	id := NewReportUserInsightMonthlyAuthenticationID("authenticationsMetricId")

	if id.AuthenticationsMetricId != "authenticationsMetricId" {
		t.Fatalf("Expected %q but got %q for Segment 'AuthenticationsMetricId'", id.AuthenticationsMetricId, "authenticationsMetricId")
	}
}

func TestFormatReportUserInsightMonthlyAuthenticationID(t *testing.T) {
	actual := NewReportUserInsightMonthlyAuthenticationID("authenticationsMetricId").ID()
	expected := "/reports/userInsights/monthly/authentications/authenticationsMetricId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserInsightMonthlyAuthenticationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyAuthenticationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/monthly",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/monthly/authentications",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/authentications/authenticationsMetricId",
			Expected: &ReportUserInsightMonthlyAuthenticationId{
				AuthenticationsMetricId: "authenticationsMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/authentications/authenticationsMetricId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyAuthenticationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationsMetricId != v.Expected.AuthenticationsMetricId {
			t.Fatalf("Expected %q but got %q for AuthenticationsMetricId", v.Expected.AuthenticationsMetricId, actual.AuthenticationsMetricId)
		}

	}
}

func TestParseReportUserInsightMonthlyAuthenticationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserInsightMonthlyAuthenticationId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/monthly",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/userInsights/monthly/authentications",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aUtHeNtIcAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userInsights/monthly/authentications/authenticationsMetricId",
			Expected: &ReportUserInsightMonthlyAuthenticationId{
				AuthenticationsMetricId: "authenticationsMetricId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userInsights/monthly/authentications/authenticationsMetricId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aUtHeNtIcAtIoNs/aUtHeNtIcAtIoNsMeTrIcId",
			Expected: &ReportUserInsightMonthlyAuthenticationId{
				AuthenticationsMetricId: "aUtHeNtIcAtIoNsMeTrIcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRiNsIgHtS/mOnThLy/aUtHeNtIcAtIoNs/aUtHeNtIcAtIoNsMeTrIcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserInsightMonthlyAuthenticationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AuthenticationsMetricId != v.Expected.AuthenticationsMetricId {
			t.Fatalf("Expected %q but got %q for AuthenticationsMetricId", v.Expected.AuthenticationsMetricId, actual.AuthenticationsMetricId)
		}

	}
}

func TestSegmentsForReportUserInsightMonthlyAuthenticationId(t *testing.T) {
	segments := ReportUserInsightMonthlyAuthenticationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserInsightMonthlyAuthenticationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
