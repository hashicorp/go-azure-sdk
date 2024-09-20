package regulatorycompliance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RegulatoryComplianceStandardId{}

func TestNewRegulatoryComplianceStandardID(t *testing.T) {
	id := NewRegulatoryComplianceStandardID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.RegulatoryComplianceStandardName != "regulatoryComplianceStandardName" {
		t.Fatalf("Expected %q but got %q for Segment 'RegulatoryComplianceStandardName'", id.RegulatoryComplianceStandardName, "regulatoryComplianceStandardName")
	}
}

func TestFormatRegulatoryComplianceStandardID(t *testing.T) {
	actual := NewRegulatoryComplianceStandardID("12345678-1234-9876-4563-123456789012", "regulatoryComplianceStandardName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseRegulatoryComplianceStandardID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RegulatoryComplianceStandardId
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
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardName",
			Expected: &RegulatoryComplianceStandardId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				RegulatoryComplianceStandardName: "regulatoryComplianceStandardName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRegulatoryComplianceStandardID(v.Input)
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

	}
}

func TestParseRegulatoryComplianceStandardIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *RegulatoryComplianceStandardId
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
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardName",
			Expected: &RegulatoryComplianceStandardId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				RegulatoryComplianceStandardName: "regulatoryComplianceStandardName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Security/regulatoryComplianceStandards/regulatoryComplianceStandardName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs/rEgUlAtOrYcOmPlIaNcEsTaNdArDnAmE",
			Expected: &RegulatoryComplianceStandardId{
				SubscriptionId:                   "12345678-1234-9876-4563-123456789012",
				RegulatoryComplianceStandardName: "rEgUlAtOrYcOmPlIaNcEsTaNdArDnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.sEcUrItY/rEgUlAtOrYcOmPlIaNcEsTaNdArDs/rEgUlAtOrYcOmPlIaNcEsTaNdArDnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseRegulatoryComplianceStandardIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForRegulatoryComplianceStandardId(t *testing.T) {
	segments := RegulatoryComplianceStandardId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("RegulatoryComplianceStandardId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
