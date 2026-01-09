package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportServicePrincipalSignInActivityId{}

func TestNewReportServicePrincipalSignInActivityID(t *testing.T) {
	id := NewReportServicePrincipalSignInActivityID("servicePrincipalSignInActivityId")

	if id.ServicePrincipalSignInActivityId != "servicePrincipalSignInActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ServicePrincipalSignInActivityId'", id.ServicePrincipalSignInActivityId, "servicePrincipalSignInActivityId")
	}
}

func TestFormatReportServicePrincipalSignInActivityID(t *testing.T) {
	actual := NewReportServicePrincipalSignInActivityID("servicePrincipalSignInActivityId").ID()
	expected := "/reports/servicePrincipalSignInActivities/servicePrincipalSignInActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportServicePrincipalSignInActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportServicePrincipalSignInActivityId
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
			Input: "/reports/servicePrincipalSignInActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/servicePrincipalSignInActivities/servicePrincipalSignInActivityId",
			Expected: &ReportServicePrincipalSignInActivityId{
				ServicePrincipalSignInActivityId: "servicePrincipalSignInActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/servicePrincipalSignInActivities/servicePrincipalSignInActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportServicePrincipalSignInActivityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalSignInActivityId != v.Expected.ServicePrincipalSignInActivityId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalSignInActivityId", v.Expected.ServicePrincipalSignInActivityId, actual.ServicePrincipalSignInActivityId)
		}

	}
}

func TestParseReportServicePrincipalSignInActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportServicePrincipalSignInActivityId
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
			Input: "/reports/servicePrincipalSignInActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/sErViCePrInCiPaLsIgNiNaCtIvItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/servicePrincipalSignInActivities/servicePrincipalSignInActivityId",
			Expected: &ReportServicePrincipalSignInActivityId{
				ServicePrincipalSignInActivityId: "servicePrincipalSignInActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/servicePrincipalSignInActivities/servicePrincipalSignInActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/sErViCePrInCiPaLsIgNiNaCtIvItIeS/sErViCePrInCiPaLsIgNiNaCtIvItYiD",
			Expected: &ReportServicePrincipalSignInActivityId{
				ServicePrincipalSignInActivityId: "sErViCePrInCiPaLsIgNiNaCtIvItYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/sErViCePrInCiPaLsIgNiNaCtIvItIeS/sErViCePrInCiPaLsIgNiNaCtIvItYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportServicePrincipalSignInActivityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ServicePrincipalSignInActivityId != v.Expected.ServicePrincipalSignInActivityId {
			t.Fatalf("Expected %q but got %q for ServicePrincipalSignInActivityId", v.Expected.ServicePrincipalSignInActivityId, actual.ServicePrincipalSignInActivityId)
		}

	}
}

func TestSegmentsForReportServicePrincipalSignInActivityId(t *testing.T) {
	segments := ReportServicePrincipalSignInActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportServicePrincipalSignInActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
