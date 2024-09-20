package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportPartnerBillingManifestId{}

func TestNewReportPartnerBillingManifestID(t *testing.T) {
	id := NewReportPartnerBillingManifestID("manifestId")

	if id.ManifestId != "manifestId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManifestId'", id.ManifestId, "manifestId")
	}
}

func TestFormatReportPartnerBillingManifestID(t *testing.T) {
	actual := NewReportPartnerBillingManifestID("manifestId").ID()
	expected := "/reports/partners/billing/manifests/manifestId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportPartnerBillingManifestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportPartnerBillingManifestId
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
			Input: "/reports/partners",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/partners/billing",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/partners/billing/manifests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/partners/billing/manifests/manifestId",
			Expected: &ReportPartnerBillingManifestId{
				ManifestId: "manifestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/partners/billing/manifests/manifestId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportPartnerBillingManifestID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManifestId != v.Expected.ManifestId {
			t.Fatalf("Expected %q but got %q for ManifestId", v.Expected.ManifestId, actual.ManifestId)
		}

	}
}

func TestParseReportPartnerBillingManifestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportPartnerBillingManifestId
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
			Input: "/reports/partners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/partners/billing",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/reports/partners/billing/manifests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg/mAnIfEsTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/partners/billing/manifests/manifestId",
			Expected: &ReportPartnerBillingManifestId{
				ManifestId: "manifestId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/partners/billing/manifests/manifestId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg/mAnIfEsTs/mAnIfEsTiD",
			Expected: &ReportPartnerBillingManifestId{
				ManifestId: "mAnIfEsTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg/mAnIfEsTs/mAnIfEsTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportPartnerBillingManifestIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManifestId != v.Expected.ManifestId {
			t.Fatalf("Expected %q but got %q for ManifestId", v.Expected.ManifestId, actual.ManifestId)
		}

	}
}

func TestSegmentsForReportPartnerBillingManifestId(t *testing.T) {
	segments := ReportPartnerBillingManifestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportPartnerBillingManifestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
