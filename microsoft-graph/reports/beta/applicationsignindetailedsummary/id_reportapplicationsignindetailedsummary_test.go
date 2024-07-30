package applicationsignindetailedsummary

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportApplicationSignInDetailedSummaryId{}

func TestNewReportApplicationSignInDetailedSummaryID(t *testing.T) {
	id := NewReportApplicationSignInDetailedSummaryID("applicationSignInDetailedSummaryIdValue")

	if id.ApplicationSignInDetailedSummaryId != "applicationSignInDetailedSummaryIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplicationSignInDetailedSummaryId'", id.ApplicationSignInDetailedSummaryId, "applicationSignInDetailedSummaryIdValue")
	}
}

func TestFormatReportApplicationSignInDetailedSummaryID(t *testing.T) {
	actual := NewReportApplicationSignInDetailedSummaryID("applicationSignInDetailedSummaryIdValue").ID()
	expected := "/reports/applicationSignInDetailedSummary/applicationSignInDetailedSummaryIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportApplicationSignInDetailedSummaryID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportApplicationSignInDetailedSummaryId
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
			Input: "/reports/applicationSignInDetailedSummary",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/applicationSignInDetailedSummary/applicationSignInDetailedSummaryIdValue",
			Expected: &ReportApplicationSignInDetailedSummaryId{
				ApplicationSignInDetailedSummaryId: "applicationSignInDetailedSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/applicationSignInDetailedSummary/applicationSignInDetailedSummaryIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportApplicationSignInDetailedSummaryID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationSignInDetailedSummaryId != v.Expected.ApplicationSignInDetailedSummaryId {
			t.Fatalf("Expected %q but got %q for ApplicationSignInDetailedSummaryId", v.Expected.ApplicationSignInDetailedSummaryId, actual.ApplicationSignInDetailedSummaryId)
		}

	}
}

func TestParseReportApplicationSignInDetailedSummaryIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportApplicationSignInDetailedSummaryId
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
			Input: "/reports/applicationSignInDetailedSummary",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/aPpLiCaTiOnSiGnInDeTaIlEdSuMmArY",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/applicationSignInDetailedSummary/applicationSignInDetailedSummaryIdValue",
			Expected: &ReportApplicationSignInDetailedSummaryId{
				ApplicationSignInDetailedSummaryId: "applicationSignInDetailedSummaryIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/applicationSignInDetailedSummary/applicationSignInDetailedSummaryIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/aPpLiCaTiOnSiGnInDeTaIlEdSuMmArY/aPpLiCaTiOnSiGnInDeTaIlEdSuMmArYiDvAlUe",
			Expected: &ReportApplicationSignInDetailedSummaryId{
				ApplicationSignInDetailedSummaryId: "aPpLiCaTiOnSiGnInDeTaIlEdSuMmArYiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/aPpLiCaTiOnSiGnInDeTaIlEdSuMmArY/aPpLiCaTiOnSiGnInDeTaIlEdSuMmArYiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportApplicationSignInDetailedSummaryIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ApplicationSignInDetailedSummaryId != v.Expected.ApplicationSignInDetailedSummaryId {
			t.Fatalf("Expected %q but got %q for ApplicationSignInDetailedSummaryId", v.Expected.ApplicationSignInDetailedSummaryId, actual.ApplicationSignInDetailedSummaryId)
		}

	}
}

func TestSegmentsForReportApplicationSignInDetailedSummaryId(t *testing.T) {
	segments := ReportApplicationSignInDetailedSummaryId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportApplicationSignInDetailedSummaryId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
