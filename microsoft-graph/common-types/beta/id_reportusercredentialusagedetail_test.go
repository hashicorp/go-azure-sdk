package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportUserCredentialUsageDetailId{}

func TestNewReportUserCredentialUsageDetailID(t *testing.T) {
	id := NewReportUserCredentialUsageDetailID("userCredentialUsageDetailsIdValue")

	if id.UserCredentialUsageDetailsId != "userCredentialUsageDetailsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserCredentialUsageDetailsId'", id.UserCredentialUsageDetailsId, "userCredentialUsageDetailsIdValue")
	}
}

func TestFormatReportUserCredentialUsageDetailID(t *testing.T) {
	actual := NewReportUserCredentialUsageDetailID("userCredentialUsageDetailsIdValue").ID()
	expected := "/reports/userCredentialUsageDetails/userCredentialUsageDetailsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportUserCredentialUsageDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserCredentialUsageDetailId
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
			Input: "/reports/userCredentialUsageDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userCredentialUsageDetails/userCredentialUsageDetailsIdValue",
			Expected: &ReportUserCredentialUsageDetailId{
				UserCredentialUsageDetailsId: "userCredentialUsageDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userCredentialUsageDetails/userCredentialUsageDetailsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserCredentialUsageDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserCredentialUsageDetailsId != v.Expected.UserCredentialUsageDetailsId {
			t.Fatalf("Expected %q but got %q for UserCredentialUsageDetailsId", v.Expected.UserCredentialUsageDetailsId, actual.UserCredentialUsageDetailsId)
		}

	}
}

func TestParseReportUserCredentialUsageDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportUserCredentialUsageDetailId
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
			Input: "/reports/userCredentialUsageDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRcReDeNtIaLuSaGeDeTaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/userCredentialUsageDetails/userCredentialUsageDetailsIdValue",
			Expected: &ReportUserCredentialUsageDetailId{
				UserCredentialUsageDetailsId: "userCredentialUsageDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/userCredentialUsageDetails/userCredentialUsageDetailsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRcReDeNtIaLuSaGeDeTaIlS/uSeRcReDeNtIaLuSaGeDeTaIlSiDvAlUe",
			Expected: &ReportUserCredentialUsageDetailId{
				UserCredentialUsageDetailsId: "uSeRcReDeNtIaLuSaGeDeTaIlSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/uSeRcReDeNtIaLuSaGeDeTaIlS/uSeRcReDeNtIaLuSaGeDeTaIlSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportUserCredentialUsageDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserCredentialUsageDetailsId != v.Expected.UserCredentialUsageDetailsId {
			t.Fatalf("Expected %q but got %q for UserCredentialUsageDetailsId", v.Expected.UserCredentialUsageDetailsId, actual.UserCredentialUsageDetailsId)
		}

	}
}

func TestSegmentsForReportUserCredentialUsageDetailId(t *testing.T) {
	segments := ReportUserCredentialUsageDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportUserCredentialUsageDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
