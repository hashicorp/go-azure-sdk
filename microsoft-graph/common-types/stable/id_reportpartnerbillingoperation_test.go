package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportPartnerBillingOperationId{}

func TestNewReportPartnerBillingOperationID(t *testing.T) {
	id := NewReportPartnerBillingOperationID("operationId")

	if id.OperationId != "operationId" {
		t.Fatalf("Expected %q but got %q for Segment 'OperationId'", id.OperationId, "operationId")
	}
}

func TestFormatReportPartnerBillingOperationID(t *testing.T) {
	actual := NewReportPartnerBillingOperationID("operationId").ID()
	expected := "/reports/partners/billing/operations/operationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseReportPartnerBillingOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportPartnerBillingOperationId
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
			Input: "/reports/partners/billing/operations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/partners/billing/operations/operationId",
			Expected: &ReportPartnerBillingOperationId{
				OperationId: "operationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/partners/billing/operations/operationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportPartnerBillingOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestParseReportPartnerBillingOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ReportPartnerBillingOperationId
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
			Input: "/reports/partners/billing/operations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg/oPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/reports/partners/billing/operations/operationId",
			Expected: &ReportPartnerBillingOperationId{
				OperationId: "operationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/reports/partners/billing/operations/operationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg/oPeRaTiOnS/oPeRaTiOnId",
			Expected: &ReportPartnerBillingOperationId{
				OperationId: "oPeRaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/rEpOrTs/pArTnErS/bIlLiNg/oPeRaTiOnS/oPeRaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseReportPartnerBillingOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OperationId != v.Expected.OperationId {
			t.Fatalf("Expected %q but got %q for OperationId", v.Expected.OperationId, actual.OperationId)
		}

	}
}

func TestSegmentsForReportPartnerBillingOperationId(t *testing.T) {
	segments := ReportPartnerBillingOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ReportPartnerBillingOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
