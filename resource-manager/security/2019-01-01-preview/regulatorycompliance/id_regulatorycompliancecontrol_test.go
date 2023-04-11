package regulatorycompliance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RegulatoryComplianceControlId{}

func TestNewRegulatoryComplianceControlID(t *testing.T) {
	id := NewRegulatoryComplianceControlID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardValue", "regulatoryComplianceControlValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.RegulatoryComplianceStandardName != "regulatoryComplianceStandardValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RegulatoryComplianceStandardName'", id.RegulatoryComplianceStandardName, "regulatoryComplianceStandardValue")
	}

	if id.RegulatoryComplianceControlName != "regulatoryComplianceControlValue" {
		t.Fatalf("Expected %q but got %q for Segment 'RegulatoryComplianceControlName'", id.RegulatoryComplianceControlName, "regulatoryComplianceControlValue")
	}
}

func TestFormatRegulatoryComplianceControlID(t *testing.T) {
	actual := NewRegulatoryComplianceControlID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardValue", "regulatoryComplianceControlValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls/regulatoryComplianceControlValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRegulatoryComplianceControlID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RegulatoryComplianceControlId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls/regulatoryComplianceControlValue",
			Expected: &RegulatoryComplianceControlId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				RegulatoryComplianceStandardName: "regulatoryComplianceStandardValue",
				RegulatoryComplianceControlName:  "regulatoryComplianceControlValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls/regulatoryComplianceControlValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRegulatoryComplianceControlID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.RegulatoryComplianceStandardName != v.Expected.RegulatoryComplianceStandardName {
			t.Fatalf("Expected %q but got %q for RegulatoryComplianceStandardName", v.Expected.RegulatoryComplianceStandardName, actual.RegulatoryComplianceStandardName)
		}

		if actual.RegulatoryComplianceControlName != v.Expected.RegulatoryComplianceControlName {
			t.Fatalf("Expected %q but got %q for RegulatoryComplianceControlName", v.Expected.RegulatoryComplianceControlName, actual.RegulatoryComplianceControlName)
		}

	}
}

func TestParseRegulatoryComplianceControlIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RegulatoryComplianceControlId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs/rEgUlAtOrYcOmPlIaNcEsTaNdArDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs/rEgUlAtOrYcOmPlIaNcEsTaNdArDvAlUe/rEgUlAtOrYcOmPlIaNcEcOnTrOlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls/regulatoryComplianceControlValue",
			Expected: &RegulatoryComplianceControlId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				RegulatoryComplianceStandardName: "regulatoryComplianceStandardValue",
				RegulatoryComplianceControlName:  "regulatoryComplianceControlValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardValue/regulatoryComplianceControls/regulatoryComplianceControlValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs/rEgUlAtOrYcOmPlIaNcEsTaNdArDvAlUe/rEgUlAtOrYcOmPlIaNcEcOnTrOlS/rEgUlAtOrYcOmPlIaNcEcOnTrOlVaLuE",
			Expected: &RegulatoryComplianceControlId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				RegulatoryComplianceStandardName: "rEgUlAtOrYcOmPlIaNcEsTaNdArDvAlUe",
				RegulatoryComplianceControlName:  "rEgUlAtOrYcOmPlIaNcEcOnTrOlVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs/rEgUlAtOrYcOmPlIaNcEsTaNdArDvAlUe/rEgUlAtOrYcOmPlIaNcEcOnTrOlS/rEgUlAtOrYcOmPlIaNcEcOnTrOlVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRegulatoryComplianceControlIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.RegulatoryComplianceStandardName != v.Expected.RegulatoryComplianceStandardName {
			t.Fatalf("Expected %q but got %q for RegulatoryComplianceStandardName", v.Expected.RegulatoryComplianceStandardName, actual.RegulatoryComplianceStandardName)
		}

		if actual.RegulatoryComplianceControlName != v.Expected.RegulatoryComplianceControlName {
			t.Fatalf("Expected %q but got %q for RegulatoryComplianceControlName", v.Expected.RegulatoryComplianceControlName, actual.RegulatoryComplianceControlName)
		}

	}
}

func TestSegmentsForRegulatoryComplianceControlId(t *testing.T) {
	segments := RegulatoryComplianceControlId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RegulatoryComplianceControlId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
